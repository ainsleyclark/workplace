// Copyright 2020 The Reddico Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package workplace

import (
	"encoding/json"
	"github.com/ainsleyclark/errors"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
)

func TestNew(t *testing.T) {
	tt := map[string]struct {
		input Config
		want  error
	}{
		"Success": {
			Config{Token: "token"},
			nil,
		},
		"Error": {
			Config{Token: ""},
			errors.New("token cannot be nil"),
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			got, err := New(test.input)
			if err != nil {
				if !reflect.DeepEqual(test.want, err) {
					t.Fatalf("expecting: %s got %s", test.want, got)
				}
			}
		})
	}
}

var (
	tx = Transmission{
		Thread:  "thread",
		Message: "hey!",
	}
	lorem = `
Lorem ipsum dolor sit amet, consectetur adipiscing elit. Donec nisi tellus, feugiat vel imperdiet eget, interdum ac est. Mauris semper consectetur arcu, id vehicula tortor venenatis interdum. Fusce mattis a velit quis fringilla. Mauris mattis consectetur mauris, ut accumsan ipsum gravida lobortis. Donec non quam ligula. Pellentesque non dolor iaculis, lobortis arcu sit amet, tempor neque. Maecenas et massa est. Quisque lectus est, iaculis vel pretium nec, suscipit sit amet tellus. Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia curae; Mauris sed rhoncus ligula. Maecenas auctor finibus eros, nec volutpat arcu elementum eu. Aenean tincidunt orci ex. Suspendisse aliquet nibh vitae tellus bibendum, suscipit pellentesque ante pretium. Donec faucibus nunc ipsum, non eleifend tellus mattis in.
Sed vulputate, metus eu viverra dapibus, justo metus placerat arcu, a volutpat risus lorem sed orci. Quisque sollicitudin odio eu sapien fringilla, sed eleifend lectus tincidunt. Pellentesque habitant morbi tristique senectus et netus et malesuada fames ac turpis egestas. Etiam nec quam luctus, tincidunt elit vitae, volutpat arcu. Phasellus eu diam eu lectus congue congue id sit amet ex. Pellentesque tincidunt lectus in est maximus tempor. Phasellus cursus sapien non turpis iaculis aliquet. Praesent venenatis bibendum ultrices. Nulla convallis vestibulum placerat. Quisque tincidunt metus eu sem facilisis, id sodales ex luctus. Nulla imperdiet tortor ut dolor imperdiet, id hendrerit velit mollis. Aenean rhoncus lorem at mi tristique, eget maximus lectus varius. Nunc lacinia sem id eros mattis, vitae pulvinar turpis ultricies. Curabitur ornare dignissim lacus, consequat consectetur ipsum tincidunt faucibus. Integer lobortis eleifend sapien convallis scelerisque.
Mauris a feugiat arcu, tempus vehicula ex. Donec et ultrices tortor, at placerat sapien. In hac habitasse platea dictumst. Fusce feugiat mi et felis elementum luctus. Phasellus efficitur vehicula mi id mollis. Proin vulputate urna in nisi aliquet dignissim. Donec vitae viverra magna, ac auctor elit. Morbi tincidunt elit et laoreet placerat. Phasellus vehicula ornare nunc. Aliquam erat volutpat. Aliquam bibendum nisl eu erat dapibus, at scelerisque mauris lacinia.
Aliquam ut bibendum purus, nec aliquam orci. Sed ultricies ut tellus id feugiat. Proin at mauris eget odio sagittis sagittis. Fusce magna ipsum, suscipit vel euismod quis, pretium eu purus. Cras sit amet volutpat risus. Sed purus lacus, lobortis quis metus et, luctus rutrum felis. Nam interdum non nisi vel pulvinar. Sed tincidunt dolor et justo vehicula, ut lacinia sapien porta. Mauris malesuada felis et porttitor aliquet. Curabitur a sagittis mauris, porttitor lobortis est. Vivamus vulputate vitae massa ut pulvinar.
Etiam mi massa, condimentum ac metus congue, volutpat interdum ipsum. In hac habitasse platea dictumst. Aenean lacinia mattis sem eget aliquam. Integer justo ligula, venenatis vel mi in, suscipit lacinia tortor. Morbi quis porta enim. Curabitur quis enim in tellus tempor fringilla. Donec quis placerat lorem. Quisque sit amet hendrerit purus. Morbi a justo ac dolor dictum hendrerit. Etiam a dolor venenatis lacus vulputate elementum a a dolor. Ut nec ultricies turpis.`
)

func TestClient_Notify(t *testing.T) {
	tt := map[string]struct {
		input      Transmission
		url        func(url string) string
		handler    http.HandlerFunc
		marshaller func(v any) ([]byte, error)
		want       string
	}{
		"Success": {
			tx,
			nil,
			func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
			},
			json.Marshal,
			"",
		},
		"Chunked": {
			Transmission{
				Thread:  "thread",
				Message: lorem,
			},
			nil,
			func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
			},
			json.Marshal,
			"",
		},
		"Marshal Error": {
			tx,
			nil,
			nil,
			func(v any) ([]byte, error) {
				return nil, errors.New("error")
			},
			"error",
		},
		"Request Error": {
			tx,
			func(url string) string {
				return "@#@#$$%$"
			},
			nil,
			json.Marshal,
			"invalid URL escape",
		},
		"Do Error": {
			tx,
			func(url string) string {
				return ""
			},
			nil,
			json.Marshal,
			"unsupported protocol scheme",
		},
		"Bad Status Code": {
			tx,
			nil,
			func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusInternalServerError)
			},
			json.Marshal,
			"invalid request",
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			server := httptest.NewServer(test.handler)
			defer server.Close()

			endpoint := server.URL
			if test.url != nil {
				endpoint = test.url(endpoint)
			}

			c := Client{
				client:     server.Client(),
				baseURL:    endpoint,
				marshaller: test.marshaller,
			}

			got := c.Notify(test.input)
			if got != nil && !strings.Contains(got.Error(), test.want) {
				t.Fatalf("expecting: %s got %s", test.want, got.Error())
			}
		})
	}
}

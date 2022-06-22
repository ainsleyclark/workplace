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
)

func TestClient_Notify(t *testing.T) {
	tt := map[string]struct {
		url        func(url string) string
		handler    http.HandlerFunc
		marshaller func(v any) ([]byte, error)
		want       string
	}{
		"Success": {
			nil,
			func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
			},
			json.Marshal,
			"",
		},
		"Marshal Error": {
			nil,
			nil,
			func(v any) ([]byte, error) {
				return nil, errors.New("error")
			},
			"error",
		},
		"Request Error": {
			func(url string) string {
				return "@#@#$$%$"
			},
			nil,
			json.Marshal,
			"invalid URL escape",
		},
		"Do Error": {
			func(url string) string {
				return ""
			},
			nil,
			json.Marshal,
			"unsupported protocol scheme",
		},
		"Bad Status Code": {
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

			got := c.Notify(tx)
			if got != nil && !strings.Contains(got.Error(), test.want) {
				t.Fatalf("expecting: %s got %s", test.want, got.Error())
			}
		})
	}
}

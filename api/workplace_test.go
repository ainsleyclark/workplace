// Copyright 2020 The Reddico Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package api

//import (
//	"encoding/json"
//	"github.com/ainsleyclark/errors"
//	"github.com/krang-backlink/api/environment"
//	"github.com/stretchr/testify/assert"
//	"net/http"
//	"net/http/httptest"
//	"testing"
//)
//
//func TestNew(t *testing.T) {
//	cfg := environment.Config{}
//	got := New(&cfg)
//	assert.NotNil(t, got)
//}
//
//var (
//	tx = Transmission{
//		Thread:  Threads.Software,
//		Message: "hey!",
//	}
//)
//
//func TestClient_Extract(t *testing.T) {
//	tt := map[string]struct {
//		url        func(url string) string
//		handler    http.HandlerFunc
//		marshaller func(v any) ([]byte, error)
//		want       any
//	}{
//		"Success": {
//			nil,
//			func(w http.ResponseWriter, r *http.Request) {
//				w.WriteHeader(http.StatusOK)
//			},
//			json.Marshal,
//			"Error marshalling Workplace transmission",
//		},
//		"Marshal Error": {
//			nil,
//			nil,
//			func(v any) ([]byte, error) {
//				return nil, errors.New("error")
//			},
//			"Error marshalling Workplace transmission",
//		},
//		"Request Error": {
//			func(url string) string {
//				return "@#@#$$%$"
//			},
//			nil,
//			json.Marshal,
//			"Error creating new request",
//		},
//		"Do Error": {
//			func(url string) string {
//				return ""
//			},
//			nil,
//			json.Marshal,
//			"Error sending Workplace request",
//		},
//		"Bad Status Code": {
//			nil,
//			func(w http.ResponseWriter, r *http.Request) {
//				w.WriteHeader(http.StatusInternalServerError)
//			},
//			json.Marshal,
//			"Error sending Workplace message status",
//		},
//	}
//
//	for name, test := range tt {
//		t.Run(name, func(t *testing.T) {
//			server := httptest.NewServer(test.handler)
//			defer server.Close()
//
//			endpoint := server.URL
//			if test.url != nil {
//				endpoint = test.url(endpoint)
//			}
//
//			c := Client{
//				client: server.Client(),
//				config: environment.Workplace{
//					URL: endpoint,
//				},
//				marshaller: test.marshaller,
//			}
//
//			err := c.Notify(tx)
//			if err != nil {
//				assert.Contains(t, errors.Message(err), test.want)
//				return
//			}
//		})
//	}
//}

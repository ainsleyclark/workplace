// Copyright 2022 Ainsley Clark. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package workplace

import (
	"errors"
	"reflect"
	"testing"
)

func TestConfig_Validate(t *testing.T) {
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
			got := test.input.Validate()
			if !reflect.DeepEqual(test.want, got) {
				t.Fatalf("expecting: %s got %s", test.want, got)
			}
		})
	}
}

// Copyright 2022 Ainsley Clark. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package workplace

import (
	"reflect"
	"testing"
)

func TestChunks(t *testing.T) {
	tt := map[string]struct {
		input string
		size  int
		want  []string
	}{
		"Empty": {
			"",
			10,
			nil,
		},
		"One": {
			"test",
			10,
			[]string{"test"},
		},
		"Simple": {
			"123456789",
			5,
			[]string{"12345", "6789"},
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			got := ChunkMessage(test.input, test.size)
			if !reflect.DeepEqual(test.want, got) {
				t.Fatalf("expecting: %s got %s", test.want, got)
			}
		})
	}
}

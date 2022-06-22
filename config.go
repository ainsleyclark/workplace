// Copyright 2022 Ainsley Clark. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package workplace

import "github.com/ainsleyclark/errors"

type (
	// Config is the configuration passed to create a new
	// Workplace client.
	Config struct {
		Token string
	}
)

// Validate sanity checks the configuration.
func (c *Config) Validate() error {
	if c.Token == "" {
		return errors.New("token cannot be nil")
	}
	return nil
}

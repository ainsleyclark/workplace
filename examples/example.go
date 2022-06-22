// Copyright 2020 The Reddico Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package examples

import "github.com/ainsleyclark/workplace"

func Example() error {
	// Create as new Workplace client.
	wp, err := workplace.New(workplace.Config{Token: "my-token"})
	if err != nil {
		return err
	}

	// Create a new Workplace Transmission that contains
	// the thread ID and message to be sent to the thread.
	tx := workplace.Transmission{
		Thread:  "thread-id",
		Message: "message",
	}

	// Send the transmission to the workplace API.
	err = wp.Notify(tx)
	if err != nil {
		return err
	}

	return nil
}

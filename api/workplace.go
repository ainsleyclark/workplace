// Copyright 2022 Ainsley Clark. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/ainsleyclark/errors"
	"net/http"
	"time"
)

type (
	// Notifier defines the single method interface that
	// transmits messages to a workplace thread.
	Notifier interface {
		// Notify sends a transmission to the corresponding thread with
		// a preformatted message.
		//
		// Returns errors.INTERNAL if the message could not be marshalled,
		// sent or there was an error creating the request.
		Notify(tx Transmission) error
	}
	// Client represents the HTTP client for posting and
	// returning data from the Workplace API.
	Client struct {
		client     *http.Client
		baseURL    string
		token      string
		marshaller func(v any) ([]byte, error)
	}
	// Transmission represents the data needed in order to send
	// a message via the Workplace API.
	Transmission struct {
		// Thread is the corresponding thread ID to send via workplace.
		Thread string `json:"thread" validate:"required" example:"t_111222333"`
		// Message is the formatted text message that will be sent
		// via workplace.
		Message string `json:"message" validate:"required" example:"Hey there!"`
	}
	// transmission is the JSON structure for workplace transmissions.
	transmission struct {
		Recipient recipient `json:"recipient"`
		Message   message   `json:"message"`
	}
	// transmission is the JSON structure for workplace messages.
	message struct {
		Text string `json:"text"`
	}
	// transmission is the JSON structure for workplace recipients.
	recipient struct {
		ThreadKey string `json:"thread_key"`
	}
)

const (
	// Timeout specifies a time limit for requests made by the
	// Workplace Client.
	Timeout = time.Second * 20
)

// New returns a new Workplace Client.
func New(token string) Notifier {
	return &Client{
		client: &http.Client{
			Timeout: Timeout,
		},
		baseURL:    "https://graph.workplace.com/me/messages",
		token:      token,
		marshaller: json.Marshal,
	}
}

func (c *Client) Notify(tx Transmission) error {
	const op = "Workplace.Notify"

	t := transmission{
		Recipient: recipient{ThreadKey: tx.Thread},
		Message:   message{Text: tx.Message},
	}

	buf, err := c.marshaller(t)
	if err != nil {
		return errors.NewInternal(err, "Error marshalling Workplace transmission", op)
	}

	req, err := http.NewRequest(http.MethodPost, c.baseURL, bytes.NewBuffer(buf))
	if err != nil {
		return errors.NewInternal(err, "Error creating new request", op)
	}

	// Add the appropriate auth headers.
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.token))

	resp, err := c.client.Do(req)
	if err != nil {
		return errors.NewInternal(err, "Error sending Workplace request", op)
	} else if resp.StatusCode != http.StatusOK {
		return errors.NewInternal(errors.New("invalid status code"), "Error sending Workplace message status: "+resp.Status, op)
	}
	defer resp.Body.Close() // nolint

	return nil
}

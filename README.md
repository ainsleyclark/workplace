<div align="center">
<img height="250" src="res/logo.svg?logo=v2" alt="Workplace Logo" />

&nbsp;

[![made-with-Go](https://img.shields.io/badge/Made%20with-Go-1f425f.svg)](http://golang.org)
[![GoDoc](https://godoc.org/github.com/ainsleyclark/workplace?status.svg)](https://pkg.go.dev/github.com/ainsleyclark/workplace)
[![Test](https://github.com/ainsleyclark/workplace/actions/workflows/test.yml/badge.svg?branch=master)](https://github.com/ainsleyclark/workplace/actions/workflows/test.yml)
[![codecov](https://codecov.io/gh/ainsleyclark/workplace/branch/master/graph/badge.svg?token=K27L8LS7DA)](https://codecov.io/gh/ainsleyclark/workplace)
[![GoReportCard](https://goreportcard.com/badge/github.com/ainsleyclark/workplace)](https://goreportcard.com/report/github.com/ainsleyclark/workplace)

</div>

# Workplace

An extremely simple Facebook Workplace client for sending transmissions to threads.

## Install

```
go get -u github.com/ainsleyclark/redigo
```

## Quick Start

```go
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
```

## Roadmap

- Add all workplace graph endpoints from [Facebook Endpoints](https://github.com/fbsamples/workplace-platform-samples/blob/main/SampleAPIEndpoints/Postman/Workplace_Graph_Collection.json)

## Contributing

Please feel free to make a pull request if you think something should be added to this package!

## Credits

Shout out to the incredible [Maria Letta](https://github.com/MariaLetta) for her excellent Gopher illustrations.

## Licence

Code Copyright 2022 Ainsley Clark. Code released under the [MIT Licence](LICENSE).

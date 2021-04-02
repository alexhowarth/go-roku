# Roku External Control Protocol Client

[![GoDoc](https://godoc.org/github.com/alexhowarth/go-oku?status.svg)](https://pkg.go.dev/github.com/alexhowarth/go-roku)
[![Go Report Card](https://goreportcard.com/badge/alexhowarth/go-roku)](https://goreportcard.com/report/github.com/alexhowarth/go-roku)

This package provides a Go library for the Roku External Control Protocol
## Installation

~~~~
go get github.com/alexhowarth/go-roku
~~~~

## Usage

```go
// search for a Roku device on the local network
roku.Search()

// build a client
client := config.NewClient("http://192.168.0.21:8060")

// build a client, overriding the default http.Client
httpClient := &http.Client{} // customise as required
client := config.NewClient("http://192.168.0.21:8060", WithHttpClient(httpClient))

// get information about the device
info, err := client.DeviceInfo()

// find the active app
apps, err := client.ActiveApp()

// launch the YouTube app
youtube := "837"
err := client.Launch(youtube)
```

## Command line tool

This package contains the command line tool `roku-cli`

To install and use it:
```bash
$ go install ./roku-cli
$ roku-cli --help
```

There are also pre-built binaries for various architectures [here](https://github.com/alexhowarth/go-roku/releases).

To extract certain data from the output, you might use jq.
## API

Documentation for the API can be found at [https://developer.roku.com](https://developer.roku.com/docs/developer-program/debugging/external-control-api.md)
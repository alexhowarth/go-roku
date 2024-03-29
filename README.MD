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
client := roku.NewClient("http://192.168.0.21:8060")

// build a client, overriding the default http.Client
httpClient := &http.Client{} // customise as required
client := roku.NewClient("http://192.168.0.21:8060", WithHttpClient(httpClient))

// information about the device
info, err := client.DeviceInfo()

// the active app
app, err := client.ActiveApp()

// all apps on the device
apps, err := client.Apps()

// launch the YouTube app with arguments
youtube := "837"
err := client.Launch(youtube, "yk8yvt5lWVc", "")
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
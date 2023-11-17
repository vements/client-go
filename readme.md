[![Build and Test](https://github.com/vements/client-go/actions/workflows/build-test.yaml/badge.svg?event=push)](https://github.com/vements/client-go/actions/workflows/build-test.yaml) [![Go Reference](https://pkg.go.dev/badge/github.com/vements/client-go@v0.0.3.svg)](https://pkg.go.dev/github.com/vements/client-go@v0.0.3) [![GitHub tag](https://img.shields.io/github/tag/vements/client-go?include_prereleases=&sort=semver&color=blue)](https://github.com/vements/client-go/releases/)
[![License](https://img.shields.io/badge/License-MIT-blue)](#license)
[![issues - client-go](https://img.shields.io/github/issues/vements/client-go)](https://github.com/vements/client-go/issues)

## Vements Client Library for Go

The Vements Client Library for Go is a client library for accessing the Vements API from applications written in the Go language. It also includes a command line tool that can be used to interact with the API in scripts or in a terminal.

### Documentation

See the [Go API docs](https://vements.io/docs/clients/golang/) for more information on how to use this library.

### Installation

To install the Vements Client Library for Go, use the following command:

```bash
$ go get github.com/vements/client-go@v0.0.3
```

### Build 

To build the command line tool, use the following command:

```bash
$ go build -o vements github.com/vements/client-go
```

### Usage

The following example shows how to use the Vements Client Library for Go to create a new Vements client, and then use that client to create a new Vements scoreboard.

```go
package main

import (
    "fmt"
    "log"

    vements " github.com/vements/client-go"
)


func init() { 
    var client *vements.Client = vements.NewClient("YOUR_API_KEY")
    var scoreboard *vements.Scoreboard = client.CreateScoreboard(display: "My Scoreboard", rankDir: "desc", public: false)
}
```


### Command Line Tool

The Vements Client Library for C# includes a command line tool that can be used to interact with the Vements API. The command line tool supports all of the same operations as the API.

#### Usage

```bash 
$ vements --help
```

The Go CLI tool supports all of the same operations as the CLI tool in other languages:

* achievement CRUD, list, leaderboard, record progress
* participant CRUD, list, progress, scores
* scoreboard CRUD, list, scoreboard, record score

The above commands all support the following options:

* `--api-key` to specify the API key
* `--verbose` to show verbose output

In addition to resource commands, these common commands are also supported:

* `api-version` to show the API version
* `client-version` to show the client library version

The library and CLI both support the following environment variables:

* `API_KEY` to specify the API key
* `SERVER_TAGS` to specify the tags used to select the server URL
* 
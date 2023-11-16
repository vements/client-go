[![Build and Test](https://github.com/vements/client-go/actions/workflows/build-test.yaml/badge.svg?event=push)](https://github.com/vements/client-go/actions/workflows/build-test.yaml)

## Vements Client Library for Go

The Vements Client Library for Go is a client library for accessing the Vements API from applications written in the Go language. It also includes a command line tool that can be used to interact with the API in scripts or in a terminal.

### Documentation

See the [Go API docs](https://vements.io/docs/clients/golang/) for more information on how to use this library.

### Installation

To install the Vements Client Library for Go, use the following command:

```bash
$ go get github.com/vements/client-libs/go
```

### Build 

To build the command line tool, use the following command:

```bash
$ go build -o vements github.com/vements/client-libs/go/cmd/vements
```

### Usage

The following example shows how to use the Vements Client Library for Go to create a new Vements client, and then use that client to create a new Vements scoreboard.

```go
package main

import (
    "fmt"
    "log"

    vements "github.com/vements/client-libs/go"
)


func init() { 
    var client *vements.Client = vements.NewClient("YOUR_API_KEY")
    var scoreboard *vements.Scoreboard = client.CreateScoreboard(display: "My Scoreboard", rankDir: "desc", public: false)
}
```


### Command Line Tool

The Vements Client Library for C# includes a command line tool that can be used to interact with the Vements API. The command line tool supports all of the same operations as the API.

#### Build 

To build the CLI:

```bash
$ go build -o ./vements-cli
```

The `./vements-cli` path is used to avoid a name collision with the `vements` package.  The file can be renamed to `vements` when moved to a directory in the `PATH`.

#### Usage

```bash 
$ vements-cli --help
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
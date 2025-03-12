# CiderGo

> [!CAUTION]
> IF YOU MAKE REQUESTS TOO FAST, THE API WILL RESPOND "status: ok" WITHOUT ACTUALLY RETURNING THE VALUE, RESULTING IN A CRASH. This will be fixed soon.

A convenient library for interacting with Cider's RPC.

## Introduction

> [!TIP]
> [Cider](https://cider.sh) is a client for Apple Music that works on all major operating systems.

CiderGo is designed for developers looking to build applications that control your local Cider app. This library simplifies interactions with the Cider RPC, offering programmatic control over music playback.

## Current Features

This library currently support all endpoints of /api/v1/playback (except for /queue GET).

## Installation

Install CiderGo via `go get`:

```sh
go get github.com/klosradieschen/cidergo
```

## Getting Started

Make sure that the Cider RPC client is running under Settings -> Connectivity -> Websocket API.

> [!IMPORTANT]
> As of right now, you have to go to "Manage External Application Access to Cider" (right underneath the Websocket API setting) and ***disable API tokens***.

## Usage Example

Here’s how you can get information about the currently playing song:

```go
package main

import (
    "fmt"
    "github.com/klosradieschen/cidergo"
)

func main() {
    song, err := cidergo.CurrentSong()
    if err != nil {
        panic(err)
    }

    fmt.Println(song) // Prints details about the currently playing song
}
```

This example demonstrates how to connect to the Cider client and retrieve information about the current song.

## Documentation

The documentation will be added very soon. For now, you can refer to the API documentation, as seen below.

## API Reference

The documentation for Cider's API can be found [here](https://cider.sh/docs/client/rpc)

## Contributing

We welcome all contributions!

## License

CiderGo is licensed under the MIT License. See the [LICENSE](LICENSE) file for more information.

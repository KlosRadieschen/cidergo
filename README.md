# cidergo

> [!WARNING]
> Status: Work in Progress (doesn't work at all as of right now)

This project is under active development and aims to provide a convenient interface for interacting with Cider's RPC/API.

## Introduction

> [!TIP]
> [Cider](https://cider.sh) is a client for Apple Music that works on all major operating systems.

CiderGo is designed for developers looking to build applications that control your local Cider app. This library acts as an RPC/API wrapper around the Cider Apple Music client, offering programmatic control over music playback.

## Current Features

- **Media Information Retrieval**: Get details about currently playing songs.

yeah, that's it

## Installation

Install cidergo via `go get`:

```sh
go get github.com/klosradieschen/cidergo
```

Ensure you have Go installed and properly configured on your system before proceeding with the installation. 

## Getting Started

Make sure that the Cider RPC client is running under setting -> Connectivity -> Websocket API. As of right now, you have to go to "Manage External Application Access to Cider" (right underneath the Websocket API setting) and disable tokens.

## Usage Example

Here’s how you can get information about the currently playing song:

```go
package main

import (
    "fmt"
    "github.com/klosradieschen/cidergo"
)

func main() {
    remote, err := cidergo.Connect("") // Enter a custom port or use "" for the default port (10767)
    if err != nil {
        panic(err)
    }

    song, err := remote.CurrentSong()
    if err != nil {
        panic(err)
    }

    fmt.Println(song) // Prints details about the currently playing song
}
```

This example demonstrates how to connect to the Cider client and retrieve information about the current song. You can specify a different port or use an empty string for the default.

## API Reference

The documentation for Cider's API can be found [here](https://cider.sh/docs/client/rpc)

## Contributing

We welcome contributions!

## License

CiderGo is licensed under the MIT License. See the [LICENSE](LICENSE) file for more information.

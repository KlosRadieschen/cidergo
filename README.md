# CiderGo

A convenient library for interacting with Cider's RPC.

## Introduction

> [!TIP]
> [Cider](https://cider.sh) is a client for Apple Music that works on all major operating systems.

CiderGo is designed for developers looking to build applications that control your local Cider app. This library simplifies interactions with the Cider RPC, offering easy control over music playback.

## Current Features

This library currently support all functional endpoints.

## Installation

Install CiderGo via `go get`:

```sh
go get github.com/klosradieschen/cidergo
```

## Getting Started

Make sure that the Cider RPC client is running under Settings -> Connectivity -> Websocket API.

> [!IMPORTANT]
> The library will not work when the API token isn't properly configured

Right under Websocket API option, you can see "Manage External Application Access to Cider". There, you can either create a new token for your application or disable tokens altogether, in which case the functions will work without using SetToken.

### Usage Example

Here’s a basic example that retrieves the name of the currently playing song:

```go
package main

import (
    "fmt"
    "github.com/klosradieschen/cidergo"
)

func main() {
	cidergo.SetToken("your-token") // Remove this line if you disabled tokens
	
    song, err := cidergo.CurrentSong()
    if err != nil {
        panic(err)
    }

    fmt.Println(song.Name) // Prints the name of the currently playing song
}
```

> [!WARNING]
> Some functions like SetShuffle do nothing in some cases, like when you are listening to a station. There is no error handling for it because Cider's API responses for these functions are always the same.

## Documentation

The Godoc documentation can be found on https://pkg.go.dev/github.com/klosradieschen/cidergo

## API Reference

The documentation for Cider's API can be found [here](https://cider.sh/docs/client/rpc)

## Contributing

We welcome all contributions!

## License

CiderGo is licensed under the MIT License. See the [LICENSE](LICENSE) file for more information.

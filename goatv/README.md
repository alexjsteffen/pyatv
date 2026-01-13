# goatv - Go Library for Apple TV and AirPlay

A Go port of the [pyatv](https://github.com/postlund/pyatv) Python library for interacting with Apple TV and AirPlay devices.

## Features

- Device discovery via mDNS/Zeroconf
- Remote control commands (play, pause, navigate, etc.)
- Metadata retrieval
- Power management
- Volume control
- App management
- Support for multiple protocols (MRP, DMAP, AirPlay, RAOP, Companion)

## Installation

```bash
go get github.com/alexjsteffen/goatv
```

## Quick Start

### Scanning for Devices

```go
package main

import (
    "context"
    "fmt"
    "log"

    "github.com/alexjsteffen/goatv/pkg/pyatv"
)

func main() {
    ctx := context.Background()
    
    // Scan for devices
    devices, err := pyatv.Scan(ctx, pyatv.DefaultScanOptions())
    if err != nil {
        log.Fatal(err)
    }

    for _, device := range devices {
        fmt.Printf("Found: %s at %s\n", device.Name, device.Address)
    }
}
```

### Connecting to a Device

```go
package main

import (
    "context"
    "fmt"
    "log"

    "github.com/alexjsteffen/goatv/pkg/pyatv"
)

func main() {
    ctx := context.Background()
    
    // Scan for devices
    devices, err := pyatv.Scan(ctx, pyatv.DefaultScanOptions())
    if err != nil {
        log.Fatal(err)
    }

    if len(devices) == 0 {
        log.Fatal("No devices found")
    }

    // Connect to first device
    atv, err := pyatv.Connect(ctx, devices[0], pyatv.ConnectOptions{})
    if err != nil {
        log.Fatal(err)
    }
    defer atv.Close()

    // Get what's playing
    playing, err := atv.Metadata().Playing(ctx)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Now playing: %s by %s\n", playing.Title, playing.Artist)
}
```

### Remote Control

```go
// Pause playback
err := atv.RemoteControl().Pause(ctx)

// Play
err := atv.RemoteControl().Play(ctx)

// Navigate
err := atv.RemoteControl().Up(ctx, pyatv.InputActionSingleTap)
err := atv.RemoteControl().Down(ctx, pyatv.InputActionSingleTap)
err := atv.RemoteControl().Left(ctx, pyatv.InputActionSingleTap)
err := atv.RemoteControl().Right(ctx, pyatv.InputActionSingleTap)
err := atv.RemoteControl().Select(ctx, pyatv.InputActionSingleTap)
err := atv.RemoteControl().Menu(ctx, pyatv.InputActionSingleTap)
```

## CLI Tool

The `atvremote` command-line tool provides quick access to device functionality:

```bash
# Build the CLI tool
cd cmd/atvremote
go build

# Scan for devices
./atvremote scan

# Send commands (uses first found device)
./atvremote playing
./atvremote pause
./atvremote play

# Target specific device
./atvremote -s 192.168.1.100 playing
./atvremote -n "Living Room" pause
```

## Supported Protocols

| Protocol | Description |
|----------|-------------|
| MRP | Media Remote Protocol (tvOS 13+) |
| DMAP | Digital Media Access Protocol (legacy) |
| AirPlay | Audio/video streaming |
| RAOP | Remote Audio Output Protocol |
| Companion | Companion Link protocol |

## Project Structure

```
goatv/
├── cmd/
│   └── atvremote/      # CLI tool
├── pkg/
│   └── pyatv/          # Core library
│       ├── const.go    # Constants and enums
│       ├── errors.go   # Error definitions
│       ├── interfaces.go # Interface definitions
│       ├── scan.go     # High-level scan/connect functions
│       ├── scanner.go  # mDNS scanner implementation
│       ├── connection.go # Device connection implementation
│       └── pairing.go  # Pairing handler
└── go.mod
```

## Status

This is a Go port of pyatv. Currently implemented:

- [x] Core type definitions and constants
- [x] mDNS device discovery
- [x] Basic connection framework
- [x] Interface definitions
- [x] CLI tool structure
- [ ] MRP protocol implementation
- [ ] DMAP protocol implementation  
- [ ] AirPlay streaming
- [ ] RAOP streaming
- [ ] Companion protocol
- [ ] Pairing implementation
- [ ] Push updates

## Contributing

Contributions are welcome! Please see the main pyatv repository for protocol documentation.

## License

MIT License - see LICENSE file.

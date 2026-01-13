package pyatv

import (
	"context"
	"fmt"
	"time"
)

// ScanOptions contains options for scanning.
type ScanOptions struct {
	Timeout    time.Duration
	Identifier string
	Protocol   *Protocol
	Hosts      []string
	Storage    Storage
}

// DefaultScanOptions returns default scan options.
func DefaultScanOptions() ScanOptions {
	return ScanOptions{
		Timeout: 5 * time.Second,
	}
}

// Scan scans for Apple TVs on the network and returns their configurations.
func Scan(ctx context.Context, opts ScanOptions) ([]*Config, error) {
	if opts.Timeout == 0 {
		opts.Timeout = 5 * time.Second
	}

	scanner := NewScanner(opts)
	return scanner.Discover(ctx)
}

// ConnectOptions contains options for connecting.
type ConnectOptions struct {
	Protocol *Protocol
	Storage  Storage
}

// Connect connects to a device based on a configuration.
func Connect(ctx context.Context, config *Config, opts ConnectOptions) (AppleTV, error) {
	if len(config.Services) == 0 {
		return nil, ErrNoService
	}

	if config.Identifier == "" {
		return nil, ErrDeviceIDMissing
	}

	// Create a new AppleTV connection
	atv := NewAppleTVConnection(config, opts)
	if err := atv.Connect(ctx); err != nil {
		return nil, err
	}

	return atv, nil
}

// PairOptions contains options for pairing.
type PairOptions struct {
	Storage Storage
}

// Pair initiates pairing with a device for a specific protocol.
func Pair(ctx context.Context, config *Config, protocol Protocol, opts PairOptions) (PairingHandler, error) {
	service := config.GetService(protocol)
	if service == nil {
		return nil, fmt.Errorf("%w: no service available for %s", ErrNoService, protocol)
	}

	handler := NewPairingHandler(config, service, protocol, opts)
	return handler, nil
}

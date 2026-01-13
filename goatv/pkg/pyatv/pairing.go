package pyatv

import (
	"context"
)

// DefaultPairingHandler provides a default pairing implementation.
type DefaultPairingHandler struct {
	config   *Config
	service  *Service
	protocol Protocol
	opts     PairOptions
	pin      string
	paired   bool
}

// NewPairingHandler creates a new pairing handler.
func NewPairingHandler(config *Config, service *Service, protocol Protocol, opts PairOptions) *DefaultPairingHandler {
	return &DefaultPairingHandler{
		config:   config,
		service:  service,
		protocol: protocol,
		opts:     opts,
	}
}

// Service returns the service being paired.
func (p *DefaultPairingHandler) Service() *Service {
	return p.service
}

// Close releases resources.
func (p *DefaultPairingHandler) Close() error {
	return nil
}

// Pin sets the PIN code for pairing.
func (p *DefaultPairingHandler) Pin(pin string) {
	p.pin = pin
}

// DeviceProvidesPin returns true if the device provides the PIN.
func (p *DefaultPairingHandler) DeviceProvidesPin() bool {
	switch p.protocol {
	case ProtocolMRP:
		return true
	case ProtocolAirPlay:
		return true
	case ProtocolCompanion:
		return true
	default:
		return false
	}
}

// HasPaired returns true if pairing was successful.
func (p *DefaultPairingHandler) HasPaired() bool {
	return p.paired
}

// Begin starts the pairing process.
func (p *DefaultPairingHandler) Begin(ctx context.Context) error {
	// TODO: Implement actual pairing logic for each protocol
	return ErrNotSupported
}

// Finish completes the pairing process.
func (p *DefaultPairingHandler) Finish(ctx context.Context) error {
	// TODO: Implement actual pairing finish logic
	if p.pin == "" {
		return ErrNoCredentials
	}

	// TODO: Verify PIN and complete pairing
	p.paired = true
	return nil
}

package pyatv

import (
	"net"
	"testing"
)

func TestConfigGetService(t *testing.T) {
	config := &Config{
		Services: []*Service{
			{Protocol: ProtocolMRP, Port: 49152},
			{Protocol: ProtocolAirPlay, Port: 7000},
		},
	}

	// Test finding existing service
	mrp := config.GetService(ProtocolMRP)
	if mrp == nil {
		t.Error("Expected to find MRP service")
	}
	if mrp.Port != 49152 {
		t.Errorf("Expected port 49152, got %d", mrp.Port)
	}

	// Test finding non-existent service
	dmap := config.GetService(ProtocolDMAP)
	if dmap != nil {
		t.Error("Expected nil for non-existent DMAP service")
	}
}

func TestConfigAddService(t *testing.T) {
	config := &Config{
		Services: []*Service{},
	}

	// Add new service
	svc1 := &Service{
		Protocol:   ProtocolMRP,
		Port:       49152,
		Properties: map[string]string{"key1": "value1"},
	}
	config.AddService(svc1)

	if len(config.Services) != 1 {
		t.Errorf("Expected 1 service, got %d", len(config.Services))
	}

	// Merge service with same protocol
	svc2 := &Service{
		Protocol:    ProtocolMRP,
		Port:        49152,
		Credentials: "creds",
		Properties:  map[string]string{"key2": "value2"},
	}
	config.AddService(svc2)

	if len(config.Services) != 1 {
		t.Errorf("Expected 1 service after merge, got %d", len(config.Services))
	}

	mrp := config.GetService(ProtocolMRP)
	if mrp.Credentials != "creds" {
		t.Errorf("Expected credentials to be updated")
	}
	if mrp.Properties["key1"] != "value1" || mrp.Properties["key2"] != "value2" {
		t.Errorf("Expected properties to be merged")
	}

	// Add different protocol
	svc3 := &Service{
		Protocol:   ProtocolAirPlay,
		Port:       7000,
		Properties: map[string]string{},
	}
	config.AddService(svc3)

	if len(config.Services) != 2 {
		t.Errorf("Expected 2 services, got %d", len(config.Services))
	}
}

func TestConfigReady(t *testing.T) {
	// Not ready - no services
	config1 := &Config{
		Services: []*Service{},
	}
	if config1.Ready() {
		t.Error("Expected config to not be ready with no services")
	}

	// Not ready - no identifier
	config2 := &Config{
		Services: []*Service{
			{Protocol: ProtocolMRP, Identifier: ""},
		},
	}
	if config2.Ready() {
		t.Error("Expected config to not be ready with empty identifier")
	}

	// Ready
	config3 := &Config{
		Services: []*Service{
			{Protocol: ProtocolMRP, Identifier: "abc123"},
		},
	}
	if !config3.Ready() {
		t.Error("Expected config to be ready")
	}
}

func TestConfigAllIdentifiers(t *testing.T) {
	config := &Config{
		Services: []*Service{
			{Protocol: ProtocolMRP, Identifier: "id1"},
			{Protocol: ProtocolAirPlay, Identifier: "id2"},
			{Protocol: ProtocolDMAP, Identifier: ""}, // Empty, should not be included
		},
	}

	ids := config.AllIdentifiers()
	if len(ids) != 2 {
		t.Errorf("Expected 2 identifiers, got %d", len(ids))
	}

	found := map[string]bool{"id1": false, "id2": false}
	for _, id := range ids {
		found[id] = true
	}

	if !found["id1"] || !found["id2"] {
		t.Error("Missing expected identifiers")
	}
}

func TestNewAppleTVConnection(t *testing.T) {
	config := &Config{
		Address:    net.ParseIP("192.168.1.100"),
		Name:       "Test Apple TV",
		Identifier: "test-id",
		Services: []*Service{
			{Protocol: ProtocolMRP, Port: 49152, Identifier: "mrp-id"},
		},
		DeviceInfo: &DeviceInfo{
			Model:   DeviceModelGen4K,
			Version: "15.0",
		},
	}

	atv := NewAppleTVConnection(config, ConnectOptions{})

	if atv == nil {
		t.Fatal("Expected non-nil connection")
	}

	if atv.DeviceInfo().Model != DeviceModelGen4K {
		t.Errorf("Expected Gen4K model, got %v", atv.DeviceInfo().Model)
	}

	if atv.RemoteControl() == nil {
		t.Error("Expected non-nil RemoteControl")
	}

	if atv.Metadata() == nil {
		t.Error("Expected non-nil Metadata")
	}

	if atv.Power() == nil {
		t.Error("Expected non-nil Power")
	}
}

func TestHTTPError(t *testing.T) {
	err := NewHTTPError("Not Found", 404)

	if err.Error() != "Not Found" {
		t.Errorf("Expected 'Not Found', got '%s'", err.Error())
	}

	if err.StatusCode != 404 {
		t.Errorf("Expected status code 404, got %d", err.StatusCode)
	}
}

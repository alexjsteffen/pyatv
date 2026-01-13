package pyatv

import (
	"context"
	"fmt"
	"net"
	"strings"
	"sync"
	"time"

	"github.com/hashicorp/mdns"
)

// ServiceInfo contains information about a discovered service.
type ServiceInfo struct {
	ServiceType string
	Protocol    Protocol
	Port        int
	Host        string
	IP          net.IP
	TxtRecords  map[string]string
}

// Scanner discovers Apple TV devices on the network.
type Scanner struct {
	opts    ScanOptions
	devices map[string]*Config
	mu      sync.Mutex
}

// NewScanner creates a new scanner with the given options.
func NewScanner(opts ScanOptions) *Scanner {
	return &Scanner{
		opts:    opts,
		devices: make(map[string]*Config),
	}
}

// Service types used for discovery.
const (
	ServiceTypeMRP       = "_mediaremotetv._tcp"
	ServiceTypeDMAP      = "_dacp._tcp"
	ServiceTypeAirPlay   = "_airplay._tcp"
	ServiceTypeRAOP      = "_raop._tcp"
	ServiceTypeCompanion = "_companion-link._tcp"
)

// Discover discovers devices on the network.
func (s *Scanner) Discover(ctx context.Context) ([]*Config, error) {
	timeout := s.opts.Timeout
	if timeout == 0 {
		timeout = 5 * time.Second
	}

	// Create context with timeout
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	// Channel to receive discovered entries
	entriesCh := make(chan *mdns.ServiceEntry, 100)

	var wg sync.WaitGroup

	// Start goroutine to process entries
	wg.Add(1)
	go func() {
		defer wg.Done()
		for entry := range entriesCh {
			s.handleEntry(entry)
		}
	}()

	// Scan for all service types
	serviceTypes := []struct {
		Type     string
		Protocol Protocol
	}{
		{ServiceTypeMRP, ProtocolMRP},
		{ServiceTypeDMAP, ProtocolDMAP},
		{ServiceTypeAirPlay, ProtocolAirPlay},
		{ServiceTypeRAOP, ProtocolRAOP},
		{ServiceTypeCompanion, ProtocolCompanion},
	}

	for _, st := range serviceTypes {
		// Skip if a specific protocol is requested and this isn't it
		if s.opts.Protocol != nil && *s.opts.Protocol != st.Protocol {
			continue
		}

		params := mdns.DefaultParams(st.Type)
		params.Entries = entriesCh
		params.Timeout = timeout
		params.WantUnicastResponse = true

		// If specific hosts are requested, do unicast queries
		if len(s.opts.Hosts) > 0 {
			for _, host := range s.opts.Hosts {
				ip := net.ParseIP(host)
				if ip != nil {
					// Unicast query to specific host
					params.WantUnicastResponse = true
				}
			}
		}

		err := mdns.Query(params)
		if err != nil && !strings.Contains(err.Error(), "no such host") {
			// Log error but continue scanning other services
			fmt.Printf("Warning: failed to query %s: %v\n", st.Type, err)
		}
	}

	// Close channel and wait for processing to complete
	close(entriesCh)
	wg.Wait()

	// Filter by identifier if specified
	var result []*Config
	for _, config := range s.devices {
		if !config.Ready() {
			continue
		}

		if s.opts.Identifier != "" {
			found := false
			for _, id := range config.AllIdentifiers() {
				if id == s.opts.Identifier {
					found = true
					break
				}
			}
			if !found {
				continue
			}
		}

		result = append(result, config)
	}

	return result, nil
}

// handleEntry processes a discovered mDNS service entry.
func (s *Scanner) handleEntry(entry *mdns.ServiceEntry) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Extract protocol from service name
	protocol := s.getProtocolFromService(entry.Name)
	if protocol == 0 {
		return
	}

	// Parse TXT records
	txtRecords := make(map[string]string)
	for _, txt := range entry.InfoFields {
		parts := strings.SplitN(txt, "=", 2)
		if len(parts) == 2 {
			txtRecords[parts[0]] = parts[1]
		} else if len(parts) == 1 {
			txtRecords[parts[0]] = ""
		}
	}

	// Get or create device config
	deviceKey := s.getDeviceKey(entry, txtRecords)
	config, exists := s.devices[deviceKey]
	if !exists {
		config = &Config{
			Address:    entry.AddrV4,
			Name:       s.extractName(entry.Name),
			Services:   make([]*Service, 0),
			Properties: make(map[string]map[string]string),
			DeviceInfo: &DeviceInfo{},
		}
		if entry.AddrV4 == nil && entry.AddrV6 != nil {
			config.Address = entry.AddrV6
		}
		s.devices[deviceKey] = config
	}

	// Update device name if not set
	if config.Name == "" {
		config.Name = s.extractName(entry.Name)
	}

	// Extract device info from properties
	s.updateDeviceInfo(config, txtRecords, protocol)

	// Create service
	service := &Service{
		Identifier: s.extractIdentifier(txtRecords, protocol),
		Protocol:   protocol,
		Port:       entry.Port,
		Properties: txtRecords,
		Enabled:    true,
		Pairing:    s.getPairingRequirement(txtRecords, protocol),
	}

	// Add or update service
	config.AddService(service)

	// Update config identifier
	if config.Identifier == "" && service.Identifier != "" {
		config.Identifier = service.Identifier
	}

	// Store properties for this service type
	config.Properties[s.getServiceTypeForProtocol(protocol)] = txtRecords
}

func (s *Scanner) getProtocolFromService(name string) Protocol {
	switch {
	case strings.Contains(name, ServiceTypeMRP):
		return ProtocolMRP
	case strings.Contains(name, ServiceTypeDMAP):
		return ProtocolDMAP
	case strings.Contains(name, ServiceTypeAirPlay):
		return ProtocolAirPlay
	case strings.Contains(name, ServiceTypeRAOP):
		return ProtocolRAOP
	case strings.Contains(name, ServiceTypeCompanion):
		return ProtocolCompanion
	default:
		return 0
	}
}

func (s *Scanner) getServiceTypeForProtocol(protocol Protocol) string {
	switch protocol {
	case ProtocolMRP:
		return ServiceTypeMRP
	case ProtocolDMAP:
		return ServiceTypeDMAP
	case ProtocolAirPlay:
		return ServiceTypeAirPlay
	case ProtocolRAOP:
		return ServiceTypeRAOP
	case ProtocolCompanion:
		return ServiceTypeCompanion
	default:
		return ""
	}
}

func (s *Scanner) getDeviceKey(entry *mdns.ServiceEntry, txtRecords map[string]string) string {
	// Try to use device ID from TXT records
	if deviceID, ok := txtRecords["deviceid"]; ok {
		return deviceID
	}
	if deviceID, ok := txtRecords["DeviceID"]; ok {
		return deviceID
	}

	// Fall back to IP address
	if entry.AddrV4 != nil {
		return entry.AddrV4.String()
	}
	if entry.AddrV6 != nil {
		return entry.AddrV6.String()
	}

	return entry.Name
}

func (s *Scanner) extractName(serviceName string) string {
	// Service name format: "DeviceName._servicetype._tcp.local."
	parts := strings.Split(serviceName, "._")
	if len(parts) > 0 {
		return parts[0]
	}
	return serviceName
}

func (s *Scanner) extractIdentifier(txtRecords map[string]string, protocol Protocol) string {
	// Different protocols use different identifier fields
	switch protocol {
	case ProtocolMRP:
		if id, ok := txtRecords["UniqueIdentifier"]; ok {
			return id
		}
	case ProtocolDMAP:
		if id, ok := txtRecords["HSGID"]; ok {
			return id
		}
	case ProtocolAirPlay:
		if id, ok := txtRecords["deviceid"]; ok {
			return id
		}
	case ProtocolRAOP:
		if id, ok := txtRecords["deviceid"]; ok {
			return id
		}
	case ProtocolCompanion:
		if id, ok := txtRecords["rpHA"]; ok {
			return id
		}
	}

	// Fallback identifiers
	if id, ok := txtRecords["deviceid"]; ok {
		return id
	}
	if id, ok := txtRecords["DeviceID"]; ok {
		return id
	}

	return ""
}

func (s *Scanner) updateDeviceInfo(config *Config, txtRecords map[string]string, protocol Protocol) {
	// Extract model info
	if model, ok := txtRecords["model"]; ok {
		config.DeviceInfo.RawModel = model
		config.DeviceInfo.Model = s.parseModel(model)
	}

	// Extract OS version
	if osv, ok := txtRecords["osvers"]; ok {
		config.DeviceInfo.Version = osv
		config.DeviceInfo.OperatingSystem = OperatingSystemTvOS
	}

	// Extract build number
	if build, ok := txtRecords["srcvers"]; ok {
		config.DeviceInfo.BuildNumber = build
	}

	// Extract MAC address
	if mac, ok := txtRecords["deviceid"]; ok {
		if strings.Contains(mac, ":") {
			config.DeviceInfo.MAC = mac
		}
	}
}

func (s *Scanner) parseModel(model string) DeviceModel {
	model = strings.ToLower(model)

	switch {
	case strings.Contains(model, "appletv3,1"):
		return DeviceModelGen3
	case strings.Contains(model, "appletv3,2"):
		return DeviceModelGen3
	case strings.Contains(model, "appletv5,3"):
		return DeviceModelGen4
	case strings.Contains(model, "appletv6,2"):
		return DeviceModelGen4K
	case strings.Contains(model, "appletv11,1"):
		return DeviceModelAppleTV4KGen2
	case strings.Contains(model, "appletv14,1"):
		return DeviceModelAppleTV4KGen3
	case strings.Contains(model, "homepod"):
		if strings.Contains(model, "mini") {
			return DeviceModelHomePodMini
		}
		return DeviceModelHomePod
	case strings.Contains(model, "airport"):
		return DeviceModelAirPortExpress
	default:
		return DeviceModelUnknown
	}
}

func (s *Scanner) getPairingRequirement(txtRecords map[string]string, protocol Protocol) PairingRequirement {
	switch protocol {
	case ProtocolMRP:
		return PairingRequirementMandatory
	case ProtocolAirPlay:
		if _, ok := txtRecords["sf"]; ok {
			return PairingRequirementOptional
		}
		return PairingRequirementNotNeeded
	case ProtocolCompanion:
		return PairingRequirementMandatory
	default:
		return PairingRequirementNotNeeded
	}
}

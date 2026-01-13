package pyatv

import (
	"context"
	"io"
	"net"
)

// ArtworkInfo represents artwork information.
type ArtworkInfo struct {
	Bytes    []byte
	MimeType string
	Width    int
	Height   int
}

// MediaMetadata represents metadata for media.
type MediaMetadata struct {
	Title    string
	Artist   string
	Album    string
	Artwork  []byte  // Raw JPEG data
	Duration float64 // Duration in seconds
}

// FeatureInfo represents feature state and options.
type FeatureInfo struct {
	State   FeatureState
	Options map[string]interface{}
}

// App represents information about an app.
type App struct {
	Name       string
	Identifier string
}

// UserAccount represents information about a user account.
type UserAccount struct {
	Name       string
	Identifier string
}

// OutputDevice represents information about an output device.
type OutputDevice struct {
	Name       string
	Identifier string
}

// Playing represents what is currently playing.
type Playing struct {
	MediaType              MediaType
	DeviceState            DeviceState
	Title                  string
	Artist                 string
	Album                  string
	Genre                  string
	TotalTime              *int // seconds, nil if unknown
	Position               *int // seconds, nil if unknown
	Shuffle                *ShuffleState
	Repeat                 *RepeatState
	Hash                   string
	SeriesName             string
	SeasonNumber           *int
	EpisodeNumber          *int
	ContentIdentifier      string
	ITunesStoreIdentifier  *int
}

// Service represents a protocol service.
type Service struct {
	Identifier       string
	Protocol         Protocol
	Port             int
	Properties       map[string]string
	Credentials      string
	Password         string
	Enabled          bool
	RequiresPassword bool
	Pairing          PairingRequirement
}

// DeviceInfo represents general device information.
type DeviceInfo struct {
	OperatingSystem OperatingSystem
	Version         string
	BuildNumber     string
	Model           DeviceModel
	RawModel        string
	MAC             string
	OutputDeviceID  string
}

// Config represents a device configuration.
type Config struct {
	Address    net.IP
	Name       string
	DeepSleep  bool
	Services   []*Service
	DeviceInfo *DeviceInfo
	Properties map[string]map[string]string
	Identifier string
}

// GetService looks up a service based on protocol.
func (c *Config) GetService(protocol Protocol) *Service {
	for _, s := range c.Services {
		if s.Protocol == protocol {
			return s
		}
	}
	return nil
}

// AddService adds a new service or merges with existing one.
func (c *Config) AddService(service *Service) {
	for i, s := range c.Services {
		if s.Protocol == service.Protocol {
			// Merge services
			if service.Credentials != "" {
				c.Services[i].Credentials = service.Credentials
			}
			if service.Password != "" {
				c.Services[i].Password = service.Password
			}
			for k, v := range service.Properties {
				c.Services[i].Properties[k] = v
			}
			return
		}
	}
	c.Services = append(c.Services, service)
}

// Ready returns if configuration is ready.
func (c *Config) Ready() bool {
	for _, s := range c.Services {
		if s.Identifier != "" {
			return true
		}
	}
	return false
}

// AllIdentifiers returns all unique identifiers for this device.
func (c *Config) AllIdentifiers() []string {
	var ids []string
	for _, s := range c.Services {
		if s.Identifier != "" {
			ids = append(ids, s.Identifier)
		}
	}
	return ids
}

// RemoteControl provides remote control functionality.
type RemoteControl interface {
	Up(ctx context.Context, action InputAction) error
	Down(ctx context.Context, action InputAction) error
	Left(ctx context.Context, action InputAction) error
	Right(ctx context.Context, action InputAction) error
	Play(ctx context.Context) error
	PlayPause(ctx context.Context) error
	Pause(ctx context.Context) error
	Stop(ctx context.Context) error
	Next(ctx context.Context) error
	Previous(ctx context.Context) error
	Select(ctx context.Context, action InputAction) error
	Menu(ctx context.Context, action InputAction) error
	VolumeUp(ctx context.Context) error
	VolumeDown(ctx context.Context) error
	Home(ctx context.Context, action InputAction) error
	HomeHold(ctx context.Context) error
	TopMenu(ctx context.Context) error
	Suspend(ctx context.Context) error
	WakeUp(ctx context.Context) error
	SkipForward(ctx context.Context, timeInterval float64) error
	SkipBackward(ctx context.Context, timeInterval float64) error
	SetPosition(ctx context.Context, pos int) error
	SetShuffle(ctx context.Context, state ShuffleState) error
	SetRepeat(ctx context.Context, state RepeatState) error
	ChannelUp(ctx context.Context) error
	ChannelDown(ctx context.Context) error
	Screensaver(ctx context.Context) error
}

// Metadata provides metadata retrieval functionality.
type Metadata interface {
	DeviceID() string
	Artwork(ctx context.Context, width, height *int) (*ArtworkInfo, error)
	ArtworkID() string
	Playing(ctx context.Context) (*Playing, error)
	App() *App
}

// PushListener is the listener interface for push updates.
type PushListener interface {
	PlaystatusUpdate(updater PushUpdater, playstatus *Playing)
	PlaystatusError(updater PushUpdater, err error)
}

// PushUpdater provides push update functionality.
type PushUpdater interface {
	Active() bool
	Start(initialDelay int)
	Stop()
	SetListener(listener PushListener)
}

// Stream provides streaming functionality.
type Stream interface {
	Close() error
	PlayURL(ctx context.Context, url string, options map[string]interface{}) error
	StreamFile(ctx context.Context, file io.Reader, metadata *MediaMetadata, overrideMissingMetadata bool, options map[string]interface{}) error
}

// DeviceListener is the listener interface for device updates.
type DeviceListener interface {
	ConnectionLost(err error)
	ConnectionClosed()
}

// PowerListener is the listener interface for power updates.
type PowerListener interface {
	PowerstateUpdate(oldState, newState PowerState)
}

// Power provides power management functionality.
type Power interface {
	PowerState() PowerState
	TurnOn(ctx context.Context, awaitNewState bool) error
	TurnOff(ctx context.Context, awaitNewState bool) error
	SetListener(listener PowerListener)
}

// Features provides feature functionality.
type Features interface {
	GetFeature(name FeatureName) *FeatureInfo
	AllFeatures(includeUnsupported bool) map[FeatureName]*FeatureInfo
	InState(states []FeatureState, names ...FeatureName) bool
}

// Apps provides app handling functionality.
type Apps interface {
	AppList(ctx context.Context) ([]App, error)
	LaunchApp(ctx context.Context, bundleIDOrURL string) error
}

// UserAccounts provides account handling functionality.
type UserAccounts interface {
	AccountList(ctx context.Context) ([]UserAccount, error)
	SwitchAccount(ctx context.Context, accountID string) error
}

// AudioListener is the listener interface for audio updates.
type AudioListener interface {
	VolumeUpdate(oldLevel, newLevel float64)
	OutputDevicesUpdate(oldDevices, newDevices []OutputDevice)
}

// Audio provides audio functionality.
type Audio interface {
	Volume() float64
	SetVolume(ctx context.Context, level float64) error
	VolumeUp(ctx context.Context) error
	VolumeDown(ctx context.Context) error
	OutputDevices() []OutputDevice
	AddOutputDevices(ctx context.Context, devices ...string) error
	RemoveOutputDevices(ctx context.Context, devices ...string) error
	SetOutputDevices(ctx context.Context, devices ...string) error
	SetListener(listener AudioListener)
}

// KeyboardListener is the listener interface for keyboard updates.
type KeyboardListener interface {
	FocusstateUpdate(oldState, newState KeyboardFocusState)
}

// Keyboard provides keyboard handling functionality.
type Keyboard interface {
	TextFocusState() KeyboardFocusState
	TextGet(ctx context.Context) (string, error)
	TextClear(ctx context.Context) error
	TextAppend(ctx context.Context, text string) error
	TextSet(ctx context.Context, text string) error
	SetListener(listener KeyboardListener)
}

// TouchGestures provides touch gesture functionality.
type TouchGestures interface {
	Swipe(ctx context.Context, startX, startY, endX, endY, durationMS int) error
	Action(ctx context.Context, x, y int, mode TouchAction) error
	Click(ctx context.Context, action InputAction) error
}

// PairingHandler provides pairing functionality.
type PairingHandler interface {
	Service() *Service
	Close() error
	Pin(pin string)
	DeviceProvidesPin() bool
	HasPaired() bool
	Begin(ctx context.Context) error
	Finish(ctx context.Context) error
}

// AppleTV represents a connection to an Apple TV.
type AppleTV interface {
	Connect(ctx context.Context) error
	Close() error
	DeviceInfo() *DeviceInfo
	Service() *Service
	RemoteControl() RemoteControl
	Metadata() Metadata
	PushUpdater() PushUpdater
	Stream() Stream
	Power() Power
	Features() Features
	Apps() Apps
	UserAccounts() UserAccounts
	Audio() Audio
	Keyboard() Keyboard
	Touch() TouchGestures
	SetDeviceListener(listener DeviceListener)
}

// Storage provides storage functionality for settings.
type Storage interface {
	Settings() []Settings
	Save(ctx context.Context) error
	Load(ctx context.Context) error
	GetSettings(ctx context.Context, config *Config) (*Settings, error)
	RemoveSettings(ctx context.Context, settings *Settings) (bool, error)
	UpdateSettings(ctx context.Context, config *Config) error
}

// Settings represents device settings.
type Settings struct {
	Identifier string
	Protocols  ProtocolSettings
	Info       InfoSettings
}

// ProtocolSettings contains settings for each protocol.
type ProtocolSettings struct {
	AirPlay   map[string]interface{}
	Companion map[string]interface{}
	DMAP      map[string]interface{}
	MRP       map[string]interface{}
	RAOP      map[string]interface{}
}

// InfoSettings contains device info settings.
type InfoSettings struct {
	Name string
	MAC  string
}

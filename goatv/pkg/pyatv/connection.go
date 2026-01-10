package pyatv

import (
	"context"
	"io"
	"sync"
)

// AppleTVConnection implements the AppleTV interface.
type AppleTVConnection struct {
	config         *Config
	opts           ConnectOptions
	connected      bool
	mu             sync.RWMutex
	deviceListener DeviceListener

	// Protocol handlers
	remote   RemoteControl
	metadata Metadata
	push     PushUpdater
	stream   Stream
	power    Power
	features Features
	apps     Apps
	accounts UserAccounts
	audio    Audio
	keyboard Keyboard
	touch    TouchGestures
}

// NewAppleTVConnection creates a new AppleTV connection.
func NewAppleTVConnection(config *Config, opts ConnectOptions) *AppleTVConnection {
	atv := &AppleTVConnection{
		config: config,
		opts:   opts,
	}

	// Initialize handlers
	atv.remote = &defaultRemoteControl{atv: atv}
	atv.metadata = &defaultMetadata{atv: atv}
	atv.push = &defaultPushUpdater{atv: atv}
	atv.stream = &defaultStream{atv: atv}
	atv.power = &defaultPower{atv: atv}
	atv.features = &defaultFeatures{atv: atv}
	atv.apps = &defaultApps{atv: atv}
	atv.accounts = &defaultUserAccounts{atv: atv}
	atv.audio = &defaultAudio{atv: atv}
	atv.keyboard = &defaultKeyboard{atv: atv}
	atv.touch = &defaultTouch{atv: atv}

	return atv
}

// Connect connects to the Apple TV.
func (a *AppleTVConnection) Connect(ctx context.Context) error {
	a.mu.Lock()
	defer a.mu.Unlock()

	if a.connected {
		return nil
	}

	// TODO: Implement actual protocol connections
	// For now, just mark as connected
	a.connected = true

	return nil
}

// Close closes the connection.
func (a *AppleTVConnection) Close() error {
	a.mu.Lock()
	defer a.mu.Unlock()

	if !a.connected {
		return nil
	}

	a.connected = false

	if a.deviceListener != nil {
		a.deviceListener.ConnectionClosed()
	}

	return nil
}

// DeviceInfo returns device information.
func (a *AppleTVConnection) DeviceInfo() *DeviceInfo {
	return a.config.DeviceInfo
}

// Service returns the main service.
func (a *AppleTVConnection) Service() *Service {
	if len(a.config.Services) > 0 {
		return a.config.Services[0]
	}
	return nil
}

// RemoteControl returns the remote control interface.
func (a *AppleTVConnection) RemoteControl() RemoteControl {
	return a.remote
}

// Metadata returns the metadata interface.
func (a *AppleTVConnection) Metadata() Metadata {
	return a.metadata
}

// PushUpdater returns the push updater interface.
func (a *AppleTVConnection) PushUpdater() PushUpdater {
	return a.push
}

// Stream returns the stream interface.
func (a *AppleTVConnection) Stream() Stream {
	return a.stream
}

// Power returns the power interface.
func (a *AppleTVConnection) Power() Power {
	return a.power
}

// Features returns the features interface.
func (a *AppleTVConnection) Features() Features {
	return a.features
}

// Apps returns the apps interface.
func (a *AppleTVConnection) Apps() Apps {
	return a.apps
}

// UserAccounts returns the user accounts interface.
func (a *AppleTVConnection) UserAccounts() UserAccounts {
	return a.accounts
}

// Audio returns the audio interface.
func (a *AppleTVConnection) Audio() Audio {
	return a.audio
}

// Keyboard returns the keyboard interface.
func (a *AppleTVConnection) Keyboard() Keyboard {
	return a.keyboard
}

// Touch returns the touch interface.
func (a *AppleTVConnection) Touch() TouchGestures {
	return a.touch
}

// SetDeviceListener sets the device listener.
func (a *AppleTVConnection) SetDeviceListener(listener DeviceListener) {
	a.deviceListener = listener
}

// Default implementations of interfaces

type defaultRemoteControl struct {
	atv *AppleTVConnection
}

func (r *defaultRemoteControl) Up(ctx context.Context, action InputAction) error {
	return ErrNotSupported
}

func (r *defaultRemoteControl) Down(ctx context.Context, action InputAction) error {
	return ErrNotSupported
}

func (r *defaultRemoteControl) Left(ctx context.Context, action InputAction) error {
	return ErrNotSupported
}

func (r *defaultRemoteControl) Right(ctx context.Context, action InputAction) error {
	return ErrNotSupported
}

func (r *defaultRemoteControl) Play(ctx context.Context) error {
	return ErrNotSupported
}

func (r *defaultRemoteControl) PlayPause(ctx context.Context) error {
	return ErrNotSupported
}

func (r *defaultRemoteControl) Pause(ctx context.Context) error {
	return ErrNotSupported
}

func (r *defaultRemoteControl) Stop(ctx context.Context) error {
	return ErrNotSupported
}

func (r *defaultRemoteControl) Next(ctx context.Context) error {
	return ErrNotSupported
}

func (r *defaultRemoteControl) Previous(ctx context.Context) error {
	return ErrNotSupported
}

func (r *defaultRemoteControl) Select(ctx context.Context, action InputAction) error {
	return ErrNotSupported
}

func (r *defaultRemoteControl) Menu(ctx context.Context, action InputAction) error {
	return ErrNotSupported
}

func (r *defaultRemoteControl) VolumeUp(ctx context.Context) error {
	return ErrNotSupported
}

func (r *defaultRemoteControl) VolumeDown(ctx context.Context) error {
	return ErrNotSupported
}

func (r *defaultRemoteControl) Home(ctx context.Context, action InputAction) error {
	return ErrNotSupported
}

func (r *defaultRemoteControl) HomeHold(ctx context.Context) error {
	return ErrNotSupported
}

func (r *defaultRemoteControl) TopMenu(ctx context.Context) error {
	return ErrNotSupported
}

func (r *defaultRemoteControl) Suspend(ctx context.Context) error {
	return ErrNotSupported
}

func (r *defaultRemoteControl) WakeUp(ctx context.Context) error {
	return ErrNotSupported
}

func (r *defaultRemoteControl) SkipForward(ctx context.Context, timeInterval float64) error {
	return ErrNotSupported
}

func (r *defaultRemoteControl) SkipBackward(ctx context.Context, timeInterval float64) error {
	return ErrNotSupported
}

func (r *defaultRemoteControl) SetPosition(ctx context.Context, pos int) error {
	return ErrNotSupported
}

func (r *defaultRemoteControl) SetShuffle(ctx context.Context, state ShuffleState) error {
	return ErrNotSupported
}

func (r *defaultRemoteControl) SetRepeat(ctx context.Context, state RepeatState) error {
	return ErrNotSupported
}

func (r *defaultRemoteControl) ChannelUp(ctx context.Context) error {
	return ErrNotSupported
}

func (r *defaultRemoteControl) ChannelDown(ctx context.Context) error {
	return ErrNotSupported
}

func (r *defaultRemoteControl) Screensaver(ctx context.Context) error {
	return ErrNotSupported
}

type defaultMetadata struct {
	atv *AppleTVConnection
}

func (m *defaultMetadata) DeviceID() string {
	return m.atv.config.Identifier
}

func (m *defaultMetadata) Artwork(ctx context.Context, width, height *int) (*ArtworkInfo, error) {
	return nil, ErrNotSupported
}

func (m *defaultMetadata) ArtworkID() string {
	return ""
}

func (m *defaultMetadata) Playing(ctx context.Context) (*Playing, error) {
	return nil, ErrNotSupported
}

func (m *defaultMetadata) App() *App {
	return nil
}

type defaultPushUpdater struct {
	atv      *AppleTVConnection
	active   bool
	listener PushListener
}

func (p *defaultPushUpdater) Active() bool {
	return p.active
}

func (p *defaultPushUpdater) Start(initialDelay int) {
	p.active = true
}

func (p *defaultPushUpdater) Stop() {
	p.active = false
}

func (p *defaultPushUpdater) SetListener(listener PushListener) {
	p.listener = listener
}

type defaultStream struct {
	atv *AppleTVConnection
}

func (s *defaultStream) Close() error {
	return nil
}

func (s *defaultStream) PlayURL(ctx context.Context, url string, options map[string]interface{}) error {
	return ErrNotSupported
}

func (s *defaultStream) StreamFile(ctx context.Context, file io.Reader, metadata *MediaMetadata, overrideMissingMetadata bool, options map[string]interface{}) error {
	return ErrNotSupported
}

type defaultPower struct {
	atv      *AppleTVConnection
	listener PowerListener
}

func (p *defaultPower) PowerState() PowerState {
	return PowerStateUnknown
}

func (p *defaultPower) TurnOn(ctx context.Context, awaitNewState bool) error {
	return ErrNotSupported
}

func (p *defaultPower) TurnOff(ctx context.Context, awaitNewState bool) error {
	return ErrNotSupported
}

func (p *defaultPower) SetListener(listener PowerListener) {
	p.listener = listener
}

type defaultFeatures struct {
	atv *AppleTVConnection
}

func (f *defaultFeatures) GetFeature(name FeatureName) *FeatureInfo {
	return &FeatureInfo{
		State:   FeatureStateUnsupported,
		Options: make(map[string]interface{}),
	}
}

func (f *defaultFeatures) AllFeatures(includeUnsupported bool) map[FeatureName]*FeatureInfo {
	features := make(map[FeatureName]*FeatureInfo)
	// Add all features as unsupported by default
	return features
}

func (f *defaultFeatures) InState(states []FeatureState, names ...FeatureName) bool {
	for _, name := range names {
		info := f.GetFeature(name)
		found := false
		for _, state := range states {
			if info.State == state {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

type defaultApps struct {
	atv *AppleTVConnection
}

func (a *defaultApps) AppList(ctx context.Context) ([]App, error) {
	return nil, ErrNotSupported
}

func (a *defaultApps) LaunchApp(ctx context.Context, bundleIDOrURL string) error {
	return ErrNotSupported
}

type defaultUserAccounts struct {
	atv *AppleTVConnection
}

func (u *defaultUserAccounts) AccountList(ctx context.Context) ([]UserAccount, error) {
	return nil, ErrNotSupported
}

func (u *defaultUserAccounts) SwitchAccount(ctx context.Context, accountID string) error {
	return ErrNotSupported
}

type defaultAudio struct {
	atv      *AppleTVConnection
	listener AudioListener
}

func (a *defaultAudio) Volume() float64 {
	return 0
}

func (a *defaultAudio) SetVolume(ctx context.Context, level float64) error {
	return ErrNotSupported
}

func (a *defaultAudio) VolumeUp(ctx context.Context) error {
	return ErrNotSupported
}

func (a *defaultAudio) VolumeDown(ctx context.Context) error {
	return ErrNotSupported
}

func (a *defaultAudio) OutputDevices() []OutputDevice {
	return nil
}

func (a *defaultAudio) AddOutputDevices(ctx context.Context, devices ...string) error {
	return ErrNotSupported
}

func (a *defaultAudio) RemoveOutputDevices(ctx context.Context, devices ...string) error {
	return ErrNotSupported
}

func (a *defaultAudio) SetOutputDevices(ctx context.Context, devices ...string) error {
	return ErrNotSupported
}

func (a *defaultAudio) SetListener(listener AudioListener) {
	a.listener = listener
}

type defaultKeyboard struct {
	atv      *AppleTVConnection
	listener KeyboardListener
}

func (k *defaultKeyboard) TextFocusState() KeyboardFocusState {
	return KeyboardFocusStateUnknown
}

func (k *defaultKeyboard) TextGet(ctx context.Context) (string, error) {
	return "", ErrNotSupported
}

func (k *defaultKeyboard) TextClear(ctx context.Context) error {
	return ErrNotSupported
}

func (k *defaultKeyboard) TextAppend(ctx context.Context, text string) error {
	return ErrNotSupported
}

func (k *defaultKeyboard) TextSet(ctx context.Context, text string) error {
	return ErrNotSupported
}

func (k *defaultKeyboard) SetListener(listener KeyboardListener) {
	k.listener = listener
}

type defaultTouch struct {
	atv *AppleTVConnection
}

func (t *defaultTouch) Swipe(ctx context.Context, startX, startY, endX, endY, durationMS int) error {
	return ErrNotSupported
}

func (t *defaultTouch) Action(ctx context.Context, x, y int, mode TouchAction) error {
	return ErrNotSupported
}

func (t *defaultTouch) Click(ctx context.Context, action InputAction) error {
	return ErrNotSupported
}

// Package pyatv provides a client library for Apple TV and AirPlay devices.
// It is a Go port of the Python pyatv library.
package pyatv

// Version information
const (
	MajorVersion  = "0"
	MinorVersion  = "1"
	PatchVersion  = "0"
	ShortVersion  = MajorVersion + "." + MinorVersion
	Version       = ShortVersion + "." + PatchVersion
)

// Protocol represents all supported protocols.
type Protocol int

const (
	// ProtocolDMAP represents the DMAP protocol.
	ProtocolDMAP Protocol = iota + 1
	// ProtocolMRP represents the MRP protocol.
	ProtocolMRP
	// ProtocolAirPlay represents the AirPlay protocol.
	ProtocolAirPlay
	// ProtocolCompanion represents the Companion link protocol.
	ProtocolCompanion
	// ProtocolRAOP represents the RAOP protocol.
	ProtocolRAOP
)

// String returns a string representation of the Protocol.
func (p Protocol) String() string {
	switch p {
	case ProtocolDMAP:
		return "DMAP"
	case ProtocolMRP:
		return "MRP"
	case ProtocolAirPlay:
		return "AirPlay"
	case ProtocolCompanion:
		return "Companion"
	case ProtocolRAOP:
		return "RAOP"
	default:
		return "Unknown"
	}
}

// MediaType represents all supported media types.
type MediaType int

const (
	// MediaTypeUnknown indicates the media type is not known.
	MediaTypeUnknown MediaType = iota
	// MediaTypeVideo indicates the media type is video.
	MediaTypeVideo
	// MediaTypeMusic indicates the media type is music.
	MediaTypeMusic
	// MediaTypeTV indicates the media type is a TV show.
	MediaTypeTV
)

// String returns a string representation of the MediaType.
func (m MediaType) String() string {
	switch m {
	case MediaTypeUnknown:
		return "Unknown"
	case MediaTypeVideo:
		return "Video"
	case MediaTypeMusic:
		return "Music"
	case MediaTypeTV:
		return "TV"
	default:
		return "Unknown"
	}
}

// DeviceState represents all supported device states.
type DeviceState int

const (
	// DeviceStateIdle indicates the device is idling.
	DeviceStateIdle DeviceState = iota
	// DeviceStateLoading indicates media is being loaded.
	DeviceStateLoading
	// DeviceStatePaused indicates media is paused.
	DeviceStatePaused
	// DeviceStatePlaying indicates media is playing.
	DeviceStatePlaying
	// DeviceStateStopped indicates media is stopped.
	DeviceStateStopped
	// DeviceStateSeeking indicates media is seeking.
	DeviceStateSeeking
)

// String returns a string representation of the DeviceState.
func (d DeviceState) String() string {
	switch d {
	case DeviceStateIdle:
		return "Idle"
	case DeviceStateLoading:
		return "Loading"
	case DeviceStatePaused:
		return "Paused"
	case DeviceStatePlaying:
		return "Playing"
	case DeviceStateStopped:
		return "Stopped"
	case DeviceStateSeeking:
		return "Seeking"
	default:
		return "Unknown"
	}
}

// RepeatState represents all supported repeat states.
type RepeatState int

const (
	// RepeatStateOff indicates repeat is off.
	RepeatStateOff RepeatState = iota
	// RepeatStateTrack indicates repeat current track.
	RepeatStateTrack
	// RepeatStateAll indicates repeat all tracks.
	RepeatStateAll
)

// String returns a string representation of the RepeatState.
func (r RepeatState) String() string {
	switch r {
	case RepeatStateOff:
		return "Off"
	case RepeatStateTrack:
		return "Track"
	case RepeatStateAll:
		return "All"
	default:
		return "Unknown"
	}
}

// ShuffleState represents all supported shuffle states.
type ShuffleState int

const (
	// ShuffleStateOff indicates shuffle is off.
	ShuffleStateOff ShuffleState = iota
	// ShuffleStateAlbums indicates shuffle on album level.
	ShuffleStateAlbums
	// ShuffleSateSongs indicates shuffle on song level.
	ShuffleStateSongs
)

// String returns a string representation of the ShuffleState.
func (s ShuffleState) String() string {
	switch s {
	case ShuffleStateOff:
		return "Off"
	case ShuffleStateAlbums:
		return "Albums"
	case ShuffleStateSongs:
		return "Songs"
	default:
		return "Unknown"
	}
}

// PowerState represents all supported power states.
type PowerState int

const (
	// PowerStateUnknown indicates the power state is unknown.
	PowerStateUnknown PowerState = iota
	// PowerStateOff indicates the device is off.
	PowerStateOff
	// PowerStateOn indicates the device is on.
	PowerStateOn
)

// String returns a string representation of the PowerState.
func (p PowerState) String() string {
	switch p {
	case PowerStateUnknown:
		return "Unknown"
	case PowerStateOff:
		return "Off"
	case PowerStateOn:
		return "On"
	default:
		return "Unknown"
	}
}

// KeyboardFocusState represents all supported keyboard focus states.
type KeyboardFocusState int

const (
	// KeyboardFocusStateUnknown indicates the focus state is unknown.
	KeyboardFocusStateUnknown KeyboardFocusState = iota
	// KeyboardFocusStateUnfocused indicates the keyboard is not focused.
	KeyboardFocusStateUnfocused
	// KeyboardFocusStateFocused indicates the keyboard is focused.
	KeyboardFocusStateFocused
)

// String returns a string representation of the KeyboardFocusState.
func (k KeyboardFocusState) String() string {
	switch k {
	case KeyboardFocusStateUnknown:
		return "Unknown"
	case KeyboardFocusStateUnfocused:
		return "Unfocused"
	case KeyboardFocusStateFocused:
		return "Focused"
	default:
		return "Unknown"
	}
}

// OperatingSystem represents the operating system on a device.
type OperatingSystem int

const (
	// OperatingSystemUnknown indicates the OS is unknown.
	OperatingSystemUnknown OperatingSystem = iota
	// OperatingSystemLegacy indicates Apple TV Software (pre-tvOS).
	OperatingSystemLegacy
	// OperatingSystemTvOS indicates tvOS.
	OperatingSystemTvOS
	// OperatingSystemAirPortOS indicates AirPortOS.
	OperatingSystemAirPortOS
	// OperatingSystemMacOS indicates macOS.
	OperatingSystemMacOS
)

// String returns a string representation of the OperatingSystem.
func (o OperatingSystem) String() string {
	switch o {
	case OperatingSystemUnknown:
		return "Unknown"
	case OperatingSystemLegacy:
		return "Legacy"
	case OperatingSystemTvOS:
		return "tvOS"
	case OperatingSystemAirPortOS:
		return "AirPortOS"
	case OperatingSystemMacOS:
		return "macOS"
	default:
		return "Unknown"
	}
}

// DeviceModel represents hardware device models.
type DeviceModel int

const (
	// DeviceModelUnknown indicates the model is unknown.
	DeviceModelUnknown DeviceModel = iota
	// DeviceModelGen2 is Apple TV 2.
	DeviceModelGen2
	// DeviceModelGen3 is Apple TV 3.
	DeviceModelGen3
	// DeviceModelGen4 is Apple TV 4.
	DeviceModelGen4
	// DeviceModelGen4K is Apple TV 4K.
	DeviceModelGen4K
	// DeviceModelHomePod is HomePod first generation.
	DeviceModelHomePod
	// DeviceModelHomePodMini is HomePod Mini.
	DeviceModelHomePodMini
	// DeviceModelAirPortExpress is AirPort Express first generation.
	DeviceModelAirPortExpress
	// DeviceModelAirPortExpressGen2 is AirPort Express second generation.
	DeviceModelAirPortExpressGen2
	// DeviceModelAppleTV4KGen2 is Apple TV 4K gen 2.
	DeviceModelAppleTV4KGen2
	// DeviceModelMusic is Music app running on desktop.
	DeviceModelMusic
	// DeviceModelAppleTV4KGen3 is Apple TV 4K gen 3.
	DeviceModelAppleTV4KGen3
	// DeviceModelHomePodGen2 is HomePod second generation.
	DeviceModelHomePodGen2
	// DeviceModelAppleTVGen1 is Apple TV first generation.
	DeviceModelAppleTVGen1
)

// String returns a string representation of the DeviceModel.
func (d DeviceModel) String() string {
	switch d {
	case DeviceModelUnknown:
		return "Unknown"
	case DeviceModelGen2:
		return "Apple TV 2"
	case DeviceModelGen3:
		return "Apple TV 3"
	case DeviceModelGen4:
		return "Apple TV 4"
	case DeviceModelGen4K:
		return "Apple TV 4K"
	case DeviceModelHomePod:
		return "HomePod"
	case DeviceModelHomePodMini:
		return "HomePod Mini"
	case DeviceModelAirPortExpress:
		return "AirPort Express"
	case DeviceModelAirPortExpressGen2:
		return "AirPort Express (gen 2)"
	case DeviceModelAppleTV4KGen2:
		return "Apple TV 4K (gen 2)"
	case DeviceModelMusic:
		return "Music"
	case DeviceModelAppleTV4KGen3:
		return "Apple TV 4K (gen 3)"
	case DeviceModelHomePodGen2:
		return "HomePod (gen 2)"
	case DeviceModelAppleTVGen1:
		return "Apple TV 1"
	default:
		return "Unknown"
	}
}

// InputAction represents the type of input when pressing a button.
type InputAction int

const (
	// InputActionSingleTap is press and release quickly.
	InputActionSingleTap InputAction = iota
	// InputActionDoubleTap is press and release twice quickly.
	InputActionDoubleTap
	// InputActionHold is press and hold before releasing.
	InputActionHold
)

// String returns a string representation of the InputAction.
func (i InputAction) String() string {
	switch i {
	case InputActionSingleTap:
		return "SingleTap"
	case InputActionDoubleTap:
		return "DoubleTap"
	case InputActionHold:
		return "Hold"
	default:
		return "Unknown"
	}
}

// PairingRequirement represents the pairing requirement for a service.
type PairingRequirement int

const (
	// PairingRequirementUnsupported means pairing is not supported.
	PairingRequirementUnsupported PairingRequirement = iota + 1
	// PairingRequirementDisabled means pairing is disabled.
	PairingRequirementDisabled
	// PairingRequirementNotNeeded means pairing is not needed.
	PairingRequirementNotNeeded
	// PairingRequirementOptional means pairing is optional.
	PairingRequirementOptional
	// PairingRequirementMandatory means pairing is mandatory.
	PairingRequirementMandatory
)

// String returns a string representation of the PairingRequirement.
func (p PairingRequirement) String() string {
	switch p {
	case PairingRequirementUnsupported:
		return "Unsupported"
	case PairingRequirementDisabled:
		return "Disabled"
	case PairingRequirementNotNeeded:
		return "NotNeeded"
	case PairingRequirementOptional:
		return "Optional"
	case PairingRequirementMandatory:
		return "Mandatory"
	default:
		return "Unknown"
	}
}

// FeatureState represents the state of a particular feature.
type FeatureState int

const (
	// FeatureStateUnknown means the feature state is unknown.
	FeatureStateUnknown FeatureState = iota
	// FeatureStateUnsupported means the device does not support this feature.
	FeatureStateUnsupported
	// FeatureStateUnavailable means the feature is not available now.
	FeatureStateUnavailable
	// FeatureStateAvailable means the feature is available.
	FeatureStateAvailable
)

// String returns a string representation of the FeatureState.
func (f FeatureState) String() string {
	switch f {
	case FeatureStateUnknown:
		return "Unknown"
	case FeatureStateUnsupported:
		return "Unsupported"
	case FeatureStateUnavailable:
		return "Unavailable"
	case FeatureStateAvailable:
		return "Available"
	default:
		return "Unknown"
	}
}

// FeatureName represents all supported features.
type FeatureName int

const (
	FeatureUp FeatureName = iota
	FeatureDown
	FeatureLeft
	FeatureRight
	FeaturePlay
	FeaturePlayPause
	FeaturePause
	FeatureStop
	FeatureNext
	FeaturePrevious
	FeatureSelect
	FeatureMenu
	FeatureVolumeUp
	FeatureVolumeDown
	FeatureHome
	FeatureHomeHold
	FeatureTopMenu
	FeatureSuspend
	FeatureWakeUp
	FeatureSetPosition
	FeatureSetShuffle
	FeatureSetRepeat
	FeatureTitle
	FeatureArtist
	FeatureAlbum
	FeatureGenre
	FeatureTotalTime
	FeaturePosition
	FeatureShuffle
	FeatureRepeat
	FeatureArtwork
	FeaturePlayURL
	FeaturePowerState
	FeatureTurnOn
	FeatureTurnOff
	FeatureApp
	FeatureSkipForward
	FeatureSkipBackward
	FeatureAppList
	FeatureLaunchApp
	FeatureSeriesName
	FeatureSeasonNumber
	FeatureEpisodeNumber
	FeaturePushUpdates
	FeatureStreamFile
	FeatureVolume
	FeatureSetVolume
	FeatureContentIdentifier
	FeatureChannelUp
	FeatureChannelDown
	FeatureITunesStoreIdentifier
	FeatureTextGet
	FeatureTextClear
	FeatureTextAppend
	FeatureTextSet
	FeatureAccountList
	FeatureSwitchAccount
	FeatureTextFocusState
	FeatureScreensaver
	FeatureOutputDevices
	FeatureAddOutputDevices
	FeatureRemoveOutputDevices
	FeatureSetOutputDevices
	FeatureSwipe
	FeatureAction
	FeatureClick
)

// String returns a string representation of the FeatureName.
func (f FeatureName) String() string {
	names := map[FeatureName]string{
		FeatureUp:                    "Up",
		FeatureDown:                  "Down",
		FeatureLeft:                  "Left",
		FeatureRight:                 "Right",
		FeaturePlay:                  "Play",
		FeaturePlayPause:             "PlayPause",
		FeaturePause:                 "Pause",
		FeatureStop:                  "Stop",
		FeatureNext:                  "Next",
		FeaturePrevious:              "Previous",
		FeatureSelect:                "Select",
		FeatureMenu:                  "Menu",
		FeatureVolumeUp:              "VolumeUp",
		FeatureVolumeDown:            "VolumeDown",
		FeatureHome:                  "Home",
		FeatureHomeHold:              "HomeHold",
		FeatureTopMenu:               "TopMenu",
		FeatureSuspend:               "Suspend",
		FeatureWakeUp:                "WakeUp",
		FeatureSetPosition:           "SetPosition",
		FeatureSetShuffle:            "SetShuffle",
		FeatureSetRepeat:             "SetRepeat",
		FeatureTitle:                 "Title",
		FeatureArtist:                "Artist",
		FeatureAlbum:                 "Album",
		FeatureGenre:                 "Genre",
		FeatureTotalTime:             "TotalTime",
		FeaturePosition:              "Position",
		FeatureShuffle:               "Shuffle",
		FeatureRepeat:                "Repeat",
		FeatureArtwork:               "Artwork",
		FeaturePlayURL:               "PlayUrl",
		FeaturePowerState:            "PowerState",
		FeatureTurnOn:                "TurnOn",
		FeatureTurnOff:               "TurnOff",
		FeatureApp:                   "App",
		FeatureSkipForward:           "SkipForward",
		FeatureSkipBackward:          "SkipBackward",
		FeatureAppList:               "AppList",
		FeatureLaunchApp:             "LaunchApp",
		FeatureSeriesName:            "SeriesName",
		FeatureSeasonNumber:          "SeasonNumber",
		FeatureEpisodeNumber:         "EpisodeNumber",
		FeaturePushUpdates:           "PushUpdates",
		FeatureStreamFile:            "StreamFile",
		FeatureVolume:                "Volume",
		FeatureSetVolume:             "SetVolume",
		FeatureContentIdentifier:     "ContentIdentifier",
		FeatureChannelUp:             "ChannelUp",
		FeatureChannelDown:           "ChannelDown",
		FeatureITunesStoreIdentifier: "iTunesStoreIdentifier",
		FeatureTextGet:               "TextGet",
		FeatureTextClear:             "TextClear",
		FeatureTextAppend:            "TextAppend",
		FeatureTextSet:               "TextSet",
		FeatureAccountList:           "AccountList",
		FeatureSwitchAccount:         "SwitchAccount",
		FeatureTextFocusState:        "TextFocusState",
		FeatureScreensaver:           "Screensaver",
		FeatureOutputDevices:         "OutputDevices",
		FeatureAddOutputDevices:      "AddOutputDevices",
		FeatureRemoveOutputDevices:   "RemoveOutputDevices",
		FeatureSetOutputDevices:      "SetOutputDevices",
		FeatureSwipe:                 "Swipe",
		FeatureAction:                "Action",
		FeatureClick:                 "Click",
	}
	if name, ok := names[f]; ok {
		return name
	}
	return "Unknown"
}

// TouchAction represents touch action constants.
type TouchAction int

const (
	// TouchActionPress is a touch press action.
	TouchActionPress TouchAction = 1
	// TouchActionHold is a touch hold action.
	TouchActionHold TouchAction = 3
	// TouchActionRelease is a touch release action.
	TouchActionRelease TouchAction = 4
	// TouchActionClick is a touch click action.
	TouchActionClick TouchAction = 5
)

// String returns a string representation of the TouchAction.
func (t TouchAction) String() string {
	switch t {
	case TouchActionPress:
		return "Press"
	case TouchActionHold:
		return "Hold"
	case TouchActionRelease:
		return "Release"
	case TouchActionClick:
		return "Click"
	default:
		return "Unknown"
	}
}

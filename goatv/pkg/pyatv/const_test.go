package pyatv

import (
	"testing"
)

func TestProtocolString(t *testing.T) {
	tests := []struct {
		protocol Protocol
		expected string
	}{
		{ProtocolDMAP, "DMAP"},
		{ProtocolMRP, "MRP"},
		{ProtocolAirPlay, "AirPlay"},
		{ProtocolCompanion, "Companion"},
		{ProtocolRAOP, "RAOP"},
		{Protocol(100), "Unknown"},
	}

	for _, tt := range tests {
		t.Run(tt.expected, func(t *testing.T) {
			if got := tt.protocol.String(); got != tt.expected {
				t.Errorf("Protocol.String() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestMediaTypeString(t *testing.T) {
	tests := []struct {
		mediaType MediaType
		expected  string
	}{
		{MediaTypeUnknown, "Unknown"},
		{MediaTypeVideo, "Video"},
		{MediaTypeMusic, "Music"},
		{MediaTypeTV, "TV"},
	}

	for _, tt := range tests {
		t.Run(tt.expected, func(t *testing.T) {
			if got := tt.mediaType.String(); got != tt.expected {
				t.Errorf("MediaType.String() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestDeviceStateString(t *testing.T) {
	tests := []struct {
		state    DeviceState
		expected string
	}{
		{DeviceStateIdle, "Idle"},
		{DeviceStateLoading, "Loading"},
		{DeviceStatePaused, "Paused"},
		{DeviceStatePlaying, "Playing"},
		{DeviceStateStopped, "Stopped"},
		{DeviceStateSeeking, "Seeking"},
	}

	for _, tt := range tests {
		t.Run(tt.expected, func(t *testing.T) {
			if got := tt.state.String(); got != tt.expected {
				t.Errorf("DeviceState.String() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestRepeatStateString(t *testing.T) {
	tests := []struct {
		state    RepeatState
		expected string
	}{
		{RepeatStateOff, "Off"},
		{RepeatStateTrack, "Track"},
		{RepeatStateAll, "All"},
	}

	for _, tt := range tests {
		t.Run(tt.expected, func(t *testing.T) {
			if got := tt.state.String(); got != tt.expected {
				t.Errorf("RepeatState.String() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestShuffleStateString(t *testing.T) {
	tests := []struct {
		state    ShuffleState
		expected string
	}{
		{ShuffleStateOff, "Off"},
		{ShuffleStateAlbums, "Albums"},
		{ShuffleStateSongs, "Songs"},
	}

	for _, tt := range tests {
		t.Run(tt.expected, func(t *testing.T) {
			if got := tt.state.String(); got != tt.expected {
				t.Errorf("ShuffleState.String() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestPowerStateString(t *testing.T) {
	tests := []struct {
		state    PowerState
		expected string
	}{
		{PowerStateUnknown, "Unknown"},
		{PowerStateOff, "Off"},
		{PowerStateOn, "On"},
	}

	for _, tt := range tests {
		t.Run(tt.expected, func(t *testing.T) {
			if got := tt.state.String(); got != tt.expected {
				t.Errorf("PowerState.String() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestDeviceModelString(t *testing.T) {
	tests := []struct {
		model    DeviceModel
		expected string
	}{
		{DeviceModelUnknown, "Unknown"},
		{DeviceModelGen2, "Apple TV 2"},
		{DeviceModelGen3, "Apple TV 3"},
		{DeviceModelGen4, "Apple TV 4"},
		{DeviceModelGen4K, "Apple TV 4K"},
		{DeviceModelHomePod, "HomePod"},
		{DeviceModelHomePodMini, "HomePod Mini"},
	}

	for _, tt := range tests {
		t.Run(tt.expected, func(t *testing.T) {
			if got := tt.model.String(); got != tt.expected {
				t.Errorf("DeviceModel.String() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestInputActionString(t *testing.T) {
	tests := []struct {
		action   InputAction
		expected string
	}{
		{InputActionSingleTap, "SingleTap"},
		{InputActionDoubleTap, "DoubleTap"},
		{InputActionHold, "Hold"},
	}

	for _, tt := range tests {
		t.Run(tt.expected, func(t *testing.T) {
			if got := tt.action.String(); got != tt.expected {
				t.Errorf("InputAction.String() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestPairingRequirementString(t *testing.T) {
	tests := []struct {
		req      PairingRequirement
		expected string
	}{
		{PairingRequirementUnsupported, "Unsupported"},
		{PairingRequirementDisabled, "Disabled"},
		{PairingRequirementNotNeeded, "NotNeeded"},
		{PairingRequirementOptional, "Optional"},
		{PairingRequirementMandatory, "Mandatory"},
	}

	for _, tt := range tests {
		t.Run(tt.expected, func(t *testing.T) {
			if got := tt.req.String(); got != tt.expected {
				t.Errorf("PairingRequirement.String() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestFeatureStateString(t *testing.T) {
	tests := []struct {
		state    FeatureState
		expected string
	}{
		{FeatureStateUnknown, "Unknown"},
		{FeatureStateUnsupported, "Unsupported"},
		{FeatureStateUnavailable, "Unavailable"},
		{FeatureStateAvailable, "Available"},
	}

	for _, tt := range tests {
		t.Run(tt.expected, func(t *testing.T) {
			if got := tt.state.String(); got != tt.expected {
				t.Errorf("FeatureState.String() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestFeatureNameString(t *testing.T) {
	tests := []struct {
		name     FeatureName
		expected string
	}{
		{FeatureUp, "Up"},
		{FeatureDown, "Down"},
		{FeaturePlay, "Play"},
		{FeaturePause, "Pause"},
	}

	for _, tt := range tests {
		t.Run(tt.expected, func(t *testing.T) {
			if got := tt.name.String(); got != tt.expected {
				t.Errorf("FeatureName.String() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestTouchActionString(t *testing.T) {
	tests := []struct {
		action   TouchAction
		expected string
	}{
		{TouchActionPress, "Press"},
		{TouchActionHold, "Hold"},
		{TouchActionRelease, "Release"},
		{TouchActionClick, "Click"},
	}

	for _, tt := range tests {
		t.Run(tt.expected, func(t *testing.T) {
			if got := tt.action.String(); got != tt.expected {
				t.Errorf("TouchAction.String() = %v, want %v", got, tt.expected)
			}
		})
	}
}

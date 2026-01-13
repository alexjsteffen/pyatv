package pyatv

import "errors"

// Common errors used by the library.
var (
	// ErrNoService is returned when connecting to a device with no usable service.
	ErrNoService = errors.New("no service to connect to")

	// ErrUnsupportedProtocol is returned when an unsupported protocol was requested.
	ErrUnsupportedProtocol = errors.New("unsupported protocol")

	// ErrConnectionFailed is returned when connection fails.
	ErrConnectionFailed = errors.New("connection failed")

	// ErrConnectionLost is returned when a connection is lost.
	ErrConnectionLost = errors.New("connection lost")

	// ErrPairing is returned when pairing fails.
	ErrPairing = errors.New("pairing failed")

	// ErrAuthentication is returned when authentication fails.
	ErrAuthentication = errors.New("authentication failed")

	// ErrNotSupported is returned when trying to perform an unsupported action.
	ErrNotSupported = errors.New("not supported")

	// ErrInvalidDMAPData is returned when invalid DMAP data is parsed.
	ErrInvalidDMAPData = errors.New("invalid DMAP data")

	// ErrUnknownMediaKind is returned when an unknown media kind is found.
	ErrUnknownMediaKind = errors.New("unknown media kind")

	// ErrUnknownPlayState is returned when an unknown play state is found.
	ErrUnknownPlayState = errors.New("unknown play state")

	// ErrNoAsyncListener is returned when starting async updater with no listener.
	ErrNoAsyncListener = errors.New("no async listener")

	// ErrNoCredentials is returned when credentials are missing.
	ErrNoCredentials = errors.New("no credentials")

	// ErrInvalidCredentials is returned when credentials are invalid.
	ErrInvalidCredentials = errors.New("invalid credentials")

	// ErrDeviceIDMissing is returned when device id is missing.
	ErrDeviceIDMissing = errors.New("device id missing")

	// ErrBackOff is returned when device mandates a backoff period.
	ErrBackOff = errors.New("backoff required")

	// ErrPlayback is returned when media playback failed.
	ErrPlayback = errors.New("playback failed")

	// ErrCommand is returned when a command failed.
	ErrCommand = errors.New("command failed")

	// ErrInvalidState is returned when trying to perform an action not possible in current state.
	ErrInvalidState = errors.New("invalid state")

	// ErrProtocol is returned when a generic protocol error occurs.
	ErrProtocol = errors.New("protocol error")

	// ErrInvalidConfig is returned when something is wrong in the config.
	ErrInvalidConfig = errors.New("invalid config")

	// ErrBlocked is returned when calling a blocked method.
	ErrBlocked = errors.New("method blocked")

	// ErrInvalidResponse is returned when a remote sends an invalid response.
	ErrInvalidResponse = errors.New("invalid response")

	// ErrOperationTimeout is returned when a timeout happens.
	ErrOperationTimeout = errors.New("operation timeout")

	// ErrSettings is returned when an error related to settings happens.
	ErrSettings = errors.New("settings error")
)

// HTTPError represents an HTTP error with status code.
type HTTPError struct {
	Message    string
	StatusCode int
}

// Error implements the error interface.
func (e *HTTPError) Error() string {
	return e.Message
}

// NewHTTPError creates a new HTTPError.
func NewHTTPError(message string, statusCode int) *HTTPError {
	return &HTTPError{
		Message:    message,
		StatusCode: statusCode,
	}
}

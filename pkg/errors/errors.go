package errors

import "fmt"

// ErrorCode represents standardized error codes
type ErrorCode string

const (
	ErrInvalidRequest    ErrorCode = "INVALID_REQUEST"
	ErrInsufficientFunds ErrorCode = "INSUFFICIENT_FUNDS"
	ErrNetworkFailure    ErrorCode = "NETWORK_FAILURE"
	ErrProofGeneration   ErrorCode = "PROOF_GENERATION_FAILED"
	ErrAgentUnavailable  ErrorCode = "AGENT_UNAVAILABLE"
	ErrTimeout           ErrorCode = "TIMEOUT"
)

// SDKError is a structured error type for better error handling
type SDKError struct {
	Code    ErrorCode
	Message string
	Cause   error
}

func (e *SDKError) Error() string {
	if e.Cause != nil {
		return fmt.Sprintf("[%s] %s: %v", e.Code, e.Message, e.Cause)
	}
	return fmt.Sprintf("[%s] %s", e.Code, e.Message)
}

func (e *SDKError) Unwrap() error {
	return e.Cause
}

// New creates a new SDK error
func New(code ErrorCode, message string) *SDKError {
	return &SDKError{
		Code:    code,
		Message: message,
	}
}

// Wrap wraps an existing error with SDK error context
func Wrap(code ErrorCode, message string, cause error) *SDKError {
	return &SDKError{
		Code:    code,
		Message: message,
		Cause:   cause,
	}
}

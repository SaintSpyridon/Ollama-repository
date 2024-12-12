package api

import (
	"fmt"
	"strings"
)

const InvalidModelNameErrMsg = "invalid model name"

// API error responses
// ErrorCode represents a standardized error code identifier
type ErrorCode string

const (
	ErrCodeUnknownKey ErrorCode = "unknown_key"
	ErrCodeGeneral    ErrorCode = "general" // Generic fallback error code
)

// ErrorResponse implements a structured error interface
type ErrorResponse struct {
	Message string         `json:"error"` // Human-readable error message
	Code    ErrorCode      `json:"code"`  // Machine-readable error code for programmatic handling, not response code
	Data    map[string]any `json:"data"`  // Additional error specific data, if any
}

func (e ErrorResponse) Error() string {
	return e.Message
}

type ErrUnknownOllamaKey struct {
	Message string
	Key     string
}

func (e ErrUnknownOllamaKey) Error() string {
	return fmt.Sprintf("unauthorized: unknown ollama key %q", strings.TrimSpace(e.Key))
}

// StatusError is an error with an HTTP status code and message,
// it is parsed on the client-side and not returned from the API
type StatusError struct {
	StatusCode   int    // e.g. 200
	Status       string // e.g. "200 OK"
	ErrorMessage string `json:"error"`
}

func (e StatusError) Error() string {
	switch {
	case e.Status != "" && e.ErrorMessage != "":
		return fmt.Sprintf("%s: %s", e.Status, e.ErrorMessage)
	case e.Status != "":
		return e.Status
	case e.ErrorMessage != "":
		return e.ErrorMessage
	default:
		// this should not happen
		return "something went wrong, please see the ollama server logs for details"
	}
}

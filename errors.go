package kokkaiapi

import "errors"

var (
	ErrRequestFailed   = errors.New("http request failed")
	ErrNonOKResponse   = errors.New("non-OK HTTP status")
	ErrParsingResponse = errors.New("failed to parse response")
)

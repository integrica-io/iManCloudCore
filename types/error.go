package types

import (
	"fmt"
)

type ErrorResponse struct {
	Error struct {
		Code        string         `json:"code"`
		CodeMessage string         `json:"code_message"`
		Messages    []errorMessage `json:"messages"`
	} `json:"error"`
}

type errorMessage struct {
	Field       string `json:"field"`
	Code        string `json:"code"`
	CodeMessage string `json:"code_message"`
}

type HTTPError struct {
	Status      int
	Code        string
	RawString	string
}

func (e *HTTPError) Error() string {
	return fmt.Sprintf("HTTP %d: %s", e.Status, e.RawString)
}
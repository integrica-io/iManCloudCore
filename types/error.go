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

type HttpError struct {
	Status      int
	Code        string
	RawString	string
}

func (e *HttpError) Error() error {
	return fmt.Errorf("status: %d, code: %s, messages: %s", e.Status, e.Code, e.RawString)
}

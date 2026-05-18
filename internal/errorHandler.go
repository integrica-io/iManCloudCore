package internal

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type errorResponse struct {
	Error struct {
		Code        string `json:"code"`
		CodeMessage string `json:"code_message"`
		Messages    []message `json:"messages"`
	} `json:"error"`
}

type message struct {
	Field       string `json:"field"`
	Code        string `json:"code"`
	CodeMessage string `json:"code_message"`
}

type httpError struct{
	status int
	code string
	codeMessage string
	messages []message
}

func (e *httpError) error() error {
	return fmt.Errorf("status: %d, code: %s, codeMessage: %s, messages: %s", e.status, e.code, e.codeMessage, e.messages)
}

func HttpErrorHandler(resp *http.Response)error{
	errorResp := new(errorResponse)
	reader, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("HttpErrorHandler, Reader, %s", err)
	}
	if err := json.Unmarshal(reader, errorResp); err != nil {
		return fmt.Errorf("HttpErrorHandler, Unmarshal, %s", err)
	}
	errStruct := httpError{
		status: resp.StatusCode,
		code: errorResp.Error.Code,
		codeMessage: errorResp.Error.CodeMessage,
		messages: errorResp.Error.Messages,
	}	
	return errStruct.error() 
}
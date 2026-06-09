package internal

import (
	"encoding/json"
	"fmt"
	"github.com/integrica-io/iManCloudCore/types"
	"io"
	"net/http"
)

func HttpErrorHandler(resp *http.Response)error{
	errorResp := new(types.ErrorResponse)
	reader, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("HttpErrorHandler, Reader, %s", err)
	}
	if err := json.Unmarshal(reader, errorResp); err != nil {
		return fmt.Errorf("HttpErrorHandler, Unmarshal, %s", err)
	}
	errStruct := types.HttpError{
		Status: resp.StatusCode,
		Code: errorResp.Error.Code,
		CodeMessage: errorResp.Error.CodeMessage,
		Messages: errorResp.Error.Messages,
	}	
	return errStruct.Error() 
}
package internal

import (
	"fmt"
	"io"
	"net/http"

	"github.com/integrica-io/iManCloudCore/types"
)

func HttpErrorHandler(resp http.Response) error {
	errorResp := new(types.ErrorResponse)
	reader, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("HttpErrorHandler, Reader, %s", err)
	}
	errStruct := types.HttpError{
		Status:     resp.StatusCode,
		Code:       errorResp.Error.Code,
		RawString: 	string(reader),
	}
	return errStruct.Error()
}

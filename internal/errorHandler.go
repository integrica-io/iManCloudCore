package internal

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/integrica-io/iManCloudCore/types"
)

func HttpErrorHandler(resp http.Response) error {
	defer resp.Body.Close()
	reader, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("HttpErrorHandler, Reader, %s", err)
	}
	var errResp types.ErrorResponse
	_ = json.Unmarshal(reader, &errResp)
	
	return &types.HTTPError{
		Status:     resp.StatusCode,
		RawString: 	string(reader),
	}
}

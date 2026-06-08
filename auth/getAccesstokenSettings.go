package auth

import (
	"context"
	"fmt"
	"github.com/integrica-io/iManCloudCore/internal"
)

type GetAccessTokenSettingsOutput struct {
	Data struct {
		AbsoluteTimeoutEnabled bool `json:"absolute_timeout_enabled"`
		AbsoluteTimeoutSeconds int  `json:"absolute_timeout_seconds"`
	} `json:"data"`
}

func GetAccessTokenSettings(ctx context.Context, client *internal.Client) error {
	if client.Token == nil{
		return fmt.Errorf("access token required")
	}
	
	return nil
}
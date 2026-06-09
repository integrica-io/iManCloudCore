package auth

import (
	"context"
	"fmt"
	"github.com/integrica-io/iManCloudCore/client"
)

type GetAccessTokenSettingsOutput struct {
	Data struct {
		AbsoluteTimeoutEnabled bool `json:"absolute_timeout_enabled"`
		AbsoluteTimeoutSeconds int  `json:"absolute_timeout_seconds"`
	} `json:"data"`
}

func GetAccessTokenSettings(ctx context.Context, client *client.Client) error {
	if client.Token == nil{
		return fmt.Errorf("access token required")
	}
	
	return nil
}
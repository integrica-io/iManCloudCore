package auth

import (
	"context"
	"fmt"
	"iManCloudCore/internal"
)

type GetAccessTokenSettingsOutput struct {
	Data struct {
		AbsoluteTimeoutEnabled bool `json:"absolute_timeout_enabled"`
		AbsoluteTimeoutSeconds int  `json:"absolute_timeout_seconds"`
	} `json:"data"`
}

func GetAccessTokenSettings(ctx context.Context, c *internal.Client) error {
	if c.Token == nil{
		return fmt.Errorf("access token required")
	}	
	return nil
}
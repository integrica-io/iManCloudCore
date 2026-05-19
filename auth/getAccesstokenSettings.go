package auth

import (
	"context"
	"fmt"
)

type GetAccessTokenSettingsOutput struct {
	Data struct {
		AbsoluteTimeoutEnabled bool `json:"absolute_timeout_enabled"`
		AbsoluteTimeoutSeconds int  `json:"absolute_timeout_seconds"`
	} `json:"data"`
}

func (c *Client) GetAccessTokenSettings(ctx context.Context) error {
	if c.Token == nil{
		return fmt.Errorf("access token required")
	}	
	return nil
}
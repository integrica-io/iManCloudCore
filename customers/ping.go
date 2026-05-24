package customers

import (
	"context"
	"iManCloudCore/internal"
)

func PingServer(ctx context.Context, c *internal.Client) error {
	endpoint := c.BaseUrl.JoinPath("work", "api", "v2", "customers", c.TokenCfg.CustomerId, "ping")
	
	req := internal.HttpRequestBuilder{}
		
	req.Url(*endpoint).Method(internal.Get).Context(ctx)

	if err := req.Exec(); err != nil {
		return err
	}	
	return nil	
}

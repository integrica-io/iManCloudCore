package customers

import (
	"context"
	"iManCloudCore/internal"
)

func PingServer(ctx context.Context, client *internal.Client) error {
	endpoint := client.BaseUrl.JoinPath("work", "api", "v2", "customers", client.TokenCfg.CustomerId, "ping")
	
	req := internal.HttpRequestBuilder{}
		
	req.Url(*endpoint).Method(internal.Get).Context(ctx)

	if err := client.Req(req); err != nil {
		return err
	}	
	return nil	
}

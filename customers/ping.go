package customers

import (
	"context"
	"github.com/integrica-io/iManCloudCore/internal"
	"github.com/integrica-io/iManCloudCore/client"
)

func PingServer(ctx context.Context, client *client.Client) error {
	endpoint := client.BaseUrl.JoinPath("work", "api", "v2", "customers", client.TokenCfg.CustomerId, "ping")
	
	req := internal.HttpRequestBuilder{}
		
	req.Url(*endpoint).Method(internal.Get).Context(ctx)

	if err := client.Req(req); err != nil {
		return err
	}	
	return nil	
}

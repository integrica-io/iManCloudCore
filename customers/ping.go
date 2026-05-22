package customers

import (
	"context"
	"iManCloudCore/internal"
	"log/slog"
	"os"
)

func PingServer(ctx context.Context, c *internal.Client) error {
	endpoint := c.BaseUrl.JoinPath("work", "api", "v2", "customers", c.TokenCfg.CustomerId, "ping")
	
	req := internal.HttpRequestBuilder{}
	
	headers := make(map[string]string)
	headers["X-Auth-Token"] = c.Token.AccessToken
		
	req.Url(*endpoint).Method(internal.Get).Headers(headers)

	if err := req.Exec(); err != nil {
		slog.Error("error","error", err)
		os.Exit(1)
	}
	return nil	
}

package auth

import (
	"context"
	"fmt"
	"net/url"

	"github.com/integrica-io/iManCloudCore/client"
	"github.com/integrica-io/iManCloudCore/internal"
)

func RevokeClientAccessToken(ctx context.Context, client *client.Client) error {
	endpoint := client.BaseUrl.JoinPath("auth", "oauth2", "revoke-token")

	if client.Token == nil {
		return fmt.Errorf("no access token available to revoke")
	}

	data := url.Values{}
	data.Add("access_token", client.Token.AccessToken)

	req := internal.HttpRequestBuilder{}

	req.Context(ctx).Url(*endpoint).Method(internal.Post).Form(data)

	if err := client.Req(req); err != nil {
		return err
	}

	return nil
}

func RevokeAccessToken(ctx context.Context, client *client.Client, accessToken string) error {
	endpoint := client.BaseUrl.JoinPath("auth", "oauth2", "revoke-token")

	data := url.Values{}
	data.Add("access_token", accessToken)

	req := internal.HttpRequestBuilder{}

	req.Context(ctx).Url(*endpoint).Method(internal.Post).Form(data)

	if err := client.Req(req); err != nil {
		return err
	}
	return nil
}

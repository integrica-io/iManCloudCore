package auth

import (
	"context"
	"fmt"
	"iManCloudCore/internal"
	"net/url"
)

func RevokeClientAccessToken(ctx context.Context, client *internal.Client) error {
	endpoint := client.BaseUrl.JoinPath("auth","oauth2","revoke-token")

	if client.Token == nil{
		return fmt.Errorf("no access token available to revoke")
	}

	data := url.Values{}
	data.Add("access_token", client.Token.AccessToken)

	req := internal.HttpRequestBuilder{}

	req.Context(ctx).Url(*endpoint).Method(internal.Post).Form(data)

	if err := req.Exec(); err != nil {
		return err
	}

	return nil
}

func RevokeAccessToken(ctx context.Context, client *internal.Client, accessToken string) error {
	endpoint := client.BaseUrl.JoinPath("auth","oauth2","revoke-token")

	data := url.Values{}
	data.Add("access_token", accessToken)

	req := internal.HttpRequestBuilder{}

	req.Context(ctx).Url(*endpoint).Method(internal.Post).Form(data)

	if err := req.Exec(); err != nil {
		return err
	}
	return nil
}
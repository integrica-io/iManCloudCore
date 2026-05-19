package auth

import (
	"context"
	"fmt"
	"strings"
	"net/http"
	"net/url"

	"iManCloudCore/internal"
)

func (c *Client)revokeClientAccessToken(ctx context.Context) error {
	endpoint := c.BaseUrl.JoinPath("auth","oauth2","revoke-token")
	data := url.Values{}

	data.Add("access_token", c.AccessToken)
	
	if c.AccessToken == ""{
		return fmt.Errorf("no access token available to revoke")
	}

	reqClient := &http.Client{}
	req, err := http.NewRequestWithContext(ctx, "POST", endpoint.String(), strings.NewReader(data.Encode()))
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := reqClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	okResp := resp.StatusCode >= 200 || resp.StatusCode < 300
	if !okResp{
		errMessage := internal.HttpErrorHandler(resp)
		return fmt.Errorf("%s", errMessage)
	}

	return nil
}
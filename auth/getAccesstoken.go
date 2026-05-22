package auth

import (
	"context"
	"fmt"
	"iManCloudCore/internal"
	"iManCloudCore/types"
	"net/url"
)

func GetAccessToken(ctx context.Context, c *internal.Client) (error) {
	endpoint := c.BaseUrl.JoinPath("auth","oauth2","token")
	data := url.Values{}

	data.Set("grant_type", string(c.TokenCfg.Grant))
	
	if c.TokenCfg.Grant == ""{
		return fmt.Errorf("empty grant type") 
	}

	if err := c.TokenCfg.IsValid(); err != nil {
		return fmt.Errorf("GetAccessToken, tokenCfg.isValid, %s", err)
	}

	if c.TokenCfg.Grant != types.RefreshToken{
		data.Set("client_id", c.TokenCfg.ClientId)
		data.Set("client_secret", c.TokenCfg.ClientSecret)
	}
		
	switch c.TokenCfg.Grant{
		case types.Password:
			data.Set("password", c.TokenCfg.Password)
			data.Set("username", c.TokenCfg.Username)
		case types.Saml2Bearer:
			data.Set("assertion", c.TokenCfg.Assertion)
	}

	req := internal.HttpRequestBuilder{}

	var GetAccesstokenOutput types.Token

	req.Context(ctx).Url(*endpoint).Method(internal.Post).ToJson(&GetAccesstokenOutput).Form(data)

	if err := req.Exec(); err != nil {
		return err
	}
	
	c.Token = &GetAccesstokenOutput
	
	return nil
}
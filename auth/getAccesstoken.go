package auth

import (
	"context"
	"fmt"
	"iManCloudCore/internal"
	"iManCloudCore/types"
	"net/url"
	"time"
)

func GetAccessToken(ctx context.Context, client *internal.Client) (error) {
	endpoint := client.BaseUrl.JoinPath("auth","oauth2","token")
	data := url.Values{}

	data.Set("grant_type", string(client.TokenCfg.Grant))
	
	if client.TokenCfg.Grant == ""{
		return fmt.Errorf("empty grant type") 
	}

	if err := client.TokenCfg.IsValid(); err != nil {
		return fmt.Errorf("GetAccessToken, tokenCfg.isValid, %s", err)
	}

	if client.TokenCfg.Grant != types.RefreshToken{
		data.Set("client_id", client.TokenCfg.ClientId)
		data.Set("client_secret", client.TokenCfg.ClientSecret)
	}
		
	switch client.TokenCfg.Grant{
		case types.Password:
			data.Set("password", client.TokenCfg.Password)
			data.Set("username", client.TokenCfg.Username)
		case types.Saml2Bearer:
			data.Set("assertion", client.TokenCfg.Assertion)
	}

	req := internal.HttpRequestBuilder{}

	var GetAccesstokenOutput types.Token

	req.Context(ctx).Url(*endpoint).Method(internal.Post).ToJson(&GetAccesstokenOutput).Form(data)

	if err := client.Req(req); err != nil {
		return err
	}
	
	client.Token = &GetAccesstokenOutput
	client.Token.TokenExpiry = time.Now().Add(time.Second * time.Duration(client.Token.ExpiresIn))
	
	return nil
}
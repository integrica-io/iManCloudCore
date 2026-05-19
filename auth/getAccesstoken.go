package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	
	"iManCloudCore/internal"
)

const (
	AuthorizationCode GrantType = "authorization_code"
	Password          GrantType = "password"
	RefreshToken      GrantType = "refresh_token"
	Saml2Bearer       GrantType = "urn:ietf:params:oauth:grant-type:saml2-bearer"
)

type Token struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    string `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	TokenType    string `json:"token_type"`
}

type AccessTokenCfg struct{
	Assertion 		string
	ClientId		string
	ClientSecret	string
	CustomerId		string
	Password		string
	Username		string
	Grant			GrantType		
}

type GrantType string

func (cfg *AccessTokenCfg) isValid() error {
	if cfg.Grant == "" {
		return fmt.Errorf("empty grant type")
	}
	
	if cfg.Grant != RefreshToken{
		if cfg.ClientId == ""{
			return fmt.Errorf("client_id required for %s grant type", cfg.Grant)
		}
		if cfg.ClientSecret == ""{
			return fmt.Errorf("client_secret required for %s grant type", cfg.Grant)
		}
	}
	
	if cfg.Grant == Saml2Bearer { 
		if cfg.CustomerId == ""{
			return fmt.Errorf("customer_id required for %s grant type", cfg.Grant)
		}
		if cfg.Assertion == ""{
			return fmt.Errorf("assertion required for %s grant type", cfg.Grant)
		}
	}
	
	if cfg.Grant == Password{
		if cfg.Password == ""{
			return fmt.Errorf("password required for %s grant type", cfg.Grant)	
		}
		if cfg.Username == ""{
			return fmt.Errorf("username required for %s grant type", cfg.Grant)
		}
	}
	
	return nil
}

func (grant GrantType) IsValid() bool{
	switch grant {
		case AuthorizationCode, Password, RefreshToken, Saml2Bearer:
			return true
	}
	return false
}

func (c *Client) GetAccessToken(ctx context.Context) (error) {
	endpoint := c.BaseUrl.JoinPath("auth","oauth2","token")
	data := url.Values{}
	
	if c.TokenCfg.Grant == ""{
		return fmt.Errorf("empty grant type") 
	}
	fmt.Println("token not empty")

	if err := c.TokenCfg.isValid(); err != nil {
		return fmt.Errorf("GetAccessToken, tokenCfg.isValid, %s", err)
	}
	fmt.Println("token config valid")

	if c.TokenCfg.Grant != RefreshToken{
		data.Set("client_id", c.TokenCfg.ClientId)
		data.Set("client_secret", c.TokenCfg.ClientSecret)
	}
		
	switch c.TokenCfg.Grant{
		case Password:
			data.Set("password", c.TokenCfg.Password)
			data.Set("username", c.TokenCfg.Username)
		case Saml2Bearer:
			data.Set("assertion", c.TokenCfg.Assertion)
	}
	
	reqClient := &http.Client{}
	req, err := http.NewRequestWithContext(ctx, "POST", endpoint.String(), strings.NewReader(data.Encode()))
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	fmt.Println(req.URL.String(), req.Header)
	
	resp, err := reqClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	fmt.Println("fetched token")

	okResp := resp.StatusCode >= 200 || resp.StatusCode < 300
	if !okResp{
		errMessage := internal.HttpErrorHandler(resp)
		return fmt.Errorf("%s", errMessage)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("GetAccesstoken, readBody, %s", err)
	}

	var tokenResp Token
	if err := json.Unmarshal(body, &tokenResp); err != nil {
		return fmt.Errorf("GetAccessToken, unmarshal response, %s", err)
	}

	c.Token = &tokenResp
	
	return nil
}
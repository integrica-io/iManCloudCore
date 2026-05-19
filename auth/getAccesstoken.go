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

type GetAccessTokenResponse struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    string `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	TokenType    string `json:"token_type"`
}

type GetAccessTokenCfg struct{
	Assertion 		string
	ClientId		string
	ClientSecret	string
	CustomerId		string
	Password		string
	Username		string
}

type GrantType string

func (cfg *GetAccessTokenCfg) isValid(grant GrantType) error {
	if grant != RefreshToken{
		if cfg.ClientId == ""{
			return fmt.Errorf("client_id required for %s grant type", grant)
		}
		if cfg.ClientSecret == ""{
			return fmt.Errorf("client_secret required for %s grant type", grant)
		}
	}
	if grant == Saml2Bearer { 
		if cfg.CustomerId == ""{
			return fmt.Errorf("customer_id required for %s grant type", grant)
		}
		if cfg.Assertion == ""{
			return fmt.Errorf("assertion required for %s grant type", grant)
		}
	}
	if grant == Password{
		if cfg.Password == ""{
			return fmt.Errorf("password required for %s grant type", grant)	
		}
		if cfg.Username == ""{
			return fmt.Errorf("username required for %s grant type", grant)
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

func (client *Client) GetAccessToken(ctx context.Context, grant GrantType, cfg *GetAccessTokenCfg) (error) {
	endpoint := client.BaseUrl.JoinPath("auth","oauth2","token")

	data := url.Values{}
	
	if !grant.IsValid(){
		return fmt.Errorf("invalid grant type") 
	}

	if err := cfg.isValid(grant); err != nil {
		return fmt.Errorf("GetAccessToken, cfg is valid, %s", err)
	}

	if grant != RefreshToken{
		data.Set("client_id", cfg.ClientId)
		data.Set("client_secret", cfg.ClientSecret)
	}
		
	switch grant{
		case Password:
			data.Set("password", cfg.Password)
			data.Set("username", cfg.Username)
		case Saml2Bearer:
			data.Set("assertion", cfg.Assertion)
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

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("GetAccesstoken, readBody, %s", err)
	}

	var tokenResp GetAccessTokenResponse
	if err := json.Unmarshal(body, &tokenResp); err != nil {
		return fmt.Errorf("GetAccessToken, unmarshal response, %s", err)
	}

	client.AccessToken = tokenResp.AccessToken
	client.ExpiresIn = tokenResp.ExpiresIn
	client.RefreshToken = tokenResp.RefreshToken
	client.TokenType = tokenResp.TokenType
	
	return nil
}
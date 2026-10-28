package auth

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/integrica-io/iManCloudCore/client"
	"github.com/integrica-io/iManCloudCore/internal"
	"github.com/integrica-io/iManCloudCore/types"
)

func GetAccessToken(ctx context.Context, client *client.Client) error {
	endpoint := client.BaseUrl.JoinPath("auth", "oauth2", "token")
	data := url.Values{}

	data.Set("grant_type", string(client.TokenCfg.Grant))

	if client.TokenCfg.Grant == "" {
		return fmt.Errorf("empty grant type")
	}

	if err := client.TokenCfg.IsValid(); err != nil {
		return fmt.Errorf("GetAccessToken, tokenCfg.isValid, %s", err)
	}

	if client.TokenCfg.Grant != types.RefreshToken {
		data.Set("client_id", client.TokenCfg.ClientId)
		data.Set("client_secret", client.TokenCfg.ClientSecret)
	}

	switch client.TokenCfg.Grant {
	case types.Password:
		data.Set("password", client.TokenCfg.Password)
		data.Set("username", client.TokenCfg.Username)
	case types.Saml2Bearer:
		data.Set("assertion", client.TokenCfg.Assertion)
	}

	fmt.Println(data.Encode())
	fmt.Println(endpoint.String())
	dataReader := bytes.NewReader([]byte(data.Encode()))
	//req := internal.HttpRequestBuilder{}
	req, err := http.NewRequestWithContext(ctx, "POST", endpoint.String(), dataReader)
	if err != nil {
		return fmt.Errorf("new request: %w", err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("request do: %w", err)
	}
	defer resp.Body.Close()
	
	respOk := resp.StatusCode >= 200 && resp.StatusCode < 300
	if !respOk {
		response := *resp
		respReader, err := io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("resp reader: %w", err)
		}
		fmt.Printf("Status: %d, response: %s", resp.StatusCode, string(respReader))
		return internal.HttpErrorHandler(response)
	}
	
	var GetAccesstokenOutput types.Token

	respReader, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("resp reader: %w", err)
	}
	
	if err := json.Unmarshal(respReader, &GetAccesstokenOutput); err != nil {
		return fmt.Errorf("resp unmarshal: %w", err)
	}

	client.Token = &GetAccesstokenOutput
	client.Token.TokenExpiry = time.Now().Add(time.Second * time.Duration(client.Token.ExpiresIn))

	return nil
}

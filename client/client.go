package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
	"math/rand/v2"

	"github.com/integrica-io/iManCloudCore/internal"
	"github.com/integrica-io/iManCloudCore/types"
)

type Client struct {
	BaseUrl  *url.URL
	TokenCfg types.AccessTokenCfg
	Token    *types.Token
}

func NewClient(hostname string, clientCfg *types.AccessTokenCfg) (*Client, error) {
	url, err := url.Parse(fmt.Sprintf("https://%s", hostname))
	if err != nil {
		return nil, err
	}

	return &Client{
		BaseUrl:  url,
		TokenCfg: *clientCfg,
	}, nil
}

func (client *Client) Req(b internal.HttpRequestBuilder) error {
	reqCfg := internal.HttpRequestCfg{}
	reqClient := &http.Client{}

	// Validate required parameters
	if b.ReqCtx == nil {
		return fmt.Errorf("context required")
	}
	reqCfg.ReqCtx = *b.ReqCtx

	if b.ReqUrl == nil {
		return fmt.Errorf("Url required")
	}
	reqCfg.ReqUrl = *b.ReqUrl

	if b.ReqMethod == nil {
		return fmt.Errorf("Method required")
	}
	reqCfg.ReqMethod = *b.ReqMethod

	req, err := http.NewRequestWithContext(reqCfg.ReqCtx, string(reqCfg.ReqMethod), reqCfg.ReqUrl.String(), http.NoBody)
	if err != nil {
		return err
	}

	switch {
	case b.ReqForm != nil:
		req.Body = io.NopCloser(strings.NewReader(b.ReqForm.Encode()))
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	case b.ReqBody != nil:
		req.Body = io.NopCloser(bytes.NewReader(*b.ReqBody))
	}

	if b.ReqHeaders != nil {
		reqCfg.ReqHeaders = *b.ReqHeaders
		for i, v := range reqCfg.ReqHeaders {
			req.Header.Set(i, v)
		}
	}

	if client.Token != nil {
		if !time.Now().Before(client.Token.TokenExpiry) {
			if err := client.RefreshAccessToken(reqCfg.ReqCtx); err != nil {
				return err
			}
		}
		req.Header.Set("X-Auth-Token", client.Token.AccessToken)
	}

	var resp *http.Response
	defer resp.Body.Close()
 	for attempt := 1; attempt <= 5; attempt++ {
		resp, err := reqClient.Do(req)
		if err == nil {
			break
		}
		slog.Error("err","err", internal.HttpErrorHandler(resp))
		jitter()
		if attempt == 5{
			return internal.HttpErrorHandler(resp)
		}
 	}
  
	respOk := resp.StatusCode >= 200 && resp.StatusCode < 300
	if !respOk {
		return internal.HttpErrorHandler(resp)
	}

	if b.ReqToFile != nil {
		reqCfg.ReqToFile = *b.ReqToFile
		bytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		err = os.WriteFile(reqCfg.ReqToFile, bytes, 0644)
		if err != nil {
		    fmt.Printf("Error writing file: %v\n", err)
		}
		return nil
	}

	if b.ReqToJson != nil {
		reqCfg.ReqToJson = *b.ReqToJson
		bytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		if err := json.Unmarshal(bytes, &reqCfg.ReqToJson); err != nil {
			return err
		}
		return nil
	}
	return nil
}

func (client *Client) RefreshAccessToken(ctx context.Context) error {
	if client.Token == nil {
		return fmt.Errorf("refresh_token required to refresh access token")
	}

	endpoint := client.BaseUrl.JoinPath("auth", "oauth2", "token")
	data := url.Values{}
	data.Set("grant_type", "refresh_token")
	data.Set("refresh_token", client.Token.RefreshToken)

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

func jitter(){
	time.Sleep(time.Millisecond*time.Duration(randRange(1, 25)))
}

func randRange(min, max int) int {
    return rand.IntN(max-min) + min
}

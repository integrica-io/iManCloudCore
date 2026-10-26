package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"math/rand/v2"
	"net/http"
	"net/url"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/integrica-io/iManCloudCore/internal"
	"github.com/integrica-io/iManCloudCore/types"
)

type Client struct {
	BaseUrl  *url.URL
	TokenCfg types.AccessTokenCfg
	Token    *types.Token
	mu		 sync.Mutex
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
	reqClient := &http.Client{}

	if b.ReqCtx == nil || *b.ReqCtx == nil {
		return fmt.Errorf("context required")
	}
	ctx := *b.ReqCtx

	if b.ReqUrl == nil {
		return fmt.Errorf("Url required")
	}

	if b.ReqMethod == nil {
		return fmt.Errorf("Method required")
	}
	
	var bodyReader io.Reader = http.NoBody
	var contentType string

	switch {
	case b.ReqForm != nil:
		bodyReader = strings.NewReader(b.ReqForm.Encode())
		contentType = "application/x-www-form-urlencoded"
	case b.ReqBody != nil:
		bodyReader = bytes.NewReader(*b.ReqBody)
	}

	req, err := http.NewRequestWithContext(ctx, string(*b.ReqMethod), b.ReqUrl.String(), bodyReader)
		if err != nil {
			return fmt.Errorf("creating request: %w", err)
	}

	if contentType != "" {
			req.Header.Set("Content-Type", contentType)
	}

	if b.ReqHeaders != nil {
		for i, v := range *b.ReqHeaders {
			req.Header.Set(i, v)
		}
	}

	if client.Token != nil {
		client.mu.Lock()
		if !time.Now().Before(client.Token.TokenExpiry) {
			if err := client.RefreshAccessToken(ctx); err != nil {
				client.mu.Unlock()
				return fmt.Errorf("refreshing access token: %w", err)
			}
		}
		client.mu.Unlock()
		req.Header.Set("X-Auth-Token", client.Token.AccessToken)
	}
	
	maxAttempts := 5
	var resp *http.Response
 	for attempt := 1; attempt <= maxAttempts; attempt++ {
		resp, err = reqClient.Do(req)
		if err == nil {
			break
		}
		fmt.Println("request attempt failed", "attempt", attempt, "err", err)
		if attempt == maxAttempts{
			return fmt.Errorf("request failed %w, attempts: %d", err, maxAttempts)
		}

		if ctx.Err() != nil {
			return fmt.Errorf("request aborted by context: %w", ctx.Err())
		}
		jitter()
 	}
  	defer resp.Body.Close()
  
	respOk := resp.StatusCode >= 200 && resp.StatusCode < 300
	if !respOk {
		return internal.HttpErrorHandler(resp)
	}

	if b.ReqToFile != nil {
		f, err := os.OpenFile(*b.ReqToFile, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
		if err != nil {
			return fmt.Errorf("opening file: %w", err)
		}
		defer f.Close()

		if _, err := io.Copy(f, resp.Body); err != nil {
			return fmt.Errorf("writing response to file: %w", err)
		}
		return f.Close()
	}

	if b.ReqToJson != nil {
		data := *b.ReqToJson
		bytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("reading response: %w", err)
		}
		if err := json.Unmarshal(bytes, data); err != nil {
			return fmt.Errorf("unmarshalling response: %w", err)
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
	time.Sleep(time.Millisecond*time.Duration(randRange(100, 1000)))
}

func randRange(min, max int) int {
    return rand.IntN(max-min) + min
}
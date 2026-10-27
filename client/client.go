package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"math/rand"
	"net/http"
	"net/url"
	"os"
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
	if b.ReqCtx == nil || *b.ReqCtx == nil {
		return fmt.Errorf("context required")
	}
	ctx := *b.ReqCtx

	if b.ReqUrl == nil {
		return fmt.Errorf("url required")
	}

	if b.ReqMethod == nil {
		return fmt.Errorf("method required")
	}

	var bodyBytes []byte
	var contentType string

	switch {
	case b.ReqForm != nil:
		bodyBytes = []byte(b.ReqForm.Encode())
		contentType = "application/x-www-form-urlencoded"
	case b.ReqBody != nil:
		bodyBytes = *b.ReqBody
	}

	var accessToken string
	client.mu.Lock()
	if client.Token != nil {
		if !time.Now().Before(client.Token.TokenExpiry) {
			if err := client.RefreshAccessToken(ctx); err != nil {
				client.mu.Unlock()
				return fmt.Errorf("refreshing access token: %w", err)
			}
		}
		accessToken = client.Token.AccessToken
	}
	client.mu.Unlock()
	
	reqClient := &http.Client{}
	maxAttempts := 5
	var resp *http.Response
 	for attempt := 1; attempt <= maxAttempts; attempt++ {
  		var bodyReader io.Reader = http.NoBody
    	if len(bodyBytes) > 0 {
     		bodyReader = bytes.NewReader(bodyBytes)
     	}
  
		req, err := http.NewRequestWithContext(ctx, string(*b.ReqMethod), b.ReqUrl.String(), bodyReader)
		if err != nil {
			return fmt.Errorf("creating request: %w", err)
		}

		if accessToken != "" {
			req.Header.Set("X-Auth-Token", accessToken)
		}
				
		if contentType != "" {
				req.Header.Set("Content-Type", contentType)
		}
		
		if b.ReqHeaders != nil {
			for i, v := range *b.ReqHeaders {
				req.Header.Set(i, v)
			}
		}
		
		resp, err = reqClient.Do(req)
		if err == nil {
			break
		}
  	
		if attempt == maxAttempts {
        	return fmt.Errorf("retries exhausted: %w", err)
        }
	
		slog.Error("request attempt failed, attempt: ", attempt, "err: ", err)

		delay := jitter(backoff(attempt))
		
		select {
        case <-time.After(delay):
        case <-ctx.Done():
            return ctx.Err()
        }
 	}
  
  	defer func(){
   		_, _ = io.Copy(io.Discard, resp.Body)
   		resp.Body.Close()
   	}()
   
	respOk := resp.StatusCode >= 200 && resp.StatusCode < 300
	if !respOk {
		response := *resp
		return internal.HttpErrorHandler(response)
	}

	switch {
		case b.ReqToFile != nil:
			if err := requestToFile(*b.ReqToFile, resp.Body); err != nil {
				return fmt.Errorf("sending request to file: %w", err)
			}
		case b.ReqToJson != nil:
			if err := requestToJson(*b.ReqToJson, resp.Body); err != nil{
				return fmt.Errorf("unmarshalling to struct: %w", err)
			}	
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

	dataReader := bytes.NewReader([]byte(data.Encode()))
	req, err := http.NewRequestWithContext(ctx, "POST", endpoint.String(), dataReader)
	if err != nil {
		return fmt.Errorf("new request: %w", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("request do: %w", err)
	}
	resp.Body.Close()

	respOk := resp.StatusCode >= 200 && resp.StatusCode < 300
	if !respOk {
		response := *resp
		return internal.HttpErrorHandler(response)
	}

	respReader, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("read response: %w", err)
	}
	
	var GetAccesstokenOutput types.Token
	if err := json.Unmarshal(respReader, &GetAccesstokenOutput); err != nil {
		return fmt.Errorf("unmarshal response: %w", err)
	}

	client.Token = &GetAccesstokenOutput
	client.Token.TokenExpiry = time.Now().Add(time.Second * time.Duration(client.Token.ExpiresIn))

	return nil
}

func requestToJson(structPointer any, body io.Reader) error {
	bytes, err := io.ReadAll(body)
	if err != nil {
		return fmt.Errorf("reading response: %w", err)
	}
	if err := json.Unmarshal(bytes, structPointer); err != nil {
		return fmt.Errorf("unmarshalling response: %w", err)
	}
	return nil
}

func requestToFile(filename string, body io.Reader) error {

	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("opening file: %w", err)
	}
	defer f.Close()

	if _, err := io.Copy(f, body); err != nil {
		return fmt.Errorf("writing response to file: %w", err)
	}
	return f.Close()
}

func jitter(d time.Duration) time.Duration {
    return time.Duration(rand.Int63n(int64(d)))
}

func backoff(attempt int) time.Duration {
    base := 100 * time.Millisecond
    max := 5 * time.Second

    d := time.Duration(1<<attempt) * base
    if d > max {
        return max
    }
    return d
}
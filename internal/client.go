package internal

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"iManCloudCore/types"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Client struct {
	BaseUrl 	 *url.URL
	TokenCfg 	 types.AccessTokenCfg
	Token	 	 *types.Token	
}

func NewClient(hostname string, clientCfg *types.AccessTokenCfg)(*Client, error){
	url, err := url.Parse(fmt.Sprintf("https://%s", hostname))
	if err != nil {
		return nil, err
	}
	
	return &Client{
		BaseUrl: url,
		TokenCfg: *clientCfg,
	}, nil
}

func (client *Client) Req(b HttpRequestBuilder) error {
	reqCfg := HttpRequestCfg{}
	reqClient := &http.Client{}
	
	//Validate required parameters
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
	if err != nil{
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
		for i, v := range reqCfg.ReqHeaders{
			req.Header.Set(i, v)
		}
	}

	if client.Token != nil {
		if !time.Now().Before(client.Token.TokenExpiry){
			if err := client.RefreshAccessToken(reqCfg.ReqCtx); err != nil {
				return err
			}			
		}
	}

 	resp, err := reqClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	respOk := resp.StatusCode >= 200 && resp.StatusCode < 300
	if !respOk {
		return HttpErrorHandler(resp)
	}

	if b.ReqToJson != nil {
		reqCfg.ReqToJson = b.ReqToJson
		bytes, err  := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		if err := json.Unmarshal(bytes, &reqCfg.ReqToJson); err != nil {
			return err
		}
	}
	return nil
}

func (client *Client) RefreshAccessToken(ctx context.Context) (error) {

	if client.Token == nil {
		return fmt.Errorf("refresh_token required to refresh access token")
	}
	
	endpoint := client.BaseUrl.JoinPath("auth","oauth2","token")
	data := url.Values{}
	data.Set("grant_type","refresh_token")
	
	req := HttpRequestBuilder{}

	var GetAccesstokenOutput types.Token

	req.Context(ctx).Url(*endpoint).Method(Post).ToJson(&GetAccesstokenOutput).Form(data)

	if err := client.Req(req); err != nil {
		return err
	}
	
	client.Token = &GetAccesstokenOutput
	client.Token.TokenExpiry = time.Now().Add(time.Second * time.Duration(client.Token.ExpiresIn))
	
	return nil
}
package internal

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type HttpMethod string

const (
	Get		HttpMethod = "GET"
	Post 	HttpMethod = "POST"
	Put		HttpMethod = "PUT"
	Patch	HttpMethod = "PATCH"
	Delete	HttpMethod = "DELETE"
)

type HttpRequestCfg struct {
	ReqUrl 		url.URL
	ReqMethod	HttpMethod
	ReqCtx		context.Context
	ReqHeaders 	map[string]string
	ReqBody 	[]byte
	ReqForm		url.Values
	ReqToJson  	any
}

type HttpRequestBuilder struct {
	ReqUrl	 	*url.URL
	ReqMethod 	*HttpMethod
	ReqCtx 		*context.Context
	ReqHeaders 	*map[string]string
	ReqBody		*[]byte
	ReqForm		*url.Values
	ReqToJson	*any
}

func (b *HttpRequestBuilder) Url(url url.URL) *HttpRequestBuilder{
	b.ReqUrl = &url
	return b
}

func (b *HttpRequestBuilder) Method(method HttpMethod) *HttpRequestBuilder{
	b.ReqMethod = &method
	return b
}

func (b *HttpRequestBuilder) Context(ctx context.Context) *HttpRequestBuilder{
	b.ReqCtx = &ctx
	return b
}

func(b *HttpRequestBuilder) Headers(headers map[string]string) *HttpRequestBuilder{
	b.ReqHeaders = &headers
	return b
}

func (b *HttpRequestBuilder) Body(body []byte) *HttpRequestBuilder{
	b.ReqBody = &body
	return b
}

func (b *HttpRequestBuilder) Form(form url.Values) *HttpRequestBuilder {
	b.ReqForm = &form
	return b
}

func (b *HttpRequestBuilder) ToJson(pointer any) *HttpRequestBuilder {
	b.ReqToJson = &pointer
	return b
}

func (b *HttpRequestBuilder) Exec() error {
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
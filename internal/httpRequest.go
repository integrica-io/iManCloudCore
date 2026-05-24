package internal

import (
	"context"
	"net/url"
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
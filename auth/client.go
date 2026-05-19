package auth

import (
	"fmt"
	"net/url"
)

type Client struct {
	BaseUrl 	 *url.URL
	TokenCfg 	 AccessTokenCfg
	Token	 	 *Token	
}

func NewClient(hostname string, tokenCfg *AccessTokenCfg)(*Client, error){
	url, err := url.Parse(fmt.Sprintf("https://%s", hostname))
	if err != nil {
		return nil, err
	}
	
	return &Client{
		BaseUrl: url,
		TokenCfg: *tokenCfg,
	}, nil
}
package internal

import (
	"fmt"
	"net/url"
	"iManCloudCore/types"
)

type Client struct {
	BaseUrl 	 *url.URL
	TokenCfg 	 types.AccessTokenCfg
	Token	 	 *types.Token	
}

func NewClient(hostname string, tokenCfg *types.AccessTokenCfg)(*Client, error){
	url, err := url.Parse(fmt.Sprintf("https://%s", hostname))
	if err != nil {
		return nil, err
	}
	
	return &Client{
		BaseUrl: url,
		TokenCfg: *tokenCfg,
	}, nil
}
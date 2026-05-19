package auth

import (
	"fmt"
	"net/url"
)

type Client struct {
	BaseUrl 	 *url.URL
	CustomerNo	 string
	AccessToken  string 
	ExpiresIn    string 
	RefreshToken string 
	TokenType    string 
}

func NewClient(hostname string, customerNo string)(*Client, error){
	url, err := url.Parse(fmt.Sprintf("https://%s", hostname))
	if err != nil {
		return nil, err
	}
	
	return &Client{
		BaseUrl: url,
		CustomerNo: customerNo,
	}, nil
}
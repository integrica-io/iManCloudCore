
package customers

import (
	"context"
	"iManCloudCore/internal"
)

type GetAccessTokenOutput struct {
	Data struct {
		AccessToken string `json:"access_token"`
	} `json:"data"`
}

func GetAccessToken(ctx context.Context, client *internal.Client, dialogToken string)(GetAccessTokenOutput, error){

	var data GetAccessTokenOutput
	endpoint := client.BaseUrl.JoinPath("work","api","v2","customers",client.TokenCfg.CustomerId, "dialog-tokens", dialogToken)

	req := internal.HttpRequestBuilder{}
	req.Context(ctx).Url(*endpoint).Method(internal.Get).ToJson(&data)

	if err := client.Req(req); err != nil {
		return data, err
	}
	
	return data, nil
}


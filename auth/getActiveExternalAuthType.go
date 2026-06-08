package auth

import (
	"context"
	"github.com/integrica-io/iManCloudCore/internal"
)

type GetActiveExternalAuthTypeOutput struct {
	Data struct {
		ID string `json:"id"`
	} `json:"data"`
}

func GetActiveExternalAuthType(ctx context.Context, client *internal.Client)(GetActiveExternalAuthTypeOutput, error){
	endpoint := client.BaseUrl.JoinPath("platform", "api", "v2","customers",client.TokenCfg.CustomerId, "settings", "system", "auth", "external-auth-type")
	var data GetActiveExternalAuthTypeOutput

	req := internal.HttpRequestBuilder{}
	req.Context(ctx).Url(*endpoint).Method(internal.Get).ToJson(&data)

	if err := client.Req(req); err != nil {
		return data, err
	}
	return data, nil
}
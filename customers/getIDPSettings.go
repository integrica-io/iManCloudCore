package customers

import (
	"context"
	"iManCloudCore/internal"

	"github.com/google/go-querystring/query"	
)

type GetIDPSettingsOutput struct {
	Data struct {
		SeamlessIdpEnabled bool `json:"seamless_idp_enabled"`
		LoginHintEnabled   bool `json:"login_hint_enabled"`
	} `json:"data"`
}

type GetIDPSettingsOptions struct{
	Cursor string `url:"cursor,omitempty"`
	Limit int `url:"limit,omitempty"`
}

func GetIDPSettings(ctx context.Context, client *internal.Client, options *GetIDPSettingsOptions)(GetIDPSettingsOutput, error){
	var data GetIDPSettingsOutput
	endpoint := client.BaseUrl.JoinPath("platform","api","v2","customers",client.TokenCfg.CustomerId, "settings", "system", "idp")

	if options != nil {
		values, err := query.Values(options)
		if err != nil {
			return data, err
		}
		endpoint.RawQuery = values.Encode()
	}

	req := internal.HttpRequestBuilder{}
	req.Context(ctx).Url(*endpoint).Method(internal.Get).ToJson(&data)

	if err := client.Req(req); err != nil {
		return data, err
	}
	return data, nil
}
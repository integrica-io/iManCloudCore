package customers

import (
	"context"
	"github.com/integrica-io/iManCloudCore/internal"
	"github.com/integrica-io/iManCloudCore/client"

	"github.com/google/go-querystring/query"	
)

type GetSecurityTemplatesOutput struct {
	Data []struct {
		ID              string `json:"id"`
		Name            string `json:"name"`
		Database        string `json:"database"`
		Description     string `json:"description"`
		DefaultSecurity string `json:"default_security"`
	} `json:"data"`
	TotalCount int `json:"total_count"`
}

type GetSecurityTemplatesOptions struct{
	Limit int `url:"limit,omitempty"`
	Offset int `url:"Offset,omitempty"`
	Total int `url:"total,omitempty"`
}

func GetSecurityTemplates(ctx context.Context, client *client.Client, options *GetSecurityTemplatesOptions)(GetSecurityTemplatesOutput, error){
	var data GetSecurityTemplatesOutput
	endpoint := client.BaseUrl.JoinPath("work","api","v2","customers",client.TokenCfg.CustomerId, "security-templates")

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
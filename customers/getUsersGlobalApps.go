package customers

import (
	"context"
	"github.com/integrica-io/iManCloudCore/internal"
	"github.com/integrica-io/iManCloudCore/client"
	"github.com/google/go-querystring/query"	
)

type GetUsersGlobalAppsOutput struct {
	Data []struct {
		ID string `json:"id"`
	} `json:"data"`
}

type GetUsersGlobalAppsOptions struct {
    ClientId string `url:"client_id,omitempty"`
    Limit int `url:"limit,omitempty"`
    Offset int `url:"offset,omitempty"`
    Total bool `url:"total,omitempty"`
}

func GetUsersGlobalApps(ctx context.Context, client *client.Client, userId string, options *GetUsersGlobalAppsOptions)(GetUsersGlobalAppsOutput, error){

	var data GetUsersGlobalAppsOutput
	endpoint := client.BaseUrl.JoinPath("work","api","v2","customers",client.TokenCfg.CustomerId, "users", userId, "apps")

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
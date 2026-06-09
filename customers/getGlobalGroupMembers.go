package customers

import (
	"context"
	"github.com/integrica-io/iManCloudCore/internal"
	"github.com/integrica-io/iManCloudCore/client"

	"github.com/google/go-querystring/query"	
)

type GetGlobalGroupMembersOutput struct {
	Data []struct {
		DirectoryID string `json:"directory_id"`
	} `json:"data"`
	Cursor string `json:"cursor"`
}

type GetGlobalGroupMembersOptions struct{
	Cursor string `url:"cursor,omitempty"`
	Limit int `url:"limit,omitempty"`
}

func GetGlobalGroupMembers(ctx context.Context, client *client.Client, customerId string, options *GetGlobalGroupMembersOptions)(GetGlobalGroupMembersOutput, error){

	var data GetGlobalGroupMembersOutput
	endpoint := client.BaseUrl.JoinPath("work","api","v2","customers",client.TokenCfg.CustomerId, "directory-sync", "groups", customerId, "members")

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
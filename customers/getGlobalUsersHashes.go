package customers

import (
	"context"
	"github.com/integrica-io/iManCloudCore/internal"
	"github.com/integrica-io/iManCloudCore/client"

	"github.com/google/go-querystring/query"	
)

type GetGlobalUsersHashesOutput struct {
	Data []struct {
		DirectoryID string `json:"directory_id"`
		DsHash      string `json:"ds_hash"`
		AllowLogon  bool   `json:"allow_logon"`
	} `json:"data"`
	Cursor string `json:"cursor"`
}

type GetGlobalUsersHashesOptions struct{
	Cursor string `url:"cursor,omitempty"`
	Limit int `url:"limit,omitempty"`
}

func GetGlobalUsersHashes(ctx context.Context, client *client.Client, options *GetGlobalRolesOptions)(GetGlobalUsersHashesOutput, error){

	var data GetGlobalUsersHashesOutput
	endpoint := client.BaseUrl.JoinPath("work","api","v2","customers",client.TokenCfg.CustomerId, "directory-sync", "users", "hashes")

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
package customers

import (
	"context"
	"github.com/integrica-io/iManCloudCore/internal"
	"github.com/integrica-io/iManCloudCore/client"

	"github.com/google/go-querystring/query"	
)

type GetGlobalGroupsHashesOutput struct {
	Data []struct {
		DirectoryID   string `json:"directory_id"`
		DsHash        string `json:"ds_hash"`
		DsMembersHash string `json:"ds_members_hash"`
		Enabled       bool   `json:"enabled"`
	} `json:"data"`
	Cursor string `json:"cursor"`
}

type GetGlobalGroupsHashesOptions struct{
	Cursor string `url:"cursor,omitempty"`
	Limit int `url:"limit,omitempty"`
}

func GetGlobalGroupsHashes(ctx context.Context, client *client.Client, options *GetGlobalGroupsHashesOptions)(GetGlobalGroupsHashesOutput, error){

	var data GetGlobalGroupsHashesOutput
	endpoint := client.BaseUrl.JoinPath("work","api","v2","customers",client.TokenCfg.CustomerId, "directory-sync", "groups", "hashes")

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
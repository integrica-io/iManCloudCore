package customers

import (
	"context"
	"github.com/integrica-io/iManCloudCore/internal"
	"github.com/integrica-io/iManCloudCore/client"
	"github.com/google/go-querystring/query"
)

type GetLibrariesUserAssignedOutput struct {
	Data []struct {
		ID        string `json:"id"`
		LibraryID int    `json:"library_id"`
		Preferred bool   `json:"preferred"`
	} `json:"data"`
	TotalCount int `json:"total_count"`
}

type GetLibrariesUserAssignedOptions struct{
	Limit 	int 	`url:"limit"`
	Offset	int 	`url:"offset"`
	Total 	bool	`url:"total"`
}

func GetLibrariesUserAssigned(ctx context.Context, client *client.Client, userId string, options *GetInfoGlobalUserOptions)(GetLibrariesUserAssignedOutput, error){
	var data GetLibrariesUserAssignedOutput
	endpoint := client.BaseUrl.JoinPath("work","api","v2","customers",client.TokenCfg.CustomerId, "users", userId, "libraries")

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
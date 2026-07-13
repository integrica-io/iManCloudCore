package library

import (
	"context"

	"github.com/google/go-querystring/query"
	"github.com/integrica-io/iManCloudCore/client"
	"github.com/integrica-io/iManCloudCore/internal"
)

func GetEmailRecipientsFiledWorkspace(ctx context.Context, client *client.Client, libraryId string, workspaceId string, options *GetEmailRecipientsFiledWorkspaceOptions) (GetEmailRecipientsFiledWorkspaceOutput, error) {
	var data GetEmailRecipientsFiledWorkspaceOutput
	endpoint := client.BaseUrl.JoinPath("work", "api", "v2", "customers", client.TokenCfg.CustomerId, "libraries", libraryId, "workspaces", workspaceId, "contacts")

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

type GetEmailRecipientsFiledWorkspaceOutput struct {
	Data []struct {
		Name    string `json:"name"`
		Address string `json:"address"`
	} `json:"data"`
}

type GetEmailRecipientsFiledWorkspaceOptions struct {
	Limit      int    `url:"limit,omitempty"`
	Query      string `url:"query,omitempty"`
	SenderOnly bool   `url:"sender_only,omitempty"`
}

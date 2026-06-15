package library

import (
	"context"

	"github.com/integrica-io/iManCloudCore/client"
	"github.com/integrica-io/iManCloudCore/internal"
)

func GetWorkspaceSecurityPolicy(ctx context.Context, client client.Client, libraryId string, workspaceId string) (GetWorkspaceSecurityPolicyOutput, error) {
	var data GetWorkspaceSecurityPolicyOutput
	endpoint := client.BaseUrl.JoinPath("work", "api", "v2", "customers", client.TokenCfg.CustomerId, "libraries", libraryId, "workspaces", workspaceId, "security")

	req := internal.HttpRequestBuilder{}
	req.Context(ctx).Url(*endpoint).Method(internal.Get).ToJson(&data)

	if err := client.Req(req); err != nil {
		return data, err
	}
	return data, nil
}

type GetWorkspaceSecurityPolicyOutput struct {
	Data struct {
		AllowList []struct {
			FullName string `json:"full_name"`
			ID       string `json:"id"`
		} `json:"allow_list"`
		DenyList []struct {
			FullName string `json:"full_name"`
			ID       string `json:"id"`
		} `json:"deny_list"`
		SecurityPolicy string `json:"security_policy"`
	} `json:"data"`
}
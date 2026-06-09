package customers

import (
	"context"
	"github.com/integrica-io/iManCloudCore/internal"
	"github.com/integrica-io/iManCloudCore/client"

	"github.com/google/go-querystring/query"	
)

type GetGlobalRolesOutput struct {
	Data []struct {
		ID                   string `json:"id"`
		Description          string `json:"description"`
		AppManagement        string `json:"app_management"`
		EncryptionManagement string `json:"encryption_management"`
		FeatureManagement    string `json:"feature_management"`
		GroupManagement      string `json:"group_management"`
		RoleManagement       string `json:"role_management"`
		SettingsManagement   string `json:"settings_management"`
		UserManagement       string `json:"user_management"`
	} `json:"data"`
	TotalCount int `json:"total_count"`
}

type GetGlobalRolesOptions struct{
	Limit int `url:"limit,omitempty"`
	Offset int `url:"offset,omitempty"`
	Query string `url:"query,omitempty"`
	QueryMatchType QueryMatchType `url:"query_match_type,omitempty"`
	Total	bool `url:"total"`
}

func GetGlobalRoles(ctx context.Context, client *client.Client, options *GetGlobalRolesOptions)(GetGlobalRolesOutput, error){

	var data GetGlobalRolesOutput
	endpoint := client.BaseUrl.JoinPath("work","api","v2","customers",client.TokenCfg.CustomerId, "roles")

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
package customers

import (
	"context"
	"github.com/integrica-io/iManCloudCore/internal"
	"github.com/integrica-io/iManCloudCore/client"
)

type GetDetailsGlobalRoleOutput struct {
	Data struct {
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
}

func GetDetailsGlobalRole(ctx context.Context, client *client.Client, roleId string)(GetDetailsGlobalRoleOutput, error){
	var data GetDetailsGlobalRoleOutput
	endpoint := client.BaseUrl.JoinPath("work","api","v2","customers",client.TokenCfg.CustomerId, "roles", roleId)

	req := internal.HttpRequestBuilder{}
	req.Context(ctx).Url(*endpoint).Method(internal.Get).ToJson(&data)

	if err := client.Req(req); err != nil {
		return data, err
	}
	return data, nil
}
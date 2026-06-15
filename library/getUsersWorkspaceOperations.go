package library

import (
	"context"
	
	"github.com/integrica-io/iManCloudCore/internal"
	"github.com/integrica-io/iManCloudCore/client"
)

func GetUsersWorkspaceOperations(ctx context.Context, client client.Client, libraryId string, workspaceId string)(GetUsersWorkspaceOperationsOutput, error){
	var data GetUsersWorkspaceOperationsOutput
	endpoint := client.BaseUrl.JoinPath("work","api","v2","customers",client.TokenCfg.CustomerId, "libraries", libraryId, "workspaces", workspaceId, "operations")

	req := internal.HttpRequestBuilder{}
	req.Context(ctx).Url(*endpoint).Method(internal.Get).ToJson(&data)

	if err := client.Req(req); err != nil {
		return data, err
	}
	return data, nil
}

type GetUsersWorkspaceOperationsOutput struct {
	Data struct {
		AddContent            bool     `json:"add_content"`
		AddFolders            bool     `json:"add_folders"`
		AddCustomFolders      bool     `json:"add_custom_folders"`
		ChangeName            bool     `json:"change_name"`
		Modify                bool     `json:"modify"`
		Remove                bool     `json:"remove"`
		RemoveFolders         bool     `json:"remove_folders"`
		SetDefaultSecurity    bool     `json:"set_default_security"`
		SetSecurity           bool     `json:"set_security"`
		UpdateContentShortcut bool     `json:"update_content_shortcut"`
		DefaultSecurityValues []string `json:"default_security_values"`
	} `json:"data"`
}
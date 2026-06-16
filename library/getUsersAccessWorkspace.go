package library

import (
	"context"
	
	"github.com/integrica-io/iManCloudCore/internal"
	"github.com/integrica-io/iManCloudCore/client"
)

func GetUsersAccessWorkspace(ctx context.Context, client *client.Client, libraryId string, workspaceId string, userId string)(GetUsersAccessWorkspaceOutput, error){
	var data GetUsersAccessWorkspaceOutput
	endpoint := client.BaseUrl.JoinPath("work","api","v2","customers",client.TokenCfg.CustomerId, "libraries", libraryId, "workspaces", workspaceId, "users", userId, "security")

	req := internal.HttpRequestBuilder{}
	req.Context(ctx).Url(*endpoint).Method(internal.Get).ToJson(&data)

	if err := client.Req(req); err != nil {
		return data, err
	}
	return data, nil
}

type GetUsersAccessWorkspaceOutput struct {
	Data struct {
		Access string `json:"access"`
	} `json:"data"`
}
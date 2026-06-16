package library

import (
	"context"
	"time"
	
	"github.com/integrica-io/iManCloudCore/internal"
	"github.com/integrica-io/iManCloudCore/client"
)

func GetWorkspaceDefaultEmailFolder(ctx context.Context, client *client.Client, libraryId string, workspaceId string)(GetWorkspaceDefaultEmailFolderOutput, error){
	var data GetWorkspaceDefaultEmailFolderOutput
	endpoint := client.BaseUrl.JoinPath("work","api","v2","customers",client.TokenCfg.CustomerId, "libraries", libraryId, "workspaces", workspaceId, "default-email-folder")

	req := internal.HttpRequestBuilder{}
	req.Context(ctx).Url(*endpoint).Method(internal.Get).ToJson(&data)

	if err := client.Req(req); err != nil {
		return data, err
	}
	return data, nil
}

type GetWorkspaceDefaultEmailFolderOutput struct {
	Data struct {
		Database               string    `json:"database"`
		DefaultSecurity        string    `json:"default_security"`
		EditDate               time.Time `json:"edit_date"`
		FolderType             string    `json:"folder_type"`
		HasDocuments           bool      `json:"has_documents"`
		HasSubfolders          bool      `json:"has_subfolders"`
		ID                     string    `json:"id"`
		IsContainerSavedSearch bool      `json:"is_container_saved_search"`
		IsContentSavedSearch   bool      `json:"is_content_saved_search"`
		IsExternalAsNormal     bool      `json:"is_external_as_normal"`
		Name                   string    `json:"name"`
		Owner                  string    `json:"owner"`
		OwnerDescription       string    `json:"owner_description"`
		ParentID               string    `json:"parent_id"`
		ViewType               string    `json:"view_type"`
		WorkspaceID            string    `json:"workspace_id"`
		WorkspaceName          string    `json:"workspace_name"`
		Wstype                 string    `json:"wstype"`
	} `json:"data"`
}
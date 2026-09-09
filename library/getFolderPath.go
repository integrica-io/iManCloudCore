package library

import (
	"context"
	"time"

	"github.com/integrica-io/iManCloudCore/client"
	"github.com/integrica-io/iManCloudCore/internal"
)

func GetFolderPath(ctx context.Context, client *client.Client, libraryId string, folderId string) (GetFolderPathOutput, error) {
	var data GetFolderPathOutput
	endpoint := client.BaseUrl.JoinPath("work", "api", "v2", "customers", client.TokenCfg.CustomerId, "libraries", libraryId, "folders", folderId, "path")

	req := internal.HttpRequestBuilder{}
	req.Context(ctx).Url(*endpoint).Method(internal.Get).ToJson(&data)

	if err := client.Req(req); err != nil {
		return data, err
	}
	return data, nil
}

type GetFolderPathOutput struct {
	Data []struct {
		CreateDate               time.Time `json:"create_date,omitempty"`
		Database                 string    `json:"database"`
		DefaultSecurity          string    `json:"default_security"`
		EditDate                 time.Time `json:"edit_date"`
		HasSubfolders            bool      `json:"has_subfolders"`
		ID                       string    `json:"id"`
		IsExternal               bool      `json:"is_external"`
		IsExternalAsNormal       bool      `json:"is_external_as_normal"`
		LastUser                 string    `json:"last_user,omitempty"`
		LastUserDescription      string    `json:"last_user_description,omitempty"`
		Name                     string    `json:"name"`
		Owner                    string    `json:"owner"`
		OwnerDescription         string    `json:"owner_description"`
		DocumentNumber           int64     `json:"document_number,omitempty"`
		IsDeclared               bool      `json:"is_declared,omitempty"`
		IsHipaa                  bool      `json:"is_hipaa,omitempty"`
		Iwl                      string    `json:"iwl,omitempty"`
		RetainDays               int       `json:"retain_days,omitempty"`
		Version                  int       `json:"version,omitempty"`
		Wstype                   string    `json:"wstype"`
		FolderType               string    `json:"folder_type,omitempty"`
		HasDocuments             bool      `json:"has_documents,omitempty"`
		InheritedDefaultSecurity string    `json:"inherited_default_security,omitempty"`
		IsContainerSavedSearch   bool      `json:"is_container_saved_search,omitempty"`
		IsContentSavedSearch     bool      `json:"is_content_saved_search,omitempty"`
		ParentID                 string    `json:"parent_id,omitempty"`
		ViewType                 string    `json:"view_type,omitempty"`
		WorkspaceName            string    `json:"workspace_name,omitempty"`
		WorkspaceID              string    `json:"workspace_id,omitempty"`
	} `json:"data"`
}

package multiLibrary

import (
	"time"
	"context"
	"iManCloudCore/internal"

	"github.com/google/go-querystring/query"	
)

type GetGlobalFoldersOutput struct {
	Data struct {
		Results []struct {
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
		} `json:"results"`
	} `json:"data"`
	TotalCount int `json:"total_count"`
}

type GetGlobalFoldersOptions struct {
    Description string `url:"description,omitempty"`
    Email string `url:"email,omitempty"`
    Libraries string `url:"libraries,omitempty"`
    Limit int `url:"limit,omitempty"`
    Name string `url:"name,omitempty"`
    NameOrDescription string `url:"name_or_description,omitempty"`
    Offset int `url:"offset,omitempty"`
    Owner string `url:"owner,omitempty"`
    ReferenceDatabase string `url:"reference_database,omitempty"`
    TimezoneOffset int `url:"timezone_offset,omitempty"`
    Total bool `url:"total,omitempty"`
}

func GetGlobalFolders(ctx context.Context, client *internal.Client, customerId string, options *GetGlobalFoldersOptions)(GetGlobalFoldersOutput, error){

	var data GetGlobalFoldersOutput
	endpoint := client.BaseUrl.JoinPath("work","api","v2","customers",client.TokenCfg.CustomerId, "folders")

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
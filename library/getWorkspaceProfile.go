package library

import (
	"time"
	"context"
	
	"github.com/integrica-io/iManCloudCore/internal"
	"github.com/integrica-io/iManCloudCore/client"
	"github.com/google/go-querystring/query"	
)

func GetWorkspaceProfile(ctx context.Context, client client.Client, libraryId string, workspaceId string, options *GetWorkspaceProfileOptions)(GetWorkspaceProfileOutput, error){
	var data GetWorkspaceProfileOutput
	endpoint := client.BaseUrl.JoinPath("work","api","v2","customers",client.TokenCfg.CustomerId, "libraries", libraryId, "workspaces", workspaceId)

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

type GetWorkspaceProfileOutput struct {
	Data struct {
		Class                  string    `json:"class"`
		ContentType            string    `json:"content_type"`
		CreateDate             time.Time `json:"create_date"`
		Custom1                string    `json:"custom1"`
		Custom1Description     string    `json:"custom1_description"`
		Custom2                string    `json:"custom2"`
		Custom2Description     string    `json:"custom2_description"`
		Custom3                string    `json:"custom3"`
		Custom3Description     string    `json:"custom3_description"`
		Custom4                string    `json:"custom4"`
		Custom4Description     string    `json:"custom4_description"`
		Database               string    `json:"database"`
		DefaultSecurity        string    `json:"default_security"`
		Description            string    `json:"description"`
		DocumentNumber         int       `json:"document_number"`
		EditDate               time.Time `json:"edit_date"`
		EditProfileDate        time.Time `json:"edit_profile_date"`
		EffectiveSecurity      string    `json:"effective_security"`
		Extension              string    `json:"extension"`
		FileCreateDate         time.Time `json:"file_create_date"`
		FileEditDate           time.Time `json:"file_edit_date"`
		HasAttachment          bool      `json:"has_attachment"`
		HasDocuments           bool      `json:"has_documents"`
		HasSubfolders          bool      `json:"has_subfolders"`
		ID                     string    `json:"id"`
		InUse                  bool      `json:"in_use"`
		Indexable              bool      `json:"indexable"`
		IsCheckedOut           bool      `json:"is_checked_out"`
		IsContainerSavedSearch bool      `json:"is_container_saved_search"`
		IsContentSavedSearch   bool      `json:"is_content_saved_search"`
		IsExternal             bool      `json:"is_external"`
		IsExternalAsNormal     bool      `json:"is_external_as_normal"`
		IsHipaa                bool      `json:"is_hipaa"`
		IsRestorable           bool      `json:"is_restorable"`
		Iwl                    string    `json:"iwl"`
		LastUser               string    `json:"last_user"`
		LastUserDescription    string    `json:"last_user_description"`
		LastUserInfo           struct {
			AllowLogon bool `json:"allow_logon"`
			IsExternal bool `json:"is_external"`
		} `json:"last_user_info"`
		Name       string `json:"name"`
		Operations struct {
			AddCustomFolders      bool     `json:"add_custom_folders"`
			AddFolders            bool     `json:"add_folders"`
			ChangeName            bool     `json:"change_name"`
			Modify                bool     `json:"modify"`
			Remove                bool     `json:"remove"`
			RemoveFolders         bool     `json:"remove_folders"`
			SetDefaultSecurity    bool     `json:"set_default_security"`
			SetSecurity           bool     `json:"set_security"`
			UpdateContentShortcut bool     `json:"update_content_shortcut"`
			DefaultSecurityValues []string `json:"default_security_values"`
		} `json:"operations"`
		Owner            string `json:"owner"`
		OwnerDescription string `json:"owner_description"`
		OwnerInfo        struct {
			AllowLogon bool `json:"allow_logon"`
			IsExternal bool `json:"is_external"`
		} `json:"owner_info"`
		ProjectCustom1 string `json:"project_custom1"`
		ProjectCustom2 string `json:"project_custom2"`
		ProjectCustom3 string `json:"project_custom3"`
		Size           int    `json:"size"`
		Subclass       string `json:"subclass"`
		Subtype        string `json:"subtype"`
		Type           string `json:"type"`
		Version        int    `json:"version"`
		WorkspaceID    string `json:"workspace_id"`
		Wstype         string `json:"wstype"`
	} `json:"data"`
}

type GetWorkspaceProfileOptions struct {
    IncludeOperations bool `url:"include_operations,omitempty"`
    ProfileCheck bool `url:"profile_check,omitempty"`
}
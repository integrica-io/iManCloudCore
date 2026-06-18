package library

import (
	"context"
	"time"
	
	"github.com/integrica-io/iManCloudCore/internal"
	"github.com/integrica-io/iManCloudCore/client"
	"github.com/google/go-querystring/query"	
)

func GetFolderProfile(ctx context.Context, client *client.Client, libraryId string, folderId string, options *GetFolderProfileOptions)(GetFolderProfileOutput, error){
	var data GetFolderProfileOutput
	endpoint := client.BaseUrl.JoinPath("work","api","v2","customers",client.TokenCfg.CustomerId, "libraries", libraryId, "folders", folderId)

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

type GetFolderProfileOutput struct {
	Data struct {
		ProjectCustom3           string    `json:"project_custom3"`
		Database                 string    `json:"database"`
		DefaultSecurity          string    `json:"default_security"`
		EditDate                 time.Time `json:"edit_date"`
		FolderType               string    `json:"folder_type"`
		HasDocuments             bool      `json:"has_documents"`
		HasSubfolders            bool      `json:"has_subfolders"`
		ID                       string    `json:"id"`
		InheritedDefaultSecurity string    `json:"inherited_default_security"`
		IsContainerSavedSearch   bool      `json:"is_container_saved_search"`
		IsContentSavedSearch     bool      `json:"is_content_saved_search"`
		IsExternalAsNormal       bool      `json:"is_external_as_normal"`
		Name                     string    `json:"name"`
		Operations               struct {
			AddContent            bool     `json:"add_content"`
			AddCustomFolders      bool     `json:"add_custom_folders"`
			AddFolders            bool     `json:"add_folders"`
			ChangeName            bool     `json:"change_name"`
			Modify                bool     `json:"modify"`
			ModifyLayout          bool     `json:"modify_layout"`
			Remove                bool     `json:"remove"`
			RemoveContent         bool     `json:"remove_content"`
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
		ParentID string `json:"parent_id"`
		Profile  struct {
			Class              string `json:"class"`
			Custom1            string `json:"custom1"`
			Custom1Description string `json:"custom1_description"`
			Custom2            string `json:"custom2"`
			Custom2Description string `json:"custom2_description"`
		} `json:"profile"`
		ViewType string `json:"view_type"`
		Wstype   string `json:"wstype"`
	} `json:"data"`
}

type GetFolderProfileOptions struct { 
    IncludeOperations bool `url:"include_operations,omitempty"`
/*
Specifies to return a list of operations allowed on a folder by the current user.<br><br>
If `true`, returns a list of operations allowed on a folder by the user.<br>
If `false`, does not return list of operations allowed on a folder by the user.<br>
If omitted, the value is `false`.

*/

    ProfileCheck bool `url:"profile_check,omitempty"`
/*
Specifies warnings for missing required fields and disabled fields based on the folder's class (if it has one).<br><br>
If `true`, returns warnings for missing required fields and disabled fields for the class set on folder.<br>
If `false`, does not return warnings for missing required fields and disabled fields for the class set on folder.<br>
For example, if `custom1` is the mandatory field for a class that is not defined, and `profile_check=true`, a warning is returned.

```
"warnings": [
    {
        "error": "required",
        "field": "custom1"
    }
],
"error": {
    "code": "NRC_INVALID_PROFILE",
    "code_message": "Document profile is invalid",
    "messages": [
        {
            "code": "required",
            "field": "custom1"
        }
    ]
}
```

*/

}



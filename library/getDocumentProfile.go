package library

import (
	"context"
	"fmt"
	"time"

	"github.com/google/go-querystring/query"
	"github.com/integrica-io/iManCloudCore/client"
	"github.com/integrica-io/iManCloudCore/internal"
)

func GetDocumentProfile(ctx context.Context, client *client.Client, libraryId string, docId string, options *GetDocumentProfileOptions) (GetDocumentProfileOutput, error) {
	var data GetDocumentProfileOutput
	endpoint := client.BaseUrl.JoinPath("work", "api", "v2", "customers", client.TokenCfg.CustomerId, "libraries", libraryId, "documents", fmt.Sprintf("%s!%s", libraryId, docId))

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

type GetDocumentProfileOutput struct {
	Data struct {
		Access            string `json:"access"`
		Author            string `json:"author"`
		AuthorDescription string `json:"author_description"`
		AuthorInfo        struct {
			AllowLogon bool `json:"allow_logon"`
			IsExternal bool `json:"is_external"`
		} `json:"author_info"`
		Class               string    `json:"class"`
		CreateDate          time.Time `json:"create_date"`
		Custom1             string    `json:"custom1"`
		Custom1Description  string    `json:"custom1_description"`
		Custom2             string    `json:"custom2"`
		Custom25            bool      `json:"custom25"`
		Custom26            bool      `json:"custom26"`
		Custom27            bool      `json:"custom27"`
		Custom28            bool      `json:"custom28"`
		Custom29            string    `json:"custom29"`
		Custom29Description string    `json:"custom29_description"`
		Custom2Description  string    `json:"custom2_description"`
		Custom3             string    `json:"custom3"`
		Custom30            string    `json:"custom30"`
		Custom30Description string    `json:"custom30_description"`
		Custom3Description  string    `json:"custom3_description"`
		Custom4             string    `json:"custom4"`
		Custom4Description  string    `json:"custom4_description"`
		Database            string    `json:"database"`
		DefaultSecurity     string    `json:"default_security"`
		DocumentNumber      int       `json:"document_number"`
		EditDate            time.Time `json:"edit_date"`
		EditProfileDate     time.Time `json:"edit_profile_date"`
		Extension           string    `json:"extension"`
		FileCreateDate      time.Time `json:"file_create_date"`
		FileEditDate        time.Time `json:"file_edit_date"`
		HasAttachment       bool      `json:"has_attachment"`
		ID                  string    `json:"id"`
		IsInUse             bool      `json:"is_in_use"`
		IsDeclared          bool      `json:"is_declared"`
		IsCheckedOut        bool      `json:"is_checked_out"`
		IsHipaa             bool      `json:"is_hipaa"`
		IsRestorable        bool      `json:"is_restorable"`
		Iwl                 string    `json:"iwl"`
		LastUser            string    `json:"last_user"`
		LastUserDescription string    `json:"last_user_description"`
		LastUserInfo        struct {
			AllowLogon bool `json:"allow_logon"`
			IsExternal bool `json:"is_external"`
		} `json:"last_user_info"`
		Name       string `json:"name"`
		Operations struct {
			Archive          bool `json:"archive"`
			Copy             bool `json:"copy"`
			CreateNewVersion bool `json:"create_new_version"`
			Declare          bool `json:"declare"`
			Delete           bool `json:"delete"`
			Freeze           bool `json:"freeze"`
			Lock             bool `json:"lock"`
			Move             bool `json:"move"`
			Relate           bool `json:"relate"`
			Replace          bool `json:"replace"`
			Restore          bool `json:"restore"`
			SetSecurity      bool `json:"set_security"`
			Undeclare        bool `json:"undeclare"`
			Unfreeze         bool `json:"unfreeze"`
			Unlock           bool `json:"unlock"`
			Update           bool `json:"update"`
		} `json:"operations"`
		Operator            string `json:"operator"`
		OperatorDescription string `json:"operator_description"`
		OperatorInfo        struct {
			AllowLogon bool `json:"allow_logon"`
			IsExternal bool `json:"is_external"`
		} `json:"operator_info"`
		Size                int      `json:"size"`
		Type                string   `json:"type"`
		Version             int      `json:"version"`
		WorkspaceID         string   `json:"workspace_id"`
		Wstype              string   `json:"wstype"`
		LockType            string   `json:"lock_type"`
		CoAuthors           []string `json:"co_authors"`
		WopiFileSizeLimit   int      `json:"wopi_file_size_limit"`
		WopiFileSizeWarning bool     `json:"wopi_file_size_warning"`
	} `json:"data"`
}

type GetDocumentProfileOptions struct { 
    IncludeOperations bool `url:"include_operations,omitempty"`
/*
Specifies to return a list of the user's allowed operations on the document.<br><br>
If `true`, returns a list of allowed operations for the user on the document.<br>
If `false`, does not return a list of allowed operations for the user on the document.
*/

    IsLatest bool `url:"is_latest,omitempty"`
/*
Specifies to return additional information about the latest version of the document in the response.

If `true`, returns `is_latest`, `latest` and `latest_version`.<br>
If `false`, does not return latest version information.
*/

    ParentId string `url:"parent_id,omitempty"`
/*
Specifies the ID of the parent container.<br><br>
This can be a folder ID or workspace ID.
The parent container can affect the operations allowed to be performed on the document.
*/

    ProfileCheck bool `url:"profile_check,omitempty"`
/*
Specifies if warnings for the missing required properties and disabled properties based on the folder's class should be returned.<br><br>
If `true`, returns warnings for missing required fields and disabled fields.<br>
If `false`, does not return warnings for missing required fields and disabled fields.
*/

}




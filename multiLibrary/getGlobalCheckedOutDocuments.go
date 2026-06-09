package multiLibrary

import (
	"context"
	"github.com/integrica-io/iManCloudCore/internal"
	"github.com/integrica-io/iManCloudCore/client"
	"github.com/google/go-querystring/query"	
)

type GetGlobalCheckedOutDocumentsOutput struct {
	Data []struct {
		Author              string `json:"author"`
		AuthorDescription   string `json:"author_description"`
		Class               string `json:"class"`
		CreateDate          string `json:"create_date"`
		Custom1             string `json:"custom1"`
		Custom1Description  string `json:"custom1_description"`
		Custom2             string `json:"custom2"`
		Custom29            string `json:"custom29"`
		Custom29Description string `json:"custom29_description"`
		Custom2Description  string `json:"custom2_description"`
		Database            string `json:"database"`
		DefaultSecurity     string `json:"default_security"`
		DocumentNumber      int    `json:"document_number"`
		EditDate            string `json:"edit_date"`
		EditProfileDate     string `json:"edit_profile_date"`
		Extension           string `json:"extension"`
		FileCreateDate      string `json:"file_create_date"`
		FileEditDate        string `json:"file_edit_date"`
		HasAttachment       bool   `json:"has_attachment"`
		ID                  string `json:"id"`
		InUse               bool   `json:"in_use"`
		InUseBy             string `json:"in_use_by,omitempty"`
		IsDeclared          bool   `json:"is_declared"`
		IsCheckedOut        bool   `json:"is_checked_out"`
		IsHipaa             bool   `json:"is_hipaa"`
		IsRestorable        bool   `json:"is_restorable"`
		Iwl                 string `json:"iwl"`
		LastUser            string `json:"last_user"`
		Name                string `json:"name"`
		Operator            string `json:"operator"`
		OperatorDescription string `json:"operator_description"`
		Size                int    `json:"size"`
		Type                string `json:"type"`
		Version             int    `json:"version"`
		WorkspaceID         string `json:"workspace_id"`
		WorkspaceName       string `json:"workspace_name"`
		Wstype              string `json:"wstype"`
		WopiFileSizeLimit   int    `json:"wopi_file_size_limit"`
		WopiFileSizeWarning bool   `json:"wopi_file_size_warning"`
	} `json:"data"`
	TotalCount int `json:"total_count"`
}

type GetGlobalCheckedOutDocumentsOptions struct {
    Libraries string `url:"libraries,omitempty"`
    Limit int `url:"limit,omitempty"`
    Offset int `url:"offset,omitempty"`
    Total bool `url:"total,omitempty"`
    Type string `url:"type,omitempty"`
    User string `url:"user,omitempty"`
}

func GetGlobalCheckedOutDocuments(ctx context.Context, client *client.Client, customerId string, options *GetGlobalCheckedOutDocumentsOptions)(GetGlobalCheckedOutDocumentsOutput, error){

	var data GetGlobalCheckedOutDocumentsOutput
	endpoint := client.BaseUrl.JoinPath("work","api","v2","customers",client.TokenCfg.CustomerId, "documents", "checked-out")

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
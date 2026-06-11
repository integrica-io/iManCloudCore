package library

import (
	"time"
	"context"
	"github.com/integrica-io/iManCloudCore/internal"
	"github.com/integrica-io/iManCloudCore/client"

	"github.com/google/go-querystring/query"	
)

type GetWorkspaceFromLibraryOutput struct {
	Cursor string `json:"cursor"`
	Data   struct {
		Results []struct {
			Author                 string    `json:"author"`
			Class                  string    `json:"class"`
			ContentType            string    `json:"content_type"`
			CreateDate             time.Time `json:"create_date"`
			Custom1                string    `json:"custom1,omitempty"`
			Custom2                string    `json:"custom2,omitempty"`
			Custom29               string    `json:"custom29,omitempty"`
			Custom3                string    `json:"custom3,omitempty"`
			Custom30               string    `json:"custom30,omitempty"`
			Custom4                string    `json:"custom4,omitempty"`
			Database               string    `json:"database"`
			DefaultSecurity        string    `json:"default_security"`
			Description            string    `json:"description,omitempty"`
			DocumentNumber         int       `json:"document_number"`
			EditDate               time.Time `json:"edit_date"`
			EditProfileDate        time.Time `json:"edit_profile_date"`
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
			Iwl                    string    `json:"iwl"`
			LastUser               string    `json:"last_user"`
			Location               string    `json:"location"`
			Name                   string    `json:"name"`
			Operator               string    `json:"operator"`
			Owner                  string    `json:"owner"`
			OwnerDescription       string    `json:"owner_description"`
			RetainDays             int       `json:"retain_days"`
			Size                   int       `json:"size"`
			Subtype                string    `json:"subtype"`
			Type                   string    `json:"type"`
			Version                int       `json:"version"`
			WorkspaceID            string    `json:"workspace_id"`
			Wstype                 string    `json:"wstype"`
		} `json:"results"`
		Facets struct {
			Senders []interface{} `json:"senders"`
			Types   []struct {
				Alias       string `json:"alias"`
				Count       int    `json:"count"`
				Description string `json:"description"`
			} `json:"types"`
		} `json:"facets"`
	} `json:"data"`
}

type GetWorkspaceFromLibraryOptions struct {
    Anywhere string `url:"anywhere,omitempty"`
    CreateDate string `url:"create_date,omitempty"`
    Cursor string `url:"cursor,omitempty"`
    Custom1 string `url:"custom1,omitempty"`
    Custom10 string `url:"custom10,omitempty"`
    Custom11 string `url:"custom11,omitempty"`
    Custom12 string `url:"custom12,omitempty"`
    Custom13 string `url:"custom13,omitempty"`
    Custom14 string `url:"custom14,omitempty"`
    Custom15 string `url:"custom15,omitempty"`
    Custom16 string `url:"custom16,omitempty"`
    Custom17 string `url:"custom17,omitempty"`
    Custom18 string `url:"custom18,omitempty"`
    Custom19 string `url:"custom19,omitempty"`
    Custom2 string `url:"custom2,omitempty"`
    Custom20 string `url:"custom20,omitempty"`
    Custom21From string `url:"custom21_from,omitempty"`
    Custom21Relative string `url:"custom21_relative,omitempty"`
    Custom21To string `url:"custom21_to,omitempty"`
    Custom22From string `url:"custom22_from,omitempty"`
    Custom22Relative string `url:"custom22_relative,omitempty"`
    Custom22To string `url:"custom22_to,omitempty"`
    Custom23From string `url:"custom23_from,omitempty"`
    Custom23Relative string `url:"custom23_relative,omitempty"`
    Custom23To string `url:"custom23_to,omitempty"`
    Custom24From string `url:"custom24_from,omitempty"`
    Custom24Relative string `url:"custom24_relative,omitempty"`
    Custom24To string `url:"custom24_to,omitempty"`
    Custom25 bool `url:"custom25,omitempty"`
    Custom26 bool `url:"custom26,omitempty"`
    Custom27 bool `url:"custom27,omitempty"`
    Custom28 bool `url:"custom28,omitempty"`
    Custom29 string `url:"custom29,omitempty"`
    Custom3 string `url:"custom3,omitempty"`
    Custom30 string `url:"custom30,omitempty"`
    Custom4 string `url:"custom4,omitempty"`
    Custom5 string `url:"custom5,omitempty"`
    Custom6 string `url:"custom6,omitempty"`
    Custom7 string `url:"custom7,omitempty"`
    Custom8 string `url:"custom8,omitempty"`
    Custom9 string `url:"custom9,omitempty"`
    Description string `url:"description,omitempty"`
    EditDate string `url:"edit_date,omitempty"`
    Facets string `url:"facets,omitempty"`
    FileTarget string `url:"file_target,omitempty"`
    IncludeIds string `url:"include_ids,omitempty"`
    Language string `url:"language,omitempty"`
    Limit int `url:"limit,omitempty"`
    Name string `url:"name,omitempty"`
    NameOrDescription string `url:"name_or_description,omitempty"`
    Offset int `url:"offset,omitempty"`
    Owner string `url:"owner,omitempty"`
    PagingMode PagingMode `url:"paging_mode,omitempty"`
    SortOrder SortOrder `url:"sort_order,omitempty"`
    TimezoneOffset int `url:"timezone_offset,omitempty"`
    Total bool `url:"total,omitempty"`
}

func getWorkspaceFromLibrary(ctx context.Context, client client.Client, libraryId string, options *GetWorkspaceFromLibraryOptions)(GetWorkspaceFromLibraryOutput, error){
	var data GetWorkspaceFromLibraryOutput
	endpoint := client.BaseUrl.JoinPath("work","api","v2","customers",client.TokenCfg.CustomerId, "libraries", libraryId, "workspaces")

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
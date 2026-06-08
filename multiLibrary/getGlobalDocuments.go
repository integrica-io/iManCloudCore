package multiLibrary

import (
	"time"
	"context"
	"github.com/integrica-io/iManCloudCore/internal"

	"github.com/google/go-querystring/query"	
)

type GetGlobalDocumentsOutput struct {
	Data struct {
		Results []struct {
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
			IsInUse             bool   `json:"is_in_use"`
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
		} `json:"results"`
		Facets struct {
			Senders []any `json:"senders"`
			Types   []struct {
				Alias       string `json:"alias"`
				Count       int    `json:"count"`
				Description string `json:"description"`
			} `json:"types"`
		} `json:"facets"`
	} `json:"data"`
	TotalCount int `json:"total_count"`
}

type GetGlobalDocumentsOptions struct {
    Anywhere string `url:"anywhere,omitempty"`
    Author string `url:"author,omitempty"`
    Body string `url:"body,omitempty"`
    CheckedOut bool `url:"checked_out,omitempty"`
    Comments string `url:"comments,omitempty"`
    CreateDate string `url:"create_date,omitempty"`
    CreateDateFrom time.Time `url:"create_date_from,omitempty"`
    CreateDateTo time.Time `url:"create_date_to,omitempty"`
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
    DocumentNumber string `url:"document_number,omitempty"`
    DocumentVersion string `url:"document_version,omitempty"`
    EditDate string `url:"edit_date,omitempty"`
    EditDateFrom string `url:"edit_date_from,omitempty"`
    EditDateTo string `url:"edit_date_to,omitempty"`
    EmailOnly bool `url:"email_only,omitempty"`
    ExcludeEmails bool `url:"exclude_emails,omitempty"`
    ExcludeShortcuts bool `url:"exclude_shortcuts,omitempty"`
    Facets string `url:"facets,omitempty"`
    FileTarget string `url:"file_target,omitempty"`
    HasAttachment bool `url:"has_attachment,omitempty"`
    InUse bool `url:"in_use,omitempty"`
    InUseBy string `url:"in_use_by,omitempty"`
    Language string `url:"language,omitempty"`
    LastUser string `url:"last_user,omitempty"`
    Libraries string `url:"libraries,omitempty"`
    Limit int `url:"limit,omitempty"`
    Name string `url:"name,omitempty"`
    Offset int `url:"offset,omitempty"`
    Operator string `url:"operator,omitempty"`
    Owner string `url:"owner,omitempty"`
    Personalized bool `url:"personalized,omitempty"`
    ReceivedDate string `url:"received_date,omitempty"`
    ReceivedDateFrom time.Time `url:"received_date_from,omitempty"`
    ReceivedDateTo time.Time `url:"received_date_to,omitempty"`
    Recipient string `url:"recipient,omitempty"`
    Results bool `url:"results,omitempty"`
    Sender string `url:"sender,omitempty"`
    SentDate string `url:"sent_date,omitempty"`
    SentDateFrom string `url:"sent_date_from,omitempty"`
    SentDateTo string `url:"sent_date_to,omitempty"`
    Subject string `url:"subject,omitempty"`
    TimezoneOffset int `url:"timezone_offset,omitempty"`
    Total bool `url:"total,omitempty"`
    Type string `url:"type,omitempty"`
    User string `url:"user,omitempty"`
}

func GetGlobalDocuments(ctx context.Context, client *internal.Client, customerId string, options *GetGlobalDocumentsOptions)(GetGlobalDocumentsOutput, error){

	var data GetGlobalDocumentsOutput
	endpoint := client.BaseUrl.JoinPath("work","api","v2","customers",client.TokenCfg.CustomerId, "documents")

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
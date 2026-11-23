package multiLibrary

import (
	"context"
	"encoding/json"
	"time"

	"github.com/integrica-io/iManCloudCore/client"
	"github.com/integrica-io/iManCloudCore/internal"
)

type SearchGlobalDocumentsOutput struct {
	Data []struct {
		ID           string    `json:"id"`
		Wstype       string    `json:"wstype"`
		Name         string    `json:"name"`
		FileEditDate time.Time `json:"file_edit_date"`
	} `json:"data"`
	Overflow bool `json:"overflow"`
}

// DocumentSearchRequest represents the request body for searching global documents.
type SearchGlobalDocumentRequest struct {
	Limit          *int                         `json:"limit,omitempty"`           // 1-9999, default 500
	TimezoneOffset *int                         `json:"timezone_offset,omitempty"` // -1080 to 1080, default 0
	ProfileFields  *DocumentSearchProfileFields `json:"profile_fields,omitempty"`
	Filters        *DocumentSearchFilters       `json:"filters,omitempty"`
}

// DocumentSearchProfileFields specifies up to 15 profile properties to return.
type DocumentSearchProfileFields struct {
	Document []string `json:"document,omitempty"`
}

// DocumentSearchFilters defines the search criteria for global documents.
// Pointer booleans (*bool) are used so omitted fields aren't serialized as false
// (important for tri-state filters like has_attachment/is_declared and fields defaulting to true like include_subtree).
type DocumentSearchFilters struct {
	// Scope & Container
	Libraries      string `json:"libraries,omitempty"`
	ContainerID    string `json:"container_id,omitempty"`
	IncludeSubtree *bool  `json:"include_subtree,omitempty"` // default true
	Personalized   *bool  `json:"personalized,omitempty"`    // default false

	// Core Document Metadata
	Anywhere        string `json:"anywhere,omitempty"`
	Body            string `json:"body,omitempty"`
	Name            string `json:"name,omitempty"`
	NameComments    string `json:"name_comments,omitempty"`
	Comments        string `json:"comments,omitempty"`
	Alias           string `json:"alias,omitempty"`
	Type            string `json:"type,omitempty"`
	DocumentNumber  any    `json:"document_number,omitempty"`  // string (e.g. "123,124") or int
	DocumentVersion any    `json:"document_version,omitempty"` // string (e.g. "1,2") or int

	// Users
	User     string `json:"user,omitempty"`
	Author   string `json:"author,omitempty"`
	Operator string `json:"operator,omitempty"`
	Owner    string `json:"owner,omitempty"`
	LastUser string `json:"last_user,omitempty"`
	InUseBy  string `json:"in_use_by,omitempty"`

	// Status & Flags
	CheckedOut       *bool `json:"checked_out,omitempty"`
	InUse            *bool `json:"in_use,omitempty"`
	ExcludeShortcuts *bool `json:"exclude_shortcuts,omitempty"` // default false
	IsDeclared       *bool `json:"is_declared,omitempty"`

	// Email Filters
	ExcludeEmails  *bool    `json:"exclude_emails,omitempty"` // default false
	EmailOnly      *bool    `json:"email_only,omitempty"`     // default false
	HasAttachment  *bool    `json:"has_attachment,omitempty"`
	Subject        string   `json:"subject,omitempty"`
	Sender         string   `json:"sender,omitempty"`
	Recipient      string   `json:"recipient,omitempty"`
	EmailAddresses []string `json:"email_addresses,omitempty"`

	// Dates (relative format e.g. "-30:0d", or ISO 8601 for _from/_to)
	CreateDate       string `json:"create_date,omitempty"`
	CreateDateFrom   string `json:"create_date_from,omitempty"`
	CreateDateTo     string `json:"create_date_to,omitempty"`
	EditDate         string `json:"edit_date,omitempty"`
	EditDateFrom     string `json:"edit_date_from,omitempty"`
	EditDateTo       string `json:"edit_date_to,omitempty"`
	SentDate         string `json:"sent_date,omitempty"`
	SentDateFrom     string `json:"sent_date_from,omitempty"`
	SentDateTo       string `json:"sent_date_to,omitempty"`
	ReceivedDate     string `json:"received_date,omitempty"`
	ReceivedDateFrom string `json:"received_date_from,omitempty"`
	ReceivedDateTo   string `json:"received_date_to,omitempty"`

	// Custom Aliases & Text (custom1–custom16, custom29–custom30)
	Custom1  string `json:"custom1,omitempty"`
	Custom2  string `json:"custom2,omitempty"`
	Custom3  string `json:"custom3,omitempty"`
	Custom4  string `json:"custom4,omitempty"`
	Custom5  string `json:"custom5,omitempty"`
	Custom6  string `json:"custom6,omitempty"`
	Custom7  string `json:"custom7,omitempty"`
	Custom8  string `json:"custom8,omitempty"`
	Custom9  string `json:"custom9,omitempty"`
	Custom10 string `json:"custom10,omitempty"`
	Custom11 string `json:"custom11,omitempty"`
	Custom12 string `json:"custom12,omitempty"`
	Custom13 string `json:"custom13,omitempty"`
	Custom14 string `json:"custom14,omitempty"`
	Custom15 string `json:"custom15,omitempty"`
	Custom16 string `json:"custom16,omitempty"`
	Custom29 string `json:"custom29,omitempty"`
	Custom30 string `json:"custom30,omitempty"`

	// Custom Numeric / Range (string e.g. "10330-10400", "<10335" or number)
	Custom17 any `json:"custom17,omitempty"`
	Custom18 any `json:"custom18,omitempty"`
	Custom19 any `json:"custom19,omitempty"`
	Custom20 any `json:"custom20,omitempty"`

	// Custom Dates
	Custom21From     string `json:"custom21_from,omitempty"`
	Custom21To       string `json:"custom21_to,omitempty"`
	Custom21Relative string `json:"custom21_relative,omitempty"`
	Custom22From     string `json:"custom22_from,omitempty"`
	Custom22To       string `json:"custom22_to,omitempty"`
	Custom22Relative string `json:"custom22_relative,omitempty"`
	Custom23From     string `json:"custom23_from,omitempty"`
	Custom23To       string `json:"custom23_to,omitempty"`
	Custom23Relative string `json:"custom23_relative,omitempty"`
	Custom24From     string `json:"custom24_from,omitempty"`
	Custom24To       string `json:"custom24_to,omitempty"`
	Custom24Relative string `json:"custom24_relative,omitempty"`

	// Custom Booleans
	Custom25 *bool `json:"custom25,omitempty"`
	Custom26 *bool `json:"custom26,omitempty"`
	Custom27 *bool `json:"custom27,omitempty"`
	Custom28 *bool `json:"custom28,omitempty"`
}

func SearchGlobalDocuments(ctx context.Context, client *client.Client, customerId string, options SearchGlobalDocumentRequest) (SearchGlobalDocumentsOutput, error) {
	var data SearchGlobalDocumentsOutput
	endpoint := client.BaseUrl.JoinPath("work", "api", "v2", "customers", client.TokenCfg.CustomerId, "documents", "search")

	bytes, err := json.Marshal(data)
	if err != nil {
		return data, err
	}

	req := internal.HttpRequestBuilder{}
	req.Context(ctx).Url(*endpoint).Method(internal.Post).ToJson(&data).Body(bytes)

	if err := client.Req(req); err != nil {
		return data, err
	}
	return data, nil
}
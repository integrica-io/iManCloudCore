package library

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/integrica-io/iManCloudCore/client"
	"github.com/integrica-io/iManCloudCore/internal"
)

// DefaultSecurity represents the allowed default access levels for a document.
type DefaultSecurity string

const (
	DefaultSecurityPublic  DefaultSecurity = "public"
	DefaultSecurityPrivate DefaultSecurity = "private"
	DefaultSecurityView    DefaultSecurity = "view"
)

// ContentType represents the document content type.
type ContentType string

const (
	ContentTypeNormal ContentType = "D"
	ContentTypeOLE    ContentType = "5"
)

// UpdateDocumentProfilePathParams holds the URL path parameters for:
// PATCH /customers/{customerId}/libraries/{libraryId}/documents/{docId}
type UpdateDocumentProfilePathParams struct {
	CustomerID string `json:"customerId"`
	LibraryID  string `json:"libraryId"`
	DocID      string `json:"docId"`
}

// Audit specifies audit information used to track document history.
// Note: The spec sets additionalProperties: true; use Extra or map[string]any
// with custom JSON marshaling if you need to send dynamic keys alongside Comments.
type Audit struct {
	Comments string `json:"comments,omitempty"`
}

// UpdateDocumentProfileRequest is the JSON request body for updating a document profile.
// Pointer fields with omitempty allow partial updates (PATCH semantics) while still
// letting callers explicitly send zero values (e.g., false, 0, "").
type UpdateDocumentProfileOptions struct {
	Alias       *string      `json:"alias,omitempty"`
	Audit       *Audit       `json:"audit,omitempty"`
	Author      *string      `json:"author,omitempty"`
	Class       *string      `json:"class,omitempty"`
	Comment     *string      `json:"comment,omitempty"` // maxLength: 8000
	ContentType *ContentType `json:"content_type,omitempty"`

	Custom1  *string `json:"custom1,omitempty"`
	Custom2  *string `json:"custom2,omitempty"`
	Custom3  *string `json:"custom3,omitempty"`
	Custom4  *string `json:"custom4,omitempty"`
	Custom5  *string `json:"custom5,omitempty"`
	Custom6  *string `json:"custom6,omitempty"`
	Custom7  *string `json:"custom7,omitempty"`
	Custom8  *string `json:"custom8,omitempty"`
	Custom9  *string `json:"custom9,omitempty"`
	Custom10 *string `json:"custom10,omitempty"`
	Custom11 *string `json:"custom11,omitempty"`
	Custom12 *string `json:"custom12,omitempty"`
	Custom13 *string `json:"custom13,omitempty"`
	Custom14 *string `json:"custom14,omitempty"`
	Custom15 *string `json:"custom15,omitempty"`
	Custom16 *string `json:"custom16,omitempty"`

	Custom17 *float64 `json:"custom17,omitempty"`
	Custom18 *float64 `json:"custom18,omitempty"`
	Custom19 *float64 `json:"custom19,omitempty"`
	Custom20 *float64 `json:"custom20,omitempty"`

	// ISO 8601 timestamps (marshaled as RFC 3339 by encoding/json)
	Custom21 *time.Time `json:"custom21,omitempty"`
	Custom22 *time.Time `json:"custom22,omitempty"`
	Custom23 *time.Time `json:"custom23,omitempty"`
	Custom24 *time.Time `json:"custom24,omitempty"`

	Custom25 *bool `json:"custom25,omitempty"`
	Custom26 *bool `json:"custom26,omitempty"`
	Custom27 *bool `json:"custom27,omitempty"`
	Custom28 *bool `json:"custom28,omitempty"`

	Custom29 *string `json:"custom29,omitempty"`
	Custom30 *string `json:"custom30,omitempty"`
	Custom31 *string `json:"custom31,omitempty"` // Reserved. Do not use.

	DefaultSecurity                     *DefaultSecurity `json:"default_security,omitempty"`
	Extension                           *string          `json:"extension,omitempty"`
	IsCheckedOut                        *bool            `json:"is_checked_out,omitempty"`
	IsDeclared                          *bool            `json:"is_declared,omitempty"`
	IsHIPAA                             *bool            `json:"is_hipaa,omitempty"`
	Name                                *string          `json:"name,omitempty"`
	Operator                            *string          `json:"operator,omitempty"`
	RetainDays                          *int             `json:"retain_days,omitempty"`
	Subclass                            *string          `json:"subclass,omitempty"`
	Type                                *string          `json:"type,omitempty"`
	WarningsForRequiredAndDisabledFields *bool           `json:"warnings_for_required_and_disabled_fields,omitempty"`
}

// UpdateDocumentProfileResponse is the HTTP 200 JSON response body.
type UpdateDocumentProfileOutput struct {
	Data     DocumentProfile `json:"data"`
	Warnings []FieldWarning  `json:"warnings,omitempty"`
}

// DocumentProfile represents the returned document profile object inside "data".
// Fields are modeled from the 200 response example (the YAML anchor *id437 was not
// expanded in the provided spec snippet).
type DocumentProfile struct {
	Author              string          `json:"author,omitempty"`
	AuthorDescription   string          `json:"author_description,omitempty"`
	Class               string          `json:"class,omitempty"`
	CreateDate          *time.Time      `json:"create_date,omitempty"`
	Custom1             string          `json:"custom1,omitempty"`
	Custom1Description  string          `json:"custom1_description,omitempty"`
	Database            string          `json:"database,omitempty"`
	DefaultSecurity     DefaultSecurity `json:"default_security,omitempty"`
	DocumentNumber      int64           `json:"document_number,omitempty"`
	EditDate            *time.Time      `json:"edit_date,omitempty"`
	EditProfileDate     *time.Time      `json:"edit_profile_date,omitempty"`
	Extension           string          `json:"extension,omitempty"`
	FileCreateDate      *time.Time      `json:"file_create_date,omitempty"`
	FileEditDate        *time.Time      `json:"file_edit_date,omitempty"`
	HasAttachment       bool            `json:"has_attachment,omitempty"`
	ID                  string          `json:"id,omitempty"`
	InUse               bool            `json:"in_use,omitempty"`
	InUseBy             string          `json:"in_use_by,omitempty"`
	IsDeclared          bool            `json:"is_declared,omitempty"`
	IsCheckedOut        bool            `json:"is_checked_out,omitempty"`
	IsHIPAA             bool            `json:"is_hipaa,omitempty"`
	IsRestorable        bool            `json:"is_restorable,omitempty"`
	IWL                 string          `json:"iwl,omitempty"`
	LastUser            string          `json:"last_user,omitempty"`
	Name                string          `json:"name,omitempty"`
	Operator            string          `json:"operator,omitempty"`
	OperatorDescription string          `json:"operator_description,omitempty"`
	Size                int64           `json:"size,omitempty"`
	Type                string          `json:"type,omitempty"`
	Version             int             `json:"version,omitempty"`
	WSType              string          `json:"wstype,omitempty"`
	WOPIFileSizeLimit   int64           `json:"wopi_file_size_limit,omitempty"`
	WOPIFileSizeWarning bool            `json:"wopi_file_size_warning,omitempty"`
}

// FieldWarning represents a warning returned when warnings_for_required_and_disabled_fields is true.
type FieldWarning struct {
	Error string `json:"error"`
	Field string `json:"field"`
}

func UpdateDocumentProfile(ctx context.Context, client *client.Client, libraryId string, docId string, options UpdateDocumentProfileOptions)(UpdateDocumentProfileOutput, error){
	var data UpdateDocumentProfileOutput
	endpoint := client.BaseUrl.JoinPath("work", "api", "v2", "customers", client.TokenCfg.CustomerId, "libraries", libraryId, "documents", fmt.Sprintf("%s!%s", libraryId, docId))

	bytes, err := json.Marshal(options)
	if err != nil {
		return data, err
	}

	req := internal.HttpRequestBuilder{}
	req.Context(ctx).Url(*endpoint).Method(internal.Patch).ToJson(&data).Body(bytes)

	fmt.Println(req)

	if err := client.Req(req); err != nil {
		return data, err
	}
	return data, nil
}
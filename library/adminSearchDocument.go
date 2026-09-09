package library

import (
	"context"
	"time"

	"github.com/google/go-querystring/query"
	"github.com/integrica-io/iManCloudCore/client"
	"github.com/integrica-io/iManCloudCore/internal"
)

// Returns all documents with the properties that an iManage Work system administrator is allowed to view.
// The minimum access permission required to implement this request: Tier 1.
func AdminSearchDocuments(ctx context.Context, client *client.Client, libraryId string, options *AdminSearchDocumentsOptions) (AdminSearchDocumentsOutput, error) {
	var data AdminSearchDocumentsOutput
	endpoint := client.BaseUrl.JoinPath("work", "api", "v2", "customers", client.TokenCfg.CustomerId, "libraries", libraryId, "documents")

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

// Result struct for AdminSearchDocuments() func
type AdminSearchDocumentsOutput struct {
	Data []struct {
		Author              string    `json:"author"`
		Class               string    `json:"class"`
		ContentType         string    `json:"content_type"`
		CreateDate          time.Time `json:"create_date"`
		Custom21            time.Time `json:"custom21,omitempty"`
		Custom22            time.Time `json:"custom22,omitempty"`
		Custom23            time.Time `json:"custom23,omitempty"`
		Custom24            time.Time `json:"custom24,omitempty"`
		Library             string    `json:"library"`
		Declared            bool      `json:"declared"`
		DefaultSecurity     string    `json:"default_security"`
		DocumentNumber      int       `json:"document_number"`
		EditDate            time.Time `json:"edit_date"`
		EditProfileDate     time.Time `json:"edit_profile_date"`
		Extension           string    `json:"extension"`
		FileCreateDate      time.Time `json:"file_create_date"`
		FileEditDate        time.Time `json:"file_edit_date"`
		ID                  string    `json:"id"`
		InUse               bool      `json:"in_use"`
		IsCheckedOut        bool      `json:"is_checked_out"`
		IsDeclared          bool      `json:"is_declared"`
		IsExternal          bool      `json:"is_external"`
		IsExternalAsNormal  bool      `json:"is_external_as_normal"`
		IsHipaa             bool      `json:"is_hipaa"`
		IsInUse             bool      `json:"is_in_use"`
		IsRelated           bool      `json:"is_related,omitempty"`
		IsRestorable        bool      `json:"is_restorable"`
		Iwl                 string    `json:"iwl"`
		LastUser            string    `json:"last_user"`
		Operator            string    `json:"operator"`
		RetainDays          int       `json:"retain_days"`
		Size                int       `json:"size"`
		Type                string    `json:"type"`
		Version             int       `json:"version"`
		WorkspaceID         string    `json:"workspace_id"`
		Wstype              string    `json:"wstype"`
		WopiFileSizeLimit   int       `json:"wopi_file_size_limit"`
		WopiFileSizeWarning bool      `json:"wopi_file_size_warning"`
		Custom3             string    `json:"custom3,omitempty"`
	} `json:"data"`
	TotalCount int `json:"total_count"`
}

type AdminSearchDocumentsOptions struct {
	Ascending bool `url:"ascending,omitempty"`
	/*
	   Indicates to return the documents in the ascending order of document numbers.<br><br>
	   If `true`, returns the documents in the ascending order.<br>
	   If `false`, returns the documents in any order.

	*/

	ContainerId string `url:"container_id,omitempty"`
	/*
	   Returns documents only from the specified container.<br><br>
	   The specified container ID can be of a workspace or folder. For example, *Archive!22*, or *Integration!1*.
	   Multiple query parameters cannot be specified as a comma-separated list.

	*/

	EmailOnly bool `url:"email_only,omitempty"`
	/*
	   Indicates whether to search for documents in the entire subtree of the specified container ID.<br><br>
	   If true, searches for documents in the entire subtree of the specified container ID.<br>
	   If false, search for documents only in the specified container ID.<br><br>

	*/

	Facets string `url:"facets,omitempty"`
	/*
	   Specifies the facets to return.

	   For more information, see <a href="#overview--facets">Facets</a> in the **Key Concepts** section.

	*/

	IncludeSubtree bool `url:"include_subtree,omitempty"`
	/*
	   Indicates whether to include single level of the container_id.<br><br>
	   If `true`, search the entire subtree of the specified container_id.<br>
	   If `false`, search a single level of the container_id.

	*/

	Latest bool `url:"latest,omitempty"`
	/*
	   Indicates to include only the latest version of each document.<br><br>
	   If `true`, only the current or latest version is included.<br>
	   If `false`, older versions of documents are also included.

	*/

	Limit int `url:"limit,omitempty"`
	/*
	   Specifies the maximum number of items to include in the response.

	   The request returns the actual number of items up to the limit value (inclusive). If there are more items than the limit value, no more items than the limit are returned and a cursor value is returned.

	   This parameter can be used in conjunction with pagination parameters for endpoints that support them. For more information, refer to <a href="#overview--pagination">Pagination</a> in the **Key Concepts** section.
	*/

	NvpName string `url:"nvp_name,omitempty"`
	/*

	 */

	NvpOp string `url:"nvp_op,omitempty"`
	/*
	   Deprecated. Do not use.
	*/

	NvpValue string `url:"nvp_value,omitempty"`
	/*

	 */

	PagingMode PagingMode `url:"paging_mode,omitempty"`
	/*
	   Specifies the paging mode to be used to retrieve the result set.

	   The following are allowed values:

	   - standard_cursor - Specifies the query parameter `cursor`can be used to identify the first item to return.
	   The cursor is a value that specifies an item by an ID within the results set to start returning from.
	   It is not an ordinal position as `offset` is for the standard paging mode.
	   For a subsequent search, set `cursor` to the cursor value returned in the previous search.
	   Do not modify the cursor value in anyway.
	   The standard_cursor mode is the preferred paging mode, as it is quicker and less stressful on system performance.
	   The parameter `offset` is ignored in this mode.<br> **Note**: In this mode the search term must match the document name exactly.
	   For example, when searching for `name=warrant`, documents with the name "warranty agreement" will not be included in the results.<br>

	   - standard -  Specifies the query parameter `offset` can be used to identify the first item to return.
	   The offset is a value that specifies the position of an item within the results set to start returning from.
	   This is relative to zero, or the absolute start of the list.
	   An offset value of zero starts at the beginning of the results set.
	   An offset value of 5, starts with the fifth item in the results set.
	   The parameter `cursor` is ignored in this mode.<br> **Note**: In this mode the search term can match the document name partially.
	   For example, when searching for `name=warrant`, documents with the name "warranty agreement" will be included in the results.<br><br>
	*/

	Personalized bool `url:"personalized,omitempty"`
	/*
	   The content that the user has accessed from My Matters in the past 30 days.<br><br>
	   If `true`, retrieves the documents from a user's Recent Matters list.<br>
	   If `false`, retrieves the documents from all sources available to the user.

	*/

	Results bool `url:"results,omitempty"`
	/*
	   Indicates whether to include search results.
	*/

	SortOrder string `url:"sort_order,omitempty"`
	/*
	   The sort order to be used when the paging mode is set to use a cursor.<br>

	   The following are allowed values:

	   |Value|Description|
	   |:---|:---|
	   |asc|The items are listed in ascending order either alphabetically or numerically.|
	   |desc|The items are listed in descending order either alphabetically or numerically.|

	   If omitted, items are returned in the order they are found.
	*/

	TimezoneOffset int `url:"timezone_offset,omitempty"`
	/*
	   Specifies the offset, in minutes, from Coordinated Universal Time (UTC).

	   iManage Work uses UTC date values. Timezone offset can be used to translate UTC to the local time for display in the client application.

	   The value can be positive (minutes ahead of UTC) or negative (minutes behind UTC).<br>
	   For example, India is 5 hours and 30 minutes ahead of UTC, which is 330 minutes, and is specified as `timezone_offset=330`.<br>
	   Houston is six hours behind UTC, which is -360 minutes, and is specified as `timezone_offset=-360`.
	*/
}

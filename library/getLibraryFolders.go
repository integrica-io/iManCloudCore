package library

import (
	"context"
	"time"
	
	"github.com/integrica-io/iManCloudCore/internal"
	"github.com/integrica-io/iManCloudCore/client"
	"github.com/google/go-querystring/query"	
)

func GetLibraryFolders(ctx context.Context, client *client.Client, libraryId string, options *GetLibraryFoldersOptions)(GetLibraryFoldersOuput, error){
	var data GetLibraryFoldersOuput
	endpoint := client.BaseUrl.JoinPath("work","api","v2","customers",client.TokenCfg.CustomerId, "libraries", libraryId, "folders")

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

type GetLibraryFoldersOuput struct {
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

type GetLibraryFoldersOptions struct { 
    ContainerId string `url:"container_id,omitempty"`
/*
Specifies the container ID of a workspace or a folder to return folders from.<br>

If the container ID is specified, the details of this container and folders under it are returned.<br>
If the container ID is not specified, all the folders from the library are returned.<br>
Only one container ID can be specified. Multiple values cannot be specified.

*/

    Cursor string `url:"cursor,omitempty"`
/*
Specifies the cursor value to retrieve the next set of results.

If a cursor is not specified, the result set is from the beginning. For more information, refer to <a href="#overview--pagination">Pagination</a> in the **Key Concepts** section.

*/

    Description string `url:"description,omitempty"`
/*
Specifies a value to be found in the `description` property of the folder.<br><br>
This parameter is case insensitive and matches whole fields.<br>
To match partial fields, include the wildcard character (&ast;) before, after, or both ends of the term.

Special characters allowed are *().&-_[]`~\|@$%^?:{}!',/\\#+<>;"=


*/

    Email string `url:"email,omitempty"`
/*
Specifies the email address to match with the folder email address.

Special characters allowed are *().&-_[]`~\|@$%^?:{}!',/\\#+<>;"=

*/

    IncludeIds string `url:"include_ids,omitempty"`
/*
Specifies to return folders only from the list of folder IDs provided.<br>

For example, `include_ids=Active_uk!2222`.<br>
Multiple folders can be included using a comma-separated list. For example, `include_ids=active_uk!22,active_uk!55,active_uk!401`.<br><br>
If this parameter is omitted, all the folders from the library are returned.

*/

    Limit int `url:"limit,omitempty"`
/*
Specifies the maximum number of items to include in the response.

The request returns the actual number of items up to the limit value (inclusive). If there are more items than the limit value, no more items than the limit are returned and a cursor value is returned.

This parameter can be used in conjunction with pagination parameters for endpoints that support them. For more information, refer to <a href="#overview--pagination">Pagination</a> in the **Key Concepts** section.
*/

    Name string `url:"name,omitempty"`
/*
Specifies a value to be found in the `name` property of the folder.<br><br>
This parameter is case insensitive and matches whole fields.
To match partial fields, include the wildcard character (&ast;) before, after, or both ends of the term.<br>
Multiple names can be included using a comma-separated list.

Special characters allowed are *().&-_[]`~\|@$%^?:{}!',/\\#+<>;"=


*/

    NameOrDescription string `url:"name_or_description,omitempty"`
/*
Specifies the text to search for in the `name` or `description` fields of iManage Work object.<br>

* This parameter is case insensitive.<br>
* Matches whole fields only and partial matches are not supported.
  For example, for an object named *Enron Completion*, the search `name_or_description=Enron` does not match.
  It must be specified as `name_or_description=Enron Completion`.
* The wildcard character (&ast;) is supported, and matches any number of contiguous characters, including spaces.<br>
  To match partial fields, include the wildcard character (&ast;) before, after, or at both ends of the term.<br>
  For example, all the following searches match with `litigation`:
  * `name_or_description=litigation`
  * `name_or_description=*gation`
  * `name_or_description=litiga*`.


Special characters allowed are *().&-_[]`~\|@$%^?:{}!',/\\#+<>;"=


*/

    Offset int `url:"offset,omitempty"`
/*
Specifies the position of the first item to be returned from the result set.

This value indicates the starting position for the first item to return from among all the possible items.
By default, this value is zero, meaning items are returned starting from the first item in the list.
For example, if the offset is ten (`offset=10`), the first item returned will be the 11th item in the list.
The `offset` is often used in combination with the `limit` parameter.
For example, if `limit=10&offset=10`, the list returns ten items starting from the 11th item through the 20th.

If the value is greater than the number of items in the return list, no items are returned but no error occurs.<br>
For more information, see <a href="#overview--pagination">Pagination</a>.
*/

    Owner string `url:"owner,omitempty"`
/*
Specifies the owner's name for this folder.<br><br>
This parameter is case insensitive and matches whole fields.
To match partial fields, include the wildcard character (&ast;) before, after, or both the ends of a search term.

*/

    PagingMode PagingMode `url:"paging_mode,omitempty"`
/*
Specifies the paging mode to be used to retrieve the result set.

The following are allowed values:
* `standard_cursor`
* `standard`

Using `standard_cursor` paging mode, the query parameter `cursor` can be used to identify the first item to return.
* The cursor is a value that specifies an item by an ID within the results set to start returning from.
    * This is not an ordinal position like `offset` is for the `standard` paging mode.
* For a subsequent search, set `cursor` to the cursor value returned in the previous search.
* Do not modify the cursor value in any way.
* The `standard_cursor` mode is the preferred paging mode, as it is quicker and less stressful on system performance.
* The parameter `offset` is ignored in this mode.

Using `standard` paging mode, the query parameter `offset` can be used to identify the first item to return.
* `offset` is a value that specifies the position of an item within the results set to start returning from.
    * This is relative to zero, the absolute start of the list.
* An offset value of zero starts at the beginning of the results set.
* An offset value of 5, starts with the fifth item in the results set.
* The parameter `cursor` is ignored in this mode.

For more information, see <a href="#overview--pagination">Pagination</a> in the **Key Concepts** section.
*/

    ReferenceDatabase string `url:"reference_database,omitempty"`
/*
Specifies the target library to which the folder shortcuts and workspace shortcuts contained in the folder are pointing to.
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

    Total bool `url:"total,omitempty"`
/*
Specifies to include the total count of items found in the response.

If `true`, the total count is included in the response.<br>
If `false`, the total count is not included in the response.<br>
The actual number of items returned may be different because of the *limit* parameter that restricts the number of items
returned for any given search.<br>

*/

}
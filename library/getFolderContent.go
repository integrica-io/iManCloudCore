package library

import (
	"context"
	"time"
	
	"github.com/integrica-io/iManCloudCore/internal"
	"github.com/integrica-io/iManCloudCore/client"
	"github.com/google/go-querystring/query"	
)

func GetFolderContent(ctx context.Context, client *client.Client, libraryId string, folderId string, options *GetFolderContentOptions)(GetFolderContentOutput, error){
	var data GetFolderContentOutput
	endpoint := client.BaseUrl.JoinPath("work","api","v2","customers",client.TokenCfg.CustomerId, "libraries", libraryId, "folders", folderId, "children")

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

type GetFolderContentOutput struct {
	Data []struct {
		Database                 string    `json:"database"`
		DefaultSecurity          string    `json:"default_security"`
		EditDate                 time.Time `json:"edit_date"`
		FolderType               string    `json:"folder_type,omitempty"`
		HasDocuments             bool      `json:"has_documents,omitempty"`
		HasSubfolders            bool      `json:"has_subfolders,omitempty"`
		ID                       string    `json:"id"`
		InheritedDefaultSecurity string    `json:"inherited_default_security,omitempty"`
		IsContainerSavedSearch   bool      `json:"is_container_saved_search,omitempty"`
		IsContentSavedSearch     bool      `json:"is_content_saved_search,omitempty"`
		IsExternal               bool      `json:"is_external,omitempty"`
		IsExternalAsNormal       bool      `json:"is_external_as_normal,omitempty"`
		Name                     string    `json:"name"`
		OwnerDescription         string    `json:"owner_description,omitempty"`
		Owner                    string    `json:"owner,omitempty"`
		ParentID                 string    `json:"parent_id,omitempty"`
		ViewType                 string    `json:"view_type,omitempty"`
		WorkspaceID              string    `json:"workspace_id"`
		WorkspaceName            string    `json:"workspace_name"`
		Wstype                   string    `json:"wstype"`
		Author                   string    `json:"author,omitempty"`
		AuthorDescription        string    `json:"author_description,omitempty"`
		Checksum                 string    `json:"checksum,omitempty"`
		ClassDescription         string    `json:"class_description,omitempty"`
		Class                    string    `json:"class,omitempty"`
		CreateDate               time.Time `json:"create_date,omitempty"`
		DocumentNumber           int       `json:"document_number,omitempty"`
		EditProfileDate          time.Time `json:"edit_profile_date,omitempty"`
		FileCreateDate           time.Time `json:"file_create_date,omitempty"`
		FileEditDate             time.Time `json:"file_edit_date,omitempty"`
		Iwl                      string    `json:"iwl,omitempty"`
		LastUser                 string    `json:"last_user,omitempty"`
		LastUserDescription      string    `json:"last_user_description,omitempty"`
		OperatorDescription      string    `json:"operator_description,omitempty"`
		Operator                 string    `json:"operator,omitempty"`
		RetainDays               int       `json:"retain_days,omitempty"`
		Size                     int       `json:"size,omitempty"`
		SubclassDescription      string    `json:"subclass_description,omitempty"`
		Type                     string    `json:"type,omitempty"`
		Extension                string    `json:"extension,omitempty"`
		TypeDescription          string    `json:"type_description,omitempty"`
		Version                  int       `json:"version,omitempty"`
	} `json:"data"`
}

type GetFolderContentOptions struct { 
    Anywhere string `url:"anywhere,omitempty"`
/*
Filters documents based on the matching text found anywhere in the document's profile or its contents.

*/

    Author string `url:"author,omitempty"`
/*
Filters documents based on the user ID of the author of the documents.
*/

    Body string `url:"body,omitempty"`
/*
Filters documents based on the matching text found in a document's content.
*/

    CheckedOut bool `url:"checked_out,omitempty"`
/*
Filters documents based on the `checked_out` status.<br><br>
If `true`, returns only checked out documents.<br>
If `false`, returns only documents that are not checked out.<br>
If omitted, returns documents regardless of whether they are checked out.

*/

    Comments string `url:"comments,omitempty"`
/*
Filters documents based on the matching text found in the document profile's comments.
* Partial matches are supported.
* The wildcard character (*) is not supported.
*/

    CreateDate string `url:"create_date,omitempty"`
/*
Filters documents by create date within a date range relative to the current date.

For more information about date range notation, see <a href="#overview--date-range-format">Date range format</a> in the **Key Concepts** section.
*/

    CreateDateFrom string `url:"create_date_from,omitempty"`
/*
Filters documents created on or after this date (in ISO 8601 format).
*/

    CreateDateTo string `url:"create_date_to,omitempty"`
/*
Filters documents created on or before this date (in ISO 8601 format).
*/

    Cursor string `url:"cursor,omitempty"`
/*
Specifies the cursor value to retrieve the next set of results.

If a cursor is not specified, the result set is from the beginning. For more information, refer to <a href="#overview--pagination">Pagination</a> in the **Key Concepts** section.

*/

    Custom1 string `url:"custom1,omitempty"`
/*
Filters documents based on the alias of the `custom1` property.<br>
Specifies a property alias to match with the custom1 property.

A property alias is an entry for a specific custom property.
For example, `custom1` can be captioned as *Client*.
It can have multiple entries, such as 001, 001abc, 001001. These are known as aliases of clients.
These aliases can be associated with client descriptions such as Ajubalaw, Microsoft, and Wallachs.

* This search is not case-sensitive.
* Partial matches are not supported.
* The wildcard character (&ast;) is not supported.

*/

    Custom10 string `url:"custom10,omitempty"`
/*
Filters documents based on the alias of the `custom10` property.<br>
* Partial matches are not supported.
* The wildcard character (*) is not supported.
*/

    Custom11 string `url:"custom11,omitempty"`
/*
Filters documents based on the alias of the `custom11` property.<br>
* Partial matches are not supported.
* The wildcard character (*) is not supported.
*/

    Custom12 string `url:"custom12,omitempty"`
/*
Filters documents based on the alias of the `custom12` property.<br>
* Partial matches are not supported.
* The wildcard character (*) is not supported.
*/

    Custom13 string `url:"custom13,omitempty"`
/*
Filters documents based on the value of the `custom13` property.<br>
* Partial matches are not supported.
* The wildcard character (*) is not supported.
*/

    Custom14 string `url:"custom14,omitempty"`
/*
Filters documents based on the value of the `custom14` property.<br>
* Partial matches are not supported.
* The wildcard character (*) is not supported.
*/

    Custom15 string `url:"custom15,omitempty"`
/*
Filters documents based on the value of the `custom15` property.<br>
* Partial matches are not supported.
* The wildcard character (*) is not supported.
*/

    Custom16 string `url:"custom16,omitempty"`
/*
Filters documents based on the value of the `custom16` property.<br>
* Partial matches are not supported.
* The wildcard character (*) is not supported.
*/

    Custom17 string `url:"custom17,omitempty"`
/*
Specifies a value or a range of values of the custom17 property to match.<br><br>

|Symbol|Value|Description|
|:---|:----|:---|
| = | Equal to | The exact custom17 value to be searched for. For example, *custom17=10335*.|
| &lt; | Less than | The custom17 value to be searched for is less than the given value. For example, *custom17=&lt;10335*.|
| &gt; | Greater than | The custom17 value to be searched for is greater than the given value. For example, *custom17=&gt;10335*.|
| - | Range | The range of custom17 value. This search includes both the lower and higher values specified in the range. This range must be in ascending order. For example, *custom17=10330-10400*.|
*/

    Custom18 string `url:"custom18,omitempty"`
/*
Specifies a value or a range of values of the custom18 property to match.<br><br>

|Symbol|Value|Description|
|:---|:----|:---|
| = | Equal to | The exact custom18 value to be searched for. For example, *custom18=10335*.|
| &lt; | Less than | The custom18 value to be searched for is less than the given value. For example, *custom18=&lt;10335*.|
| &gt; | Greater than | The custom18 value to be searched for is greater than the given value. For example, *custom18=&gt;10335*.|
| - | Range | The range of custom18 value. This search includes both the lower and higher values specified in the range. This range must be in ascending order. For example, *custom18=10330-10400*.|
*/

    Custom19 string `url:"custom19,omitempty"`
/*
Specifies a value or a range of values of the custom19 property to match.<br><br>

|Symbol|Value|Description|
|:---|:----|:---|
| = | Equal to | The exact custom19 value to be searched for. For example, *custom19=10335*.|
| &lt; | Less than | The custom19 value to be searched for is less than the given value. For example, *custom19=&lt;10335*.|
| &gt; | Greater than | The custom19 value to be searched for is greater than the given value. For example, *custom19=&gt;10335*.|
| - | Range | The range of custom19 values. This search includes both the lower and higher values specified in the range. This range must be in ascending order. For example, *custom19=10330-10400*.|
*/

    Custom2 string `url:"custom2,omitempty"`
/*
Filters documents based on the alias of the `custom2` property.<br>
* Partial matches are not supported.
* The wildcard character (*) is not supported.
*/

    Custom20 string `url:"custom20,omitempty"`
/*
Specifies a value or a range of values of the custom20 property to match.<br><br>

|Symbol|Value|Description|
|:---|:----|:---|
| = | Equal to | The exact custom20 value to be searched for. For example, *custom20=10335*.|
| &lt; | Less than | The custom20 value to be searched for is less than the given value. For example, *custom20=&lt;10335*.|
| &gt; | Greater than | The custom20 value to be searched for is greater than the given value. For example, *custom20=&gt;10335*.|
| - | Range | The range of custom20 values. This search includes both the lower and higher values specified in the range. This range must be in ascending order. For example, *custom20=10330-10400*.|
*/

    Custom21From string `url:"custom21_from,omitempty"`
/*
Filters documents based on `custom21` values dated on or after this date (in ISO 8601 format).
*/

    Custom21Relative string `url:"custom21_relative,omitempty"`
/*
Filters documents by `custom21` value within a date range relative to the current date.

For more information about date range notation, see <a href="#overview--date-range-format">Date range format</a> in the **Key Concepts** section.
*/

    Custom21To string `url:"custom21_to,omitempty"`
/*
Filters documents based on `custom21` values dated on or before this date (in ISO 8601 format).
*/

    Custom22From string `url:"custom22_from,omitempty"`
/*
Filters documents based on `custom22` values dated on or after this date (in ISO 8601 format).
*/

    Custom22Relative string `url:"custom22_relative,omitempty"`
/*
Filters documents by `custom22` value within a date range relative to the current date.

For more information about date range notation, see <a href="#overview--date-range-format">Date range format</a> in the **Key Concepts** section.
*/

    Custom22To string `url:"custom22_to,omitempty"`
/*
Filters documents based on `custom22` values dated on or before this date (in ISO 8601 format).
*/

    Custom23From string `url:"custom23_from,omitempty"`
/*
Filters documents based on `custom23` values dated on or after this date (in ISO 8601 format).
*/

    Custom23Relative string `url:"custom23_relative,omitempty"`
/*
Filters documents by `custom23` value within a date range relative to the current date.

For more information about date range notation, see <a href="#overview--date-range-format">Date range format</a> in the **Key Concepts** section.
*/

    Custom23To string `url:"custom23_to,omitempty"`
/*
Filters documents based on `custom23` values dated on or before this date (in ISO 8601 format).
*/

    Custom24From string `url:"custom24_from,omitempty"`
/*
Filters documents based on `custom24` values dated on or after this date (in ISO 8601 format).
*/

    Custom24Relative string `url:"custom24_relative,omitempty"`
/*
Filters documents by `custom24` value within a date range relative to the current date.

For more information about date range notation, see <a href="#overview--date-range-format">Date range format</a> in the **Key Concepts** section.
*/

    Custom24To string `url:"custom24_to,omitempty"`
/*
Filters documents based on `custom24` values dated on or before this date (in ISO 8601 format).
*/

    Custom25 bool `url:"custom25,omitempty"`
/*
Filters documents based on the value of the `custom25` property.
*/

    Custom26 bool `url:"custom26,omitempty"`
/*
Filters documents based on the value of the `custom26` property.
*/

    Custom27 bool `url:"custom27,omitempty"`
/*
Filters documents based on the value of the `custom27` property.
*/

    Custom28 bool `url:"custom28,omitempty"`
/*
Filters documents based on the value of the `custom28` property.
*/

    Custom29 string `url:"custom29,omitempty"`
/*
Filters documents based on the alias of the `custom29` property.<br>
*/

    Custom3 string `url:"custom3,omitempty"`
/*
Filters documents based on the alias of the `custom3` property.<br>
* Partial matches are not supported.
* The wildcard character (*) is not supported.
*/

    Custom30 string `url:"custom30,omitempty"`
/*
Filters documents based on the alias of the `custom30` property.<br>
*/

    Custom4 string `url:"custom4,omitempty"`
/*
Filters documents based on the alias of the `custom4` property.<br>
* Partial matches are not supported.
* The wildcard character (*) is not supported.
*/

    Custom5 string `url:"custom5,omitempty"`
/*
Filters documents based on the alias of the `custom5` property.<br>
* Partial matches are not supported.
* The wildcard character (*) is not supported.
*/

    Custom6 string `url:"custom6,omitempty"`
/*
Filters documents based on the alias of the `custom6` property.<br>
* Partial matches are not supported.
* The wildcard character (*) is not supported.
*/

    Custom7 string `url:"custom7,omitempty"`
/*
Filters documents based on the alias of the `custom7` property.<br>
* Partial matches are not supported.
* The wildcard character (*) is not supported.
*/

    Custom8 string `url:"custom8,omitempty"`
/*
Filters documents based on the alias of the `custom8` property.<br>
* Partial matches are not supported.
* The wildcard character (*) is not supported.
*/

    Custom9 string `url:"custom9,omitempty"`
/*
Filters documents based on the alias of the `custom9` property.<br>
* Partial matches are not supported.
* The wildcard character (*) is not supported.
*/

    DocumentNumber string `url:"document_number,omitempty"`
/*
Filters documents based on their numeric identifier, known as a document number.<br><br>
The document number appears as part of the document ID.
For example, if the document ID is `active!72634.3`, the document number is 72634.<br>

This filter supports specifying document numbers as follows:
* A single value.
* Multiple document numbers as a comma-separated list.
  For example, `document_number=7545,674332,6543`.
* A value less than (<), greater than(>), less than or equal(<=), and greater than or equal (>=)
  For example, `document_number=<=7545`.
* A range of document numbers. For example, `document_number=7545-674332`.

*/

    DocumentVersion string `url:"document_version,omitempty"`
/*
Filters documents based on the document version.<br>
When specifying document versions, either a single value (`document_number=3`), or a comma-separated list (`document_number=1,3,5`) can be used to specify multiple document version numbers.

This filter supports specifying document versions as follows:
* A single value.
* Multiple document versions as a comma-separated list.
  For example, `document_version=2,5,8`.
* A value less than (<), greater than(>), less than or equal(<=), and greater than or equal (>=)
  For example, `document_version=>=4`.
* A range of document versions. For example, `document_version=2-5`.
*/

    EditDate string `url:"edit_date,omitempty"`
/*
Filters documents by edit date within a date range relative to the current date.

For more information about date range notation, see <a href="#overview--date-range-format">Date range format</a> in the **Key Concepts** section.
*/

    EditDateFrom string `url:"edit_date_from,omitempty"`
/*
Filters documents that were edited on or after this date (in ISO 8601 format).
*/

    EditDateTo string `url:"edit_date_to,omitempty"`
/*
Filters documents that were edited on or before this date (in ISO 8601 format).
*/

    EmailOnly bool `url:"email_only,omitempty"`
/*
Filters documents to include only emails and to exclude documents.<br><br>
If `true`, only emails are included.<br>
If `false`, include emails and documents.<br>

To return only documents, see the `exclude_emails` parameter.

*/

    ExcludeDocs bool `url:"exclude_docs,omitempty"`
/*
Specifies to exclude documents from the results set.<br>

If `true`, excludes documents from the results set.<br>
If `false`, includes documents in the results set.<br>

See `email_only` parameter to include only emails in the results set.
See `exclude_emails` to exclude emails from the results set.

*/

    ExcludeEmails bool `url:"exclude_emails,omitempty"`
/*
Filters documents to exclude emails.<br><br>
If `true`, exclude emails and return only documents.<br>
If `false`, include emails and documents.<br>

To return only emails, see the `email_only` parameter.

*/

    ExcludeShortcuts bool `url:"exclude_shortcuts,omitempty"`
/*
Filters documents to exclude document shortcuts.<br><br>
If `true`, exclude document shortcuts.<br>
If `false`, include document shortcuts.

*/

    FileTarget string `url:"file_target,omitempty"`
/*
Filters documents based on the value of any of the following properties:
* `custom1`
* `custom2`
* `custom29`
* `custom30`
* `workspace_name`

*/

    HasAttachment bool `url:"has_attachment,omitempty"`
/*
Filters emails that have attachments.<br><br>
If `true`, returns emails with at least one attachment.<br>
If `false`, returns all emails with or without attachments.

*/

    InUse bool `url:"in_use,omitempty"`
/*
Filters documents based on whether or not they are in use.

If `true`, returns documents marked as in use.<br>
If `false`, returns documents marked as not in use.<br>
If omitted, returns documents regardless of whether they are in use.

*/

    InUseBy string `url:"in_use_by,omitempty"`
/*
Filters documents that are in use by the specified user ID.
*/

    LastUser string `url:"last_user,omitempty"`
/*
Filters documents based on the user ID of the last user of documents.
*/

    Limit int `url:"limit,omitempty"`
/*
Specifies the maximum number of items to include in the response.

The request returns the actual number of items up to the limit value (inclusive). If there are more items than the limit value, no more items than the limit are returned and a cursor value is returned.

This parameter can be used in conjunction with pagination parameters for endpoints that support them. For more information, refer to <a href="#overview--pagination">Pagination</a> in the **Key Concepts** section.
*/

    Name string `url:"name,omitempty"`
/*
Filters documents based on the matching text found in the name of documents.

* The search matches with whole words only.<br>
* The wildcard character (*) is supported.

*/

    Offset int `url:"offset,omitempty"`
/*
Deprecated. Do not use.
*/

    Operator string `url:"operator,omitempty"`
/*
Filters documents based on the user ID of the operator of the documents.
*/

    Owner string `url:"owner,omitempty"`
/*
Filters documents based on the user ID of the owner of the documents.
*/

    PagingMode string `url:"paging_mode,omitempty"`
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

**NOTE:** The following exceptions to the above occur when the `paging_mode` is specified as `standard_cursor`:
* The filter `total` is ignored.
* The response field `total_count` is not returned.

**NOTE:** `standard_cursor` mode is not supported for the following:
* Saved search folders
* Index searches

For more information about index searches, refer to <a href="#overview--indexing-based-on-document-class">Indexing based on document class</a> in the **Key Concepts** section.

*/

    ReceivedDate string `url:"received_date,omitempty"`
/*
Filters emails by received date within a date range relative to the current date.

For more information about date range notation, see <a href="#overview--date-range-format">Date range format</a> in the **Key Concepts** section.
*/

    ReceivedDateFrom string `url:"received_date_from,omitempty"`
/*
Filters emails received on or after this date (in ISO 8601 format).
Specifies to search for emails received on or after this date (in ISO 8601 format).
*/

    ReceivedDateTo string `url:"received_date_to,omitempty"`
/*
Filters emails received on or before this date (in ISO 8601 format).
Specifies to search for emails received on or before this date (in ISO 8601 format).
*/

    Recipient string `url:"recipient,omitempty"`
/*
Filters documents based on the text that appears in recipient field of emails.
*/

    Sender string `url:"sender,omitempty"`
/*
Filters documents based on the text that appears in the sender field of emails.
*/

    SentDate string `url:"sent_date,omitempty"`
/*
Filters emails by sent date within a date range relative to the current date.

For more information about date range notation, see <a href="#overview--date-range-format">Date range format</a> in the **Key Concepts** section.
*/

    SentDateFrom string `url:"sent_date_from,omitempty"`
/*
Filters emails sent on or after this date (in ISO 8601 format).
*/

    SentDateTo string `url:"sent_date_to,omitempty"`
/*
Filters emails sent on or before this date (in ISO 8601 format).
*/

    Subject string `url:"subject,omitempty"`
/*
Filters emails based on the matching text found in the email subject.
 * The wildcard character (*) is supported.
 * The search is not case-sensitive.
 * The search is fuzzy and returns results in the order of relevance.
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

    Type string `url:"type,omitempty"`
/*
Specifies the document type.

Each iManage Work system has a customizable list of document types per the organizational requirements.
For example, *Word*, *WordX*, *Acrobat*, *Powerpoint*. <br>
These values are not necessarily the same as the document type's extension field, which may be *doc*, *docx*, *pdf*, or *ppt*.<br>
Document type is not case sensitive, that is, *Word* and *WORD* matches with *word*.<br>
There is no stemming and the wildcard character (&ast;) is not supported.
For example, *acrob* or *acrob&#42;* does not match *acrobat*.<br>
Multiple document types can be included using a comma-separated list.
For example, *type=WORD,WORDX,WORDXT*

To get a complete list of document types for a library, see GET `/customers/{customerId}/libraries/{libraryId}/types`.
*/

    User string `url:"user,omitempty"`
/*
Filters documents based on the document's user ID.
*/

}
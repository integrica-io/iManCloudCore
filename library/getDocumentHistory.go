package library

import (
	"context"
	"time"
	"fmt"

	"github.com/integrica-io/iManCloudCore/client"
	"github.com/integrica-io/iManCloudCore/internal"
)

func GetDocumentHistory(ctx context.Context, client *client.Client, libraryId string, documentId string, options *GetDocumentHistoryOptions) (GetDocumentHistoryOutput, error) {
	var data GetDocumentHistoryOutput
	endpoint := client.BaseUrl.JoinPath("work", "api", "v2", "customers", client.TokenCfg.CustomerId, "libraries", libraryId, "documents", fmt.Sprintf("%s!%s", libraryId, documentId), "history")

	req := internal.HttpRequestBuilder{}
	req.Context(ctx).Url(*endpoint).Method(internal.Get).ToJson(&data)

	if err := client.Req(req); err != nil {
		return data, err
	}
	return data, nil
}

type GetDocumentHistoryOutput struct {
	Data []GetDocumentHistoryOutputData `json:"data"`
	TotalCount int `json:"total_count"`
}

type GetDocumentHistoryOutputData struct{
	Activity        string    `json:"activity"`
	ActivityCode    int       `json:"activity_code"`
	ActivityDate    time.Time `json:"activity_date"`
	ApplicationID   string    `json:"application_id"`
	ApplicationName string    `json:"application_name"`
	DocumentNumber  int       `json:"document_number"`
	HasJournal      bool      `json:"has_journal"`
	Location        string    `json:"location"`
	Num1            string    `json:"num1"`
	ID              string    `json:"id"`
	User            string    `json:"user"`
	UserFullname    string    `json:"user_fullname"`
	Version         int       `json:"version"`
	Comments        string    `json:"comments,omitempty"`
}

type GetDocumentHistoryActivity string

const (
	Copy GetDocumentHistoryActivity 	= "copy"
	Download GetDocumentHistoryActivity	= "download"
	Edit GetDocumentHistoryActivity		= "edit"
	Email GetDocumentHistoryActivity	= "email"
	Print GetDocumentHistoryActivity	= "print"
	Share GetDocumentHistoryActivity	= "share"
	Version GetDocumentHistoryActivity	= "version"
	View GetDocumentHistoryActivity		= "view"
	All GetDocumentHistoryActivity		= "all"
)

type GetDocumentHistoryOptions struct { 
    Activity GetDocumentHistoryActivity `url:"activity,omitempty"`
/*
Filters documents based on the activities that were performed on them.

The following table contains the list of allowed values for the `activity` query parameter:

|Query parameter|Document activities|Description|
|---|---|---|
|copy|COPYDOC|The document was copied.|
|download|EXPORT|The document was exported.|
|edit|NEWDOCUMENT, NEWDOCVER, MODIFY|The document was created, its version was updated, or its content was modified.|
|email|MAIL|The document was emailed.|
|print|PRINT|The document was printed.|
|share|SHARED, SHARE_EXPIRATION_MODIFIED, SHARE_REVOKED|The document was shared, the share's expiry has changed, or the share has been revoked|
|version|NEWDOCVER|The version of the document was updated.|
|view|VIEW|The document was viewed.|
|all| |All allowable activity values.|

*/

    ActivityDateEnd string `url:"activity_date_end,omitempty"`
/*
Filters document activities based on the end date and time (in ISO 8601 format) of an activity.

The date can be specified in one of two formats.
* The standard iManage ISO 8601 format. This is `yyyy-mm-ddThh:mm:ss.µZ` where `µ` represents zero or more microseconds and `Z` is the zone
designator for zero UTC offset. For example: `2019-10-29T11:44:24.229Z`
* The standard iManage ISO 8601 format with offset. This is `yyyy-mm-ddThh:mm:ss.µ∓hh.mm` where `µ` represents zero or more microseconds. For example: `2019-10-29T11:44:24.229+02.00`.
    * The `∓hh.mm` value at the end represents the timezone offset. For example, India is five and half hours ahead of London, so a date from London can be set match India Standard Time using `+05.30` as in `2019-10-29T11:44:24.229+05.30`.

*/

    ActivityDateStart string `url:"activity_date_start,omitempty"`
/*
Filters document activities based on the start date and time (in ISO 8601 format) of an activity.

The date can be specified in one of two formats.
* The standard iManage ISO 8601 format. This is `yyyy-mm-ddThh:mm:ss.µZ` where `µ` represents zero or more microseconds and `Z` is the zone
designator for zero UTC offset. For example: `2019-10-29T11:44:24.229Z`
* The standard iManage ISO 8601 format with offset. This is `yyyy-mm-ddThh:mm:ss.µ∓hh.mm` where `µ` represents zero or more microseconds. For example: `2019-10-29T11:44:24.229+02.00`.
    * The `∓hh.mm` value at the end represents the timezone offset. For example, India is five and half hours ahead of London, so a date from London can be set match India Standard Time using `+05.30` as in `2019-10-29T11:44:24.229+05.30`.

*/

    AllVersions bool `url:"all_versions,omitempty"`
/*
Returns the history of all the versions of a document.

If `true`, returns history of all the versions of a document.<br>
If `false`, returns history of only the version specified in the `docId`.

*/

    Limit int `url:"limit,omitempty"`
/*
Specifies the maximum number of items to include in the response.

The request returns the actual number of items up to the limit value (inclusive). If there are more items than the limit value, no more items than the limit are returned and a cursor value is returned.

This parameter can be used in conjunction with pagination parameters for endpoints that support them. For more information, refer to <a href="#overview--pagination">Pagination</a> in the **Key Concepts** section.
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

    User string `url:"user,omitempty"`
/*
Filters document activities based on the user who performed them.
*/

}
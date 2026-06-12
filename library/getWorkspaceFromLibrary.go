package library

import (
	"time"
	"context"
	
	"github.com/integrica-io/iManCloudCore/internal"
	"github.com/integrica-io/iManCloudCore/client"
	"github.com/google/go-querystring/query"	
)

func GetWorkspaceFromLibrary(ctx context.Context, client *client.Client, libraryId string, options *GetWorkspaceFromLibraryOptions)(GetWorkspaceFromLibraryOutput, error){
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
/*
Specifies the text to be searched anywhere in the profile or contents.<br><br>
This applies only to workspace and document searches.
For workspace template searches, this parameter is the same as the parameter `name_or_description`.

This search is case insensitive and matches whole words. The wildcard character (&ast;) is not supported.
This parameter searches indexed properties and the user is required to have their role's **Allow Full Text Search** privilege enabled.

*/

    CreateDate string `url:"create_date,omitempty"`
/*
Specifies a date range relative to the current date in which the workspace was created.

For more information about date range notation, see <a href="#overview--date-range-format">Date range format</a> in the **Key Concepts** section.

*/

    Cursor string `url:"cursor,omitempty"`
/*
Specifies the cursor value to retrieve the next set of results.

If a cursor is not specified, the result set is from the beginning. For more information, refer to <a href="#overview--pagination">Pagination</a> in the **Key Concepts** section.

*/

    Custom1 string `url:"custom1,omitempty"`
/*
Specifies an alias to match with this custom property.

An alias is an entry for that custom property.
For example, custom1 may be captioned as *Client*.
Client may have three entries&#58; 001, 001abc, and 001001.
Each of these is an alias of client.

This search is case insensitive.
Partial matches are not supported.
The wildcard character (&ast;) is not supported.
*/

    Custom10 string `url:"custom10,omitempty"`
/*
Specifies an alias to match with this custom property.

An alias is an entry for that custom property.
For example, custom1 may be captioned as *Client*.
Client may have three entries&#58; 001, 001abc, and 001001.
Each of these is an alias of client.

This search is case insensitive.
Partial matches are not supported.
The wildcard character (&ast;) is not supported.
*/

    Custom11 string `url:"custom11,omitempty"`
/*
Specifies an alias to match with this custom property.

An alias is an entry for that custom property.
For example, custom1 may be captioned as *Client*.
Client may have three entries&#58; 001, 001abc, and 001001.
Each of these is an alias of client.

This search is case insensitive.
Partial matches are not supported.
The wildcard character (&ast;) is not supported.
*/

    Custom12 string `url:"custom12,omitempty"`
/*
Specifies an alias to match with this custom property.

An alias is an entry for that custom property.
For example, custom1 may be captioned as *Client*.
Client may have three entries&#58; 001, 001abc, and 001001.
Each of these is an alias of client.

This search is case insensitive.
Partial matches are not supported.
The wildcard character (&ast;) is not supported.
*/

    Custom13 string `url:"custom13,omitempty"`
/*
Specifies the text to search for in custom13 property.<br><br>
For workspaces, the use of this property is customized for your iManage Work installation.

*/

    Custom14 string `url:"custom14,omitempty"`
/*
Specifies the text to search for in custom14 property.<br><br>
For workspaces, the use of this property is customized for your iManage Work installation.

*/

    Custom15 string `url:"custom15,omitempty"`
/*
Specifies the text to search for in custom15 property.<br><br>
For workspaces, the use of this property is customized for your iManage Work installation.

*/

    Custom16 string `url:"custom16,omitempty"`
/*
Specifies the text to search for in custom16 property.<br><br>
 For workspaces, the use of this property is customized for your iManage Work installation.

*/

    Custom17 string `url:"custom17,omitempty"`
/*
Specifies a value for the custom17 property.<br><br>
The use of this property is customized for your iManage Work installation.<br><br>
The date can be specified in one of two formats.
* The standard iManage ISO 8601 format. This is `yyyy-mm-ddThh:mm:ss.µZ` where `µ` represents zero or more microseconds and `Z` is the zone
designator for zero UTC offset. For example: `2019-10-29T11:44:24.229Z`
* The standard iManage ISO 8601 format with offset. This is `yyyy-mm-ddThh:mm:ss.µ∓hh.mm` where `µ` represents zero or more microseconds. For example: `2019-10-29T11:44:24.229+02.00`.
    * The `∓hh.mm` value at the end represents the timezone offset. For example, India is five and half hours ahead of London, so a date from London can be set match India Standard Time using `+05.30` as in `2019-10-29T11:44:24.229+05.30`.

*/

    Custom18 string `url:"custom18,omitempty"`
/*
Specifies a value for the custom18 property.<br><br>
The use of this property is customized for your iManage Work installation.<br><br>
The date can be specified in one of two formats.
* The standard iManage ISO 8601 format. This is `yyyy-mm-ddThh:mm:ss.µZ` where `µ` represents zero or more microseconds and `Z` is the zone
designator for zero UTC offset. For example: `2019-10-29T11:44:24.229Z`
* The standard iManage ISO 8601 format with offset. This is `yyyy-mm-ddThh:mm:ss.µ∓hh.mm` where `µ` represents zero or more microseconds. For example: `2019-10-29T11:44:24.229+02.00`.
    * The `∓hh.mm` value at the end represents the timezone offset. For example, India is five and half hours ahead of London, so a date from London can be set match India Standard Time using `+05.30` as in `2019-10-29T11:44:24.229+05.30`.

*/

    Custom19 string `url:"custom19,omitempty"`
/*
Specifies a value for the custom19 property.<br><br>
The use of this property is customized for your iManage Work installation.<br><br>
The date can be specified in one of two formats.
* The standard iManage ISO 8601 format. This is `yyyy-mm-ddThh:mm:ss.µZ` where `µ` represents zero or more microseconds and `Z` is the zone
designator for zero UTC offset. For example: `2019-10-29T11:44:24.229Z`
* The standard iManage ISO 8601 format with offset. This is `yyyy-mm-ddThh:mm:ss.µ∓hh.mm` where `µ` represents zero or more microseconds. For example: `2019-10-29T11:44:24.229+02.00`.
    * The `∓hh.mm` value at the end represents the timezone offset. For example, India is five and half hours ahead of London, so a date from London can be set match India Standard Time using `+05.30` as in `2019-10-29T11:44:24.229+05.30`.

*/

    Custom2 string `url:"custom2,omitempty"`
/*
Specifies an alias to match with this custom property.

An alias is an entry for that custom property.
For example, custom1 may be captioned as *Client*.
Client may have three entries&#58; 001, 001abc, and 001001.
Each of these is an alias of client.

This search is case insensitive.
Partial matches are not supported.
The wildcard character (&ast;) is not supported.
*/

    Custom20 string `url:"custom20,omitempty"`
/*
Specifies a value for the custom20 property.<br><br>
The use of this property is customized for your iManage Work installation.<br><br>
The date can be specified in one of two formats.
* The standard iManage ISO 8601 format. This is `yyyy-mm-ddThh:mm:ss.µZ` where `µ` represents zero or more microseconds and `Z` is the zone
designator for zero UTC offset. For example: `2019-10-29T11:44:24.229Z`
* The standard iManage ISO 8601 format with offset. This is `yyyy-mm-ddThh:mm:ss.µ∓hh.mm` where `µ` represents zero or more microseconds. For example: `2019-10-29T11:44:24.229+02.00`.
    * The `∓hh.mm` value at the end represents the timezone offset. For example, India is five and half hours ahead of London, so a date from London can be set match India Standard Time using `+05.30` as in `2019-10-29T11:44:24.229+05.30`.

*/

    Custom21From string `url:"custom21_from,omitempty"`
/*
Specifies the lower limit of the custom21 date to search from.<br><br>
This property is not reserved by iManage Work and is not a validated value.
It is intended to be customized for your company's requirements.
This is a date/time property in an ISO 8601 format value. For example, `2018-05-24T00:00:00-08:00`.<br>

When searching, this parameter specifies the earliest value to search from.
For example, specifying 2018-05-24T00:00:00-08:00 (Saturday May 24, 2018 00:00:00 PST) returns only items
created on or after 2018-05-24T00:00:00-08:00.
If omitted or is invalid, there is no lower limit, returning items based on other search criteria.
You can specify a date/time range by using *property_name*_from for the lower limit and *property_name*_to for the upper limit.
This is applicable for each custom21, custom22, custom23, custom24.
For example, `custom21_from=2018-05-24T00:00:00-08:00&custom21_to=2018-11-24T00:00:00-08:00`.

*/

    Custom21Relative string `url:"custom21_relative,omitempty"`
/*
Specifies a date/time range to search, relative to the current date/time.<br><br>
This property is not reserved by iManage Work and is not a validated value.
It is intended to be customized for your company's requirements.

For more information about date range notation, see <a href="#overview--date-range-format">Date range format</a> in the **Key Concepts** section.

*/

    Custom21To string `url:"custom21_to,omitempty"`
/*
Specifies the lower limit of the custom21 date to search to.<br><br>
This property is not reserved by iManage Work and is not a validated value.
It is intended to be customized for your company's requirements.
This is a date/time property in an ISO 8601 format value. For example, `2018-05-24T00:00:00-08:00`.<br>

When searching, this parameter specifies the latest value to search to.
For example, specifying 2018-05-24T00:00:00-08:00 (Saturday May 24, 2018 00:00:00 PST) returns only items
with a property value before 2018-05-24T00:00:00-08:00.
If omitted or is invalid, there is no upper limit, returning items based on other search criteria.
You can specify a date/time range by using *property_name*_from for the lower limit and *property_name*_to for the upper limit.
This is applicable for each custom21, custom22, custom23, custom24.
For example, `custom21_from=2018-05-24T00:00:00-08:00&custom21_to=2018-11-24T00:00:00-08:00`.

*/

    Custom22From string `url:"custom22_from,omitempty"`
/*
Specifies the lower limit of the custom22 date to search from.<br><br>
This property is not reserved by iManage Work and is not a validated value.
It is intended to be customized for your company's requirements.
This is a date/time property in an ISO 8601 format value. For example, `2018-05-24T00:00:00-08:00`.<br>

When searching, this parameter specifies the earliest value to search from.
For example, specifying 2018-05-24T00:00:00-08:00 (Saturday May 24, 2018 00:00:00 PST) returns only items
created on or after 2018-05-24T00:00:00-08:00.
If omitted or is invalid, there is no lower limit, returning items based on other search criteria.
You can specify a date/time range by using *property_name*_from for the lower limit and *property_name*_to for the upper limit.
This is applicable for each custom21, custom22, custom23, custom24.
For example, `custom21_from=2018-05-24T00:00:00-08:00&custom21_to=2018-11-24T00:00:00-08:00`.

*/

    Custom22Relative string `url:"custom22_relative,omitempty"`
/*
Specifies a date/time range to search, relative to the current date/time.<br><br>
This property is not reserved by iManage Work and is not a validated value.
It is intended to be customized for your company's requirements.

For more information about date range notation, see <a href="#overview--date-range-format">Date range format</a> in the **Key Concepts** section.

*/

    Custom22To string `url:"custom22_to,omitempty"`
/*
Specifies the lower limit of the custom22 date to search to.<br><br>
This property is not reserved by iManage Work and is not a validated value.
It is intended to be customized for your company's requirements.
This is a date/time property in an ISO 8601 format value. For example, `2018-05-24T00:00:00-08:00`.<br>

When searching, this parameter specifies the latest value to search to.
For example, specifying 2018-05-24T00:00:00-08:00 (Saturday May 24, 2018 00:00:00 PST) returns only items
with a property value before 2018-05-24T00:00:00-08:00.
If omitted or is invalid, there is no upper limit, returning items based on other search criteria.
You can specify a date/time range by using *property_name*_from for the lower limit and *property_name*_to for the upper limit.
This is applicable for each custom21, custom22, custom23, custom24.
For example, `custom21_from=2018-05-24T00:00:00-08:00&custom21_to=2018-11-24T00:00:00-08:00`.

*/

    Custom23From string `url:"custom23_from,omitempty"`
/*
Specifies the lower limit of the custom23 date to search from.<br><br>
This property is not reserved by iManage Work and is not a validated value.
It is intended to be customized for your company's requirements.
This is a date/time property in an ISO 8601 format value. For example, `2018-05-24T00:00:00-08:00`.<br>

When searching, this parameter specifies the earliest value to search from.
For example, specifying 2018-05-24T00:00:00-08:00 (Saturday May 24, 2018 00:00:00 PST) returns only items
created on or after 2018-05-24T00:00:00-08:00.
If omitted or is invalid, there is no lower limit, returning items based on other search criteria.
You can specify a date/time range by using *property_name*_from for the lower limit and *property_name*_to for the upper limit.
This is applicable for each custom21, custom22, custom23, custom24.
For example, `custom21_from=2018-05-24T00:00:00-08:00&custom21_to=2018-11-24T00:00:00-08:00`.

*/

    Custom23Relative string `url:"custom23_relative,omitempty"`
/*
Specifies a date/time range to search, relative to the current date/time.<br><br>
This property is not reserved by iManage Work and is not a validated value.
It is intended to be customized for your company's requirements.

For more information about date range notation, see <a href="#overview--date-range-format">Date range format</a> in the **Key Concepts** section.

*/

    Custom23To string `url:"custom23_to,omitempty"`
/*
Specifies the lower limit of the custom23 date to search to.<br><br>
This property is not reserved by iManage Work and is not a validated value.
It is intended to be customized for your company's requirements.
This is a date/time property in an ISO 8601 format value. For example, `2018-05-24T00:00:00-08:00`.<br>

When searching, this parameter specifies the latest value to search to.
For example, specifying 2018-05-24T00:00:00-08:00 (Saturday May 24, 2018 00:00:00 PST) returns only items
with a property value before 2018-05-24T00:00:00-08:00.
If omitted or is invalid, there is no upper limit, returning items based on other search criteria.
You can specify a date/time range by using *property_name*_from for the lower limit and *property_name*_to for the upper limit.
This is applicable for each custom21, custom22, custom23, custom24.
For example, `custom21_from=2018-05-24T00:00:00-08:00&custom21_to=2018-11-24T00:00:00-08:00`.

*/

    Custom24From string `url:"custom24_from,omitempty"`
/*
Specifies the lower limit of the custom24 date to search from.<br><br>
This property is not reserved by iManage Work and is not a validated value.
It is intended to be customized for your company's requirements.
This is a date/time property in an ISO 8601 format value. For example, `2018-05-24T00:00:00-08:00`.<br>

When searching, this parameter specifies the earliest value to search from.
For example, specifying 2018-05-24T00:00:00-08:00 (Saturday May 24, 2018 00:00:00 PST) returns only items
created on or after 2018-05-24T00:00:00-08:00.
If omitted or is invalid, there is no lower limit, returning items based on other search criteria.
You can specify a date/time range by using *property_name*_from for the lower limit and *property_name*_to for the upper limit.
This is applicable for each custom21, custom22, custom23, custom24.
For example, `custom21_from=2018-05-24T00:00:00-08:00&custom21_to=2018-11-24T00:00:00-08:00`.

*/

    Custom24Relative string `url:"custom24_relative,omitempty"`
/*
Specifies a date/time range to search, relative to the current date/time.<br><br>
This property is not reserved by iManage Work and is not a validated value.
It is intended to be customized for your company's requirements.

For more information about date range notation, see <a href="#overview--date-range-format">Date range format</a> in the **Key Concepts** section.

*/

    Custom24To string `url:"custom24_to,omitempty"`
/*
Specifies the lower limit of the custom24 date to search to.<br><br>
This property is not reserved by iManage Work and is not a validated value.
It is intended to be customized for your company's requirements.
This is a date/time property in an ISO 8601 format value. For example, `2018-05-24T00:00:00-08:00`.<br>

When searching, this parameter specifies the latest value to search to.
For example, specifying 2018-05-24T00:00:00-08:00 (Saturday May 24, 2018 00:00:00 PST) returns only items
with a property value before 2018-05-24T00:00:00-08:00.
If omitted or is invalid, there is no upper limit, returning items based on other search criteria.
You can specify a date/time range by using *property_name*_from for the lower limit and *property_name*_to for the upper limit.
This is applicable for each custom21, custom22, custom23, custom24.
For example, `custom21_from=2018-05-24T00:00:00-08:00&custom21_to=2018-11-24T00:00:00-08:00`.

*/

    Custom25 bool `url:"custom25,omitempty"`
/*
Specifies the custom25 property value to search for.<br><br>
This property is not reserved by iManage Work and is not a validated value.
It is intended to be customized for your company's requirements.
This is a boolean value but can also be null.<br>

Examples:<br>
* `custom25=true` returns items if the custom25 value is `true`.
* `custom25=false`, returns items if the custom25 value is `false`.

*/

    Custom26 bool `url:"custom26,omitempty"`
/*
Specifies the custom26 property value to search for.<br><br>
This property is not reserved by iManage Work and is not a validated value.
It is intended to be customized for your company's requirements.
This is a boolean value but can also be null.<br>

Examples:<br>
* `custom26=true` returns items if the custom26 value is `true`.
* `custom26=false`, returns items if the custom26 value is `false`.

*/

    Custom27 bool `url:"custom27,omitempty"`
/*
Specifies the custom27 property value to search for.<br><br>
This property is not reserved by iManage Work and is not a validated value.
It is intended to be customized for your company's requirements.
This is a boolean value but can also be null.<br>

Examples:<br>
* `custom27=true` returns items if the custom27 value is `true`.
* `custom27=false`, returns items if the custom27 value is `false`.

*/

    Custom28 bool `url:"custom28,omitempty"`
/*
Specifies the custom28 property value to search for.<br><br>
This property is not reserved by iManage Work and is not a validated value.
It is intended to be customized for your company's requirements.
This is a boolean value but can also be null.<br>

Examples:<br>
* `custom28=true` returns items if the custom28 value is `true`.
* `custom28=false`, returns items if the custom28 value is `false`.

*/

    Custom29 string `url:"custom29,omitempty"`
/*
Specifies an alias to match with this custom property.

An alias is an entry for that custom property.
For example, custom1 may be captioned as *Client*.
Client may have three entries&#58; 001, 001abc, and 001001.
Each of these is an alias of client.

This search is case insensitive.
Partial matches are not supported.
The wildcard character (&ast;) is not supported.
*/

    Custom3 string `url:"custom3,omitempty"`
/*
Specifies an alias to match with this custom property.

An alias is an entry for that custom property.
For example, custom1 may be captioned as *Client*.
Client may have three entries&#58; 001, 001abc, and 001001.
Each of these is an alias of client.

This search is case insensitive.
Partial matches are not supported.
The wildcard character (&ast;) is not supported.
*/

    Custom30 string `url:"custom30,omitempty"`
/*
Specifies an alias to match with this custom property.

An alias is an entry for that custom property.
For example, custom1 may be captioned as *Client*.
Client may have three entries&#58; 001, 001abc, and 001001.
Each of these is an alias of client.

This search is case insensitive.
Partial matches are not supported.
The wildcard character (&ast;) is not supported.
*/

    Custom4 string `url:"custom4,omitempty"`
/*
Specifies an alias to match with this custom property.

An alias is an entry for that custom property.
For example, custom1 may be captioned as *Client*.
Client may have three entries&#58; 001, 001abc, and 001001.
Each of these is an alias of client.

This search is case insensitive.
Partial matches are not supported.
The wildcard character (&ast;) is not supported.
*/

    Custom5 string `url:"custom5,omitempty"`
/*
Specifies an alias to match with this custom property.

An alias is an entry for that custom property.
For example, custom1 may be captioned as *Client*.
Client may have three entries&#58; 001, 001abc, and 001001.
Each of these is an alias of client.

This search is case insensitive.
Partial matches are not supported.
The wildcard character (&ast;) is not supported.
*/

    Custom6 string `url:"custom6,omitempty"`
/*
Specifies an alias to match with this custom property.

An alias is an entry for that custom property.
For example, custom1 may be captioned as *Client*.
Client may have three entries&#58; 001, 001abc, and 001001.
Each of these is an alias of client.

This search is case insensitive.
Partial matches are not supported.
The wildcard character (&ast;) is not supported.
*/

    Custom7 string `url:"custom7,omitempty"`
/*
Specifies an alias to match with this custom property.

An alias is an entry for that custom property.
For example, custom1 may be captioned as *Client*.
Client may have three entries&#58; 001, 001abc, and 001001.
Each of these is an alias of client.

This search is case insensitive.
Partial matches are not supported.
The wildcard character (&ast;) is not supported.
*/

    Custom8 string `url:"custom8,omitempty"`
/*
Specifies an alias to match with this custom property.

An alias is an entry for that custom property.
For example, custom1 may be captioned as *Client*.
Client may have three entries&#58; 001, 001abc, and 001001.
Each of these is an alias of client.

This search is case insensitive.
Partial matches are not supported.
The wildcard character (&ast;) is not supported.
*/

    Custom9 string `url:"custom9,omitempty"`
/*
Specifies an alias to match with this custom property.

An alias is an entry for that custom property.
For example, custom1 may be captioned as *Client*.
Client may have three entries&#58; 001, 001abc, and 001001.
Each of these is an alias of client.

This search is case insensitive.
Partial matches are not supported.
The wildcard character (&ast;) is not supported.
*/

    Description string `url:"description,omitempty"`
/*
Specifies the text to search for in the `description` field of the workspace's profile.<br><br>
This parameter is case insensitive.
Matches whole words (partial matches are not supported).
For example, (`description=litigation Enron`) matches either *litigation* or *Enron* in any of the fields.
The search with `description=litigation` matches only *litigation*.
The search with `description=litiga` does not match any variation.
Exact words may be matched with quotation marks (`name="litigation Enron"`).
The wildcard character (&ast;) is not supported.
For example, `file_target=*gation`, or `file_target=litiga*` does not match *litigation*.
The underscore character is treated as a space.
The search with `file_target=litigation_Enron` is identical to `file_target=litigation Enron`.

Special characters allowed are *().&-_[]`~\|@$%^?:{}!',/\\#+<>;"=


*/

    EditDate string `url:"edit_date,omitempty"`
/*
Specifies a date range relative to the current date in which the workspace was edited.

For more information about date range notation, see <a href="#overview--date-range-format">Date range format</a> in the **Key Concepts** section.

*/

    Facets string `url:"facets,omitempty"`
/*
Specifies which facets to return.

Facets are a summary list of the most common occurrences for each facet type listed.
For example, if `facets` specifies custom1, all items containing any value for custom1 are counted.
The five items with the highest counts are returned.
Facets allow users to anticipate selections, or to present a list of suggestions that the user will likely select.

Multiple facets can be included using a comma-separated list.
For example: `facets=custom1,custom2,custom9`.
If this parameter is omitted, which is the default value, no facets are returned.
Possible values are: custom1, custom2, custom3, custom4, custom5, custom6, custom7, custom8, custom9, custom10, custom11, custom12, custom29, and custom30.

This parameter is often used along with the `results` parameter.
If neither of the parameters are included, only the search results are returned.
If `results=false`, and a facets value is specified, only the facets information is returned.
If both the parameters are specified, both facets and search results are returned.
If `results=false` and no valid facets parameters are specified, no data of either type is returned. However, an empty response object is returned.<br>

<font style="color:black;font-size: 15px;font-weight: bold;">Sample request</font><br>
In this example, the search is GET `https://www.ajubalaw.com/work/api/v2/customers/100/recent-workspaces?results=false&facets=custom1,custom9`.
The response JSON object does not include any search results.
The facets section specified returning custom9 but there are no custom9 values to return.
However, three custom1 values are returned, with *Microsoft* being the most common with 132 instances.

<pre>
{
  "data": {
    "facets": {
      "custom1": [
        {
          "alias"      : "Microsoft",
          "count"      : 132,
          "description": "Microsoft Corporation"
        },
        {
          "alias"      : "Amazon",
          "count"      : 20,
          "description": "Amazon.com, Inc"
        },
        {
          "alias"      : "Wallachs",
          "count"      : 1,
          "description": "Wallachs"
        }
      ],
      "custom9": []
    }
  }
}</pre>

*/

    FileTarget string `url:"file_target,omitempty"`
/*
Specifies  text to find in any of custom1, custom2, custom29, custom30, or the workspace name.<br><br>
This parameter is case insensitive.<br>
Matches whole words (partial matches are not supported).<br>
For example, `file_target=litigation Enron` matches either litigation or Enron in any of the fields.<br>
The search with `file_target=litigation` matches only *litigation*.<br>
The search with `file_target=litiga` does not match any variation.<br>
The wildcard character (&ast;) is not supported.
For example, `file_target=*gation`, or `file_target=litiga*` does not match *litigation*.<br>
The underscore character is treated as a space.
The search with `file_target=litigation_Enron` is identical to `file_target=litigation Enron`.

*/

    IncludeIds string `url:"include_ids,omitempty"`
/*
Specifies to return workspaces only from the list of workspace IDs provided.<br><br>
If this parameter is omitted, returns all workspaces from the user's available libraries.
Multiple workspaces can be included using a comma-separated list.
For example, `include_ids=active_us!22,active_us!55,active_us!401`.

*/

    Language string `url:"language,omitempty"`
/*
Specifies the language to be used for searching the content.<br><br>
It is specfied in ISO 639-1 language code and is not case-sensitive.

Documents do not have a language property. However, specifying a language code along with your
search terms can help improve the search accuracy and efficiency.
For example, if iManage Work knows the search term "人身伤害" is Chinese and not English, it performs a better search.

This field supports multiple languages only if the feature flag `multi_lang_document_search` is `true` in your environment. This flag 
is returned by the request <a href="#get-/work/api/v2/customers/-customerId-/features"> GET /customers/{customerId}/features</a>.<br>

Multiple language parameters are specified in a comma-separated string. For example, `"zh,en,de"`.

The following are the allowed values:

| Language   | Code |
|------------|------|
| Chinese    | zh   |
| English    | en   |
| German     | de   |
| French     | fr   |
| Japanese   | ja   |
| Spanish    | es   |
| Portuguese | pt   |

*/

    Limit int `url:"limit,omitempty"`
/*
Specifies the maximum number of items to include in the response.

The request returns the actual number of items up to the limit value (inclusive). If there are more items than the limit value, no more items than the limit are returned and a cursor value is returned.

This parameter can be used in conjunction with pagination parameters for endpoints that support them. For more information, refer to <a href="#overview--pagination">Pagination</a> in the **Key Concepts** section.
*/

    Name string `url:"name,omitempty"`
/*
Specifies the text to search for in the `name` field of the workspaces.<br><br>
This parameter is case insensitive.<br>
Matches whole fields only (partial matches are not supported).<br>
For workspace or document searches, multiple words can be included using a comma-separated list (`name=litigation,Enron`) or with spaces (`name=litigation Enron`).<br>
Exact words may be matched with quotation marks (`name="litigation Enron"`).<br>
The wildcard character (&ast;) is not supported.
For example, `name=*gation`, or `name=litiga*` does not match `litigation`.

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
Specifies the owner's user ID.<br>

The parameter value:
* Is case insensitive.<br>
* Matches whole fields only (partial matches are not supported).<br>
* Supports the wildcard character (&ast;).
The wildcard character matches any number of contiguous characters.
To match partial fields, include the wildcard character before, after, or on both ends of the user ID.<br>
For example, the following all match&#58; `owner=ACASE`, `owner=acase`, `owner=*C*s*`.
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

    SortOrder SortOrder `url:"sort_order,omitempty"`
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



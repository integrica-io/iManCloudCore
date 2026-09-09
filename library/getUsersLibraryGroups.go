package library

import (
	"context"
	"time"

	"github.com/google/go-querystring/query"
	"github.com/integrica-io/iManCloudCore/client"
	"github.com/integrica-io/iManCloudCore/internal"
)

func GetUsersLibraryGroups(ctx context.Context, client *client.Client, libraryId string, userId string, options *GetUsersLibraryGroupsOptions) (GetUsersLibraryGroupsOutput, error) {
	var data GetUsersLibraryGroupsOutput
	endpoint := client.BaseUrl.JoinPath("work", "api", "v2", "customers", client.TokenCfg.CustomerId, "libraries", libraryId, "users", userId, "groups")

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

type GetUsersLibraryGroupsOutput struct {
	Data []struct {
		Database          string    `json:"database"`
		DistinguishedName string    `json:"distinguished_name"`
		Enabled           bool      `json:"enabled"`
		FullName          string    `json:"full_name"`
		GroupDomain       string    `json:"group_domain"`
		GroupNos          int       `json:"group_nos"`
		GroupNumber       int64     `json:"group_number"`
		ID                string    `json:"id"`
		IsExternal        bool      `json:"is_external"`
		LastSyncTs        time.Time `json:"last_sync_ts"`
		SyncID            string    `json:"sync_id"`
	} `json:"data"`
	TotalCount int `json:"total_count"`
}

type GetUsersLibraryGroupsOptions struct {
	Alias string `url:"alias,omitempty"`
	/*
	   Filters groups based on group alias.<br>
	   The group alias is the group `id`.

	   * The wildcard character (&ast;) is supported and matches any number of contiguous characters, including spaces.
	   * If no wildcard character is included, only the group that matches the exact string is returned.
	   * Special characters allowed are *().&-_[]`~\|@$%^?:{}!',/\\#+<>;"=


	*/

	Enabled bool `url:"enabled,omitempty"`
	/*
	   Filters groups based on their enabled status.

	   If `true`, only enabled groups are returned.<br>
	   If `false`, only disabled groups are returned.<br>
	   If omitted, both enabled and disabled groups are returned.

	*/

	External bool `url:"external,omitempty"`
	/*
	   Filters groups based on their `is_external` value.

	   If `true`, only external groups are returned.<br>
	   If `false`, only internal groups are returned.<br>
	   If omitted, both internal and external groups are returned.

	*/

	FullName string `url:"full_name,omitempty"`
	/*
	   Filters groups based on the name.

	   * The wildcard character (&ast;) is supported and matches any number of contiguous characters, including spaces.
	   * To match partial fields, include the asterisk wildcard before or after the query value.
	   * If no wildcard character is included, only the group that matches the exact string is returned.
	   * Special characters allowed are *().&-_[]`~\|@$%^?:{}!',/\\#+<>;"=


	*/

	LastSyncTs string `url:"last_sync_ts,omitempty"`
	/*
	   Filters groups based on  on their last sync time.

	   The last sync time is the date when it was last synchronized with an external directory (such as LDAP or Microsoft
	   Active Directory). The parameters `last_sync_ts`  and `last_sync_ts_end` can be used to form a date/time range.
	   For more information, see the parameter `last_sync_ts_end`.

	   * If omitted, the last sync time is not used for filtering results.
	   * If both parameters are used, items whose last sync time falls within the range will be returned.
	   * If only `last_sync_ts` is provided, then it does not act as the start of a range, but as an exact timestamp filter. Only items whose last sync timestamp matches the specified value are returned.
	   * If only `last_sync_ts_end` is provided, then only items with the last sync time before the specified value are returned.

	*/

	LastSyncTsEnd string `url:"last_sync_ts_end,omitempty"`
	/*
	   Filters groups based on the upper date/time limit for the last sync time.

	   Returns the items that are synced between `last_sync_ts` and `last_sync_ts_end`.<br>
	   The last sync time is the date when it was last synchronized with an external directory (such as LDAP or Microsoft
	   Active Directory). The parameters `last_sync_ts` and `last_sync_ts_end` can be used to form a date/time range.

	   * If omitted, a date/time range is not used, and uses only the `last_sync_ts`, if available.
	   * If both parameters are used, the last sync time that falls between these values inclusively are returned.
	   * If only `last_sync_ts` is provided, then it does not act as the start of a range, but as an exact timestamp filter. Only items whose last sync timestamp matches the specified value are returned.
	   * If only `last_sync_ts_end` is provided, then only items with the last sync time before the specified value are returned.

	*/

	Limit int `url:"limit,omitempty"`
	/*
	   Specifies the maximum number of items to include in the response.

	   The request returns the actual number of items up to the limit value (inclusive). If there are more items than the limit value, no more items than the limit are returned and a cursor value is returned.

	   This parameter can be used in conjunction with pagination parameters for endpoints that support them. For more information, refer to <a href="#overview--pagination">Pagination</a> in the **Key Concepts** section.
	*/

	Nos string `url:"nos,omitempty"`
	/*
	   Filters groups based on the type of group.

	   The following are valid nos values for groups:

	   | Value | Description |
	   | :---: | ----------- |
	   | 2 | Group for virtual users. |
	   | 6 | Group for enterprise users. |

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

	SyncId string `url:"sync_id,omitempty"`
	/*
	   Filters groups based on the ID from an external directory (such as AD FS), which is used for group profile synchronization.
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

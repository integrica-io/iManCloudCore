package library

import (
	"context"

	"github.com/google/go-querystring/query"
	"github.com/integrica-io/iManCloudCore/client"
	"github.com/integrica-io/iManCloudCore/internal"
)

func GetMyMattersSubscriptions(ctx context.Context, client *client.Client, libraryId string, userId string, options *GetMyMattersSubscriptionsOptions) (GetMyMattersSubscriptionsOutput, error) {
	var data GetMyMattersSubscriptionsOutput
	endpoint := client.BaseUrl.JoinPath("work", "api", "v2", "customers", client.TokenCfg.CustomerId, "libraries", libraryId, "users", userId, "my-matters", "subscriptions")

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

type GetMyMattersSubscriptionsOutput struct {
	Data []struct {
		Database       string `json:"database"`
		FullName       string `json:"full_name"`
		SubscriptionID string `json:"subscription_id"`
		Userid         string `json:"userid"`
		Wstype         string `json:"wstype"`
		IsExternal     bool   `json:"is_external"`
		AllowLogon     bool   `json:"allow_logon"`
	} `json:"data"`
	TotalCount int `json:"total_count"`
}

type GetMyMattersSubscriptionsOptions struct {
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

	Total bool `url:"total,omitempty"`
	/*
	   Specifies to include the total count of items found in the response.

	   If `true`, the total count is included in the response.<br>
	   If `false`, the total count is not included in the response.<br>
	   The actual number of items returned may be different because of the *limit* parameter that restricts the number of items
	   returned for any given search.<br>
	*/
}

package customers

import (
	"context"
	"time"
	"github.com/integrica-io/iManCloudCore/internal"

	"github.com/google/go-querystring/query"	
)

type GetUsersGlobalRoleOutput struct {
	Data []struct {
		AllowLogon       bool      `json:"allow_logon"`
		CreateDate       time.Time `json:"create_date"`
		EditDate         time.Time `json:"edit_date"`
		Email            string    `json:"email"`
		FailedLogins     int       `json:"failed_logins"`
		FullName         string    `json:"full_name"`
		ID               string    `json:"id"`
		IPRange          string    `json:"ip_range,omitempty"`
		IPRangeEnabled   bool      `json:"ip_range_enabled"`
		IsExternal       bool      `json:"is_external"`
		Location         string    `json:"location"`
		Ssid             string    `json:"ssid"`
		UserIDEx         string    `json:"user_id_ex"`
		UserNos          int       `json:"user_nos"`
		UserNum          int64     `json:"user_num"`
		PreferredLibrary string    `json:"preferred_library,omitempty"`
	} `json:"data"`
	TotalCount int  `json:"total_count"`
	Overflow   bool `json:"overflow"`
}

type GetUsersGlobalRoleOptions struct{
	AllowLogon bool `url:"allow_logon,omitempty"`
	FullName string `url:"full_name,omitempty"`
	Limit int `url:"limit,omitempty"`
	Location string `url:"location,omitempty"`
	Offset int `url:"offset,omitempty"`
	PreferredLibrary string `url:"preferred_library,omitempty"`
	Query string `url:"query,omitempty"`
	QueryMatchType QueryMatchType `url:"query_match_type,omitempty"`
	Total	bool `url:"total"`
}

func GetUsersGlobalRole(ctx context.Context, client *internal.Client, roleId string, options *GetUsersGlobalRoleOptions)(GetUsersGlobalRoleOutput, error){
	var data GetUsersGlobalRoleOutput
	endpoint := client.BaseUrl.JoinPath("work","api","v2","customers",client.TokenCfg.CustomerId, "roles", roleId, "members")

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
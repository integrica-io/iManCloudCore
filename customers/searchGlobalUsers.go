package customers

import (
	"time"
	"context"
	"iManCloudCore/internal"

	"github.com/google/go-querystring/query"	
)

type SearchGlobalUsersOutput struct {
	Data []struct {
		AllowLogon       bool      `json:"allow_logon"`
		CreateDate       time.Time `json:"create_date"`
		EditDate         time.Time `json:"edit_date"`
		Email            string    `json:"email"`
		FullName         string    `json:"full_name"`
		ID               string    `json:"id"`
		IPRange          string    `json:"ip_range"`
		IPRangeEnabled   bool      `json:"ip_range_enabled"`
		IsExternal       bool      `json:"is_external"`
		Location         string    `json:"location"`
		PreferredLibrary string    `json:"preferred_library"`
		Ssid             string    `json:"ssid"`
		UserIDEx         string    `json:"user_id_ex"`
		UserNos          int       `json:"user_nos"`
		UserNum          int64     `json:"user_num"`
		Libraries        []struct {
			ID        string `json:"id"`
			LibraryID int    `json:"library_id"`
			Preferred bool   `json:"preferred"`
		} `json:"libraries"`
	} `json:"data"`
	TotalCount int `json:"total_count"`
}

type SearchGlobalUsersOptions struct {
    Alias string `url:"alias,omitempty"`
    AllowLogon bool `url:"allow_logon,omitempty"`
    DirectoryId string `url:"directory_id,omitempty"`
    Email string `url:"email,omitempty"`
    FullName string `url:"full_name,omitempty"`
    IncludeLibraries bool `url:"include_libraries,omitempty"`
    IpRangeEnabled bool `url:"ip_range_enabled,omitempty"`
    IsExternal bool `url:"is_external,omitempty"`
    Limit int `url:"limit,omitempty"`
    Location string `url:"location,omitempty"`
    Nos int `url:"nos,omitempty"`
    Offset int `url:"offset,omitempty"`
    PasswordNeverExpire bool `url:"password_never_expire,omitempty"`
    PreferredLibrary string `url:"preferred_library,omitempty"`
    Query string `url:"query,omitempty"`
    Ssid string `url:"ssid,omitempty"`
    Total bool `url:"total,omitempty"`
}

func SearchGlobalUsers(ctx context.Context, client *internal.Client, options *SearchGlobalUsersOptions)(SearchGlobalUsersOutput, error){

	var data SearchGlobalUsersOutput
	endpoint := client.BaseUrl.JoinPath("work","api","v2","customers",client.TokenCfg.CustomerId, "users")

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
package customers

import (
	"time"
	"context"
	"github.com/integrica-io/iManCloudCore/internal"
	"github.com/integrica-io/iManCloudCore/client"
	"github.com/google/go-querystring/query"	
)

type GetInformationGlobalUserOutput struct {
	Data struct {
		AllowLogon       bool      `json:"allow_logon"`
		CreateDate       time.Time `json:"create_date"`
		EditDate         time.Time `json:"edit_date"`
		Email            string    `json:"email"`
		ExchAutodiscover string    `json:"exch_autodiscover"`
		FullName         string    `json:"full_name"`
		ID               string    `json:"id"`
		IPRange          string    `json:"ip_range"`
		IPRangeEnabled   bool      `json:"ip_range_enabled"`
		IsExternal       bool      `json:"is_external"`
		PreferredLibrary string    `json:"preferred_library"`
		Ssid             string    `json:"ssid"`
		UserDomain       string    `json:"user_domain"`
		UserIDEx         string    `json:"user_id_ex"`
		UserNos          int       `json:"user_nos"`
		UserNum          int64     `json:"user_num"`
		Libraries        []struct {
			ID        string `json:"id"`
			LibraryID int    `json:"library_id"`
			Preferred bool   `json:"preferred"`
		} `json:"libraries"`
	} `json:"data"`
}

type GetInformationGlobalUserOptions struct {
    IncludeLibraries bool `url:"include_libraries,omitempty"`
}

func GetInformationGlobalUser(ctx context.Context, client *client.Client, userId string, options *GetInformationGlobalUserOptions)(GetInformationGlobalUserOutput, error){

	var data GetInformationGlobalUserOutput
	endpoint := client.BaseUrl.JoinPath("work","api","v2","customers",client.TokenCfg.CustomerId, "users", userId)

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
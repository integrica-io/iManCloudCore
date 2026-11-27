package customers

import (
	"context"
	"encoding/json"
	"time"

	"github.com/integrica-io/iManCloudCore/client"
	"github.com/integrica-io/iManCloudCore/internal"
)

type UpdateUsersInformationOptions struct {
	Email string `json:"email,omitempty"`
	FullName string `json:"full_name,omitempty"`
	Id string `json:"id,omitempty"`
	IpRange string `json:"ip_range,omitempty"`
	IpRangeEnabled bool `json:"ip_range_enabled,omitempty"`
	IsExternal bool `json:"is_external,omitempty"`
	AllowLogon bool `json:"allow_logon,omitempty"`
	PreferredLibrary string `json:"preferred_library,omitempty"`
	UserDomain string `json:"user_domain,omitempty"`
	UserIdEx string `json:"user_id_ex,omitempty"`
	UserNos int `json:"user_nos,omitempty"`
}

type UpdateUsersInformationResponse struct {
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
	} `json:"data"`
}

func UpdateDocumentProfile(ctx context.Context, client *client.Client, userId string, options UpdateUsersInformationOptions)(UpdateUsersInformationResponse, error){
	var data UpdateUsersInformationResponse
	endpoint := client.BaseUrl.JoinPath("work", "api", "v2", "customers", client.TokenCfg.CustomerId, "users", userId)

	bytes, err := json.Marshal(options)
	if err != nil {
		return data, err
	}
	
	req := internal.HttpRequestBuilder{}
	req.Context(ctx).Url(*endpoint).Method(internal.Patch).ToJson(&data).Body(bytes)

	if err := client.Req(req); err != nil {
		return data, err
	}
	return data, nil
}
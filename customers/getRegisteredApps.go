package customers

import (
	"github.com/integrica-io/iManCloudCore/internal"
	"github.com/integrica-io/iManCloudCore/client"
	"context"
	"time"

	"github.com/google/go-querystring/query"
)

const (
	Available	Status = "available"
	Disabled 	Status = "disabled"
	Enabled		Status = "enabled"
)

type Status string

type GetRegisteredAppsOptions struct {
    HostAppId string `url:"host_app_id,omitempty"`
    Limit int `url:"limit,omitempty"`
    Name string `url:"name,omitempty"`
    Offset int `url:"offset,omitempty"`
    Status Status `url:"status,omitempty"`
    Total bool `url:"total,omitempty"`
}

type GetRegisteredAppsOutput struct {
	Data []struct {
		ID    string `json:"id"`
		Oauth struct {
			ClientIDIssuedAt      int      `json:"client_id_issued_at"`
			ApplicationType       string   `json:"application_type"`
			ClientSecret          string   `json:"client_secret"`
			ClientSecretExpiresAt int      `json:"client_secret_expires_at"`
			ClientType            string   `json:"client_type"`
			Contacts              []string `json:"contacts"`
			RedirectUris          []string `json:"redirect_uris"`
		} `json:"oauth"`
		EditDate time.Time `json:"edit_date"`
		Links    struct {
			Package struct {
				Href string `json:"href"`
			} `json:"package"`
		} `json:"_links"`
		Status  string `json:"status"`
		HostApp struct {
			ID string `json:"id"`
		} `json:"host_app"`
		DefaultSecurity string `json:"default_security"`
		Type            string `json:"type"`
		Name            string `json:"name"`
		Description     string `json:"description"`
		Publisher       struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"publisher"`
	} `json:"data"`
	TotalCount int `json:"total_count"`
}

func GetRegisteredApps(ctx context.Context, client *client.Client, options *GetRegisteredAppsOptions)(GetRegisteredAppsOutput, error){
	var data GetRegisteredAppsOutput
	endpoint := client.BaseUrl.JoinPath("work","api","v2","customers",client.TokenCfg.CustomerId, "apps")

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
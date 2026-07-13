package library

import (
	"context"

	"github.com/integrica-io/iManCloudCore/client"
	"github.com/integrica-io/iManCloudCore/internal"
)

func GetFolderSecurity(ctx context.Context, client *client.Client, libraryId string, folderId string) (GetFolderSecurityOutput, error) {
	var data GetFolderSecurityOutput
	endpoint := client.BaseUrl.JoinPath("work", "api", "v2", "customers", client.TokenCfg.CustomerId, "libraries", libraryId, "folders", folderId, "security")

	req := internal.HttpRequestBuilder{}
	req.Context(ctx).Url(*endpoint).Method(internal.Get).ToJson(&data)

	if err := client.Req(req); err != nil {
		return data, err
	}
	return data, nil
}

type GetFolderSecurityOutput struct {
	Data []struct {
		Access              int    `json:"access"`
		AccessLevel         string `json:"access_level"`
		ID                  string `json:"id"`
		Name                string `json:"name"`
		Sid                 string `json:"sid"`
		Type                string `json:"type"`
		IsExternal          bool   `json:"is_external"`
		AllowLogon          bool   `json:"allow_logon,omitempty"`
		Enabled             bool   `json:"enabled,omitempty"`
		HasRestrictedMember bool   `json:"has_restricted_member,omitempty"`
	} `json:"data"`
	DefaultSecurity string `json:"default_security"`
}

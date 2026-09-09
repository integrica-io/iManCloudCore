package library

import (
	"context"
	"time"

	"github.com/integrica-io/iManCloudCore/client"
	"github.com/integrica-io/iManCloudCore/internal"
)

func GetGroupsDetails(ctx context.Context, client *client.Client, libraryId string, groupId string) (GetGroupDetailsOutput, error) {
	var data GetGroupDetailsOutput
	endpoint := client.BaseUrl.JoinPath("work", "api", "v2", "customers", client.TokenCfg.CustomerId, "libraries", libraryId, "groups", groupId)

	req := internal.HttpRequestBuilder{}
	req.Context(ctx).Url(*endpoint).Method(internal.Get).ToJson(&data)

	if err := client.Req(req); err != nil {
		return data, err
	}
	return data, nil
}

type GetGroupDetailsOutput struct {
	Data struct {
		Database          string    `json:"database"`
		DistinguishedName string    `json:"distinguished_name"`
		Enabled           bool      `json:"enabled"`
		FullName          string    `json:"full_name"`
		GlobalID          int64     `json:"global_id"`
		GroupDomain       string    `json:"group_domain"`
		GroupNos          int       `json:"group_nos"`
		GroupNumber       int64     `json:"group_number"`
		ID                string    `json:"id"`
		IsExternal        bool      `json:"is_external"`
		LastSyncTs        time.Time `json:"last_sync_ts"`
		SyncID            string    `json:"sync_id"`
	} `json:"data"`
}

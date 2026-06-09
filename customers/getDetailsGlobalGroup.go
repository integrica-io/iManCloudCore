package customers

import (
	"time"
	"context"
	"github.com/integrica-io/iManCloudCore/internal"
	"github.com/integrica-io/iManCloudCore/client"
)

type GetDetailsGlobalGroupOutput struct {
	Data struct {
		CreateDate        time.Time `json:"create_date"`
		DirectoryID       string    `json:"directory_id"`
		DistinguishedName string    `json:"distinguished_name"`
		EditDate          time.Time `json:"edit_date"`
		Enabled           bool      `json:"enabled"`
		FullName          string    `json:"full_name"`
		GlobalID          int64     `json:"global_id"`
		GroupDomain       string    `json:"group_domain"`
		GroupNos          int       `json:"group_nos"`
		GroupNumber       int64     `json:"group_number"`
		ID                string    `json:"id"`
		IsExternal        bool      `json:"is_external"`
		LastSyncTs        time.Time `json:"last_sync_ts"`
	} `json:"data"`
}

func GetDetailsGlobalGroup(ctx context.Context, client *client.Client, groupId string)(GetDetailsGlobalGroupOutput, error){

	var data GetDetailsGlobalGroupOutput
	endpoint := client.BaseUrl.JoinPath("work","api","v2","customers",client.TokenCfg.CustomerId, "groups", groupId)

	req := internal.HttpRequestBuilder{}
	req.Context(ctx).Url(*endpoint).Method(internal.Get).ToJson(&data)

	if err := client.Req(req); err != nil {
		return data, err
	}
	return data, nil
}
package library

import (
	"context"
	"fmt"

	"github.com/integrica-io/iManCloudCore/client"
	"github.com/integrica-io/iManCloudCore/internal"
)

func GetUserAccessFolder(ctx context.Context, client *client.Client, libraryId string, folderId string, userId string) (GetUserAccessFolderOutput, error) {
	var data GetUserAccessFolderOutput
	endpoint := client.BaseUrl.JoinPath("work", "api", "v2", "customers", client.TokenCfg.CustomerId, "libraries", libraryId, "folders", fmt.Sprintf("%s!%s", libraryId, folderId), "users", userId, "security")

	req := internal.HttpRequestBuilder{}
	req.Context(ctx).Url(*endpoint).Method(internal.Get).ToJson(&data)

	if err := client.Req(req); err != nil {
		return data, err
	}
	return data, nil
}

type GetUserAccessFolderOutput struct {
	Data struct {
		Access string `json:"access"`
	} `json:"data"`
}

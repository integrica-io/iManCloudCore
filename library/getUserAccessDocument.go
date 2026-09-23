package library

import (
	"context"
	"fmt"

	"github.com/integrica-io/iManCloudCore/client"
	"github.com/integrica-io/iManCloudCore/internal"
)

func GetUserAccessDocument(ctx context.Context, client *client.Client, libraryId string, docId string, userId string) (GetUserAccessDocumentOutput, error) {
	var data GetUserAccessDocumentOutput
	endpoint := client.BaseUrl.JoinPath("work", "api", "v2", "customers", client.TokenCfg.CustomerId, "libraries", libraryId, "documents", fmt.Sprintf("%s!%s", libraryId, docId), "users", userId, "security")

	req := internal.HttpRequestBuilder{}
	req.Context(ctx).Url(*endpoint).Method(internal.Get).ToJson(&data)

	if err := client.Req(req); err != nil {
		return data, err
	}
	return data, nil
}

type GetUserAccessDocumentOutput struct {
	Data struct {
		Access string `json:"access"`
	} `json:"data"`
}

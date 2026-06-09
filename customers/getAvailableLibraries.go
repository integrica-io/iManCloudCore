package customers

import (
	"context"
	"github.com/integrica-io/iManCloudCore/internal"
	"github.com/integrica-io/iManCloudCore/client"
)

type GetAvailableLibrariesOutput struct {
	Data []struct {
		ID       string `json:"id"`
		Type     string `json:"type"`
		IsHidden bool   `json:"is_hidden"`
	} `json:"data"`
}

func GetAvailableLibraries(ctx context.Context, client *client.Client)(GetAvailableLibrariesOutput, error){
	endpoint := client.BaseUrl.JoinPath("work", "api", "v2","customers",client.TokenCfg.CustomerId, "libraries")
	var data GetAvailableLibrariesOutput

	req := internal.HttpRequestBuilder{}
	req.Context(ctx).Url(*endpoint).Method(internal.Get).ToJson(&data)

	if err := client.Req(req); err != nil {
		return data, err
	}
	return data, nil
}
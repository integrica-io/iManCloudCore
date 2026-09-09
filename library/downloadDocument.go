package library

import (
	"context"
	
	"github.com/google/go-querystring/query"
	"github.com/integrica-io/iManCloudCore/client"
	"github.com/integrica-io/iManCloudCore/internal"
)

func DownloadDocument(ctx context.Context, client *client.Client, libraryId string, docId string, options GetFolderProfileOptions, exportFilename string) (error) {
	endpoint := client.BaseUrl.JoinPath("work", "api", "v2", "customers", client.TokenCfg.CustomerId, "libraries", libraryId, "documents", docId, "download")

	values, err := query.Values(options)
	if err != nil {
		return err
	}
	endpoint.RawQuery = values.Encode()

	req := internal.HttpRequestBuilder{}
	req.Context(ctx).Url(*endpoint).Method(internal.Get).ToFile(exportFilename)

	if err := client.Req(req); err != nil {
		return err
	}
	return nil
}

type DownloadDocumentOptions struct {
	Activity string `url:"activity"`//Default View
	/*
	Specifies the activity type for the document download action.

	For the document download action, the parameter activity must be specified as export, so that the activity is tracked correctly in the document 
	history timeline. If it is not specified, the default activity type appears as view in the document history.
	*/
	Latest bool `url:"latest,omitempty"` //Default false
	/* 	
	Specifies to download the latest version of a document.
	
	If true, downloads the latest version irrespective of the version number specified in the docId.
	If false, downloads the version specified in the endpoint parameter docId.
	*/
}

// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. 
// =================================================================================

package sdk

import (
	"context"

	"github.com/gogf/gf/contrib/sdk/httpclient/v2"
	"github.com/gogf/gf/v2/text/gstr"

	"github.com/patient-fyd/jxust-softhub-api/api/file"
	"github.com/patient-fyd/jxust-softhub-api/api/file/v1"
)

type implementerFileV1 struct {
	*httpclient.Client
}

func (i *implementer) FileV1() file.IFileV1 {
	var (
		client = httpclient.New(i.config)
		prefix = gstr.TrimRight(i.config.URL, "/") + ""
	)
	client.Client = client.Prefix(prefix)
	return &implementerFileV1{client}
}

func (i *implementerFileV1) Upload(ctx context.Context, req *v1.UploadReq) (res *v1.UploadRes, err error) {
	err = i.Request(ctx, req, &res)
	return
}

func (i *implementerFileV1) List(ctx context.Context, req *v1.ListReq) (res *v1.ListRes, err error) {
	err = i.Request(ctx, req, &res)
	return
}

func (i *implementerFileV1) Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error) {
	err = i.Request(ctx, req, &res)
	return
}

func (i *implementerFileV1) Download(ctx context.Context, req *v1.DownloadReq) (res *v1.DownloadRes, err error) {
	err = i.Request(ctx, req, &res)
	return
}


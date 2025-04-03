// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. 
// =================================================================================

package sdk

import (
	"context"

	"github.com/gogf/gf/contrib/sdk/httpclient/v2"
	"github.com/gogf/gf/v2/text/gstr"

	"github.com/patient-fyd/jxust-softhub-api/api/news"
	"github.com/patient-fyd/jxust-softhub-api/api/news/v1"
)

type implementerNewsV1 struct {
	*httpclient.Client
}

func (i *implementer) NewsV1() news.INewsV1 {
	var (
		client = httpclient.New(i.config)
		prefix = gstr.TrimRight(i.config.URL, "/") + ""
	)
	client.Client = client.Prefix(prefix)
	return &implementerNewsV1{client}
}

func (i *implementerNewsV1) List(ctx context.Context, req *v1.ListReq) (res *v1.ListRes, err error) {
	err = i.Request(ctx, req, &res)
	return
}

func (i *implementerNewsV1) Detail(ctx context.Context, req *v1.DetailReq) (res *v1.DetailRes, err error) {
	err = i.Request(ctx, req, &res)
	return
}

func (i *implementerNewsV1) Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error) {
	err = i.Request(ctx, req, &res)
	return
}

func (i *implementerNewsV1) Update(ctx context.Context, req *v1.UpdateReq) (res *v1.UpdateRes, err error) {
	err = i.Request(ctx, req, &res)
	return
}

func (i *implementerNewsV1) Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error) {
	err = i.Request(ctx, req, &res)
	return
}


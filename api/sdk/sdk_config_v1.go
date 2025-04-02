// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. 
// =================================================================================

package sdk

import (
	"context"

	"github.com/gogf/gf/contrib/sdk/httpclient/v2"
	"github.com/gogf/gf/v2/text/gstr"

	"github.com/patient-fyd/jxust-softhub-api/api/config"
	"github.com/patient-fyd/jxust-softhub-api/api/config/v1"
)

type implementerConfigV1 struct {
	*httpclient.Client
}

func (i *implementer) ConfigV1() config.IConfigV1 {
	var (
		client = httpclient.New(i.config)
		prefix = gstr.TrimRight(i.config.URL, "/") + ""
	)
	client.Client = client.Prefix(prefix)
	return &implementerConfigV1{client}
}

func (i *implementerConfigV1) Get(ctx context.Context, req *v1.GetReq) (res *v1.GetRes, err error) {
	err = i.Request(ctx, req, &res)
	return
}

func (i *implementerConfigV1) List(ctx context.Context, req *v1.ListReq) (res *v1.ListRes, err error) {
	err = i.Request(ctx, req, &res)
	return
}

func (i *implementerConfigV1) Set(ctx context.Context, req *v1.SetReq) (res *v1.SetRes, err error) {
	err = i.Request(ctx, req, &res)
	return
}

func (i *implementerConfigV1) Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error) {
	err = i.Request(ctx, req, &res)
	return
}


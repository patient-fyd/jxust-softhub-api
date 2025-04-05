// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. 
// =================================================================================

package sdk

import (
	"context"

	"github.com/gogf/gf/contrib/sdk/httpclient/v2"
	"github.com/gogf/gf/v2/text/gstr"

	"github.com/patient-fyd/jxust-softhub-api/api/circle"
	"github.com/patient-fyd/jxust-softhub-api/api/circle/v1"
)

type implementerCircleV1 struct {
	*httpclient.Client
}

func (i *implementer) CircleV1() circle.ICircleV1 {
	var (
		client = httpclient.New(i.config)
		prefix = gstr.TrimRight(i.config.URL, "/") + ""
	)
	client.Client = client.Prefix(prefix)
	return &implementerCircleV1{client}
}

func (i *implementerCircleV1) List(ctx context.Context, req *v1.ListReq) (res *v1.ListRes, err error) {
	err = i.Request(ctx, req, &res)
	return
}

func (i *implementerCircleV1) Detail(ctx context.Context, req *v1.DetailReq) (res *v1.DetailRes, err error) {
	err = i.Request(ctx, req, &res)
	return
}


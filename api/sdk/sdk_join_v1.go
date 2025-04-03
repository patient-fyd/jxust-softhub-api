// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. 
// =================================================================================

package sdk

import (
	"context"

	"github.com/gogf/gf/contrib/sdk/httpclient/v2"
	"github.com/gogf/gf/v2/text/gstr"

	"github.com/patient-fyd/jxust-softhub-api/api/join"
	"github.com/patient-fyd/jxust-softhub-api/api/join/v1"
)

type implementerJoinV1 struct {
	*httpclient.Client
}

func (i *implementer) JoinV1() join.IJoinV1 {
	var (
		client = httpclient.New(i.config)
		prefix = gstr.TrimRight(i.config.URL, "/") + ""
	)
	client.Client = client.Prefix(prefix)
	return &implementerJoinV1{client}
}

func (i *implementerJoinV1) Apply(ctx context.Context, req *v1.ApplyReq) (res *v1.ApplyRes, err error) {
	err = i.Request(ctx, req, &res)
	return
}

func (i *implementerJoinV1) List(ctx context.Context, req *v1.ListReq) (res *v1.ListRes, err error) {
	err = i.Request(ctx, req, &res)
	return
}

func (i *implementerJoinV1) Detail(ctx context.Context, req *v1.DetailReq) (res *v1.DetailRes, err error) {
	err = i.Request(ctx, req, &res)
	return
}

func (i *implementerJoinV1) Review(ctx context.Context, req *v1.ReviewReq) (res *v1.ReviewRes, err error) {
	err = i.Request(ctx, req, &res)
	return
}

func (i *implementerJoinV1) ApplicationList(ctx context.Context, req *v1.ApplicationListReq) (res *v1.ApplicationListRes, err error) {
	err = i.Request(ctx, req, &res)
	return
}


// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. 
// =================================================================================

package sdk

import (
	"context"

	"github.com/gogf/gf/contrib/sdk/httpclient/v2"
	"github.com/gogf/gf/v2/text/gstr"

	"github.com/patient-fyd/jxust-softhub-api/api/stat"
	"github.com/patient-fyd/jxust-softhub-api/api/stat/v1"
)

type implementerStatV1 struct {
	*httpclient.Client
}

func (i *implementer) StatV1() stat.IStatV1 {
	var (
		client = httpclient.New(i.config)
		prefix = gstr.TrimRight(i.config.URL, "/") + ""
	)
	client.Client = client.Prefix(prefix)
	return &implementerStatV1{client}
}

func (i *implementerStatV1) VisitStat(ctx context.Context, req *v1.VisitStatReq) (res *v1.VisitStatRes, err error) {
	err = i.Request(ctx, req, &res)
	return
}

func (i *implementerStatV1) ActivityStat(ctx context.Context, req *v1.ActivityStatReq) (res *v1.ActivityStatRes, err error) {
	err = i.Request(ctx, req, &res)
	return
}

func (i *implementerStatV1) UserStat(ctx context.Context, req *v1.UserStatReq) (res *v1.UserStatRes, err error) {
	err = i.Request(ctx, req, &res)
	return
}


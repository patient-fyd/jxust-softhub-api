// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. 
// =================================================================================

package sdk

import (
	"context"

	"github.com/gogf/gf/contrib/sdk/httpclient/v2"
	"github.com/gogf/gf/v2/text/gstr"

	"github.com/patient-fyd/jxust-softhub-api/api/user"
	"github.com/patient-fyd/jxust-softhub-api/api/user/v1"
)

type implementerUserV1 struct {
	*httpclient.Client
}

func (i *implementer) UserV1() user.IUserV1 {
	var (
		client = httpclient.New(i.config)
		prefix = gstr.TrimRight(i.config.URL, "/") + ""
	)
	client.Client = client.Prefix(prefix)
	return &implementerUserV1{client}
}

func (i *implementerUserV1) Info(ctx context.Context, req *v1.InfoReq) (res *v1.InfoRes, err error) {
	err = i.Request(ctx, req, &res)
	return
}

func (i *implementerUserV1) List(ctx context.Context, req *v1.ListReq) (res *v1.ListRes, err error) {
	err = i.Request(ctx, req, &res)
	return
}

func (i *implementerUserV1) AssignRole(ctx context.Context, req *v1.AssignRoleReq) (res *v1.AssignRoleRes, err error) {
	err = i.Request(ctx, req, &res)
	return
}


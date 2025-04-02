// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. 
// =================================================================================

package sdk

import (
	"context"

	"github.com/gogf/gf/contrib/sdk/httpclient/v2"
	"github.com/gogf/gf/v2/text/gstr"

	"github.com/patient-fyd/jxust-softhub-api/api/auth"
	"github.com/patient-fyd/jxust-softhub-api/api/auth/v1"
)

type implementerAuthV1 struct {
	*httpclient.Client
}

func (i *implementer) AuthV1() auth.IAuthV1 {
	var (
		client = httpclient.New(i.config)
		prefix = gstr.TrimRight(i.config.URL, "/") + ""
	)
	client.Client = client.Prefix(prefix)
	return &implementerAuthV1{client}
}

func (i *implementerAuthV1) Register(ctx context.Context, req *v1.RegisterReq) (res *v1.RegisterRes, err error) {
	err = i.Request(ctx, req, &res)
	return
}

func (i *implementerAuthV1) Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error) {
	err = i.Request(ctx, req, &res)
	return
}

func (i *implementerAuthV1) RefreshToken(ctx context.Context, req *v1.RefreshTokenReq) (res *v1.RefreshTokenRes, err error) {
	err = i.Request(ctx, req, &res)
	return
}


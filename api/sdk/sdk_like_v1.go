// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. 
// =================================================================================

package sdk

import (
	"context"

	"github.com/gogf/gf/contrib/sdk/httpclient/v2"
	"github.com/gogf/gf/v2/text/gstr"

	"github.com/patient-fyd/jxust-softhub-api/api/like"
	"github.com/patient-fyd/jxust-softhub-api/api/like/v1"
)

type implementerLikeV1 struct {
	*httpclient.Client
}

func (i *implementer) LikeV1() like.ILikeV1 {
	var (
		client = httpclient.New(i.config)
		prefix = gstr.TrimRight(i.config.URL, "/") + ""
	)
	client.Client = client.Prefix(prefix)
	return &implementerLikeV1{client}
}

func (i *implementerLikeV1) LikeToggle(ctx context.Context, req *v1.LikeToggleReq) (res *v1.LikeToggleRes, err error) {
	err = i.Request(ctx, req, &res)
	return
}


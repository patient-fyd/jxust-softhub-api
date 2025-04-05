// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. 
// =================================================================================

package sdk

import (
	"context"

	"github.com/gogf/gf/contrib/sdk/httpclient/v2"
	"github.com/gogf/gf/v2/text/gstr"

	"github.com/patient-fyd/jxust-softhub-api/api/follow"
	"github.com/patient-fyd/jxust-softhub-api/api/follow/v1"
)

type implementerFollowV1 struct {
	*httpclient.Client
}

func (i *implementer) FollowV1() follow.IFollowV1 {
	var (
		client = httpclient.New(i.config)
		prefix = gstr.TrimRight(i.config.URL, "/") + ""
	)
	client.Client = client.Prefix(prefix)
	return &implementerFollowV1{client}
}

func (i *implementerFollowV1) FollowToggle(ctx context.Context, req *v1.FollowToggleReq) (res *v1.FollowToggleRes, err error) {
	err = i.Request(ctx, req, &res)
	return
}

func (i *implementerFollowV1) FollowingList(ctx context.Context, req *v1.FollowingListReq) (res *v1.FollowingListRes, err error) {
	err = i.Request(ctx, req, &res)
	return
}

func (i *implementerFollowV1) FollowerList(ctx context.Context, req *v1.FollowerListReq) (res *v1.FollowerListRes, err error) {
	err = i.Request(ctx, req, &res)
	return
}


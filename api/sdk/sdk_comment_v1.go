// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. 
// =================================================================================

package sdk

import (
	"context"

	"github.com/gogf/gf/contrib/sdk/httpclient/v2"
	"github.com/gogf/gf/v2/text/gstr"

	"github.com/patient-fyd/jxust-softhub-api/api/comment"
	"github.com/patient-fyd/jxust-softhub-api/api/comment/v1"
)

type implementerCommentV1 struct {
	*httpclient.Client
}

func (i *implementer) CommentV1() comment.ICommentV1 {
	var (
		client = httpclient.New(i.config)
		prefix = gstr.TrimRight(i.config.URL, "/") + ""
	)
	client.Client = client.Prefix(prefix)
	return &implementerCommentV1{client}
}

func (i *implementerCommentV1) CommentList(ctx context.Context, req *v1.CommentListReq) (res *v1.CommentListRes, err error) {
	err = i.Request(ctx, req, &res)
	return
}

func (i *implementerCommentV1) CommentCreate(ctx context.Context, req *v1.CommentCreateReq) (res *v1.CommentCreateRes, err error) {
	err = i.Request(ctx, req, &res)
	return
}

func (i *implementerCommentV1) CommentDelete(ctx context.Context, req *v1.CommentDeleteReq) (res *v1.CommentDeleteRes, err error) {
	err = i.Request(ctx, req, &res)
	return
}


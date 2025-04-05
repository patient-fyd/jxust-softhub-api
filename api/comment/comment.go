// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. 
// =================================================================================

package comment

import (
	"context"
	
	"github.com/patient-fyd/jxust-softhub-api/api/comment/v1"
)

type ICommentV1 interface {
	CommentList(ctx context.Context, req *v1.CommentListReq) (res *v1.CommentListRes, err error)
	CommentCreate(ctx context.Context, req *v1.CommentCreateReq) (res *v1.CommentCreateRes, err error)
	CommentDelete(ctx context.Context, req *v1.CommentDeleteReq) (res *v1.CommentDeleteRes, err error)
}



// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. 
// =================================================================================

package comment

import (
	"context"
	
	"github.com/patient-fyd/jxust-softhub-api/api/comment/v1"
)

type ICommentV1 interface {
	// 获取评论列表
	List(ctx context.Context, req *v1.ListReq) (res *v1.ListRes, err error)
	// 创建评论
	Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error)
	// 删除评论
	Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error)
} 
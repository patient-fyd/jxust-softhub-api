// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. 
// =================================================================================

package post

import (
	"context"
	
	"github.com/patient-fyd/jxust-softhub-api/api/post/v1"
)

type IPostV1 interface {
	// 获取帖子列表
	List(ctx context.Context, req *v1.ListReq) (res *v1.ListRes, err error)
	// 获取帖子详情
	Detail(ctx context.Context, req *v1.DetailReq) (res *v1.DetailRes, err error)
	// 创建帖子
	Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error)
	// 删除帖子
	Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error)
} 
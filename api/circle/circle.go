// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. 
// =================================================================================

package circle

import (
	"context"
	
	"github.com/patient-fyd/jxust-softhub-api/api/circle/v1"
)

type ICircleV1 interface {
	// 获取圈子列表
	List(ctx context.Context, req *v1.ListReq) (res *v1.ListRes, err error)
	// 获取圈子详情
	Detail(ctx context.Context, req *v1.DetailReq) (res *v1.DetailRes, err error)
} 
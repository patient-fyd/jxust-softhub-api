// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. 
// =================================================================================

package topic

import (
	"context"
	
	"github.com/patient-fyd/jxust-softhub-api/api/topic/v1"
)

type ITopicV1 interface {
	// 获取话题列表
	List(ctx context.Context, req *v1.ListReq) (res *v1.ListRes, err error)
} 
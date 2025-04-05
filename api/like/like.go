// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. 
// =================================================================================

package like

import (
	"context"
	
	"github.com/patient-fyd/jxust-softhub-api/api/like/v1"
)

type ILikeV1 interface {
	// 点赞/取消点赞
	Toggle(ctx context.Context, req *v1.ToggleReq) (res *v1.ToggleRes, err error)
} 
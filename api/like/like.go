// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. 
// =================================================================================

package like

import (
	"context"
	
	"github.com/patient-fyd/jxust-softhub-api/api/like/v1"
)

type ILikeV1 interface {
	LikeToggle(ctx context.Context, req *v1.LikeToggleReq) (res *v1.LikeToggleRes, err error)
}



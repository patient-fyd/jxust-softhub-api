// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. 
// =================================================================================

package follow

import (
	"context"
	
	"github.com/patient-fyd/jxust-softhub-api/api/follow/v1"
)

type IFollowV1 interface {
	Toggle(ctx context.Context, req *v1.ToggleReq) (res *v1.ToggleRes, err error)
	FollowingList(ctx context.Context, req *v1.FollowingListReq) (res *v1.FollowingListRes, err error)
	FollowerList(ctx context.Context, req *v1.FollowerListReq) (res *v1.FollowerListRes, err error)
}



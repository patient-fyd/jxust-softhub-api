// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. 
// =================================================================================

package follow

import (
	"context"
	
	"github.com/patient-fyd/jxust-softhub-api/api/follow/v1"
)

type IFollowV1 interface {
	FollowToggle(ctx context.Context, req *v1.FollowToggleReq) (res *v1.FollowToggleRes, err error)
	FollowingList(ctx context.Context, req *v1.FollowingListReq) (res *v1.FollowingListRes, err error)
	FollowerList(ctx context.Context, req *v1.FollowerListReq) (res *v1.FollowerListRes, err error)
}



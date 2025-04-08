// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. 
// =================================================================================

package circle

import (
	"context"
	
	"github.com/patient-fyd/jxust-softhub-api/api/circle/v1"
)

type ICircleV1 interface {
	List(ctx context.Context, req *v1.ListReq) (res *v1.ListRes, err error)
	Detail(ctx context.Context, req *v1.DetailReq) (res *v1.DetailRes, err error)
	Join(ctx context.Context, req *v1.JoinReq) (res *v1.JoinRes, err error)
	MyCircles(ctx context.Context, req *v1.MyCirclesReq) (res *v1.MyCirclesRes, err error)
	CircleStat(ctx context.Context, req *v1.CircleStatReq) (res *v1.CircleStatRes, err error)
}



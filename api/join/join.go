// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. 
// =================================================================================

package join

import (
	"context"
	
	"github.com/patient-fyd/jxust-softhub-api/api/join/v1"
)

type IJoinV1 interface {
	Apply(ctx context.Context, req *v1.ApplyReq) (res *v1.ApplyRes, err error)
	List(ctx context.Context, req *v1.ListReq) (res *v1.ListRes, err error)
	Detail(ctx context.Context, req *v1.DetailReq) (res *v1.DetailRes, err error)
	Review(ctx context.Context, req *v1.ReviewReq) (res *v1.ReviewRes, err error)
	ApplicationList(ctx context.Context, req *v1.ApplicationListReq) (res *v1.ApplicationListRes, err error)
}



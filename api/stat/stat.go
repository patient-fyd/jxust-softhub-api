// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. 
// =================================================================================

package stat

import (
	"context"
	
	"github.com/patient-fyd/jxust-softhub-api/api/stat/v1"
)

type IStatV1 interface {
	VisitStat(ctx context.Context, req *v1.VisitStatReq) (res *v1.VisitStatRes, err error)
	ActivityStat(ctx context.Context, req *v1.ActivityStatReq) (res *v1.ActivityStatRes, err error)
	UserStat(ctx context.Context, req *v1.UserStatReq) (res *v1.UserStatRes, err error)
}



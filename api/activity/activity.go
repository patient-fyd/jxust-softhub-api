// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. 
// =================================================================================

package activity

import (
	"context"
	
	"github.com/patient-fyd/jxust-softhub-api/api/activity/v1"
)

type IActivityV1 interface {
	// 活动管理接口
	Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error)
	Update(ctx context.Context, req *v1.UpdateReq) (res *v1.UpdateRes, err error)
	Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error)
	List(ctx context.Context, req *v1.ListReq) (res *v1.ListRes, err error)
	Detail(ctx context.Context, req *v1.DetailReq) (res *v1.DetailRes, err error)
	
	// 活动报名接口
	Register(ctx context.Context, req *v1.RegisterReq) (res *v1.RegisterRes, err error)
	RegisterList(ctx context.Context, req *v1.RegisterListReq) (res *v1.RegisterListRes, err error)
	ApproveRegistration(ctx context.Context, req *v1.ApproveRegistrationReq) (res *v1.ApproveRegistrationRes, err error)
	RejectRegistration(ctx context.Context, req *v1.RejectRegistrationReq) (res *v1.RejectRegistrationRes, err error)
} 
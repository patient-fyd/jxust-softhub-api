// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. 
// =================================================================================

package user

import (
	"context"
	
	"github.com/patient-fyd/jxust-softhub-api/api/user/v1"
)

type IUserV1 interface {
	Info(ctx context.Context, req *v1.InfoReq) (res *v1.InfoRes, err error)
	List(ctx context.Context, req *v1.ListReq) (res *v1.ListRes, err error)
	AssignRole(ctx context.Context, req *v1.AssignRoleReq) (res *v1.AssignRoleRes, err error)
}



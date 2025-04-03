// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. 
// =================================================================================

package config

import (
	"context"
	
	"github.com/patient-fyd/jxust-softhub-api/api/config/v1"
)

type IConfigV1 interface {
	Get(ctx context.Context, req *v1.GetReq) (res *v1.GetRes, err error)
	List(ctx context.Context, req *v1.ListReq) (res *v1.ListRes, err error)
	Set(ctx context.Context, req *v1.SetReq) (res *v1.SetRes, err error)
	Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error)
}



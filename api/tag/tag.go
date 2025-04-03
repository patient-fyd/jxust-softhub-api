// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. 
// =================================================================================

package tag

import (
	"context"
	
	"github.com/patient-fyd/jxust-softhub-api/api/tag/v1"
)

type ITagV1 interface {
	List(ctx context.Context, req *v1.ListReq) (res *v1.ListRes, err error)
	Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error)
	Update(ctx context.Context, req *v1.UpdateReq) (res *v1.UpdateRes, err error)
	Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error)
	ContentTags(ctx context.Context, req *v1.ContentTagsReq) (res *v1.ContentTagsRes, err error)
	SetContentTags(ctx context.Context, req *v1.SetContentTagsReq) (res *v1.SetContentTagsRes, err error)
}



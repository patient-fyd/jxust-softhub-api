// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. 
// =================================================================================

package file

import (
	"context"
	
	"github.com/patient-fyd/jxust-softhub-api/api/file/v1"
)

type IFileV1 interface {
	Upload(ctx context.Context, req *v1.UploadReq) (res *v1.UploadRes, err error)
	List(ctx context.Context, req *v1.ListReq) (res *v1.ListRes, err error)
	Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error)
	Download(ctx context.Context, req *v1.DownloadReq) (res *v1.DownloadRes, err error)
}



package file

import (
	"context"

	"github.com/gogf/gf/v2/net/ghttp"
	v1 "github.com/patient-fyd/jxust-softhub-api/api/file/v1"
)

// IFileV1 文件管理服务V1接口定义
type IFileV1 interface {
	// Upload 上传文件
	Upload(ctx context.Context, req *v1.UploadReq) (res *v1.UploadRes, err error)

	// List 获取文件列表
	List(ctx context.Context, req *v1.ListReq) (res *v1.ListRes, err error)

	// Delete 删除文件
	Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error)

	// Download 下载文件
	Download(ctx context.Context, req *v1.DownloadReq) (res interface{}, err error)
}

// Controller 控制器接口
type Controller interface {
	// Register 注册路由
	Register(group *ghttp.RouterGroup)

	// V1 获取 V1 接口
	V1() IFileV1
}

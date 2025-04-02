package tag

import (
	"context"

	"github.com/gogf/gf/v2/net/ghttp"
	v1 "github.com/patient-fyd/jxust-softhub-api/api/tag/v1"
)

// ITagV1 标签管理服务V1接口定义
type ITagV1 interface {
	// List 获取标签列表
	List(ctx context.Context, req *v1.ListReq) (res *v1.ListRes, err error)

	// Create 创建标签
	Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error)

	// Update 更新标签
	Update(ctx context.Context, req *v1.UpdateReq) (res *v1.UpdateRes, err error)

	// Delete 删除标签
	Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error)

	// ContentTags 获取内容关联的标签
	ContentTags(ctx context.Context, req *v1.ContentTagsReq) (res *v1.ContentTagsRes, err error)

	// SetContentTags 设置内容关联的标签
	SetContentTags(ctx context.Context, req *v1.SetContentTagsReq) (res *v1.SetContentTagsRes, err error)
}

// Controller 控制器接口
type Controller interface {
	// Register 注册路由
	Register(group *ghttp.RouterGroup)

	// V1 获取 V1 接口
	V1() ITagV1
}

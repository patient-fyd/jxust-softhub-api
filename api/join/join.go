package join

import (
	"context"

	"github.com/gogf/gf/v2/net/ghttp"
	v1 "github.com/patient-fyd/jxust-softhub-api/api/join/v1"
)

// IJoinV1 入会申请服务V1接口定义
type IJoinV1 interface {
	// Apply 提交入会申请
	Apply(ctx context.Context, req *v1.ApplyReq) (res *v1.ApplyRes, err error)

	// List 获取入会申请列表
	List(ctx context.Context, req *v1.ListReq) (res *v1.ListRes, err error)

	// Detail 获取入会申请详情
	Detail(ctx context.Context, req *v1.DetailReq) (res *v1.DetailRes, err error)

	// Review 审核入会申请
	Review(ctx context.Context, req *v1.ReviewReq) (res *v1.ReviewRes, err error)
}

// Controller 控制器接口
type Controller interface {
	// Register 注册路由
	Register(group *ghttp.RouterGroup)

	// V1 获取 V1 接口
	V1() IJoinV1
}

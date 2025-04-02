package member

import (
	"context"

	"github.com/gogf/gf/v2/net/ghttp"
	v1 "github.com/patient-fyd/jxust-softhub-api/api/member/v1"
)

// IMemberV1 成员管理服务V1接口定义
type IMemberV1 interface {
	// List 获取成员列表
	List(ctx context.Context, req *v1.ListReq) (res *v1.ListRes, err error)

	// Detail 获取成员详情
	Detail(ctx context.Context, req *v1.DetailReq) (res *v1.DetailRes, err error)

	// Create 创建成员
	Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error)

	// Update 更新成员信息
	Update(ctx context.Context, req *v1.UpdateReq) (res *v1.UpdateRes, err error)

	// Delete 删除成员
	Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error)
}

// Controller 控制器接口
type Controller interface {
	// Register 注册路由
	Register(group *ghttp.RouterGroup)

	// V1 获取 V1 接口
	V1() IMemberV1
}

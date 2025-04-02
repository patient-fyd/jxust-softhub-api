package config

import (
	"context"

	"github.com/gogf/gf/v2/net/ghttp"
	v1 "github.com/patient-fyd/jxust-softhub-api/api/config/v1"
)

// IConfigV1 系统配置服务V1接口定义
type IConfigV1 interface {
	// Get 获取系统配置
	Get(ctx context.Context, req *v1.GetReq) (res *v1.GetRes, err error)

	// List 获取配置列表
	List(ctx context.Context, req *v1.ListReq) (res *v1.ListRes, err error)

	// Set 设置系统配置
	Set(ctx context.Context, req *v1.SetReq) (res *v1.SetRes, err error)

	// Delete 删除系统配置
	Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error)
}

// Controller 控制器接口
type Controller interface {
	// Register 注册路由
	Register(group *ghttp.RouterGroup)

	// V1 获取 V1 接口
	V1() IConfigV1
}

package stat

import (
	"context"

	"github.com/gogf/gf/v2/net/ghttp"
	v1 "github.com/patient-fyd/jxust-softhub-api/api/stat/v1"
)

// IStatV1 统计分析服务V1接口定义
type IStatV1 interface {
	// VisitStat 获取网站访问统计数据
	VisitStat(ctx context.Context, req *v1.VisitStatReq) (res *v1.VisitStatRes, err error)

	// ActivityStat 获取活动统计数据
	ActivityStat(ctx context.Context, req *v1.ActivityStatReq) (res *v1.ActivityStatRes, err error)

	// UserStat 获取用户统计数据
	UserStat(ctx context.Context, req *v1.UserStatReq) (res *v1.UserStatRes, err error)
}

// Controller 控制器接口
type Controller interface {
	// Register 注册路由
	Register(group *ghttp.RouterGroup)

	// V1 获取 V1 接口
	V1() IStatV1
}

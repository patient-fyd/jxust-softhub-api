package openapi

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

// Controller OpenAPI控制器
type Controller struct{}

// New 创建OpenAPI控制器实例
func New() *Controller {
	return &Controller{}
}

// GetOpenAPIDoc 获取OpenAPI文档
func (c *Controller) GetOpenAPIDoc(r *ghttp.Request) {
	// 返回API信息，避免重定向
	r.Response.WriteJson(g.Map{
		"code":    0,
		"msg":     "API文档已自动生成",
		"traceid": r.GetCtxVar("TraceId", "").String(),
		"data": g.Map{
			"title":       "JXUST SoftHub API",
			"description": "JXUST SoftHub 前台与后台管理系统接口文档",
			"version":     "1.0.0",
		},
	})
}

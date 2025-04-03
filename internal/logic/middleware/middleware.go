package middleware

import (
	"github.com/patient-fyd/jxust-softhub-api/internal/service"
)

// sMiddleware 中间件服务
type sMiddleware struct{}

// New 创建中间件服务
func New() *sMiddleware {
	return &sMiddleware{}
}

// init 初始化
func init() {
	service.RegisterMiddleware(New())
}

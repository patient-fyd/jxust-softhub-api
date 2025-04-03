package middleware

import (
	"context"
	"strings"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/golang-jwt/jwt/v4"

	"github.com/patient-fyd/jxust-softhub-api/internal/codes"
)

// AuthMiddleware JWT认证中间件
func (s *sMiddleware) AuthMiddleware(r *ghttp.Request) {
	// 不需要认证的路径直接放行
	if isPublicPath(r.URL.Path) {
		r.Middleware.Next()
		return
	}

	// 获取Authorization头
	authHeader := r.GetHeader("Authorization")
	if authHeader == "" {
		r.Response.WriteJson(g.Map{
			"code":    codes.CodeNotAuthorized.Code(),
			"msg":     "未授权访问",
			"traceid": "",
			"data":    nil,
		})
		r.ExitAll()
		return
	}

	// 校验Bearer格式
	parts := strings.SplitN(authHeader, " ", 2)
	if !(len(parts) == 2 && parts[0] == "Bearer") {
		r.Response.WriteJson(g.Map{
			"code":    codes.CodeNotAuthorized.Code(),
			"msg":     "认证格式错误",
			"traceid": "",
			"data":    nil,
		})
		r.ExitAll()
		return
	}

	// 解析JWT Token
	tokenString := parts[1]
	claims := jwt.MapClaims{}

	// 解析token
	jwtSecret := g.Cfg().MustGet(r.Context(), "jwt.secret").String()
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})

	if err != nil {
		g.Log().Error(r.Context(), "Parse token error:", err)
		r.Response.WriteJson(g.Map{
			"code":    codes.CodeNotAuthorized.Code(),
			"msg":     "无效的认证令牌",
			"traceid": "",
			"data":    nil,
		})
		r.ExitAll()
		return
	}

	if !token.Valid {
		r.Response.WriteJson(g.Map{
			"code":    codes.CodeNotAuthorized.Code(),
			"msg":     "认证令牌已过期",
			"traceid": "",
			"data":    nil,
		})
		r.ExitAll()
		return
	}

	// 从token中提取userId并设置到上下文中
	if userId, ok := claims["userId"].(float64); ok {
		// 创建新的上下文，将用户ID放入
		ctx := context.WithValue(r.Context(), "userId", uint(userId))
		r.SetCtx(ctx)
	} else {
		r.Response.WriteJson(g.Map{
			"code":    codes.CodeNotAuthorized.Code(),
			"msg":     "无效的用户信息",
			"traceid": "",
			"data":    nil,
		})
		r.ExitAll()
		return
	}

	// 继续处理请求
	r.Middleware.Next()
}

// isPublicPath 判断是否为公开路径
func isPublicPath(path string) bool {
	// 定义公开路径列表
	publicPaths := []string{
		"/api/auth/v1/login",
		"/api/auth/v1/register",
		"/api/auth/v1/admin/login",
		"/api/auth/v1/refresh",
		"/api.json",
		"/swagger",
		"/openapi.json", // 自定义OpenAPI文档路径
		"/swagger",      // Swagger UI路径
		"/api.json",     // GoFrame默认OpenAPI文档路径
	}

	for _, p := range publicPaths {
		if strings.HasPrefix(path, p) {
			return true
		}
	}
	return false
}

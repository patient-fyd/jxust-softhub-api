package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

// RegisterReq 用户注册请求
type RegisterReq struct {
	g.Meta   `path:"/api/auth/v1/register" method:"post" tags:"AuthService" summary:"用户注册"`
	UserName string `p:"userName" v:"required#用户名不能为空"`
	Password string `p:"password" v:"required|length:6,30#密码不能为空|密码长度应当在6到30之间"`
	Name     string `p:"name" v:"required#姓名不能为空"`
	Email    string `p:"email" v:"required|email#邮箱不能为空|邮箱格式不正确"`
	Phone    string `p:"phone" v:"required|phone#手机号不能为空|手机号格式不正确"`
}

// RegisterRes 用户注册响应
type RegisterRes struct {
	User *UserInfo `json:"user"` // 用户信息
}

// LoginReq 用户登录请求
type LoginReq struct {
	g.Meta   `path:"/api/auth/v1/login" method:"post" tags:"AuthService" summary:"用户登录"`
	UserName string `p:"userName" v:"required#用户名不能为空"`
	Password string `p:"password" v:"required#密码不能为空"`
}

// LoginRes 用户登录响应
type LoginRes struct {
	Token string    `json:"token"` // JWT Token
	User  *UserInfo `json:"user"`  // 用户信息
}

// UserInfo 用户信息
type UserInfo struct {
	UserId   uint   `json:"userId"`   // 用户ID
	UserName string `json:"userName"` // 用户名
	Name     string `json:"name"`     // 真实姓名
	RoleId   uint   `json:"roleId"`   // 用户角色ID
}

// AdminLoginReq 管理员登录请求
type AdminLoginReq struct {
	g.Meta   `path:"/api/auth/v1/admin/login" method:"post" tags:"AuthService" summary:"终极管理员登录"`
	UserName string `p:"userName" v:"required#用户名不能为空"`
	Password string `p:"password" v:"required#密码不能为空"`
}

// AdminLoginRes 管理员登录响应
type AdminLoginRes struct {
	Token string    `json:"token"` // JWT Token
	User  *UserInfo `json:"user"`  // 用户信息
}

// RefreshTokenReq 刷新Token请求
type RefreshTokenReq struct {
	g.Meta `path:"/api/auth/v1/refresh" method:"post" tags:"AuthService" summary:"刷新Token"`
	Token  string `p:"token" v:"required#Token不能为空"`
}

// RefreshTokenRes 刷新Token响应
type RefreshTokenRes struct {
	Token string `json:"token"` // 新Token
}

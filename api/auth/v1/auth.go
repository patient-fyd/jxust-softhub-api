package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

// RegisterReq 用户注册请求
type RegisterReq struct {
	g.Meta   `path:"/api/auth/register" method:"post" tags:"AuthService" summary:"用户注册"`
	UserName string `p:"userName" v:"required|length:4,30#用户名不能为空|用户名长度应当在:min到:max之间"`
	Password string `p:"password" v:"required|length:6,30#密码不能为空|密码长度应当在:min到:max之间"`
	Name     string `p:"name" v:"required|length:2,50#真实姓名不能为空|真实姓名长度应当在:min到:max之间"`
	Email    string `p:"email" v:"required|email#邮箱不能为空|邮箱格式不正确"`
	Phone    string `p:"phone" v:"required|phone#手机号不能为空|手机号格式不正确"`
}

// RegisterRes 用户注册响应
type RegisterRes struct {
	UserId     uint   `json:"userId"`     // 用户ID
	UserName   string `json:"userName"`   // 用户名
	CreateTime string `json:"createTime"` // 创建时间
}

// LoginReq 用户登录请求
type LoginReq struct {
	g.Meta   `path:"/api/auth/login" method:"post" tags:"AuthService" summary:"用户登录"`
	UserName string `p:"userName" v:"required#用户名不能为空"`
	Password string `p:"password" v:"required#密码不能为空"`
}

// LoginRes 用户登录响应
type LoginRes struct {
	Token string    `json:"token"` // JWT令牌
	User  LoginUser `json:"user"`  // 用户基本信息
}

// LoginUser 登录用户信息
type LoginUser struct {
	UserId   uint   `json:"userId"`   // 用户ID
	UserName string `json:"userName"` // 用户名
	Name     string `json:"name"`     // 真实姓名
}

// RefreshTokenReq Token刷新请求
type RefreshTokenReq struct {
	g.Meta `path:"/api/auth/refresh" method:"post" tags:"AuthService" summary:"刷新Token"`
	Token  string `p:"token" v:"required#Token不能为空"`
}

// RefreshTokenRes Token刷新响应
type RefreshTokenRes struct {
	Token string `json:"token"` // 新的JWT令牌
}

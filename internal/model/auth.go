package model

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// UserRegisterInput 用户注册输入
type UserRegisterInput struct {
	UserName string // 用户名
	Password string // 密码
	Name     string // 真实姓名
	Email    string // 电子邮箱
	Phone    string // 联系电话
}

// UserRegisterOutput 用户注册输出
type UserRegisterOutput struct {
	UserId   uint   // 用户ID
	UserName string // 用户名
	Name     string // 真实姓名
	RoleId   uint   // 角色ID
}

// UserLoginInput 用户登录输入
type UserLoginInput struct {
	UserName string // 用户名
	Password string // 密码
}

// UserLoginOutput 用户登录输出
type UserLoginOutput struct {
	Token     string      // JWT Token
	UserId    uint        // 用户ID
	UserName  string      // 用户名
	Name      string      // 真实姓名
	RoleId    uint        // 角色ID
	LoginTime *gtime.Time // 登录时间
}

// AdminLoginInput 管理员登录输入
type AdminLoginInput struct {
	UserName string // 用户名
	Password string // 密码
}

// AdminLoginOutput 管理员登录输出
type AdminLoginOutput struct {
	Token     string      // JWT Token
	UserId    uint        // 用户ID
	UserName  string      // 用户名
	Name      string      // 真实姓名
	RoleId    uint        // 角色ID
	LoginTime *gtime.Time // 登录时间
}

// RefreshTokenInput 刷新Token输入
type RefreshTokenInput struct {
	Token string // 原Token
}

// RefreshTokenOutput 刷新Token输出
type RefreshTokenOutput struct {
	Token string // 新Token
}

// UserRoleEnum 用户角色枚举
type UserRoleEnum int

const (
	UserRoleUser  UserRoleEnum = 1 // 普通用户
	UserRoleAdmin UserRoleEnum = 2 // 管理员
)

// UserSessionInfo 用户会话信息
type UserSessionInfo struct {
	UserId   uint         // 用户ID
	UserName string       // 用户名
	Name     string       // 真实姓名
	RoleId   UserRoleEnum // 角色ID
}

// ClaimsUser JWT令牌中的用户信息
type ClaimsUser struct {
	UserId   uint         `json:"uid"`   // 用户ID
	UserName string       `json:"uname"` // 用户名
	RoleId   UserRoleEnum `json:"rid"`   // 角色ID
}

// ClaimsOutput JWT解析结果
type ClaimsOutput struct {
	UserId    uint   `json:"userId"`
	UserName  string `json:"userName"`
	Name      string `json:"name"`
	RoleId    uint   `json:"roleId"`
	ExpiresAt int64  `json:"expiresAt"`
}

// CodeInput 验证码输入
type CodeInput struct {
	Code string `json:"code" v:"required#验证码不能为空"`
}

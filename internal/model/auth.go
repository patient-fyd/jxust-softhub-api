package model

// RegisterInput 用户注册输入参数
type RegisterInput struct {
	UserName string // 用户名
	Password string // 密码
	Name     string // 真实姓名
	Email    string // 邮箱
	Phone    string // 手机号
}

// RegisterOutput 用户注册输出参数
type RegisterOutput struct {
	UserId     uint   // 用户ID
	UserName   string // 用户名
	CreateTime string // 创建时间
}

// LoginInput 用户登录输入参数
type LoginInput struct {
	UserName string // 用户名
	Password string // 密码
}

// LoginOutput 用户登录输出参数
type LoginOutput struct {
	Token string     // JWT令牌
	User  *LoginUser // 用户信息
}

// LoginUser 登录用户信息
type LoginUser struct {
	UserId   uint   // 用户ID
	UserName string // 用户名
	Name     string // 真实姓名
}

// TokenRefreshInput 刷新Token输入参数
type TokenRefreshInput struct {
	Token string // JWT令牌
}

// TokenRefreshOutput 刷新Token输出参数
type TokenRefreshOutput struct {
	Token string // 新的JWT令牌
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

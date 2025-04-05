package consts

const (
	CtxAccessUserKey     = "user"
	CtxAccessUserMailKey = "mail"

	AccessUserHeader     = "X-Consumer-Custom-ID"
	AccessUserMailHeader = "X-Consumer-Username"
)

const (
	CliLoggerName = "cli"
)

// 用户认证相关常量
const (
	CtxUserIdKey   = "userId"   // 用户ID上下文键
	CtxUserNameKey = "userName" // 用户名上下文键
	CtxUserRoleKey = "roleId"   // 用户角色上下文键

	TokenType = "Bearer" // Token类型
)

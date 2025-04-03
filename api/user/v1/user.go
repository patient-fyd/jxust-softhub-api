package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

// InfoReq 获取用户信息请求
type InfoReq struct {
	g.Meta `path:"/api/user/v1/info" method:"get" tags:"UserService" summary:"获取当前用户信息"`
}

// InfoRes 获取用户信息响应
type InfoRes struct {
	UserId    uint   `json:"userId"`    // 用户ID
	UserName  string `json:"userName"`  // 用户名
	Name      string `json:"name"`      // 真实姓名
	Email     string `json:"email"`     // 邮箱
	Phone     string `json:"phone"`     // 手机号
	Avatar    string `json:"avatar"`    // 头像
	RoleId    uint   `json:"roleId"`    // 角色ID
	RoleName  string `json:"roleName"`  // 角色名称
	CreatedAt string `json:"createdAt"` // 创建时间
	UpdatedAt string `json:"updatedAt"` // 更新时间
}

// ListReq 获取用户列表请求
type ListReq struct {
	g.Meta   `path:"/api/user/v1/list" method:"get" tags:"UserService" summary:"获取用户列表(管理员)"`
	PageNum  int    `p:"page_num" v:"min:0#分页号码错误" dc:"分页号码, 默认1"`
	PageSize int    `p:"page_size" v:"max:500#每页数量最大500条" dc:"分页数量, 最大500"`
	Keyword  string `p:"keyword" dc:"搜索关键词"`
}

// ListRes 获取用户列表响应
type ListRes struct {
	List     []UserInfo `json:"list"`      // 用户列表
	Total    int        `json:"total"`     // 总数
	PageNum  int        `json:"page_num"`  // 分页号码
	PageSize int        `json:"page_size"` // 分页数量
}

// UserInfo 用户详情
type UserInfo struct {
	UserId    uint   `json:"userId"`    // 用户ID
	UserName  string `json:"userName"`  // 用户名
	Name      string `json:"name"`      // 真实姓名
	Email     string `json:"email"`     // 邮箱
	Phone     string `json:"phone"`     // 手机号
	Avatar    string `json:"avatar"`    // 头像
	RoleId    uint   `json:"roleId"`    // 角色ID
	RoleName  string `json:"roleName"`  // 角色名称
	CreatedAt string `json:"createdAt"` // 创建时间
	UpdatedAt string `json:"updatedAt"` // 更新时间
}

// AssignRoleReq 分配用户角色请求
type AssignRoleReq struct {
	g.Meta `path:"/api/user/v1/{userId}/role" method:"put" tags:"UserService" summary:"分配用户角色"`
	UserId uint `in:"path" v:"required#用户ID不能为空"`
	RoleId uint `p:"roleId" v:"required#角色ID不能为空"`
}

// AssignRoleRes 分配用户角色响应
type AssignRoleRes struct {
	Success bool `json:"success" dc:"是否成功"`
}

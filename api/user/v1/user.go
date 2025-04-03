package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

// InfoReq 获取当前用户信息请求
type InfoReq struct {
	g.Meta `path:"/api/user/info" method:"get" tags:"UserService" summary:"获取当前用户信息"`
}

// InfoRes 获取当前用户信息响应
type InfoRes struct {
	UserId    uint   `json:"userId"`    // 用户ID
	UserName  string `json:"userName"`  // 用户名
	Name      string `json:"name"`      // 真实姓名
	Email     string `json:"email"`     // 电子邮箱
	Phone     string `json:"phone"`     // 联系电话
	Avatar    string `json:"avatar"`    // 头像
	RoleId    uint   `json:"roleId"`    // 角色ID
	RoleName  string `json:"roleName"`  // 角色名称
	CreatedAt string `json:"createdAt"` // 创建时间
	UpdatedAt string `json:"updatedAt"` // 更新时间
}

// ListReq 获取用户列表请求（管理员接口）
type ListReq struct {
	g.Meta   `path:"/api/user/list" method:"get" tags:"UserService" summary:"获取用户列表（管理员接口）"`
	PageNum  int    `p:"page_num" v:"min:0#分页号码错误" dc:"分页号码, 默认1"`
	PageSize int    `p:"page_size" v:"max:500#每页数量最大500条" dc:"分页数量, 最大500"`
	Keyword  string `p:"keyword" dc:"搜索关键字"`
}

// ListRes 获取用户列表响应
type ListRes struct {
	List     []UserInfo `json:"list"`      // 用户列表
	Total    int        `json:"total"`     // 总数
	PageNum  int        `json:"page_num"`  // 分页号码
	PageSize int        `json:"page_size"` // 分页数量
}

// UserInfo 用户信息
type UserInfo struct {
	UserId    uint   `json:"userId"`    // 用户ID
	UserName  string `json:"userName"`  // 用户名
	Name      string `json:"name"`      // 真实姓名
	Email     string `json:"email"`     // 电子邮箱
	Phone     string `json:"phone"`     // 联系电话
	Avatar    string `json:"avatar"`    // 头像
	RoleId    uint   `json:"roleId"`    // 角色ID
	RoleName  string `json:"roleName"`  // 角色名称
	CreatedAt string `json:"createdAt"` // 创建时间
	UpdatedAt string `json:"updatedAt"` // 更新时间
}

// 添加用户角色分配相关的API定义
type AssignRoleReq struct {
	g.Meta `path:"/user/{userId}/role" method:"put" tags:"UserService" summary:"分配用户角色"`
	UserId uint `v:"required|min:1" json:"userId" in:"path" dc:"用户ID"`
	RoleId uint `v:"required|min:1" json:"roleId" dc:"角色ID"`
}

type AssignRoleRes struct {
	g.Meta  `mime:"application/json"`
	Success bool `json:"success" dc:"是否成功"`
}

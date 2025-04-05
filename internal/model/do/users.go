// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Users is the golang structure of table users for DAO operations like Where/Data.
type Users struct {
	g.Meta         `orm:"table:users, do:true"`
	UserId         interface{} // 用户ID，主键，自增
	UserName       interface{} // 用户名，登录和显示名称
	Password       interface{} // 用户密码，存储加密后的密码
	Name           interface{} // 真实姓名
	RoleId         interface{} // 角色ID，关联 roles 表，标识用户所属角色
	Avatar         interface{} // 用户头像图片URL
	JoinYear       interface{} // 入会年份，格式如2025
	Email          interface{} // 用户邮箱
	Phone          interface{} // 用户联系电话
	CreateTime     *gtime.Time // 记录创建时间
	UpdateTime     *gtime.Time // 记录最后更新时间
	FollowerCount  interface{} // 粉丝数量
	FollowingCount interface{} // 关注数量
}

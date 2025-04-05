// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Users is the golang structure for table users.
type Users struct {
	UserId         uint        `json:"user_id"         description:"用户ID，主键，自增"`
	UserName       string      `json:"user_name"       description:"用户名，登录和显示名称"`
	Password       string      `json:"password"        description:"用户密码，存储加密后的密码"`
	Name           string      `json:"name"            description:"真实姓名"`
	RoleId         uint        `json:"role_id"         description:"角色ID，关联 roles 表，标识用户所属角色"`
	Avatar         string      `json:"avatar"          description:"用户头像图片URL"`
	JoinYear       string      `json:"join_year"       description:"入会年份，格式如2025"`
	Email          string      `json:"email"           description:"用户邮箱"`
	Phone          string      `json:"phone"           description:"用户联系电话"`
	CreateTime     *gtime.Time `json:"create_time"     description:"记录创建时间"`
	UpdateTime     *gtime.Time `json:"update_time"     description:"记录最后更新时间"`
	FollowerCount  uint        `json:"follower_count"  description:"粉丝数量"`
	FollowingCount uint        `json:"following_count" description:"关注数量"`
}

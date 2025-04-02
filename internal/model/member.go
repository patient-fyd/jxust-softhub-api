package model

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// MemberListInput 成员列表查询输入
type MemberListInput struct {
	Department string // 部门筛选
	Grade      string // 年级筛选
	IsCore     int    // 是否核心成员：0-否，1-是
	Status     int    // 成员状态筛选
	PageNum    int    // 分页号码
	PageSize   int    // 分页数量
}

// MemberListOutput 成员列表查询输出
type MemberListOutput struct {
	List     []*MemberInfo // 成员列表
	Total    int           // 总数
	PageNum  int           // 分页号码
	PageSize int           // 分页数量
}

// MemberInfo 成员信息
type MemberInfo struct {
	MemberId     uint        // 成员ID
	UserId       uint        // 用户ID
	UserName     string      // 用户名
	Name         string      // 真实姓名
	Avatar       string      // 头像
	Grade        string      // 年级
	JoinYear     string      // 入会年份
	Department   string      // 所属部门
	Position     string      // 职位
	Skills       string      // 技能描述
	Introduction string      // 个人简介
	IsCore       int         // 是否核心成员
	Status       int         // 成员状态
	CreateTime   *gtime.Time // 创建时间
	UpdateTime   *gtime.Time // 更新时间
}

// MemberDetailInput 成员详情查询输入
type MemberDetailInput struct {
	MemberId uint // 成员ID
}

// MemberDetailOutput 成员详情查询输出
type MemberDetailOutput struct {
	*MemberInfo
}

// MemberCreateInput 创建成员输入
type MemberCreateInput struct {
	UserId       uint   // 用户ID
	Grade        string // 年级
	JoinYear     string // 入会年份
	Department   string // 所属部门
	Position     string // 职位
	Skills       string // 技能描述
	Introduction string // 个人简介
	IsCore       int    // 是否核心成员
	Status       int    // 成员状态
}

// MemberCreateOutput 创建成员输出
type MemberCreateOutput struct {
	MemberId uint // 成员ID
}

// MemberUpdateInput 更新成员输入
type MemberUpdateInput struct {
	MemberId     uint   // 成员ID
	Grade        string // 年级
	Department   string // 所属部门
	Position     string // 职位
	Skills       string // 技能描述
	Introduction string // 个人简介
	IsCore       int    // 是否核心成员
	Status       int    // 成员状态
}

// MemberUpdateOutput 更新成员输出
type MemberUpdateOutput struct{}

// MemberDeleteInput 删除成员输入
type MemberDeleteInput struct {
	MemberId uint // 成员ID
}

// MemberDeleteOutput 删除成员输出
type MemberDeleteOutput struct{}

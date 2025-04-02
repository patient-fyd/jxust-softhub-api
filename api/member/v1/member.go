package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

// ListReq 成员列表请求
type ListReq struct {
	g.Meta     `path:"/api/member/list" method:"get" tags:"MemberService" summary:"获取成员列表"`
	Department string `p:"department" dc:"部门筛选"`
	Grade      string `p:"grade" dc:"年级筛选"`
	IsCore     int    `p:"is_core" dc:"是否核心成员：0-否，1-是"`
	Status     int    `p:"status" dc:"成员状态筛选"`
	// 分页参数
	PageNum  int `p:"page_num" v:"min:0#分页号码错误" dc:"分页号码, 默认1"`
	PageSize int `p:"page_size" v:"max:500#每页数量最大500条" dc:"分页数量, 最大500"`
}

// ListRes 成员列表响应
type ListRes struct {
	List     []MemberInfo `json:"list"`      // 成员列表
	Total    int          `json:"total"`     // 总数
	PageNum  int          `json:"page_num"`  // 分页号码
	PageSize int          `json:"page_size"` // 分页数量
}

// MemberInfo 成员信息
type MemberInfo struct {
	MemberId     uint   `json:"memberId"`     // 成员ID
	UserId       uint   `json:"userId"`       // 用户ID
	UserName     string `json:"userName"`     // 用户名
	Name         string `json:"name"`         // 真实姓名
	Avatar       string `json:"avatar"`       // 头像
	Grade        string `json:"grade"`        // 年级
	JoinYear     int    `json:"joinYear"`     // 入会年份
	Department   string `json:"department"`   // 所属部门
	Position     string `json:"position"`     // 职位
	Skills       string `json:"skills"`       // 技能描述
	Introduction string `json:"introduction"` // 个人简介
	IsCore       int    `json:"isCore"`       // 是否核心成员
	Status       int    `json:"status"`       // 成员状态
}

// DetailReq 成员详情请求
type DetailReq struct {
	g.Meta   `path:"/api/member/detail" method:"get" tags:"MemberService" summary:"获取成员详情"`
	MemberId uint `p:"memberId" v:"required#成员ID不能为空"`
}

// DetailRes 成员详情响应
type DetailRes struct {
	MemberInfo
}

// CreateReq 创建成员请求
type CreateReq struct {
	g.Meta       `path:"/api/member/create" method:"post" tags:"MemberService" summary:"创建成员"`
	UserId       uint   `p:"userId" v:"required#用户ID不能为空"`
	Grade        string `p:"grade" v:"required#年级不能为空"`
	JoinYear     int    `p:"joinYear" v:"required#入会年份不能为空"`
	Department   string `p:"department"`
	Position     string `p:"position"`
	Skills       string `p:"skills"`
	Introduction string `p:"introduction"`
	IsCore       int    `p:"isCore"`
	Status       int    `p:"status" v:"required#成员状态不能为空"`
}

// CreateRes 创建成员响应
type CreateRes struct {
	MemberId uint `json:"memberId"` // 成员ID
}

// UpdateReq 更新成员请求
type UpdateReq struct {
	g.Meta       `path:"/api/member/update" method:"put" tags:"MemberService" summary:"更新成员信息"`
	MemberId     uint   `p:"memberId" v:"required#成员ID不能为空"`
	Grade        string `p:"grade"`
	Department   string `p:"department"`
	Position     string `p:"position"`
	Skills       string `p:"skills"`
	Introduction string `p:"introduction"`
	IsCore       int    `p:"isCore"`
	Status       int    `p:"status"`
}

// UpdateRes 更新成员响应
type UpdateRes struct{}

// DeleteReq 删除成员请求
type DeleteReq struct {
	g.Meta   `path:"/api/member/delete" method:"delete" tags:"MemberService" summary:"删除成员"`
	MemberId uint `p:"memberId" v:"required#成员ID不能为空"`
}

// DeleteRes 删除成员响应
type DeleteRes struct{}

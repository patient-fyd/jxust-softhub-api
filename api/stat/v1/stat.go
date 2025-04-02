package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

// VisitStatReq 网站访问统计数据请求
type VisitStatReq struct {
	g.Meta    `path:"/api/stat/visit" method:"get" tags:"StatService" summary:"获取网站访问统计数据"`
	StartDate string `p:"start_date" v:"date#开始日期格式不正确" dc:"开始日期，格式：YYYY-MM-DD"`
	EndDate   string `p:"end_date" v:"date#结束日期格式不正确" dc:"结束日期，格式：YYYY-MM-DD"`
}

// VisitStatRes 网站访问统计数据响应
type VisitStatRes struct {
	List []VisitStatInfo `json:"list"` // 统计数据列表
}

// VisitStatInfo 访问统计数据
type VisitStatInfo struct {
	VisitDate     string  `json:"visitDate"`     // 访问日期
	PageView      int     `json:"pageView"`      // 页面浏览量
	UniqueVisitor int     `json:"uniqueVisitor"` // 独立访客数
	NewVisitor    int     `json:"newVisitor"`    // 新访客数
	AvgStayTime   int     `json:"avgStayTime"`   // 平均停留时间（秒）
	BounceRate    float64 `json:"bounceRate"`    // 跳出率（百分比）
}

// ActivityStatReq 活动统计数据请求
type ActivityStatReq struct {
	g.Meta    `path:"/api/stat/activity" method:"get" tags:"StatService" summary:"获取活动统计数据"`
	StartDate string `p:"start_date" v:"date#开始日期格式不正确" dc:"开始日期，格式：YYYY-MM-DD"`
	EndDate   string `p:"end_date" v:"date#结束日期格式不正确" dc:"结束日期，格式：YYYY-MM-DD"`
}

// ActivityStatRes 活动统计数据响应
type ActivityStatRes struct {
	TotalActivity     int                `json:"totalActivity"`     // 活动总数
	TotalRegistration int                `json:"totalRegistration"` // 报名总人数
	AvgRegistration   float64            `json:"avgRegistration"`   // 平均每个活动报名人数
	ActivityList      []ActivityStatInfo `json:"activityList"`      // 活动统计列表
}

// ActivityStatInfo 活动统计数据
type ActivityStatInfo struct {
	ActivityId      uint   `json:"activityId"`      // 活动ID
	Title           string `json:"title"`           // 活动标题
	StartTime       string `json:"startTime"`       // 活动开始时间
	EndTime         string `json:"endTime"`         // 活动结束时间
	RegistrationNum int    `json:"registrationNum"` // 报名人数
	ApprovedNum     int    `json:"approvedNum"`     // 审核通过人数
	RejectedNum     int    `json:"rejectedNum"`     // 拒绝人数
	PendingNum      int    `json:"pendingNum"`      // 待审核人数
}

// UserStatReq 用户统计数据请求
type UserStatReq struct {
	g.Meta `path:"/api/stat/user" method:"get" tags:"StatService" summary:"获取用户统计数据"`
}

// UserStatRes 用户统计数据响应
type UserStatRes struct {
	TotalUser            int             `json:"totalUser"`            // 用户总数
	NewUserLast30Days    int             `json:"newUserLast30Days"`    // 最近30天新增用户数
	ActiveUserLast30Days int             `json:"activeUserLast30Days"` // 最近30天活跃用户数
	UserByRole           map[string]int  `json:"userByRole"`           // 按角色统计用户数
	MemberStats          MemberStatsInfo `json:"memberStats"`          // 协会成员统计
}

// MemberStatsInfo 成员统计信息
type MemberStatsInfo struct {
	TotalMember  int            `json:"totalMember"`  // 成员总数
	CoreMember   int            `json:"coreMember"`   // 核心成员数
	ByDepartment map[string]int `json:"byDepartment"` // 按部门统计
	ByGrade      map[string]int `json:"byGrade"`      // 按年级统计
}

package model

import (
	"time"
)

// 活动状态常量
const (
	ActivityStatusNotStarted = 0 // 未开始
	ActivityStatusOngoing    = 1 // 进行中
	ActivityStatusEnded      = 2 // 已结束
)

// 报名状态常量
const (
	RegistrationStatusPending  = 0 // 待审核
	RegistrationStatusApproved = 1 // 已通过
	RegistrationStatusRejected = 2 // 已拒绝
)

// 活动创建输入
type ActivityCreateInput struct {
	Title           string    // 活动标题
	Description     string    // 活动详细描述
	StartTime       time.Time // 活动开始时间
	EndTime         time.Time // 活动结束时间
	Location        string    // 活动举办地点
	MaxParticipants int       // 最大参与人数限制
}

// 活动创建输出
type ActivityCreateOutput struct {
	ActivityId uint // 活动ID
}

// 活动更新输入
type ActivityUpdateInput struct {
	ActivityId      uint       // 活动ID
	Title           *string    // 活动标题
	Description     *string    // 活动详细描述
	StartTime       *time.Time // 活动开始时间
	EndTime         *time.Time // 活动结束时间
	Location        *string    // 活动举办地点
	MaxParticipants *int       // 最大参与人数限制
	Status          *int       // 活动状态
}

// 活动更新输出
type ActivityUpdateOutput struct {
	ActivityId uint // 活动ID
}

// 活动删除输入
type ActivityDeleteInput struct {
	ActivityId uint // 活动ID
}

// 活动列表查询输入
type ActivityListInput struct {
	Status   *int // 活动状态
	Page     int  // 页码
	PageSize int  // 每页数量
}

// 活动信息
type ActivityInfo struct {
	ActivityId      uint      // 活动ID
	Title           string    // 活动标题
	Description     string    // 活动详细描述
	StartTime       time.Time // 活动开始时间
	EndTime         time.Time // 活动结束时间
	Location        string    // 活动举办地点
	MaxParticipants int       // 最大参与人数
	Status          int       // 活动状态：0-未开始, 1-进行中, 2-已结束
	CreateTime      time.Time // 创建时间
}

// 活动列表查询输出
type ActivityListOutput struct {
	List     []ActivityInfo // 活动列表
	Total    int            // 总数
	Page     int            // 当前页码
	PageSize int            // 每页数量
}

// 活动详情查询输入
type ActivityDetailInput struct {
	ActivityId uint // 活动ID
}

// 活动详情查询输出
type ActivityDetailOutput struct {
	Activity ActivityInfo // 活动详情
}

// 活动报名输入
type ActivityRegisterInput struct {
	ActivityId uint   // 活动ID
	UserId     uint   // 用户ID，可选，如果有登录用户
	Name       string // 报名者姓名
	StudentId  string // 报名者学号
	Contact    string // 报名者联系方式
}

// 活动报名输出
type ActivityRegisterOutput struct {
	RegistrationId uint // 报名记录ID
}

// 报名信息
type RegistrationInfo struct {
	RegistrationId uint      // 报名ID
	ActivityId     uint      // 活动ID
	UserId         uint      // 用户ID
	Name           string    // 报名者姓名
	StudentId      string    // 报名者学号
	Contact        string    // 联系方式
	Status         int       // 状态：0-待审核, 1-通过, 2-拒绝
	CreateTime     time.Time // 创建时间
	UpdateTime     time.Time // 更新时间
}

// 报名列表查询输入
type RegistrationListInput struct {
	ActivityId uint // 活动ID
	Status     *int // 报名状态
	Page       int  // 页码
	PageSize   int  // 每页数量
}

// 报名列表查询输出
type RegistrationListOutput struct {
	List     []RegistrationInfo // 报名列表
	Total    int                // 总数
	Page     int                // 当前页码
	PageSize int                // 每页数量
}

// 审核报名输入
type RegistrationReviewInput struct {
	RegistrationId uint   // 报名ID
	Status         int    // 状态：1-通过, 2-拒绝
	Reason         string // 拒绝原因，仅在拒绝时有效
	ReviewerId     uint   // 审核人ID
}

// 审核报名输出
type RegistrationReviewOutput struct {
	RegistrationId uint // 报名ID
}

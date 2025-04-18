package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/patient-fyd/jxust-softhub-api/api"
)

// 活动基本信息
type ActivityInfo struct {
	ActivityId      uint   `json:"activityId"`      // 活动ID
	Title           string `json:"title"`           // 活动标题
	Description     string `json:"description"`     // 活动详细描述
	StartTime       string `json:"startTime"`       // 活动开始时间
	EndTime         string `json:"endTime"`         // 活动结束时间
	Location        string `json:"location"`        // 活动举办地点
	MaxParticipants int    `json:"maxParticipants"` // 最大参与人数
	Status          int    `json:"status"`          // 活动状态：0-未开始, 1-进行中, 2-已结束
	CreateTime      string `json:"createTime"`      // 创建时间
}

// 活动报名信息
type RegistrationInfo struct {
	RegistrationId uint   `json:"registrationId"` // 报名ID
	ActivityId     uint   `json:"activityId"`     // 活动ID
	UserId         uint   `json:"userId"`         // 用户ID
	Name           string `json:"name"`           // 报名者姓名
	StudentId      string `json:"studentId"`      // 报名者学号
	Contact        string `json:"contact"`        // 联系方式
	Status         int    `json:"status"`         // 状态：0-待审核, 1-通过, 2-拒绝
	CreateTime     string `json:"createTime"`     // 创建时间
}

// 创建活动请求
type CreateReq struct {
	g.Meta          `path:"/api/activity/v1/create" method:"post" tags:"ActivityService" summary:"创建活动"`
	Title           string `p:"title" v:"required#活动标题不能为空"`
	Description     string `p:"description" v:"required#活动描述不能为空"`
	StartTime       string `p:"startTime" v:"required|datetime#开始时间不能为空|开始时间格式错误"`
	EndTime         string `p:"endTime" v:"required|datetime#结束时间不能为空|结束时间格式错误"`
	Location        string `p:"location" v:"required#活动地点不能为空"`
	MaxParticipants int    `p:"maxParticipants" v:"min:0#参与人数不能小于0"`
}

// 创建活动响应
type CreateRes struct {
	ActivityId uint `json:"activityId"` // 新创建的活动ID
}

// 更新活动请求
type UpdateReq struct {
	g.Meta          `path:"/api/activity/v1/update" method:"put" tags:"ActivityService" summary:"更新活动"`
	ActivityId      uint   `p:"activityId" v:"required#活动ID不能为空"`
	Title           string `p:"title"`
	Description     string `p:"description"`
	StartTime       string `p:"startTime" v:"datetime#开始时间格式错误"`
	EndTime         string `p:"endTime" v:"datetime#结束时间格式错误"`
	Location        string `p:"location"`
	MaxParticipants int    `p:"maxParticipants" v:"min:0#参与人数不能小于0"`
	Status          int    `p:"status" v:"in:0,1,2#状态值错误"`
}

// 更新活动响应
type UpdateRes struct {
	ActivityId uint `json:"activityId"` // 更新的活动ID
}

// 删除活动请求
type DeleteReq struct {
	g.Meta     `path:"/api/activity/v1/delete/{activityId}" method:"delete" tags:"ActivityService" summary:"删除活动"`
	ActivityId uint `p:"activityId" v:"required#活动ID不能为空"`
}

// 删除活动响应
type DeleteRes struct{}

// 活动列表请求
type ListReq struct {
	g.Meta `path:"/api/activity/v1/list" method:"get" tags:"ActivityService" summary:"获取活动列表"`
	Status int `p:"status" in:"query" dc:"活动状态：0-未开始, 1-进行中, 2-已结束, 不传则查询所有"`
	api.CommonPaginationReq
}

// 活动列表响应
type ListRes struct {
	List []ActivityInfo `json:"list"` // 活动列表
	api.CommonPaginationRes
}

// 活动详情请求
type DetailReq struct {
	g.Meta     `path:"/api/activity/v1/detail/{activityId}" method:"get" tags:"ActivityService" summary:"获取活动详情"`
	ActivityId uint `p:"activityId" v:"required#活动ID不能为空"`
}

// 活动详情响应
type DetailRes struct {
	Activity ActivityInfo `json:"activity"` // 活动详情
}

// 活动报名请求
type RegisterReq struct {
	g.Meta     `path:"/api/activity/v1/register" method:"post" tags:"ActivityService" summary:"活动报名"`
	ActivityId uint   `p:"activityId" v:"required#活动ID不能为空"`
	Name       string `p:"name" v:"required#姓名不能为空"`
	StudentId  string `p:"studentId" v:"required#学号不能为空"`
	Contact    string `p:"contact" v:"required#联系方式不能为空"`
}

// 活动报名响应
type RegisterRes struct {
	RegistrationId uint `json:"registrationId"` // 报名记录ID
}

// 报名列表请求
type RegisterListReq struct {
	g.Meta     `path:"/api/activity/v1/register/list" method:"get" tags:"ActivityService" summary:"获取活动报名列表"`
	ActivityId uint `p:"activityId" v:"required#活动ID不能为空" in:"query"`
	Status     int  `p:"status" in:"query" dc:"报名状态：0-待审核, 1-通过, 2-拒绝, 不传则查询所有"`
	api.CommonPaginationReq
}

// 报名列表响应
type RegisterListRes struct {
	List []RegistrationInfo `json:"list"` // 报名列表
	api.CommonPaginationRes
}

// 审核通过请求
type ApproveRegistrationReq struct {
	g.Meta         `path:"/api/activity/v1/register/approve" method:"post" tags:"ActivityService" summary:"审核通过活动报名"`
	RegistrationId uint `p:"registrationId" v:"required#报名ID不能为空"`
}

// 审核通过响应
type ApproveRegistrationRes struct {
	RegistrationId uint `json:"registrationId"` // 报名记录ID
}

// 审核拒绝请求
type RejectRegistrationReq struct {
	g.Meta         `path:"/api/activity/v1/register/reject" method:"post" tags:"ActivityService" summary:"拒绝活动报名"`
	RegistrationId uint   `p:"registrationId" v:"required#报名ID不能为空"`
	Reason         string `p:"reason" dc:"拒绝原因"`
}

// 审核拒绝响应
type RejectRegistrationRes struct {
	RegistrationId uint `json:"registrationId"` // 报名记录ID
}

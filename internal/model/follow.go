package model

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// FollowToggleInput 关注/取消关注输入参数
type FollowToggleInput struct {
	FollowedId int `json:"followedId" v:"required#被关注对象ID不能为空"` // 被关注对象ID
	FollowType int `json:"followType" v:"required#关注类型不能为空"`    // 关注类型：1-用户，2-圈子，3-话题
}

// FollowToggleOutput 关注/取消关注输出参数
type FollowToggleOutput struct {
	Status        int `json:"status"`        // 操作后状态：1-已关注，0-已取消关注
	FollowerCount int `json:"followerCount"` // 粉丝总数
}

// FollowingListInput 获取关注列表输入参数
type FollowingListInput struct {
	UserId     int `json:"userId,omitempty"`                           // 用户ID，不传则查询当前登录用户
	Page       int `json:"page" v:"required|min:1#页码不能为空|页码必须大于0"`     // 页码
	Size       int `json:"size" v:"required|min:1#每页条数不能为空|每页条数必须大于0"` // 每页条数
	FollowType int `json:"followType" v:"required#关注类型不能为空"`           // 关注类型：1-用户，2-圈子，3-话题
}

// FollowingListOutput 获取关注列表输出参数
type FollowingListOutput struct {
	List  []FollowingListItem `json:"list"`  // 关注列表
	Total int                 `json:"total"` // 总数
	Page  int                 `json:"page"`  // 页码
	Size  int                 `json:"size"`  // 每页条数
}

// FollowingListItem 关注列表项
type FollowingListItem struct {
	FollowedId   int         `json:"followedId"`   // 被关注对象ID
	FollowType   int         `json:"followType"`   // 关注类型：1-用户，2-圈子，3-话题
	Name         string      `json:"name"`         // 名称（用户名/圈子名/话题名）
	Avatar       string      `json:"avatar"`       // 头像
	Introduction string      `json:"introduction"` // 简介
	CreateTime   *gtime.Time `json:"createTime"`   // 关注时间
}

// FollowerListInput 获取粉丝列表输入参数
type FollowerListInput struct {
	UserId int `json:"userId,omitempty"`                           // 用户ID，不传则查询当前登录用户
	Page   int `json:"page" v:"required|min:1#页码不能为空|页码必须大于0"`     // 页码
	Size   int `json:"size" v:"required|min:1#每页条数不能为空|每页条数必须大于0"` // 每页条数
}

// FollowerListOutput 获取粉丝列表输出参数
type FollowerListOutput struct {
	List  []FollowerListItem `json:"list"`  // 粉丝列表
	Total int                `json:"total"` // 总数
	Page  int                `json:"page"`  // 页码
	Size  int                `json:"size"`  // 每页条数
}

// FollowerListItem 粉丝列表项
type FollowerListItem struct {
	UserId       int         `json:"userId"`       // 用户ID
	UserName     string      `json:"userName"`     // 用户名
	Avatar       string      `json:"avatar"`       // 头像
	Introduction string      `json:"introduction"` // 简介
	IsFollowed   bool        `json:"isFollowed"`   // 当前用户是否已关注该粉丝
	CreateTime   *gtime.Time `json:"createTime"`   // 关注时间
}

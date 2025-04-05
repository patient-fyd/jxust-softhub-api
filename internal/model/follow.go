package model

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// FollowToggleReq 关注/取消关注输入
type FollowToggleReq struct {
	FollowedId int `json:"followedId" v:"required#被关注对象ID不能为空"` // 被关注对象ID
	FollowType int `json:"followType" v:"required#关注类型不能为空"`    // 关注类型：1-用户，2-圈子
}

// FollowToggleOutput 关注/取消关注输出
type FollowToggleOutput struct {
	FollowId   uint        `json:"followId"`   // 关注记录ID
	IsFollowed bool        `json:"isFollowed"` // 是否已关注
	FollowTime *gtime.Time `json:"followTime"` // 关注时间
}

// FollowingListReq 获取关注列表输入
type FollowingListReq struct {
	UserId     int `json:"userId" v:"required#用户ID不能为空"`     // 用户ID
	FollowType int `json:"followType" v:"required#关注类型不能为空"` // 关注类型：1-用户，2-圈子
	Page       int `json:"page"`                             // 页码
	Size       int `json:"size"`                             // 每页数量
}

// FollowingItem 关注列表项
type FollowingItem struct {
	ID            int         `json:"id"`            // ID（用户ID或圈子ID）
	Name          string      `json:"name"`          // 名称
	Avatar        string      `json:"avatar"`        // 头像
	Description   string      `json:"description"`   // 描述
	FollowerCount int         `json:"followerCount"` // 粉丝数
	IsFollowed    bool        `json:"isFollowed"`    // 当前用户是否已关注
	FollowTime    *gtime.Time `json:"followTime"`    // 关注时间
}

// FollowingListOutput 获取关注列表输出
type FollowingListOutput struct {
	List  []FollowingItem `json:"list"`  // 关注列表
	Total int             `json:"total"` // 总数
	Page  int             `json:"page"`  // 页码
	Size  int             `json:"size"`  // 每页数量
}

// FollowerListReq 获取粉丝列表输入
type FollowerListReq struct {
	FollowedId int `json:"followedId" v:"required#被关注对象ID不能为空"` // 被关注对象ID
	FollowType int `json:"followType" v:"required#关注类型不能为空"`    // 关注类型：1-用户，2-圈子
	Page       int `json:"page"`                                // 页码
	Size       int `json:"size"`                                // 每页数量
}

// FollowerItem 粉丝列表项
type FollowerItem struct {
	ID                int         `json:"id"`                // 用户ID
	Name              string      `json:"name"`              // 用户名
	Avatar            string      `json:"avatar"`            // 头像
	Description       string      `json:"description"`       // 描述
	FollowerCount     int         `json:"followerCount"`     // 粉丝数
	IsFollowed        bool        `json:"isFollowed"`        // 当前用户是否已关注此粉丝
	IsFollowEachOther bool        `json:"isFollowEachOther"` // 是否互相关注
	FollowTime        *gtime.Time `json:"followTime"`        // 关注时间
}

// FollowerListOutput 获取粉丝列表输出
type FollowerListOutput struct {
	List  []FollowerItem `json:"list"`  // 粉丝列表
	Total int            `json:"total"` // 总数
	Page  int            `json:"page"`  // 页码
	Size  int            `json:"size"`  // 每页数量
}

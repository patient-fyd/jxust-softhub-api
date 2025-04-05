package v1

import (
	"github.com/gogf/gf/v2/frame/g"

	"github.com/patient-fyd/jxust-softhub-api/internal/model"
)

// FollowToggleReq 关注/取消关注请求
type FollowToggleReq struct {
	g.Meta     `path:"/follow-toggle" method:"post" tags:"关注管理" summary:"关注/取消关注"`
	FollowedId int `json:"followedId" v:"required#被关注对象ID不能为空"` // 被关注对象ID
	FollowType int `json:"followType" v:"required#关注类型不能为空"`    // 关注类型：1-用户，2-圈子
}

// FollowToggleRes 关注/取消关注响应
type FollowToggleRes struct {
	FollowId   uint   `json:"followId"`   // 关注记录ID
	IsFollowed bool   `json:"isFollowed"` // 是否已关注
	FollowTime string `json:"followTime"` // 关注时间
}

// FollowingListReq 获取关注列表请求
type FollowingListReq struct {
	g.Meta     `path:"/following/list" method:"get" tags:"关注管理" summary:"获取关注列表"`
	UserId     int `json:"userId" v:"required#用户ID不能为空"`     // 用户ID
	FollowType int `json:"followType" v:"required#关注类型不能为空"` // 关注类型：1-用户，2-圈子
	Page       int `json:"page" d:"1"`                       // 页码
	Size       int `json:"size" d:"10"`                      // 每页数量
}

// FollowingListRes 获取关注列表响应
type FollowingListRes struct {
	List  []model.FollowingItem `json:"list"`  // 关注列表
	Total int                   `json:"total"` // 总数
	Page  int                   `json:"page"`  // 页码
	Size  int                   `json:"size"`  // 每页数量
}

// FollowerListReq 获取粉丝列表请求
type FollowerListReq struct {
	g.Meta     `path:"/follower/list" method:"get" tags:"关注管理" summary:"获取粉丝列表"`
	FollowedId int `json:"followedId" v:"required#被关注对象ID不能为空"` // 被关注对象ID
	FollowType int `json:"followType" v:"required#关注类型不能为空"`    // 关注类型：1-用户
	Page       int `json:"page" d:"1"`                          // 页码
	Size       int `json:"size" d:"10"`                         // 每页数量
}

// FollowerListRes 获取粉丝列表响应
type FollowerListRes struct {
	List  []model.FollowerItem `json:"list"`  // 粉丝列表
	Total int                  `json:"total"` // 总数
	Page  int                  `json:"page"`  // 页码
	Size  int                  `json:"size"`  // 每页数量
}

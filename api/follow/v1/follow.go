package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

// ToggleReq 关注/取消关注请求
type ToggleReq struct {
	g.Meta     `path:"/api/follow/v1/toggle" method:"post" tags:"FollowService" summary:"关注/取消关注"`
	FollowedId int `p:"followedId" v:"required#被关注对象ID不能为空" dc:"被关注对象ID"`
	FollowType int `p:"followType" v:"required#关注类型不能为空" dc:"关注类型：1-用户，2-圈子，3-话题"`
}

// ToggleRes 关注/取消关注响应
type ToggleRes struct {
	Status        int `json:"status"`        // 操作后状态：1-已关注，0-已取消
	FollowerCount int `json:"followerCount"` // 粉丝总数
}

// FollowingListReq 获取关注列表请求
type FollowingListReq struct {
	g.Meta     `path:"/api/follow/v1/following-list" method:"get" tags:"FollowService" summary:"获取关注列表"`
	UserId     int `p:"userId" dc:"用户ID，不传则查询当前登录用户"`
	Page       int `p:"page" v:"required|min:1#页码不能为空|页码必须大于0" dc:"页码"`
	Size       int `p:"size" v:"required|min:1#每页条数不能为空|每页条数必须大于0" dc:"每页条数"`
	FollowType int `p:"followType" v:"required#关注类型不能为空" dc:"关注类型：1-用户，2-圈子，3-话题"`
}

// FollowingListRes 获取关注列表响应
type FollowingListRes struct {
	List  []FollowingItem `json:"list"`  // 关注列表
	Total int             `json:"total"` // 总数
	Page  int             `json:"page"`  // 页码
	Size  int             `json:"size"`  // 每页条数
}

// FollowingItem 关注列表项
type FollowingItem struct {
	FollowedId   int    `json:"followedId"`   // 被关注对象ID
	FollowType   int    `json:"followType"`   // 关注类型：1-用户，2-圈子，3-话题
	Name         string `json:"name"`         // 名称（用户名/圈子名/话题名）
	Avatar       string `json:"avatar"`       // 头像
	Introduction string `json:"introduction"` // 简介
	CreateTime   string `json:"createTime"`   // 关注时间
}

// FollowerListReq 获取粉丝列表请求
type FollowerListReq struct {
	g.Meta `path:"/api/follow/v1/follower-list" method:"get" tags:"FollowService" summary:"获取粉丝列表"`
	UserId int `p:"userId" dc:"用户ID，不传则查询当前登录用户"`
	Page   int `p:"page" v:"required|min:1#页码不能为空|页码必须大于0" dc:"页码"`
	Size   int `p:"size" v:"required|min:1#每页条数不能为空|每页条数必须大于0" dc:"每页条数"`
}

// FollowerListRes 获取粉丝列表响应
type FollowerListRes struct {
	List  []FollowerItem `json:"list"`  // 粉丝列表
	Total int            `json:"total"` // 总数
	Page  int            `json:"page"`  // 页码
	Size  int            `json:"size"`  // 每页条数
}

// FollowerItem 粉丝列表项
type FollowerItem struct {
	UserId       int    `json:"userId"`       // 用户ID
	UserName     string `json:"userName"`     // 用户名
	Avatar       string `json:"avatar"`       // 头像
	Introduction string `json:"introduction"` // 简介
	IsFollowed   bool   `json:"isFollowed"`   // 当前用户是否已关注该粉丝
	CreateTime   string `json:"createTime"`   // 关注时间
}

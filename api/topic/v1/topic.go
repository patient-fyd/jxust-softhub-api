package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

// ListReq 获取话题列表请求
type ListReq struct {
	g.Meta  `path:"/api/topic/v1/list" method:"get" tags:"TopicService" summary:"获取话题列表"`
	Page    int    `p:"page" v:"required|min:1#页码不能为空|页码必须大于0" dc:"页码"`
	Size    int    `p:"size" v:"required|min:1#每页条数不能为空|每页条数必须大于0" dc:"每页条数"`
	Keyword string `p:"keyword" dc:"关键词搜索"`
	UserId  int    `p:"userId" dc:"用户ID，查询该用户关注的话题"`
	IsHot   int    `p:"isHot" dc:"是否热门：0-否，1-是"`
}

// ListRes 获取话题列表响应
type ListRes struct {
	List  []TopicItem `json:"list"`  // 话题列表
	Total int         `json:"total"` // 总数
	Page  int         `json:"page"`  // 页码
	Size  int         `json:"size"`  // 每页条数
}

// TopicItem 话题列表项
type TopicItem struct {
	TopicId     int    `json:"topicId"`     // 话题ID
	Name        string `json:"name"`        // 话题名称
	Description string `json:"description"` // 话题描述
	Icon        string `json:"icon"`        // 话题图标
	PostCount   int    `json:"postCount"`   // 帖子数
	IsHot       int    `json:"isHot"`       // 是否热门：0-否，1-是
	IsFollowed  bool   `json:"isFollowed"`  // 当前用户是否已关注
	CreateTime  string `json:"createTime"`  // 创建时间
}

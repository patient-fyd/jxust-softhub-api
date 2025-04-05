package model

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TopicListInput 获取话题列表输入参数
type TopicListInput struct {
	Page    int    `json:"page" v:"required|min:1#页码不能为空|页码必须大于0"`     // 页码
	Size    int    `json:"size" v:"required|min:1#每页条数不能为空|每页条数必须大于0"` // 每页条数
	Keyword string `json:"keyword,omitempty"`                          // 关键词搜索
	UserId  int    `json:"userId,omitempty"`                           // 用户ID，查询该用户关注的话题
	IsHot   int    `json:"isHot,omitempty"`                            // 是否热门：0-否，1-是
}

// TopicListOutput 获取话题列表输出参数
type TopicListOutput struct {
	List  []TopicListItem `json:"list"`  // 话题列表
	Total int             `json:"total"` // 总数
	Page  int             `json:"page"`  // 页码
	Size  int             `json:"size"`  // 每页条数
}

// TopicListItem 话题列表项
type TopicListItem struct {
	TopicId     int         `json:"topicId"`     // 话题ID
	Name        string      `json:"name"`        // 话题名称
	Description string      `json:"description"` // 话题描述
	Icon        string      `json:"icon"`        // 话题图标
	PostCount   int         `json:"postCount"`   // 帖子数
	IsHot       int         `json:"isHot"`       // 是否热门：0-否，1-是
	IsFollowed  bool        `json:"isFollowed"`  // 当前用户是否已关注
	CreateTime  *gtime.Time `json:"createTime"`  // 创建时间
}

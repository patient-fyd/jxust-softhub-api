package model

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// CircleListInput 获取圈子列表输入参数
type CircleListInput struct {
	Page    int    `json:"page" v:"required|min:1#页码不能为空|页码必须大于0"`     // 页码
	Size    int    `json:"size" v:"required|min:1#每页条数不能为空|每页条数必须大于0"` // 每页条数
	Keyword string `json:"keyword,omitempty"`                          // 关键词搜索
	UserId  int    `json:"userId,omitempty"`                           // 用户ID，查询该用户关注的圈子
	OrderBy string `json:"orderBy,omitempty"`                          // 排序方式：new(最新)、hot(最热)
}

// CircleListOutput 获取圈子列表输出参数
type CircleListOutput struct {
	List  []CircleListItem `json:"list"`  // 圈子列表
	Total int              `json:"total"` // 总数
	Page  int              `json:"page"`  // 页码
	Size  int              `json:"size"`  // 每页条数
}

// CircleListItem 圈子列表项
type CircleListItem struct {
	CircleId    int         `json:"circleId"`    // 圈子ID
	Name        string      `json:"name"`        // 圈子名称
	Description string      `json:"description"` // 圈子描述
	Icon        string      `json:"icon"`        // 圈子图标
	PostCount   int         `json:"postCount"`   // 帖子数
	MemberCount int         `json:"memberCount"` // 成员数
	IsOfficial  int         `json:"isOfficial"`  // 是否官方圈子：0-否，1-是
	IsFollowed  bool        `json:"isFollowed"`  // 当前用户是否已关注
	CreateTime  *gtime.Time `json:"createTime"`  // 创建时间
}

// CircleDetailInput 获取圈子详情输入参数
type CircleDetailInput struct {
	CircleId int `json:"circleId" v:"required#圈子ID不能为空"` // 圈子ID
}

// CircleDetailOutput 获取圈子详情输出参数
type CircleDetailOutput struct {
	CircleId    int         `json:"circleId"`    // 圈子ID
	Name        string      `json:"name"`        // 圈子名称
	Description string      `json:"description"` // 圈子描述
	Icon        string      `json:"icon"`        // 圈子图标
	PostCount   int         `json:"postCount"`   // 帖子数
	MemberCount int         `json:"memberCount"` // 成员数
	CreatorId   int         `json:"creatorId"`   // 创建者ID
	CreatorName string      `json:"creatorName"` // 创建者名称
	IsOfficial  int         `json:"isOfficial"`  // 是否官方圈子：0-否，1-是
	IsFollowed  bool        `json:"isFollowed"`  // 当前用户是否已关注
	CreateTime  *gtime.Time `json:"createTime"`  // 创建时间
	UpdateTime  *gtime.Time `json:"updateTime"`  // 更新时间
}

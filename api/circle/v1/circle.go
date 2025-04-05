package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

// ListReq 获取圈子列表请求
type ListReq struct {
	g.Meta  `path:"/api/circle/v1/list" method:"get" tags:"CircleService" summary:"获取圈子列表"`
	Page    int    `p:"page" v:"required|min:1#页码不能为空|页码必须大于0" dc:"页码"`
	Size    int    `p:"size" v:"required|min:1#每页条数不能为空|每页条数必须大于0" dc:"每页条数"`
	Keyword string `p:"keyword" dc:"关键词搜索"`
	UserId  int    `p:"userId" dc:"用户ID，查询该用户关注的圈子"`
	OrderBy string `p:"orderBy" dc:"排序方式：new(最新)、hot(最热)"`
}

// ListRes 获取圈子列表响应
type ListRes struct {
	List  []CircleItem `json:"list"`  // 圈子列表
	Total int          `json:"total"` // 总数
	Page  int          `json:"page"`  // 页码
	Size  int          `json:"size"`  // 每页条数
}

// CircleItem 圈子列表项
type CircleItem struct {
	CircleId    int    `json:"circleId"`    // 圈子ID
	Name        string `json:"name"`        // 圈子名称
	Description string `json:"description"` // 圈子描述
	Icon        string `json:"icon"`        // 圈子图标
	PostCount   int    `json:"postCount"`   // 帖子数
	MemberCount int    `json:"memberCount"` // 成员数
	IsOfficial  int    `json:"isOfficial"`  // 是否官方圈子：0-否，1-是
	IsFollowed  bool   `json:"isFollowed"`  // 当前用户是否已关注
	CreateTime  string `json:"createTime"`  // 创建时间
}

// DetailReq 获取圈子详情请求
type DetailReq struct {
	g.Meta   `path:"/api/circle/v1/detail" method:"get" tags:"CircleService" summary:"获取圈子详情"`
	CircleId int `p:"circleId" v:"required#圈子ID不能为空" dc:"圈子ID"`
}

// DetailRes 获取圈子详情响应
type DetailRes struct {
	CircleId    int    `json:"circleId"`    // 圈子ID
	Name        string `json:"name"`        // 圈子名称
	Description string `json:"description"` // 圈子描述
	Icon        string `json:"icon"`        // 圈子图标
	PostCount   int    `json:"postCount"`   // 帖子数
	MemberCount int    `json:"memberCount"` // 成员数
	CreatorId   int    `json:"creatorId"`   // 创建者ID
	CreatorName string `json:"creatorName"` // 创建者名称
	IsOfficial  int    `json:"isOfficial"`  // 是否官方圈子：0-否，1-是
	IsFollowed  bool   `json:"isFollowed"`  // 当前用户是否已关注
	CreateTime  string `json:"createTime"`  // 创建时间
	UpdateTime  string `json:"updateTime"`  // 更新时间
}

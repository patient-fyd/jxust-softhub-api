package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

// ListReq 获取帖子列表请求
type ListReq struct {
	g.Meta   `path:"/api/post/v1/list" method:"get" tags:"PostService" summary:"获取帖子列表"`
	Page     int    `p:"page" v:"required|min:1#页码不能为空|页码必须大于0" dc:"页码"`
	Size     int    `p:"size" v:"required|min:1#每页条数不能为空|每页条数必须大于0" dc:"每页条数"`
	Keyword  string `p:"keyword" dc:"关键词搜索"`
	CircleId int    `p:"circleId" dc:"圈子ID"`
	TopicId  int    `p:"topicId" dc:"话题ID"`
	UserId   int    `p:"userId" dc:"用户ID，查询指定用户的帖子"`
	OrderBy  string `p:"orderBy" dc:"排序字段：time(最新)、hot(最热)"`
}

// ListRes 获取帖子列表响应
type ListRes struct {
	List  []PostItem `json:"list"`  // 帖子列表
	Total int        `json:"total"` // 总数
	Page  int        `json:"page"`  // 页码
	Size  int        `json:"size"`  // 每页条数
}

// PostItem 帖子列表项
type PostItem struct {
	PostId       int      `json:"postId"`       // 帖子ID
	UserId       int      `json:"userId"`       // 用户ID
	UserName     string   `json:"userName"`     // 用户名
	UserAvatar   string   `json:"userAvatar"`   // 用户头像
	Content      string   `json:"content"`      // 内容
	Images       []string `json:"images"`       // 图片列表
	CircleId     int      `json:"circleId"`     // 圈子ID
	CircleName   string   `json:"circleName"`   // 圈子名称
	TopicId      int      `json:"topicId"`      // 话题ID
	TopicName    string   `json:"topicName"`    // 话题名称
	ViewCount    int      `json:"viewCount"`    // 浏览量
	LikeCount    int      `json:"likeCount"`    // 点赞数
	CommentCount int      `json:"commentCount"` // 评论数
	ShareCount   int      `json:"shareCount"`   // 分享数
	IsTop        int      `json:"isTop"`        // 是否置顶：0-否，1-是
	IsLiked      bool     `json:"isLiked"`      // 当前用户是否已点赞
	CreateTime   string   `json:"createTime"`   // 创建时间
}

// DetailReq 获取帖子详情请求
type DetailReq struct {
	g.Meta `path:"/api/post/v1/detail" method:"get" tags:"PostService" summary:"获取帖子详情"`
	PostId int `p:"postId" v:"required#帖子ID不能为空" dc:"帖子ID"`
}

// DetailRes 获取帖子详情响应
type DetailRes struct {
	PostId       int      `json:"postId"`       // 帖子ID
	UserId       int      `json:"userId"`       // 用户ID
	UserName     string   `json:"userName"`     // 用户名
	UserAvatar   string   `json:"userAvatar"`   // 用户头像
	Content      string   `json:"content"`      // 内容
	Images       []string `json:"images"`       // 图片列表
	CircleId     int      `json:"circleId"`     // 圈子ID
	CircleName   string   `json:"circleName"`   // 圈子名称
	TopicId      int      `json:"topicId"`      // 话题ID
	TopicName    string   `json:"topicName"`    // 话题名称
	ViewCount    int      `json:"viewCount"`    // 浏览量
	LikeCount    int      `json:"likeCount"`    // 点赞数
	CommentCount int      `json:"commentCount"` // 评论数
	ShareCount   int      `json:"shareCount"`   // 分享数
	IsTop        int      `json:"isTop"`        // 是否置顶：0-否，1-是
	IsLiked      bool     `json:"isLiked"`      // 当前用户是否已点赞
	CreateTime   string   `json:"createTime"`   // 创建时间
	UpdateTime   string   `json:"updateTime"`   // 更新时间
}

// CreateReq 创建帖子请求
type CreateReq struct {
	g.Meta   `path:"/api/post/v1/create" method:"post" tags:"PostService" summary:"创建帖子"`
	Content  string   `p:"content" v:"required#内容不能为空" dc:"内容"`
	Images   []string `p:"images" dc:"图片数组，包含图片URL"`
	CircleId int      `p:"circleId" dc:"圈子ID"`
	TopicId  int      `p:"topicId" dc:"话题ID"`
}

// CreateRes 创建帖子响应
type CreateRes struct {
	PostId int `json:"postId"` // 帖子ID
}

// DeleteReq 删除帖子请求
type DeleteReq struct {
	g.Meta `path:"/api/post/v1/delete" method:"post" tags:"PostService" summary:"删除帖子"`
	PostId int `p:"postId" v:"required#帖子ID不能为空" dc:"帖子ID"`
}

// DeleteRes 删除帖子响应
type DeleteRes struct {
	Success bool `json:"success"` // 是否成功
}

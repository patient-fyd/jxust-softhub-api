package model

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PostListInput 获取帖子列表输入参数
type PostListInput struct {
	Page     int    `json:"page" v:"required|min:1#页码不能为空|页码必须大于0"`     // 页码
	Size     int    `json:"size" v:"required|min:1#每页条数不能为空|每页条数必须大于0"` // 每页条数
	Keyword  string `json:"keyword,omitempty"`                          // 关键词搜索
	CircleId int    `json:"circleId,omitempty"`                         // 圈子ID
	TopicId  int    `json:"topicId,omitempty"`                          // 话题ID
	UserId   int    `json:"userId,omitempty"`                           // 用户ID，查询指定用户的帖子
	OrderBy  string `json:"orderBy,omitempty"`                          // 排序字段：time(最新)、hot(最热)
}

// PostListOutput 获取帖子列表输出参数
type PostListOutput struct {
	List  []PostListItem `json:"list"`  // 帖子列表
	Total int            `json:"total"` // 总数
	Page  int            `json:"page"`  // 页码
	Size  int            `json:"size"`  // 每页条数
}

// PostListItem 帖子列表项
type PostListItem struct {
	PostId       int         `json:"postId"`       // 帖子ID
	UserId       int         `json:"userId"`       // 用户ID
	UserName     string      `json:"userName"`     // 用户名
	UserAvatar   string      `json:"userAvatar"`   // 用户头像
	Content      string      `json:"content"`      // 内容
	Images       []string    `json:"images"`       // 图片列表
	CircleId     int         `json:"circleId"`     // 圈子ID
	CircleName   string      `json:"circleName"`   // 圈子名称
	TopicId      int         `json:"topicId"`      // 话题ID
	TopicName    string      `json:"topicName"`    // 话题名称
	ViewCount    int         `json:"viewCount"`    // 浏览量
	LikeCount    int         `json:"likeCount"`    // 点赞数
	CommentCount int         `json:"commentCount"` // 评论数
	ShareCount   int         `json:"shareCount"`   // 分享数
	IsTop        int         `json:"isTop"`        // 是否置顶：0-否，1-是
	IsLiked      bool        `json:"isLiked"`      // 当前用户是否已点赞
	CreateTime   *gtime.Time `json:"createTime"`   // 创建时间
}

// PostDetailInput 获取帖子详情输入参数
type PostDetailInput struct {
	PostId int `json:"postId" v:"required#帖子ID不能为空"` // 帖子ID
}

// PostDetailOutput 获取帖子详情输出参数
type PostDetailOutput struct {
	PostId       int         `json:"postId"`       // 帖子ID
	UserId       int         `json:"userId"`       // 用户ID
	UserName     string      `json:"userName"`     // 用户名
	UserAvatar   string      `json:"userAvatar"`   // 用户头像
	Content      string      `json:"content"`      // 内容
	Images       []string    `json:"images"`       // 图片列表
	CircleId     int         `json:"circleId"`     // 圈子ID
	CircleName   string      `json:"circleName"`   // 圈子名称
	TopicId      int         `json:"topicId"`      // 话题ID
	TopicName    string      `json:"topicName"`    // 话题名称
	ViewCount    int         `json:"viewCount"`    // 浏览量
	LikeCount    int         `json:"likeCount"`    // 点赞数
	CommentCount int         `json:"commentCount"` // 评论数
	ShareCount   int         `json:"shareCount"`   // 分享数
	IsTop        int         `json:"isTop"`        // 是否置顶：0-否，1-是
	IsLiked      bool        `json:"isLiked"`      // 当前用户是否已点赞
	CreateTime   *gtime.Time `json:"createTime"`   // 创建时间
	UpdateTime   *gtime.Time `json:"updateTime"`   // 更新时间
}

// PostCreateInput 创建帖子输入参数
type PostCreateInput struct {
	Content  string   `json:"content" v:"required#内容不能为空"` // 内容
	Images   []string `json:"images,omitempty"`            // 图片数组，包含图片URL
	CircleId int      `json:"circleId,omitempty"`          // 圈子ID
	TopicId  int      `json:"topicId,omitempty"`           // 话题ID
}

// PostCreateOutput 创建帖子输出参数
type PostCreateOutput struct {
	PostId int `json:"postId"` // 帖子ID
}

// PostDeleteInput 删除帖子输入参数
type PostDeleteInput struct {
	PostId int `json:"postId" v:"required#帖子ID不能为空"` // 帖子ID
}

// PostDeleteOutput 删除帖子输出参数
type PostDeleteOutput struct {
	Success bool `json:"success"` // 是否成功
}

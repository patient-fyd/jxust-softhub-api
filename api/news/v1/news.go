package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

// ListReq 获取新闻列表请求
type ListReq struct {
	g.Meta   `path:"/api/news/v1/list" method:"get" tags:"NewsService" summary:"获取新闻列表"`
	Category string `p:"category" dc:"新闻分类"`
	Page     int    `p:"page" v:"min:0#分页号码错误" dc:"分页号码, 默认1"`
	PageSize int    `p:"pageSize" v:"max:500#每页数量最大500条" dc:"分页数量, 最大500"`
}

// ListRes 获取新闻列表响应
type ListRes struct {
	List     []NewsInfo `json:"list"`     // 新闻列表
	Total    int        `json:"total"`    // 总数
	Page     int        `json:"page"`     // 当前页码
	PageSize int        `json:"pageSize"` // 每页数量
}

// NewsInfo 新闻信息
type NewsInfo struct {
	Id         uint   `json:"id"`         // 新闻ID
	Title      string `json:"title"`      // 新闻标题
	Category   string `json:"category"`   // 新闻分类
	CoverImage string `json:"coverImage"` // 封面图片URL
	ViewCount  uint   `json:"viewCount"`  // 查看次数
	IsTop      uint8  `json:"isTop"`      // 是否置顶：0-否，1-是
	Status     uint8  `json:"status"`     // 状态：0-草稿，1-已发布，2-已下架
	CreatedAt  string `json:"createdAt"`  // 创建时间
	UpdatedAt  string `json:"updatedAt"`  // 更新时间
}

// DetailReq 获取新闻详情请求
type DetailReq struct {
	g.Meta `path:"/api/news/v1/detail/{newsId}" method:"get" tags:"NewsService" summary:"获取新闻详情"`
	NewsId uint `in:"path" v:"min:1#新闻ID必须大于0" dc:"新闻ID"`
}

// DetailRes 获取新闻详情响应
type DetailRes struct {
	Id         uint   `json:"id"`         // 新闻ID
	Title      string `json:"title"`      // 新闻标题
	Content    string `json:"content"`    // 新闻内容
	Category   string `json:"category"`   // 新闻分类
	CoverImage string `json:"coverImage"` // 封面图片URL
	ViewCount  uint   `json:"viewCount"`  // 查看次数
	IsTop      uint8  `json:"isTop"`      // 是否置顶：0-否，1-是
	Status     uint8  `json:"status"`     // 状态：0-草稿，1-已发布，2-已下架
	CreatedAt  string `json:"createdAt"`  // 创建时间
	UpdatedAt  string `json:"updatedAt"`  // 更新时间
}

// CreateReq 创建新闻请求
type CreateReq struct {
	g.Meta     `path:"/api/news/v1/create" method:"post" tags:"NewsService" summary:"创建新闻"`
	Title      string `p:"title" v:"required#新闻标题不能为空"`
	Content    string `p:"content" v:"required#新闻内容不能为空"`
	Category   string `p:"category" v:"required#新闻分类不能为空"`
	CoverImage string `p:"coverImage"`
	IsTop      uint8  `p:"isTop" v:"in:0,1#是否置顶参数错误"`
	Status     uint8  `p:"status" v:"in:0,1,2#状态参数错误"`
}

// CreateRes 创建新闻响应
type CreateRes struct {
	Id uint `json:"id"` // 新闻ID
}

// UpdateReq 更新新闻请求
type UpdateReq struct {
	g.Meta     `path:"/api/news/v1/update/{newsId}" method:"put" tags:"NewsService" summary:"更新新闻"`
	NewsId     uint   `in:"path" v:"min:1#新闻ID必须大于0" dc:"新闻ID"`
	Title      string `p:"title"`
	Content    string `p:"content"`
	Category   string `p:"category"`
	CoverImage string `p:"coverImage"`
	IsTop      uint8  `p:"isTop" v:"in:0,1#是否置顶参数错误"`
	Status     uint8  `p:"status" v:"in:0,1,2#状态参数错误"`
}

// UpdateRes 更新新闻响应
type UpdateRes struct {
	Success bool `json:"success"` // 是否成功
}

// DeleteReq 删除新闻请求
type DeleteReq struct {
	g.Meta `path:"/api/news/v1/delete/{newsId}" method:"delete" tags:"NewsService" summary:"删除新闻"`
	NewsId uint `in:"path" v:"min:1#新闻ID必须大于0" dc:"新闻ID"`
}

// DeleteRes 删除新闻响应
type DeleteRes struct {
	Success bool `json:"success"` // 是否成功
}

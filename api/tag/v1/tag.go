package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

// ListReq 标签列表请求
type ListReq struct {
	g.Meta  `path:"/api/tag/list" method:"get" tags:"TagService" summary:"获取标签列表"`
	TagType string `p:"tag_type" dc:"标签类型，如news、project、resource等"`
	// 分页参数
	PageNum  int `p:"page_num" v:"min:0#分页号码错误" dc:"分页号码, 默认1"`
	PageSize int `p:"page_size" v:"max:500#每页数量最大500条" dc:"分页数量, 最大500"`
}

// ListRes 标签列表响应
type ListRes struct {
	List     []TagInfo `json:"list"`      // 标签列表
	Total    int       `json:"total"`     // 总数
	PageNum  int       `json:"page_num"`  // 分页号码
	PageSize int       `json:"page_size"` // 分页数量
}

// TagInfo 标签信息
type TagInfo struct {
	TagId      uint   `json:"tagId"`      // 标签ID
	TagName    string `json:"tagName"`    // 标签名称
	TagType    string `json:"tagType"`    // 标签类型
	UsageCount int    `json:"usageCount"` // 使用次数
	CreateTime string `json:"createTime"` // 创建时间
}

// CreateReq 创建标签请求
type CreateReq struct {
	g.Meta  `path:"/api/tag/create" method:"post" tags:"TagService" summary:"创建标签"`
	TagName string `p:"tagName" v:"required#标签名称不能为空"`
	TagType string `p:"tagType" v:"required#标签类型不能为空"`
}

// CreateRes 创建标签响应
type CreateRes struct {
	TagId uint `json:"tagId"` // 标签ID
}

// UpdateReq 更新标签请求
type UpdateReq struct {
	g.Meta  `path:"/api/tag/update" method:"put" tags:"TagService" summary:"更新标签信息"`
	TagId   uint   `p:"tagId" v:"required#标签ID不能为空"`
	TagName string `p:"tagName" v:"required#标签名称不能为空"`
}

// UpdateRes 更新标签响应
type UpdateRes struct{}

// DeleteReq 删除标签请求
type DeleteReq struct {
	g.Meta `path:"/api/tag/delete" method:"delete" tags:"TagService" summary:"删除标签"`
	TagId  uint `p:"tagId" v:"required#标签ID不能为空"`
}

// DeleteRes 删除标签响应
type DeleteRes struct{}

// ContentTagsReq 获取内容关联标签请求
type ContentTagsReq struct {
	g.Meta      `path:"/api/tag/content" method:"get" tags:"TagService" summary:"获取内容关联的标签"`
	ContentType string `p:"contentType" v:"required#内容类型不能为空"`
	ContentId   uint   `p:"contentId" v:"required#内容ID不能为空"`
}

// ContentTagsRes 获取内容关联标签响应
type ContentTagsRes struct {
	Tags []TagInfo `json:"tags"` // 标签列表
}

// SetContentTagsReq 设置内容关联标签请求
type SetContentTagsReq struct {
	g.Meta      `path:"/api/tag/content/set" method:"post" tags:"TagService" summary:"设置内容关联的标签"`
	ContentType string `p:"contentType" v:"required#内容类型不能为空"`
	ContentId   uint   `p:"contentId" v:"required#内容ID不能为空"`
	TagIds      []uint `p:"tagIds" v:"required#标签ID列表不能为空"`
}

// SetContentTagsRes 设置内容关联标签响应
type SetContentTagsRes struct{}

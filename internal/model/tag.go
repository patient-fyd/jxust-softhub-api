package model

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TagListInput 标签列表查询输入
type TagListInput struct {
	TagType  string // 标签类型，如news、project、resource等
	PageNum  int    // 分页号码
	PageSize int    // 分页数量
}

// TagListOutput 标签列表查询输出
type TagListOutput struct {
	List     []*TagInfo // 标签列表
	Total    int        // 总数
	PageNum  int        // 分页号码
	PageSize int        // 分页数量
}

// TagInfo 标签信息
type TagInfo struct {
	TagId      uint        // 标签ID
	TagName    string      // 标签名称
	TagType    string      // 标签类型
	UsageCount int         // 使用次数
	CreateTime *gtime.Time // 创建时间
}

// TagCreateInput 创建标签输入
type TagCreateInput struct {
	TagName string // 标签名称
	TagType string // 标签类型
}

// TagCreateOutput 创建标签输出
type TagCreateOutput struct {
	TagId uint // 标签ID
}

// TagUpdateInput 更新标签输入
type TagUpdateInput struct {
	TagId   uint   // 标签ID
	TagName string // 标签名称
}

// TagUpdateOutput 更新标签输出
type TagUpdateOutput struct{}

// TagDeleteInput 删除标签输入
type TagDeleteInput struct {
	TagId uint // 标签ID
}

// TagDeleteOutput 删除标签输出
type TagDeleteOutput struct{}

// ContentTagsInput 获取内容关联标签输入
type ContentTagsInput struct {
	ContentType string // 内容类型
	ContentId   uint   // 内容ID
}

// ContentTagsOutput 获取内容关联标签输出
type ContentTagsOutput struct {
	Tags []*TagInfo // 标签列表
}

// SetContentTagsInput 设置内容关联标签输入
type SetContentTagsInput struct {
	ContentType string // 内容类型
	ContentId   uint   // 内容ID
	TagIds      []uint // 标签ID列表
}

// SetContentTagsOutput 设置内容关联标签输出
type SetContentTagsOutput struct{}

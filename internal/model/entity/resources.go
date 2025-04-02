// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Resources is the golang structure for table resources.
type Resources struct {
	ResourceId  uint        `json:"resource_id" description:"资源ID，主键，自增"`
	Title       string      `json:"title"       description:"资源标题"`
	Description string      `json:"description" description:"资源描述信息"`
	Category    string      `json:"category"    description:"资源分类，如编程语言、算法、工具等"`
	FilePath    string      `json:"file_path"   description:"资源文件存储路径或 URL"`
	UploaderId  uint        `json:"uploader_id" description:"上传者ID，关联 users 表"`
	Downloads   int         `json:"downloads"   description:"下载次数"`
	CreateTime  *gtime.Time `json:"create_time" description:"记录创建时间"`
}

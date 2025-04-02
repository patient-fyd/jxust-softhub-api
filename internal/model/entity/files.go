// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Files is the golang structure for table files.
type Files struct {
	FileId      uint        `json:"file_id"      description:"文件ID，主键，自增"`
	FileName    string      `json:"file_name"    description:"文件名称"`
	FileSize    uint64      `json:"file_size"    description:"文件大小（字节）"`
	FileType    string      `json:"file_type"    description:"文件类型/MIME类型"`
	FilePath    string      `json:"file_path"    description:"文件存储路径"`
	UploaderId  uint        `json:"uploader_id"  description:"上传者ID，关联users表"`
	UploadTime  *gtime.Time `json:"upload_time"  description:"上传时间"`
	Md5Hash     string      `json:"md_5_hash"    description:"文件MD5哈希值，用于去重"`
	RelatedType string      `json:"related_type" description:"关联类型，如news、activity、resource等"`
	RelatedId   uint        `json:"related_id"   description:"关联ID，指向具体的内容ID"`
}

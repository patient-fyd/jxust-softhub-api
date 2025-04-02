// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Files is the golang structure of table files for DAO operations like Where/Data.
type Files struct {
	g.Meta      `orm:"table:files, do:true"`
	FileId      interface{} // 文件ID，主键，自增
	FileName    interface{} // 文件名称
	FileSize    interface{} // 文件大小（字节）
	FileType    interface{} // 文件类型/MIME类型
	FilePath    interface{} // 文件存储路径
	UploaderId  interface{} // 上传者ID，关联users表
	UploadTime  *gtime.Time // 上传时间
	Md5Hash     interface{} // 文件MD5哈希值，用于去重
	RelatedType interface{} // 关联类型，如news、activity、resource等
	RelatedId   interface{} // 关联ID，指向具体的内容ID
}

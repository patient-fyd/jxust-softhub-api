// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Resources is the golang structure of table resources for DAO operations like Where/Data.
type Resources struct {
	g.Meta      `orm:"table:resources, do:true"`
	ResourceId  interface{} // 资源ID，主键，自增
	Title       interface{} // 资源标题
	Description interface{} // 资源描述信息
	Category    interface{} // 资源分类，如编程语言、算法、工具等
	FilePath    interface{} // 资源文件存储路径或 URL
	UploaderId  interface{} // 上传者ID，关联 users 表
	Downloads   interface{} // 下载次数
	CreateTime  *gtime.Time // 记录创建时间
}

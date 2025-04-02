// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SystemConfigs is the golang structure of table system_configs for DAO operations like Where/Data.
type SystemConfigs struct {
	g.Meta      `orm:"table:system_configs, do:true"`
	ConfigKey   interface{} // 配置键名
	ConfigValue interface{} // 配置值
	Description interface{} // 配置说明
	CreateTime  *gtime.Time // 创建时间
	UpdateTime  *gtime.Time // 更新时间
}

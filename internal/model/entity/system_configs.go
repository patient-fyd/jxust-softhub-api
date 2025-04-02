// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SystemConfigs is the golang structure for table system_configs.
type SystemConfigs struct {
	ConfigKey   string      `json:"config_key"   description:"配置键名"`
	ConfigValue string      `json:"config_value" description:"配置值"`
	Description string      `json:"description"  description:"配置说明"`
	CreateTime  *gtime.Time `json:"create_time"  description:"创建时间"`
	UpdateTime  *gtime.Time `json:"update_time"  description:"更新时间"`
}

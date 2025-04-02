package model

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ConfigGetInput 获取系统配置输入
type ConfigGetInput struct {
	ConfigKey string // 配置键名
}

// ConfigGetOutput 获取系统配置输出
type ConfigGetOutput struct {
	Config *ConfigInfo // 配置信息
}

// ConfigListInput 配置列表查询输入
type ConfigListInput struct {
	PageNum  int // 分页号码
	PageSize int // 分页数量
}

// ConfigListOutput 配置列表查询输出
type ConfigListOutput struct {
	List     []*ConfigInfo // 配置列表
	Total    int           // 总数
	PageNum  int           // 分页号码
	PageSize int           // 分页数量
}

// ConfigInfo 配置信息
type ConfigInfo struct {
	ConfigKey   string      // 配置键名
	ConfigValue string      // 配置值
	Description string      // 配置说明
	CreateTime  *gtime.Time // 创建时间
	UpdateTime  *gtime.Time // 更新时间
}

// ConfigSetInput 设置系统配置输入
type ConfigSetInput struct {
	ConfigKey   string // 配置键名
	ConfigValue string // 配置值
	Description string // 配置说明
}

// ConfigSetOutput 设置系统配置输出
type ConfigSetOutput struct{}

// ConfigDeleteInput 删除系统配置输入
type ConfigDeleteInput struct {
	ConfigKey string // 配置键名
}

// ConfigDeleteOutput 删除系统配置输出
type ConfigDeleteOutput struct{}

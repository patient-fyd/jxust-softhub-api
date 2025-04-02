package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

// GetReq 获取系统配置请求
type GetReq struct {
	g.Meta    `path:"/api/config/get" method:"get" tags:"ConfigService" summary:"获取系统配置"`
	ConfigKey string `p:"configKey" v:"required#配置键名不能为空"`
}

// GetRes 获取系统配置响应
type GetRes struct {
	Config ConfigInfo `json:"config"` // 配置信息
}

// ListReq 配置列表请求
type ListReq struct {
	g.Meta `path:"/api/config/list" method:"get" tags:"ConfigService" summary:"获取配置列表"`
	// 分页参数
	PageNum  int `p:"page_num" v:"min:0#分页号码错误" dc:"分页号码, 默认1"`
	PageSize int `p:"page_size" v:"max:500#每页数量最大500条" dc:"分页数量, 最大500"`
}

// ListRes 配置列表响应
type ListRes struct {
	List     []ConfigInfo `json:"list"`      // 配置列表
	Total    int          `json:"total"`     // 总数
	PageNum  int          `json:"page_num"`  // 分页号码
	PageSize int          `json:"page_size"` // 分页数量
}

// ConfigInfo 配置信息
type ConfigInfo struct {
	ConfigKey   string `json:"configKey"`   // 配置键名
	ConfigValue string `json:"configValue"` // 配置值
	Description string `json:"description"` // 配置说明
	CreateTime  string `json:"createTime"`  // 创建时间
	UpdateTime  string `json:"updateTime"`  // 更新时间
}

// SetReq 设置系统配置请求
type SetReq struct {
	g.Meta      `path:"/api/config/set" method:"post" tags:"ConfigService" summary:"设置系统配置"`
	ConfigKey   string `p:"configKey" v:"required#配置键名不能为空"`
	ConfigValue string `p:"configValue" v:"required#配置值不能为空"`
	Description string `p:"description" dc:"配置说明"`
}

// SetRes 设置系统配置响应
type SetRes struct{}

// DeleteReq 删除系统配置请求
type DeleteReq struct {
	g.Meta    `path:"/api/config/delete" method:"delete" tags:"ConfigService" summary:"删除系统配置"`
	ConfigKey string `p:"configKey" v:"required#配置键名不能为空"`
}

// DeleteRes 删除系统配置响应
type DeleteRes struct{}

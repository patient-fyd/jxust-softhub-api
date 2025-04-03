package user

import (
	"github.com/patient-fyd/jxust-softhub-api/api/user"
)

// ControllerV1 v1版本控制器
type ControllerV1 struct{}

// NewV1 创建v1版本控制器实例
func NewV1() user.IUserV1 {
	return &ControllerV1{}
}

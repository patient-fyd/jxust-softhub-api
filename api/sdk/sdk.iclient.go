// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package sdk

import (
	"github.com/patient-fyd/jxust-softhub-api/api/auth"
)

type IClient interface {
	// 在此处添加新的服务接口
	AuthV1() auth.IAuthV1
}

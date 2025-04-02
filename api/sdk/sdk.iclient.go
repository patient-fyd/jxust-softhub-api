// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package sdk

import (
	"github.com/patient-fyd/jxust-softhub-api/api/auth"
	"github.com/patient-fyd/jxust-softhub-api/api/config"
	"github.com/patient-fyd/jxust-softhub-api/api/file"
	"github.com/patient-fyd/jxust-softhub-api/api/join"
	"github.com/patient-fyd/jxust-softhub-api/api/member"
	"github.com/patient-fyd/jxust-softhub-api/api/stat"
	"github.com/patient-fyd/jxust-softhub-api/api/tag"
)

type IClient interface {
	// 在此处添加新的服务接口
	AuthV1() auth.IAuthV1
	ConfigV1() config.IConfigV1
	FileV1() file.IFileV1
	JoinV1() join.IJoinV1
	MemberV1() member.IMemberV1
	StatV1() stat.IStatV1
	TagV1() tag.ITagV1
}

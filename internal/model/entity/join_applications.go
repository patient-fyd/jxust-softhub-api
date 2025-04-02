// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// JoinApplications is the golang structure for table join_applications.
type JoinApplications struct {
	ApplicationId    uint        `json:"application_id"    description:"申请ID，主键，自增"`
	Name             string      `json:"name"              description:"申请人姓名"`
	StudentId        string      `json:"student_id"        description:"学号"`
	Grade            string      `json:"grade"             description:"年级，如2020级"`
	College          string      `json:"college"           description:"学院"`
	Major            string      `json:"major"             description:"专业"`
	Phone            string      `json:"phone"             description:"联系电话"`
	Email            string      `json:"email"             description:"邮箱"`
	Reason           string      `json:"reason"            description:"申请理由"`
	Skills           string      `json:"skills"            description:"技能介绍"`
	ExpectDepartment string      `json:"expect_department" description:"期望加入的部门"`
	Status           int         `json:"status"            description:"申请状态：0-待审核，1-通过，2-拒绝"`
	ReviewerId       uint        `json:"reviewer_id"       description:"审核人ID，关联users表"`
	ReviewComment    string      `json:"review_comment"    description:"审核意见"`
	CreateTime       *gtime.Time `json:"create_time"       description:"申请时间"`
	UpdateTime       *gtime.Time `json:"update_time"       description:"最后更新时间"`
}

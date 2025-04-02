// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// JoinApplications is the golang structure of table join_applications for DAO operations like Where/Data.
type JoinApplications struct {
	g.Meta           `orm:"table:join_applications, do:true"`
	ApplicationId    interface{} // 申请ID，主键，自增
	Name             interface{} // 申请人姓名
	StudentId        interface{} // 学号
	Grade            interface{} // 年级，如2020级
	College          interface{} // 学院
	Major            interface{} // 专业
	Phone            interface{} // 联系电话
	Email            interface{} // 邮箱
	Reason           interface{} // 申请理由
	Skills           interface{} // 技能介绍
	ExpectDepartment interface{} // 期望加入的部门
	Status           interface{} // 申请状态：0-待审核，1-通过，2-拒绝
	ReviewerId       interface{} // 审核人ID，关联users表
	ReviewComment    interface{} // 审核意见
	CreateTime       *gtime.Time // 申请时间
	UpdateTime       *gtime.Time // 最后更新时间
}

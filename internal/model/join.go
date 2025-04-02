package model

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// JoinApplyInput 提交入会申请输入
type JoinApplyInput struct {
	Name             string // 申请人姓名
	StudentId        string // 学号
	Grade            string // 年级
	College          string // 学院
	Major            string // 专业
	Phone            string // 联系电话
	Email            string // 邮箱
	Reason           string // 申请理由
	Skills           string // 技能介绍
	ExpectDepartment string // 期望加入的部门
}

// JoinApplyOutput 提交入会申请输出
type JoinApplyOutput struct {
	ApplicationId uint // 申请ID
}

// JoinListInput 入会申请列表查询输入
type JoinListInput struct {
	Status           int    // 申请状态筛选
	Grade            string // 年级筛选
	ExpectDepartment string // 期望部门筛选
	PageNum          int    // 分页号码
	PageSize         int    // 分页数量
}

// JoinListOutput 入会申请列表查询输出
type JoinListOutput struct {
	List     []*JoinApplicationInfo // 申请列表
	Total    int                    // 总数
	PageNum  int                    // 分页号码
	PageSize int                    // 分页数量
}

// JoinApplicationInfo 申请信息
type JoinApplicationInfo struct {
	ApplicationId    uint        // 申请ID
	Name             string      // 申请人姓名
	StudentId        string      // 学号
	Grade            string      // 年级
	College          string      // 学院
	Major            string      // 专业
	Phone            string      // 联系电话
	Email            string      // 邮箱
	Reason           string      // 申请理由
	Skills           string      // 技能介绍
	ExpectDepartment string      // 期望加入的部门
	Status           int         // 申请状态
	ReviewerId       uint        // 审核人ID
	ReviewerName     string      // 审核人姓名
	ReviewComment    string      // 审核意见
	CreateTime       *gtime.Time // 申请时间
	UpdateTime       *gtime.Time // 最后更新时间
}

// JoinDetailInput 申请详情查询输入
type JoinDetailInput struct {
	ApplicationId uint // 申请ID
}

// JoinDetailOutput 申请详情查询输出
type JoinDetailOutput struct {
	*JoinApplicationInfo
}

// JoinReviewInput 审核入会申请输入
type JoinReviewInput struct {
	ApplicationId uint   // 申请ID
	ReviewerId    uint   // 审核人ID
	Status        int    // 审核状态：1-通过，2-拒绝
	ReviewComment string // 审核意见
}

// JoinReviewOutput 审核入会申请输出
type JoinReviewOutput struct{}

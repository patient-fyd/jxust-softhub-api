package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

// ApplyReq 提交入会申请请求
type ApplyReq struct {
	g.Meta           `path:"/api/join/v1/apply" method:"post" tags:"JoinService" summary:"提交入会申请"`
	Name             string `p:"name" v:"required#申请人姓名不能为空"`
	StudentId        string `p:"studentId" v:"required#学号不能为空"`
	Grade            string `p:"grade" v:"required#年级不能为空"`
	College          string `p:"college" v:"required#学院不能为空"`
	Major            string `p:"major" v:"required#专业不能为空"`
	Phone            string `p:"phone" v:"required|phone#联系电话不能为空|联系电话格式不正确"`
	Email            string `p:"email" v:"email#邮箱格式不正确"`
	Reason           string `p:"reason" v:"required#申请理由不能为空"`
	Skills           string `p:"skills"`
	ExpectDepartment string `p:"expectDepartment"`
}

// ApplyRes 提交入会申请响应
type ApplyRes struct {
	ApplicationId uint `json:"applicationId"` // 申请ID
}

// ListReq 入会申请列表请求
type ListReq struct {
	g.Meta           `path:"/api/join/v1/list" method:"get" tags:"JoinService" summary:"获取入会申请列表"`
	Status           int    `p:"status" dc:"申请状态筛选"`
	Grade            string `p:"grade" dc:"年级筛选"`
	ExpectDepartment string `p:"expectDepartment" dc:"期望部门筛选"`
	// 分页参数
	PageNum  int `p:"page_num" v:"min:0#分页号码错误" dc:"分页号码, 默认1"`
	PageSize int `p:"page_size" v:"max:500#每页数量最大500条" dc:"分页数量, 最大500"`
}

// ListRes 入会申请列表响应
type ListRes struct {
	List     []ApplicationInfo `json:"list"`      // 申请列表
	Total    int               `json:"total"`     // 总数
	PageNum  int               `json:"page_num"`  // 分页号码
	PageSize int               `json:"page_size"` // 分页数量
}

// ApplicationInfo 申请信息
type ApplicationInfo struct {
	ApplicationId    uint   `json:"applicationId"`    // 申请ID
	Name             string `json:"name"`             // 申请人姓名
	StudentId        string `json:"studentId"`        // 学号
	Grade            string `json:"grade"`            // 年级
	College          string `json:"college"`          // 学院
	Major            string `json:"major"`            // 专业
	Phone            string `json:"phone"`            // 联系电话
	Email            string `json:"email"`            // 邮箱
	Reason           string `json:"reason"`           // 申请理由
	Skills           string `json:"skills"`           // 技能介绍
	ExpectDepartment string `json:"expectDepartment"` // 期望加入的部门
	Status           int    `json:"status"`           // 申请状态
	ReviewerId       uint   `json:"reviewerId"`       // 审核人ID
	ReviewerName     string `json:"reviewerName"`     // 审核人姓名
	ReviewComment    string `json:"reviewComment"`    // 审核意见
	CreateTime       string `json:"createTime"`       // 申请时间
	UpdateTime       string `json:"updateTime"`       // 最后更新时间
}

// DetailReq 申请详情请求
type DetailReq struct {
	g.Meta        `path:"/api/join/v1/detail" method:"get" tags:"JoinService" summary:"获取入会申请详情"`
	ApplicationId uint `p:"applicationId" v:"required#申请ID不能为空"`
}

// DetailRes 申请详情响应
type DetailRes struct {
	ApplicationInfo
}

// ReviewReq 审核入会申请请求
type ReviewReq struct {
	g.Meta        `path:"/api/join/v1/review" method:"post" tags:"JoinService" summary:"审核入会申请"`
	ApplicationId uint   `p:"applicationId" v:"required#申请ID不能为空"`
	Status        int    `p:"status" v:"required|in:1,2#审核状态不能为空|审核状态只能为通过或拒绝"`
	ReviewComment string `p:"reviewComment"`
}

// ReviewRes 审核入会申请响应
type ReviewRes struct{}

// ApplicationListReq 获取入会申请列表请求（符合API文档路径）
type ApplicationListReq struct {
	g.Meta           `path:"/api/join/v1/application-list" method:"get" tags:"JoinService" summary:"获取入会申请列表"`
	Status           int    `p:"status" dc:"申请状态筛选"`
	Grade            string `p:"grade" dc:"年级筛选"`
	ExpectDepartment string `p:"expectDepartment" dc:"期望部门筛选"`
	// 分页参数
	Page     int `p:"page" v:"min:0#分页号码错误" dc:"分页号码, 默认1"`
	PageSize int `p:"pageSize" v:"max:500#每页数量最大500条" dc:"分页数量, 最大500"`
}

// ApplicationListRes 获取入会申请列表响应
type ApplicationListRes struct {
	List     []ApplicationInfo `json:"list"`     // 申请列表
	Total    int               `json:"total"`    // 总数
	Page     int               `json:"page"`     // 当前页码
	PageSize int               `json:"pageSize"` // 每页数量
}

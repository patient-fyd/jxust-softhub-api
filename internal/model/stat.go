package model

// VisitStatInput 网站访问统计数据查询输入
type VisitStatInput struct {
	StartDate string // 开始日期，格式：YYYY-MM-DD
	EndDate   string // 结束日期，格式：YYYY-MM-DD
}

// VisitStatOutput 网站访问统计数据查询输出
type VisitStatOutput struct {
	List []*VisitStatInfo // 统计数据列表
}

// VisitStatInfo 访问统计数据
type VisitStatInfo struct {
	VisitDate     string  // 访问日期
	PageView      int     // 页面浏览量
	UniqueVisitor int     // 独立访客数
	NewVisitor    int     // 新访客数
	AvgStayTime   int     // 平均停留时间（秒）
	BounceRate    float64 // 跳出率（百分比）
}

// ActivityStatInput 活动统计数据查询输入
type ActivityStatInput struct {
	StartDate string // 开始日期，格式：YYYY-MM-DD
	EndDate   string // 结束日期，格式：YYYY-MM-DD
}

// ActivityStatOutput 活动统计数据查询输出
type ActivityStatOutput struct {
	TotalActivity     int                 // 活动总数
	TotalRegistration int                 // 报名总人数
	AvgRegistration   float64             // 平均每个活动报名人数
	ActivityList      []*ActivityStatInfo // 活动统计列表
}

// ActivityStatInfo 活动统计数据
type ActivityStatInfo struct {
	ActivityId      uint   // 活动ID
	Title           string // 活动标题
	StartTime       string // 活动开始时间
	EndTime         string // 活动结束时间
	RegistrationNum int    // 报名人数
	ApprovedNum     int    // 审核通过人数
	RejectedNum     int    // 拒绝人数
	PendingNum      int    // 待审核人数
}

// UserStatInput 用户统计数据查询输入
type UserStatInput struct{}

// UserStatOutput 用户统计数据查询输出
type UserStatOutput struct {
	TotalUser            int              // 用户总数
	NewUserLast30Days    int              // 最近30天新增用户数
	ActiveUserLast30Days int              // 最近30天活跃用户数
	UserByRole           map[string]int   // 按角色统计用户数
	MemberStats          *MemberStatsInfo // 协会成员统计
}

// MemberStatsInfo 成员统计信息
type MemberStatsInfo struct {
	TotalMember  int            // 成员总数
	CoreMember   int            // 核心成员数
	ByDepartment map[string]int // 按部门统计
	ByGrade      map[string]int // 按年级统计
}

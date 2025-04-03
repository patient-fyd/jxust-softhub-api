package model

// UserDetailInfo 用户详细信息
type UserDetailInfo struct {
	UserId    uint   // 用户ID
	UserName  string // 用户名
	Name      string // 真实姓名
	Email     string // 电子邮箱
	Phone     string // 联系电话
	Avatar    string // 头像
	RoleId    uint   // 角色ID
	RoleName  string // 角色名称
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
}

// UserListInput 用户列表查询输入
type UserListInput struct {
	PageNum  int    // 页码
	PageSize int    // 每页数量
	Keyword  string // 搜索关键词
}

// UserListOutput 用户列表查询输出
type UserListOutput struct {
	List     []*UserDetailInfo // 用户列表
	Total    int               // 总数
	PageNum  int               // 页码
	PageSize int               // 每页数量
}

// 用户角色分配输入参数
type UserAssignRoleInput struct {
	UserId uint // 用户ID
	RoleId uint // 角色ID
}

// 用户角色分配输出参数
type UserAssignRoleOutput struct {
	Success bool // 是否成功
}

package model

// NewsListInput 获取新闻列表的输入参数
type NewsListInput struct {
	Category string // 新闻分类
	Page     int    // 页码
	PageSize int    // 每页数量
}

// NewsListOutput 获取新闻列表的输出参数
type NewsListOutput struct {
	List     []NewsInfo // 新闻列表
	Total    int        // 总数
	Page     int        // 当前页码
	PageSize int        // 每页数量
}

// NewsInfo 新闻信息
type NewsInfo struct {
	Id         uint   // 新闻ID
	Title      string // 新闻标题
	Category   string // 新闻分类
	CoverImage string // 封面图片URL
	ViewCount  uint   // 查看次数
	IsTop      uint8  // 是否置顶：0-否，1-是
	Status     uint8  // 状态：0-草稿，1-已发布，2-已下架
	CreatedAt  string // 创建时间
	UpdatedAt  string // 更新时间
}

// NewsDetailInput 获取新闻详情的输入参数
type NewsDetailInput struct {
	NewsId uint // 新闻ID
}

// NewsDetailOutput 获取新闻详情的输出参数
type NewsDetailOutput struct {
	Id         uint   // 新闻ID
	Title      string // 新闻标题
	Content    string // 新闻内容
	Category   string // 新闻分类
	CoverImage string // 封面图片URL
	ViewCount  uint   // 查看次数
	IsTop      uint8  // 是否置顶：0-否，1-是
	Status     uint8  // 状态：0-草稿，1-已发布，2-已下架
	CreatedAt  string // 创建时间
	UpdatedAt  string // 更新时间
}

// NewsCreateInput 创建新闻的输入参数
type NewsCreateInput struct {
	Title      string // 新闻标题
	Content    string // 新闻内容
	Category   string // 新闻分类
	CoverImage string // 封面图片URL
	IsTop      uint8  // 是否置顶：0-否，1-是
	Status     uint8  // 状态：0-草稿，1-已发布，2-已下架
}

// NewsCreateOutput 创建新闻的输出参数
type NewsCreateOutput struct {
	Id uint // 新闻ID
}

// NewsUpdateInput 更新新闻的输入参数
type NewsUpdateInput struct {
	NewsId     uint   // 新闻ID
	Title      string // 新闻标题
	Content    string // 新闻内容
	Category   string // 新闻分类
	CoverImage string // 封面图片URL
	IsTop      uint8  // 是否置顶：0-否，1-是
	Status     uint8  // 状态：0-草稿，1-已发布，2-已下架
}

// NewsUpdateOutput 更新新闻的输出参数
type NewsUpdateOutput struct {
	Success bool // 是否成功
}

// NewsDeleteInput 删除新闻的输入参数
type NewsDeleteInput struct {
	NewsId uint // 新闻ID
}

// NewsDeleteOutput 删除新闻的输出参数
type NewsDeleteOutput struct {
	Success bool // 是否成功
}

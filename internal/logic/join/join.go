package join

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"

	"github.com/patient-fyd/jxust-softhub-api/internal/dao"
	"github.com/patient-fyd/jxust-softhub-api/internal/model"
	"github.com/patient-fyd/jxust-softhub-api/internal/service"
)

// sJoin 入会申请服务实现
type sJoin struct{}

// New 创建入会申请服务
func New() *sJoin {
	return &sJoin{}
}

// init 初始化
func init() {
	service.RegisterJoin(New())
}

// Apply 提交入会申请
func (s *sJoin) Apply(ctx context.Context, in model.JoinApplyInput) (*model.JoinApplyOutput, error) {
	// 获取当前登录用户
	userId := service.Auth().GetLoginUserId(ctx)
	if userId == 0 {
		return nil, gerror.New("用户未登录")
	}

	// 检查是否已经提交过申请
	count, err := dao.JoinApplications.Ctx(ctx).
		Where("status", 0). // 待审核状态
		Count()
	if err != nil {
		g.Log().Error(ctx, "查询申请记录失败:", err)
		return nil, gerror.Wrap(err, "检查申请记录失败")
	}
	if count > 0 {
		return nil, gerror.New("已有待审核的申请，请勿重复提交")
	}

	// 创建申请记录
	data := g.Map{
		"name":              in.Name,
		"student_id":        in.StudentId,
		"grade":             in.Grade,
		"college":           in.College,
		"major":             in.Major,
		"phone":             in.Phone,
		"email":             in.Email,
		"reason":            in.Reason,
		"skills":            in.Skills,
		"expect_department": in.ExpectDepartment,
		"status":            0, // 默认待审核
	}
	result, err := dao.JoinApplications.Ctx(ctx).Data(data).Insert()
	if err != nil {
		g.Log().Error(ctx, "创建申请记录失败:", err)
		return nil, gerror.Wrap(err, "提交申请失败")
	}

	// 获取插入的ID
	lastInsertId, err := result.LastInsertId()
	if err != nil {
		g.Log().Error(ctx, "获取申请ID失败:", err)
		return nil, gerror.Wrap(err, "提交申请失败")
	}

	return &model.JoinApplyOutput{
		ApplicationId: uint(lastInsertId),
	}, nil
}

// List 获取入会申请列表
func (s *sJoin) List(ctx context.Context, in model.JoinListInput) (*model.JoinListOutput, error) {
	// 获取当前登录用户，检查是否为管理员
	userId := service.Auth().GetLoginUserId(ctx)
	if userId == 0 {
		return nil, gerror.New("用户未登录")
	}

	// 查询用户角色
	var user struct {
		RoleId uint
	}
	err := dao.Users.Ctx(ctx).
		Fields("roleId").
		Where("userId", userId).
		Scan(&user)
	if err != nil {
		return nil, gerror.Wrap(err, "查询用户角色失败")
	}

	// 判断是否为管理员角色（假设roleId=1为管理员）
	if user.RoleId != 1 {
		return nil, gerror.New("权限不足，只有管理员可以查看申请列表")
	}

	// 构建查询条件
	m := dao.JoinApplications.Ctx(ctx)

	// 根据状态筛选
	if in.Status > 0 {
		m = m.Where("status", in.Status)
	}

	// 根据年级筛选
	if in.Grade != "" {
		m = m.Where("grade", in.Grade)
	}

	// 根据期望部门筛选
	if in.ExpectDepartment != "" {
		m = m.Where("expect_department", in.ExpectDepartment)
	}

	// 统计总数
	total, err := m.Count()
	if err != nil {
		return nil, gerror.Wrap(err, "统计申请总数失败")
	}

	// 处理分页参数
	if in.PageNum <= 0 {
		in.PageNum = 1
	}
	if in.PageSize <= 0 {
		in.PageSize = 10
	}

	// 查询列表数据
	var applications []struct {
		Id               uint        `json:"id"`
		UserId           uint        `json:"userId"`
		Name             string      `json:"name"`
		StudentId        string      `json:"studentId"`
		Grade            string      `json:"grade"`
		College          string      `json:"college"`
		Major            string      `json:"major"`
		Phone            string      `json:"phone"`
		Email            string      `json:"email"`
		Reason           string      `json:"reason"`
		Skills           string      `json:"skills"`
		ExpectDepartment string      `json:"expectDepartment"`
		Status           int         `json:"status"`
		ReviewerId       uint        `json:"reviewerId"`
		ReviewerName     string      `json:"reviewerName"`
		ReviewComment    string      `json:"reviewComment"`
		CreateTime       *gtime.Time `json:"createTime"`
		UpdateTime       *gtime.Time `json:"updateTime"`
	}

	err = m.
		Order("applicationId DESC").
		Page(in.PageNum, in.PageSize).
		Scan(&applications)
	if err != nil {
		return nil, gerror.Wrap(err, "查询申请列表失败")
	}

	// 实例化返回对象
	result := &model.JoinListOutput{
		Total:    total,
		PageNum:  in.PageNum,
		PageSize: in.PageSize,
		List:     make([]*model.JoinApplicationInfo, 0, len(applications)),
	}

	// 填充数据
	for _, app := range applications {
		info := &model.JoinApplicationInfo{
			ApplicationId:    app.Id,
			Name:             app.Name,
			StudentId:        app.StudentId,
			Grade:            app.Grade,
			College:          app.College,
			Major:            app.Major,
			Phone:            app.Phone,
			Email:            app.Email,
			Reason:           app.Reason,
			Skills:           app.Skills,
			ExpectDepartment: app.ExpectDepartment,
			Status:           app.Status,
			ReviewerId:       app.ReviewerId,
			ReviewerName:     app.ReviewerName,
			ReviewComment:    app.ReviewComment,
			CreateTime:       app.CreateTime,
			UpdateTime:       app.UpdateTime,
		}
		result.List = append(result.List, info)
	}

	return result, nil
}

// ApplicationList 获取入会申请列表（API文档中的路径）
func (s *sJoin) ApplicationList(ctx context.Context, in model.JoinApplicationListInput) (*model.JoinApplicationListOutput, error) {
	// 转换为 List 方法的输入参数
	listInput := model.JoinListInput{
		Status:           in.Status,
		Grade:            in.Grade,
		ExpectDepartment: in.ExpectDepartment,
		PageNum:          in.Page,
		PageSize:         in.PageSize,
	}

	// 调用 List 方法
	listResult, err := s.List(ctx, listInput)
	if err != nil {
		return nil, err
	}

	// 创建返回结果
	result := &model.JoinApplicationListOutput{
		Total:    listResult.Total,
		Page:     listResult.PageNum,
		PageSize: listResult.PageSize,
		List:     make([]*model.JoinApplicationInfo, 0, len(listResult.List)),
	}

	// 复制数据
	for _, item := range listResult.List {
		result.List = append(result.List, item)
	}

	return result, nil
}

// Detail 获取入会申请详情
func (s *sJoin) Detail(ctx context.Context, in model.JoinDetailInput) (*model.JoinDetailOutput, error) {
	// TODO: 实现获取入会申请详情功能
	return nil, gerror.New("功能未实现")
}

// Review 审核入会申请
func (s *sJoin) Review(ctx context.Context, in model.JoinReviewInput) (*model.JoinReviewOutput, error) {
	// TODO: 实现审核入会申请功能
	return nil, gerror.New("功能未实现")
}

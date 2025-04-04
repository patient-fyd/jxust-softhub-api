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

	// 获取用户信息
	var userInfo struct {
		Name  string
		Phone string
	}
	err := dao.Users.Ctx(ctx).
		Fields("name", "phone").
		Where("userId", userId).
		Scan(&userInfo)
	if err != nil {
		return nil, gerror.Wrap(err, "查询用户信息失败")
	}

	// 检查是否已经提交过申请（只检查当前用户的申请）
	count, err := dao.JoinApplications.Ctx(ctx).
		Where("status", 0). // 待审核状态
		Where("name", userInfo.Name).
		Where("phone", userInfo.Phone).
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
		"name":             in.Name,
		"studentId":        in.StudentId,
		"grade":            in.Grade,
		"college":          in.College,
		"major":            in.Major,
		"phone":            in.Phone,
		"email":            in.Email,
		"reason":           in.Reason,
		"skills":           in.Skills,
		"expectDepartment": in.ExpectDepartment,
		"status":           0, // 默认待审核
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

	// 构建查询条件
	m := dao.JoinApplications.Ctx(ctx)

	// 判断角色权限
	if !(user.RoleId == 1 || user.RoleId == 2 || user.RoleId == 3) { // 不是管理员、内容管理员、会长
		// 非管理员只能查看自己提交的申请，根据姓名和电话查询
		// 获取用户信息
		var userInfo struct {
			Name  string
			Phone string
		}
		err := dao.Users.Ctx(ctx).
			Fields("name", "phone").
			Where("userId", userId).
			Scan(&userInfo)
		if err != nil {
			return nil, gerror.Wrap(err, "查询用户信息失败")
		}

		if userInfo.Name != "" && userInfo.Phone != "" {
			m = m.Where("name", userInfo.Name).Where("phone", userInfo.Phone)
		} else {
			// 如果用户未设置姓名和电话，则无法关联查询申请记录
			return &model.JoinListOutput{
				Total:    0,
				PageNum:  in.PageNum,
				PageSize: in.PageSize,
				List:     make([]*model.JoinApplicationInfo, 0),
			}, nil
		}
	}

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
		m = m.Where("expectDepartment", in.ExpectDepartment)
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
		ApplicationId    uint        `json:"applicationId"` // 改成 ApplicationId 而不是 Id
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

	// 获取审核人姓名
	reviewerMap := make(map[uint]string)
	if len(applications) > 0 {
		// 收集所有审核人ID
		reviewerIds := make([]uint, 0)
		for _, app := range applications {
			if app.ReviewerId > 0 {
				reviewerIds = append(reviewerIds, app.ReviewerId)
			}
		}

		// 如果有审核人，查询他们的姓名
		if len(reviewerIds) > 0 {
			var reviewers []struct {
				UserId uint
				Name   string
			}
			err = dao.Users.Ctx(ctx).
				Fields("userId", "name").
				WhereIn("userId", reviewerIds).
				Scan(&reviewers)
			if err == nil {
				for _, r := range reviewers {
					reviewerMap[r.UserId] = r.Name
				}
			}
		}
	}

	// 填充数据
	for _, app := range applications {
		info := &model.JoinApplicationInfo{
			ApplicationId:    app.ApplicationId,
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
			ReviewerName:     reviewerMap[app.ReviewerId], // 使用查询到的审核人姓名
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
	// 获取当前登录用户
	userId := service.Auth().GetLoginUserId(ctx)
	if userId == 0 {
		return nil, gerror.New("用户未登录")
	}

	// 查询用户角色
	var user struct {
		RoleId uint
		Name   string
		Phone  string
	}
	err := dao.Users.Ctx(ctx).
		Fields("roleId", "name", "phone").
		Where("userId", userId).
		Scan(&user)
	if err != nil {
		return nil, gerror.Wrap(err, "查询用户信息失败")
	}

	// 查询申请详情
	var application struct {
		ApplicationId    uint        `json:"applicationId"`
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
		ReviewComment    string      `json:"reviewComment"`
		CreateTime       *gtime.Time `json:"createTime"`
		UpdateTime       *gtime.Time `json:"updateTime"`
	}

	m := dao.JoinApplications.Ctx(ctx).Where("applicationId", in.ApplicationId)
	err = m.Scan(&application)
	if err != nil {
		return nil, gerror.Wrap(err, "查询申请详情失败")
	}

	// 如果没有找到记录
	if application.ApplicationId == 0 {
		return nil, gerror.New("申请记录不存在")
	}

	// 判断权限：只有管理员或本人可以查看申请详情
	isAdmin := user.RoleId == 1 || user.RoleId == 2 || user.RoleId == 3
	isSelf := user.Name == application.Name && user.Phone == application.Phone

	if !isAdmin && !isSelf {
		return nil, gerror.New("权限不足，不能查看该申请详情")
	}

	// 查询审核人信息
	reviewerName := ""
	if application.ReviewerId > 0 {
		var reviewer struct {
			Name string
		}
		err = dao.Users.Ctx(ctx).
			Fields("name").
			Where("userId", application.ReviewerId).
			Scan(&reviewer)
		if err == nil && reviewer.Name != "" {
			reviewerName = reviewer.Name
		}
	}

	// 构建返回数据
	applicationInfo := &model.JoinApplicationInfo{
		ApplicationId:    application.ApplicationId,
		Name:             application.Name,
		StudentId:        application.StudentId,
		Grade:            application.Grade,
		College:          application.College,
		Major:            application.Major,
		Phone:            application.Phone,
		Email:            application.Email,
		Reason:           application.Reason,
		Skills:           application.Skills,
		ExpectDepartment: application.ExpectDepartment,
		Status:           application.Status,
		ReviewerId:       application.ReviewerId,
		ReviewerName:     reviewerName,
		ReviewComment:    application.ReviewComment,
		CreateTime:       application.CreateTime,
		UpdateTime:       application.UpdateTime,
	}

	return &model.JoinDetailOutput{
		JoinApplicationInfo: applicationInfo,
	}, nil
}

// Review 审核入会申请
func (s *sJoin) Review(ctx context.Context, in model.JoinReviewInput) (*model.JoinReviewOutput, error) {
	// 获取当前登录用户
	userId := service.Auth().GetLoginUserId(ctx)
	if userId == 0 {
		return nil, gerror.New("用户未登录")
	}

	// 查询用户角色
	var user struct {
		RoleId uint
		Name   string
	}
	err := dao.Users.Ctx(ctx).
		Fields("roleId", "name").
		Where("userId", userId).
		Scan(&user)
	if err != nil {
		return nil, gerror.Wrap(err, "查询用户信息失败")
	}

	// 判断权限：只有管理员可以审核申请
	if !(user.RoleId == 1 || user.RoleId == 2 || user.RoleId == 3) {
		return nil, gerror.New("权限不足，只有管理员可以审核申请")
	}

	// 验证申请ID是否存在
	var application struct {
		ApplicationId uint
		Status        int
	}
	err = dao.JoinApplications.Ctx(ctx).
		Fields("applicationId", "status").
		Where("applicationId", in.ApplicationId).
		Scan(&application)
	if err != nil {
		return nil, gerror.Wrap(err, "查询申请记录失败")
	}

	if application.ApplicationId == 0 {
		return nil, gerror.New("申请记录不存在")
	}

	// 判断申请状态是否为待审核
	if application.Status != 0 {
		return nil, gerror.New("该申请已审核，不能重复审核")
	}

	// 更新申请状态
	data := g.Map{
		"status":        in.Status,
		"reviewerId":    userId,
		"reviewComment": in.ReviewComment,
		"updateTime":    gtime.Now(),
	}

	_, err = dao.JoinApplications.Ctx(ctx).
		Data(data).
		Where("applicationId", in.ApplicationId).
		Update()
	if err != nil {
		return nil, gerror.Wrap(err, "更新申请状态失败")
	}

	// 如果审核通过，可以在这里执行其他操作，如创建会员记录等
	if in.Status == 1 {
		// 查询申请详情，获取申请人信息
		var applicationInfo struct {
			Name             string
			StudentId        string
			Grade            string
			College          string
			Major            string
			Phone            string
			Email            string
			Skills           string
			ExpectDepartment string
		}
		err = dao.JoinApplications.Ctx(ctx).
			Where("applicationId", in.ApplicationId).
			Scan(&applicationInfo)
		if err == nil {
			// 记录审计日志
			_, _ = dao.AuditLogs.Ctx(ctx).Data(g.Map{
				"userId": userId,
				"action": "approve_join_application",
				"description": g.Map{
					"applicationId": in.ApplicationId,
					"name":          applicationInfo.Name,
					"studentId":     applicationInfo.StudentId,
				},
			}).Insert()

			// 查询对应的用户并将角色改为会员(假设会员的roleId=5)
			// 根据申请中的姓名和电话查找对应用户
			var userInfo struct {
				UserId uint
				RoleId uint
			}
			err = dao.Users.Ctx(ctx).
				Fields("userId", "roleId").
				Where("name", applicationInfo.Name).
				Where("phone", applicationInfo.Phone).
				Scan(&userInfo)

			if err == nil && userInfo.UserId > 0 {
				// 如果用户当前是普通用户(roleId=4)，则将其升级为会员(roleId=5)
				if userInfo.RoleId == 4 {
					_, updateErr := dao.Users.Ctx(ctx).
						Data(g.Map{"roleId": 5}). // 假设roleId=5是会员角色
						Where("userId", userInfo.UserId).
						Update()

					if updateErr != nil {
						g.Log().Error(ctx, "更新用户角色失败:", updateErr)
					} else {
						g.Log().Info(ctx, "用户已从普通用户升级为会员:", userInfo.UserId)
					}
				}
			}
		}
	} else if in.Status == 2 {
		// 如果拒绝，记录审计日志
		_, _ = dao.AuditLogs.Ctx(ctx).Data(g.Map{
			"userId": userId,
			"action": "reject_join_application",
			"description": g.Map{
				"applicationId": in.ApplicationId,
				"reviewComment": in.ReviewComment,
			},
		}).Insert()
	}

	return &model.JoinReviewOutput{}, nil
}

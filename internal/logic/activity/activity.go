package activity

import (
	"context"
	"time"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"

	"github.com/patient-fyd/jxust-softhub-api/internal/codes"
	"github.com/patient-fyd/jxust-softhub-api/internal/model"
	"github.com/patient-fyd/jxust-softhub-api/internal/service"
)

// 服务实现
type sActivity struct{}

// 创建活动服务实例
func New() *sActivity {
	return &sActivity{}
}

// 注册服务
func init() {
	service.RegisterActivity(New())
}

// 创建活动
func (s *sActivity) Create(ctx context.Context, in model.ActivityCreateInput) (*model.ActivityCreateOutput, error) {
	// 检查结束时间是否晚于开始时间
	if in.EndTime.Before(in.StartTime) {
		return nil, gerror.New("结束时间不能早于开始时间")
	}

	// 判断活动状态
	status := model.ActivityStatusNotStarted
	now := time.Now()
	if now.After(in.StartTime) && now.Before(in.EndTime) {
		status = model.ActivityStatusOngoing
	} else if now.After(in.EndTime) {
		status = model.ActivityStatusEnded
	}

	// 创建活动
	r, err := g.DB().Insert(ctx, "activities", g.Map{
		"title":           in.Title,
		"description":     in.Description,
		"startTime":       in.StartTime,
		"endTime":         in.EndTime,
		"location":        in.Location,
		"maxParticipants": in.MaxParticipants,
		"status":          status,
		"createTime":      gtime.Now(),
	})
	if err != nil {
		return nil, gerror.Wrap(err, "创建活动失败")
	}

	// 获取新创建的活动ID
	lastInsertId, err := r.LastInsertId()
	if err != nil {
		return nil, gerror.Wrap(err, "获取活动ID失败")
	}

	return &model.ActivityCreateOutput{
		ActivityId: uint(lastInsertId),
	}, nil
}

// 更新活动
func (s *sActivity) Update(ctx context.Context, in model.ActivityUpdateInput) (*model.ActivityUpdateOutput, error) {
	// 检查活动是否存在
	activity, err := s.getActivityById(ctx, in.ActivityId)
	if err != nil {
		return nil, err
	}
	if activity == nil {
		return nil, gerror.NewCode(codes.CodeNotFound, "活动不存在")
	}

	// 构建更新数据
	data := g.Map{}

	if in.Title != nil {
		data["title"] = *in.Title
	}
	if in.Description != nil {
		data["description"] = *in.Description
	}
	if in.StartTime != nil {
		data["startTime"] = *in.StartTime
	}
	if in.EndTime != nil {
		data["endTime"] = *in.EndTime
	}
	if in.Location != nil {
		data["location"] = *in.Location
	}
	if in.MaxParticipants != nil {
		data["maxParticipants"] = *in.MaxParticipants
	}
	if in.Status != nil {
		data["status"] = *in.Status
	} else {
		// 如果没有明确指定状态，根据时间自动判断
		if in.StartTime != nil || in.EndTime != nil {
			startTime := activity.StartTime
			endTime := activity.EndTime

			if in.StartTime != nil {
				startTime = *in.StartTime
			}
			if in.EndTime != nil {
				endTime = *in.EndTime
			}

			now := time.Now()
			if now.Before(startTime) {
				data["status"] = model.ActivityStatusNotStarted
			} else if now.After(startTime) && now.Before(endTime) {
				data["status"] = model.ActivityStatusOngoing
			} else {
				data["status"] = model.ActivityStatusEnded
			}
		}
	}

	// 检查开始时间和结束时间是否合理
	if in.StartTime != nil && in.EndTime != nil && in.EndTime.Before(*in.StartTime) {
		return nil, gerror.New("结束时间不能早于开始时间")
	} else if in.StartTime != nil && in.EndTime == nil && activity.EndTime.Before(*in.StartTime) {
		return nil, gerror.New("结束时间不能早于开始时间")
	} else if in.StartTime == nil && in.EndTime != nil && in.EndTime.Before(activity.StartTime) {
		return nil, gerror.New("结束时间不能早于开始时间")
	}

	// 如果没有需要更新的字段，直接返回
	if len(data) == 0 {
		return &model.ActivityUpdateOutput{ActivityId: in.ActivityId}, nil
	}

	// 更新活动
	_, err = g.DB().Update(ctx, "activities", data, "activityId", in.ActivityId)
	if err != nil {
		return nil, gerror.Wrap(err, "更新活动失败")
	}

	return &model.ActivityUpdateOutput{
		ActivityId: in.ActivityId,
	}, nil
}

// 删除活动
func (s *sActivity) Delete(ctx context.Context, in model.ActivityDeleteInput) error {
	// 检查活动是否存在
	activity, err := s.getActivityById(ctx, in.ActivityId)
	if err != nil {
		return err
	}
	if activity == nil {
		return gerror.NewCode(codes.CodeNotFound, "活动不存在")
	}

	// 删除活动
	_, err = g.DB().Delete(ctx, "activities", "activityId", in.ActivityId)
	if err != nil {
		return gerror.Wrap(err, "删除活动失败")
	}

	return nil
}

// 获取活动列表
func (s *sActivity) List(ctx context.Context, in model.ActivityListInput) (*model.ActivityListOutput, error) {
	// 设置默认分页参数
	if in.Page <= 0 {
		in.Page = 1
	}
	if in.PageSize <= 0 {
		in.PageSize = 10
	}
	if in.PageSize > 100 {
		in.PageSize = 100
	}

	// 构建查询条件
	m := g.DB().Model("activities")
	if in.Status != nil {
		m = m.Where("status", *in.Status)
	}

	// 获取总数
	count, err := m.Count()
	if err != nil {
		return nil, gerror.Wrap(err, "获取活动总数失败")
	}

	// 构建查询
	m = m.Order("createTime DESC").Page(in.Page, in.PageSize)

	// 直接查询到结构体切片中
	var result []model.ActivityInfo
	if err = m.Scan(&result); err != nil {
		return nil, gerror.Wrap(err, "查询活动列表失败")
	}

	return &model.ActivityListOutput{
		List:     result,
		Total:    count,
		Page:     in.Page,
		PageSize: in.PageSize,
	}, nil
}

// 获取活动详情
func (s *sActivity) Detail(ctx context.Context, in model.ActivityDetailInput) (*model.ActivityDetailOutput, error) {
	// 查询活动
	activity, err := s.getActivityById(ctx, in.ActivityId)
	if err != nil {
		return nil, err
	}
	if activity == nil {
		return nil, gerror.NewCode(codes.CodeNotFound, "活动不存在")
	}

	return &model.ActivityDetailOutput{
		Activity: *activity,
	}, nil
}

// 活动报名
func (s *sActivity) Register(ctx context.Context, in model.ActivityRegisterInput) (*model.ActivityRegisterOutput, error) {
	// 检查活动是否存在
	activity, err := s.getActivityById(ctx, in.ActivityId)
	if err != nil {
		return nil, err
	}
	if activity == nil {
		return nil, gerror.NewCode(codes.CodeNotFound, "活动不存在")
	}

	// 检查活动是否已结束
	if activity.Status == model.ActivityStatusEnded {
		return nil, gerror.New("活动已结束，无法报名")
	}

	// 检查报名人数是否已满
	count, err := g.DB().Model("activity_registrations").
		Where("activityId", in.ActivityId).
		Where("status", model.RegistrationStatusApproved).
		Count()
	if err != nil {
		return nil, gerror.Wrap(err, "检查报名人数失败")
	}

	if activity.MaxParticipants > 0 && count >= activity.MaxParticipants {
		return nil, gerror.New("活动报名人数已满")
	}

	// 如果有用户ID，检查是否已经报名
	if in.UserId > 0 {
		exists, err := g.DB().Model("activity_registrations").
			Where("activityId", in.ActivityId).
			Where("userId", in.UserId).
			Count()
		if err != nil {
			return nil, gerror.Wrap(err, "检查是否重复报名失败")
		}
		if exists > 0 {
			return nil, gerror.New("您已经报名过该活动")
		}
	}

	// 创建报名记录
	r, err := g.DB().Insert(ctx, "activity_registrations", g.Map{
		"activityId": in.ActivityId,
		"userId":     in.UserId,
		"name":       in.Name,
		"studentId":  in.StudentId,
		"contact":    in.Contact,
		"status":     model.RegistrationStatusPending,
		"createTime": gtime.Now(),
	})
	if err != nil {
		return nil, gerror.Wrap(err, "创建报名记录失败")
	}

	// 获取新创建的报名ID
	lastInsertId, err := r.LastInsertId()
	if err != nil {
		return nil, gerror.Wrap(err, "获取报名ID失败")
	}

	return &model.ActivityRegisterOutput{
		RegistrationId: uint(lastInsertId),
	}, nil
}

// 获取报名列表
func (s *sActivity) RegisterList(ctx context.Context, in model.RegistrationListInput) (*model.RegistrationListOutput, error) {
	// 检查活动是否存在
	activity, err := s.getActivityById(ctx, in.ActivityId)
	if err != nil {
		return nil, err
	}
	if activity == nil {
		return nil, gerror.NewCode(codes.CodeNotFound, "活动不存在")
	}

	// 设置默认分页参数
	if in.Page <= 0 {
		in.Page = 1
	}
	if in.PageSize <= 0 {
		in.PageSize = 10
	}
	if in.PageSize > 100 {
		in.PageSize = 100
	}

	// 构建查询条件
	m := g.DB().Model("activity_registrations").Where("activityId", in.ActivityId)
	if in.Status != nil {
		m = m.Where("status", *in.Status)
	}

	// 获取总数
	count, err := m.Count()
	if err != nil {
		return nil, gerror.Wrap(err, "获取报名总数失败")
	}

	// 查询数据
	m = m.Order("createTime DESC").Page(in.Page, in.PageSize)

	// 直接查询到结构体切片中
	var result []model.RegistrationInfo
	if err = m.Scan(&result); err != nil {
		return nil, gerror.Wrap(err, "查询报名列表失败")
	}

	return &model.RegistrationListOutput{
		List:     result,
		Total:    count,
		Page:     in.Page,
		PageSize: in.PageSize,
	}, nil
}

// 审核报名
func (s *sActivity) ReviewRegistration(ctx context.Context, in model.RegistrationReviewInput) (*model.RegistrationReviewOutput, error) {
	// 检查报名记录是否存在
	registration, err := s.getRegistrationById(ctx, in.RegistrationId)
	if err != nil {
		return nil, err
	}
	if registration == nil {
		return nil, gerror.NewCode(codes.CodeNotFound, "报名记录不存在")
	}

	// 检查活动是否存在
	activity, err := s.getActivityById(ctx, registration.ActivityId)
	if err != nil {
		return nil, err
	}
	if activity == nil {
		return nil, gerror.NewCode(codes.CodeNotFound, "活动不存在")
	}

	// 检查活动是否已结束
	if activity.Status == model.ActivityStatusEnded {
		return nil, gerror.New("活动已结束，无法审核报名")
	}

	// 如果要通过报名，检查报名人数是否已满
	if in.Status == model.RegistrationStatusApproved {
		count, err := g.DB().Model("activity_registrations").
			Where("activityId", registration.ActivityId).
			Where("status", model.RegistrationStatusApproved).
			Count()
		if err != nil {
			return nil, gerror.Wrap(err, "检查报名人数失败")
		}

		if activity.MaxParticipants > 0 && count >= activity.MaxParticipants {
			return nil, gerror.New("活动报名人数已满")
		}
	}

	// 更新报名状态
	data := g.Map{
		"status":     in.Status,
		"updateTime": gtime.Now(),
	}

	_, err = g.DB().Update(ctx, "activity_registrations", data, "registrationId", in.RegistrationId)
	if err != nil {
		return nil, gerror.Wrap(err, "更新报名状态失败")
	}

	return &model.RegistrationReviewOutput{
		RegistrationId: in.RegistrationId,
	}, nil
}

// 根据ID获取活动信息(内部使用)
func (s *sActivity) getActivityById(ctx context.Context, id uint) (*model.ActivityInfo, error) {
	var activity model.ActivityInfo
	err := g.DB().Model("activities").Where("activityId", id).Scan(&activity)
	if err != nil {
		return nil, gerror.Wrap(err, "查询活动信息失败")
	}
	if activity.ActivityId == 0 {
		return nil, nil
	}
	return &activity, nil
}

// 根据ID获取报名记录(内部使用)
func (s *sActivity) getRegistrationById(ctx context.Context, id uint) (*model.RegistrationInfo, error) {
	var registration model.RegistrationInfo
	err := g.DB().Model("activity_registrations").Where("registrationId", id).Scan(&registration)
	if err != nil {
		return nil, gerror.Wrap(err, "查询报名信息失败")
	}
	if registration.RegistrationId == 0 {
		return nil, nil
	}
	return &registration, nil
}

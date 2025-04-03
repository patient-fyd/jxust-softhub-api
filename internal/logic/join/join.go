package join

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"

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
	// TODO: 实现获取入会申请列表功能
	return nil, gerror.New("功能未实现")
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

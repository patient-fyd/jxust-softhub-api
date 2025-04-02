package member

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"

	"github.com/patient-fyd/jxust-softhub-api/internal/codes"
	"github.com/patient-fyd/jxust-softhub-api/internal/dao"
	"github.com/patient-fyd/jxust-softhub-api/internal/model"
	"github.com/patient-fyd/jxust-softhub-api/internal/model/do"
	"github.com/patient-fyd/jxust-softhub-api/internal/model/entity"
	"github.com/patient-fyd/jxust-softhub-api/internal/service"
)

// sMember 成员管理服务实现
type sMember struct{}

// New 创建成员管理服务
func New() *sMember {
	return &sMember{}
}

// init 初始化
func init() {
	service.RegisterMember(New())
}

// List 获取成员列表
func (s *sMember) List(ctx context.Context, in model.MemberListInput) (*model.MemberListOutput, error) {
	// 基本的输入参数处理
	if in.PageNum <= 0 {
		in.PageNum = 1
	}
	if in.PageSize <= 0 {
		in.PageSize = 10
	}
	if in.PageSize > 500 {
		in.PageSize = 500
	}

	// 构建查询条件
	model := dao.Members.Ctx(ctx)
	if in.Department != "" {
		model = model.Where("department", in.Department)
	}
	if in.Grade != "" {
		model = model.Where("grade", in.Grade)
	}
	if in.IsCore > 0 {
		model = model.Where("is_core", in.IsCore)
	}
	if in.Status > 0 {
		model = model.Where("status", in.Status)
	}

	// 查询总数
	total, err := model.Count()
	if err != nil {
		g.Log().Error(ctx, "Query members count error:", err)
		return nil, gerror.Wrap(err, "获取成员总数失败")
	}

	if total == 0 {
		return &model.MemberListOutput{
			Total:    0,
			PageNum:  in.PageNum,
			PageSize: in.PageSize,
			List:     []*model.MemberInfo{},
		}, nil
	}

	// 查询分页数据
	var members []*entity.Members
	err = model.Page(in.PageNum, in.PageSize).OrderDesc("create_time").Scan(&members)
	if err != nil {
		g.Log().Error(ctx, "Query members error:", err)
		return nil, gerror.Wrap(err, "获取成员列表失败")
	}

	// 构建返回数据
	var list []*model.MemberInfo
	for _, member := range members {
		// 查询关联的用户信息
		var user *entity.Users
		err = dao.Users.Ctx(ctx).Where("user_id", member.UserId).Scan(&user)
		if err != nil {
			g.Log().Error(ctx, "Query user error:", err)
			continue
		}

		list = append(list, &model.MemberInfo{
			MemberId:     member.MemberId,
			UserId:       member.UserId,
			UserName:     user.UserName,
			Name:         user.Name,
			Avatar:       user.Avatar,
			Grade:        member.Grade,
			JoinYear:     member.JoinYear,
			Department:   member.Department,
			Position:     member.Position,
			Skills:       member.Skills,
			Introduction: member.Introduction,
			IsCore:       member.IsCore,
			Status:       member.Status,
			CreateTime:   member.CreateTime,
			UpdateTime:   member.UpdateTime,
		})
	}

	return &model.MemberListOutput{
		List:     list,
		Total:    total,
		PageNum:  in.PageNum,
		PageSize: in.PageSize,
	}, nil
}

// Detail 获取成员详情
func (s *sMember) Detail(ctx context.Context, in model.MemberDetailInput) (*model.MemberDetailOutput, error) {
	// 查询成员信息
	var member *entity.Members
	err := dao.Members.Ctx(ctx).Where("member_id", in.MemberId).Scan(&member)
	if err != nil {
		g.Log().Error(ctx, "Query member error:", err)
		return nil, gerror.Wrap(err, "获取成员信息失败")
	}

	if member == nil {
		return nil, gerror.NewCode(codes.CodeNotFound, "成员不存在")
	}

	// 查询关联的用户信息
	var user *entity.Users
	err = dao.Users.Ctx(ctx).Where("user_id", member.UserId).Scan(&user)
	if err != nil {
		g.Log().Error(ctx, "Query user error:", err)
		return nil, gerror.Wrap(err, "获取用户信息失败")
	}

	return &model.MemberDetailOutput{
		MemberInfo: &model.MemberInfo{
			MemberId:     member.MemberId,
			UserId:       member.UserId,
			UserName:     user.UserName,
			Name:         user.Name,
			Avatar:       user.Avatar,
			Grade:        member.Grade,
			JoinYear:     member.JoinYear,
			Department:   member.Department,
			Position:     member.Position,
			Skills:       member.Skills,
			Introduction: member.Introduction,
			IsCore:       member.IsCore,
			Status:       member.Status,
			CreateTime:   member.CreateTime,
			UpdateTime:   member.UpdateTime,
		},
	}, nil
}

// Create 创建成员
func (s *sMember) Create(ctx context.Context, in model.MemberCreateInput) (*model.MemberCreateOutput, error) {
	// 检查用户是否存在
	count, err := dao.Users.Ctx(ctx).Where("user_id", in.UserId).Count()
	if err != nil {
		g.Log().Error(ctx, "Check user error:", err)
		return nil, gerror.Wrap(err, "检查用户失败")
	}
	if count == 0 {
		return nil, gerror.NewCode(codes.CodeValidationFailed, "用户不存在")
	}

	// 检查是否已经是成员
	count, err = dao.Members.Ctx(ctx).Where("user_id", in.UserId).Count()
	if err != nil {
		g.Log().Error(ctx, "Check member error:", err)
		return nil, gerror.Wrap(err, "检查成员失败")
	}
	if count > 0 {
		return nil, gerror.NewCode(codes.CodeValidationFailed, "该用户已是协会成员")
	}

	// 创建成员
	member := do.Members{
		UserId:       in.UserId,
		Grade:        in.Grade,
		JoinYear:     in.JoinYear,
		Department:   in.Department,
		Position:     in.Position,
		Skills:       in.Skills,
		Introduction: in.Introduction,
		IsCore:       in.IsCore,
		Status:       in.Status,
	}

	// 插入数据
	result, err := dao.Members.Ctx(ctx).Data(member).Insert()
	if err != nil {
		g.Log().Error(ctx, "Create member error:", err)
		return nil, gerror.Wrap(err, "创建成员失败")
	}

	// 获取成员ID
	memberId, err := result.LastInsertId()
	if err != nil {
		g.Log().Error(ctx, "Get last insert id error:", err)
		return nil, gerror.Wrap(err, "获取成员ID失败")
	}

	return &model.MemberCreateOutput{
		MemberId: uint(memberId),
	}, nil
}

// Update 更新成员信息
func (s *sMember) Update(ctx context.Context, in model.MemberUpdateInput) (*model.MemberUpdateOutput, error) {
	// 检查成员是否存在
	count, err := dao.Members.Ctx(ctx).Where("member_id", in.MemberId).Count()
	if err != nil {
		g.Log().Error(ctx, "Check member error:", err)
		return nil, gerror.Wrap(err, "检查成员失败")
	}
	if count == 0 {
		return nil, gerror.NewCode(codes.CodeNotFound, "成员不存在")
	}

	// 更新成员信息
	_, err = dao.Members.Ctx(ctx).
		Data(do.Members{
			Grade:        in.Grade,
			Department:   in.Department,
			Position:     in.Position,
			Skills:       in.Skills,
			Introduction: in.Introduction,
			IsCore:       in.IsCore,
			Status:       in.Status,
		}).
		Where("member_id", in.MemberId).
		Update()
	if err != nil {
		g.Log().Error(ctx, "Update member error:", err)
		return nil, gerror.Wrap(err, "更新成员信息失败")
	}

	return &model.MemberUpdateOutput{}, nil
}

// Delete 删除成员
func (s *sMember) Delete(ctx context.Context, in model.MemberDeleteInput) (*model.MemberDeleteOutput, error) {
	// 检查成员是否存在
	count, err := dao.Members.Ctx(ctx).Where("member_id", in.MemberId).Count()
	if err != nil {
		g.Log().Error(ctx, "Check member error:", err)
		return nil, gerror.Wrap(err, "检查成员失败")
	}
	if count == 0 {
		return nil, gerror.NewCode(codes.CodeNotFound, "成员不存在")
	}

	// 删除成员
	_, err = dao.Members.Ctx(ctx).Where("member_id", in.MemberId).Delete()
	if err != nil {
		g.Log().Error(ctx, "Delete member error:", err)
		return nil, gerror.Wrap(err, "删除成员失败")
	}

	return &model.MemberDeleteOutput{}, nil
}

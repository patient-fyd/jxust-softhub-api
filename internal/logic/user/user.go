package user

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"

	"github.com/patient-fyd/jxust-softhub-api/internal/codes"
	"github.com/patient-fyd/jxust-softhub-api/internal/dao"
	"github.com/patient-fyd/jxust-softhub-api/internal/model"
	"github.com/patient-fyd/jxust-softhub-api/internal/model/entity"
	"github.com/patient-fyd/jxust-softhub-api/internal/service"
)

// sUser 用户服务实现
type sUser struct{}

// New 创建用户服务
func New() *sUser {
	return &sUser{}
}

// init 初始化
func init() {
	service.RegisterUser(New())
}

// GetUserInfo 获取用户详细信息
func (s *sUser) GetUserInfo(ctx context.Context, userId uint) (*model.UserDetailInfo, error) {
	var user *entity.Users
	err := dao.Users.Ctx(ctx).Where("userId", userId).Scan(&user)
	if err != nil {
		g.Log().Error(ctx, "Query user error:", err)
		return nil, gerror.Wrap(err, "获取用户信息失败")
	}

	if user == nil {
		return nil, gerror.New("用户不存在")
	}

	// 获取角色名称
	var role *entity.Roles
	err = dao.Roles.Ctx(ctx).Where("roleId", user.RoleId).Scan(&role)
	if err != nil {
		g.Log().Error(ctx, "Query role error:", err)
		return nil, gerror.Wrap(err, "获取角色信息失败")
	}

	roleName := ""
	if role != nil {
		roleName = role.RoleName
	}

	return &model.UserDetailInfo{
		UserId:    user.UserId,
		UserName:  user.UserName,
		Name:      user.Name,
		Email:     user.Email,
		Phone:     user.Phone,
		Avatar:    user.Avatar,
		RoleId:    user.RoleId,
		RoleName:  roleName,
		CreatedAt: user.CreateTime.String(),
		UpdatedAt: user.UpdateTime.String(),
	}, nil
}

// GetUserList 获取用户列表
func (s *sUser) GetUserList(ctx context.Context, in model.UserListInput) (*model.UserListOutput, error) {
	m := dao.Users.Ctx(ctx)

	// 搜索条件
	if in.Keyword != "" {
		m = m.WhereLike("user_name", "%"+in.Keyword+"%").
			WhereOr("name", "%"+in.Keyword+"%").
			WhereOr("email", "%"+in.Keyword+"%").
			WhereOr("phone", "%"+in.Keyword+"%")
	}

	// 获取总数
	total, err := m.Count()
	if err != nil {
		g.Log().Error(ctx, "Count users error:", err)
		return nil, gerror.Wrap(err, "获取用户总数失败")
	}

	// 分页
	if in.PageNum <= 0 {
		in.PageNum = 1
	}
	if in.PageSize <= 0 {
		in.PageSize = 10
	}
	m = m.Page(in.PageNum, in.PageSize)

	// 查询数据
	var users []*entity.Users
	err = m.Scan(&users)
	if err != nil {
		g.Log().Error(ctx, "Query users error:", err)
		return nil, gerror.Wrap(err, "查询用户列表失败")
	}

	// 查询角色信息
	var roleMap = make(map[uint]string)
	var roleIds []uint
	for _, user := range users {
		roleIds = append(roleIds, user.RoleId)
	}
	if len(roleIds) > 0 {
		var roles []*entity.Roles
		err = dao.Roles.Ctx(ctx).WhereIn("roleId", roleIds).Scan(&roles)
		if err != nil {
			g.Log().Error(ctx, "Query roles error:", err)
			return nil, gerror.Wrap(err, "查询角色列表失败")
		}
		for _, role := range roles {
			roleMap[role.RoleId] = role.RoleName
		}
	}

	// 组装数据
	var list []*model.UserDetailInfo
	for _, user := range users {
		list = append(list, &model.UserDetailInfo{
			UserId:    user.UserId,
			UserName:  user.UserName,
			Name:      user.Name,
			Email:     user.Email,
			Phone:     user.Phone,
			Avatar:    user.Avatar,
			RoleId:    user.RoleId,
			RoleName:  roleMap[user.RoleId],
			CreatedAt: user.CreateTime.String(),
			UpdatedAt: user.UpdateTime.String(),
		})
	}

	return &model.UserListOutput{
		List:     list,
		Total:    int(total),
		PageNum:  in.PageNum,
		PageSize: in.PageSize,
	}, nil
}

// AssignRole 分配用户角色
func (s *sUser) AssignRole(ctx context.Context, in model.UserAssignRoleInput) (*model.UserAssignRoleOutput, error) {
	// 获取当前登录用户
	loginUserId := service.Auth().GetLoginUserId(ctx)
	if loginUserId == 0 {
		return nil, gerror.NewCode(codes.CodePermissionDenied, "未登录或登录已过期")
	}

	// 从上下文获取用户角色信息，判断是否有权限分配角色
	// 实际应用中，这里应该检查当前登录用户是否有权限分配角色（如：管理员角色）
	// 如果当前用户不是管理员，则返回错误
	// TODO: 实际业务中应该判断当前登录用户是否有权限分配角色

	g.Log().Infof(ctx, "用户 %d 正在为用户 %d 分配角色 %d", loginUserId, in.UserId, in.RoleId)

	// 检查被分配角色的用户是否存在
	count, err := dao.Users.Ctx(ctx).Where(dao.Users.Columns().UserId, in.UserId).Count()
	if err != nil {
		return nil, err
	}
	if count == 0 {
		return nil, gerror.NewCode(codes.CodeNotFound, "用户不存在")
	}

	// 更新用户角色
	_, err = dao.Users.Ctx(ctx).Data(g.Map{
		dao.Users.Columns().RoleId:     in.RoleId,
		dao.Users.Columns().UpdateTime: gtime.Now(),
	}).Where(dao.Users.Columns().UserId, in.UserId).Update()
	if err != nil {
		return nil, err
	}

	return &model.UserAssignRoleOutput{
		Success: true,
	}, nil
}

package auth

import (
	"context"

	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"

	"github.com/patient-fyd/jxust-softhub-api/internal/codes"
	"github.com/patient-fyd/jxust-softhub-api/internal/dao"
	"github.com/patient-fyd/jxust-softhub-api/internal/model"
	"github.com/patient-fyd/jxust-softhub-api/internal/model/do"
	"github.com/patient-fyd/jxust-softhub-api/internal/model/entity"
	"github.com/patient-fyd/jxust-softhub-api/internal/service"
	"github.com/patient-fyd/jxust-softhub-api/utility"
)

type sAuth struct {
	jwtService *utility.JWTService
}

// New 创建认证服务
func New() service.IAuth {
	return &sAuth{
		jwtService: utility.NewJWTService(),
	}
}

// init 初始化
func init() {
	service.RegisterAuth(New())
}

// Register 用户注册
func (s *sAuth) Register(ctx context.Context, in model.RegisterInput) (*model.RegisterOutput, error) {
	// 1. 检查用户名是否已存在
	count, err := dao.Users.Ctx(ctx).Where(do.Users{
		UserName: in.UserName,
	}).Count()
	if err != nil {
		g.Log().Error(ctx, "Check username error:", err)
		return nil, gerror.Wrap(err, "检查用户名失败")
	}
	if count > 0 {
		return nil, gerror.NewCode(codes.CodeValidationFailed, "用户名已存在")
	}

	// 2. 检查邮箱是否已存在
	count, err = dao.Users.Ctx(ctx).Where(do.Users{
		Email: in.Email,
	}).Count()
	if err != nil {
		g.Log().Error(ctx, "Check email error:", err)
		return nil, gerror.Wrap(err, "检查邮箱失败")
	}
	if count > 0 {
		return nil, gerror.NewCode(codes.CodeValidationFailed, "邮箱已存在")
	}

	// 3. 检查手机号是否已存在
	count, err = dao.Users.Ctx(ctx).Where(do.Users{
		Phone: in.Phone,
	}).Count()
	if err != nil {
		g.Log().Error(ctx, "Check phone error:", err)
		return nil, gerror.Wrap(err, "检查手机号失败")
	}
	if count > 0 {
		return nil, gerror.NewCode(codes.CodeValidationFailed, "手机号已存在")
	}

	// 4. 创建用户
	now := gtime.Now()
	user := do.Users{
		UserName:   in.UserName,
		Password:   gmd5.MustEncryptString(in.Password), // MD5加密密码
		Name:       in.Name,
		Email:      in.Email,
		Phone:      in.Phone,
		RoleId:     1, // 1:普通用户
		CreateTime: now,
		UpdateTime: now,
	}

	// 5. 插入数据
	result, err := dao.Users.Ctx(ctx).Data(user).Insert()
	if err != nil {
		g.Log().Error(ctx, "Create user error:", err)
		return nil, gerror.Wrap(err, "创建用户失败")
	}

	// 6. 获取用户ID
	userId, err := result.LastInsertId()
	if err != nil {
		g.Log().Error(ctx, "Get last insert id error:", err)
		return nil, gerror.Wrap(err, "获取用户ID失败")
	}

	// 7. 插入用户-角色关系
	// 注：由于我们直接在用户表中已经设置了角色，这里省略了角色权限表的操作
	// role-permission关系需要在后续的权限系统中完善

	// 8. 返回用户信息
	return &model.RegisterOutput{
		UserId:     uint(userId),
		UserName:   in.UserName,
		CreateTime: now.String(),
	}, nil
}

// Login 用户登录
func (s *sAuth) Login(ctx context.Context, in model.LoginInput) (*model.LoginOutput, error) {
	// 1. 查询用户
	user := entity.Users{}
	err := dao.Users.Ctx(ctx).Where(do.Users{
		UserName: in.UserName,
	}).Scan(&user)
	if err != nil {
		g.Log().Error(ctx, "Query user error:", err)
		return nil, gerror.Wrap(err, "查询用户失败")
	}

	// 2. 检查用户是否存在
	if user.UserId == 0 {
		return nil, gerror.NewCode(codes.CodeValidationFailed, "用户名或密码错误")
	}

	// 3. 校验密码
	if user.Password != gmd5.MustEncryptString(in.Password) {
		return nil, gerror.NewCode(codes.CodeValidationFailed, "用户名或密码错误")
	}

	// 4. 生成 Token
	token, err := s.jwtService.GenerateToken(ctx, &model.ClaimsUser{
		UserId:   user.UserId,
		UserName: user.UserName,
		RoleId:   model.UserRoleEnum(user.RoleId),
	})
	if err != nil {
		g.Log().Error(ctx, "Generate token error:", err)
		return nil, gerror.Wrap(err, "生成令牌失败")
	}

	// 5. 返回登录结果
	return &model.LoginOutput{
		Token: token,
		User: &model.LoginUser{
			UserId:   user.UserId,
			UserName: user.UserName,
			Name:     user.Name,
		},
	}, nil
}

// RefreshToken 刷新Token
func (s *sAuth) RefreshToken(ctx context.Context, in model.TokenRefreshInput) (*model.TokenRefreshOutput, error) {
	// 1. 刷新令牌
	newToken, err := s.jwtService.RefreshToken(ctx, in.Token)
	if err != nil {
		g.Log().Error(ctx, "Refresh token error:", err)
		return nil, gerror.NewCode(codes.CodeNotAvailable, "刷新令牌失败")
	}

	// 2. 返回新令牌
	return &model.TokenRefreshOutput{
		Token: newToken,
	}, nil
}

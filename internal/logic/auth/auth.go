package auth

import (
	"context"
	"time"

	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/golang-jwt/jwt/v4"

	"github.com/patient-fyd/jxust-softhub-api/internal/codes"
	"github.com/patient-fyd/jxust-softhub-api/internal/dao"
	"github.com/patient-fyd/jxust-softhub-api/internal/model"
	"github.com/patient-fyd/jxust-softhub-api/internal/model/do"
	"github.com/patient-fyd/jxust-softhub-api/internal/model/entity"
	"github.com/patient-fyd/jxust-softhub-api/internal/service"
)

// sAuth 认证服务实现
type sAuth struct {
	// jwtSecret JWT密钥
	jwtSecret string
	// jwtExpire JWT过期时间（秒）
	jwtExpire int
}

// New 创建认证服务
func New() *sAuth {
	return &sAuth{
		jwtSecret: g.Cfg().MustGet(context.Background(), "jwt.secret").String(),
		jwtExpire: g.Cfg().MustGet(context.Background(), "jwt.expire").Int(),
	}
}

// init 初始化
func init() {
	service.RegisterAuth(New())
}

// Register 用户注册
func (s *sAuth) Register(ctx context.Context, in model.UserRegisterInput) (*model.UserRegisterOutput, error) {
	// 检查用户名是否已存在
	count, err := dao.Users.Ctx(ctx).Where(do.Users{
		UserName: in.UserName,
	}).Count()
	if err != nil {
		g.Log().Error(ctx, "Check username error:", err)
		return nil, gerror.Wrap(err, "检查用户名失败")
	}
	if count > 0 {
		return nil, gerror.NewCode(codes.CodeValidationFailed, "用户名已被注册")
	}

	// 检查邮箱是否已存在
	count, err = dao.Users.Ctx(ctx).Where(do.Users{
		Email: in.Email,
	}).Count()
	if err != nil {
		g.Log().Error(ctx, "Check email error:", err)
		return nil, gerror.Wrap(err, "检查邮箱失败")
	}
	if count > 0 {
		return nil, gerror.NewCode(codes.CodeValidationFailed, "邮箱已被注册")
	}

	// 密码加密
	password, err := gmd5.EncryptString(in.Password)
	if err != nil {
		g.Log().Error(ctx, "Password encrypt error:", err)
		return nil, gerror.Wrap(err, "密码加密失败")
	}

	// 创建用户
	result, err := dao.Users.Ctx(ctx).Data(do.Users{
		UserName: in.UserName,
		Password: password,
		Name:     in.Name,
		Email:    in.Email,
		Phone:    in.Phone,
		RoleId:   4, // 默认为普通学生，角色ID为4
	}).Insert()
	if err != nil {
		g.Log().Error(ctx, "Insert user error:", err)
		return nil, gerror.Wrap(err, "创建用户失败")
	}

	// 获取用户ID
	userId, err := result.LastInsertId()
	if err != nil {
		g.Log().Error(ctx, "Get last insert id error:", err)
		return nil, gerror.Wrap(err, "获取用户ID失败")
	}

	return &model.UserRegisterOutput{
		UserId:   uint(userId),
		UserName: in.UserName,
		Name:     in.Name,
		RoleId:   4, // 默认为普通学生，角色ID为4
	}, nil
}

// Login 用户登录
func (s *sAuth) Login(ctx context.Context, in model.UserLoginInput) (*model.UserLoginOutput, error) {
	// 密码加密
	password, err := gmd5.EncryptString(in.Password)
	if err != nil {
		g.Log().Error(ctx, "Password encrypt error:", err)
		return nil, gerror.Wrap(err, "密码加密失败")
	}

	// 查询用户
	var user *entity.Users
	err = dao.Users.Ctx(ctx).Where(do.Users{
		UserName: in.UserName,
		Password: password,
	}).Scan(&user)
	if err != nil {
		g.Log().Error(ctx, "Query user error:", err)
		return nil, gerror.Wrap(err, "查询用户失败")
	}
	if user == nil {
		return nil, gerror.NewCode(codes.CodeValidationFailed, "用户名或密码错误")
	}

	// 生成Token
	token, err := s.generateToken(user.UserId)
	if err != nil {
		g.Log().Error(ctx, "Generate token error:", err)
		return nil, gerror.Wrap(err, "生成Token失败")
	}

	return &model.UserLoginOutput{
		Token:     token,
		UserId:    user.UserId,
		UserName:  user.UserName,
		Name:      user.Name,
		RoleId:    user.RoleId,
		LoginTime: gtime.Now(),
	}, nil
}

// AdminLogin 管理员登录
func (s *sAuth) AdminLogin(ctx context.Context, in model.AdminLoginInput) (*model.AdminLoginOutput, error) {
	// 密码加密
	password, err := gmd5.EncryptString(in.Password)
	if err != nil {
		g.Log().Error(ctx, "Password encrypt error:", err)
		return nil, gerror.Wrap(err, "密码加密失败")
	}

	// 查询用户
	var user *entity.Users
	err = dao.Users.Ctx(ctx).Where(do.Users{
		UserName: in.UserName,
		Password: password,
	}).Scan(&user)
	if err != nil {
		g.Log().Error(ctx, "Query user error:", err)
		return nil, gerror.Wrap(err, "查询用户失败")
	}
	if user == nil {
		return nil, gerror.NewCode(codes.CodeValidationFailed, "用户名或密码错误")
	}

	// 检查是否为管理员
	if user.RoleId < 1 {
		return nil, gerror.NewCode(codes.CodeNotAuthorized, "无管理员权限")
	}

	// 生成Token
	token, err := s.generateToken(user.UserId)
	if err != nil {
		g.Log().Error(ctx, "Generate token error:", err)
		return nil, gerror.Wrap(err, "生成Token失败")
	}

	return &model.AdminLoginOutput{
		Token:     token,
		UserId:    user.UserId,
		UserName:  user.UserName,
		Name:      user.Name,
		RoleId:    user.RoleId,
		LoginTime: gtime.Now(),
	}, nil
}

// RefreshToken 刷新Token
func (s *sAuth) RefreshToken(ctx context.Context, in model.RefreshTokenInput) (*model.RefreshTokenOutput, error) {
	// 验证原Token
	token, err := jwt.Parse(in.Token, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.jwtSecret), nil
	})
	if err != nil {
		g.Log().Error(ctx, "Parse token error:", err)
		return nil, gerror.NewCode(codes.CodeValidationFailed, "Token无效")
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		// 获取用户ID
		userId := uint(claims["userId"].(float64))

		// 生成新的Token
		newToken, err := s.generateToken(userId)
		if err != nil {
			g.Log().Error(ctx, "Generate token error:", err)
			return nil, gerror.Wrap(err, "生成Token失败")
		}

		return &model.RefreshTokenOutput{
			Token: newToken,
		}, nil
	}

	return nil, gerror.NewCode(codes.CodeValidationFailed, "Token无效")
}

// generateToken 生成JWT Token
func (s *sAuth) generateToken(userId uint) (string, error) {
	claims := jwt.MapClaims{
		"userId": userId,
		"exp":    time.Now().Add(time.Duration(s.jwtExpire) * time.Second).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.jwtSecret))
}

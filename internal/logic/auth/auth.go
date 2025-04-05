package auth

import (
	"context"
	"time"

	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
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

// GetLoginUserId 获取当前登录用户ID
func (s *sAuth) GetLoginUserId(ctx context.Context) uint {
	value := ctx.Value("userId")
	if value == nil {
		return 0
	}
	userId, ok := value.(uint)
	if !ok {
		return 0
	}
	return userId
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

	// 检查是否为超级管理员（角色ID为1）
	if user.RoleId != 1 {
		return nil, gerror.NewCode(codes.CodeNotAuthorized, "无超级管理员权限")
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

// IsAdmin 检查当前用户是否为管理员
func (s *sAuth) IsAdmin(ctx context.Context) (bool, error) {
	// 获取当前登录用户ID
	userId := s.GetLoginUserId(ctx)
	if userId <= 0 {
		return false, gerror.New("用户未登录")
	}

	// 查询用户角色
	user, err := dao.Users.Ctx(ctx).Where("userId", userId).One()
	if err != nil {
		return false, err
	}
	if user.IsEmpty() {
		return false, gerror.New("用户不存在")
	}

	// 检查是否为管理员（角色ID为1-3）
	roleId := user["roleId"].Uint()
	return roleId >= 1 && roleId <= 3, nil
}

// CheckCapcha 检查验证码
func (s *sAuth) CheckCapcha(ctx context.Context, in model.CodeInput) bool {
	// 在实际应用中，这里应该从缓存中获取验证码并进行比对
	// 由于当前系统没有实现验证码功能，这里默认返回true表示验证通过
	return true
}

// CheckPasswordInFile 检查密码是否在文件中
func (s *sAuth) CheckPasswordInFile(ctx context.Context, password string) bool {
	// 在实际应用中，这里应该检查密码是否在常见密码文件中
	// 由于当前系统没有实现此功能，这里默认返回false表示密码不在常见密码文件中
	return false
}

// GenerateToken 生成令牌
func (s *sAuth) GenerateToken(ctx context.Context, userId int, expire ...int) (string, error) {
	expireTime := s.jwtExpire
	if len(expire) > 0 && expire[0] > 0 {
		expireTime = expire[0]
	}
	return s.generateTokenWithExpire(uint(userId), expireTime)
}

// GenerateTokenFromDevice 生成设备令牌
func (s *sAuth) GenerateTokenFromDevice(ctx context.Context, userId int) (string, error) {
	// 设备令牌有更长的过期时间，这里设置为30天
	return s.generateTokenWithExpire(uint(userId), 30*24*60*60)
}

// 用于生成JWT令牌，支持自定义过期时间
func (s *sAuth) generateTokenWithExpire(userId uint, expireSeconds int) (string, error) {
	// 设置JWT的过期时间
	expireTime := time.Now().Add(time.Duration(expireSeconds) * time.Second)

	// 创建JWT claims
	claims := jwt.MapClaims{
		"userId": userId,
		"exp":    expireTime.Unix(),
	}

	// 创建token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 签名token
	return token.SignedString([]byte(s.jwtSecret))
}

// VerifyToken 验证令牌
func (s *sAuth) VerifyToken(ctx context.Context, tokenString string) (*model.ClaimsOutput, error) {
	// 解析JWT token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// 验证签名算法
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, gerror.New("无效的令牌签名算法")
		}
		return []byte(s.jwtSecret), nil
	})

	if err != nil {
		return nil, gerror.New("令牌验证失败: " + err.Error())
	}

	// 验证token是否有效
	if !token.Valid {
		return nil, gerror.New("无效的令牌")
	}

	// 获取token中的claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, gerror.New("无效的令牌内容")
	}

	// 获取用户ID
	userId, ok := claims["userId"].(float64)
	if !ok {
		return nil, gerror.New("无效的用户ID")
	}

	// 查询用户信息
	user, err := dao.Users.Ctx(ctx).Where("userId", uint(userId)).One()
	if err != nil {
		return nil, err
	}
	if user.IsEmpty() {
		return nil, gerror.New("用户不存在")
	}

	// 返回用户信息
	expiresAt, _ := claims["exp"].(float64)
	return &model.ClaimsOutput{
		UserId:    uint(userId),
		UserName:  user["userName"].String(),
		Name:      user["name"].String(),
		RoleId:    user["roleId"].Uint(),
		ExpiresAt: int64(expiresAt),
	}, nil
}

// VerifyTokenMiddleware 令牌验证中间件
func (s *sAuth) VerifyTokenMiddleware(r *ghttp.Request) {
	// 获取请求头中的token
	token := r.Header.Get("Authorization")
	if token == "" {
		r.Middleware.Next()
		return
	}

	// 去除Bearer前缀
	if len(token) > 7 && token[:7] == "Bearer " {
		token = token[7:]
	}

	// 验证token
	claims, err := s.VerifyToken(r.Context(), token)
	if err != nil {
		r.Middleware.Next()
		return
	}

	// 将用户ID存入上下文
	ctx := context.WithValue(r.Context(), "userId", claims.UserId)
	r.SetCtx(ctx)

	r.Middleware.Next()
}

// GetLoginUser 获取当前登录用户
func (s *sAuth) GetLoginUser(ctx context.Context) (*entity.Users, error) {
	userId := s.GetLoginUserId(ctx)
	if userId <= 0 {
		return nil, gerror.New("用户未登录")
	}

	var user *entity.Users
	err := dao.Users.Ctx(ctx).Where("userId", userId).Scan(&user)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, gerror.New("用户不存在")
	}

	return user, nil
}

// LoginState 获取登录状态
func (s *sAuth) LoginState(ctx context.Context) bool {
	return s.GetLoginUserId(ctx) > 0
}

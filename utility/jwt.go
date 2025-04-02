package utility

import (
	"context"
	"errors"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/golang-jwt/jwt/v5"
	"github.com/patient-fyd/jxust-softhub-api/internal/model"
)

var (
	ErrInvalidToken  = errors.New("invalid token")
	ErrExpiredToken  = errors.New("token has expired")
	ErrInvalidClaims = errors.New("invalid token claims")
)

const (
	// 默认过期时间：24小时
	defaultExpireTime = 24 * time.Hour
	// 刷新Token时间：22小时
	refreshTime = 22 * time.Hour
	// JWT密钥
	jwtKey = "jxust-softhub-api-jwt-key"
)

// CustomClaims 自定义JWT声明
type CustomClaims struct {
	UID   uint               `json:"uid"`
	UName string             `json:"uname"`
	RID   model.UserRoleEnum `json:"rid"`
	jwt.RegisteredClaims
}

// JWTService JWT服务
type JWTService struct{}

// NewJWTService 创建JWT服务
func NewJWTService() *JWTService {
	return &JWTService{}
}

// GenerateToken 生成JWT令牌
func (s *JWTService) GenerateToken(ctx context.Context, user *model.ClaimsUser) (string, error) {
	// 创建自定义声明
	claims := CustomClaims{
		UID:   user.UserId,
		UName: user.UserName,
		RID:   user.RoleId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(defaultExpireTime)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	// 创建令牌
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 使用密钥签名令牌
	tokenStr, err := token.SignedString([]byte(jwtKey))
	if err != nil {
		g.Log().Error(ctx, "Generate token error:", err)
		return "", err
	}
	return tokenStr, nil
}

// ParseToken 解析JWT令牌
func (s *JWTService) ParseToken(ctx context.Context, tokenString string) (*model.ClaimsUser, error) {
	// 解析令牌
	token, err := jwt.ParseWithClaims(
		tokenString,
		&CustomClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		},
	)

	if err != nil {
		// 检查错误类型
		if errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet) {
			return nil, ErrExpiredToken
		}
		g.Log().Error(ctx, "Parse token error:", err)
		return nil, ErrInvalidToken
	}

	// 从令牌中获取声明
	claims, ok := token.Claims.(*CustomClaims)
	if !ok || !token.Valid {
		return nil, ErrInvalidClaims
	}

	// 转换声明为用户信息
	return &model.ClaimsUser{
		UserId:   claims.UID,
		UserName: claims.UName,
		RoleId:   claims.RID,
	}, nil
}

// RefreshToken 刷新JWT令牌
func (s *JWTService) RefreshToken(ctx context.Context, tokenString string) (string, error) {
	// 解析旧令牌
	user, err := s.ParseToken(ctx, tokenString)
	if err != nil && !errors.Is(err, ErrExpiredToken) {
		return "", err
	}

	// 生成新令牌
	return s.GenerateToken(ctx, user)
}

// GetTokenExpireTime 获取令牌过期时间
func (s *JWTService) GetTokenExpireTime() time.Duration {
	return defaultExpireTime
}

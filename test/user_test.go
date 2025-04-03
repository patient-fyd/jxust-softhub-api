package test

import (
	"context"
	"fmt"
	"testing"

	_ "github.com/patient-fyd/jxust-softhub-api/internal/logic/auth" // 确保init函数被执行
	_ "github.com/patient-fyd/jxust-softhub-api/internal/logic/user" // 确保init函数被执行
	"github.com/patient-fyd/jxust-softhub-api/internal/service"
)

// TestAuthGetLoginUserID 测试从上下文中获取用户ID
func TestAuthGetLoginUserID(t *testing.T) {
	// 创建包含用户ID的上下文
	ctx := context.WithValue(context.Background(), "userId", uint(1))

	// 获取用户ID
	userId := service.Auth().GetLoginUserId(ctx)
	if userId != 1 {
		t.Errorf("Expected userId to be 1, got %d", userId)
	}

	fmt.Println("从上下文中成功获取用户ID:", userId)
}

// TestUserGetUserInfo 测试获取用户信息
func TestUserGetUserInfo(t *testing.T) {
	// 创建上下文
	ctx := context.Background()

	// 获取用户信息
	userInfo, err := service.User().GetUserInfo(ctx, 1)
	if err != nil {
		t.Fatalf("GetUserInfo error: %v", err)
	}

	fmt.Printf("成功获取用户信息: %+v\n", userInfo)
}

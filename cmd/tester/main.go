package main

import (
	"context"
	"fmt"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/os/gctx"

	_ "github.com/patient-fyd/jxust-softhub-api/internal/logic/auth"
	_ "github.com/patient-fyd/jxust-softhub-api/internal/logic/user"
	"github.com/patient-fyd/jxust-softhub-api/internal/model"
	"github.com/patient-fyd/jxust-softhub-api/internal/service"
)

func main() {
	ctx := gctx.New()

	// 设置配置文件目录
	g.Cfg().GetAdapter().(*gcfg.AdapterFile).SetPath("manifest/config")

	// 测试Auth服务
	testAuthService(ctx)

	// 测试User服务
	testUserService(ctx)
}

// 测试Auth服务
func testAuthService(ctx context.Context) {
	fmt.Println("=== 测试Auth服务 ===")

	// 创建包含用户ID的上下文
	ctxWithUserId := context.WithValue(ctx, "userId", uint(1))

	// 获取用户ID
	userId := service.Auth().GetLoginUserId(ctxWithUserId)
	fmt.Println("从上下文中获取的用户ID:", userId)

	// 尝试获取空上下文的用户ID
	emptyUserId := service.Auth().GetLoginUserId(ctx)
	fmt.Println("从空上下文中获取的用户ID:", emptyUserId)
}

// 测试User服务
func testUserService(ctx context.Context) {
	fmt.Println("\n=== 测试User服务 ===")

	// 获取用户信息
	userInfo, err := service.User().GetUserInfo(ctx, 1)
	if err != nil {
		fmt.Printf("获取用户信息失败: %v\n", err)
		return
	}

	fmt.Printf("成功获取用户信息: %+v\n", userInfo)

	// 获取用户列表
	userList, err := service.User().GetUserList(ctx, model.UserListInput{
		PageNum:  1,
		PageSize: 10,
	})
	if err != nil {
		fmt.Printf("获取用户列表失败: %v\n", err)
		return
	}

	fmt.Printf("成功获取用户列表，总数: %d\n", userList.Total)
}

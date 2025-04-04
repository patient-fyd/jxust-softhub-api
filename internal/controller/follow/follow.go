// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package follow

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"

	"github.com/patient-fyd/jxust-softhub-api/internal/model"
	"github.com/patient-fyd/jxust-softhub-api/internal/service"
)

// 关注API控制器
type Controller struct{}

// Toggle 关注/取消关注
func (c *Controller) Toggle(ctx context.Context, req *model.FollowToggleReq) (res *model.FollowToggleOutput, err error) {
	// 判断用户是否登录
	if !service.Auth().LoginState(ctx) {
		return nil, gerror.New("用户未登录")
	}

	return service.Follow().Toggle(ctx, *req)
}

// FollowingList 获取关注列表
func (c *Controller) FollowingList(ctx context.Context, req *model.FollowingListReq) (res *model.FollowingListOutput, err error) {
	return service.Follow().FollowingList(ctx, *req)
}

// FollowerList 获取粉丝列表
func (c *Controller) FollowerList(ctx context.Context, req *model.FollowerListReq) (res *model.FollowerListOutput, err error) {
	return service.Follow().FollowerList(ctx, *req)
}

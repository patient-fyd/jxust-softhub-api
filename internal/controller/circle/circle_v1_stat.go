package circle

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"

	v1 "github.com/patient-fyd/jxust-softhub-api/api/circle/v1"
	"github.com/patient-fyd/jxust-softhub-api/internal/model"
	"github.com/patient-fyd/jxust-softhub-api/internal/service"
	"github.com/patient-fyd/jxust-softhub-api/utility/auth"
)

// CircleStat 获取圈子统计信息
func (c *ControllerV1) CircleStat(ctx context.Context, req *v1.CircleStatReq) (res *v1.CircleStatRes, err error) {
	// 判断用户是否已登录
	userId := auth.GetLoginUserId(ctx)
	if userId <= 0 {
		return nil, gerror.New("请先登录")
	}

	// 构造服务层输入参数
	input := model.CircleStatInput{
		UserId: int(userId),
	}

	// 调用服务层
	output, err := service.Circle().CircleStat(ctx, input)
	if err != nil {
		return nil, gerror.Wrap(err, "获取圈子统计信息失败")
	}

	// 构造响应数据
	res = &v1.CircleStatRes{
		TotalCount:     output.TotalCount,
		FollowingCount: output.FollowingCount,
		RecentActive:   make([]v1.CircleItem, 0, len(output.RecentActive)),
	}

	// 转换最近活跃圈子列表数据
	for _, item := range output.RecentActive {
		res.RecentActive = append(res.RecentActive, v1.CircleItem{
			CircleId:    item.CircleId,
			Name:        item.Name,
			Description: item.Description,
			Icon:        item.Icon,
			PostCount:   item.PostCount,
			MemberCount: item.MemberCount,
			IsOfficial:  item.IsOfficial,
			IsFollowed:  item.IsFollowed,
			CreateTime:  item.CreateTime.String(),
		})
	}

	return res, nil
}

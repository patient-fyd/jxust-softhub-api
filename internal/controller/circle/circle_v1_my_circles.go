package circle

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"

	v1 "github.com/patient-fyd/jxust-softhub-api/api/circle/v1"
	"github.com/patient-fyd/jxust-softhub-api/internal/model"
	"github.com/patient-fyd/jxust-softhub-api/internal/service"
	"github.com/patient-fyd/jxust-softhub-api/utility/auth"
)

// MyCircles 获取用户关注的圈子列表
func (c *ControllerV1) MyCircles(ctx context.Context, req *v1.MyCirclesReq) (res *v1.MyCirclesRes, err error) {
	// 判断用户是否已登录
	userId := auth.GetLoginUserId(ctx)
	if userId <= 0 {
		return nil, gerror.New("请先登录")
	}

	// 构造服务层输入参数
	input := model.CircleMyCirclesInput{
		UserId: int(userId),
		Page:   req.Page,
		Size:   req.Size,
	}

	// 调用服务层
	output, err := service.Circle().MyCircles(ctx, input)
	if err != nil {
		return nil, gerror.Wrap(err, "获取我的圈子列表失败")
	}

	// 构造响应数据
	res = &v1.MyCirclesRes{
		List:  make([]v1.CircleItem, 0, len(output.List)),
		Page:  output.Page,
		Size:  output.Size,
		Total: output.Total,
	}

	// 转换圈子列表数据
	for _, item := range output.List {
		res.List = append(res.List, v1.CircleItem{
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

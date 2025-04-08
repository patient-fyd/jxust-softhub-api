package circle

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"

	v1 "github.com/patient-fyd/jxust-softhub-api/api/circle/v1"
	"github.com/patient-fyd/jxust-softhub-api/internal/model"
	"github.com/patient-fyd/jxust-softhub-api/internal/service"
)

func (c *ControllerV1) List(ctx context.Context, req *v1.ListReq) (res *v1.ListRes, err error) {
	// 转换请求参数
	input := model.CircleListInput{
		Page:    req.Page,
		Size:    req.Size,
		Keyword: req.Keyword,
		UserId:  req.UserId,
		OrderBy: req.OrderBy,
	}

	// 调用服务层
	output, err := service.Circle().List(ctx, input)
	if err != nil {
		return nil, gerror.Wrap(err, "获取圈子列表失败")
	}

	// 转换响应参数
	res = &v1.ListRes{
		Total: output.Total,
		Page:  output.Page,
		Size:  output.Size,
		List:  make([]v1.CircleItem, 0, len(output.List)),
	}

	// 转换列表数据
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

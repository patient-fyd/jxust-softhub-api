package circle

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"

	v1 "github.com/patient-fyd/jxust-softhub-api/api/circle/v1"
	"github.com/patient-fyd/jxust-softhub-api/internal/model"
	"github.com/patient-fyd/jxust-softhub-api/internal/service"
)

func (c *ControllerV1) Detail(ctx context.Context, req *v1.DetailReq) (res *v1.DetailRes, err error) {
	// 转换请求参数
	input := model.CircleDetailInput{
		CircleId: req.CircleId,
	}

	// 调用服务层
	output, err := service.Circle().Detail(ctx, input)
	if err != nil {
		return nil, gerror.Wrap(err, "获取圈子详情失败")
	}

	// 如果没有找到圈子
	if output == nil {
		return nil, gerror.New("圈子不存在或已删除")
	}

	// 转换响应参数
	res = &v1.DetailRes{
		CircleId:    output.CircleId,
		Name:        output.Name,
		Description: output.Description,
		Icon:        output.Icon,
		PostCount:   output.PostCount,
		MemberCount: output.MemberCount,
		CreatorId:   output.CreatorId,
		CreatorName: output.CreatorName,
		IsOfficial:  output.IsOfficial,
		IsFollowed:  output.IsFollowed,
		CreateTime:  output.CreateTime.String(),
		UpdateTime:  output.UpdateTime.String(),
	}

	return res, nil
}

package follow

import (
	"context"

	v1 "github.com/patient-fyd/jxust-softhub-api/api/follow/v1"
	"github.com/patient-fyd/jxust-softhub-api/internal/model"
	"github.com/patient-fyd/jxust-softhub-api/internal/service"
)

func (c *ControllerV1) FollowerList(ctx context.Context, req *v1.FollowerListReq) (res *v1.FollowerListRes, err error) {
	input := model.FollowerListReq{
		FollowedId: req.FollowedId,
		FollowType: req.FollowType,
		Page:       req.Page,
		Size:       req.Size,
	}

	output, err := service.Follow().FollowerList(ctx, input)
	if err != nil {
		return nil, err
	}

	return &v1.FollowerListRes{
		List:  output.List,
		Total: output.Total,
		Page:  output.Page,
		Size:  output.Size,
	}, nil
}

package follow

import (
	"context"

	v1 "github.com/patient-fyd/jxust-softhub-api/api/follow/v1"
	"github.com/patient-fyd/jxust-softhub-api/internal/model"
	"github.com/patient-fyd/jxust-softhub-api/internal/service"
)

func (c *ControllerV1) FollowingList(ctx context.Context, req *v1.FollowingListReq) (res *v1.FollowingListRes, err error) {
	input := model.FollowingListReq{
		UserId:     req.UserId,
		FollowType: req.FollowType,
		Page:       req.Page,
		Size:       req.Size,
	}

	output, err := service.Follow().FollowingList(ctx, input)
	if err != nil {
		return nil, err
	}

	return &v1.FollowingListRes{
		List:  output.List,
		Total: output.Total,
		Page:  output.Page,
		Size:  output.Size,
	}, nil
}

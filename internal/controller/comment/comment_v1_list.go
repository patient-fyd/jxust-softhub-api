package comment

import (
	"context"

	v1 "github.com/patient-fyd/jxust-softhub-api/api/comment/v1"
	"github.com/patient-fyd/jxust-softhub-api/internal/model"
	"github.com/patient-fyd/jxust-softhub-api/internal/service"
)

func (c *ControllerV1) CommentList(ctx context.Context, req *v1.CommentListReq) (res *v1.CommentListRes, err error) {
	input := model.CommentListInput{
		ContentType: req.ContentType,
		ContentId:   req.ContentId,
		Page:        req.Page,
		Size:        req.Size,
	}

	output, err := service.Comment().List(ctx, input)
	if err != nil {
		return nil, err
	}

	return &v1.CommentListRes{
		List:  output.List,
		Page:  output.Page,
		Size:  output.Size,
		Total: output.Total,
	}, nil
}

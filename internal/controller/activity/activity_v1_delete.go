package activity

import (
	"context"

	v1 "github.com/patient-fyd/jxust-softhub-api/api/activity/v1"
	"github.com/patient-fyd/jxust-softhub-api/internal/model"
	"github.com/patient-fyd/jxust-softhub-api/internal/service"
)

// 删除活动
func (c *ControllerV1) Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error) {
	// 调用服务删除活动
	err = service.Activity().Delete(ctx, model.ActivityDeleteInput{
		ActivityId: req.ActivityId,
	})
	if err != nil {
		return nil, err
	}

	return &v1.DeleteRes{}, nil
}

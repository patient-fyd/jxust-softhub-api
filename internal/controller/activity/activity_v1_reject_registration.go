package activity

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"

	v1 "github.com/patient-fyd/jxust-softhub-api/api/activity/v1"
	"github.com/patient-fyd/jxust-softhub-api/internal/model"
	"github.com/patient-fyd/jxust-softhub-api/internal/service"
)

// 拒绝报名
func (c *ControllerV1) RejectRegistration(ctx context.Context, req *v1.RejectRegistrationReq) (res *v1.RejectRegistrationRes, err error) {
	// 获取当前登录用户ID
	var reviewerId uint = 0
	if v := g.RequestFromCtx(ctx).GetCtxVar("userId"); !v.IsNil() {
		reviewerId = v.Uint()
	}

	// 调用服务审核报名
	output, err := service.Activity().ReviewRegistration(ctx, model.RegistrationReviewInput{
		RegistrationId: req.RegistrationId,
		Status:         model.RegistrationStatusRejected,
		Reason:         req.Reason,
		ReviewerId:     reviewerId,
	})
	if err != nil {
		return nil, err
	}

	return &v1.RejectRegistrationRes{
		RegistrationId: output.RegistrationId,
	}, nil
}

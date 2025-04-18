package activity

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"

	v1 "github.com/patient-fyd/jxust-softhub-api/api/activity/v1"
	"github.com/patient-fyd/jxust-softhub-api/internal/model"
	"github.com/patient-fyd/jxust-softhub-api/internal/service"
)

// 活动报名
func (c *ControllerV1) Register(ctx context.Context, req *v1.RegisterReq) (res *v1.RegisterRes, err error) {
	// 获取当前登录用户ID，如果未登录则为0
	var userId uint = 0
	if v := g.RequestFromCtx(ctx).GetCtxVar("userId"); !v.IsNil() {
		userId = v.Uint()
	}

	// 调用服务创建报名
	output, err := service.Activity().Register(ctx, model.ActivityRegisterInput{
		ActivityId: req.ActivityId,
		UserId:     userId,
		Name:       req.Name,
		StudentId:  req.StudentId,
		Contact:    req.Contact,
	})
	if err != nil {
		return nil, err
	}

	return &v1.RegisterRes{
		RegistrationId: output.RegistrationId,
	}, nil
}

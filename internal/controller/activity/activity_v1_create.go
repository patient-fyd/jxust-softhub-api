package activity

import (
	"context"

	"github.com/gogf/gf/v2/os/gtime"

	v1 "github.com/patient-fyd/jxust-softhub-api/api/activity/v1"
	"github.com/patient-fyd/jxust-softhub-api/internal/model"
	"github.com/patient-fyd/jxust-softhub-api/internal/service"
)

// 创建活动
func (c *ControllerV1) Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error) {
	// 解析时间
	startTime, err := gtime.StrToTime(req.StartTime)
	if err != nil {
		return nil, err
	}
	endTime, err := gtime.StrToTime(req.EndTime)
	if err != nil {
		return nil, err
	}

	// 调用服务创建活动
	output, err := service.Activity().Create(ctx, model.ActivityCreateInput{
		Title:           req.Title,
		Description:     req.Description,
		StartTime:       startTime.Time,
		EndTime:         endTime.Time,
		Location:        req.Location,
		MaxParticipants: req.MaxParticipants,
	})
	if err != nil {
		return nil, err
	}

	return &v1.CreateRes{
		ActivityId: output.ActivityId,
	}, nil
}

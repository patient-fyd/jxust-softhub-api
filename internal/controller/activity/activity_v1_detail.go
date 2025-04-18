package activity

import (
	"context"

	v1 "github.com/patient-fyd/jxust-softhub-api/api/activity/v1"
	"github.com/patient-fyd/jxust-softhub-api/internal/model"
	"github.com/patient-fyd/jxust-softhub-api/internal/service"
)

// 获取活动详情
func (c *ControllerV1) Detail(ctx context.Context, req *v1.DetailReq) (res *v1.DetailRes, err error) {
	// 调用服务查询活动详情
	output, err := service.Activity().Detail(ctx, model.ActivityDetailInput{
		ActivityId: req.ActivityId,
	})
	if err != nil {
		return nil, err
	}

	// 转换活动数据
	return &v1.DetailRes{
		Activity: v1.ActivityInfo{
			ActivityId:      output.Activity.ActivityId,
			Title:           output.Activity.Title,
			Description:     output.Activity.Description,
			StartTime:       output.Activity.StartTime.Format("2006-01-02 15:04:05"),
			EndTime:         output.Activity.EndTime.Format("2006-01-02 15:04:05"),
			Location:        output.Activity.Location,
			MaxParticipants: output.Activity.MaxParticipants,
			Status:          output.Activity.Status,
			CreateTime:      output.Activity.CreateTime.Format("2006-01-02 15:04:05"),
		},
	}, nil
}

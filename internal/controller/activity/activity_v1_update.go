package activity

import (
	"context"

	"github.com/gogf/gf/v2/os/gtime"

	v1 "github.com/patient-fyd/jxust-softhub-api/api/activity/v1"
	"github.com/patient-fyd/jxust-softhub-api/internal/model"
	"github.com/patient-fyd/jxust-softhub-api/internal/service"
)

// 更新活动
func (c *ControllerV1) Update(ctx context.Context, req *v1.UpdateReq) (res *v1.UpdateRes, err error) {
	// 构建更新输入参数
	input := model.ActivityUpdateInput{
		ActivityId: req.ActivityId,
	}

	// 有传值时才更新
	if req.Title != "" {
		title := req.Title
		input.Title = &title
	}
	if req.Description != "" {
		description := req.Description
		input.Description = &description
	}
	if req.Location != "" {
		location := req.Location
		input.Location = &location
	}
	if req.MaxParticipants >= 0 {
		maxParticipants := req.MaxParticipants
		input.MaxParticipants = &maxParticipants
	}
	if req.Status >= 0 {
		status := req.Status
		input.Status = &status
	}

	// 处理时间
	if req.StartTime != "" {
		startTime, err := gtime.StrToTime(req.StartTime)
		if err != nil {
			return nil, err
		}
		t := startTime.Time
		input.StartTime = &t
	}
	if req.EndTime != "" {
		endTime, err := gtime.StrToTime(req.EndTime)
		if err != nil {
			return nil, err
		}
		t := endTime.Time
		input.EndTime = &t
	}

	// 调用服务更新活动
	output, err := service.Activity().Update(ctx, input)
	if err != nil {
		return nil, err
	}

	return &v1.UpdateRes{
		ActivityId: output.ActivityId,
	}, nil
}

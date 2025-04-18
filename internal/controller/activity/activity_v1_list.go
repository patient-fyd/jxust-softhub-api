package activity

import (
	"context"

	"github.com/patient-fyd/jxust-softhub-api/api"
	v1 "github.com/patient-fyd/jxust-softhub-api/api/activity/v1"
	"github.com/patient-fyd/jxust-softhub-api/internal/model"
	"github.com/patient-fyd/jxust-softhub-api/internal/service"
)

// 获取活动列表
func (c *ControllerV1) List(ctx context.Context, req *v1.ListReq) (res *v1.ListRes, err error) {
	// 构建查询参数
	input := model.ActivityListInput{
		Page:     req.PageNum,
		PageSize: req.PageSize,
	}

	// 添加状态过滤条件
	if req.Status >= 0 {
		status := req.Status
		input.Status = &status
	}

	// 调用服务查询活动列表
	output, err := service.Activity().List(ctx, input)
	if err != nil {
		return nil, err
	}

	// 构建返回结果
	res = &v1.ListRes{
		List: make([]v1.ActivityInfo, 0),
		CommonPaginationRes: api.CommonPaginationRes{
			Total:    output.Total,
			PageNum:  output.Page,
			PageSize: output.PageSize,
		},
	}

	// 转换活动数据
	for _, item := range output.List {
		res.List = append(res.List, v1.ActivityInfo{
			ActivityId:      item.ActivityId,
			Title:           item.Title,
			Description:     item.Description,
			StartTime:       item.StartTime.Format("2006-01-02 15:04:05"),
			EndTime:         item.EndTime.Format("2006-01-02 15:04:05"),
			Location:        item.Location,
			MaxParticipants: item.MaxParticipants,
			Status:          item.Status,
			CreateTime:      item.CreateTime.Format("2006-01-02 15:04:05"),
		})
	}

	return res, nil
}

package activity

import (
	"context"

	"github.com/patient-fyd/jxust-softhub-api/api"
	v1 "github.com/patient-fyd/jxust-softhub-api/api/activity/v1"
	"github.com/patient-fyd/jxust-softhub-api/internal/model"
	"github.com/patient-fyd/jxust-softhub-api/internal/service"
)

// 获取报名列表
func (c *ControllerV1) RegisterList(ctx context.Context, req *v1.RegisterListReq) (res *v1.RegisterListRes, err error) {
	// 构建查询参数
	input := model.RegistrationListInput{
		ActivityId: req.ActivityId,
		Page:       req.PageNum,
		PageSize:   req.PageSize,
	}

	// 添加状态过滤条件
	if req.Status >= 0 {
		status := req.Status
		input.Status = &status
	}

	// 调用服务查询报名列表
	output, err := service.Activity().RegisterList(ctx, input)
	if err != nil {
		return nil, err
	}

	// 构建返回结果
	res = &v1.RegisterListRes{
		List: make([]v1.RegistrationInfo, 0),
		CommonPaginationRes: api.CommonPaginationRes{
			Total:    output.Total,
			PageNum:  output.Page,
			PageSize: output.PageSize,
		},
	}

	// 转换报名数据
	for _, item := range output.List {
		res.List = append(res.List, v1.RegistrationInfo{
			RegistrationId: item.RegistrationId,
			ActivityId:     item.ActivityId,
			UserId:         item.UserId,
			Name:           item.Name,
			StudentId:      item.StudentId,
			Contact:        item.Contact,
			Status:         item.Status,
			CreateTime:     item.CreateTime.Format("2006-01-02 15:04:05"),
		})
	}

	return res, nil
}

package join

import (
	"context"

	v1 "github.com/patient-fyd/jxust-softhub-api/api/join/v1"
	"github.com/patient-fyd/jxust-softhub-api/internal/model"
	"github.com/patient-fyd/jxust-softhub-api/internal/service"
)

func (c *ControllerV1) ApplicationList(ctx context.Context, req *v1.ApplicationListReq) (res *v1.ApplicationListRes, err error) {
	// 调用服务层方法获取申请列表
	result, err := service.Join().ApplicationList(ctx, model.JoinApplicationListInput{
		Status:           req.Status,
		Grade:            req.Grade,
		ExpectDepartment: req.ExpectDepartment,
		Page:             req.Page,
		PageSize:         req.PageSize,
	})
	if err != nil {
		return nil, err
	}

	// 构造响应数据
	res = &v1.ApplicationListRes{
		List:     make([]v1.ApplicationInfo, 0, len(result.List)),
		Total:    result.Total,
		Page:     result.Page,
		PageSize: result.PageSize,
	}

	// 转换数据格式
	for _, item := range result.List {
		res.List = append(res.List, v1.ApplicationInfo{
			ApplicationId:    item.ApplicationId,
			Name:             item.Name,
			StudentId:        item.StudentId,
			Grade:            item.Grade,
			College:          item.College,
			Major:            item.Major,
			Phone:            item.Phone,
			Email:            item.Email,
			Reason:           item.Reason,
			Skills:           item.Skills,
			ExpectDepartment: item.ExpectDepartment,
			Status:           item.Status,
			ReviewerId:       item.ReviewerId,
			ReviewerName:     item.ReviewerName,
			ReviewComment:    item.ReviewComment,
		})
	}

	return res, nil
}

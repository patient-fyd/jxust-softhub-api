package join

import (
	"context"

	v1 "github.com/patient-fyd/jxust-softhub-api/api/join/v1"
	"github.com/patient-fyd/jxust-softhub-api/internal/model"
	"github.com/patient-fyd/jxust-softhub-api/internal/service"
)

func (c *ControllerV1) Apply(ctx context.Context, req *v1.ApplyReq) (res *v1.ApplyRes, err error) {
	// 转换请求参数为业务层输入
	input := model.JoinApplyInput{
		Name:             req.Name,
		StudentId:        req.StudentId,
		Grade:            req.Grade,
		College:          req.College,
		Major:            req.Major,
		Phone:            req.Phone,
		Email:            req.Email,
		Reason:           req.Reason,
		Skills:           req.Skills,
		ExpectDepartment: req.ExpectDepartment,
	}

	// 调用业务层提交申请
	output, err := service.Join().Apply(ctx, input)
	if err != nil {
		return nil, err
	}

	// 返回结果
	return &v1.ApplyRes{
		ApplicationId: output.ApplicationId,
	}, nil
}

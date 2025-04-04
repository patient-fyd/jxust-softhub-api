package join

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	v1 "github.com/patient-fyd/jxust-softhub-api/api/join/v1"
	"github.com/patient-fyd/jxust-softhub-api/internal/model"
	"github.com/patient-fyd/jxust-softhub-api/internal/service"
)

func (c *ControllerV1) Detail(ctx context.Context, req *v1.DetailReq) (res *v1.DetailRes, err error) {
	// 构造业务参数
	input := model.JoinDetailInput{
		ApplicationId: req.ApplicationId,
	}

	// 调用服务
	output, err := service.Join().Detail(ctx, input)
	if err != nil {
		return nil, gerror.NewCode(gcode.CodeInternalError, err.Error())
	}

	// 构造响应
	res = &v1.DetailRes{
		ApplicationInfo: v1.ApplicationInfo{
			ApplicationId:    output.ApplicationId,
			Name:             output.Name,
			StudentId:        output.StudentId,
			Grade:            output.Grade,
			College:          output.College,
			Major:            output.Major,
			Phone:            output.Phone,
			Email:            output.Email,
			Reason:           output.Reason,
			Skills:           output.Skills,
			ExpectDepartment: output.ExpectDepartment,
			Status:           output.Status,
			ReviewerId:       output.ReviewerId,
			ReviewerName:     output.ReviewerName,
			ReviewComment:    output.ReviewComment,
			CreateTime:       output.CreateTime.String(),
			UpdateTime:       output.UpdateTime.String(),
		},
	}

	return res, nil
}

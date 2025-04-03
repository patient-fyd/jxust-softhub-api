package user

import (
	"context"

	v1 "github.com/patient-fyd/jxust-softhub-api/api/user/v1"
	"github.com/patient-fyd/jxust-softhub-api/internal/model"
	"github.com/patient-fyd/jxust-softhub-api/internal/service"
)

// AssignRole 为用户分配角色
func (c *ControllerV1) AssignRole(ctx context.Context, req *v1.AssignRoleReq) (res *v1.AssignRoleRes, err error) {
	// 构造业务逻辑输入参数
	input := model.UserAssignRoleInput{
		UserId: req.UserId,
		RoleId: req.RoleId,
	}

	// 调用服务层接口
	output, err := service.User().AssignRole(ctx, input)
	if err != nil {
		return nil, err
	}

	// 返回结果
	return &v1.AssignRoleRes{
		Success: output.Success,
	}, nil
}

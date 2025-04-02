package config

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"github.com/patient-fyd/jxust-softhub-api/api/config/v1"
)

func (c *ControllerV1) Set(ctx context.Context, req *v1.SetReq) (res *v1.SetRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

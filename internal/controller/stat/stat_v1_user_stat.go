package stat

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"github.com/patient-fyd/jxust-softhub-api/api/stat/v1"
)

func (c *ControllerV1) UserStat(ctx context.Context, req *v1.UserStatReq) (res *v1.UserStatRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

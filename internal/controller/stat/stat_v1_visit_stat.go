package stat

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"github.com/patient-fyd/jxust-softhub-api/api/stat/v1"
)

func (c *ControllerV1) VisitStat(ctx context.Context, req *v1.VisitStatReq) (res *v1.VisitStatRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

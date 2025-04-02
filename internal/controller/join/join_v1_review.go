package join

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"github.com/patient-fyd/jxust-softhub-api/api/join/v1"
)

func (c *ControllerV1) Review(ctx context.Context, req *v1.ReviewReq) (res *v1.ReviewRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

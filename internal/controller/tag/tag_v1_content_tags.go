package tag

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"github.com/patient-fyd/jxust-softhub-api/api/tag/v1"
)

func (c *ControllerV1) ContentTags(ctx context.Context, req *v1.ContentTagsReq) (res *v1.ContentTagsRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

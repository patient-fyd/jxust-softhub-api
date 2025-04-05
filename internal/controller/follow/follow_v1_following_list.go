package follow

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"github.com/patient-fyd/jxust-softhub-api/api/follow/v1"
)

func (c *ControllerV1) FollowingList(ctx context.Context, req *v1.FollowingListReq) (res *v1.FollowingListRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

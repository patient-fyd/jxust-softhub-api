package news

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	v1 "github.com/patient-fyd/jxust-softhub-api/api/news/v1"
	"github.com/patient-fyd/jxust-softhub-api/internal/model"
	"github.com/patient-fyd/jxust-softhub-api/internal/service"
)

// Delete 删除新闻
func (c *ControllerV1) Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error) {
	// 构造删除参数
	input := model.NewsDeleteInput{
		NewsId: req.NewsId,
	}

	// 调用服务
	output, err := service.News().Delete(ctx, input)
	if err != nil {
		return nil, gerror.NewCode(gcode.CodeInternalError, err.Error())
	}

	// 如果操作失败
	if !output.Success {
		return nil, gerror.NewCode(gcode.CodeNotFound, "新闻不存在")
	}

	// 转换响应
	res = &v1.DeleteRes{
		Success: output.Success,
	}

	return res, nil
}

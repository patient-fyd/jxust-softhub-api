package news

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	v1 "github.com/patient-fyd/jxust-softhub-api/api/news/v1"
	"github.com/patient-fyd/jxust-softhub-api/internal/model"
	"github.com/patient-fyd/jxust-softhub-api/internal/service"
)

// Update 更新新闻
func (c *ControllerV1) Update(ctx context.Context, req *v1.UpdateReq) (res *v1.UpdateRes, err error) {
	// 构造更新参数
	input := model.NewsUpdateInput{
		NewsId:     req.NewsId,
		Title:      req.Title,
		Content:    req.Content,
		Category:   req.Category,
		CoverImage: req.CoverImage,
		IsTop:      req.IsTop,
		Status:     req.Status,
	}

	// 调用服务
	output, err := service.News().Update(ctx, input)
	if err != nil {
		return nil, gerror.NewCode(gcode.CodeInternalError, err.Error())
	}

	// 如果操作失败
	if !output.Success {
		return nil, gerror.NewCode(gcode.CodeNotFound, "新闻不存在")
	}

	// 转换响应
	res = &v1.UpdateRes{
		Success: output.Success,
	}

	return res, nil
}

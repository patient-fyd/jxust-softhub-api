package news

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	v1 "github.com/patient-fyd/jxust-softhub-api/api/news/v1"
	"github.com/patient-fyd/jxust-softhub-api/internal/model"
	"github.com/patient-fyd/jxust-softhub-api/internal/service"
)

// Create 创建新闻
func (c *ControllerV1) Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error) {
	// 构造创建参数
	input := model.NewsCreateInput{
		Title:      req.Title,
		Content:    req.Content,
		Category:   req.Category,
		CoverImage: req.CoverImage,
		IsTop:      req.IsTop,
		Status:     req.Status,
	}

	// 调用服务
	output, err := service.News().Create(ctx, input)
	if err != nil {
		return nil, gerror.NewCode(gcode.CodeInternalError, err.Error())
	}

	// 转换响应
	res = &v1.CreateRes{
		Id: output.Id,
	}

	return res, nil
}

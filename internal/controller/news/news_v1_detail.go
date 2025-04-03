package news

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	v1 "github.com/patient-fyd/jxust-softhub-api/api/news/v1"
	"github.com/patient-fyd/jxust-softhub-api/internal/model"
	"github.com/patient-fyd/jxust-softhub-api/internal/service"
)

// Detail 获取新闻详情
func (c *ControllerV1) Detail(ctx context.Context, req *v1.DetailReq) (res *v1.DetailRes, err error) {
	// 构造查询参数
	input := model.NewsDetailInput{
		NewsId: req.NewsId,
	}

	// 调用服务
	output, err := service.News().Detail(ctx, input)
	if err != nil {
		return nil, gerror.NewCode(gcode.CodeInternalError, err.Error())
	}

	// 如果没有找到数据
	if output == nil {
		return nil, gerror.NewCode(gcode.CodeNotFound, "新闻不存在")
	}

	// 转换响应
	res = &v1.DetailRes{
		Id:         output.Id,
		Title:      output.Title,
		Content:    output.Content,
		Category:   output.Category,
		CoverImage: output.CoverImage,
		ViewCount:  output.ViewCount,
		IsTop:      0, // 暂不实现置顶功能
		Status:     output.Status,
		CreatedAt:  output.CreatedAt,
		UpdatedAt:  output.UpdatedAt,
	}

	return res, nil
}

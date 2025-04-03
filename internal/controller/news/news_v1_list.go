package news

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	v1 "github.com/patient-fyd/jxust-softhub-api/api/news/v1"
	"github.com/patient-fyd/jxust-softhub-api/internal/model"
	"github.com/patient-fyd/jxust-softhub-api/internal/service"
)

// List 获取新闻列表
func (c *ControllerV1) List(ctx context.Context, req *v1.ListReq) (res *v1.ListRes, err error) {
	// 参数验证
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 10
	}

	// 构造查询参数
	input := model.NewsListInput{
		Category: req.Category,
		Page:     req.Page,
		PageSize: req.PageSize,
	}

	// 调用服务
	output, err := service.News().List(ctx, input)
	if err != nil {
		return nil, gerror.NewCode(gcode.CodeInternalError, err.Error())
	}

	// 转换响应
	res = &v1.ListRes{
		Total:    output.Total,
		Page:     output.Page,
		PageSize: output.PageSize,
		List:     make([]v1.NewsInfo, len(output.List)),
	}

	// 填充列表数据
	for i, item := range output.List {
		res.List[i] = v1.NewsInfo{
			Id:         item.Id,
			Title:      item.Title,
			Category:   item.Category,
			CoverImage: item.CoverImage,
			ViewCount:  item.ViewCount,
			IsTop:      0, // 暂不实现置顶功能
			Status:     item.Status,
			CreatedAt:  item.CreatedAt,
			UpdatedAt:  item.UpdatedAt,
		}
	}

	return res, nil
}

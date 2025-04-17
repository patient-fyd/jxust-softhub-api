package blog

import (
	"context"

	v1 "github.com/patient-fyd/jxust-softhub-api/api/blog/v1"
	"github.com/patient-fyd/jxust-softhub-api/internal/model"
	"github.com/patient-fyd/jxust-softhub-api/internal/service"
)

func (c *ControllerV1) BlogList(ctx context.Context, req *v1.BlogListReq) (res *v1.BlogListRes, err error) {
	// 构建查询参数
	input := model.BlogListInput{
		Page:        req.Page,
		Size:        req.Size,
		Category:    req.Category,
		Tag:         req.Tag,
		Keyword:     req.Keyword,
		AuthorId:    req.AuthorId,
		Status:      req.Status,
		IsRecommend: req.IsRecommend,
	}

	// 调用Service层获取博客列表
	output, err := service.Blog().GetList(ctx, input)
	if err != nil {
		return nil, err
	}

	// 构建响应数据
	res = &v1.BlogListRes{
		List:  output.List,
		Page:  output.Page,
		Size:  output.Size,
		Total: output.Total,
	}

	return res, nil
}

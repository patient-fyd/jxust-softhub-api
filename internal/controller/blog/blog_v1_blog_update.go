package blog

import (
	"context"

	v1 "github.com/patient-fyd/jxust-softhub-api/api/blog/v1"
	"github.com/patient-fyd/jxust-softhub-api/internal/model"
	"github.com/patient-fyd/jxust-softhub-api/internal/service"
)

func (c *ControllerV1) BlogUpdate(ctx context.Context, req *v1.BlogUpdateReq) (res *v1.BlogUpdateRes, err error) {
	// 构建更新参数
	input := model.BlogUpdateInput{
		BlogId:     req.BlogId,
		Title:      req.Title,
		Content:    req.Content,
		Summary:    req.Summary,
		Category:   req.Category,
		Tags:       req.Tags,
		CoverImage: req.CoverImage,
		Status:     req.Status,
	}

	// 调用Service层更新博客
	err = service.Blog().Update(ctx, input)
	if err != nil {
		return nil, err
	}

	// 构建响应数据
	res = &v1.BlogUpdateRes{}

	return res, nil
}

package blog

import (
	"context"

	v1 "github.com/patient-fyd/jxust-softhub-api/api/blog/v1"
	"github.com/patient-fyd/jxust-softhub-api/internal/model"
	"github.com/patient-fyd/jxust-softhub-api/internal/service"
)

func (c *ControllerV1) BlogCreate(ctx context.Context, req *v1.BlogCreateReq) (res *v1.BlogCreateRes, err error) {
	// 构建创建参数
	input := model.BlogCreateInput{
		Title:      req.Title,
		Content:    req.Content,
		Summary:    req.Summary,
		Category:   req.Category,
		Tags:       req.Tags,
		CoverImage: req.CoverImage,
		Status:     req.Status,
	}

	// 调用Service层创建博客
	output, err := service.Blog().Create(ctx, input)
	if err != nil {
		return nil, err
	}

	// 构建响应数据
	res = &v1.BlogCreateRes{
		BlogId: output.BlogId,
	}

	return res, nil
}

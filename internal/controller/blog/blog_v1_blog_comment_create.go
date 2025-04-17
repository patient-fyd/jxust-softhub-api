package blog

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	v1 "github.com/patient-fyd/jxust-softhub-api/api/blog/v1"
	"github.com/patient-fyd/jxust-softhub-api/internal/model"
	"github.com/patient-fyd/jxust-softhub-api/internal/service"
)

func (c *ControllerV1) BlogCommentCreate(ctx context.Context, req *v1.BlogCommentCreateReq) (res *v1.BlogCommentCreateRes, err error) {
	// 构建创建参数
	input := model.BlogCommentCreateInput{
		BlogId:   req.BlogId,
		Content:  req.Content,
		ParentId: req.ParentId,
	}

	// 调用Service层创建博客评论
	output, err := service.Blog().CreateComment(ctx, input)
	if err != nil {
		g.Log().Error(ctx, "创建博客评论失败:", err)
		return nil, err
	}

	// 构建响应数据
	res = &v1.BlogCommentCreateRes{
		CommentId: output.CommentId,
	}

	return res, nil
}

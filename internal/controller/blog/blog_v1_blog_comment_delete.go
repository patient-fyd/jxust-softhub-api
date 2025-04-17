package blog

import (
	"context"

	v1 "github.com/patient-fyd/jxust-softhub-api/api/blog/v1"
	"github.com/patient-fyd/jxust-softhub-api/internal/model"
	"github.com/patient-fyd/jxust-softhub-api/internal/service"
)

func (c *ControllerV1) BlogCommentDelete(ctx context.Context, req *v1.BlogCommentDeleteReq) (res *v1.BlogCommentDeleteRes, err error) {
	// 构建删除参数
	input := model.BlogCommentDeleteInput{
		CommentId: req.CommentId,
	}

	// 调用Service层删除博客评论
	err = service.Blog().DeleteComment(ctx, input)
	if err != nil {
		return nil, err
	}

	// 构建响应数据
	res = &v1.BlogCommentDeleteRes{}

	return res, nil
}

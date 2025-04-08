package post

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"

	v1 "github.com/patient-fyd/jxust-softhub-api/api/post/v1"
	"github.com/patient-fyd/jxust-softhub-api/internal/model"
	"github.com/patient-fyd/jxust-softhub-api/internal/service"
)

func (c *ControllerV1) Detail(ctx context.Context, req *v1.DetailReq) (res *v1.DetailRes, err error) {
	// 转换请求参数
	input := model.PostDetailInput{
		PostId: req.PostId,
	}

	// 调用业务逻辑层
	output, err := service.Post().Detail(ctx, input)
	if err != nil {
		return nil, gerror.Wrap(err, "获取帖子详情失败")
	}

	// 转换响应参数
	return &v1.DetailRes{
		PostId:       output.PostId,
		UserId:       output.UserId,
		UserName:     output.UserName,
		UserAvatar:   output.UserAvatar,
		Content:      output.Content,
		Images:       output.Images,
		CircleId:     output.CircleId,
		CircleName:   output.CircleName,
		TopicId:      output.TopicId,
		TopicName:    output.TopicName,
		ViewCount:    output.ViewCount,
		LikeCount:    output.LikeCount,
		CommentCount: output.CommentCount,
		ShareCount:   output.ShareCount,
		IsTop:        output.IsTop,
		IsLiked:      output.IsLiked,
		CreateTime:   output.CreateTime.String(),
		UpdateTime:   output.UpdateTime.String(),
	}, nil
}

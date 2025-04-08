package post

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"

	v1 "github.com/patient-fyd/jxust-softhub-api/api/post/v1"
	"github.com/patient-fyd/jxust-softhub-api/internal/model"
	"github.com/patient-fyd/jxust-softhub-api/internal/service"
)

func (c *ControllerV1) List(ctx context.Context, req *v1.ListReq) (res *v1.ListRes, err error) {
	// 转换请求参数
	input := model.PostListInput{
		Page:     req.Page,
		Size:     req.Size,
		Keyword:  req.Keyword,
		CircleId: req.CircleId,
		TopicId:  req.TopicId,
		UserId:   req.UserId,
		OrderBy:  req.OrderBy,
	}

	// 调用业务逻辑层
	output, err := service.Post().List(ctx, input)
	if err != nil {
		return nil, gerror.Wrap(err, "获取帖子列表失败")
	}

	// 转换响应参数
	res = &v1.ListRes{
		Total: output.Total,
		Page:  output.Page,
		Size:  output.Size,
		List:  make([]v1.PostItem, len(output.List)),
	}

	// 将model.PostListItem转换为v1.PostItem
	for i, item := range output.List {
		res.List[i] = v1.PostItem{
			PostId:       item.PostId,
			UserId:       item.UserId,
			UserName:     item.UserName,
			UserAvatar:   item.UserAvatar,
			Content:      item.Content,
			Images:       item.Images,
			CircleId:     item.CircleId,
			CircleName:   item.CircleName,
			TopicId:      item.TopicId,
			TopicName:    item.TopicName,
			ViewCount:    item.ViewCount,
			LikeCount:    item.LikeCount,
			CommentCount: item.CommentCount,
			ShareCount:   item.ShareCount,
			IsTop:        item.IsTop,
			IsLiked:      item.IsLiked,
			CreateTime:   item.CreateTime.String(),
		}
	}

	return res, nil
}

package service

import (
	"context"

	"github.com/patient-fyd/jxust-softhub-api/internal/model"
)

// IComment 评论服务接口
type IComment interface {
	// List 获取评论列表
	List(ctx context.Context, in model.CommentListInput) (*model.CommentListOutput, error)
	// Create 创建评论
	Create(ctx context.Context, in model.CommentCreateInput) (*model.CommentCreateOutput, error)
	// Delete 删除评论
	Delete(ctx context.Context, in model.CommentDeleteInput) (*model.CommentDeleteOutput, error)
}

var (
	localComment IComment
)

// Comment 获取评论服务实例
func Comment() IComment {
	if localComment == nil {
		panic("implement not found for interface IComment, forgot register?")
	}
	return localComment
}

// RegisterComment 注册评论服务实现
func RegisterComment(i IComment) {
	localComment = i
}

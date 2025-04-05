package service

import (
	"context"

	"github.com/patient-fyd/jxust-softhub-api/internal/model"
)

// IPost 帖子服务接口
type IPost interface {
	// List 获取帖子列表
	List(ctx context.Context, in model.PostListInput) (*model.PostListOutput, error)
	// Detail 获取帖子详情
	Detail(ctx context.Context, in model.PostDetailInput) (*model.PostDetailOutput, error)
	// Create 创建帖子
	Create(ctx context.Context, in model.PostCreateInput) (*model.PostCreateOutput, error)
	// Delete 删除帖子
	Delete(ctx context.Context, in model.PostDeleteInput) (*model.PostDeleteOutput, error)
}

// 获取帖子服务
func Post() IPost {
	return localPost
}

// 注册帖子服务
func RegisterPost(i IPost) {
	localPost = i
}

var localPost IPost

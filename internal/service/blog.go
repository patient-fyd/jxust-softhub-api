// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"

	"github.com/patient-fyd/jxust-softhub-api/internal/model"
)

type (
	IBlog interface {
		// 获取博客列表
		GetList(ctx context.Context, in model.BlogListInput) (out *model.BlogListOutput, err error)
		// 获取博客详情
		GetDetail(ctx context.Context, in model.BlogDetailInput) (out *model.BlogDetailOutput, err error)
		// 创建博客
		Create(ctx context.Context, in model.BlogCreateInput) (out *model.BlogCreateOutput, err error)
		// 更新博客
		Update(ctx context.Context, in model.BlogUpdateInput) error
		// 删除博客
		Delete(ctx context.Context, in model.BlogDeleteInput) error
		// 设置博客推荐状态
		SetRecommend(ctx context.Context, in model.BlogRecommendInput) error
		// 点赞博客
		Like(ctx context.Context, in model.BlogLikeInput) error
		// 取消点赞博客
		Unlike(ctx context.Context, in model.BlogUnlikeInput) error
		// 获取博客评论列表
		GetCommentList(ctx context.Context, in model.BlogCommentListInput) (out *model.BlogCommentListOutput, err error)
		// 创建博客评论
		CreateComment(ctx context.Context, in model.BlogCommentCreateInput) (out *model.BlogCommentCreateOutput, err error)
		// 删除博客评论
		DeleteComment(ctx context.Context, in model.BlogCommentDeleteInput) error
		// 获取博客分类列表
		GetCategoryList(ctx context.Context) (out *model.BlogCategoryListOutput, err error)
		// 获取博客标签列表
		GetTagList(ctx context.Context) (out *model.BlogTagListOutput, err error)
	}
)

var (
	localBlog IBlog
)

func Blog() IBlog {
	if localBlog == nil {
		panic("implement not found for interface IBlog, forgot register?")
	}
	return localBlog
}

func RegisterBlog(i IBlog) {
	localBlog = i
}

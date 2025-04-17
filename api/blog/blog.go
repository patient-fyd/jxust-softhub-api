package blog

import (
	"context"

	v1 "github.com/patient-fyd/jxust-softhub-api/api/blog/v1"
)

// IBlog 博客管理API接口
type IBlog interface {
	// 博客管理
	BlogList(ctx context.Context, req *v1.BlogListReq) (res *v1.BlogListRes, err error)
	BlogDetail(ctx context.Context, req *v1.BlogDetailReq) (res *v1.BlogDetailRes, err error)
	BlogCreate(ctx context.Context, req *v1.BlogCreateReq) (res *v1.BlogCreateRes, err error)
	BlogUpdate(ctx context.Context, req *v1.BlogUpdateReq) (res *v1.BlogUpdateRes, err error)
	BlogDelete(ctx context.Context, req *v1.BlogDeleteReq) (res *v1.BlogDeleteRes, err error)
	BlogRecommend(ctx context.Context, req *v1.BlogRecommendReq) (res *v1.BlogRecommendRes, err error)
	BlogLike(ctx context.Context, req *v1.BlogLikeReq) (res *v1.BlogLikeRes, err error)
	BlogUnlike(ctx context.Context, req *v1.BlogUnlikeReq) (res *v1.BlogUnlikeRes, err error)

	// 博客评论
	BlogCommentList(ctx context.Context, req *v1.BlogCommentListReq) (res *v1.BlogCommentListRes, err error)
	BlogCommentCreate(ctx context.Context, req *v1.BlogCommentCreateReq) (res *v1.BlogCommentCreateRes, err error)
	BlogCommentDelete(ctx context.Context, req *v1.BlogCommentDeleteReq) (res *v1.BlogCommentDeleteRes, err error)

	// 博客分类和标签
	BlogCategoryList(ctx context.Context, req *v1.BlogCategoryListReq) (res *v1.BlogCategoryListRes, err error)
	BlogTagList(ctx context.Context, req *v1.BlogTagListReq) (res *v1.BlogTagListRes, err error)
}

// IBlogV1 博客管理API v1版本接口
type IBlogV1 interface {
	IBlog
}

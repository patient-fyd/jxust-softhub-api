// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. 
// =================================================================================

package sdk

import (
	"context"

	"github.com/gogf/gf/contrib/sdk/httpclient/v2"
	"github.com/gogf/gf/v2/text/gstr"

	"github.com/patient-fyd/jxust-softhub-api/api/blog"
	"github.com/patient-fyd/jxust-softhub-api/api/blog/v1"
)

type implementerBlogV1 struct {
	*httpclient.Client
}

func (i *implementer) BlogV1() blog.IBlogV1 {
	var (
		client = httpclient.New(i.config)
		prefix = gstr.TrimRight(i.config.URL, "/") + ""
	)
	client.Client = client.Prefix(prefix)
	return &implementerBlogV1{client}
}

func (i *implementerBlogV1) BlogList(ctx context.Context, req *v1.BlogListReq) (res *v1.BlogListRes, err error) {
	err = i.Request(ctx, req, &res)
	return
}

func (i *implementerBlogV1) BlogDetail(ctx context.Context, req *v1.BlogDetailReq) (res *v1.BlogDetailRes, err error) {
	err = i.Request(ctx, req, &res)
	return
}

func (i *implementerBlogV1) BlogCreate(ctx context.Context, req *v1.BlogCreateReq) (res *v1.BlogCreateRes, err error) {
	err = i.Request(ctx, req, &res)
	return
}

func (i *implementerBlogV1) BlogUpdate(ctx context.Context, req *v1.BlogUpdateReq) (res *v1.BlogUpdateRes, err error) {
	err = i.Request(ctx, req, &res)
	return
}

func (i *implementerBlogV1) BlogDelete(ctx context.Context, req *v1.BlogDeleteReq) (res *v1.BlogDeleteRes, err error) {
	err = i.Request(ctx, req, &res)
	return
}

func (i *implementerBlogV1) BlogRecommend(ctx context.Context, req *v1.BlogRecommendReq) (res *v1.BlogRecommendRes, err error) {
	err = i.Request(ctx, req, &res)
	return
}

func (i *implementerBlogV1) BlogLike(ctx context.Context, req *v1.BlogLikeReq) (res *v1.BlogLikeRes, err error) {
	err = i.Request(ctx, req, &res)
	return
}

func (i *implementerBlogV1) BlogUnlike(ctx context.Context, req *v1.BlogUnlikeReq) (res *v1.BlogUnlikeRes, err error) {
	err = i.Request(ctx, req, &res)
	return
}

func (i *implementerBlogV1) BlogCommentList(ctx context.Context, req *v1.BlogCommentListReq) (res *v1.BlogCommentListRes, err error) {
	err = i.Request(ctx, req, &res)
	return
}

func (i *implementerBlogV1) BlogCommentCreate(ctx context.Context, req *v1.BlogCommentCreateReq) (res *v1.BlogCommentCreateRes, err error) {
	err = i.Request(ctx, req, &res)
	return
}

func (i *implementerBlogV1) BlogCommentDelete(ctx context.Context, req *v1.BlogCommentDeleteReq) (res *v1.BlogCommentDeleteRes, err error) {
	err = i.Request(ctx, req, &res)
	return
}

func (i *implementerBlogV1) BlogCategoryList(ctx context.Context, req *v1.BlogCategoryListReq) (res *v1.BlogCategoryListRes, err error) {
	err = i.Request(ctx, req, &res)
	return
}

func (i *implementerBlogV1) BlogTagList(ctx context.Context, req *v1.BlogTagListReq) (res *v1.BlogTagListRes, err error) {
	err = i.Request(ctx, req, &res)
	return
}


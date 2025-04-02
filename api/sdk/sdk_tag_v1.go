// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. 
// =================================================================================

package sdk

import (
	"context"

	"github.com/gogf/gf/contrib/sdk/httpclient/v2"
	"github.com/gogf/gf/v2/text/gstr"

	"github.com/patient-fyd/jxust-softhub-api/api/tag"
	"github.com/patient-fyd/jxust-softhub-api/api/tag/v1"
)

type implementerTagV1 struct {
	*httpclient.Client
}

func (i *implementer) TagV1() tag.ITagV1 {
	var (
		client = httpclient.New(i.config)
		prefix = gstr.TrimRight(i.config.URL, "/") + ""
	)
	client.Client = client.Prefix(prefix)
	return &implementerTagV1{client}
}

func (i *implementerTagV1) List(ctx context.Context, req *v1.ListReq) (res *v1.ListRes, err error) {
	err = i.Request(ctx, req, &res)
	return
}

func (i *implementerTagV1) Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error) {
	err = i.Request(ctx, req, &res)
	return
}

func (i *implementerTagV1) Update(ctx context.Context, req *v1.UpdateReq) (res *v1.UpdateRes, err error) {
	err = i.Request(ctx, req, &res)
	return
}

func (i *implementerTagV1) Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error) {
	err = i.Request(ctx, req, &res)
	return
}

func (i *implementerTagV1) ContentTags(ctx context.Context, req *v1.ContentTagsReq) (res *v1.ContentTagsRes, err error) {
	err = i.Request(ctx, req, &res)
	return
}

func (i *implementerTagV1) SetContentTags(ctx context.Context, req *v1.SetContentTagsReq) (res *v1.SetContentTagsRes, err error) {
	err = i.Request(ctx, req, &res)
	return
}


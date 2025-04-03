package service

import (
	"context"

	"github.com/patient-fyd/jxust-softhub-api/internal/model"
)

// INews 新闻服务接口
type INews interface {
	// List 获取新闻列表
	List(ctx context.Context, in model.NewsListInput) (*model.NewsListOutput, error)
	// Detail 获取新闻详情
	Detail(ctx context.Context, in model.NewsDetailInput) (*model.NewsDetailOutput, error)
	// Create 创建新闻
	Create(ctx context.Context, in model.NewsCreateInput) (*model.NewsCreateOutput, error)
	// Update 更新新闻
	Update(ctx context.Context, in model.NewsUpdateInput) (*model.NewsUpdateOutput, error)
	// Delete 删除新闻
	Delete(ctx context.Context, in model.NewsDeleteInput) (*model.NewsDeleteOutput, error)
}

// 获取新闻服务
func News() INews {
	return localNews
}

// 注册新闻服务
func RegisterNews(i INews) {
	localNews = i
}

var localNews INews

package news

import (
	"context"

	"github.com/gogf/gf/v2/os/gtime"

	"github.com/patient-fyd/jxust-softhub-api/internal/dao"
	"github.com/patient-fyd/jxust-softhub-api/internal/model"
	"github.com/patient-fyd/jxust-softhub-api/internal/model/do"
	"github.com/patient-fyd/jxust-softhub-api/internal/model/entity"
	"github.com/patient-fyd/jxust-softhub-api/internal/service"
)

type sNews struct{}

func init() {
	service.RegisterNews(New())
}

func New() *sNews {
	return &sNews{}
}

// List 获取新闻列表
func (s *sNews) List(ctx context.Context, in model.NewsListInput) (*model.NewsListOutput, error) {
	var (
		m         = dao.News.Ctx(ctx)
		condition = do.News{}
		out       = &model.NewsListOutput{
			Page:     in.Page,
			PageSize: in.PageSize,
		}
	)

	if in.Page <= 0 {
		in.Page = 1
	}
	if in.PageSize <= 0 {
		in.PageSize = 10
	}

	// 添加分类条件
	if in.Category != "" {
		condition.Category = in.Category
	}

	// 状态默认为已发布
	condition.Status = 1

	// 查询符合条件的记录总数
	count, err := m.Where(condition).Count()
	if err != nil {
		return nil, err
	}
	out.Total = count

	// 没有数据，直接返回
	if count == 0 {
		return out, nil
	}

	// 查询列表数据
	var list []*entity.News
	err = m.Where(condition).
		Page(in.Page, in.PageSize).
		Order("createTime DESC").
		Scan(&list)
	if err != nil {
		return nil, err
	}

	// 转换结果格式
	out.List = make([]model.NewsInfo, len(list))
	for i, v := range list {
		out.List[i] = model.NewsInfo{
			Id:         v.NewsId,
			Title:      v.Title,
			Category:   v.Category,
			CoverImage: v.CoverImage,
			ViewCount:  uint(v.ViewCount),
			Status:     uint8(v.Status),
			CreatedAt:  v.CreateTime.String(),
			UpdatedAt:  v.UpdateTime.String(),
		}
	}

	return out, nil
}

// Detail 获取新闻详情
func (s *sNews) Detail(ctx context.Context, in model.NewsDetailInput) (*model.NewsDetailOutput, error) {
	var (
		m = dao.News.Ctx(ctx)
	)

	// 查询新闻详情
	news, err := m.Where(do.News{
		NewsId: in.NewsId,
	}).One()
	if err != nil {
		return nil, err
	}
	if news.IsEmpty() {
		return nil, nil
	}

	var result entity.News
	err = news.Struct(&result)
	if err != nil {
		return nil, err
	}

	// 增加浏览次数
	_, err = m.Where(do.News{
		NewsId: in.NewsId,
	}).Increment("viewCount", 1)
	if err != nil {
		return nil, err
	}

	return &model.NewsDetailOutput{
		Id:         result.NewsId,
		Title:      result.Title,
		Content:    result.Content,
		Category:   result.Category,
		CoverImage: result.CoverImage,
		ViewCount:  uint(result.ViewCount + 1), // +1 是因为刚增加的查看次数
		Status:     uint8(result.Status),
		CreatedAt:  result.CreateTime.String(),
		UpdatedAt:  result.UpdateTime.String(),
	}, nil
}

// Create 创建新闻
func (s *sNews) Create(ctx context.Context, in model.NewsCreateInput) (*model.NewsCreateOutput, error) {
	var (
		m = dao.News.Ctx(ctx)
	)

	// 从上下文获取当前用户ID
	userId := uint(1) // 暂时默认为1，后面从上下文获取

	// 创建新闻记录
	result, err := m.Insert(do.News{
		Title:      in.Title,
		Content:    in.Content,
		Category:   in.Category,
		CoverImage: in.CoverImage,
		AuthorId:   userId,
		ViewCount:  0,
		Status:     int(in.Status),
		CreateTime: gtime.Now(),
		UpdateTime: gtime.Now(),
	})
	if err != nil {
		return nil, err
	}

	// 获取创建的新闻ID
	newsId, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	return &model.NewsCreateOutput{
		Id: uint(newsId),
	}, nil
}

// Update 更新新闻
func (s *sNews) Update(ctx context.Context, in model.NewsUpdateInput) (*model.NewsUpdateOutput, error) {
	var (
		m = dao.News.Ctx(ctx)
	)

	// 检查新闻是否存在
	news, err := m.Where(do.News{
		NewsId: in.NewsId,
	}).One()
	if err != nil {
		return nil, err
	}
	if news.IsEmpty() {
		return &model.NewsUpdateOutput{Success: false}, nil
	}

	// 构建更新数据
	updateData := do.News{
		UpdateTime: gtime.Now(),
	}

	if in.Title != "" {
		updateData.Title = in.Title
	}
	if in.Content != "" {
		updateData.Content = in.Content
	}
	if in.Category != "" {
		updateData.Category = in.Category
	}
	if in.CoverImage != "" {
		updateData.CoverImage = in.CoverImage
	}

	updateData.Status = int(in.Status)

	// 更新新闻记录
	_, err = m.Where(do.News{
		NewsId: in.NewsId,
	}).Update(updateData)
	if err != nil {
		return nil, err
	}

	return &model.NewsUpdateOutput{
		Success: true,
	}, nil
}

// Delete 删除新闻
func (s *sNews) Delete(ctx context.Context, in model.NewsDeleteInput) (*model.NewsDeleteOutput, error) {
	var (
		m = dao.News.Ctx(ctx)
	)

	// 检查新闻是否存在
	news, err := m.Where(do.News{
		NewsId: in.NewsId,
	}).One()
	if err != nil {
		return nil, err
	}
	if news.IsEmpty() {
		return &model.NewsDeleteOutput{Success: false}, nil
	}

	// 删除新闻记录
	_, err = m.Where(do.News{
		NewsId: in.NewsId,
	}).Delete()
	if err != nil {
		return nil, err
	}

	return &model.NewsDeleteOutput{
		Success: true,
	}, nil
}

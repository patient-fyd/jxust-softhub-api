package blog

import (
	"context"
	"sort"
	"strings"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"

	"github.com/patient-fyd/jxust-softhub-api/internal/dao"
	"github.com/patient-fyd/jxust-softhub-api/internal/model"
	"github.com/patient-fyd/jxust-softhub-api/internal/model/entity"
	"github.com/patient-fyd/jxust-softhub-api/internal/service"
)

type sBlog struct{}

func New() *sBlog {
	return &sBlog{}
}

func init() {
	service.RegisterBlog(New())
}

// 获取博客列表
func (s *sBlog) GetList(ctx context.Context, in model.BlogListInput) (out *model.BlogListOutput, err error) {
	m := dao.Blogs.Ctx(ctx)

	// 1. 构建查询条件
	if in.Category != "" {
		m = m.Where(dao.Blogs.Columns().Category, in.Category)
	}

	// 2. 根据标签过滤
	if in.Tag != "" {
		m = m.Where(dao.Blogs.Columns().Tags+" LIKE ?", "%"+in.Tag+"%")
	}

	// 3. 关键词搜索
	if in.Keyword != "" {
		m = m.Where(
			"("+dao.Blogs.Columns().Title+" LIKE ? OR "+dao.Blogs.Columns().Summary+" LIKE ?)",
			"%"+in.Keyword+"%",
			"%"+in.Keyword+"%",
		)
	}

	// 4. 按作者过滤
	if in.AuthorId > 0 {
		m = m.Where(dao.Blogs.Columns().AuthorId, in.AuthorId)
	}

	// 5. 按状态过滤
	m = m.Where(dao.Blogs.Columns().Status, in.Status)

	// 6. 按推荐过滤
	if in.IsRecommend == 1 {
		m = m.Where(dao.Blogs.Columns().IsRecommend, 1)
	}

	// 7. 统计总数
	count, err := m.Count()
	if err != nil {
		return nil, err
	}

	out = &model.BlogListOutput{
		Page:  in.Page,
		Size:  in.Size,
		Total: count,
		List:  make([]*model.BlogListItem, 0),
	}

	// 8. 如果没有数据，直接返回
	if count == 0 {
		return out, nil
	}

	// 9. 查询数据
	var list []*entity.Blogs
	err = m.OrderDesc(dao.Blogs.Columns().CreateTime).
		Page(in.Page, in.Size).
		Scan(&list)
	if err != nil {
		return nil, err
	}

	// 10. 提取所有作者ID
	var authorIds []uint
	for _, v := range list {
		if v.AuthorId > 0 {
			authorIds = append(authorIds, v.AuthorId)
		}
	}

	// 11. 查询作者信息
	var authorMap map[uint]entity.Users
	if len(authorIds) > 0 {
		var authors []entity.Users
		err = dao.Users.Ctx(ctx).
			Fields(dao.Users.Columns().UserId, dao.Users.Columns().UserName).
			Where(dao.Users.Columns().UserId+" IN(?)", authorIds).
			Scan(&authors)
		if err != nil {
			return nil, err
		}

		authorMap = make(map[uint]entity.Users, len(authors))
		for _, v := range authors {
			authorMap[v.UserId] = v
		}
	}

	// 12. 转换数据
	for _, v := range list {
		item := &model.BlogListItem{
			BlogId:       v.BlogId,
			Title:        v.Title,
			Summary:      v.Summary,
			Category:     v.Category,
			Tags:         v.Tags,
			CoverImage:   v.CoverImage,
			AuthorId:     v.AuthorId,
			ViewCount:    v.ViewCount,
			LikeCount:    v.LikeCount,
			CommentCount: v.CommentCount,
			IsRecommend:  v.IsRecommend,
			Status:       v.Status,
			CreateTime:   v.CreateTime,
		}

		// 设置作者名称
		if author, ok := authorMap[v.AuthorId]; ok {
			item.AuthorName = author.UserName
		}

		out.List = append(out.List, item)
	}

	return out, nil
}

// 获取博客详情
func (s *sBlog) GetDetail(ctx context.Context, in model.BlogDetailInput) (out *model.BlogDetailOutput, err error) {
	var blog entity.Blogs
	err = dao.Blogs.Ctx(ctx).
		Where(dao.Blogs.Columns().BlogId, in.BlogId).
		Scan(&blog)
	if err != nil {
		return nil, err
	}

	// 文章不存在
	if blog.BlogId == 0 {
		return nil, nil
	}

	out = &model.BlogDetailOutput{
		Blogs: blog,
	}

	// 增加浏览次数
	dao.Blogs.Ctx(ctx).
		Data(g.Map{
			dao.Blogs.Columns().ViewCount: "view_count + 1",
		}).
		Where(dao.Blogs.Columns().BlogId, in.BlogId).
		Update()

	// 如果有作者，查询作者信息
	if blog.AuthorId > 0 {
		var author entity.Users
		err = dao.Users.Ctx(ctx).
			Fields(dao.Users.Columns().UserName, dao.Users.Columns().Avatar).
			Where(dao.Users.Columns().UserId, blog.AuthorId).
			Scan(&author)
		if err != nil {
			return nil, err
		}

		if author.UserId > 0 {
			out.AuthorName = author.UserName
			out.AuthorAvatar = author.Avatar
		}
	}

	// 检查当前用户是否已点赞
	userId := gconv.Uint(ctx.Value("userId"))
	if userId > 0 {
		count, err := dao.BlogLikes.Ctx(ctx).
			Where(dao.BlogLikes.Columns().BlogId, in.BlogId).
			Where(dao.BlogLikes.Columns().UserId, userId).
			Count()
		if err != nil {
			return nil, err
		}
		out.IsLiked = count > 0
	}

	return out, nil
}

// 创建博客
func (s *sBlog) Create(ctx context.Context, in model.BlogCreateInput) (out *model.BlogCreateOutput, err error) {
	// 1. 获取当前用户ID
	userId := gconv.Uint(ctx.Value("userId"))
	if userId == 0 {
		return nil, gerror.New("用户未登录")
	}

	// 2. 如果未提供摘要，自动截取内容前150个字符作为摘要
	if in.Summary == "" && in.Content != "" {
		// 去除Markdown标记和空白字符
		content := gstr.Replace(in.Content, "\n", " ")
		content = gstr.Replace(content, "\r", " ")
		content = gstr.Replace(content, "#", "")
		content = gstr.Replace(content, "*", "")
		content = gstr.Replace(content, "```", "")
		content = strings.TrimSpace(content)

		// 截取前150个字符作为摘要
		if len(content) > 150 {
			in.Summary = content[:150] + "..."
		} else {
			in.Summary = content
		}
	}

	// 3. 插入数据
	data := entity.Blogs{
		Title:      in.Title,
		Content:    in.Content,
		Summary:    in.Summary,
		Category:   in.Category,
		Tags:       in.Tags,
		CoverImage: in.CoverImage,
		AuthorId:   userId,
		Status:     in.Status,
	}

	id, err := dao.Blogs.Ctx(ctx).Data(data).InsertAndGetId()
	if err != nil {
		return nil, err
	}

	return &model.BlogCreateOutput{
		BlogId: uint(id),
	}, nil
}

// 更新博客
func (s *sBlog) Update(ctx context.Context, in model.BlogUpdateInput) error {
	// 1. 获取当前用户ID
	userId := gconv.Uint(ctx.Value("userId"))
	if userId == 0 {
		return gerror.New("用户未登录")
	}

	// 2. 查询博客是否存在且是否为当前用户所有
	var blog entity.Blogs
	err := dao.Blogs.Ctx(ctx).
		Where(dao.Blogs.Columns().BlogId, in.BlogId).
		Scan(&blog)
	if err != nil {
		return err
	}

	// 3. 博客不存在
	if blog.BlogId == 0 {
		return gerror.New("博客不存在")
	}

	// 4. 检查是否有权限更新（仅作者本人或管理员可以更新）
	isAdmin := ctx.Value("isAdmin") != nil && gconv.Bool(ctx.Value("isAdmin"))
	if blog.AuthorId != userId && !isAdmin {
		return gerror.New("无权限操作此博客")
	}

	// 5. 如果未提供摘要，自动截取内容前150个字符作为摘要
	if in.Summary == "" && in.Content != "" {
		// 去除Markdown标记和空白字符
		content := gstr.Replace(in.Content, "\n", " ")
		content = gstr.Replace(content, "\r", " ")
		content = gstr.Replace(content, "#", "")
		content = gstr.Replace(content, "*", "")
		content = gstr.Replace(content, "```", "")
		content = strings.TrimSpace(content)

		// 截取前150个字符作为摘要
		if len(content) > 150 {
			in.Summary = content[:150] + "..."
		} else {
			in.Summary = content
		}
	}

	// 6. 更新数据
	data := g.Map{
		dao.Blogs.Columns().Title:      in.Title,
		dao.Blogs.Columns().Content:    in.Content,
		dao.Blogs.Columns().Summary:    in.Summary,
		dao.Blogs.Columns().Category:   in.Category,
		dao.Blogs.Columns().Tags:       in.Tags,
		dao.Blogs.Columns().CoverImage: in.CoverImage,
		dao.Blogs.Columns().Status:     in.Status,
	}

	_, err = dao.Blogs.Ctx(ctx).
		Data(data).
		Where(dao.Blogs.Columns().BlogId, in.BlogId).
		Update()

	return err
}

// 删除博客
func (s *sBlog) Delete(ctx context.Context, in model.BlogDeleteInput) error {
	// 1. 获取当前用户ID
	userId := gconv.Uint(ctx.Value("userId"))
	if userId == 0 {
		return gerror.New("用户未登录")
	}

	// 2. 查询博客是否存在且是否为当前用户所有
	var blog entity.Blogs
	err := dao.Blogs.Ctx(ctx).
		Where(dao.Blogs.Columns().BlogId, in.BlogId).
		Scan(&blog)
	if err != nil {
		return err
	}

	// 3. 博客不存在
	if blog.BlogId == 0 {
		return gerror.New("博客不存在")
	}

	// 4. 检查是否有权限删除（仅作者本人或管理员可以删除）
	isAdmin := ctx.Value("isAdmin") != nil && gconv.Bool(ctx.Value("isAdmin"))
	if blog.AuthorId != userId && !isAdmin {
		return gerror.New("无权限操作此博客")
	}

	// 5. 删除数据
	_, err = dao.Blogs.Ctx(ctx).
		Where(dao.Blogs.Columns().BlogId, in.BlogId).
		Delete()

	return err
}

// 设置博客推荐状态
func (s *sBlog) SetRecommend(ctx context.Context, in model.BlogRecommendInput) error {
	// 检查是否为管理员（只有管理员可以设置推荐）
	isAdmin := ctx.Value("isAdmin") != nil && gconv.Bool(ctx.Value("isAdmin"))
	if !isAdmin {
		return gerror.New("无权限执行此操作")
	}

	// 更新数据
	_, err := dao.Blogs.Ctx(ctx).
		Data(g.Map{
			dao.Blogs.Columns().IsRecommend: in.IsRecommend,
		}).
		Where(dao.Blogs.Columns().BlogId, in.BlogId).
		Update()

	return err
}

// 点赞博客
func (s *sBlog) Like(ctx context.Context, in model.BlogLikeInput) error {
	// 1. 获取当前用户ID
	userId := gconv.Uint(ctx.Value("userId"))
	if userId == 0 {
		return gerror.New("用户未登录")
	}

	// 2. 检查博客是否存在
	count, err := dao.Blogs.Ctx(ctx).
		Where(dao.Blogs.Columns().BlogId, in.BlogId).
		Count()
	if err != nil {
		return err
	}
	if count == 0 {
		return gerror.New("博客不存在")
	}

	// 3. 检查是否已点赞
	count, err = dao.BlogLikes.Ctx(ctx).
		Where(dao.BlogLikes.Columns().BlogId, in.BlogId).
		Where(dao.BlogLikes.Columns().UserId, userId).
		Count()
	if err != nil {
		return err
	}
	if count > 0 {
		return gerror.New("已经点过赞了")
	}

	// 4. 添加点赞记录
	_, err = dao.BlogLikes.Ctx(ctx).
		Data(entity.BlogLikes{
			BlogId: in.BlogId,
			UserId: userId,
		}).
		Insert()

	return err
}

// 取消点赞博客
func (s *sBlog) Unlike(ctx context.Context, in model.BlogUnlikeInput) error {
	// 1. 获取当前用户ID
	userId := gconv.Uint(ctx.Value("userId"))
	if userId == 0 {
		return gerror.New("用户未登录")
	}

	// 2. 删除点赞记录
	_, err := dao.BlogLikes.Ctx(ctx).
		Where(dao.BlogLikes.Columns().BlogId, in.BlogId).
		Where(dao.BlogLikes.Columns().UserId, userId).
		Delete()

	return err
}

// 获取博客评论列表
func (s *sBlog) GetCommentList(ctx context.Context, in model.BlogCommentListInput) (out *model.BlogCommentListOutput, err error) {
	m := dao.BlogComments.Ctx(ctx)

	// 1. 根据博客ID筛选
	m = m.Where(dao.BlogComments.Columns().BlogId, in.BlogId)

	// 2. 只查询状态正常的评论
	m = m.Where(dao.BlogComments.Columns().Status, 1)

	// 3. 默认只查询顶级评论
	m = m.Where(dao.BlogComments.Columns().ParentId + " IS NULL")

	// 4. 统计总数
	count, err := m.Count()
	if err != nil {
		return nil, err
	}

	out = &model.BlogCommentListOutput{
		Page:  in.Page,
		Size:  in.Size,
		Total: count,
		List:  make([]*model.BlogCommentItem, 0),
	}

	// 5. 如果没有数据，直接返回
	if count == 0 {
		return out, nil
	}

	// 6. 查询顶级评论
	var list []*entity.BlogComments
	err = m.OrderDesc(dao.BlogComments.Columns().CreateTime).
		Page(in.Page, in.Size).
		Scan(&list)
	if err != nil {
		return nil, err
	}

	// 7. 提取评论用户ID列表
	var userIds []uint
	var commentIds []uint
	for _, v := range list {
		userIds = append(userIds, v.UserId)
		commentIds = append(commentIds, v.CommentId)
	}

	// 8. 查询用户信息
	var userMap map[uint]entity.Users
	if len(userIds) > 0 {
		var users []entity.Users
		err = dao.Users.Ctx(ctx).
			Fields(dao.Users.Columns().UserId, dao.Users.Columns().UserName, dao.Users.Columns().Avatar).
			Where(dao.Users.Columns().UserId+" IN(?)", userIds).
			Scan(&users)
		if err != nil {
			return nil, err
		}

		userMap = make(map[uint]entity.Users, len(users))
		for _, v := range users {
			userMap[v.UserId] = v
		}
	}

	// 9. 查询每个顶级评论的回复数
	replyCountMap := make(map[uint]int)
	if len(commentIds) > 0 {
		var replyCounts []struct {
			ParentId uint `json:"parent_id"`
			Count    int  `json:"count"`
		}
		err = dao.BlogComments.Ctx(ctx).
			Fields("parent_id, COUNT(1) as count").
			Where(dao.BlogComments.Columns().ParentId+" IN(?)", commentIds).
			Where(dao.BlogComments.Columns().Status, 1).
			Group(dao.BlogComments.Columns().ParentId).
			Scan(&replyCounts)
		if err != nil {
			return nil, err
		}

		for _, v := range replyCounts {
			replyCountMap[v.ParentId] = v.Count
		}
	}

	// 10. 转换数据
	for _, v := range list {
		item := &model.BlogCommentItem{
			CommentId:  v.CommentId,
			BlogId:     v.BlogId,
			UserId:     v.UserId,
			Content:    v.Content,
			ParentId:   v.ParentId,
			LikeCount:  v.LikeCount,
			ReplyCount: replyCountMap[v.CommentId],
			CreateTime: v.CreateTime,
		}

		// 设置用户信息
		if user, ok := userMap[v.UserId]; ok {
			item.UserName = user.UserName
			item.UserAvatar = user.Avatar
		}

		out.List = append(out.List, item)
	}

	return out, nil
}

// 创建博客评论
func (s *sBlog) CreateComment(ctx context.Context, in model.BlogCommentCreateInput) (out *model.BlogCommentCreateOutput, err error) {
	// 1. 获取当前用户ID
	userId := gconv.Uint(ctx.Value("userId"))
	if userId == 0 {
		return nil, gerror.New("用户未登录")
	}

	// 2. 检查博客是否存在
	count, err := dao.Blogs.Ctx(ctx).
		Where(dao.Blogs.Columns().BlogId, in.BlogId).
		Count()
	if err != nil {
		return nil, err
	}
	if count == 0 {
		return nil, gerror.New("博客不存在")
	}

	// 3. 如果有父评论，检查父评论是否存在
	if in.ParentId > 0 {
		count, err = dao.BlogComments.Ctx(ctx).
			Where(dao.BlogComments.Columns().CommentId, in.ParentId).
			Where(dao.BlogComments.Columns().BlogId, in.BlogId).
			Count()
		if err != nil {
			return nil, err
		}
		if count == 0 {
			return nil, gerror.New("父评论不存在")
		}
	}

	// 4. 添加评论记录
	data := entity.BlogComments{
		BlogId:   in.BlogId,
		UserId:   userId,
		Content:  in.Content,
		ParentId: in.ParentId,
		Status:   1,
	}

	id, err := dao.BlogComments.Ctx(ctx).Data(data).InsertAndGetId()
	if err != nil {
		return nil, err
	}

	return &model.BlogCommentCreateOutput{
		CommentId: uint(id),
	}, nil
}

// 删除博客评论
func (s *sBlog) DeleteComment(ctx context.Context, in model.BlogCommentDeleteInput) error {
	// 1. 获取当前用户ID
	userId := gconv.Uint(ctx.Value("userId"))
	if userId == 0 {
		return gerror.New("用户未登录")
	}

	// 2. 查询评论是否存在且是否为当前用户所有
	var comment entity.BlogComments
	err := dao.BlogComments.Ctx(ctx).
		Where(dao.BlogComments.Columns().CommentId, in.CommentId).
		Scan(&comment)
	if err != nil {
		return err
	}

	// 3. 评论不存在
	if comment.CommentId == 0 {
		return gerror.New("评论不存在")
	}

	// 4. 检查是否有权限删除（仅评论作者本人或管理员可以删除）
	isAdmin := ctx.Value("isAdmin") != nil && gconv.Bool(ctx.Value("isAdmin"))
	if comment.UserId != userId && !isAdmin {
		return gerror.New("无权限操作此评论")
	}

	// 5. 软删除评论（将状态设为0）
	_, err = dao.BlogComments.Ctx(ctx).
		Data(g.Map{
			dao.BlogComments.Columns().Status: 0,
		}).
		Where(dao.BlogComments.Columns().CommentId, in.CommentId).
		Update()

	return err
}

// 获取博客分类列表
func (s *sBlog) GetCategoryList(ctx context.Context) (out *model.BlogCategoryListOutput, err error) {
	var categories []string
	err = dao.Blogs.Ctx(ctx).
		Fields("DISTINCT "+dao.Blogs.Columns().Category).
		OrderAsc(dao.Blogs.Columns().Category).
		Where(dao.Blogs.Columns().Status, 1).
		Scan(&categories)
	if err != nil {
		return nil, err
	}

	return &model.BlogCategoryListOutput{
		List: categories,
	}, nil
}

// 获取博客标签列表
func (s *sBlog) GetTagList(ctx context.Context) (out *model.BlogTagListOutput, err error) {
	// 初始化输出对象
	out = &model.BlogTagListOutput{
		List: make([]string, 0),
	}

	// 1. 查询所有标签字段
	var tagStrings []string
	err = dao.Blogs.Ctx(ctx).
		Fields(dao.Blogs.Columns().Tags).
		Where(dao.Blogs.Columns().Status, 1).
		Where(dao.Blogs.Columns().Tags + " != ''").
		Scan(&tagStrings)
	if err != nil {
		return out, err
	}

	// 2. 解析并合并标签
	tagMap := make(map[string]bool)
	for _, tagStr := range tagStrings {
		if tagStr == "" {
			continue
		}
		tags := strings.Split(tagStr, ",")
		for _, tag := range tags {
			tag = strings.TrimSpace(tag)
			if tag != "" {
				tagMap[tag] = true
			}
		}
	}

	// 3. 转换为列表
	for tag := range tagMap {
		out.List = append(out.List, tag)
	}

	// 4. 按字母顺序排序
	sort.Strings(out.List)

	return out, nil
}

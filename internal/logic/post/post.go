package post

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/patient-fyd/jxust-softhub-api/internal/dao"
	"github.com/patient-fyd/jxust-softhub-api/internal/model"
	"github.com/patient-fyd/jxust-softhub-api/internal/model/do"
	"github.com/patient-fyd/jxust-softhub-api/internal/model/entity"
	"github.com/patient-fyd/jxust-softhub-api/internal/service"
	"github.com/patient-fyd/jxust-softhub-api/utility/auth"
)

type sPost struct{}

func init() {
	service.RegisterPost(New())
}

func New() *sPost {
	return &sPost{}
}

// List 获取帖子列表
func (s *sPost) List(ctx context.Context, in model.PostListInput) (*model.PostListOutput, error) {
	var (
		m         = dao.Posts.Ctx(ctx)
		condition = g.Map{}
		out       = &model.PostListOutput{
			Page: in.Page,
			Size: in.Size,
		}
	)

	// 默认分页参数
	if in.Page <= 0 {
		in.Page = 1
	}
	if in.Size <= 0 {
		in.Size = 10
	}

	// 添加查询条件
	if in.Keyword != "" {
		condition["content LIKE ?"] = "%" + in.Keyword + "%"
	}
	if in.CircleId > 0 {
		condition["circleId"] = in.CircleId
	}
	if in.TopicId > 0 {
		condition["topicId"] = in.TopicId
	}
	if in.UserId > 0 {
		condition["userId"] = in.UserId
	}

	// 只查询状态为已发布的帖子
	condition["status"] = 1

	// 查询符合条件的记录总数
	count, err := m.Where(condition).Count()
	if err != nil {
		return nil, err
	}
	out.Total = count

	// 没有数据，直接返回
	if count == 0 {
		out.List = make([]model.PostListItem, 0)
		return out, nil
	}

	// 排序规则
	orderBy := "createTime DESC"
	if in.OrderBy == "hot" {
		orderBy = "likeCount DESC, commentCount DESC, createTime DESC"
	}

	// 查询列表数据
	var list []*entity.Posts
	err = m.Where(condition).
		Page(in.Page, in.Size).
		Order(orderBy).
		Scan(&list)
	if err != nil {
		return nil, err
	}

	// 当前登录用户ID
	loginUserId := auth.GetLoginUserId(ctx)

	// 转换结果格式
	out.List = make([]model.PostListItem, len(list))
	for i, v := range list {
		// 获取用户信息
		user, err := dao.Users.Ctx(ctx).Where("userId", v.UserId).One()
		if err != nil {
			return nil, err
		}

		// 获取圈子信息
		var circleName string
		if v.CircleId > 0 {
			circle, err := dao.Circles.Ctx(ctx).Where("circleId", v.CircleId).One()
			if err != nil {
				return nil, err
			}
			if !circle.IsEmpty() {
				circleName = circle["name"].String()
			}
		}

		// 获取话题信息
		var topicName string
		if v.TopicId > 0 {
			topic, err := dao.Topics.Ctx(ctx).Where("topicId", v.TopicId).One()
			if err != nil {
				return nil, err
			}
			if !topic.IsEmpty() {
				topicName = topic["name"].String()
			}
		}

		// 获取帖子图片
		var images []string
		err = dao.PostImages.Ctx(ctx).
			Where("postId", v.PostId).
			Fields("imageUrl").
			Order("sortOrder ASC").
			Scan(&images, "imageUrl")
		if err != nil {
			return nil, err
		}

		// 判断当前用户是否已点赞该帖子
		isLiked := false
		if loginUserId > 0 {
			likeCount, err := dao.Likes.Ctx(ctx).
				Where(g.Map{
					"userId":     loginUserId,
					"targetId":   v.PostId,
					"targetType": 1, // 1为帖子
				}).
				Count()
			if err != nil {
				return nil, err
			}
			isLiked = likeCount > 0
		}

		out.List[i] = model.PostListItem{
			PostId:       int(v.PostId),
			UserId:       int(v.UserId),
			UserName:     user["userName"].String(),
			UserAvatar:   user["avatar"].String(),
			Content:      v.Content,
			Images:       images,
			CircleId:     int(v.CircleId),
			CircleName:   circleName,
			TopicId:      int(v.TopicId),
			TopicName:    topicName,
			ViewCount:    int(v.ViewCount),
			LikeCount:    int(v.LikeCount),
			CommentCount: int(v.CommentCount),
			ShareCount:   int(v.ShareCount),
			IsTop:        v.IsTop,
			IsLiked:      isLiked,
			CreateTime:   v.CreateTime,
		}
	}

	return out, nil
}

// Detail 获取帖子详情
func (s *sPost) Detail(ctx context.Context, in model.PostDetailInput) (*model.PostDetailOutput, error) {
	var (
		m = dao.Posts.Ctx(ctx)
	)

	// 查询帖子详情
	post, err := m.Where(do.Posts{
		PostId: in.PostId,
		Status: 1, // 只查询已发布的帖子
	}).One()
	if err != nil {
		return nil, err
	}
	if post.IsEmpty() {
		return nil, gerror.New("帖子不存在或已删除")
	}

	var result entity.Posts
	err = post.Struct(&result)
	if err != nil {
		return nil, err
	}

	// 增加浏览次数
	_, err = m.Where(do.Posts{
		PostId: in.PostId,
	}).Increment("viewCount", 1)
	if err != nil {
		return nil, err
	}

	// 当前登录用户ID
	loginUserId := auth.GetLoginUserId(ctx)

	// 获取用户信息
	user, err := dao.Users.Ctx(ctx).Where("userId", result.UserId).One()
	if err != nil {
		return nil, err
	}

	// 获取圈子信息
	var circleName string
	if result.CircleId > 0 {
		circle, err := dao.Circles.Ctx(ctx).Where("circleId", result.CircleId).One()
		if err != nil {
			return nil, err
		}
		if !circle.IsEmpty() {
			circleName = circle["name"].String()
		}
	}

	// 获取话题信息
	var topicName string
	if result.TopicId > 0 {
		topic, err := dao.Topics.Ctx(ctx).Where("topicId", result.TopicId).One()
		if err != nil {
			return nil, err
		}
		if !topic.IsEmpty() {
			topicName = topic["name"].String()
		}
	}

	// 获取帖子图片
	var images []string
	err = dao.PostImages.Ctx(ctx).
		Where("postId", result.PostId).
		Fields("imageUrl").
		Order("sortOrder ASC").
		Scan(&images, "imageUrl")
	if err != nil {
		return nil, err
	}

	// 判断当前用户是否已点赞该帖子
	isLiked := false
	if loginUserId > 0 {
		likeCount, err := dao.Likes.Ctx(ctx).
			Where(g.Map{
				"userId":     loginUserId,
				"targetId":   result.PostId,
				"targetType": 1, // 1为帖子
			}).
			Count()
		if err != nil {
			return nil, err
		}
		isLiked = likeCount > 0
	}

	return &model.PostDetailOutput{
		PostId:       int(result.PostId),
		UserId:       int(result.UserId),
		UserName:     user["userName"].String(),
		UserAvatar:   user["avatar"].String(),
		Content:      result.Content,
		Images:       images,
		CircleId:     int(result.CircleId),
		CircleName:   circleName,
		TopicId:      int(result.TopicId),
		TopicName:    topicName,
		ViewCount:    int(result.ViewCount + 1), // +1 是因为刚增加的查看次数
		LikeCount:    int(result.LikeCount),
		CommentCount: int(result.CommentCount),
		ShareCount:   int(result.ShareCount),
		IsTop:        result.IsTop,
		IsLiked:      isLiked,
		CreateTime:   result.CreateTime,
		UpdateTime:   result.UpdateTime,
	}, nil
}

// Create 创建帖子
func (s *sPost) Create(ctx context.Context, in model.PostCreateInput) (*model.PostCreateOutput, error) {
	// 获取当前登录用户ID
	userId := auth.GetLoginUserId(ctx)
	if userId <= 0 {
		return nil, gerror.New("用户未登录")
	}

	var (
		postId int64
		err    error
	)

	// 开启事务
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 创建帖子记录
		result, err := dao.Posts.Ctx(ctx).TX(tx).Insert(do.Posts{
			UserId:       int(userId),
			Content:      in.Content,
			CircleId:     in.CircleId,
			TopicId:      in.TopicId,
			ViewCount:    0,
			LikeCount:    0,
			CommentCount: 0,
			ShareCount:   0,
			IsTop:        0,
			Status:       1, // 状态：1-已发布
			CreateTime:   gtime.Now(),
			UpdateTime:   gtime.Now(),
		})
		if err != nil {
			return err
		}

		// 获取创建的帖子ID
		postId, err = result.LastInsertId()
		if err != nil {
			return err
		}

		// 处理图片
		if len(in.Images) > 0 {
			for i, imageUrl := range in.Images {
				_, err = dao.PostImages.Ctx(ctx).TX(tx).Insert(do.PostImages{
					PostId:     int(postId),
					ImageUrl:   imageUrl,
					SortOrder:  i,
					CreateTime: gtime.Now(),
				})
				if err != nil {
					return err
				}
			}
		}

		// 更新圈子帖子数量
		if in.CircleId > 0 {
			_, err = dao.Circles.Ctx(ctx).TX(tx).
				Where("circleId", in.CircleId).
				Increment("postCount", 1)
			if err != nil {
				return err
			}
		}

		// 更新话题帖子数量
		if in.TopicId > 0 {
			_, err = dao.Topics.Ctx(ctx).TX(tx).
				Where("topicId", in.TopicId).
				Increment("postCount", 1)
			if err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &model.PostCreateOutput{
		PostId: int(postId),
	}, nil
}

// Delete 删除帖子
func (s *sPost) Delete(ctx context.Context, in model.PostDeleteInput) (*model.PostDeleteOutput, error) {
	// 获取当前登录用户ID
	userId := auth.GetLoginUserId(ctx)
	if userId <= 0 {
		return nil, gerror.New("用户未登录")
	}

	// 查询帖子信息
	post, err := dao.Posts.Ctx(ctx).Where("postId", in.PostId).One()
	if err != nil {
		return nil, err
	}
	if post.IsEmpty() {
		return nil, gerror.New("帖子不存在")
	}

	// 验证当前用户是否是帖子作者
	if post["userId"].Int() != int(userId) {
		// 检查当前用户是否有管理员权限
		userRole, err := dao.Users.Ctx(ctx).Where("userId", userId).Value("roleId")
		if err != nil {
			return nil, err
		}
		if userRole.Int() > 3 { // 假设1-3是管理员角色
			return nil, gerror.New("无权删除此帖子")
		}
	}

	var (
		success bool
	)

	// 开启事务
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 记录要更新的话题和圈子ID
		circleId := post["circleId"].Int()
		topicId := post["topicId"].Int()

		// 软删除帖子（将状态改为已删除）
		_, err := dao.Posts.Ctx(ctx).TX(tx).
			Data(do.Posts{
				Status:     2, // 状态：2-已删除
				UpdateTime: gtime.Now(),
			}).
			Where("postId", in.PostId).
			Update()
		if err != nil {
			return err
		}

		// 更新圈子帖子数量
		if circleId > 0 {
			_, err = dao.Circles.Ctx(ctx).TX(tx).
				Where("circleId", circleId).
				Where("postCount > 0"). // 防止减为负数
				Decrement("postCount", 1)
			if err != nil {
				return err
			}
		}

		// 更新话题帖子数量
		if topicId > 0 {
			_, err = dao.Topics.Ctx(ctx).TX(tx).
				Where("topicId", topicId).
				Where("postCount > 0"). // 防止减为负数
				Decrement("postCount", 1)
			if err != nil {
				return err
			}
		}

		success = true
		return nil
	})

	if err != nil {
		return nil, err
	}

	return &model.PostDeleteOutput{
		Success: success,
	}, nil
}

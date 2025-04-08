package post

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/patient-fyd/jxust-softhub-api/internal/dao"
	"github.com/patient-fyd/jxust-softhub-api/internal/model"
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
	out.List = make([]model.PostListItem, 0, len(list))
	if len(list) > 0 {
		// 收集所有帖子ID和用户ID
		postIds := make([]interface{}, 0, len(list))
		userIds := make([]interface{}, 0, len(list))
		for _, v := range list {
			postIds = append(postIds, v.PostId)
			userIds = append(userIds, v.UserId)
		}

		// 批量查询帖子图片
		var postImages []struct {
			PostId   uint
			ImageUrl string
		}
		err = dao.PostImages.Ctx(ctx).
			Fields("postId", "imageUrl").
			WhereIn("postId", postIds).
			Order("sortOrder ASC").
			Scan(&postImages)
		if err != nil {
			return nil, err
		}

		// 整理图片数据
		postImagesMap := make(map[uint][]string)
		for _, img := range postImages {
			postImagesMap[img.PostId] = append(postImagesMap[img.PostId], img.ImageUrl)
		}

		// 批量查询用户信息
		var users []struct {
			UserId   uint
			UserName string
			Avatar   string
		}
		err = dao.Users.Ctx(ctx).
			Fields("userId", "userName", "avatar").
			WhereIn("userId", userIds).
			Scan(&users)
		if err != nil {
			return nil, err
		}

		// 整理用户数据
		userMap := make(map[uint]struct {
			UserName string
			Avatar   string
		})
		for _, u := range users {
			userMap[u.UserId] = struct {
				UserName string
				Avatar   string
			}{
				UserName: u.UserName,
				Avatar:   u.Avatar,
			}
		}

		// 查询圈子和话题信息
		circleIds := make([]interface{}, 0)
		topicIds := make([]interface{}, 0)
		for _, v := range list {
			if v.CircleId > 0 {
				circleIds = append(circleIds, v.CircleId)
			}
			if v.TopicId > 0 {
				topicIds = append(topicIds, v.TopicId)
			}
		}

		// 整理圈子数据
		circleMap := make(map[uint]string)
		if len(circleIds) > 0 {
			var circles []struct {
				CircleId   uint
				CircleName string
			}
			err = dao.Circles.Ctx(ctx).
				Fields("circleId", "circleName").
				WhereIn("circleId", circleIds).
				Scan(&circles)
			if err == nil {
				for _, c := range circles {
					circleMap[c.CircleId] = c.CircleName
				}
			}
		}

		// 整理话题数据
		topicMap := make(map[uint]string)
		if len(topicIds) > 0 {
			var topics []struct {
				TopicId   uint
				TopicName string
			}
			err = dao.Topics.Ctx(ctx).
				Fields("topicId", "topicName").
				WhereIn("topicId", topicIds).
				Scan(&topics)
			if err == nil {
				for _, t := range topics {
					topicMap[t.TopicId] = t.TopicName
				}
			}
		}

		// 查询用户点赞状态
		userLikedMap := make(map[uint]bool)
		if loginUserId > 0 {
			var userLikes []struct {
				PostId uint
			}
			err = dao.Likes.Ctx(ctx).
				Fields("targetId AS postId").
				Where("userId", loginUserId).
				Where("targetType", 1). // 1为帖子
				WhereIn("targetId", postIds).
				Scan(&userLikes)
			if err == nil {
				for _, like := range userLikes {
					userLikedMap[like.PostId] = true
				}
			}
		}

		// 组装最终数据
		for _, v := range list {
			item := model.PostListItem{
				PostId:       int(v.PostId),
				UserId:       int(v.UserId),
				UserName:     userMap[v.UserId].UserName,
				UserAvatar:   userMap[v.UserId].Avatar,
				Content:      v.Content,
				Images:       postImagesMap[v.PostId],
				CircleId:     int(v.CircleId),
				CircleName:   circleMap[v.CircleId],
				TopicId:      int(v.TopicId),
				TopicName:    topicMap[v.TopicId],
				ViewCount:    int(v.ViewCount),
				LikeCount:    int(v.LikeCount),
				CommentCount: int(v.CommentCount),
				ShareCount:   int(v.ShareCount),
				IsTop:        int(v.IsTop),
				IsLiked:      userLikedMap[v.PostId],
				CreateTime:   v.CreateTime,
			}
			out.List = append(out.List, item)
		}
	}

	return out, nil
}

// Detail 获取帖子详情
func (s *sPost) Detail(ctx context.Context, in model.PostDetailInput) (*model.PostDetailOutput, error) {
	// 查询帖子信息
	var result *entity.Posts
	err := dao.Posts.Ctx(ctx).
		Where("postId", in.PostId).
		Where("status", 1). // 只查询状态为已发布的帖子
		Scan(&result)
	if err != nil {
		return nil, err
	}
	if result == nil {
		return nil, gerror.New("帖子不存在或已删除")
	}

	// 递增浏览量
	_, err = dao.Posts.Ctx(ctx).
		Where("postId", in.PostId).
		Increment("viewCount", 1)
	if err != nil {
		g.Log().Warning(ctx, "增加帖子浏览量失败:", err)
	}

	// 获取帖子图片
	var postImages []struct {
		ImageUrl string
	}
	err = dao.PostImages.Ctx(ctx).
		Fields("imageUrl").
		Where("postId", in.PostId).
		Order("sortOrder ASC").
		Scan(&postImages)
	if err != nil {
		return nil, err
	}

	// 提取图片URL列表
	images := make([]string, 0, len(postImages))
	for _, img := range postImages {
		images = append(images, img.ImageUrl)
	}

	// 获取用户信息
	var userInfo struct {
		UserName string
		Avatar   string
	}
	err = dao.Users.Ctx(ctx).
		Fields("userName", "avatar").
		Where("userId", result.UserId).
		Scan(&userInfo)
	if err != nil {
		return nil, err
	}

	// 获取圈子名称
	circleName := ""
	if result.CircleId > 0 {
		var circle struct {
			CircleName string
		}
		err = dao.Circles.Ctx(ctx).
			Fields("circleName").
			Where("circleId", result.CircleId).
			Scan(&circle)
		if err == nil {
			circleName = circle.CircleName
		}
	}

	// 获取话题名称
	topicName := ""
	if result.TopicId > 0 {
		var topic struct {
			TopicName string
		}
		err = dao.Topics.Ctx(ctx).
			Fields("topicName").
			Where("topicId", result.TopicId).
			Scan(&topic)
		if err == nil {
			topicName = topic.TopicName
		}
	}

	// 查询当前用户是否已点赞
	isLiked := false
	loginUserId := auth.GetLoginUserId(ctx)
	if loginUserId > 0 {
		count, _ := dao.Likes.Ctx(ctx).
			Where("targetId", in.PostId).
			Where("userId", loginUserId).
			Where("targetType", 1). // 1为帖子
			Count()
		isLiked = count > 0
	}

	return &model.PostDetailOutput{
		PostId:       int(result.PostId),
		UserId:       int(result.UserId),
		UserName:     userInfo.UserName,
		UserAvatar:   userInfo.Avatar,
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
	// 检查输入参数
	if ctx == nil {
		g.Log().Error(ctx, "创建帖子失败: 上下文为空")
		return nil, gerror.New("系统错误: 上下文为空")
	}
	g.Log().Debugf(ctx, "开始创建帖子,输入参数: %+v", in)

	// 初始化Images字段，防止空指针
	if in.Images == nil {
		in.Images = make([]string, 0)
		g.Log().Debug(ctx, "初始化空Images数组")
	}

	// 获取当前登录用户ID
	userId := auth.GetLoginUserId(ctx)
	if userId == 0 {
		g.Log().Error(ctx, "创建帖子失败: 用户未登录")
		return nil, gerror.New("用户未登录")
	}
	g.Log().Debugf(ctx, "当前登录用户ID: %d", userId)

	// 检查数据库连接
	if dao.Posts.DB() == nil {
		g.Log().Error(ctx, "创建帖子失败: 数据库连接为空")
		return nil, gerror.New("系统错误: 数据库连接未初始化")
	}
	g.Log().Debug(ctx, "数据库连接正常")

	// 检查圈子是否存在
	if in.CircleId > 0 {
		count, err := dao.Circles.Ctx(ctx).Where("circleId", in.CircleId).Count()
		if err != nil {
			g.Log().Errorf(ctx, "检查圈子是否存在失败: %v", err)
			return nil, err
		}
		if count == 0 {
			g.Log().Errorf(ctx, "圈子不存在: %d", in.CircleId)
			return nil, gerror.New("圈子不存在")
		}
	}

	// 检查话题是否存在
	if in.TopicId > 0 {
		count, err := dao.Topics.Ctx(ctx).Where("topicId", in.TopicId).Count()
		if err != nil {
			g.Log().Errorf(ctx, "检查话题是否存在失败: %v", err)
			return nil, err
		}
		if count == 0 {
			g.Log().Errorf(ctx, "话题不存在: %d", in.TopicId)
			return nil, gerror.New("话题不存在")
		}
	}

	// 准备帖子数据
	data := g.Map{
		"userId":     userId,
		"content":    in.Content,
		"circleId":   in.CircleId,
		"topicId":    in.TopicId,
		"status":     1, // 1-已发布
		"createTime": gtime.Now(),
		"updateTime": gtime.Now(),
	}
	g.Log().Debugf(ctx, "准备插入的帖子数据: %+v", data)

	// 简化版本 - 直接插入而不使用事务
	// 这样可以排除事务处理带来的复杂性
	g.Log().Debug(ctx, "开始插入帖子数据(不使用事务)")
	lastInsertId, err := dao.Posts.Ctx(ctx).Data(data).InsertAndGetId()
	if err != nil {
		g.Log().Errorf(ctx, "插入帖子失败: %v", err)
		return nil, gerror.Newf("插入帖子失败: %v", err)
	}
	g.Log().Debugf(ctx, "帖子创建成功,ID: %d", lastInsertId)

	// 如果有图片,插入图片记录
	if len(in.Images) > 0 {
		g.Log().Debugf(ctx, "开始插入帖子图片,数量: %d", len(in.Images))
		imageData := make([]g.Map, 0, len(in.Images))
		for i, url := range in.Images {
			imageData = append(imageData, g.Map{
				"postId":     lastInsertId,
				"imageUrl":   url,
				"sortOrder":  i,
				"createTime": gtime.Now(),
			})
		}
		_, err = dao.PostImages.Ctx(ctx).Data(imageData).Insert()
		if err != nil {
			g.Log().Errorf(ctx, "插入帖子图片失败: %v", err)
			// 图片插入失败不影响帖子创建结果
			g.Log().Warning(ctx, "图片插入失败,但帖子已成功创建")
		} else {
			g.Log().Debugf(ctx, "插入帖子图片成功,数量: %d", len(in.Images))
		}
	}

	// 检查结果
	if lastInsertId <= 0 {
		g.Log().Error(ctx, "帖子创建失败,返回的ID无效")
		return nil, gerror.New("帖子创建失败,返回的ID无效")
	}

	g.Log().Debugf(ctx, "帖子创建完成,返回ID: %d", lastInsertId)
	return &model.PostCreateOutput{
		PostId: int(lastInsertId),
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
	var post *entity.Posts
	err := dao.Posts.Ctx(ctx).
		Where("postId", in.PostId).
		Where("status", 1). // 只处理状态为已发布的帖子
		Scan(&post)
	if err != nil {
		return nil, err
	}
	if post == nil {
		return nil, gerror.New("帖子不存在或已删除")
	}

	// 检查当前用户是否有权限删除该帖子
	// 只有管理员或帖子作者可以删除
	var roleInfo struct {
		RoleId uint
	}
	err = dao.Users.Ctx(ctx).
		Fields("roleId").
		Where("userId", userId).
		Scan(&roleInfo)
	if err != nil {
		return nil, err
	}
	roleId := roleInfo.RoleId

	// 检查权限：管理员(roleId=1,2,3)或帖子作者
	if roleId != 1 && roleId != 2 && roleId != 3 && userId != uint(post.UserId) {
		return nil, gerror.New("您没有权限删除该帖子")
	}

	// 开启事务
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 软删除帖子（将状态设为已删除）
		_, err := dao.Posts.Ctx(ctx).TX(tx).
			Data(g.Map{"status": 2, "updateTime": gtime.Now()}).
			Where("postId", in.PostId).
			Update()
		if err != nil {
			return err
		}

		// 更新圈子帖子数量
		if post.CircleId > 0 {
			_, err = dao.Circles.Ctx(ctx).TX(tx).
				Where("circleId", post.CircleId).
				Where("postCount > 0"). // 防止减到负数
				Decrement("postCount", 1)
			if err != nil {
				return err
			}
		}

		// 更新话题帖子数量
		if post.TopicId > 0 {
			_, err = dao.Topics.Ctx(ctx).TX(tx).
				Where("topicId", post.TopicId).
				Where("postCount > 0"). // 防止减到负数
				Decrement("postCount", 1)
			if err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &model.PostDeleteOutput{
		Success: true,
	}, nil
}

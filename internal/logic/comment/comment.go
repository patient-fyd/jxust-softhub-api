package comment

import (
	"context"
	"strings"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"

	"github.com/patient-fyd/jxust-softhub-api/internal/dao"
	"github.com/patient-fyd/jxust-softhub-api/internal/model"
	"github.com/patient-fyd/jxust-softhub-api/internal/service"
	"github.com/patient-fyd/jxust-softhub-api/utility/auth"
)

type sComment struct{}

func init() {
	service.RegisterComment(New())
}

func New() *sComment {
	return &sComment{}
}

// List 获取评论列表
func (s *sComment) List(ctx context.Context, in model.CommentListInput) (*model.CommentListOutput, error) {
	// 参数处理
	if in.Page <= 0 {
		in.Page = 1
	}
	if in.Size <= 0 {
		in.Size = 10
	}

	// 获取当前登录用户ID
	userId := auth.GetLoginUserId(ctx)

	// 查询评论总数
	total, err := dao.Comments.Ctx(ctx).
		Where("contentType", in.ContentType).
		Where("contentId", in.ContentId).
		Count()
	if err != nil {
		return nil, err
	}

	// 查询评论列表
	var (
		list = make([]model.CommentListItem, 0)
		m    = dao.Comments.Ctx(ctx).
			Where("contentType", in.ContentType).
			Where("contentId", in.ContentId).
			Order("createTime DESC").
			Page(in.Page, in.Size)
	)
	result, err := m.All()
	if err != nil {
		return nil, err
	}

	// 获取所有评论的用户ID
	userIds := make([]int, 0)
	for _, item := range result {
		userIds = append(userIds, item["userId"].Int())
	}

	// 获取用户信息
	userMap := make(map[int]gdb.Record)
	if len(userIds) > 0 {
		users, err := dao.Users.Ctx(ctx).
			Where("userId IN(?)", userIds).
			All()
		if err != nil {
			return nil, err
		}
		for _, user := range users {
			userMap[user["userId"].Int()] = user
		}
	}

	// 转换评论列表
	for _, item := range result {
		// 获取用户信息
		var (
			username   = ""
			userAvatar = ""
		)
		if user, ok := userMap[item["userId"].Int()]; ok {
			username = user["userName"].String()
			userAvatar = user["avatar"].String()
		}

		// 获取点赞数
		likeCount, err := dao.Likes.Ctx(ctx).
			Where("targetId", item["commentId"]).
			Where("targetType", 2). // 2为评论
			Count()
		if err != nil {
			return nil, err
		}

		// 检查当前用户是否已点赞
		isLiked := false
		if userId > 0 {
			count, err := dao.Likes.Ctx(ctx).
				Where("targetId", item["commentId"]).
				Where("targetType", 2). // 2为评论
				Where("userId", userId).
				Count()
			if err != nil {
				return nil, err
			}
			isLiked = count > 0
		}

		list = append(list, model.CommentListItem{
			CommentId:   item["commentId"].Int(),
			ContentType: item["contentType"].String(),
			ContentId:   item["contentId"].Int(),
			UserId:      item["userId"].Int(),
			UserName:    username,
			UserAvatar:  userAvatar,
			Content:     item["content"].String(),
			LikeCount:   int(likeCount),
			IsLiked:     isLiked,
			CreateTime:  item["createTime"].GTime(),
		})
	}

	// 返回结果
	out := &model.CommentListOutput{
		List:  list,
		Total: int(total),
		Page:  in.Page,
		Size:  in.Size,
	}

	return out, nil
}

// Create 创建评论
func (s *sComment) Create(ctx context.Context, in model.CommentCreateInput) (*model.CommentCreateOutput, error) {
	// 获取当前登录用户ID
	userId := auth.GetLoginUserId(ctx)
	if userId <= 0 {
		return nil, gerror.New("用户未登录")
	}

	// 验证评论内容
	if len(in.Content) == 0 {
		return nil, gerror.New("评论内容不能为空")
	}
	if len(in.Content) > 500 {
		return nil, gerror.New("评论内容不能超过500个字符")
	}

	// 过滤敏感词
	filteredContent, hasSensitiveWord := s.filterSensitiveWords(ctx, in.Content)
	if hasSensitiveWord {
		return nil, gerror.New("评论内容包含敏感词，请修改后重试")
	}

	var (
		out       = &model.CommentCreateOutput{}
		commentId int64
	)

	// 开启事务
	err := g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 插入评论
		result, err := dao.Comments.Ctx(ctx).TX(tx).Insert(g.Map{
			"contentType": in.ContentType,
			"contentId":   in.ContentId,
			"userId":      userId,
			"content":     filteredContent,
			"createTime":  gtime.Now(),
		})
		if err != nil {
			return err
		}

		// 获取评论ID
		commentId, err = result.LastInsertId()
		if err != nil {
			return err
		}

		// 如果是帖子评论，增加帖子的评论数
		if in.ContentType == "post" {
			_, err = dao.Posts.Ctx(ctx).TX(tx).
				Where("postId", in.ContentId).
				Increment("commentCount", 1)
			if err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	out.CommentId = int(commentId)
	return out, nil
}

// filterSensitiveWords 过滤敏感词
func (s *sComment) filterSensitiveWords(ctx context.Context, content string) (string, bool) {
	// TODO: 实现敏感词过滤功能
	// 方案1: 使用本地敏感词库，采用DFA(确定有限自动机)算法实现高效过滤
	// 方案2: 使用Redis存储敏感词库，使用Lua脚本实现过滤
	// 方案3: 调用第三方内容审核API

	// 示例实现：简单替换一些敏感词（实际应用中应使用更完善的敏感词库和算法）
	sensitiveWords := []string{"傻逼", "操", "草", "fuck", "shit", "dick"}
	filteredContent := content
	hasSensitiveWord := false

	for _, word := range sensitiveWords {
		if strings.Contains(filteredContent, word) {
			hasSensitiveWord = true
			filteredContent = strings.ReplaceAll(filteredContent, word, "***")
		}
	}

	return filteredContent, hasSensitiveWord
}

// Delete 删除评论
func (s *sComment) Delete(ctx context.Context, in model.CommentDeleteInput) (*model.CommentDeleteOutput, error) {
	// 获取当前登录用户ID
	userId := auth.GetLoginUserId(ctx)
	if userId <= 0 {
		return nil, gerror.New("用户未登录")
	}

	// 查询评论是否存在
	comment, err := dao.Comments.Ctx(ctx).Where("commentId", in.CommentId).One()
	if err != nil {
		return nil, err
	}
	if comment.IsEmpty() {
		return nil, gerror.New("评论不存在")
	}

	// 检查是否有权限删除（评论作者或管理员）
	if comment["userId"].Int() != int(userId) {
		// 检查是否为管理员
		isAdmin, err := service.Auth().IsAdmin(ctx)
		if err != nil {
			return nil, err
		}
		if !isAdmin {
			return nil, gerror.New("没有权限删除该评论")
		}
	}

	var success bool
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 删除评论
		_, err := dao.Comments.Ctx(ctx).TX(tx).
			Where("commentId", in.CommentId).
			Delete()
		if err != nil {
			return err
		}

		// 删除评论相关的点赞记录
		_, err = dao.Likes.Ctx(ctx).TX(tx).
			Where(g.Map{
				"targetId":   in.CommentId,
				"targetType": 2, // 2为评论
			}).
			Delete()
		if err != nil {
			return err
		}

		// 如果是帖子评论，减少帖子的评论数
		contentType := comment["contentType"].String()
		contentId := comment["contentId"].Int()
		if contentType == "post" {
			_, err = dao.Posts.Ctx(ctx).TX(tx).
				Where("postId", contentId).
				Decrement("commentCount", 1)
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

	return &model.CommentDeleteOutput{
		Success: success,
	}, nil
}

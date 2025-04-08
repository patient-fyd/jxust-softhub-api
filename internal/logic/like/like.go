package like

import (
	"context"
	"fmt"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"

	"github.com/patient-fyd/jxust-softhub-api/internal/dao"
	"github.com/patient-fyd/jxust-softhub-api/internal/model"
	"github.com/patient-fyd/jxust-softhub-api/internal/service"
	"github.com/patient-fyd/jxust-softhub-api/utility/auth"
)

type sLike struct{}

func init() {
	service.RegisterLike(New())
}

func New() *sLike {
	return &sLike{}
}

// Toggle 点赞/取消点赞
func (s *sLike) Toggle(ctx context.Context, in model.LikeToggleReq) (*model.LikeToggleOutput, error) {
	// 获取当前登录用户ID
	userId := auth.GetLoginUserId(ctx)
	if userId <= 0 {
		return nil, gerror.New("用户未登录")
	}

	// 验证目标类型
	if in.TargetType != 1 && in.TargetType != 2 {
		return nil, gerror.New("无效的目标类型")
	}

	// 验证目标是否存在
	var exists bool
	var err error

	if in.TargetType == 1 { // 帖子
		count, err := dao.Posts.Ctx(ctx).
			Where("postId", in.TargetId).
			Where("status", 1). // 已发布状态
			Count()
		if err != nil {
			return nil, err
		}
		exists = count > 0
	} else if in.TargetType == 2 { // 评论
		count, err := dao.Comments.Ctx(ctx).
			Where("commentId", in.TargetId).
			Count()
		if err != nil {
			return nil, err
		}
		exists = count > 0
	}

	if !exists {
		return nil, gerror.New("目标不存在或已删除")
	}

	// 查询是否已点赞
	likeRecord, err := dao.Likes.Ctx(ctx).
		Where("userId", userId).
		Where("targetId", in.TargetId).
		Where("targetType", in.TargetType).
		One()
	if err != nil {
		return nil, err
	}

	var (
		isLiked  bool
		likeId   uint
		likeTime *gtime.Time
	)

	// 开启事务
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		if likeRecord.IsEmpty() {
			// 未点赞，添加点赞记录
			result, err := dao.Likes.Ctx(ctx).TX(tx).Insert(g.Map{
				"userId":     userId,
				"targetId":   in.TargetId,
				"targetType": in.TargetType,
				"createTime": gtime.Now(),
			})
			if err != nil {
				return err
			}

			lastInsertId, err := result.LastInsertId()
			if err != nil {
				return err
			}
			likeId = uint(lastInsertId)
			likeTime = gtime.Now()
			isLiked = true

			// 更新目标的点赞数
			if in.TargetType == 1 { // 帖子
				_, err = dao.Posts.Ctx(ctx).TX(tx).
					Where("postId", in.TargetId).
					Increment("likeCount", 1)
				if err != nil {
					return err
				}
			} else if in.TargetType == 2 { // 评论
				_, err = dao.Comments.Ctx(ctx).TX(tx).
					Where("commentId", in.TargetId).
					Increment("likeCount", 1)
				if err != nil {
					return err
				}
			}
		} else {
			// 已点赞，取消点赞
			_, err = dao.Likes.Ctx(ctx).TX(tx).
				Where("userId", userId).
				Where("targetId", in.TargetId).
				Where("targetType", in.TargetType).
				Delete()
			if err != nil {
				return err
			}

			isLiked = false

			// 更新目标的点赞数
			if in.TargetType == 1 { // 帖子
				_, err = dao.Posts.Ctx(ctx).TX(tx).
					Where("postId", in.TargetId).
					Where("likeCount > 0"). // 防止减为负数
					Decrement("likeCount", 1)
				if err != nil {
					return err
				}
			} else if in.TargetType == 2 { // 评论
				_, err = dao.Comments.Ctx(ctx).TX(tx).
					Where("commentId", in.TargetId).
					Where("likeCount > 0"). // 防止减为负数
					Decrement("likeCount", 1)
				if err != nil {
					return err
				}
			}
		}

		// 更新 Redis 热门排序
		redis := g.Redis()
		key := fmt.Sprintf("hot:%d", in.TargetType)
		score := float64(1)
		if !isLiked {
			score = -1
		}
		_, err = redis.ZIncrBy(ctx, key, score, gconv.String(in.TargetId))
		if err != nil {
			g.Log().Errorf(ctx, "更新 Redis 热门排序失败: %v", err)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &model.LikeToggleOutput{
		LikeId:   likeId,
		IsLiked:  isLiked,
		LikeTime: likeTime,
	}, nil
}

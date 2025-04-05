package follow

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"

	"github.com/patient-fyd/jxust-softhub-api/internal/dao"
	"github.com/patient-fyd/jxust-softhub-api/internal/model"
	"github.com/patient-fyd/jxust-softhub-api/internal/service"
	"github.com/patient-fyd/jxust-softhub-api/utility/auth"
)

type sFollow struct{}

func init() {
	service.RegisterFollow(New())
}

func New() *sFollow {
	return &sFollow{}
}

// Toggle 关注/取消关注
func (s *sFollow) Toggle(ctx context.Context, in model.FollowToggleReq) (*model.FollowToggleOutput, error) {
	// 获取当前登录用户ID
	userId := auth.GetLoginUserId(ctx)
	if userId <= 0 {
		return nil, gerror.New("用户未登录")
	}

	// 验证被关注的对象类型
	if in.FollowType != 1 && in.FollowType != 2 {
		return nil, gerror.New("无效的关注类型")
	}

	// 不能关注自己
	if in.FollowType == 1 && uint(in.FollowedId) == userId {
		return nil, gerror.New("不能关注自己")
	}

	// 验证被关注的对象是否存在
	var exists bool
	var err error

	if in.FollowType == 1 { // 用户
		count, err := dao.Users.Ctx(ctx).
			Where("userId", in.FollowedId).
			Count()
		if err != nil {
			return nil, err
		}
		exists = count > 0
	} else if in.FollowType == 2 { // 圈子
		count, err := dao.Circles.Ctx(ctx).
			Where("circleId", in.FollowedId).
			Count()
		if err != nil {
			return nil, err
		}
		exists = count > 0
	}

	if !exists {
		return nil, gerror.New("关注对象不存在")
	}

	// 查询是否已关注
	followRecord, err := dao.Follows.Ctx(ctx).
		Where("userId", userId).
		Where("followedId", in.FollowedId).
		Where("followType", in.FollowType).
		One()
	if err != nil {
		return nil, err
	}

	var (
		isFollowed bool
		followId   uint
		followTime *gtime.Time
	)

	// 开启事务
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		if followRecord.IsEmpty() {
			// 未关注，添加关注记录
			result, err := dao.Follows.Ctx(ctx).TX(tx).Insert(g.Map{
				"userId":     userId,
				"followedId": in.FollowedId,
				"followType": in.FollowType,
				"createTime": gtime.Now(),
			})
			if err != nil {
				return err
			}

			lastInsertId, err := result.LastInsertId()
			if err != nil {
				return err
			}
			followId = uint(lastInsertId)
			followTime = gtime.Now()
			isFollowed = true

			// 对于用户关注，表中已有触发器会自动更新粉丝数和关注数
			// 如果是圈子，需要更新圈子的关注数
			if in.FollowType == 2 { // 圈子
				_, err = dao.Circles.Ctx(ctx).TX(tx).
					Where("circleId", in.FollowedId).
					Increment("followerCount", 1)
				if err != nil {
					return err
				}
			}
		} else {
			// 已关注，取消关注
			_, err = dao.Follows.Ctx(ctx).TX(tx).
				Where("userId", userId).
				Where("followedId", in.FollowedId).
				Where("followType", in.FollowType).
				Delete()
			if err != nil {
				return err
			}

			isFollowed = false

			// 对于用户关注，表中已有触发器会自动更新粉丝数和关注数
			// 如果是圈子，需要更新圈子的关注数
			if in.FollowType == 2 { // 圈子
				_, err = dao.Circles.Ctx(ctx).TX(tx).
					Where("circleId", in.FollowedId).
					Where("followerCount > 0"). // 防止减为负数
					Decrement("followerCount", 1)
				if err != nil {
					return err
				}
			}
		}

		// TODO: 如果需要用Redis实现热门用户或圈子推荐，可以在这里更新Redis数据
		// 使用Redis的有序集合(Sorted Set)存储热门用户或圈子，以粉丝数作为分值
		// 示例代码:
		// key := fmt.Sprintf("hot_following:%s", in.FollowType)
		// redis.ZIncrBy(key, 1 或 -1, in.FollowedId)

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &model.FollowToggleOutput{
		FollowId:   followId,
		IsFollowed: isFollowed,
		FollowTime: followTime,
	}, nil
}

// FollowingList 获取关注列表
func (s *sFollow) FollowingList(ctx context.Context, in model.FollowingListReq) (*model.FollowingListOutput, error) {
	// 处理分页参数
	if in.Page <= 0 {
		in.Page = 1
	}
	if in.Size <= 0 {
		in.Size = 10
	}

	// 获取当前登录用户ID
	loginUserId := auth.GetLoginUserId(ctx)

	// 查询用户是否存在
	userCount, err := dao.Users.Ctx(ctx).
		Where("userId", in.UserId).
		Count()
	if err != nil {
		return nil, err
	}
	if userCount == 0 {
		return nil, gerror.New("用户不存在")
	}

	// 查询关注类型是否有效
	if in.FollowType != 1 && in.FollowType != 2 {
		return nil, gerror.New("无效的关注类型")
	}

	// 查询关注总数
	total, err := dao.Follows.Ctx(ctx).
		Where("userId", in.UserId).
		Where("followType", in.FollowType).
		Count()
	if err != nil {
		return nil, err
	}

	// 查询关注列表
	var records []gdb.Record
	err = dao.Follows.Ctx(ctx).
		Where("userId", in.UserId).
		Where("followType", in.FollowType).
		Order("createTime DESC").
		Page(in.Page, in.Size).
		Scan(&records)
	if err != nil {
		return nil, err
	}

	// 处理关注列表数据
	list := make([]model.FollowingItem, 0, len(records))
	for _, record := range records {
		followedId := record["followedId"].Int()

		// 根据关注类型获取被关注对象信息
		if in.FollowType == 1 { // 用户
			userInfo, err := dao.Users.Ctx(ctx).
				Where("userId", followedId).
				One()
			if err != nil {
				return nil, err
			}
			if userInfo.IsEmpty() {
				continue
			}

			// 检查当前登录用户是否关注了此用户
			isFollowed := false
			if loginUserId > 0 && uint(followedId) != loginUserId {
				count, err := dao.Follows.Ctx(ctx).
					Where("userId", loginUserId).
					Where("followedId", followedId).
					Where("followType", 1).
					Count()
				if err != nil {
					return nil, err
				}
				isFollowed = count > 0
			}

			list = append(list, model.FollowingItem{
				ID:            followedId,
				Name:          userInfo["name"].String(),
				Avatar:        userInfo["avatar"].String(),
				Description:   userInfo["introduction"].String(),
				FollowerCount: userInfo["followerCount"].Int(),
				IsFollowed:    isFollowed,
				FollowTime:    record["createTime"].GTime(),
			})
		} else if in.FollowType == 2 { // 圈子
			circleInfo, err := dao.Circles.Ctx(ctx).
				Where("circleId", followedId).
				One()
			if err != nil {
				return nil, err
			}
			if circleInfo.IsEmpty() {
				continue
			}

			// 检查当前登录用户是否关注了此圈子
			isFollowed := false
			if loginUserId > 0 {
				count, err := dao.Follows.Ctx(ctx).
					Where("userId", loginUserId).
					Where("followedId", followedId).
					Where("followType", 2).
					Count()
				if err != nil {
					return nil, err
				}
				isFollowed = count > 0
			}

			list = append(list, model.FollowingItem{
				ID:            followedId,
				Name:          circleInfo["name"].String(),
				Avatar:        circleInfo["logo"].String(),
				Description:   circleInfo["description"].String(),
				FollowerCount: circleInfo["followerCount"].Int(),
				IsFollowed:    isFollowed,
				FollowTime:    record["createTime"].GTime(),
			})
		}
	}

	// 返回结果
	out := &model.FollowingListOutput{
		List:  list,
		Total: int(total),
		Page:  in.Page,
		Size:  in.Size,
	}

	return out, nil
}

// FollowerList 获取粉丝列表
func (s *sFollow) FollowerList(ctx context.Context, in model.FollowerListReq) (*model.FollowerListOutput, error) {
	// 处理分页参数
	if in.Page <= 0 {
		in.Page = 1
	}
	if in.Size <= 0 {
		in.Size = 10
	}

	// 只支持查询用户的粉丝
	if in.FollowType != 1 {
		return nil, gerror.New("仅支持查询用户的粉丝")
	}

	// 获取当前登录用户ID
	loginUserId := auth.GetLoginUserId(ctx)

	// 查询用户是否存在
	userCount, err := dao.Users.Ctx(ctx).
		Where("userId", in.FollowedId).
		Count()
	if err != nil {
		return nil, err
	}
	if userCount == 0 {
		return nil, gerror.New("用户不存在")
	}

	// 查询粉丝总数
	total, err := dao.Follows.Ctx(ctx).
		Where("followedId", in.FollowedId).
		Where("followType", in.FollowType).
		Count()
	if err != nil {
		return nil, err
	}

	// 查询粉丝列表
	var records []gdb.Record
	err = dao.Follows.Ctx(ctx).
		Where("followedId", in.FollowedId).
		Where("followType", in.FollowType).
		Order("createTime DESC").
		Page(in.Page, in.Size).
		Scan(&records)
	if err != nil {
		return nil, err
	}

	// 处理粉丝列表数据
	list := make([]model.FollowerItem, 0, len(records))
	for _, record := range records {
		userId := record["userId"].Int()

		// 获取粉丝用户信息
		userInfo, err := dao.Users.Ctx(ctx).
			Where("userId", userId).
			One()
		if err != nil {
			return nil, err
		}
		if userInfo.IsEmpty() {
			continue
		}

		// 检查当前登录用户是否关注了此粉丝
		isFollowed := false
		if loginUserId > 0 && uint(userId) != loginUserId {
			count, err := dao.Follows.Ctx(ctx).
				Where("userId", loginUserId).
				Where("followedId", userId).
				Where("followType", 1).
				Count()
			if err != nil {
				return nil, err
			}
			isFollowed = count > 0
		}

		// 检查此粉丝是否是互相关注
		isFollowEachOther := false
		if in.FollowedId != userId {
			count, err := dao.Follows.Ctx(ctx).
				Where("userId", in.FollowedId).
				Where("followedId", userId).
				Where("followType", 1).
				Count()
			if err != nil {
				return nil, err
			}
			isFollowEachOther = count > 0
		}

		list = append(list, model.FollowerItem{
			ID:                userId,
			Name:              userInfo["name"].String(),
			Avatar:            userInfo["avatar"].String(),
			Description:       userInfo["introduction"].String(),
			FollowerCount:     userInfo["followerCount"].Int(),
			IsFollowed:        isFollowed,
			IsFollowEachOther: isFollowEachOther,
			FollowTime:        record["createTime"].GTime(),
		})
	}

	// 返回结果
	out := &model.FollowerListOutput{
		List:  list,
		Total: int(total),
		Page:  in.Page,
		Size:  in.Size,
	}

	return out, nil
}

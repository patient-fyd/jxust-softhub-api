package circle

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/patient-fyd/jxust-softhub-api/internal/dao"
	"github.com/patient-fyd/jxust-softhub-api/internal/model"
	"github.com/patient-fyd/jxust-softhub-api/internal/model/entity"
	"github.com/patient-fyd/jxust-softhub-api/internal/service"
	"github.com/patient-fyd/jxust-softhub-api/utility/auth"
)

type sCircle struct{}

func init() {
	service.RegisterCircle(New())
}

func New() *sCircle {
	return &sCircle{}
}

// List 获取圈子列表
func (s *sCircle) List(ctx context.Context, in model.CircleListInput) (*model.CircleListOutput, error) {
	var (
		m         = dao.Circles.Ctx(ctx)
		condition = g.Map{}
		out       = &model.CircleListOutput{
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
		condition["circleName LIKE ?"] = "%" + in.Keyword + "%"
	}

	// 只查询正常状态的圈子
	condition["status"] = 1

	// 如果指定了用户ID，则查询该用户关注的圈子
	if in.UserId > 0 {
		var followResults []struct {
			FollowedId int
		}
		err := dao.Follows.Ctx(ctx).
			Fields("followedId").
			Where("userId", in.UserId).
			Where("followType", 2). // 2表示关注圈子
			Scan(&followResults)
		if err != nil {
			return nil, err
		}

		// 提取圈子ID
		followCircleIds := make([]int, 0, len(followResults))
		for _, v := range followResults {
			followCircleIds = append(followCircleIds, v.FollowedId)
		}

		if len(followCircleIds) > 0 {
			condition["circleId IN(?)"] = followCircleIds
		} else {
			// 如果用户没有关注任何圈子，返回空结果
			out.List = make([]model.CircleListItem, 0)
			out.Total = 0
			return out, nil
		}
	}

	// 查询符合条件的记录总数
	count, err := m.Where(condition).Count()
	if err != nil {
		return nil, err
	}
	out.Total = count

	// 没有数据，直接返回
	if count == 0 {
		out.List = make([]model.CircleListItem, 0)
		return out, nil
	}

	// 排序规则
	orderBy := "createTime DESC"
	if in.OrderBy == "hot" {
		orderBy = "memberCount DESC, postCount DESC, createTime DESC"
	}

	// 查询列表数据
	var list []*entity.Circles
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
	out.List = make([]model.CircleListItem, 0, len(list))
	if len(list) > 0 {
		// 收集所有圈子ID
		circleIds := make([]int, 0, len(list))
		for _, v := range list {
			circleIds = append(circleIds, gconv.Int(v.CircleId))
		}

		// 获取当前用户已关注的圈子
		followedCircleMap := make(map[int]bool)
		if loginUserId > 0 && len(circleIds) > 0 {
			var followResults []struct {
				FollowedId int
			}
			err = dao.Follows.Ctx(ctx).
				Fields("followedId").
				Where("userId", loginUserId).
				Where("followType", 2). // 2表示关注圈子
				Where("followedId IN(?)", circleIds).
				Scan(&followResults)
			if err != nil {
				return nil, err
			}
			for _, v := range followResults {
				followedCircleMap[v.FollowedId] = true
			}
		}

		// 转换数据
		for _, v := range list {
			item := model.CircleListItem{
				CircleId:    gconv.Int(v.CircleId),
				Name:        v.CircleName,
				Description: v.Description,
				Icon:        v.Icon,
				PostCount:   gconv.Int(v.PostCount),
				MemberCount: gconv.Int(v.MemberCount),
				IsOfficial:  gconv.Int(v.IsOfficial),
				IsFollowed:  followedCircleMap[gconv.Int(v.CircleId)],
				CreateTime:  v.CreateTime,
			}
			out.List = append(out.List, item)
		}
	}

	return out, nil
}

// Detail 获取圈子详情
func (s *sCircle) Detail(ctx context.Context, in model.CircleDetailInput) (*model.CircleDetailOutput, error) {
	// 查询圈子信息
	var circle *entity.Circles
	err := dao.Circles.Ctx(ctx).
		Where("circleId", in.CircleId).
		Where("status", 1). // 只查询正常状态的圈子
		Scan(&circle)
	if err != nil {
		return nil, err
	}
	if circle == nil {
		return nil, nil
	}

	// 获取创建者信息
	var creator struct {
		UserName string
	}
	err = dao.Users.Ctx(ctx).
		Fields("userName").
		Where("userId", circle.UserId).
		Scan(&creator)
	if err != nil {
		return nil, err
	}

	// 当前登录用户ID
	loginUserId := auth.GetLoginUserId(ctx)

	// 判断当前用户是否已关注该圈子
	isFollowed := false
	if loginUserId > 0 {
		count, err := dao.Follows.Ctx(ctx).
			Where("userId", loginUserId).
			Where("followedId", in.CircleId).
			Where("followType", 2). // 2表示关注圈子
			Count()
		if err != nil {
			return nil, err
		}
		isFollowed = count > 0
	}

	// 转换结果
	return &model.CircleDetailOutput{
		CircleId:    gconv.Int(circle.CircleId),
		Name:        circle.CircleName,
		Description: circle.Description,
		Icon:        circle.Icon,
		PostCount:   gconv.Int(circle.PostCount),
		MemberCount: gconv.Int(circle.MemberCount),
		CreatorId:   gconv.Int(circle.UserId),
		CreatorName: creator.UserName,
		IsOfficial:  gconv.Int(circle.IsOfficial),
		IsFollowed:  isFollowed,
		CreateTime:  circle.CreateTime,
		UpdateTime:  circle.UpdateTime,
	}, nil
}

// Join 关注/取消关注圈子
func (s *sCircle) Join(ctx context.Context, in model.CircleJoinInput) (*model.CircleJoinOutput, error) {
	// 获取当前登录用户ID
	loginUserId := auth.GetLoginUserId(ctx)
	if loginUserId <= 0 {
		return nil, gerror.New("请先登录")
	}

	// 检查圈子是否存在
	var circle *entity.Circles
	err := dao.Circles.Ctx(ctx).
		Where("circleId", in.CircleId).
		Where("status", 1). // 只查询正常状态的圈子
		Scan(&circle)
	if err != nil {
		return nil, err
	}
	if circle == nil {
		return nil, gerror.New("圈子不存在或已删除")
	}

	// 判断当前用户是否已关注该圈子
	var isFollowed bool
	count, err := dao.Follows.Ctx(ctx).
		Where("userId", loginUserId).
		Where("followedId", in.CircleId).
		Where("followType", 2). // 2表示关注圈子
		Count()
	if err != nil {
		return nil, err
	}
	isFollowed = count > 0

	// 已关注，则取消关注；未关注，则添加关注
	if isFollowed {
		// 取消关注
		_, err = dao.Follows.Ctx(ctx).
			Where("userId", loginUserId).
			Where("followedId", in.CircleId).
			Where("followType", 2).
			Delete()
		if err != nil {
			return nil, err
		}

		// 更新圈子成员数量（减1）
		_, err = dao.Circles.Ctx(ctx).
			Data("memberCount=memberCount-1").
			Where("circleId", in.CircleId).
			Where("memberCount > 0").
			Update()
		if err != nil {
			return nil, err
		}

		return &model.CircleJoinOutput{
			IsFollowed: false,
		}, nil
	} else {
		// 添加关注
		_, err = dao.Follows.Ctx(ctx).Insert(g.Map{
			"userId":     loginUserId,
			"followedId": in.CircleId,
			"followType": 2, // 2表示关注圈子
			"createTime": gtime.Now(),
		})
		if err != nil {
			return nil, err
		}

		// 更新圈子成员数量（加1）
		_, err = dao.Circles.Ctx(ctx).
			Data("memberCount=memberCount+1").
			Where("circleId", in.CircleId).
			Update()
		if err != nil {
			return nil, err
		}

		return &model.CircleJoinOutput{
			IsFollowed: true,
		}, nil
	}
}

// MyCircles 获取我关注的圈子列表
func (s *sCircle) MyCircles(ctx context.Context, in model.CircleMyCirclesInput) (*model.CircleMyCirclesOutput, error) {
	// 默认分页参数
	if in.Page <= 0 {
		in.Page = 1
	}
	if in.Size <= 0 {
		in.Size = 10
	}

	// 要查询的用户ID
	userId := in.UserId
	if userId <= 0 {
		// 获取当前登录用户ID
		userId = gconv.Int(auth.GetLoginUserId(ctx))
		if userId <= 0 {
			return nil, gerror.New("请先登录")
		}
	}

	// 获取用户关注的圈子ID列表
	var followResults []struct {
		FollowedId int
	}
	err := dao.Follows.Ctx(ctx).
		Fields("followedId").
		Where("userId", userId).
		Where("followType", 2). // 2表示关注圈子
		Scan(&followResults)
	if err != nil {
		return nil, err
	}

	// 提取圈子ID
	circleIds := make([]int, 0, len(followResults))
	for _, v := range followResults {
		circleIds = append(circleIds, v.FollowedId)
	}

	// 准备输出结构
	out := &model.CircleMyCirclesOutput{
		Page: in.Page,
		Size: in.Size,
		List: make([]model.CircleListItem, 0),
	}

	// 如果没有关注任何圈子，直接返回空结果
	if len(circleIds) == 0 {
		out.Total = 0
		return out, nil
	}

	// 查询符合条件的圈子总数
	count, err := dao.Circles.Ctx(ctx).
		Where("circleId IN(?)", circleIds).
		Where("status", 1). // 只查询正常状态的圈子
		Count()
	if err != nil {
		return nil, err
	}
	out.Total = count

	// 如果没有数据，直接返回
	if count == 0 {
		return out, nil
	}

	// 查询圈子数据
	var list []*entity.Circles
	err = dao.Circles.Ctx(ctx).
		Where("circleId IN(?)", circleIds).
		Where("status", 1). // 只查询正常状态的圈子
		Page(in.Page, in.Size).
		Order("createTime DESC").
		Scan(&list)
	if err != nil {
		return nil, err
	}

	// 转换结果格式
	for _, v := range list {
		item := model.CircleListItem{
			CircleId:    gconv.Int(v.CircleId),
			Name:        v.CircleName,
			Description: v.Description,
			Icon:        v.Icon,
			PostCount:   gconv.Int(v.PostCount),
			MemberCount: gconv.Int(v.MemberCount),
			IsOfficial:  gconv.Int(v.IsOfficial),
			IsFollowed:  true, // 这里一定是已关注的
			CreateTime:  v.CreateTime,
		}
		out.List = append(out.List, item)
	}

	return out, nil
}

// CircleStat 获取圈子统计信息
func (s *sCircle) CircleStat(ctx context.Context, in model.CircleStatInput) (*model.CircleStatOutput, error) {
	// 要查询的用户ID
	userId := in.UserId
	if userId <= 0 {
		// 获取当前登录用户ID
		userId = gconv.Int(auth.GetLoginUserId(ctx))
		if userId <= 0 {
			return nil, gerror.New("请先登录")
		}
	}

	// 准备输出结构
	out := &model.CircleStatOutput{
		RecentActive: make([]model.CircleListItem, 0),
	}

	// 查询圈子总数
	totalCount, err := dao.Circles.Ctx(ctx).
		Where("status", 1). // 只查询正常状态的圈子
		Count()
	if err != nil {
		return nil, err
	}
	out.TotalCount = totalCount

	// 查询用户关注的圈子数量
	followingCount, err := dao.Follows.Ctx(ctx).
		Where("userId", userId).
		Where("followType", 2). // 2表示关注圈子
		Count()
	if err != nil {
		return nil, err
	}
	out.FollowingCount = followingCount

	// 查询最近活跃的圈子（按成员数和帖子数排序）
	var list []*entity.Circles
	err = dao.Circles.Ctx(ctx).
		Where("status", 1). // 只查询正常状态的圈子
		Order("memberCount DESC, postCount DESC, createTime DESC").
		Limit(5).
		Scan(&list)
	if err != nil {
		return nil, err
	}

	// 获取当前用户已关注的圈子
	followedCircleMap := make(map[int]bool)
	if userId > 0 && len(list) > 0 {
		circleIds := make([]int, 0, len(list))
		for _, v := range list {
			circleIds = append(circleIds, gconv.Int(v.CircleId))
		}

		var followResults []struct {
			FollowedId int
		}
		err = dao.Follows.Ctx(ctx).
			Fields("followedId").
			Where("userId", userId).
			Where("followType", 2). // 2表示关注圈子
			Where("followedId IN(?)", circleIds).
			Scan(&followResults)
		if err != nil {
			return nil, err
		}
		for _, v := range followResults {
			followedCircleMap[v.FollowedId] = true
		}
	}

	// 转换结果格式
	for _, v := range list {
		item := model.CircleListItem{
			CircleId:    gconv.Int(v.CircleId),
			Name:        v.CircleName,
			Description: v.Description,
			Icon:        v.Icon,
			PostCount:   gconv.Int(v.PostCount),
			MemberCount: gconv.Int(v.MemberCount),
			IsOfficial:  gconv.Int(v.IsOfficial),
			IsFollowed:  followedCircleMap[gconv.Int(v.CircleId)],
			CreateTime:  v.CreateTime,
		}
		out.RecentActive = append(out.RecentActive, item)
	}

	return out, nil
}

package member

import (
	"context"
	"strconv"

	v1 "github.com/patient-fyd/jxust-softhub-api/api/member/v1"
	"github.com/patient-fyd/jxust-softhub-api/internal/model"
	"github.com/patient-fyd/jxust-softhub-api/internal/service"
)

// 将字符串类型的入会年份转成int类型
func parseJoinYear(yearStr string) int {
	year, err := strconv.Atoi(yearStr)
	if err != nil {
		// 如果转换失败，默认返回当前年份
		return 2025
	}
	return year
}

func (c *ControllerV1) List(ctx context.Context, req *v1.ListReq) (res *v1.ListRes, err error) {
	// 调用逻辑层获取成员列表
	output, err := service.Member().List(ctx, model.MemberListInput{
		Department: req.Department,
		Grade:      req.Grade,
		IsCore:     req.IsCore,
		Status:     req.Status,
		PageNum:    req.PageNum,
		PageSize:   req.PageSize,
	})
	if err != nil {
		return nil, err
	}

	// 构建响应数据
	res = &v1.ListRes{
		Total:    output.Total,
		PageNum:  output.PageNum,
		PageSize: output.PageSize,
		List:     make([]v1.MemberInfo, 0, len(output.List)),
	}

	// 转换成员信息列表
	for _, item := range output.List {
		res.List = append(res.List, v1.MemberInfo{
			MemberId:     item.MemberId,
			UserId:       item.UserId,
			UserName:     item.UserName,
			Name:         item.Name,
			Avatar:       item.Avatar,
			Grade:        item.Grade,
			JoinYear:     parseJoinYear(item.JoinYear),
			Department:   item.Department,
			Position:     item.Position,
			Skills:       item.Skills,
			Introduction: item.Introduction,
			IsCore:       item.IsCore,
			Status:       item.Status,
		})
	}

	return res, nil
}

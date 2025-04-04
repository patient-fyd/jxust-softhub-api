package news

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	v1 "github.com/patient-fyd/jxust-softhub-api/api/news/v1"
	"github.com/patient-fyd/jxust-softhub-api/internal/model"
	"github.com/patient-fyd/jxust-softhub-api/internal/service"
)

// Update 更新新闻
func (c *ControllerV1) Update(ctx context.Context, req *v1.UpdateReq) (res *v1.UpdateRes, err error) {
	// 检查新闻类型是否有变化
	// 先获取原有新闻详情
	detail, err := service.News().Detail(ctx, model.NewsDetailInput{NewsId: req.NewsId})
	if err != nil {
		return nil, gerror.NewCode(gcode.CodeInternalError, "获取新闻详情失败: "+err.Error())
	}
	if detail == nil {
		return nil, gerror.NewCode(gcode.CodeNotFound, "新闻不存在")
	}

	// 如果修改了新闻类型，需要重新检查权限
	if req.NewsType > 0 && req.NewsType != detail.NewsType {
		if req.NewsType == 1 { // 协会通知
			// 检查是否拥有发布协会通知的权限（权限ID: 6）
			hasPermission, err := service.Permission().HasPermission(ctx, 6)
			if err != nil {
				return nil, gerror.NewCode(gcode.CodeInternalError, "检查权限失败: "+err.Error())
			}
			if !hasPermission {
				return nil, gerror.NewCode(gcode.CodeOperationFailed, "您没有发布协会通知的权限")
			}
		} else if req.NewsType == 2 { // 技术分享
			// 检查是否拥有发布技术文章的权限（权限ID: 5）
			hasPermission, err := service.Permission().HasPermission(ctx, 5)
			if err != nil {
				return nil, gerror.NewCode(gcode.CodeInternalError, "检查权限失败: "+err.Error())
			}
			if !hasPermission {
				return nil, gerror.NewCode(gcode.CodeOperationFailed, "您没有发布技术文章的权限")
			}
		}
	}

	// 构造更新参数
	input := model.NewsUpdateInput{
		NewsId:     req.NewsId,
		Title:      req.Title,
		Content:    req.Content,
		Category:   req.Category,
		NewsType:   req.NewsType,
		CoverImage: req.CoverImage,
		IsTop:      req.IsTop,
		Status:     req.Status,
	}

	// 调用服务
	output, err := service.News().Update(ctx, input)
	if err != nil {
		return nil, gerror.NewCode(gcode.CodeInternalError, err.Error())
	}

	// 如果操作失败
	if !output.Success {
		return nil, gerror.NewCode(gcode.CodeNotFound, "新闻不存在")
	}

	// 转换响应
	res = &v1.UpdateRes{
		Success: output.Success,
	}

	return res, nil
}

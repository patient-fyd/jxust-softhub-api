package news

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	v1 "github.com/patient-fyd/jxust-softhub-api/api/news/v1"
	"github.com/patient-fyd/jxust-softhub-api/internal/model"
	"github.com/patient-fyd/jxust-softhub-api/internal/service"
)

// Create 创建新闻
func (c *ControllerV1) Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error) {
	// 检查权限
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

	// 构造创建参数
	input := model.NewsCreateInput{
		Title:      req.Title,
		Content:    req.Content,
		Category:   req.Category,
		NewsType:   req.NewsType,
		CoverImage: req.CoverImage,
		IsTop:      req.IsTop,
		Status:     req.Status,
	}

	// 调用服务
	output, err := service.News().Create(ctx, input)
	if err != nil {
		return nil, gerror.NewCode(gcode.CodeInternalError, err.Error())
	}

	// 转换响应
	res = &v1.CreateRes{
		Id: output.Id,
	}

	return res, nil
}

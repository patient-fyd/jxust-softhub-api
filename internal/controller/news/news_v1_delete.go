package news

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	v1 "github.com/patient-fyd/jxust-softhub-api/api/news/v1"
	"github.com/patient-fyd/jxust-softhub-api/internal/model"
	"github.com/patient-fyd/jxust-softhub-api/internal/service"
)

// Delete 删除新闻
func (c *ControllerV1) Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error) {
	// 获取新闻详情，检查类型并验证权限
	detail, err := service.News().Detail(ctx, model.NewsDetailInput{NewsId: req.NewsId})
	if err != nil {
		return nil, gerror.NewCode(gcode.CodeInternalError, "获取新闻详情失败: "+err.Error())
	}
	if detail == nil {
		return nil, gerror.NewCode(gcode.CodeNotFound, "新闻不存在")
	}

	// 根据新闻类型检查权限
	if detail.NewsType == 1 { // 协会通知
		// 检查是否拥有管理协会通知的权限（权限ID: 6）
		hasPermission, err := service.Permission().HasPermission(ctx, 6)
		if err != nil {
			return nil, gerror.NewCode(gcode.CodeInternalError, "检查权限失败: "+err.Error())
		}
		if !hasPermission {
			return nil, gerror.NewCode(gcode.CodeOperationFailed, "您没有管理协会通知的权限")
		}
	} else if detail.NewsType == 2 { // 技术分享
		// 检查是否拥有管理技术文章的权限（权限ID: 5）
		hasPermission, err := service.Permission().HasPermission(ctx, 5)
		if err != nil {
			return nil, gerror.NewCode(gcode.CodeInternalError, "检查权限失败: "+err.Error())
		}
		if !hasPermission {
			return nil, gerror.NewCode(gcode.CodeOperationFailed, "您没有管理技术文章的权限")
		}
	}

	// 构造删除参数
	input := model.NewsDeleteInput{
		NewsId: req.NewsId,
	}

	// 调用服务
	output, err := service.News().Delete(ctx, input)
	if err != nil {
		return nil, gerror.NewCode(gcode.CodeInternalError, err.Error())
	}

	// 如果操作失败
	if !output.Success {
		return nil, gerror.NewCode(gcode.CodeNotFound, "新闻不存在")
	}

	// 转换响应
	res = &v1.DeleteRes{
		Success: output.Success,
	}

	return res, nil
}

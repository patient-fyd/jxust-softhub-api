package post

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"

	v1 "github.com/patient-fyd/jxust-softhub-api/api/post/v1"
	"github.com/patient-fyd/jxust-softhub-api/internal/model"
	"github.com/patient-fyd/jxust-softhub-api/internal/service"
)

func (c *ControllerV1) Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error) {
	// 安全检查
	if ctx == nil {
		return nil, gerror.New("上下文为空")
	}

	if c == nil {
		g.Log().Error(ctx, "控制器实例为空")
		return nil, gerror.New("系统错误: 控制器实例为空")
	}

	// 检查请求体
	if req == nil {
		g.Log().Error(ctx, "请求体为空")
		return nil, gerror.New("请求体不能为空")
	}

	// 添加详细调试日志
	g.Log().Debug(ctx, "接收到帖子创建请求")
	g.Log().Debugf(ctx, "请求内容: %+v", req)
	g.Log().Debugf(ctx, "内容长度: %d", len(req.Content))

	// 检查必填字段
	if len(req.Content) == 0 {
		g.Log().Error(ctx, "帖子内容为空")
		return nil, gerror.New("帖子内容不能为空")
	}

	// 转换请求参数
	input := model.PostCreateInput{
		Content:  req.Content,
		Images:   req.Images,
		CircleId: req.CircleId,
		TopicId:  req.TopicId,
	}

	// 初始化空图片数组，防止空指针
	if input.Images == nil {
		input.Images = make([]string, 0)
	}

	// 打印处理后的输入参数
	g.Log().Debugf(ctx, "处理后的输入参数: %+v", input)

	// 安全检查服务层实例
	postService := service.Post()
	if postService == nil {
		g.Log().Error(ctx, "帖子服务实例为空")
		return nil, gerror.New("系统错误: 服务未正确初始化")
	}
	g.Log().Debug(ctx, "帖子服务实例获取成功")

	// 调用业务逻辑层
	g.Log().Debug(ctx, "开始调用Post服务创建帖子")

	// 使用recover处理可能的panic
	var output *model.PostCreateOutput
	func() {
		defer func() {
			if r := recover(); r != nil {
				g.Log().Errorf(ctx, "创建帖子时发生panic: %v", r)
				err = gerror.Newf("系统异常: %v", r)
			}
		}()
		output, err = postService.Create(ctx, input)
	}()

	if err != nil {
		g.Log().Errorf(ctx, "帖子创建失败: %v", err)
		return nil, gerror.Wrap(err, "创建帖子失败")
	}

	// 验证输出结果
	if output == nil {
		g.Log().Error(ctx, "帖子创建返回nil")
		return nil, gerror.New("帖子创建失败，返回结果为空")
	}

	// 转换响应参数
	g.Log().Debugf(ctx, "帖子创建成功，ID: %d", output.PostId)
	return &v1.CreateRes{
		PostId: output.PostId,
	}, nil
}

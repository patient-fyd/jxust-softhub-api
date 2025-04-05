package apiserver

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/glog"

	"github.com/patient-fyd/jxust-softhub-api/internal/controller/auth"
	"github.com/patient-fyd/jxust-softhub-api/internal/controller/comment"
	configCtrl "github.com/patient-fyd/jxust-softhub-api/internal/controller/config"
	"github.com/patient-fyd/jxust-softhub-api/internal/controller/file"
	"github.com/patient-fyd/jxust-softhub-api/internal/controller/follow"
	"github.com/patient-fyd/jxust-softhub-api/internal/controller/join"
	"github.com/patient-fyd/jxust-softhub-api/internal/controller/like"
	"github.com/patient-fyd/jxust-softhub-api/internal/controller/member"
	"github.com/patient-fyd/jxust-softhub-api/internal/controller/news"
	"github.com/patient-fyd/jxust-softhub-api/internal/controller/post"
	"github.com/patient-fyd/jxust-softhub-api/internal/controller/stat"
	"github.com/patient-fyd/jxust-softhub-api/internal/controller/tag"
	"github.com/patient-fyd/jxust-softhub-api/internal/controller/user"

	// 确保auth/user逻辑包被导入并执行其init函数
	_ "github.com/patient-fyd/jxust-softhub-api/internal/logic/auth"
	_ "github.com/patient-fyd/jxust-softhub-api/internal/logic/join"
	_ "github.com/patient-fyd/jxust-softhub-api/internal/logic/news"
	_ "github.com/patient-fyd/jxust-softhub-api/internal/logic/user"

	"github.com/patient-fyd/jxust-softhub-api/internal/service"
	"github.com/patient-fyd/jxust-softhub-api/utility"
)

var (
	Main = gcmd.Command{
		Name:        "jxust-softhub-api-api",
		Brief:       "An API Server Demo",
		Description: "An API server demo using GoFrame V2",
		Usage:       "jxust-softhub-api-api [OPTION]",
		Examples: `
			Dev:
				./jxust-softhub-api-api

			Test:
				./jxust-softhub-api-api -c config.test.yaml
				or 
				GF_GCFG_FILE=config.test.yaml GF_GERROR_BRIEF=true ./jxust-softhub-api-api

			Prod:
				./jxust-softhub-api-api -c config.prod.yaml
				or 
				GF_GCFG_FILE=config.prod.yaml GF_GERROR_BRIEF=true ./jxust-softhub-api-api`,
		Additional: "Find more information at: https://github.com/patient-fyd/jxust-softhub-api",
		Arguments: []gcmd.Argument{
			{
				Name:   "version",
				Short:  "v",
				Brief:  "print version info",
				IsArg:  false,
				Orphan: true,
			},
			{
				Name:   "config",
				Short:  "c",
				Brief:  "config file (default config.yaml)",
				IsArg:  false,
				Orphan: false,
			},
		},
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			// 判断不带数据的选项是否存在时，可以通过GetOpt(name) != nil方式
			ver := parser.GetOpt("version")
			if ver != nil {
				utility.PrintVersionInfo()
				return nil
			}

			config := parser.GetOpt("config").String()
			if config != "" {
				//nolint: forcetypeassert
				g.Cfg().GetAdapter().(*gcfg.AdapterFile).SetFileName(config)
			}

			// json格式日志
			logFormat, err := g.Cfg().Get(ctx, "logger.format")
			if err == nil {
				if logFormat.String() == "json" {
					glog.SetDefaultHandler(glog.HandlerJson)
				}
			}

			// 异步打印日志 & 显示打印错误的文件行号, 对访问日志无效
			g.Log().SetFlags(glog.F_ASYNC | glog.F_TIME_STD | glog.F_FILE_LONG)

			//nolint:forcetypeassert
			configFile := g.Cfg().GetAdapter().(*gcfg.AdapterFile).GetFileName()
			g.Log().Debugf(ctx, "use config file: %s", configFile)

			s := g.Server()
			s.Use(
				service.Middleware().TraceID,
				service.Middleware().AccessUser,
				service.Middleware().AuthMiddleware,
				service.Middleware().ResponseHandler,
			)

			// 注册OpenAPI文档接口，不需要鉴权
			s.Group("/", func(group *ghttp.RouterGroup) {
				// 注册认证相关接口
				group.Bind(
					auth.NewV1(),
				)

				// 注册成员管理相关接口
				group.Bind(
					member.NewV1(),
				)

				// 注册入会申请相关接口
				group.Bind(
					join.NewV1(),
				)

				// 注册文件管理相关接口
				group.Bind(
					file.NewV1(),
				)

				// 注册新闻管理相关接口
				group.Bind(
					news.NewV1(),
				)

				// 注册统计分析相关接口
				group.Bind(
					stat.NewV1(),
				)

				// 注册标签管理相关接口
				group.Bind(
					tag.NewV1(),
				)

				// 注册系统配置相关接口
				group.Bind(
					configCtrl.NewV1(),
				)

				// 注册用户相关接口
				group.Bind(
					user.NewV1(),
				)

				// 注册帖子相关接口
				group.Bind(
					post.NewV1(),
				)

				// 注册评论相关接口
				group.Bind(
					comment.NewV1(),
				)

				// 注册点赞相关接口
				group.Bind(
					like.NewV1(),
				)

				// 注册关注相关接口
				group.Bind(
					follow.NewV1(),
				)
			})

			s.Run()
			return nil
		},
	}
)

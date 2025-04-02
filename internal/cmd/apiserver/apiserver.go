package apiserver

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/glog"

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
				service.Middleware().ResponseHandler,
			)
			s.Group("/v1", func(group *ghttp.RouterGroup) {
				// TODO: 在此处添加新的控制器绑定
				// group.Bind(
				//     // 在此处添加新的控制器
				// )
			})
			s.Run()
			return nil
		},
	}
)

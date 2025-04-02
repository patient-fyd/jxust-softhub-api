package cli

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/glog"

	"github.com/patient-fyd/jxust-softhub-api/internal/consts"
	"github.com/patient-fyd/jxust-softhub-api/utility"
)

var (
	Main = gcmd.Command{
		Name:        "jxust-softhub-api-cli",
		Brief:       "A command-line tool demo",
		Description: "A command-line tool demo using GoFrame V2",
		Usage:       "jxust-softhub-api-cli [OPTION]",
		Examples: `
			Dev:
				./jxust-softhub-api-cli

			Test:
				./jxust-softhub-api-cli -c config.test.yaml
				or 
				GF_GCFG_FILE=config.test.yaml ./jxust-softhub-api-cli

			Prod:
				./jxust-softhub-api-cli -c config.prod.yaml
				or 
				GF_GCFG_FILE=config.prod.yaml ./jxust-softhub-api-cli`,
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
			logFormat, err := g.Cfg().Get(ctx, "logger.cli.format")
			if err == nil {
				if logFormat.String() == "json" {
					glog.SetDefaultHandler(glog.HandlerJson)
				}
			}

			// 显示打印错误的文件行号
			g.Log(consts.CliLoggerName).SetFlags(glog.F_TIME_STD | glog.F_FILE_LONG)

			//nolint:forcetypeassert
			configFile := g.Cfg().GetAdapter().(*gcfg.AdapterFile).GetFileName()
			g.Log(consts.CliLoggerName).Debugf(ctx, "use config file: %s", configFile)

			// ****************** 以下部分为业务逻辑

			g.Log(consts.CliLoggerName).Info(ctx, "foo")

			return nil
		},
	}
)

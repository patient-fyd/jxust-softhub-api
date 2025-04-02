package main

import (
	"github.com/gogf/gf/v2/os/gctx"

	_ "github.com/patient-fyd/jxust-softhub-api/internal/logic"
	_ "github.com/patient-fyd/jxust-softhub-api/internal/packed"

	"github.com/patient-fyd/jxust-softhub-api/internal/cmd/cli"
)

func main() {
	cli.Main.Run(gctx.New())
}

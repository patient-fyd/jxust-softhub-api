package main

import (
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"

	"github.com/gogf/gf/v2/os/gctx"

	_ "github.com/patient-fyd/jxust-softhub-api/internal/logic"
	_ "github.com/patient-fyd/jxust-softhub-api/internal/packed"

	"github.com/patient-fyd/jxust-softhub-api/internal/cmd/apiserver"
)

func main() {
	apiserver.Main.Run(gctx.New())
}

package middleware

import (
	"net/http"
	"strings"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gvalid"

	"github.com/patient-fyd/jxust-softhub-api/internal/codes"
)

type response struct {
	Code    int         `json:"code"`
	Message string      `json:"msg"`
	TraceID string      `json:"traceid"`
	Data    interface{} `json:"data"`
}

// ResponseHandler custom response format.
func (s *sMiddleware) ResponseHandler(r *ghttp.Request) {
	r.Middleware.Next()

	// 跳过以下路径的响应处理
	skipPaths := []string{
		"/api.json",
		"/swagger",
		"/openapi.json",
	}

	for _, path := range skipPaths {
		if strings.HasPrefix(r.Request.URL.Path, path) {
			return
		}
	}

	// For PProf
	if g.Cfg().MustGet(r.Context(), "server.pprofEnabled").Bool() {
		rPath := g.Cfg().MustGet(r.Context(), "server.pprofPattern").String()
		if rPath == "" {
			rPath = "/debug/pprof"
		}
		if strings.HasPrefix(r.Request.URL.Path, rPath) {
			return
		}
	}

	// Clean exist response info.
	// i.e.:
	//	 1) gf exception recovered info
	//	 2) gf 404 Not Found content
	r.Response.ClearBuffer()

	var (
		err     = r.GetError()
		res     = r.GetHandlerResponse()
		msg     string
		bizCode codes.BizCode
		ok      bool
	)

	if err != nil {
		if _, ok = err.(gvalid.Error); ok { // validation error
			err = gerror.WrapCode(codes.CodeValidationFailed, err)
			code1 := gerror.Code(err)
			//nolint: errcheck
			bizCode, _ = code1.(codes.BizCode)
		} else {
			code := gerror.Code(err)
			bizCode, ok = code.(codes.BizCode) // custom error codes
			if !ok {
				err = gerror.NewCode(codes.CodeInternal) // internal error
				code1 := gerror.Code(err)
				//nolint: errcheck
				bizCode, _ = code1.(codes.BizCode)
			}
		}

		msg = err.Error()
	} else {
		if r.Response.Status == http.StatusNotFound { // gf internal 404 error
			//nolint: errcheck
			bizCode, _ = codes.CodeNotFound.(codes.BizCode)
		} else {
			//nolint: errcheck
			bizCode, _ = codes.CodeOK.(codes.BizCode)
		}

		msg = bizCode.Message()
	}

	r.Response.WriteHeader(bizCode.BizDetail().HTTPCode)

	// 解析bizCode的Code字段为int类型
	var codeValue int
	if bizCode.BizDetail().Code == "OK" {
		codeValue = 0
	} else {
		// 其他错误情况设置为对应的错误码，这里统一设置为-1，如有需要可以细化
		codeValue = -1
	}

	r.Response.WriteJsonExit(response{
		Code:    codeValue,
		Message: msg,
		TraceID: gctx.CtxId(r.Context()),
		Data:    res,
	})
}

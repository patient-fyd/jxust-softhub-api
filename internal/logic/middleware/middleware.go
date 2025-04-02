package middleware

import "github.com/patient-fyd/jxust-softhub-api/internal/service"

type (
	sMiddleware struct{}
)

func init() {
	service.RegisterMiddleware(New())
}

func New() *sMiddleware {
	return &sMiddleware{}
}

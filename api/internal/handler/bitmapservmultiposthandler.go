package handler

import (
	"net/http"

	"github.com/weblfe/bitmap_service/api/internal/svc"
)

func BitMapServMultiPostHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return BitMapServMultiGetHandler(ctx)
}

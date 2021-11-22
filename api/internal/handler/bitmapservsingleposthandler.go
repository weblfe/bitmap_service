package handler

import (
	"net/http"

	"github.com/weblfe/bitmap_service/api/internal/svc"
)

func BitMapServSinglePostHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return BitMapServSingleGetHandler(ctx)
}

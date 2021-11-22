package handler

import (
		"github.com/tal-tech/go-zero/rest/httpx"
		"github.com/weblfe/bitmap_service/api/internal/logic"
		"github.com/weblfe/bitmap_service/api/internal/svc"
		"github.com/weblfe/bitmap_service/api/internal/types"
		"net/http"
)

func GetIndexHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewGetIndexLogic(r.Context(), ctx)
		resp, err := l.GetIndex()
		if err != nil {
			httpx.OkJson(w, errorRes(err))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

// 错误
func errorRes(err error) interface{} {
	return &types.BaseResponse{
		Msg:     err.Error(),
		Code:    200,
		ErrorNo: 500,
	}
}

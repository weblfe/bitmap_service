package handler

import (
	"net/http"

	"github.com/weblfe/bitmap_service/api/internal/logic"
	"github.com/weblfe/bitmap_service/api/internal/svc"
	"github.com/weblfe/bitmap_service/api/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func BitMapServSingleGetHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SingleRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, &types.SingleResponse{
				BaseResponse: types.BaseResponse{
					Code:    200,
					Msg:     err.Error(),
					ErrorNo: 1400,
				},
				Data: nil,
			})
			return
		}

		l := logic.NewBitMapServSingleLogic(r.Context(), ctx)
		resp, err := l.BitMapServSingle(req)
		if err != nil {
			httpx.OkJson(w, &types.SingleResponse{
				BaseResponse: types.BaseResponse{
					Code:    200,
					Msg:     err.Error(),
					ErrorNo: 500,
				},
				Data: nil,
			})
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

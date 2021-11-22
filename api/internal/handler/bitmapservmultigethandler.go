package handler

import (
	"net/http"

	"github.com/weblfe/bitmap_service/api/internal/logic"
	"github.com/weblfe/bitmap_service/api/internal/svc"
	"github.com/weblfe/bitmap_service/api/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func BitMapServMultiGetHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.MulitRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, &types.MulitResponse{
				BaseResponse: types.BaseResponse{
					Code:    200,
					Msg:     err.Error(),
					ErrorNo: 1400,
				},
				Data: nil,
			})
			return
		}

		l := logic.NewBitMapServMultiLogic(r.Context(), ctx)
		resp, err := l.BitMapServMulti(req)
		if err != nil {
			httpx.OkJson(w, &types.MulitResponse{
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

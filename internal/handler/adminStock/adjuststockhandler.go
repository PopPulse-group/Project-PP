// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package adminStock

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"poppulse/internal/logic/adminStock"
	"poppulse/internal/svc"
	"poppulse/internal/types"
)

// M20 调整可售库存
func AdjustStockHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AdjustStockReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := adminStock.NewAdjustStockLogic(r.Context(), svcCtx)
		resp, err := l.AdjustStock(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

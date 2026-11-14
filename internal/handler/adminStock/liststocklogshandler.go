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

// M21 库存流水
func ListStockLogsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.StockLogListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := adminStock.NewListStockLogsLogic(r.Context(), svcCtx)
		resp, err := l.ListStockLogs(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package adminProduct

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"poppulse/internal/logic/adminProduct"
	"poppulse/internal/svc"
	"poppulse/internal/types"
)

// M17 新增 SKU
func CreateSkuHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateSkuReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := adminProduct.NewCreateSkuLogic(r.Context(), svcCtx)
		resp, err := l.CreateSku(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package adminAftersale

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"poppulse/internal/logic/adminAftersale"
	"poppulse/internal/svc"
	"poppulse/internal/types"
)

// M37 审核售后
func AdminReviewAfterSaleHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AdminReviewAfterSaleReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := adminAftersale.NewAdminReviewAfterSaleLogic(r.Context(), svcCtx)
		resp, err := l.AdminReviewAfterSale(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

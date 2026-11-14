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

// M38 模拟退款成功
func AdminMockRefundConfirmHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AdminMockRefundConfirmReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := adminAftersale.NewAdminMockRefundConfirmLogic(r.Context(), svcCtx)
		resp, err := l.AdminMockRefundConfirm(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

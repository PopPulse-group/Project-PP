// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package payment

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"poppulse/internal/logic/payment"
	"poppulse/internal/svc"
	"poppulse/internal/types"
)

// M26 模拟支付成功
func MockPayConfirmHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.MockPayConfirmReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := payment.NewMockPayConfirmLogic(r.Context(), svcCtx)
		resp, err := l.MockPayConfirm(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

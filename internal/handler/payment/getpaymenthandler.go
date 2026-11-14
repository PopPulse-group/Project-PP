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

// M27 查询本人支付结果
func GetPaymentHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetPaymentReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := payment.NewGetPaymentLogic(r.Context(), svcCtx)
		resp, err := l.GetPayment(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

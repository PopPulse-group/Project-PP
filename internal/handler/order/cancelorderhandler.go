// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package order

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"vagetable-gozero/internal/logic/order"
	"vagetable-gozero/internal/svc"
	"vagetable-gozero/internal/types"
)

// 取消待支付订单（一次性释放库存）
func CancelOrderHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.OrderIdReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := order.NewCancelOrderLogic(r.Context(), svcCtx)
		resp, err := l.CancelOrder(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

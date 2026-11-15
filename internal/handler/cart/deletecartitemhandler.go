// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package cart

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"vagetable-gozero/internal/logic/cart"
	"vagetable-gozero/internal/svc"
	"vagetable-gozero/internal/types"
)

// 删除明细
func DeleteCartItemHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CartItemIdReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := cart.NewDeleteCartItemLogic(r.Context(), svcCtx)
		resp, err := l.DeleteCartItem(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

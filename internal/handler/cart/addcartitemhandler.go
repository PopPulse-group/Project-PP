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

// 加购（同商品合并数量）
func AddCartItemHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddCartItemReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := cart.NewAddCartItemLogic(r.Context(), svcCtx)
		resp, err := l.AddCartItem(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

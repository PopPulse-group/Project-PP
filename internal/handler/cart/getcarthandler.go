// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package cart

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"vagetable-gozero/internal/logic/cart"
	"vagetable-gozero/internal/svc"
)

// 查看购物车（实时价格与库存）
func GetCartHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := cart.NewGetCartLogic(r.Context(), svcCtx)
		resp, err := l.GetCart()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

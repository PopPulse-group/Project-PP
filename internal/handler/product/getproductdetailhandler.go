// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package product

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"vagetable-gozero/internal/logic/product"
	"vagetable-gozero/internal/svc"
	"vagetable-gozero/internal/types"
)

// 商品详情（含可售库存）
func GetProductDetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ProductDetailReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := product.NewGetProductDetailLogic(r.Context(), svcCtx)
		resp, err := l.GetProductDetail(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

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

// 商品列表（分类/关键词/分页）
func ListProductsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ProductListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := product.NewListProductsLogic(r.Context(), svcCtx)
		resp, err := l.ListProducts(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

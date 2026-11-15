// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package product

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"poppulse/internal/logic/product"
	"poppulse/internal/svc"
)

// 分类列表
func ListCategoriesHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := product.NewListCategoriesLogic(r.Context(), svcCtx)
		resp, err := l.ListCategories()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

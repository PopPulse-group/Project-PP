// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package product

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"poppulse/internal/logic/product"
	"poppulse/internal/svc"
	"poppulse/internal/types"
)

// M06 IP 列表
func ListIpsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IpListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := product.NewListIpsLogic(r.Context(), svcCtx)
		resp, err := l.ListIps(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

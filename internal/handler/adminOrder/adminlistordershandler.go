// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package adminOrder

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"poppulse/internal/logic/adminOrder"
	"poppulse/internal/svc"
	"poppulse/internal/types"
)

// M28 后台订单列表
func AdminListOrdersHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AdminOrderListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := adminOrder.NewAdminListOrdersLogic(r.Context(), svcCtx)
		resp, err := l.AdminListOrders(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

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

// M29 后台订单详情
func AdminGetOrderDetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.OrderDetailReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := adminOrder.NewAdminGetOrderDetailLogic(r.Context(), svcCtx)
		resp, err := l.AdminGetOrderDetail(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

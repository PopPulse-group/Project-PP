// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package aftersale

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"poppulse/internal/logic/aftersale"
	"poppulse/internal/svc"
	"poppulse/internal/types"
)

// M33 本人售后详情
func GetAfterSaleDetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AfterSaleDetailReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := aftersale.NewGetAfterSaleDetailLogic(r.Context(), svcCtx)
		resp, err := l.GetAfterSaleDetail(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

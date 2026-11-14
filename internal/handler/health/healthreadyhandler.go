// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package health

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"poppulse/internal/logic/health"
	"poppulse/internal/svc"
)

// M02 服务就绪
func HealthReadyHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := health.NewHealthReadyLogic(r.Context(), svcCtx)
		resp, err := l.HealthReady()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

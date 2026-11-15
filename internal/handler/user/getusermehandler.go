// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package user

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"vagetable-gozero/internal/logic/user"
	"vagetable-gozero/internal/svc"
)

// 当前用户
func GetUserMeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := user.NewGetUserMeLogic(r.Context(), svcCtx)
		resp, err := l.GetUserMe()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

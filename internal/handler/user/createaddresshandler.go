// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package user

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"vagetable-gozero/internal/logic/user"
	"vagetable-gozero/internal/svc"
	"vagetable-gozero/internal/types"
)

// 新增地址
func CreateAddressHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateAddressReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewCreateAddressLogic(r.Context(), svcCtx)
		resp, err := l.CreateAddress(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

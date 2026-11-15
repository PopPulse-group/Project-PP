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

// 删除地址（仅本人）
func DeleteAddressHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddressIdReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewDeleteAddressLogic(r.Context(), svcCtx)
		err := l.DeleteAddress(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}

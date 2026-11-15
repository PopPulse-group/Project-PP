package middleware

import (
	"encoding/json"
	"net/http"

	"vagetable-gozero/internal/ctxdata"
	"vagetable-gozero/internal/xerr"
)

// OperatorMiddleware 校验当前登录用户是否具备运营（operator）角色。
// 注意：角色来自 JWT claims，不能由请求参数决定；业务层仍需校验数据归属。
type OperatorMiddleware struct {
}

func NewOperatorMiddleware() *OperatorMiddleware {
	return &OperatorMiddleware{}
}

func (m *OperatorMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		role, ok := ctxdata.RoleFromCtx(r.Context())
		if !ok || role != "operator" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusForbidden)
			_ = json.NewEncoder(w).Encode(map[string]any{
				"code":    xerr.ErrForbidden.Code,
				"message": xerr.ErrForbidden.Msg,
				"data":    nil,
			})
			return
		}
		next(w, r)
	}
}

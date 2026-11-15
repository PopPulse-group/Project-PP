// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package ai

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"vagetable-gozero/internal/logic/ai"
	"vagetable-gozero/internal/svc"
	"vagetable-gozero/internal/types"
)

// 确认推荐并加入购物车（幂等合并，不直接扣款）
func ConfirmMealPlanHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ConfirmMealPlanReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := ai.NewConfirmMealPlanLogic(r.Context(), svcCtx)
		resp, err := l.ConfirmMealPlan(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package ai

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"poppulse/internal/logic/ai"
	"poppulse/internal/svc"
	"poppulse/internal/types"
)

// 生成菜谱推荐（人数/预算/口味/忌口 → 结构化菜谱 + 在售 SKU 匹配）
func CreateMealPlanHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateMealPlanReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := ai.NewCreateMealPlanLogic(r.Context(), svcCtx)
		resp, err := l.CreateMealPlan(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

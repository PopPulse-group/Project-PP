// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package ai

import (
	"context"

	"vagetable-gozero/internal/svc"
	"vagetable-gozero/internal/types"
	"vagetable-gozero/internal/xerr"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateMealPlanLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 生成菜谱推荐（人数/预算/口味/忌口 → 结构化菜谱 + 在售 SKU 匹配）
func NewCreateMealPlanLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateMealPlanLogic {
	return &CreateMealPlanLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateMealPlanLogic) CreateMealPlan(req *types.CreateMealPlanReq) (resp *types.MealPlanResp, err error) {
	return nil, xerr.ErrNotImplemented
}

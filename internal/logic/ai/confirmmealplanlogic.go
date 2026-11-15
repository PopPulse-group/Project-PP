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

type ConfirmMealPlanLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 确认推荐并加入购物车（幂等合并，不直接扣款）
func NewConfirmMealPlanLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ConfirmMealPlanLogic {
	return &ConfirmMealPlanLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ConfirmMealPlanLogic) ConfirmMealPlan(req *types.ConfirmMealPlanReq) (resp *types.ConfirmMealPlanResp, err error) {
	return nil, xerr.ErrNotImplemented
}

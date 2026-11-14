// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package adminStock

import (
	"context"

	"poppulse/internal/svc"
	"poppulse/internal/types"
	"poppulse/internal/xerr"

	"github.com/zeromicro/go-zero/core/logx"
)

type AdjustStockLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// M20 调整可售库存
func NewAdjustStockLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdjustStockLogic {
	return &AdjustStockLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AdjustStockLogic) AdjustStock(req *types.AdjustStockReq) (resp *types.AdjustStockResp, err error) {
	return nil, xerr.ErrNotImplemented
}

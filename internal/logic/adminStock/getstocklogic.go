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

type GetStockLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// M19 库存余额
func NewGetStockLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetStockLogic {
	return &GetStockLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetStockLogic) GetStock(req *types.GetStockReq) (resp *types.StockResp, err error) {
	return nil, xerr.ErrNotImplemented
}

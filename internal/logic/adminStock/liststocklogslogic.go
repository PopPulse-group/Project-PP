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

type ListStockLogsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// M21 库存流水
func NewListStockLogsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListStockLogsLogic {
	return &ListStockLogsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListStockLogsLogic) ListStockLogs(req *types.StockLogListReq) (resp *types.StockLogListResp, err error) {
	return nil, xerr.ErrNotImplemented
}

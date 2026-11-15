// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package order

import (
	"context"

	"vagetable-gozero/internal/svc"
	"vagetable-gozero/internal/types"
	"vagetable-gozero/internal/xerr"

	"github.com/zeromicro/go-zero/core/logx"
)

type CancelOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 取消待支付订单（一次性释放库存）
func NewCancelOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CancelOrderLogic {
	return &CancelOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CancelOrderLogic) CancelOrder(req *types.OrderIdReq) (resp *types.CancelOrderResp, err error) {
	return nil, xerr.ErrNotImplemented
}

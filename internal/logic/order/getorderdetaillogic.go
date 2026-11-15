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

type GetOrderDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 本人订单详情
func NewGetOrderDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOrderDetailLogic {
	return &GetOrderDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetOrderDetailLogic) GetOrderDetail(req *types.OrderIdReq) (resp *types.OrderDetailResp, err error) {
	return nil, xerr.ErrNotImplemented
}

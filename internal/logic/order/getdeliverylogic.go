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

type GetDeliveryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 查询模拟配送进度
func NewGetDeliveryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDeliveryLogic {
	return &GetDeliveryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetDeliveryLogic) GetDelivery(req *types.OrderIdReq) (resp *types.DeliveryResp, err error) {
	return nil, xerr.ErrNotImplemented
}

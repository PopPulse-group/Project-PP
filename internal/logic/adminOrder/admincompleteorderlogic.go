// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package adminOrder

import (
	"context"

	"poppulse/internal/svc"
	"poppulse/internal/types"
	"poppulse/internal/xerr"

	"github.com/zeromicro/go-zero/core/logx"
)

type AdminCompleteOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// M30 运营确认订单完成
func NewAdminCompleteOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdminCompleteOrderLogic {
	return &AdminCompleteOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AdminCompleteOrderLogic) AdminCompleteOrder(req *types.AdminCompleteOrderReq) (resp *types.AdminCompleteOrderResp, err error) {
	return nil, xerr.ErrNotImplemented
}

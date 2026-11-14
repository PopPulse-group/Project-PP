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

type AdminListOrdersLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// M28 后台订单列表
func NewAdminListOrdersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdminListOrdersLogic {
	return &AdminListOrdersLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AdminListOrdersLogic) AdminListOrders(req *types.AdminOrderListReq) (resp *types.AdminOrderListResp, err error) {
	return nil, xerr.ErrNotImplemented
}

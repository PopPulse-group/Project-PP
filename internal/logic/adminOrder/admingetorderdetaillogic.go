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

type AdminGetOrderDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// M29 后台订单详情
func NewAdminGetOrderDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdminGetOrderDetailLogic {
	return &AdminGetOrderDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AdminGetOrderDetailLogic) AdminGetOrderDetail(req *types.OrderDetailReq) (resp *types.AdminOrderDetailResp, err error) {
	return nil, xerr.ErrNotImplemented
}

// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package adminAftersale

import (
	"context"

	"poppulse/internal/svc"
	"poppulse/internal/types"
	"poppulse/internal/xerr"

	"github.com/zeromicro/go-zero/core/logx"
)

type AdminGetRefundLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// M39 后台退款查询
func NewAdminGetRefundLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdminGetRefundLogic {
	return &AdminGetRefundLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AdminGetRefundLogic) AdminGetRefund(req *types.AdminGetRefundReq) (resp *types.RefundResp, err error) {
	return nil, xerr.ErrNotImplemented
}

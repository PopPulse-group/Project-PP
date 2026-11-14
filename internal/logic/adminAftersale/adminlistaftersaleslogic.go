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

type AdminListAfterSalesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// M35 后台售后列表
func NewAdminListAfterSalesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdminListAfterSalesLogic {
	return &AdminListAfterSalesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AdminListAfterSalesLogic) AdminListAfterSales(req *types.AdminAfterSaleListReq) (resp *types.AdminAfterSaleListResp, err error) {
	return nil, xerr.ErrNotImplemented
}

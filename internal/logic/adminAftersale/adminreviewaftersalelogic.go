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

type AdminReviewAfterSaleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// M37 审核售后
func NewAdminReviewAfterSaleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdminReviewAfterSaleLogic {
	return &AdminReviewAfterSaleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AdminReviewAfterSaleLogic) AdminReviewAfterSale(req *types.AdminReviewAfterSaleReq) (resp *types.AdminReviewAfterSaleResp, err error) {
	return nil, xerr.ErrNotImplemented
}

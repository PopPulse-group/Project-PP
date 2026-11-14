// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package aftersale

import (
	"context"

	"poppulse/internal/svc"
	"poppulse/internal/types"
	"poppulse/internal/xerr"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListAfterSalesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// M32 本人售后列表
func NewListAfterSalesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListAfterSalesLogic {
	return &ListAfterSalesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListAfterSalesLogic) ListAfterSales(req *types.AfterSaleListReq) (resp *types.AfterSaleListResp, err error) {
	return nil, xerr.ErrNotImplemented
}

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

type CreateAfterSaleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// M31 申请整单退款
func NewCreateAfterSaleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateAfterSaleLogic {
	return &CreateAfterSaleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateAfterSaleLogic) CreateAfterSale(req *types.CreateAfterSaleReq) (resp *types.CreateAfterSaleResp, err error) {
	return nil, xerr.ErrNotImplemented
}

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

type GetRefundLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// M34 本人退款结果
func NewGetRefundLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRefundLogic {
	return &GetRefundLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetRefundLogic) GetRefund(req *types.GetRefundReq) (resp *types.RefundResp, err error) {
	return nil, xerr.ErrNotImplemented
}

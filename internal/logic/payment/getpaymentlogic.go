// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package payment

import (
	"context"

	"poppulse/internal/svc"
	"poppulse/internal/types"
	"poppulse/internal/xerr"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPaymentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// M27 查询本人支付结果
func NewGetPaymentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPaymentLogic {
	return &GetPaymentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPaymentLogic) GetPayment(req *types.GetPaymentReq) (resp *types.PaymentResp, err error) {
	return nil, xerr.ErrNotImplemented
}

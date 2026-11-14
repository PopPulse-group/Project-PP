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

type AdminMockRefundConfirmLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// M38 模拟退款成功
func NewAdminMockRefundConfirmLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdminMockRefundConfirmLogic {
	return &AdminMockRefundConfirmLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AdminMockRefundConfirmLogic) AdminMockRefundConfirm(req *types.AdminMockRefundConfirmReq) (resp *types.AdminMockRefundConfirmResp, err error) {
	return nil, xerr.ErrNotImplemented
}

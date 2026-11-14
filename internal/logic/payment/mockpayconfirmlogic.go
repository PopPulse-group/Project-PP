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

type MockPayConfirmLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// M26 模拟支付成功
func NewMockPayConfirmLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MockPayConfirmLogic {
	return &MockPayConfirmLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MockPayConfirmLogic) MockPayConfirm(req *types.MockPayConfirmReq) (resp *types.MockPayConfirmResp, err error) {
	return nil, xerr.ErrNotImplemented
}

// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package order

import (
	"context"

	"poppulse/internal/svc"
	"poppulse/internal/types"
	"poppulse/internal/xerr"

	"github.com/zeromicro/go-zero/core/logx"
)

type MockPayLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 模拟支付（仅 Mock.Enabled 时可用）
func NewMockPayLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MockPayLogic {
	return &MockPayLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MockPayLogic) MockPay(req *types.OrderIdReq) (resp *types.MockPayResp, err error) {
	return nil, xerr.ErrNotImplemented
}

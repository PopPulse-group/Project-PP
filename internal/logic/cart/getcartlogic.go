// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package cart

import (
	"context"

	"poppulse/internal/svc"
	"poppulse/internal/types"
	"poppulse/internal/xerr"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCartLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 查看购物车（实时价格与库存）
func NewGetCartLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCartLogic {
	return &GetCartLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCartLogic) GetCart() (resp *types.CartResp, err error) {
	return nil, xerr.ErrNotImplemented
}

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

type AddCartItemLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 加购（同商品合并数量）
func NewAddCartItemLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddCartItemLogic {
	return &AddCartItemLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddCartItemLogic) AddCartItem(req *types.AddCartItemReq) (resp *types.CartResp, err error) {
	return nil, xerr.ErrNotImplemented
}

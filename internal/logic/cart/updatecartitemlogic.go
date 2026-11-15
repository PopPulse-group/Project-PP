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

type UpdateCartItemLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 改数量（0 视为删除）
func NewUpdateCartItemLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCartItemLogic {
	return &UpdateCartItemLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateCartItemLogic) UpdateCartItem(req *types.UpdateCartItemReq) (resp *types.CartResp, err error) {
	return nil, xerr.ErrNotImplemented
}

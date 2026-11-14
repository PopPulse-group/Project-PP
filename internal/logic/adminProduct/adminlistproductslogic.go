// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package adminProduct

import (
	"context"

	"poppulse/internal/svc"
	"poppulse/internal/types"
	"poppulse/internal/xerr"

	"github.com/zeromicro/go-zero/core/logx"
)

type AdminListProductsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// M13 后台商品列表
func NewAdminListProductsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdminListProductsLogic {
	return &AdminListProductsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AdminListProductsLogic) AdminListProducts(req *types.AdminProductListReq) (resp *types.AdminProductListResp, err error) {
	return nil, xerr.ErrNotImplemented
}

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

type AdminGetProductDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// M14 后台商品详情
func NewAdminGetProductDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdminGetProductDetailLogic {
	return &AdminGetProductDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AdminGetProductDetailLogic) AdminGetProductDetail(req *types.AdminProductDetailReq) (resp *types.AdminProductDetailResp, err error) {
	return nil, xerr.ErrNotImplemented
}

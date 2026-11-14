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

type UpdateSkuLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// M18 修改 SKU
func NewUpdateSkuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSkuLogic {
	return &UpdateSkuLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateSkuLogic) UpdateSku(req *types.UpdateSkuReq) (resp *types.CreateEntityResp, err error) {
	return nil, xerr.ErrNotImplemented
}

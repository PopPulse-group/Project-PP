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

type AdminGetAfterSaleDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// M36 后台售后详情
func NewAdminGetAfterSaleDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdminGetAfterSaleDetailLogic {
	return &AdminGetAfterSaleDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AdminGetAfterSaleDetailLogic) AdminGetAfterSaleDetail(req *types.AfterSaleDetailReq) (resp *types.AfterSaleDetailResp, err error) {
	return nil, xerr.ErrNotImplemented
}

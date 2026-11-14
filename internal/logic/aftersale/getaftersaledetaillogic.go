// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package aftersale

import (
	"context"

	"poppulse/internal/svc"
	"poppulse/internal/types"
	"poppulse/internal/xerr"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAfterSaleDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// M33 本人售后详情
func NewGetAfterSaleDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAfterSaleDetailLogic {
	return &GetAfterSaleDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAfterSaleDetailLogic) GetAfterSaleDetail(req *types.AfterSaleDetailReq) (resp *types.AfterSaleDetailResp, err error) {
	return nil, xerr.ErrNotImplemented
}

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

type CreateSeriesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// M11 创建系列
func NewCreateSeriesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateSeriesLogic {
	return &CreateSeriesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateSeriesLogic) CreateSeries(req *types.CreateSeriesReq) (resp *types.CreateEntityResp, err error) {
	return nil, xerr.ErrNotImplemented
}

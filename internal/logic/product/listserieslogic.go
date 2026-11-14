// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package product

import (
	"context"

	"poppulse/internal/svc"
	"poppulse/internal/types"
	"poppulse/internal/xerr"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListSeriesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// M07 系列列表
func NewListSeriesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListSeriesLogic {
	return &ListSeriesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListSeriesLogic) ListSeries(req *types.SeriesListReq) (resp *types.SeriesListResp, err error) {
	return nil, xerr.ErrNotImplemented
}

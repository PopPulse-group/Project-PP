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

type ListIpsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// M06 IP 列表
func NewListIpsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListIpsLogic {
	return &ListIpsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListIpsLogic) ListIps(req *types.IpListReq) (resp *types.IpListResp, err error) {
	return nil, xerr.ErrNotImplemented
}

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

type CreateIpLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// M10 创建 IP
func NewCreateIpLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateIpLogic {
	return &CreateIpLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateIpLogic) CreateIp(req *types.CreateIpReq) (resp *types.CreateEntityResp, err error) {
	return nil, xerr.ErrNotImplemented
}

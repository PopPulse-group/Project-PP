// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package health

import (
	"context"
	"time"

	"vagetable-gozero/internal/svc"
	"vagetable-gozero/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HealthLiveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 进程存活
func NewHealthLiveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HealthLiveLogic {
	return &HealthLiveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HealthLiveLogic) HealthLive() (resp *types.HealthResp, err error) {
	return &types.HealthResp{Status: "ok", Time: time.Now().Format(time.RFC3339)}, nil
}

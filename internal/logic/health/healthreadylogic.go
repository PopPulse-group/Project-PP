package health

import (
	"context"
	"time"

	"poppulse/internal/svc"
	"poppulse/internal/types"
	"poppulse/internal/xerr"

	"github.com/zeromicro/go-zero/core/logx"
)

type HealthReadyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// M02 服务就绪：检查 MySQL 等必要依赖，不可用返回 503/90002
func NewHealthReadyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HealthReadyLogic {
	return &HealthReadyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HealthReadyLogic) HealthReady() (resp *types.HealthResp, err error) {
	var one int
	if err := l.svcCtx.DB.QueryRowCtx(l.ctx, &one, "SELECT 1"); err != nil {
		l.Errorf("health ready check failed: %v", err)
		return nil, xerr.ErrServiceDegrade
	}
	return &types.HealthResp{Status: "ok", Time: time.Now().Format(time.RFC3339)}, nil
}

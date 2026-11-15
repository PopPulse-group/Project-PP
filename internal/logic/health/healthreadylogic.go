// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package health

import (
	"context"
	"time"

	"vagetable-gozero/internal/svc"
	"vagetable-gozero/internal/types"
	"vagetable-gozero/internal/xerr"

	"github.com/zeromicro/go-zero/core/logx"
)

type HealthReadyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 服务就绪（MySQL/Redis）
func NewHealthReadyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HealthReadyLogic {
	return &HealthReadyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HealthReadyLogic) HealthReady() (resp *types.HealthResp, err error) {
	// MySQL 必须可用
	var one int
	if err := l.svcCtx.DB.QueryRowCtx(l.ctx, &one, "SELECT 1"); err != nil {
		l.Errorf("health ready mysql check failed: %v", err)
		return nil, xerr.ErrServiceDegrade
	}
	// Redis 启用时必须可用（缓存加速用，不承载库存事实）
	if l.svcCtx.Redis != nil {
		if !l.svcCtx.Redis.Ping() {
			l.Errorf("health ready redis check failed: ping=false")
			return nil, xerr.ErrServiceDegrade
		}
	}
	return &types.HealthResp{Status: "ok", Time: time.Now().Format(time.RFC3339)}, nil
}

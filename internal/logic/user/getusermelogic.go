package user

import (
	"context"
	"strconv"

	"poppulse/internal/ctxdata"
	"poppulse/internal/model"
	"poppulse/internal/svc"
	"poppulse/internal/types"
	"poppulse/internal/xerr"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type GetUserMeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// M04 当前用户
func NewGetUserMeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserMeLogic {
	return &GetUserMeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserMeLogic) GetUserMe() (resp *types.UserInfo, err error) {
	userId, ok := ctxdata.UserIdFromCtx(l.ctx)
	if !ok {
		return nil, xerr.ErrUnauthorized
	}
	id, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		return nil, xerr.ErrUnauthorized
	}

	u, err := l.svcCtx.UserModel.FindById(l.ctx, id)
	if err != nil {
		if err == sqlx.ErrNotFound {
			return nil, xerr.ErrNotFound
		}
		l.Errorf("find user failed, id=%s: %v", userId, err)
		return nil, xerr.ErrInternal
	}

	info := model.ToUserInfo(u)
	return &info, nil
}

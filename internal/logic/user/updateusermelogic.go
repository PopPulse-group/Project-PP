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

type UpdateUserMeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// M05 修改个人资料：PATCH 语义，未传字段保留原值
func NewUpdateUserMeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserMeLogic {
	return &UpdateUserMeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateUserMeLogic) UpdateUserMe(req *types.UpdateUserMeReq) (resp *types.UserInfo, err error) {
	userId, ok := ctxdata.UserIdFromCtx(l.ctx)
	if !ok {
		return nil, xerr.ErrUnauthorized
	}
	id, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		return nil, xerr.ErrUnauthorized
	}

	// 参数校验：昵称长度限制
	if req.Nickname != nil && (len(*req.Nickname) == 0 || len(*req.Nickname) > 64) {
		return nil, xerr.ErrInvalidParam
	}

	if err := l.svcCtx.UserModel.UpdateProfile(l.ctx, id, req.Nickname, req.Avatar, req.Bio); err != nil {
		l.Errorf("update profile failed, id=%s: %v", userId, err)
		return nil, xerr.ErrInternal
	}

	u, err := l.svcCtx.UserModel.FindById(l.ctx, id)
	if err != nil {
		if err == sqlx.ErrNotFound {
			return nil, xerr.ErrNotFound
		}
		return nil, xerr.ErrInternal
	}

	info := model.ToUserInfo(u)
	return &info, nil
}

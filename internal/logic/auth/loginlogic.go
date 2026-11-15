// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package auth

import (
	"context"

	"poppulse/internal/ctxdata"
	"poppulse/internal/model"
	"poppulse/internal/svc"
	"poppulse/internal/types"
	"poppulse/internal/xerr"

	"github.com/zeromicro/go-zero/core/logx"
	"golang.org/x/crypto/bcrypt"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 登录
func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	// 参数校验
	if req.Username == "" || req.Password == "" {
		return nil, xerr.ErrInvalidParam
	}

	// 查询用户：失败不区分账号不存在或密码错误
	u, err := l.svcCtx.UserModel.FindByUsername(l.ctx, req.Username)
	if err != nil {
		l.Logger.Infof("login user not found, username=%s", req.Username)
		return nil, xerr.ErrLoginFailed
	}

	// 密码哈希校验
	if err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(req.Password)); err != nil {
		return nil, xerr.ErrLoginFailed
	}

	// 账号状态校验
	if u.Status != "active" {
		return nil, xerr.ErrForbidden
	}

	// 签发 JWT，claims 携带 userId 与 role
	token, expiresAt, err := ctxdata.BuildToken(l.svcCtx.Config.Auth.AccessSecret, l.svcCtx.Config.Auth.AccessExpire, model.FormatId(u.Id), "user")
	if err != nil {
		l.Errorf("build token failed: %v", err)
		return nil, xerr.ErrInternal
	}

	return &types.LoginResp{
		AccessToken: token,
		ExpiresAt:   expiresAt,
		User:        model.ToUserInfo(u),
	}, nil
}

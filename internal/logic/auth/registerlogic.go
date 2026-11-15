// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package auth

import (
	"context"
	"strings"

	"poppulse/internal/ctxdata"
	"poppulse/internal/model"
	"poppulse/internal/svc"
	"poppulse/internal/types"
	"poppulse/internal/xerr"

	"github.com/zeromicro/go-zero/core/logx"
	"golang.org/x/crypto/bcrypt"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 注册（成功即登录，返回 token）
func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.LoginResp, err error) {
	// 参数校验
	username := strings.TrimSpace(req.Username)
	if len(username) < 4 || len(username) > 32 {
		return nil, xerr.NewBiz(40011, "用户名长度需为 4-32 字符")
	}
	if len(req.Password) < 8 || len(req.Password) > 64 {
		return nil, xerr.NewBiz(40012, "密码长度需为 8-64 字符")
	}
	nickname := strings.TrimSpace(req.Nickname)
	if nickname == "" {
		nickname = username
	}

	// 用户名重复检查
	if _, err := l.svcCtx.UserModel.FindByUsername(l.ctx, username); err == nil {
		return nil, xerr.NewBiz(40904, "用户名已存在")
	}

	// 创建用户（bcrypt 哈希存储）
	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		l.Errorf("hash password failed: %v", err)
		return nil, xerr.ErrInternal
	}
	u, err := l.svcCtx.UserModel.Create(l.ctx, username, string(hash), nickname)
	if err != nil {
		// 唯一约束兜底（并发注册同名）
		l.Errorf("create user failed: %v", err)
		return nil, xerr.NewBiz(40904, "用户名已存在")
	}

	// 注册成功即登录：签发 JWT
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

// Package ctxdata 定义 JWT 声明在 context 中的键以及 token 生成工具。
package ctxdata

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// go-zero 将 JWT claims 以 claim 名为键注入请求 context
const (
	CtxUserIdKey = "userId"
	CtxRoleKey   = "role"
)

// UserIdFromCtx 从请求 context 中取当前登录用户 ID（int64 字符串形式）
func UserIdFromCtx(ctx context.Context) (string, bool) {
	v, ok := ctx.Value(CtxUserIdKey).(string)
	return v, ok && v != ""
}

// RoleFromCtx 从请求 context 中取当前登录用户角色
func RoleFromCtx(ctx context.Context) (string, bool) {
	v, ok := ctx.Value(CtxRoleKey).(string)
	return v, ok && v != ""
}

// BuildToken 生成 HS256 JWT，claims 含 userId 与 role
func BuildToken(secret string, expireSeconds int64, userId, role string) (token string, expiresAt int64, err error) {
	now := time.Now().Unix()
	expiresAt = now + expireSeconds
	claims := make(jwt.MapClaims)
	claims["iat"] = now
	claims["exp"] = expiresAt
	claims["nbf"] = now
	claims[CtxUserIdKey] = userId
	claims[CtxRoleKey] = role
	t := jwt.New(jwt.SigningMethodHS256)
	t.Claims = claims
	signed, err := t.SignedString([]byte(secret))
	return signed, expiresAt, err
}

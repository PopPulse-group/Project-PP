// Package model 数据访问层（goctl model 生成 + 手写补充）。
package model

import (
	"context"
	"strconv"
	"time"

	"poppulse/internal/types"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type User struct {
	Id           int64     `db:"id"`
	Username     string    `db:"username"`
	PasswordHash string    `db:"password_hash"`
	Nickname     string    `db:"nickname"`
	Avatar       string    `db:"avatar"`
	Bio          string    `db:"bio"`
	Role         string    `db:"role"`   // user | operator
	Status       string    `db:"status"` // active | disabled
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
}

type UserModel struct {
	conn sqlx.SqlConn
}

func NewUserModel(conn sqlx.SqlConn) *UserModel {
	return &UserModel{conn: conn}
}

func (m *UserModel) FindByUsername(ctx context.Context, username string) (*User, error) {
	var u User
	err := m.conn.QueryRowCtx(ctx, &u,
		"SELECT id, username, password_hash, nickname, avatar, bio, role, status, created_at, updated_at FROM users WHERE username = ? LIMIT 1", username)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (m *UserModel) FindById(ctx context.Context, id int64) (*User, error) {
	var u User
	err := m.conn.QueryRowCtx(ctx, &u,
		"SELECT id, username, password_hash, nickname, avatar, bio, role, status, created_at, updated_at FROM users WHERE id = ? LIMIT 1", id)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

// FormatId 将数据库自增 ID 转为字符串（JSON 传输统一字符串，避免大整数精度问题）
func FormatId(id int64) string {
	return strconv.FormatInt(id, 10)
}

// ToUserInfo 转换为接口输出结构
func ToUserInfo(u *User) types.UserInfo {
	return types.UserInfo{
		Id:        FormatId(u.Id),
		Username:  u.Username,
		Nickname:  u.Nickname,
		Role:      u.Role,
		CreatedAt: u.CreatedAt.Format(time.RFC3339),
	}
}

// UpdateProfile 按传入的非空指针更新用户资料
func (m *UserModel) UpdateProfile(ctx context.Context, id int64, nickname, avatar, bio *string) error {
	if nickname == nil && avatar == nil && bio == nil {
		return nil
	}
	sql := "UPDATE users SET "
	var args []any
	if nickname != nil {
		sql += "nickname = ?, "
		args = append(args, *nickname)
	}
	if avatar != nil {
		sql += "avatar = ?, "
		args = append(args, *avatar)
	}
	if bio != nil {
		sql += "bio = ?, "
		args = append(args, *bio)
	}
	sql = sql[:len(sql)-2] + " WHERE id = ?"
	args = append(args, id)
	_, err := m.conn.ExecCtx(ctx, sql, args...)
	return err
}

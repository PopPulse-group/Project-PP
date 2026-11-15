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
		CreatedAt: u.CreatedAt.Format(time.RFC3339),
	}
}

func (m *UserModel) FindByUsername(ctx context.Context, username string) (*User, error) {
	var u User
	err := m.conn.QueryRowCtx(ctx, &u,
		"SELECT id, username, password_hash, nickname, status, created_at, updated_at FROM users WHERE username = ? LIMIT 1", username)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (m *UserModel) FindById(ctx context.Context, id int64) (*User, error) {
	var u User
	err := m.conn.QueryRowCtx(ctx, &u,
		"SELECT id, username, password_hash, nickname, status, created_at, updated_at FROM users WHERE id = ? LIMIT 1", id)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

// Create 注册新用户；username 唯一约束冲突由调用方转为业务错误
func (m *UserModel) Create(ctx context.Context, username, passwordHash, nickname string) (*User, error) {
	res, err := m.conn.ExecCtx(ctx,
		"INSERT INTO users (username, password_hash, nickname, status) VALUES (?, ?, ?, 'active')",
		username, passwordHash, nickname)
	if err != nil {
		return nil, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	return m.FindById(ctx, id)
}

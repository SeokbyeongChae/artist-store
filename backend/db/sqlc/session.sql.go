// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: session.sql

package db

import (
	"context"
	"database/sql"
	"time"
)

const createSession = `-- name: CreateSession :execresult
INSERT INTO sessions (
  uuid,
  account_id,
  refresh_token,
  user_agent,
  client_ip,
  is_blocked,
  expire_at
) VALUES (
  ?, ?, ?, ?, ?, ?, ?
)
`

type CreateSessionParams struct {
	Uuid         string    `json:"uuid"`
	AccountID    int64     `json:"account_id"`
	RefreshToken string    `json:"refresh_token"`
	UserAgent    string    `json:"user_agent"`
	ClientIp     string    `json:"client_ip"`
	IsBlocked    bool      `json:"is_blocked"`
	ExpireAt     time.Time `json:"expire_at"`
}

func (q *Queries) CreateSession(ctx context.Context, arg CreateSessionParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createSession,
		arg.Uuid,
		arg.AccountID,
		arg.RefreshToken,
		arg.UserAgent,
		arg.ClientIp,
		arg.IsBlocked,
		arg.ExpireAt,
	)
}

const getSession = `-- name: GetSession :one
SELECT uuid, account_id, refresh_token, user_agent, client_ip, is_blocked, expire_at, created_at FROM sessions
WHERE uuid = ? LIMIT 1
`

func (q *Queries) GetSession(ctx context.Context, uuid string) (Session, error) {
	row := q.db.QueryRowContext(ctx, getSession, uuid)
	var i Session
	err := row.Scan(
		&i.Uuid,
		&i.AccountID,
		&i.RefreshToken,
		&i.UserAgent,
		&i.ClientIp,
		&i.IsBlocked,
		&i.ExpireAt,
		&i.CreatedAt,
	)
	return i, err
}

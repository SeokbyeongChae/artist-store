// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package db

import (
	"context"
	"database/sql"
)

type Querier interface {
	CreateAccount(ctx context.Context, arg CreateAccountParams) (sql.Result, error)
	GetAccount(ctx context.Context, id int64) (Account, error)
	GetAccountByEmail(ctx context.Context, email string) (Account, error)
	ListAccounts(ctx context.Context, arg ListAccountsParams) ([]Account, error)
	UpdateAccount(ctx context.Context, arg UpdateAccountParams) error
}

var _ Querier = (*Queries)(nil)

// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: accounts.sql

package db

import (
	"context"
)

const AddAccountBalancer = `-- name: AddAccountBalancer :one
UPDATE accounts
SET balance = balance + $1
WHERE id = $2
RETURNING id, owner, balance, currency, created_at
`

type AddAccountBalancerParams struct {
	Amount int64 `json:"amount"`
	ID     int64 `json:"id"`
}

// AddAccountBalancer
//
//	UPDATE accounts
//	SET balance = balance + $1
//	WHERE id = $2
//	RETURNING id, owner, balance, currency, created_at
func (q *Queries) AddAccountBalancer(ctx context.Context, arg AddAccountBalancerParams) (Accounts, error) {
	row := q.db.QueryRow(ctx, AddAccountBalancer, arg.Amount, arg.ID)
	var i Accounts
	err := row.Scan(
		&i.ID,
		&i.Owner,
		&i.Balance,
		&i.Currency,
		&i.CreatedAt,
	)
	return i, err
}

const CreateAccount = `-- name: CreateAccount :one
INSERT INTO accounts(owner, balance, currency)
VALUES ($1, $2, $3)
RETURNING id, owner, balance, currency, created_at
`

type CreateAccountParams struct {
	Owner    string `json:"owner"`
	Balance  int64  `json:"balance"`
	Currency string `json:"currency"`
}

// CreateAccount
//
//	INSERT INTO accounts(owner, balance, currency)
//	VALUES ($1, $2, $3)
//	RETURNING id, owner, balance, currency, created_at
func (q *Queries) CreateAccount(ctx context.Context, arg CreateAccountParams) (Accounts, error) {
	row := q.db.QueryRow(ctx, CreateAccount, arg.Owner, arg.Balance, arg.Currency)
	var i Accounts
	err := row.Scan(
		&i.ID,
		&i.Owner,
		&i.Balance,
		&i.Currency,
		&i.CreatedAt,
	)
	return i, err
}

const DeleteAccount = `-- name: DeleteAccount :exec
DELETE
FROM accounts
WHERE id = $1
`

// DeleteAccount
//
//	DELETE
//	FROM accounts
//	WHERE id = $1
func (q *Queries) DeleteAccount(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, DeleteAccount, id)
	return err
}

const GetAccount = `-- name: GetAccount :one
SELECT id, owner, balance, currency, created_at
FROM accounts
WHERE id = $1
ORDER BY id
`

// GetAccount
//
//	SELECT id, owner, balance, currency, created_at
//	FROM accounts
//	WHERE id = $1
//	ORDER BY id
func (q *Queries) GetAccount(ctx context.Context, id int64) (Accounts, error) {
	row := q.db.QueryRow(ctx, GetAccount, id)
	var i Accounts
	err := row.Scan(
		&i.ID,
		&i.Owner,
		&i.Balance,
		&i.Currency,
		&i.CreatedAt,
	)
	return i, err
}

const GetAccountForUpdate = `-- name: GetAccountForUpdate :one
SELECT id, owner, balance, currency, created_at
FROM accounts
WHERE id = $1
    FOR NO KEY UPDATE
`

// GetAccountForUpdate
//
//	SELECT id, owner, balance, currency, created_at
//	FROM accounts
//	WHERE id = $1
//	    FOR NO KEY UPDATE
func (q *Queries) GetAccountForUpdate(ctx context.Context, id int64) (Accounts, error) {
	row := q.db.QueryRow(ctx, GetAccountForUpdate, id)
	var i Accounts
	err := row.Scan(
		&i.ID,
		&i.Owner,
		&i.Balance,
		&i.Currency,
		&i.CreatedAt,
	)
	return i, err
}

const ListAccounts = `-- name: ListAccounts :many
SELECT id, owner, balance, currency, created_at
FROM accounts
ORDER BY id
LIMIT $1 OFFSET $2
`

type ListAccountsParams struct {
	Limit  int64 `json:"limit"`
	Offset int64 `json:"offset"`
}

// ListAccounts
//
//	SELECT id, owner, balance, currency, created_at
//	FROM accounts
//	ORDER BY id
//	LIMIT $1 OFFSET $2
func (q *Queries) ListAccounts(ctx context.Context, arg ListAccountsParams) ([]Accounts, error) {
	rows, err := q.db.Query(ctx, ListAccounts, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Accounts{}
	for rows.Next() {
		var i Accounts
		if err := rows.Scan(
			&i.ID,
			&i.Owner,
			&i.Balance,
			&i.Currency,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const UpdateAccount = `-- name: UpdateAccount :one
UPDATE accounts
SET balance = $2
WHERE id = $1
RETURNING id, owner, balance, currency, created_at
`

type UpdateAccountParams struct {
	ID      int64 `json:"id"`
	Balance int64 `json:"balance"`
}

// UpdateAccount
//
//	UPDATE accounts
//	SET balance = $2
//	WHERE id = $1
//	RETURNING id, owner, balance, currency, created_at
func (q *Queries) UpdateAccount(ctx context.Context, arg UpdateAccountParams) (Accounts, error) {
	row := q.db.QueryRow(ctx, UpdateAccount, arg.ID, arg.Balance)
	var i Accounts
	err := row.Scan(
		&i.ID,
		&i.Owner,
		&i.Balance,
		&i.Currency,
		&i.CreatedAt,
	)
	return i, err
}

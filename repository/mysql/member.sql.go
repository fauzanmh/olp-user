package mysql

import (
	"context"

	"github.com/fauzanmh/olp-user/entity"
)

const checkEmail = `-- name: CheckEmail :one
SELECT EXISTS(SELECT 1 FROM members WHERE email = ? LIMIT 1) AS exis
`

func (q *Queries) CheckEmail(ctx context.Context, email string) (bool, error) {
	row := q.queryRow(ctx, q.checkEmailStmt, checkEmail, email)
	var exist bool
	err := row.Scan(&exist)
	return exist, err
}

const register = `-- name: Register :exec
INSERT INTO members (name, email, address, created_at, updated_at) 
VALUES (?, ?, ?, ?, ?)
`

func (q *Queries) Register(ctx context.Context, arg *entity.RegisterParams) error {
	_, err := q.exec(ctx, q.registerStmt, register,
		arg.Name,
		arg.Email,
		arg.Address,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	return err
}

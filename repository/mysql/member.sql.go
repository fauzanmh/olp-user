package mysql

import (
	"context"
	"database/sql"

	"github.com/fauzanmh/olp-user/constant"
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

func (q *Queries) Register(ctx context.Context, arg *entity.RegisterParams) (int64, error) {
	var id int64
	res, err := q.exec(ctx, q.registerStmt, register,
		arg.Name,
		arg.Email,
		arg.Address,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	if err != nil {
		return id, err
	}

	id, err = res.LastInsertId()
	if err != nil {
		return id, err
	}

	return id, err
}

const checkMember = `-- name: CheckMember :one
SELECT name, email, address, deleted_at FROM members WHERE id = ?
`

func (q *Queries) CheckMember(ctx context.Context, id int64) (entity.CheckMemberRow, error) {
	row := q.queryRow(ctx, q.checkMemberStmt, checkMember, id)
	var i entity.CheckMemberRow
	err := row.Scan(
		&i.Name,
		&i.Email,
		&i.Address,
		&i.DeletedAt,
	)
	if err == sql.ErrNoRows {
		err = constant.ErrorMysqlDataNotFound
	}
	return i, err
}

const deleteMember = `-- name: DeleteMember :exec
UPDATE members SET deleted_at = ?
WHERE id = ?
`

func (q *Queries) DeleteMember(ctx context.Context, arg *entity.DeleteMemberParams) error {
	_, err := q.exec(ctx, q.deleteMemberStmt, deleteMember, arg.DeletedAt, arg.ID)
	return err
}

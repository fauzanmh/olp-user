package mysql

import (
	"context"
)

const getTotalCourse = `-- name: GetTotalCourse :one
SELECT count(id) FROM courses 
WHERE deleted_at IS NULL
`

func (q *Queries) GetTotalCourse(ctx context.Context) (int64, error) {
	row := q.queryRow(ctx, q.getTotalCourseStmt, getTotalCourse)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const getTotalCourseIsFree = `-- name: GetTotalCourseIsFree :one
SELECT count(id) FROM courses 
WHERE price = 0 AND deleted_at IS NULL
`

func (q *Queries) GetTotalCourseIsFree(ctx context.Context) (int64, error) {
	row := q.queryRow(ctx, q.getTotalCourseIsFreeStmt, getTotalCourseIsFree)
	var count int64
	err := row.Scan(&count)
	return count, err
}

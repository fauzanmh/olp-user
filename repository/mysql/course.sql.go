package mysql

import (
	"context"
	"database/sql"

	"github.com/fauzanmh/olp-user/constant"
	"github.com/fauzanmh/olp-user/entity"
)

const createCourse = `-- name: CreateCourse :exec
INSERT INTO courses (course_category_id, name, description, price, created_at, updated_at)
VALUES (?, ?, ?, ?, ?, ?)
`

func (q *Queries) CreateCourse(ctx context.Context, arg *entity.CreateCourseParams) error {
	_, err := q.exec(ctx, q.createCourseStmt, createCourse,
		arg.CourseCategoryID,
		arg.Name,
		arg.Description,
		arg.Price,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	return err
}

const deleteCourse = `-- name: DeleteCourse :exec
UPDATE courses SET deleted_at = ?
WHERE id = ?
`

func (q *Queries) DeleteCourse(ctx context.Context, arg *entity.DeleteCourseParams) error {
	_, err := q.exec(ctx, q.deleteCourseStmt, deleteCourse, arg.DeletedAt, arg.ID)
	return err
}

const getAllCourses = `-- name: GetAllCourses :many
SELECT id, course_category_id, name, description, price
FROM courses
`

func (q *Queries) GetAllCourses(ctx context.Context) ([]entity.GetAllCoursesRow, error) {
	rows, err := q.query(ctx, q.getAllCoursesStmt, getAllCourses)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []entity.GetAllCoursesRow{}
	for rows.Next() {
		var i entity.GetAllCoursesRow
		if err := rows.Scan(
			&i.ID,
			&i.CourseCategoryID,
			&i.Name,
			&i.Description,
			&i.Price,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getOneCourse = `-- name: GetOneCourse :one
SELECT id, course_category_id, name, description, price FROM courses
WHERE id = ? AND deleted_at IS NULL
`

func (q *Queries) GetOneCourse(ctx context.Context, id int64) (entity.GetOneCourseRow, error) {
	row := q.queryRow(ctx, q.getOneCourseStmt, getOneCourse, id)
	var i entity.GetOneCourseRow
	err := row.Scan(
		&i.ID,
		&i.CourseCategoryID,
		&i.Name,
		&i.Description,
		&i.Price,
	)
	if err == sql.ErrNoRows {
		err = constant.ErrorMysqlDataNotFound
	}
	return i, err
}

const updateCourse = `-- name: UpdateCourse :exec
UPDATE courses SET course_category_id = ?, name = ?, description = ?, price = ?, updated_at = ?
WHERE id = ?
`

func (q *Queries) UpdateCourse(ctx context.Context, arg *entity.UpdateCourseParams) error {
	_, err := q.exec(ctx, q.updateCourseStmt, updateCourse,
		arg.CourseCategoryID,
		arg.Name,
		arg.Description,
		arg.Price,
		arg.UpdatedAt,
		arg.ID,
	)
	return err
}

package mysql

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/fauzanmh/olp-user/constant"
	"github.com/fauzanmh/olp-user/entity"
	"go.uber.org/zap"
)

const getCourseDetail = `-- name: GetCourseDetail :one
SELECT c.id, course_category_id, c.name, description, price, course_category_name
FROM courses c
INNER JOIN (
    SELECT id as id_course_category, name as course_category_name FROM course_categories
) cc ON c.course_category_id = cc.id_course_category 
WHERE c.id = ? AND deleted_at IS NULL
`

func (q *Queries) GetCourseDetail(ctx context.Context, id int64) (entity.GetCourseDetailRow, error) {
	row := q.queryRow(ctx, q.getCourseDetailStmt, getCourseDetail, id)
	var i entity.GetCourseDetailRow
	err := row.Scan(
		&i.ID,
		&i.CourseCategoryID,
		&i.Name,
		&i.Description,
		&i.Price,
		&i.CourseCategoryName,
	)
	if err == sql.ErrNoRows {
		err = constant.ErrorMysqlDataNotFound
	}
	return i, err
}

const getCourses = `-- name: GetCourses :many
SELECT c.id, course_category_id, c.name, description, price, course_category_name
FROM courses c
INNER JOIN (
    SELECT id as id_course_category, name as course_category_name FROM course_categories
) cc ON c.course_category_id = cc.id_course_category
WHERE deleted_at IS NULL AND name LIKE "%s"
`

func (q *Queries) GetCourses(ctx context.Context, search string) ([]entity.GetCoursesRow, error) {
	query := fmt.Sprintf(getCourses, search)
	zap.S().Error(query)
	rows, err := q.query(ctx, q.getCoursesStmt, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []entity.GetCoursesRow{}
	for rows.Next() {
		var i entity.GetCoursesRow
		if err := rows.Scan(
			&i.ID,
			&i.CourseCategoryID,
			&i.Name,
			&i.Description,
			&i.Price,
			&i.CourseCategoryName,
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

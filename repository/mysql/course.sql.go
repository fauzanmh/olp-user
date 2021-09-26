package mysql

import (
	"context"

	"github.com/fauzanmh/olp-user/entity"
)

const getCourses = `-- name: GetCourses :many
SELECT c.id, course_category_id, c.name, description, price, course_category_name
FROM courses c
INNER JOIN (
    SELECT id as id_course_category, name as course_category_name FROM course_categories
) cc ON c.course_category_id = cc.id_course_category
`

func (q *Queries) GetCourses(ctx context.Context) ([]entity.GetCoursesRow, error) {
	rows, err := q.query(ctx, q.getCoursesStmt, getCourses)
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

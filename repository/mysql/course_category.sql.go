package mysql

import (
	"context"

	"github.com/fauzanmh/olp-user/entity"
)

const getAllCourseCategory = `-- name: GetAllCourseCategory :many
SELECT id, name, total_used FROM course_categories
WHERE deleted_at IS NULL
`

func (q *Queries) GetAllCourseCategory(ctx context.Context) ([]entity.GetAllCourseCategoryRow, error) {
	rows, err := q.query(ctx, q.getAllCourseCategoryStmt, getAllCourseCategory)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []entity.GetAllCourseCategoryRow{}
	for rows.Next() {
		var i entity.GetAllCourseCategoryRow
		if err := rows.Scan(&i.ID, &i.Name, &i.TotalUsed); err != nil {
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

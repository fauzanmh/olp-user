package mysql

import (
	"context"
	"database/sql"

	"github.com/fauzanmh/olp-user/entity"
)

type Repository interface {
	// Courses
	GetCourseDetail(ctx context.Context, id int64) (entity.GetCourseDetailRow, error)
	GetCourses(ctx context.Context, search, sort string) ([]entity.GetCoursesRow, error)

	// Course Category
	GetAllCourseCategory(ctx context.Context) ([]entity.GetAllCourseCategoryRow, error)
	GetPopularCourseCategory(ctx context.Context) (entity.GetPopularCourseCategoryRow, error)

	//Tx
	BeginTx(ctx context.Context) (*sql.Tx, error)
	WithTx(tx *sql.Tx) *Queries
	RollbackTx(tx *sql.Tx) error
	CommitTx(tx *sql.Tx) error
}

package mysql

import (
	"context"
	"database/sql"

	"github.com/fauzanmh/olp-user/entity"
)

type Repository interface {
	// Courses
	GetCourses(ctx context.Context) ([]entity.GetCoursesRow, error)

	// Course Category
	GetAllCourseCategory(ctx context.Context) ([]entity.GetAllCourseCategoryRow, error)
	GetPopularCourseCategory(ctx context.Context) (entity.GetPopularCourseCategoryRow, error)

	//Tx
	BeginTx(ctx context.Context) (*sql.Tx, error)
	WithTx(tx *sql.Tx) *Queries
	RollbackTx(tx *sql.Tx) error
	CommitTx(tx *sql.Tx) error
}

package mysql

import (
	"context"
	"database/sql"

	"github.com/fauzanmh/olp-user/entity"
)

type Repository interface {
	// Courses
	CreateCourse(ctx context.Context, args *entity.CreateCourseParams) (err error)
	DeleteCourse(ctx context.Context, arg *entity.DeleteCourseParams) error
	GetAllCourses(ctx context.Context) ([]entity.GetAllCoursesRow, error)
	GetOneCourse(ctx context.Context, id int64) (entity.GetOneCourseRow, error)
	UpdateCourse(ctx context.Context, arg *entity.UpdateCourseParams) error

	// Course Category
	GetOneCourseCategory(ctx context.Context, id int32) (entity.GetOneCourseCategoryRow, error)

	// Statistic
	GetTotalCourse(ctx context.Context) (int64, error)
	GetTotalCourseIsFree(ctx context.Context) (int64, error)

	//Tx
	BeginTx(ctx context.Context) (*sql.Tx, error)
	WithTx(tx *sql.Tx) *Queries
	RollbackTx(tx *sql.Tx) error
	CommitTx(tx *sql.Tx) error
}

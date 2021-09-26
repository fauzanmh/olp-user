package mysql

import (
	"context"
	"database/sql"

	"github.com/fauzanmh/olp-user/entity"
)

type Repository interface {

	// Course Category
	GetAllCourseCategory(ctx context.Context) ([]entity.GetAllCourseCategoryRow, error)

	//Tx
	BeginTx(ctx context.Context) (*sql.Tx, error)
	WithTx(tx *sql.Tx) *Queries
	RollbackTx(tx *sql.Tx) error
	CommitTx(tx *sql.Tx) error
}

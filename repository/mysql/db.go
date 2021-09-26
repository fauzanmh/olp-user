package mysql

import (
	"context"
	"database/sql"
	"fmt"
)

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

func New(db DBTX) *Queries {
	return &Queries{db: db}
}

func Prepare(ctx context.Context, db DBTX) (*Queries, error) {
	q := Queries{db: db}
	var err error
	if q.createCourseStmt, err = db.PrepareContext(ctx, createCourse); err != nil {
		return nil, fmt.Errorf("error preparing query CreateCourse: %w", err)
	}
	if q.deleteCourseStmt, err = db.PrepareContext(ctx, deleteCourse); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteCourse: %w", err)
	}
	if q.getAllCoursesStmt, err = db.PrepareContext(ctx, getAllCourses); err != nil {
		return nil, fmt.Errorf("error preparing query GetAllCourses: %w", err)
	}
	if q.getOneCourseStmt, err = db.PrepareContext(ctx, getOneCourse); err != nil {
		return nil, fmt.Errorf("error preparing query GetOneCourse: %w", err)
	}
	if q.getOneCourseCategoryStmt, err = db.PrepareContext(ctx, getOneCourseCategory); err != nil {
		return nil, fmt.Errorf("error preparing query GetOneCourseCategory: %w", err)
	}
	if q.getTotalCourseStmt, err = db.PrepareContext(ctx, getTotalCourse); err != nil {
		return nil, fmt.Errorf("error preparing query GetTotalCourse: %w", err)
	}
	if q.getTotalCourseIsFreeStmt, err = db.PrepareContext(ctx, getTotalCourseIsFree); err != nil {
		return nil, fmt.Errorf("error preparing query GetTotalCourseIsFree: %w", err)
	}
	if q.updateCourseStmt, err = db.PrepareContext(ctx, updateCourse); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateCourse: %w", err)
	}
	return &q, nil
}

func (q *Queries) Close() error {
	var err error
	if q.createCourseStmt != nil {
		if cerr := q.createCourseStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createCourseStmt: %w", cerr)
		}
	}
	if q.deleteCourseStmt != nil {
		if cerr := q.deleteCourseStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteCourseStmt: %w", cerr)
		}
	}
	if q.getAllCoursesStmt != nil {
		if cerr := q.getAllCoursesStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getAllCoursesStmt: %w", cerr)
		}
	}
	if q.getOneCourseStmt != nil {
		if cerr := q.getOneCourseStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getOneCourseStmt: %w", cerr)
		}
	}
	if q.getOneCourseCategoryStmt != nil {
		if cerr := q.getOneCourseCategoryStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getOneCourseCategoryStmt: %w", cerr)
		}
	}
	if q.getTotalCourseStmt != nil {
		if cerr := q.getTotalCourseStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getTotalCourseStmt: %w", cerr)
		}
	}
	if q.getTotalCourseIsFreeStmt != nil {
		if cerr := q.getTotalCourseIsFreeStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getTotalCourseIsFreeStmt: %w", cerr)
		}
	}
	if q.updateCourseStmt != nil {
		if cerr := q.updateCourseStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateCourseStmt: %w", cerr)
		}
	}
	return err
}

func (q *Queries) exec(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (sql.Result, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).ExecContext(ctx, args...)
	case stmt != nil:
		return stmt.ExecContext(ctx, args...)
	default:
		return q.db.ExecContext(ctx, query, args...)
	}
}

func (q *Queries) query(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (*sql.Rows, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryContext(ctx, args...)
	default:
		return q.db.QueryContext(ctx, query, args...)
	}
}

func (q *Queries) queryRow(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) *sql.Row {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryRowContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryRowContext(ctx, args...)
	default:
		return q.db.QueryRowContext(ctx, query, args...)
	}
}

type Queries struct {
	db                       DBTX
	tx                       *sql.Tx
	createCourseStmt         *sql.Stmt
	deleteCourseStmt         *sql.Stmt
	getAllCoursesStmt        *sql.Stmt
	getOneCourseStmt         *sql.Stmt
	getOneCourseCategoryStmt *sql.Stmt
	getTotalCourseStmt       *sql.Stmt
	getTotalCourseIsFreeStmt *sql.Stmt
	updateCourseStmt         *sql.Stmt
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db:                       tx,
		tx:                       tx,
		createCourseStmt:         q.createCourseStmt,
		deleteCourseStmt:         q.deleteCourseStmt,
		getAllCoursesStmt:        q.getAllCoursesStmt,
		getOneCourseStmt:         q.getOneCourseStmt,
		getOneCourseCategoryStmt: q.getOneCourseCategoryStmt,
		getTotalCourseStmt:       q.getTotalCourseStmt,
		getTotalCourseIsFreeStmt: q.getTotalCourseIsFreeStmt,
		updateCourseStmt:         q.updateCourseStmt,
	}
}

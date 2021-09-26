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
	if q.getAllCourseCategoryStmt, err = db.PrepareContext(ctx, getAllCourseCategory); err != nil {
		return nil, fmt.Errorf("error preparing query GetAllCourseCategory: %w", err)
	}
	if q.getCoursesStmt, err = db.PrepareContext(ctx, getCourses); err != nil {
		return nil, fmt.Errorf("error preparing query GetCourses: %w", err)
	}
	if q.getPopularCourseCategoryStmt, err = db.PrepareContext(ctx, getPopularCourseCategory); err != nil {
		return nil, fmt.Errorf("error preparing query GetPopularCourseCategory: %w", err)
	}
	return &q, nil
}

func (q *Queries) Close() error {
	var err error
	if q.getAllCourseCategoryStmt != nil {
		if cerr := q.getAllCourseCategoryStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getAllCourseCategoryStmt: %w", cerr)
		}
	}
	if q.getCoursesStmt != nil {
		if cerr := q.getCoursesStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getCoursesStmt: %w", cerr)
		}
	}
	if q.getPopularCourseCategoryStmt != nil {
		if cerr := q.getPopularCourseCategoryStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getPopularCourseCategoryStmt: %w", cerr)
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
	db                           DBTX
	tx                           *sql.Tx
	getAllCourseCategoryStmt     *sql.Stmt
	getCoursesStmt               *sql.Stmt
	getPopularCourseCategoryStmt *sql.Stmt
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db:                           tx,
		tx:                           tx,
		getAllCourseCategoryStmt:     q.getAllCourseCategoryStmt,
		getCoursesStmt:               q.getCoursesStmt,
		getPopularCourseCategoryStmt: q.getPopularCourseCategoryStmt,
	}
}

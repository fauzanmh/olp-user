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
	if q.checkEmailStmt, err = db.PrepareContext(ctx, checkEmail); err != nil {
		return nil, fmt.Errorf("error preparing query CheckEmail: %w", err)
	}
	if q.checkMemberStmt, err = db.PrepareContext(ctx, checkMember); err != nil {
		return nil, fmt.Errorf("error preparing query CheckMember: %w", err)
	}
	if q.deleteMemberStmt, err = db.PrepareContext(ctx, deleteMember); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteMember: %w", err)
	}
	if q.getAllCourseCategoryStmt, err = db.PrepareContext(ctx, getAllCourseCategory); err != nil {
		return nil, fmt.Errorf("error preparing query GetAllCourseCategory: %w", err)
	}
	if q.getCourseDetailStmt, err = db.PrepareContext(ctx, getCourseDetail); err != nil {
		return nil, fmt.Errorf("error preparing query GetCourseDetail: %w", err)
	}
	if q.getCoursesStmt, err = db.PrepareContext(ctx, getCourses); err != nil {
		return nil, fmt.Errorf("error preparing query GetCourses: %w", err)
	}
	if q.getPopularCourseCategoryStmt, err = db.PrepareContext(ctx, getPopularCourseCategory); err != nil {
		return nil, fmt.Errorf("error preparing query GetPopularCourseCategory: %w", err)
	}
	if q.registerStmt, err = db.PrepareContext(ctx, register); err != nil {
		return nil, fmt.Errorf("error preparing query Register: %w", err)
	}
	return &q, nil
}

func (q *Queries) Close() error {
	var err error
	if q.checkEmailStmt != nil {
		if cerr := q.checkEmailStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing checkEmailStmt: %w", cerr)
		}
	}
	if q.checkMemberStmt != nil {
		if cerr := q.checkMemberStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing checkMemberStmt: %w", cerr)
		}
	}
	if q.deleteMemberStmt != nil {
		if cerr := q.deleteMemberStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteMemberStmt: %w", cerr)
		}
	}
	if q.getAllCourseCategoryStmt != nil {
		if cerr := q.getAllCourseCategoryStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getAllCourseCategoryStmt: %w", cerr)
		}
	}
	if q.getCourseDetailStmt != nil {
		if cerr := q.getCourseDetailStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getCourseDetailStmt: %w", cerr)
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
	if q.registerStmt != nil {
		if cerr := q.registerStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing registerStmt: %w", cerr)
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
	checkEmailStmt               *sql.Stmt
	checkMemberStmt              *sql.Stmt
	deleteMemberStmt             *sql.Stmt
	getAllCourseCategoryStmt     *sql.Stmt
	getCourseDetailStmt          *sql.Stmt
	getCoursesStmt               *sql.Stmt
	getPopularCourseCategoryStmt *sql.Stmt
	registerStmt                 *sql.Stmt
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db:                           tx,
		tx:                           tx,
		checkEmailStmt:               q.checkEmailStmt,
		checkMemberStmt:              q.checkMemberStmt,
		deleteMemberStmt:             q.deleteMemberStmt,
		getAllCourseCategoryStmt:     q.getAllCourseCategoryStmt,
		getCourseDetailStmt:          q.getCourseDetailStmt,
		getCoursesStmt:               q.getCoursesStmt,
		getPopularCourseCategoryStmt: q.getPopularCourseCategoryStmt,
		registerStmt:                 q.registerStmt,
	}
}

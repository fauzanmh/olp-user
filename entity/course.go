package entity

import (
	"database/sql"
)

// model
type Course struct {
	ID               int64         `json:"id"`
	CourseCategoryID int32         `json:"course_category_id"`
	Name             string        `json:"name"`
	Description      string        `json:"description"`
	Price            string        `json:"price"`
	CreatedAt        int64         `json:"created_at"`
	UpdatedAt        sql.NullInt64 `json:"updated_at"`
	DeletedAt        sql.NullInt64 `json:"deleted_at"`
}

// --- params and rows --- //

// for create course
type CreateCourseParams struct {
	CourseCategoryID int32         `json:"course_category_id"`
	Name             string        `json:"name"`
	Description      string        `json:"description"`
	Price            string        `json:"price"`
	CreatedAt        int64         `json:"created_at"`
	UpdatedAt        sql.NullInt64 `json:"updated_at"`
}

// for soft delete course
type DeleteCourseParams struct {
	DeletedAt sql.NullInt64 `json:"deleted_at"`
	ID        int64         `json:"id"`
}

// for get all courses
type GetAllCoursesRow struct {
	ID               int64  `json:"id"`
	CourseCategoryID int32  `json:"course_category_id"`
	Name             string `json:"name"`
	Description      string `json:"description"`
	Price            string `json:"price"`
}

// for get one course
type GetOneCourseRow struct {
	ID               int64  `json:"id"`
	CourseCategoryID int32  `json:"course_category_id"`
	Name             string `json:"name"`
	Description      string `json:"description"`
	Price            string `json:"price"`
}

// for update course
type UpdateCourseParams struct {
	CourseCategoryID int32         `json:"course_category_id"`
	Name             string        `json:"name"`
	Description      string        `json:"description"`
	Price            string        `json:"price"`
	UpdatedAt        sql.NullInt64 `json:"updated_at"`
	ID               int64         `json:"id"`
}

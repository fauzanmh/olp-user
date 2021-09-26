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

// for get courses
type GetCoursesRow struct {
	ID                 int64  `json:"id"`
	CourseCategoryID   int32  `json:"course_category_id"`
	Name               string `json:"name"`
	Description        string `json:"description"`
	Price              string `json:"price"`
	CourseCategoryName string `json:"course_category_name"`
}

// for get course detail
type GetCourseDetailRow struct {
	ID                 int64  `json:"id"`
	CourseCategoryID   int32  `json:"course_category_id"`
	Name               string `json:"name"`
	Description        string `json:"description"`
	Price              string `json:"price"`
	CourseCategoryName string `json:"course_category_name"`
}

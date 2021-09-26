package entity

import "database/sql"

// model
type CourseCategory struct {
	ID        int32         `json:"id"`
	Name      string        `json:"name"`
	TotalUsed int32         `json:"total_used"`
	CreatedAt int64         `json:"created_at"`
	UpdatedAt sql.NullInt64 `json:"updated_at"`
	DeletedAt sql.NullInt64 `json:"deleted_at"`
}

// --- params and rows --- //
type GetAllCourseCategoryRow struct {
	ID        int32  `json:"id"`
	Name      string `json:"name"`
	TotalUsed int32  `json:"total_used"`
}

type GetPopularCourseCategoryRow struct {
	ID        int32  `json:"id"`
	Name      string `json:"name"`
	TotalUsed int32  `json:"total_used"`
}

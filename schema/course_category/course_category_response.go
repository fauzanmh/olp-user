package course_category

type GetCourseCategory struct {
	ID        int32  `json:"id"`
	Name      string `json:"name"`
	TotalUsed int32  `json:"total_used"`
}

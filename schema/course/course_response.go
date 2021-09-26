package course

type GetCoursesResponse struct {
	ID                 int64  `json:"id"`
	CourseCategoryID   int32  `json:"course_category_id"`
	Name               string `json:"name"`
	Description        string `json:"description"`
	Price              string `json:"price"`
	CourseCategoryName string `json:"course_category_name"`
}

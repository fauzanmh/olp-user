package course

type GetAllCoursesResponse struct {
	ID               int64  `json:"id"`
	CourseCategoryID int32  `json:"course_category_id"`
	Name             string `json:"name"`
	Description      string `json:"description"`
	Price            string `json:"price"`
}

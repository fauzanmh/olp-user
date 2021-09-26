package statistic

type GetStatisticResponse struct {
	TotalUser       int64 `json:"total_user"`
	TotalCourse     int64 `json:"total_course"`
	TotalCourseFree int64 `json:"total_course_free"`
}

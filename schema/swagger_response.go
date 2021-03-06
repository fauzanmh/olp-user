package schema

import (
	"github.com/fauzanmh/olp-user/schema/course"
	"github.com/fauzanmh/olp-user/schema/course_category"
)

type SwaggerGetCoursesResponse struct {
	Base
	Data []course.CourseResponse `json:"data"`
}

type SwaggerGetAllCourseCategoriesResponse struct {
	Base
	Data []course_category.GetCourseCategory `json:"data"`
}

type SwaggerGetPopularCourseCategoryResponse struct {
	Base
	Data course_category.GetCourseCategory `json:"data"`
}

type SwaggerGetCourseDetailResponse struct {
	Base
	Data course.CourseResponse `json:"data"`
}

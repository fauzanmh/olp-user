package schema

import (
	"github.com/fauzanmh/olp-user/schema/course"
	"github.com/fauzanmh/olp-user/schema/course_category"
)

type SwaggerGetAllCoursesResponse struct {
	Base
	Data course.GetAllCoursesResponse `json:"data"`
}

type SwaggerGetAllCourseCategoriesResponse struct {
	Base
	Data []course_category.GetAllCourseCategoriesResponse `json:"data"`
}

package schema

import (
	"github.com/fauzanmh/olp-user/schema/course"
	"github.com/fauzanmh/olp-user/schema/statistic"
)

type SwaggerGetAllCoursesResponse struct {
	Base
	Data course.GetAllCoursesResponse `json:"data"`
}

type SwaggerGetStatisticResponse struct {
	Base
	Data statistic.GetStatisticResponse `json:"data"`
}

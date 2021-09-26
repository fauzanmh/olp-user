package course

import (
	"context"

	"github.com/fauzanmh/olp-user/schema/course"
)

type Usecase interface {
	Get(ctx context.Context) (res []course.GetAllCoursesResponse, err error)
	Create(ctx context.Context, req *course.CourseCreateRequest) (err error)
	Update(ctx context.Context, req *course.CourseUpdateRequest) (err error)
	Delete(ctx context.Context, req *course.CourseDeleteRequest) (err error)
}

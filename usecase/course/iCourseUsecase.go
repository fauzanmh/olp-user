package course

import (
	"context"

	"github.com/fauzanmh/olp-user/schema/course"
)

type Usecase interface {
	Get(ctx context.Context) (res []course.CourseResponse, err error)
	GetDetail(ctx context.Context, req *course.CourseDetailRequest) (res course.CourseResponse, err error)
}

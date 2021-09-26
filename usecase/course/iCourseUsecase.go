package course

import (
	"context"

	"github.com/fauzanmh/olp-user/schema/course"
)

type Usecase interface {
	Get(ctx context.Context) (res []course.GetCoursesResponse, err error)
}

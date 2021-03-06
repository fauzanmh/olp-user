package course_category

import (
	"context"

	"github.com/fauzanmh/olp-user/schema/course_category"
)

type Usecase interface {
	Get(ctx context.Context) (res []course_category.GetCourseCategory, err error)
	GetPopular(ctx context.Context) (res course_category.GetCourseCategory, err error)
}

package course

import (
	"context"

	appInit "github.com/fauzanmh/olp-user/init"
	mysqlRepo "github.com/fauzanmh/olp-user/repository/mysql"
	"github.com/fauzanmh/olp-user/schema/course"
)

type usecase struct {
	config    *appInit.Config
	mysqlRepo mysqlRepo.Repository
}

func NewCourseUseCase(config *appInit.Config, mysqlRepo mysqlRepo.Repository) Usecase {
	return &usecase{
		config:    config,
		mysqlRepo: mysqlRepo,
	}
}

// --- get all course --- ///
func (u *usecase) Get(ctx context.Context) (res []course.GetAllCoursesResponse, err error) {

	return
}

// --- create course --- ///
func (u *usecase) Create(ctx context.Context, req *course.CourseCreateRequest) (err error) {

	return
}

// --- update course --- ///
func (u *usecase) Update(ctx context.Context, req *course.CourseUpdateRequest) (err error) {

	return
}

// --- delete course --- ///
func (u *usecase) Delete(ctx context.Context, req *course.CourseDeleteRequest) (err error) {

	return
}

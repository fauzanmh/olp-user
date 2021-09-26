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
func (u *usecase) Get(ctx context.Context) (res []course.GetCoursesResponse, err error) {
	// get data from database
	data, err := u.mysqlRepo.GetCourses(ctx)
	if err != nil {
		return
	}

	for idx := range data {
		res = append(res, course.GetCoursesResponse{
			ID:                 data[idx].ID,
			CourseCategoryID:   data[idx].CourseCategoryID,
			Name:               data[idx].Name,
			Description:        data[idx].Description,
			Price:              data[idx].Price,
			CourseCategoryName: data[idx].CourseCategoryName,
		})
	}

	return
}

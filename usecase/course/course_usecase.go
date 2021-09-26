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

// --- get courses --- ///
func (u *usecase) Get(ctx context.Context) (res []course.CourseResponse, err error) {
	// get data from database
	data, err := u.mysqlRepo.GetCourses(ctx)
	if err != nil {
		return
	}

	// convert from entity to schema
	for idx := range data {
		res = append(res, course.CourseResponse{
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

// --- get course detail --- ///
func (u *usecase) GetDetail(ctx context.Context, req *course.CourseDetailRequest) (res course.CourseResponse, err error) {
	// get data from database
	data, err := u.mysqlRepo.GetCourseDetail(ctx, req.ID)
	if err != nil {
		return
	}

	res = course.CourseResponse{
		ID:                 data.ID,
		CourseCategoryID:   data.CourseCategoryID,
		Name:               data.Name,
		Description:        data.Description,
		Price:              data.Price,
		CourseCategoryName: data.CourseCategoryName,
	}

	return
}

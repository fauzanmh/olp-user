package course_category

import (
	"context"

	appInit "github.com/fauzanmh/olp-user/init"
	mysqlRepo "github.com/fauzanmh/olp-user/repository/mysql"
	"github.com/fauzanmh/olp-user/schema/course_category"
)

type usecase struct {
	config    *appInit.Config
	mysqlRepo mysqlRepo.Repository
}

func NewCourseCategoryUseCase(config *appInit.Config, mysqlRepo mysqlRepo.Repository) Usecase {
	return &usecase{
		config:    config,
		mysqlRepo: mysqlRepo,
	}
}

// --- get all course categories --- ///
func (u *usecase) Get(ctx context.Context) (res []course_category.GetAllCourseCategoriesResponse, err error) {
	// get from database
	data, err := u.mysqlRepo.GetAllCourseCategory(ctx)
	if err != nil {
		return
	}

	for idx := range data {
		res = append(res, course_category.GetAllCourseCategoriesResponse{
			ID:        data[idx].ID,
			Name:      data[idx].Name,
			TotalUsed: data[idx].TotalUsed,
		})
	}

	return
}

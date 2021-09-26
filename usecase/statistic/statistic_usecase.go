package statistic

import (
	"context"

	appInit "github.com/fauzanmh/olp-user/init"
	mysqlRepo "github.com/fauzanmh/olp-user/repository/mysql"
	"github.com/fauzanmh/olp-user/schema/statistic"
)

type usecase struct {
	config    *appInit.Config
	mysqlRepo mysqlRepo.Repository
}

func NewStatisticUseCase(config *appInit.Config, mysqlRepo mysqlRepo.Repository) Usecase {
	return &usecase{
		config:    config,
		mysqlRepo: mysqlRepo,
	}
}

// --- get statistic --- ///
func (u *usecase) Get(ctx context.Context) (res statistic.GetStatisticResponse, err error) {
	// get total course
	totalCourse, err := u.mysqlRepo.GetTotalCourse(ctx)
	if err != nil {
		return
	}

	// get total course is free
	totalCourseFree, err := u.mysqlRepo.GetTotalCourseIsFree(ctx)
	if err != nil {
		return
	}

	// get total user
	totalUser := int64(0)

	res = statistic.GetStatisticResponse{
		TotalCourseFree: totalCourseFree,
		TotalCourse:     totalCourse,
		TotalUser:       totalUser,
	}

	return
}

package course

import (
	"context"
	"fmt"

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
func (u *usecase) Get(ctx context.Context, req *course.CourseGetRequest) (res []course.CourseResponse, err error) {
	res = []course.CourseResponse{}

	// check if search not null
	search := fmt.Sprintf("%s%s", "%", "%")
	if req.Search != "" {
		search = fmt.Sprintf("%s%s%s", "%", req.Search, "%")
	}

	// check if sort not null and default sort is updated_at
	sort := "ORDER BY updated_at DESC"
	if req.Sort != "" {
		switch req.Sort {
		case "lowest_price":
			sort = `AND price != "0" ORDER BY price ASC`
		case "highest_price":
			sort = `AND price != "0" ORDER BY price DESC`
		default:
			sort = `AND price = "0" ORDER BY updated_at DESC`
		}
	}

	// get data from database
	data, err := u.mysqlRepo.GetCourses(ctx, search, sort)
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

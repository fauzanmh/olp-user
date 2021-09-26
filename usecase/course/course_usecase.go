package course

import (
	"context"
	"database/sql"
	"time"

	"github.com/fauzanmh/olp-user/entity"
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
	// get from database
	data, err := u.mysqlRepo.GetAllCourses(ctx)
	if err != nil {
		return
	}

	// convert from entity to schema
	for idx := range data {
		res = append(res, course.GetAllCoursesResponse{
			ID:               data[idx].ID,
			CourseCategoryID: data[idx].CourseCategoryID,
			Name:             data[idx].Name,
			Description:      data[idx].Description,
			Price:            data[idx].Price,
		})
	}

	return
}

// --- create course --- ///
func (u *usecase) Create(ctx context.Context, req *course.CourseCreateRequest) (err error) {
	// check if course category is exists
	_, err = u.mysqlRepo.GetOneCourseCategory(ctx, req.CourseCategoryID)
	if err != nil {
		return
	}

	// arguments
	createCourseParams := &entity.CreateCourseParams{
		CourseCategoryID: req.CourseCategoryID,
		Name:             req.Name,
		Description:      req.Description,
		Price:            req.Price,
		CreatedAt:        time.Now().Unix(),
		UpdatedAt:        sql.NullInt64{Int64: time.Now().Unix(), Valid: true},
	}

	// store to database
	err = u.mysqlRepo.CreateCourse(ctx, createCourseParams)
	if err != nil {
		return
	}

	return
}

// --- update course --- ///
func (u *usecase) Update(ctx context.Context, req *course.CourseUpdateRequest) (err error) {
	// check if course is exists
	_, err = u.mysqlRepo.GetOneCourse(ctx, req.ID)
	if err != nil {
		return
	}

	// check if course category is exists
	_, err = u.mysqlRepo.GetOneCourseCategory(ctx, req.CourseCategoryID)
	if err != nil {
		return
	}

	// arguments
	updateCourseParams := &entity.UpdateCourseParams{
		ID:               req.ID,
		CourseCategoryID: req.CourseCategoryID,
		Name:             req.Name,
		Description:      req.Description,
		Price:            req.Price,
		UpdatedAt:        sql.NullInt64{Int64: time.Now().Unix(), Valid: true},
	}

	err = u.mysqlRepo.UpdateCourse(ctx, updateCourseParams)
	if err != nil {
		return
	}

	return
}

// --- delete course --- ///
func (u *usecase) Delete(ctx context.Context, req *course.CourseDeleteRequest) (err error) {
	// check if course is exists
	_, err = u.mysqlRepo.GetOneCourse(ctx, req.ID)
	if err != nil {
		return
	}

	// arguments
	deleteCourseParams := &entity.DeleteCourseParams{
		ID:        req.ID,
		DeletedAt: sql.NullInt64{Int64: time.Now().Unix(), Valid: true},
	}

	// delete from database
	err = u.mysqlRepo.DeleteCourse(ctx, deleteCourseParams)
	if err != nil {
		return
	}

	return
}

package http

import (
	"github.com/fauzanmh/olp-user/pkg/util"
	"github.com/fauzanmh/olp-user/schema/course"
	usecaseCourse "github.com/fauzanmh/olp-user/usecase/course"
	"github.com/labstack/echo/v4"
)

type CourseHandler struct {
	usecase usecaseCourse.Usecase
}

func NewCourseHandler(e *echo.Group, usecase usecaseCourse.Usecase) {
	handler := &CourseHandler{
		usecase: usecase,
	}

	routerV1 := e.Group("/v1")
	routerV1.GET("/course", handler.Get)
	routerV1.GET("/course/:id/detail", handler.GetDetail)
}

// Get godoc
// @Summary Get Courses
// @Description Get Courses
// @Tags Course
// @Accept json
// @Produce json
// @Param search query string false "search course by name"
// @Param sort query string false "sort by {lowest price|highest price|free}"
// @Success 200 {object} schema.SwaggerGetCoursesResponse
// @Failure 400 {object} schema.Base
// @Failure 401 {object} schema.Base
// @Failure 404 {object} schema.Base
// @Failure 500 {object} schema.Base
// @Router /v1/course [get]
func (h *CourseHandler) Get(c echo.Context) error {
	req := course.CourseGetRequest{}
	ctx := c.Request().Context()

	// parsing
	err := util.ParsingParameter(c, &req)
	if err != nil {
		return util.ErrorParsing(c, err, nil)
	}

	// validate
	err = util.ValidateParameter(c, &req)
	if err != nil {
		return util.ErrorValidate(c, err, nil)
	}

	data, err := h.usecase.Get(ctx, &req)
	if err != nil {
		return util.ErrorResponse(c, err, nil)
	}

	return util.SuccessResponse(c, "success get courses", data)
}

// GetDetail godoc
// @Summary Get Course Detail
// @Description Get Course Detail
// @Tags Course
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} schema.SwaggerGetCourseDetailResponse
// @Failure 400 {object} schema.Base
// @Failure 401 {object} schema.Base
// @Failure 404 {object} schema.Base
// @Failure 500 {object} schema.Base
// @Router /v1/course/{id}/detail [get]
func (h *CourseHandler) GetDetail(c echo.Context) error {
	req := course.CourseDetailRequest{}
	ctx := c.Request().Context()

	// parsing
	err := util.ParsingParameter(c, &req)
	if err != nil {
		return util.ErrorParsing(c, err, nil)
	}

	// validate
	err = util.ValidateParameter(c, &req)
	if err != nil {
		return util.ErrorValidate(c, err, nil)
	}

	data, err := h.usecase.GetDetail(ctx, &req)
	if err != nil {
		return util.ErrorResponse(c, err, nil)
	}

	return util.SuccessResponse(c, "success get course detail", data)
}

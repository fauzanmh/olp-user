package http

import (
	"github.com/fauzanmh/olp-user/pkg/util"
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
}

// Get godoc
// @Summary Get All Courses
// @Description Get All Courses
// @Tags Course
// @Accept json
// @Produce json
// @Success 200 {object} schema.SwaggerGetCoursesResponse
// @Failure 400 {object} schema.Base
// @Failure 401 {object} schema.Base
// @Failure 404 {object} schema.Base
// @Failure 500 {object} schema.Base
// @Router /v1/course [get]
func (h *CourseHandler) Get(c echo.Context) error {
	ctx := c.Request().Context()

	data, err := h.usecase.Get(ctx)
	if err != nil {
		return util.ErrorResponse(c, err, nil)
	}

	return util.SuccessResponse(c, "success get courses", data)
}

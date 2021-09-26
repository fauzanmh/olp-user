package http

import (
	"github.com/fauzanmh/olp-user/pkg/util"
	usecase "github.com/fauzanmh/olp-user/usecase/course_category"
	"github.com/labstack/echo/v4"
)

type CourseCategoryHandler struct {
	usecase usecase.Usecase
}

func NewCourseCategoryHandler(e *echo.Group, usecase usecase.Usecase) {
	handler := &CourseCategoryHandler{
		usecase: usecase,
	}

	routerV1 := e.Group("/v1")
	routerV1.GET("/course/category", handler.Get)
	routerV1.GET("/course/category/popular", handler.GetPopular)
}

// Get godoc
// @Summary Get Course Category
// @Description Get Course Category
// @Tags Course Category
// @Accept json
// @Produce json
// @Success 200 {object} schema.SwaggerGetAllCourseCategoriesResponse
// @Failure 400 {object} schema.Base
// @Failure 401 {object} schema.Base
// @Failure 404 {object} schema.Base
// @Failure 500 {object} schema.Base
// @Router /v1/course/category [get]
func (h *CourseCategoryHandler) Get(c echo.Context) error {
	ctx := c.Request().Context()

	data, err := h.usecase.Get(ctx)
	if err != nil {
		return util.ErrorResponse(c, err, nil)
	}

	return util.SuccessResponse(c, "success get course categories", data)
}

// GetPopular godoc
// @Summary Get Popular Course Category
// @Description Get Popular Course Category
// @Tags Course Category
// @Accept json
// @Produce json
// @Success 200 {object} schema.SwaggerGetPopularCourseCategoryResponse
// @Failure 400 {object} schema.Base
// @Failure 401 {object} schema.Base
// @Failure 404 {object} schema.Base
// @Failure 500 {object} schema.Base
// @Router /v1/course/category/popular [get]
func (h *CourseCategoryHandler) GetPopular(c echo.Context) error {
	ctx := c.Request().Context()

	data, err := h.usecase.GetPopular(ctx)
	if err != nil {
		return util.ErrorResponse(c, err, nil)
	}

	return util.SuccessResponse(c, "success get popular course category", data)
}

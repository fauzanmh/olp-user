package http

import (
	"github.com/fauzanmh/olp-user/pkg/util"
	usecase "github.com/fauzanmh/olp-user/usecase/statistic"
	"github.com/labstack/echo/v4"
)

type StatisticHandler struct {
	usecase usecase.Usecase
}

func NewStatisticHandler(e *echo.Group, usecase usecase.Usecase) {
	handler := &StatisticHandler{
		usecase: usecase,
	}

	routerV1 := e.Group("/v1")
	routerV1.GET("/statistic", handler.Get)
}

// Get godoc
// @Summary Get Statistic
// @Description Get Statistic
// @Tags Statistic
// @Accept json
// @Produce json
// @Success 200 {object} schema.SwaggerGetStatisticResponse
// @Failure 400 {object} schema.Base
// @Failure 401 {object} schema.Base
// @Failure 404 {object} schema.Base
// @Failure 500 {object} schema.Base
// @Router /v1/statistic [get]
func (h *StatisticHandler) Get(c echo.Context) error {
	ctx := c.Request().Context()

	data, err := h.usecase.Get(ctx)
	if err != nil {
		return util.ErrorResponse(c, err, nil)
	}

	return util.SuccessResponse(c, "success get statistic", data)
}

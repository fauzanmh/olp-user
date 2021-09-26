package http

import (
	"github.com/fauzanmh/olp-user/pkg/util"
	"github.com/fauzanmh/olp-user/schema/member"
	usecase "github.com/fauzanmh/olp-user/usecase/member"
	"github.com/labstack/echo/v4"
)

type MemberHandler struct {
	usecase usecase.Usecase
}

func NewMemberHandler(e *echo.Group, usecase usecase.Usecase) {
	handler := &MemberHandler{
		usecase: usecase,
	}

	routerV1 := e.Group("/v1")
	routerV1.POST("/register", handler.Register)
}

// Register godoc
// @Summary Register
// @Description Register
// @Tags Member
// @Accept json
// @Produce json
// @Param request body member.RegisterRequest{} true "Request Body"
// @Success 200 {object} schema.Base
// @Failure 400 {object} schema.Base
// @Failure 401 {object} schema.Base
// @Failure 404 {object} schema.Base
// @Failure 500 {object} schema.Base
// @Router /v1/register [post]
func (h *MemberHandler) Register(c echo.Context) error {
	req := member.RegisterRequest{}
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

	err = h.usecase.Register(ctx, &req)
	if err != nil {
		return util.ErrorResponse(c, err, nil)
	}

	return util.SuccessResponse(c, "success register", nil)
}

package http

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/malikkhoiri/auth-svc/domain"
	"github.com/malikkhoiri/auth-svc/helper"
)

type UserHandler struct {
	UUsecase domain.UserUsecase
}

func NewUserHandler(e *echo.Echo, u domain.UserUsecase) {
	handler := &UserHandler{
		UUsecase: u,
	}
	e.GET("/user/:id", handler.GetByID)
	e.POST("/user", handler.Store)
}

func (uh *UserHandler) GetByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return err
	}

	user, err := uh.UUsecase.GetByID(c.Request().Context(), int64(id))

	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.BadResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    user,
	})
}

func (uh *UserHandler) Store(c echo.Context) error {
	user := &domain.User{}
	err := c.Bind(user)

	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, helper.BadResponse{
			Status:  http.StatusUnprocessableEntity,
			Message: err.Error(),
		})
	}

	err = c.Validate(user)

	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.BadResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	ctx := c.Request().Context()
	err = uh.UUsecase.Store(ctx, user)

	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.BadResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse{
		Status:  http.StatusOK,
		Message: "Success",
	})
}

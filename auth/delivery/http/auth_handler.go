package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/malikkhoiri/auth-svc/domain"
	"github.com/malikkhoiri/auth-svc/helper"
)

// AuthHandler represent the httphandler for auth
type AuthHandler struct {
	AUsecase domain.AuthUsecase
}

// NewAuthHandler will initialize the auth/ resources endpoint
func NewAuthHandler(e *echo.Echo, au domain.AuthUsecase) {
	handler := &AuthHandler{
		AUsecase: au,
	}
	e.POST("/login", handler.LoginByUsernameAndPassword)
}

func (ah *AuthHandler) LoginByUsernameAndPassword(c echo.Context) error {
	auth := &domain.Auth{}
	err := c.Bind(auth)

	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, helper.BadResponse{
			Status:  http.StatusUnprocessableEntity,
			Message: err.Error(),
		})
	}

	err = c.Validate(auth)

	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.BadResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	ctx := c.Request().Context()
	res, err := ah.AUsecase.LoginByUsernameAndPassword(ctx, auth.Username, auth.Password)

	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.BadResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse{
		Status: http.StatusOK,
		Data:   res,
	})
}

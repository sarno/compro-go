package handler

import (
	"compro/internal/adapter/handler/request"
	"compro/internal/adapter/handler/response"
	"compro/internal/core/domain/entity"
	"compro/internal/core/service"
	"compro/utils/conv"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type IUserHandler interface {
	LoginAdmin(c echo.Context) error
}

type UserHandler struct {
	userService service.IUserService
}

// LoginAdmin implements [IUserHandler].
func (u *UserHandler) LoginAdmin(c echo.Context) error {
	var (
		req       = request.LoginRequest{}
		respLogin = response.LoginResponse{}
		resp      = response.DefaultSuccessResponse{}
		respError = response.ErrorResponseDefault{}
		ctx       = c.Request().Context()
	)
	
	if err = c.Bind(&req); err != nil {
		code = "[HANDLER] LoginAdmin - 1"
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusUnprocessableEntity, respError)
	}

	if err = c.Validate(req); err != nil {
		code = "[HANDLER] LoginAdmin - 2"
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusBadRequest, respError)
	}

	reqEntity := entity.UserEntity{
		Email:    req.Email,
		Password: req.Password,
	}

	token, err := u.userService.LoginAdmin(ctx, reqEntity)
	if err != nil {
		code = "[HANDLER] LoginAdmin - 3"
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(conv.SetHTTPStatusCode(err), respError)
	}

	respLogin.Token = token
	resp.Meta.Status = true
	resp.Meta.Message = "Success login"
	resp.Data = respLogin
	resp.Pagination = nil
	return c.JSON(http.StatusOK, resp)
}

var (
	err  error
	code string
)

func NewUserHandler(e *echo.Echo, userService service.IUserService) IUserHandler {
	userHandler := &UserHandler{
		userService: userService,
	}

	e.Use(middleware.Recover())
	e.POST("/login", userHandler.LoginAdmin)

	return userHandler
}

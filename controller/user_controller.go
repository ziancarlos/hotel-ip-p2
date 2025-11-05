package controller

import (
	"hotel_ip-p2/exception"
	"hotel_ip-p2/helper"
	"hotel_ip-p2/mapper"
	"hotel_ip-p2/model/web"
	"hotel_ip-p2/model/web/request"
	"hotel_ip-p2/model/web/response"
	"hotel_ip-p2/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	UserService service.UserService
}

func NewUserController(userService service.UserService) *UserController {
	return &UserController{
		UserService: userService,
	}
}

func (controller *UserController) Register(c echo.Context) error {
	var req request.UserRequest

	if err := c.Bind(&req); err != nil {
		return exception.NewCustomError(http.StatusBadRequest, "Invalid request body")
	}

	if err := c.Validate(&req); err != nil {
		return exception.NewCustomError(http.StatusBadRequest, err.Error())
	}

	user := mapper.ToUserDomain(req)

	result, err := controller.UserService.Register(user)
	if err != nil {
		return exception.NewCustomError(http.StatusInternalServerError, err.Error())
	}

	userResponse := mapper.ToUserResponse(result)

	return c.JSON(http.StatusCreated, web.WebResponse{
		Message: "User registered successfully",
		Data:    userResponse,
	})
}

func (controller *UserController) Login(c echo.Context) error {
	var req request.LoginRequest

	if err := c.Bind(&req); err != nil {
		return exception.NewCustomError(http.StatusBadRequest, "Invalid request body")
	}

	if err := c.Validate(&req); err != nil {
		return exception.NewCustomError(http.StatusBadRequest, err.Error())
	}

	user, err := controller.UserService.Login(req.Email, req.Password)
	if err != nil {
		return exception.NewCustomError(http.StatusUnauthorized, "Invalid credentials")
	}

	token, err := helper.GenerateToken(user.ID)
	if err != nil {
		return exception.NewCustomError(http.StatusInternalServerError, "Failed to generate token")
	}

	loginResponse := response.LoginResponse{
		Token: token,
	}

	return c.JSON(http.StatusOK, web.WebResponse{
		Message: "Login successful",
		Data:    loginResponse,
	})
}

func (controller *UserController) GetMe(c echo.Context) error {
	userID := c.Get("user_id").(int)

	user, err := controller.UserService.GetById(userID)
	if err != nil {
		return exception.NewCustomError(http.StatusNotFound, "User not found")
	}

	userResponse := mapper.ToUserResponse(user)

	return c.JSON(http.StatusOK, web.WebResponse{
		Message: "User retrieved successfully",
		Data:    userResponse,
	})
}

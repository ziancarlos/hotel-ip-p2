package controller

import (
	"errors"
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

// Register godoc
// @Summary Register a new user
// @Description Register a new user account
// @Tags users
// @Accept json
// @Produce json
// @Param request body request.UserRequest true "User registration details"
// @Success 201 {object} web.WebResponse{data=response.UserResponse} "User registered successfully"
// @Failure 400 {object} web.WebResponse "Invalid request body or validation error"
// @Failure 500 {object} web.WebResponse "Internal server error"
// @Router /users/register [post]
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
		var customErr *exception.CustomError
		if errors.As(err, &customErr) {
			return customErr
		}
		return exception.NewCustomError(http.StatusInternalServerError, err.Error())
	}

	userResponse := mapper.ToUserResponse(result)

	return c.JSON(http.StatusCreated, web.WebResponse{
		Message: "User registered successfully",
		Data:    userResponse,
	})
}

// Login godoc
// @Summary Login user
// @Description Authenticate user and get JWT token
// @Tags users
// @Accept json
// @Produce json
// @Param request body request.LoginRequest true "Login credentials"
// @Success 200 {object} web.WebResponse{data=response.LoginResponse} "Login successful"
// @Failure 400 {object} web.WebResponse "Invalid request body or validation error"
// @Failure 401 {object} web.WebResponse "Invalid credentials"
// @Failure 500 {object} web.WebResponse "Internal server error"
// @Router /users/login [post]
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
		return err
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

// GetMe godoc
// @Summary Get current user
// @Description Get current authenticated user information
// @Tags users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} web.WebResponse{data=response.UserResponse} "User retrieved successfully"
// @Failure 401 {object} web.WebResponse "Unauthorized"
// @Failure 404 {object} web.WebResponse "User not found"
// @Router /users/me [get]
func (controller *UserController) GetMe(c echo.Context) error {
	userID := c.Get("user_id").(int)

	user, err := controller.UserService.GetById(userID)
	if err != nil {
		return err
	}

	userResponse := mapper.ToUserResponse(user)

	return c.JSON(http.StatusOK, web.WebResponse{
		Message: "User retrieved successfully",
		Data:    userResponse,
	})
}

func (controller *UserController) GetById(c echo.Context) error {
	return c.JSON(http.StatusOK, web.WebResponse{
		Message: "Get user by id",
	})
}

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
	"log"
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
func (controller *UserController) Register(c echo.Context) error {
	log.Println("Request to register new user")
	var req request.UserRequest

	if err := c.Bind(&req); err != nil {
		log.Printf("Failed to bind request body: %v", err)
		return exception.NewCustomError(http.StatusBadRequest, "Invalid request body")
	}

	if err := c.Validate(&req); err != nil {
		log.Printf("Validation failed: %v", err)
		return exception.NewCustomError(http.StatusBadRequest, err.Error())
	}

	user := mapper.ToUserDomain(req)

	result, err := controller.UserService.Register(user)
	if err != nil {
		log.Printf("Failed to register user: %v", err)
		var customErr *exception.CustomError
		if errors.As(err, &customErr) {
			return customErr
		}
		return exception.NewCustomError(http.StatusInternalServerError, err.Error())
	}

	log.Printf("User registered successfully with ID: %d", result.ID)
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
func (controller *UserController) Login(c echo.Context) error {
	log.Println("Request to login user")
	var req request.LoginRequest

	if err := c.Bind(&req); err != nil {
		log.Printf("Failed to bind request body: %v", err)
		return exception.NewCustomError(http.StatusBadRequest, "Invalid request body")
	}

	if err := c.Validate(&req); err != nil {
		log.Printf("Validation failed: %v", err)
		return exception.NewCustomError(http.StatusBadRequest, err.Error())
	}

	user, err := controller.UserService.Login(req.Email, req.Password)
	if err != nil {
		log.Printf("Login failed for email %s: %v", req.Email, err)
		return err
	}

	token, err := helper.GenerateToken(user.ID)
	if err != nil {
		log.Printf("Failed to generate token: %v", err)
		return exception.NewCustomError(http.StatusInternalServerError, "Failed to generate token")
	}

	log.Printf("User logged in successfully with ID: %d", user.ID)

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
	log.Printf("Request to retrieve user info for ID: %d", userID)

	user, err := controller.UserService.GetById(userID)
	if err != nil {
		log.Printf("Failed to retrieve user: %v", err)
		return err
	}

	log.Printf("User info retrieved successfully for ID: %d", userID)
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

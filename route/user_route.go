package route

import (
	"hotel_ip-p2/controller"
	"hotel_ip-p2/middleware"

	"github.com/labstack/echo/v4"
)

func UserRoutes(e *echo.Group, userController *controller.UserController, topupController *controller.TopupController) {
	users := e.Group("/users")
	users.POST("/register", userController.Register)
	users.POST("/login", userController.Login)
	users.GET("/me", userController.GetMe, middleware.AuthMiddleware)
	users.POST("/topup", topupController.TopupWebhook)
}

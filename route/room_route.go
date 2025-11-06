package route

import (
	"hotel_ip-p2/controller"
	"hotel_ip-p2/middleware"

	"github.com/labstack/echo/v4"
)

func RoomRoutes(e *echo.Group, roomController *controller.RoomController) {
	rooms := e.Group("/rooms")

	rooms.POST("", roomController.Create, middleware.AuthMiddleware)
	rooms.GET("", roomController.FindAll, middleware.AuthMiddleware)
	rooms.GET("/:id", roomController.FindById, middleware.AuthMiddleware)
	rooms.PUT("/:id", roomController.Update, middleware.AuthMiddleware)
	rooms.DELETE("/:id", roomController.Delete, middleware.AuthMiddleware)
}

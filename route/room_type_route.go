package route

import (
	"hotel_ip-p2/controller"
	"hotel_ip-p2/middleware"

	"github.com/labstack/echo/v4"
)

func RoomTypeRoutes(e *echo.Group, roomTypeController *controller.RoomTypeController) {
	roomTypes := e.Group("/room-types")

	roomTypes.POST("", roomTypeController.Create, middleware.AuthMiddleware)
	roomTypes.GET("", roomTypeController.FindAll, middleware.AuthMiddleware)
	roomTypes.GET("/:id", roomTypeController.FindById, middleware.AuthMiddleware)
	roomTypes.PUT("/:id", roomTypeController.Update, middleware.AuthMiddleware)
	roomTypes.DELETE("/:id", roomTypeController.Delete, middleware.AuthMiddleware)
}

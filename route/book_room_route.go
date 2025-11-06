package route

import (
	"hotel_ip-p2/controller"
	"hotel_ip-p2/middleware"

	"github.com/labstack/echo/v4"
)

func BookRoomRoutes(e *echo.Group, bookRoomController *controller.BookRoomController) {
	bookRooms := e.Group("/book-rooms")

	bookRooms.POST("", bookRoomController.Create, middleware.AuthMiddleware)
	bookRooms.GET("/my-bookings", bookRoomController.FindByUserId, middleware.AuthMiddleware)
}

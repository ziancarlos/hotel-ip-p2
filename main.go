package main

import (
	"hotel_ip-p2/controller"
	"hotel_ip-p2/helper"
	"hotel_ip-p2/middleware"
	"hotel_ip-p2/repository"
	"hotel_ip-p2/route"
	"hotel_ip-p2/service"
	"log"

	_ "hotel_ip-p2/docs"

	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Hotel Booking API
// @version 1.0
// @description This is a hotel booking management API server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.email support@hotel.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.

func main() {
	helper.InitConfig()
	db := helper.InitDB()

	userRepository := repository.NewUserRepository()
	topupRepository := repository.NewTopupRepository()
	roomTypeRepository := repository.NewRoomTypeRepository()
	roomRepository := repository.NewRoomRepository()
	bookRoomRepository := repository.NewBookRoomRepository()

	userService := service.NewUserService(userRepository, db)
	topupService := service.NewTopupService(topupRepository, userRepository, db)
	roomTypeService := service.NewRoomTypeService(roomTypeRepository, roomRepository, db)
	roomService := service.NewRoomService(roomRepository, roomTypeRepository, db)
	bookRoomService := service.NewBookRoomService(bookRoomRepository, roomRepository, userRepository, db)

	userController := controller.NewUserController(userService)
	topupController := controller.NewTopupController(topupService)
	roomTypeController := controller.NewRoomTypeController(roomTypeService)
	roomController := controller.NewRoomController(roomService)
	bookRoomController := controller.NewBookRoomController(bookRoomService)

	e := echo.New()

	e.Use(echomiddleware.Logger())
	e.Use(echomiddleware.Recover())
	e.Use(echomiddleware.CORS())

	e.Validator = helper.NewValidator()
	e.HTTPErrorHandler = middleware.ErrorHandler

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	api := e.Group("/api")
	route.UserRoutes(api, userController, topupController)
	route.RoomTypeRoutes(api, roomTypeController)
	route.RoomRoutes(api, roomController)
	route.BookRoomRoutes(api, bookRoomController)

	port := ":8080"
	log.Printf("Server starting on port %s", port)
	if err := e.Start(port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}

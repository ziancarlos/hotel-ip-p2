package main

import (
	"hotel_ip-p2/controller"
	"hotel_ip-p2/helper"
	"hotel_ip-p2/middleware"
	"hotel_ip-p2/repository"
	"hotel_ip-p2/route"
	"hotel_ip-p2/service"
	"log"

	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
)

func main() {
	helper.InitConfig()
	db := helper.InitDB()

	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(userRepository, db)
	userController := controller.NewUserController(userService)

	topupRepository := repository.NewTopupRepository()
	topupService := service.NewTopupService(topupRepository, userRepository, db)
	topupController := controller.NewTopupController(topupService)

	e := echo.New()

	e.Use(echomiddleware.Logger())
	e.Use(echomiddleware.Recover())
	e.Use(echomiddleware.CORS())

	e.Validator = helper.NewValidator()
	e.HTTPErrorHandler = middleware.ErrorHandler

	api := e.Group("/api")
	route.UserRoutes(api, userController, topupController)

	port := ":8080"
	log.Printf("Server starting on port %s", port)
	if err := e.Start(port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}

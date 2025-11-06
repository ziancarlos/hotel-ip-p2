package route

import (
	"hotel_ip-p2/controller"

	"github.com/labstack/echo/v4"
)

func TopupRoutes(e *echo.Group, topupController *controller.TopupController) {
	e.POST("/topup/webhook", topupController.TopupWebhook)
}

package controller

import (
	"hotel_ip-p2/exception"
	"hotel_ip-p2/mapper"
	"hotel_ip-p2/model/web"
	"hotel_ip-p2/model/web/request"
	"hotel_ip-p2/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type TopupController struct {
	TopupService service.TopupService
}

func NewTopupController(topupService service.TopupService) *TopupController {
	return &TopupController{
		TopupService: topupService,
	}
}

func (controller *TopupController) TopupWebhook(c echo.Context) error {
	var req request.TopupWebhookRequest

	if err := c.Bind(&req); err != nil {
		return exception.NewCustomError(http.StatusBadRequest, "Invalid request body")
	}
	topupDomain, err := mapper.ToTopupDomain(req)
	if err != nil {
		return exception.NewCustomError(http.StatusBadRequest, "Invalid amount format")
	}

	result, err := controller.TopupService.ProcessWebhook(topupDomain)
	if err != nil {
		return err
	}

	if result.ID == 0 {
		return exception.NewCustomError(http.StatusOK, "Notification ignored - not settlement")
	}

	topupResponse := mapper.ToTopupResponse(result)

	return c.JSON(http.StatusOK, web.WebResponse{
		Message: "Topup processed successfully",
		Data:    topupResponse,
	})
}

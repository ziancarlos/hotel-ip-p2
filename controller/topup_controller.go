package controller

import (
	"fmt"
	"hotel_ip-p2/exception"
	"hotel_ip-p2/helper"
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

// TopupWebhook godoc
// @Summary Process topup webhook
// @Description Process Midtrans payment webhook for user balance topup
// @Tags topup
// @Accept json
// @Produce json
// @Param request body request.TopupWebhookRequest true "Midtrans webhook payload"
// @Success 200 {object} web.WebResponse{data=response.TopupResponse} "Topup processed successfully"
// @Failure 400 {object} web.WebResponse "Invalid request body"
// @Router /users/topup [post]
func (controller *TopupController) TopupWebhook(c echo.Context) error {
	var req request.TopupWebhookRequest

	if err := c.Bind(&req); err != nil {
		fmt.Println("Error binding request:", err)
		return exception.NewCustomError(http.StatusBadRequest, "Invalid request body")
	}

	if !helper.ValidateMidtransSignature(req.OrderID, req.StatusCode, req.GrossAmount, req.SignatureKey) {
		return exception.NewCustomError(http.StatusUnauthorized, "Invalid signature key")
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

package controller

import (
	"hotel_ip-p2/exception"
	"hotel_ip-p2/mapper"
	"hotel_ip-p2/model/web"
	"hotel_ip-p2/model/web/request"
	"hotel_ip-p2/service"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type BookRoomController struct {
	BookRoomService service.BookRoomService
}

func NewBookRoomController(bookRoomService service.BookRoomService) *BookRoomController {
	return &BookRoomController{
		BookRoomService: bookRoomService,
	}
}

// Create godoc
// @Summary Book a room
// @Description Create a new room booking
// @Tags bookings
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body request.BookRoomRequest true "Booking details"
// @Success 201 {object} web.WebResponse{data=response.BookRoomResponse} "Room booked successfully"
// @Failure 400 {object} web.WebResponse "Invalid request body or validation error"
// @Failure 401 {object} web.WebResponse "Unauthorized"
// @Router /book-rooms [post]
func (controller *BookRoomController) Create(c echo.Context) error {
	var req request.BookRoomRequest

	if err := c.Bind(&req); err != nil {
		return exception.NewCustomError(http.StatusBadRequest, "Invalid request body")
	}

	if err := c.Validate(&req); err != nil {
		return exception.NewCustomError(http.StatusBadRequest, err.Error())
	}

	bookingDate, err := time.Parse("2006-01-02", req.Date)
	if err != nil {
		return exception.NewCustomError(http.StatusBadRequest, "Invalid date format, use YYYY-MM-DD")
	}

	today := time.Now().Truncate(24 * time.Hour)
	if bookingDate.Before(today) {
		return exception.NewCustomError(http.StatusBadRequest, "Booking date must be today or in the future")
	}

	userID := c.Get("user_id").(int)

	bookRoomDomain, err := mapper.ToBookRoomDomain(req, userID)
	if err != nil {
		return exception.NewCustomError(http.StatusBadRequest, "Invalid date format")
	}

	result, err := controller.BookRoomService.Create(bookRoomDomain)
	if err != nil {
		return err
	}

	bookRoomResponse := mapper.ToBookRoomResponse(result)

	return c.JSON(http.StatusCreated, web.WebResponse{
		Message: "Room booked successfully",
		Data:    bookRoomResponse,
	})
}

// FindByUserId godoc
// @Summary Get my bookings
// @Description Get all bookings for the authenticated user
// @Tags bookings
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} web.WebResponse{data=[]response.BookRoomResponse} "Bookings retrieved successfully"
// @Failure 401 {object} web.WebResponse "Unauthorized"
// @Router /book-rooms/my-bookings [get]
func (controller *BookRoomController) FindByUserId(c echo.Context) error {
	userID := c.Get("user_id").(int)

	result, err := controller.BookRoomService.FindByUserId(userID)
	if err != nil {
		return err
	}

	bookRoomResponses := mapper.ToBookRoomResponses(result)

	return c.JSON(http.StatusOK, web.WebResponse{
		Message: "Bookings retrieved successfully",
		Data:    bookRoomResponses,
	})
}

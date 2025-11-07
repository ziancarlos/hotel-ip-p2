package controller

import (
	"hotel_ip-p2/exception"
	"hotel_ip-p2/mapper"
	"hotel_ip-p2/model/web"
	"hotel_ip-p2/model/web/request"
	"hotel_ip-p2/service"
	"log"
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
func (controller *BookRoomController) Create(c echo.Context) error {
	log.Println("Request to create new room booking")
	var req request.BookRoomRequest

	if err := c.Bind(&req); err != nil {
		log.Printf("Failed to bind request body: %v", err)
		return exception.NewCustomError(http.StatusBadRequest, "Invalid request body")
	}

	if err := c.Validate(&req); err != nil {
		log.Printf("Validation failed: %v", err)
		return exception.NewCustomError(http.StatusBadRequest, err.Error())
	}

	bookingDate, err := time.Parse("2006-01-02", req.Date)
	if err != nil {
		log.Printf("Invalid date format: %v", err)
		return exception.NewCustomError(http.StatusBadRequest, "Invalid date format, use YYYY-MM-DD")
	}

	today := time.Now().Truncate(24 * time.Hour)
	if bookingDate.Before(today) {
		log.Println("Booking date is in the past")
		return exception.NewCustomError(http.StatusBadRequest, "Booking date must be today or in the future")
	}

	userID := c.Get("user_id").(int)
	log.Printf("Creating booking for user ID: %d", userID)

	bookRoomDomain, err := mapper.ToBookRoomDomain(req, userID)
	if err != nil {
		log.Printf("Failed to map request to domain: %v", err)
		return exception.NewCustomError(http.StatusBadRequest, "Invalid date format")
	}

	result, err := controller.BookRoomService.Create(bookRoomDomain)
	if err != nil {
		log.Printf("Failed to create room booking: %v", err)
		return err
	}

	log.Printf("Room booking created successfully with ID: %d", result.ID)
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
	log.Printf("Request to retrieve bookings for user ID: %d", userID)

	result, err := controller.BookRoomService.FindByUserId(userID)
	if err != nil {
		log.Printf("Failed to retrieve bookings: %v", err)
		return err
	}

	log.Printf("Successfully retrieved %d bookings for user ID: %d", len(result), userID)
	bookRoomResponses := mapper.ToBookRoomResponses(result)

	return c.JSON(http.StatusOK, web.WebResponse{
		Message: "Bookings retrieved successfully",
		Data:    bookRoomResponses,
	})
}

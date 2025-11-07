package controller

import (
	"hotel_ip-p2/exception"
	"hotel_ip-p2/mapper"
	"hotel_ip-p2/model/web"
	"hotel_ip-p2/model/web/request"
	"hotel_ip-p2/service"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type RoomTypeController struct {
	RoomTypeService service.RoomTypeService
}

func NewRoomTypeController(roomTypeService service.RoomTypeService) *RoomTypeController {
	return &RoomTypeController{
		RoomTypeService: roomTypeService,
	}
}

// Create godoc
// @Summary Create a new room type
// @Description Create a new room type
// @Tags room-types
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body request.RoomTypeRequest true "Room type details"
// @Success 201 {object} web.WebResponse{data=response.RoomTypeResponse} "Room type created successfully"
// @Failure 400 {object} web.WebResponse "Invalid request body or validation error"
// @Failure 401 {object} web.WebResponse "Unauthorized"
func (controller *RoomTypeController) Create(c echo.Context) error {
	log.Println("Request to create new room type")
	var req request.RoomTypeRequest

	if err := c.Bind(&req); err != nil {
		log.Printf("Failed to bind request body: %v", err)
		return exception.NewCustomError(http.StatusBadRequest, "Invalid request body")
	}

	if err := c.Validate(&req); err != nil {
		log.Printf("Validation failed: %v", err)
		return exception.NewCustomError(http.StatusBadRequest, err.Error())
	}

	roomTypeDomain := mapper.ToRoomTypeDomain(req)

	result, err := controller.RoomTypeService.Create(roomTypeDomain)
	if err != nil {
		log.Printf("Failed to create room type: %v", err)
		return err
	}

	log.Printf("Room type created successfully with ID: %d", result.ID)
	roomTypeResponse := mapper.ToRoomTypeResponse(result)

	return c.JSON(http.StatusCreated, web.WebResponse{
		Message: "Room type created successfully",
		Data:    roomTypeResponse,
	})
}

// FindAll godoc
// @Summary Get all room types
// @Description Get a list of all room types
// @Tags room-types
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} web.WebResponse{data=[]response.RoomTypeResponse} "Room types retrieved successfully"
// @Failure 401 {object} web.WebResponse "Unauthorized"
// @Router /room-types [get]
func (controller *RoomTypeController) FindAll(c echo.Context) error {
	log.Println("Request to retrieve all room types")
	result, err := controller.RoomTypeService.FindAll()
	if err != nil {
		log.Printf("Failed to retrieve room types: %v", err)
		return err
	}

	log.Printf("Successfully retrieved %d room types", len(result))
	roomTypeResponses := mapper.ToRoomTypeResponses(result)

	return c.JSON(http.StatusOK, web.WebResponse{
		Message: "Room types retrieved successfully",
		Data:    roomTypeResponses,
	})
}

// FindById godoc
// @Summary Get room type by ID
// @Description Get a room type by its ID
// @Tags room-types
// FindById godoc
// @Summary Get room type by ID
// @Description Get a room type by its ID
// @Tags room-types
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Room Type ID"
// @Success 200 {object} web.WebResponse{data=response.RoomTypeResponse} "Room type retrieved successfully"
// @Failure 400 {object} web.WebResponse "Invalid ID"
// @Failure 401 {object} web.WebResponse "Unauthorized"
// @Failure 404 {object} web.WebResponse "Room type not found"
// @Router /room-types/{id} [get]
func (controller *RoomTypeController) FindById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Printf("Invalid room type ID parameter: %v", err)
		return exception.NewCustomError(http.StatusBadRequest, "Invalid ID")
	}

	log.Printf("Request to retrieve room type with ID: %d", id)
	result, err := controller.RoomTypeService.FindById(id)
	if err != nil {
		log.Printf("Failed to retrieve room type: %v", err)
		return err
	}

	log.Printf("Room type retrieved successfully with ID: %d", id)
	roomTypeResponse := mapper.ToRoomTypeResponse(result)

	return c.JSON(http.StatusOK, web.WebResponse{
		Message: "Room type retrieved successfully",
		Data:    roomTypeResponse,
	})
}

// Update godoc
// @Summary Update a room type
// @Description Update an existing room type by ID
// @Tags room-types
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Room Type ID"
// @Param request body request.RoomTypeRequest true "Updated room type details"
// @Success 200 {object} web.WebResponse{data=response.RoomTypeResponse} "Room type updated successfully"
// @Failure 400 {object} web.WebResponse "Invalid ID or request body"
// @Failure 401 {object} web.WebResponse "Unauthorized"
// @Failure 404 {object} web.WebResponse "Room type not found"
// @Router /room-types/{id} [put]
func (controller *RoomTypeController) Update(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Printf("Invalid room type ID parameter: %v", err)
		return exception.NewCustomError(http.StatusBadRequest, "Invalid ID")
	}

	log.Printf("Request to update room type with ID: %d", id)
	var req request.RoomTypeRequest

	if err := c.Bind(&req); err != nil {
		log.Printf("Failed to bind request body: %v", err)
		return exception.NewCustomError(http.StatusBadRequest, "Invalid request body")
	}

	if err := c.Validate(&req); err != nil {
		log.Printf("Validation failed: %v", err)
		return exception.NewCustomError(http.StatusBadRequest, err.Error())
	}

	roomTypeDomain := mapper.ToRoomTypeDomain(req)
	roomTypeDomain.ID = id

	result, err := controller.RoomTypeService.Update(roomTypeDomain)
	if err != nil {
		log.Printf("Failed to update room type: %v", err)
		return err
	}

	log.Printf("Room type updated successfully with ID: %d", id)
	roomTypeResponse := mapper.ToRoomTypeResponse(result)

	return c.JSON(http.StatusOK, web.WebResponse{
		Message: "Room type updated successfully",
		Data:    roomTypeResponse,
	})
}

// Delete godoc
// @Summary Delete a room type
// @Description Delete a room type by ID
// @Tags room-types
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Room Type ID"
// @Success 200 {object} web.WebResponse "Room type deleted successfully"
// @Failure 400 {object} web.WebResponse "Invalid ID"
// @Failure 401 {object} web.WebResponse "Unauthorized"
// @Failure 404 {object} web.WebResponse "Room type not found"
// @Router /room-types/{id} [delete]
func (controller *RoomTypeController) Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Printf("Invalid room type ID parameter: %v", err)
		return exception.NewCustomError(http.StatusBadRequest, "Invalid ID")
	}

	log.Printf("Request to delete room type with ID: %d", id)
	err = controller.RoomTypeService.Delete(id)
	if err != nil {
		log.Printf("Failed to delete room type: %v", err)
		return err
	}

	log.Printf("Room type deleted successfully with ID: %d", id)
	return c.JSON(http.StatusOK, web.WebResponse{
		Message: "Room type deleted successfully",
	})
}

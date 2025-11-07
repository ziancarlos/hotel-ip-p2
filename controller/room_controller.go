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

type RoomController struct {
	RoomService service.RoomService
}

func NewRoomController(roomService service.RoomService) *RoomController {
	return &RoomController{
		RoomService: roomService,
	}
}

// Create godoc
// @Summary Create a new room
// @Description Create a new room
// @Tags rooms
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body request.RoomRequest true "Room details"
// @Success 201 {object} web.WebResponse{data=response.RoomResponse} "Room created successfully"
// @Failure 400 {object} web.WebResponse "Invalid request body or validation error"
// @Failure 401 {object} web.WebResponse "Unauthorized"
func (controller *RoomController) Create(c echo.Context) error {
	log.Println("Request to create new room")
	var req request.RoomRequest

	if err := c.Bind(&req); err != nil {
		log.Printf("Failed to bind request body: %v", err)
		return exception.NewCustomError(http.StatusBadRequest, "Invalid request body")
	}

	if err := c.Validate(&req); err != nil {
		log.Printf("Validation failed: %v", err)
		return exception.NewCustomError(http.StatusBadRequest, err.Error())
	}

	roomDomain := mapper.ToRoomDomain(req)

	result, err := controller.RoomService.Create(roomDomain)
	if err != nil {
		log.Printf("Failed to create room: %v", err)
		return err
	}

	log.Printf("Room created successfully with ID: %d", result.ID)
	roomResponse := mapper.ToRoomResponse(result)

	return c.JSON(http.StatusCreated, web.WebResponse{
		Message: "Room created successfully",
		Data:    roomResponse,
	})
}

// FindAll godoc
// @Summary Get all rooms
// @Description Get a list of all rooms
// @Tags rooms
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} web.WebResponse{data=[]response.RoomResponse} "Rooms retrieved successfully"
// @Failure 401 {object} web.WebResponse "Unauthorized"
// @Router /rooms [get]
func (controller *RoomController) FindAll(c echo.Context) error {
	log.Println("Request to retrieve all rooms")
	result, err := controller.RoomService.FindAll()
	if err != nil {
		log.Printf("Failed to retrieve rooms: %v", err)
		return err
	}

	log.Printf("Successfully retrieved %d rooms", len(result))
	roomResponses := mapper.ToRoomResponses(result)

	return c.JSON(http.StatusOK, web.WebResponse{
		Message: "Rooms retrieved successfully",
		Data:    roomResponses,
	})
}

// FindById godoc
// @Summary Get room by ID
// @Description Get a room by its ID
// @Tags rooms
// FindById godoc
// @Summary Get room by ID
// @Description Get a room by its ID
// @Tags rooms
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Room ID"
// @Success 200 {object} web.WebResponse{data=response.RoomResponse} "Room retrieved successfully"
// @Failure 400 {object} web.WebResponse "Invalid ID"
// @Failure 401 {object} web.WebResponse "Unauthorized"
// @Failure 404 {object} web.WebResponse "Room not found"
// @Router /rooms/{id} [get]
func (controller *RoomController) FindById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Printf("Invalid room ID parameter: %v", err)
		return exception.NewCustomError(http.StatusBadRequest, "Invalid ID")
	}

	log.Printf("Request to retrieve room with ID: %d", id)
	result, err := controller.RoomService.FindById(id)
	if err != nil {
		log.Printf("Failed to retrieve room: %v", err)
		return err
	}

	log.Printf("Room retrieved successfully with ID: %d", id)
	roomResponse := mapper.ToRoomResponse(result)

	return c.JSON(http.StatusOK, web.WebResponse{
		Message: "Room retrieved successfully",
		Data:    roomResponse,
	})
}

// Update godoc
// @Summary Update a room
// @Description Update an existing room by ID
// @Tags rooms
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Room ID"
// @Param request body request.RoomRequest true "Updated room details"
// @Success 200 {object} web.WebResponse{data=response.RoomResponse} "Room updated successfully"
// @Failure 400 {object} web.WebResponse "Invalid ID or request body"
// @Failure 401 {object} web.WebResponse "Unauthorized"
// @Failure 404 {object} web.WebResponse "Room not found"
// @Router /rooms/{id} [put]
func (controller *RoomController) Update(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Printf("Invalid room ID parameter: %v", err)
		return exception.NewCustomError(http.StatusBadRequest, "Invalid ID")
	}

	log.Printf("Request to update room with ID: %d", id)
	var req request.RoomRequest

	if err := c.Bind(&req); err != nil {
		log.Printf("Failed to bind request body: %v", err)
		return exception.NewCustomError(http.StatusBadRequest, "Invalid request body")
	}

	if err := c.Validate(&req); err != nil {
		log.Printf("Validation failed: %v", err)
		return exception.NewCustomError(http.StatusBadRequest, err.Error())
	}

	roomDomain := mapper.ToRoomDomain(req)
	roomDomain.ID = id

	result, err := controller.RoomService.Update(roomDomain)
	if err != nil {
		log.Printf("Failed to update room: %v", err)
		return err
	}

	log.Printf("Room updated successfully with ID: %d", id)
	roomResponse := mapper.ToRoomResponse(result)

	return c.JSON(http.StatusOK, web.WebResponse{
		Message: "Room updated successfully",
		Data:    roomResponse,
	})
}

// Delete godoc
// @Summary Delete a room
// @Description Delete a room by ID
// @Tags rooms
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Room ID"
// @Success 200 {object} web.WebResponse "Room deleted successfully"
// @Failure 400 {object} web.WebResponse "Invalid ID"
// @Failure 401 {object} web.WebResponse "Unauthorized"
// @Failure 404 {object} web.WebResponse "Room not found"
// @Router /rooms/{id} [delete]
func (controller *RoomController) Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Printf("Invalid room ID parameter: %v", err)
		return exception.NewCustomError(http.StatusBadRequest, "Invalid ID")
	}

	log.Printf("Request to delete room with ID: %d", id)
	err = controller.RoomService.Delete(id)
	if err != nil {
		log.Printf("Failed to delete room: %v", err)
		return err
	}

	log.Printf("Room deleted successfully with ID: %d", id)
	return c.JSON(http.StatusOK, web.WebResponse{
		Message: "Room deleted successfully",
	})
}

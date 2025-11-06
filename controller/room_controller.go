package controller

import (
	"hotel_ip-p2/exception"
	"hotel_ip-p2/mapper"
	"hotel_ip-p2/model/web"
	"hotel_ip-p2/model/web/request"
	"hotel_ip-p2/service"
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

func (controller *RoomController) Create(c echo.Context) error {
	var req request.RoomRequest

	if err := c.Bind(&req); err != nil {
		return exception.NewCustomError(http.StatusBadRequest, "Invalid request body")
	}

	if err := c.Validate(&req); err != nil {
		return exception.NewCustomError(http.StatusBadRequest, err.Error())
	}

	roomDomain := mapper.ToRoomDomain(req)

	result, err := controller.RoomService.Create(roomDomain)
	if err != nil {
		return err
	}

	roomResponse := mapper.ToRoomResponse(result)

	return c.JSON(http.StatusCreated, web.WebResponse{
		Message: "Room created successfully",
		Data:    roomResponse,
	})
}

func (controller *RoomController) FindAll(c echo.Context) error {
	result, err := controller.RoomService.FindAll()
	if err != nil {
		return err
	}

	roomResponses := mapper.ToRoomResponses(result)

	return c.JSON(http.StatusOK, web.WebResponse{
		Message: "Rooms retrieved successfully",
		Data:    roomResponses,
	})
}

func (controller *RoomController) FindById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return exception.NewCustomError(http.StatusBadRequest, "Invalid ID")
	}

	result, err := controller.RoomService.FindById(id)
	if err != nil {
		return err
	}

	roomResponse := mapper.ToRoomResponse(result)

	return c.JSON(http.StatusOK, web.WebResponse{
		Message: "Room retrieved successfully",
		Data:    roomResponse,
	})
}

func (controller *RoomController) Update(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return exception.NewCustomError(http.StatusBadRequest, "Invalid ID")
	}

	var req request.RoomRequest

	if err := c.Bind(&req); err != nil {
		return exception.NewCustomError(http.StatusBadRequest, "Invalid request body")
	}

	if err := c.Validate(&req); err != nil {
		return exception.NewCustomError(http.StatusBadRequest, err.Error())
	}

	roomDomain := mapper.ToRoomDomain(req)
	roomDomain.ID = id

	result, err := controller.RoomService.Update(roomDomain)
	if err != nil {
		return err
	}

	roomResponse := mapper.ToRoomResponse(result)

	return c.JSON(http.StatusOK, web.WebResponse{
		Message: "Room updated successfully",
		Data:    roomResponse,
	})
}

func (controller *RoomController) Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return exception.NewCustomError(http.StatusBadRequest, "Invalid ID")
	}

	err = controller.RoomService.Delete(id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, web.WebResponse{
		Message: "Room deleted successfully",
	})
}

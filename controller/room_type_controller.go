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

type RoomTypeController struct {
	RoomTypeService service.RoomTypeService
}

func NewRoomTypeController(roomTypeService service.RoomTypeService) *RoomTypeController {
	return &RoomTypeController{
		RoomTypeService: roomTypeService,
	}
}

func (controller *RoomTypeController) Create(c echo.Context) error {
	var req request.RoomTypeRequest

	if err := c.Bind(&req); err != nil {
		return exception.NewCustomError(http.StatusBadRequest, "Invalid request body")
	}

	if err := c.Validate(&req); err != nil {
		return exception.NewCustomError(http.StatusBadRequest, err.Error())
	}

	roomTypeDomain := mapper.ToRoomTypeDomain(req)

	result, err := controller.RoomTypeService.Create(roomTypeDomain)
	if err != nil {
		return err
	}

	roomTypeResponse := mapper.ToRoomTypeResponse(result)

	return c.JSON(http.StatusCreated, web.WebResponse{
		Message: "Room type created successfully",
		Data:    roomTypeResponse,
	})
}

func (controller *RoomTypeController) FindAll(c echo.Context) error {
	result, err := controller.RoomTypeService.FindAll()
	if err != nil {
		return err
	}

	roomTypeResponses := mapper.ToRoomTypeResponses(result)

	return c.JSON(http.StatusOK, web.WebResponse{
		Message: "Room types retrieved successfully",
		Data:    roomTypeResponses,
	})
}

func (controller *RoomTypeController) FindById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return exception.NewCustomError(http.StatusBadRequest, "Invalid ID")
	}

	result, err := controller.RoomTypeService.FindById(id)
	if err != nil {
		return err
	}

	roomTypeResponse := mapper.ToRoomTypeResponse(result)

	return c.JSON(http.StatusOK, web.WebResponse{
		Message: "Room type retrieved successfully",
		Data:    roomTypeResponse,
	})
}

func (controller *RoomTypeController) Update(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return exception.NewCustomError(http.StatusBadRequest, "Invalid ID")
	}

	var req request.RoomTypeRequest

	if err := c.Bind(&req); err != nil {
		return exception.NewCustomError(http.StatusBadRequest, "Invalid request body")
	}

	if err := c.Validate(&req); err != nil {
		return exception.NewCustomError(http.StatusBadRequest, err.Error())
	}

	roomTypeDomain := mapper.ToRoomTypeDomain(req)
	roomTypeDomain.ID = id

	result, err := controller.RoomTypeService.Update(roomTypeDomain)
	if err != nil {
		return err
	}

	roomTypeResponse := mapper.ToRoomTypeResponse(result)

	return c.JSON(http.StatusOK, web.WebResponse{
		Message: "Room type updated successfully",
		Data:    roomTypeResponse,
	})
}

func (controller *RoomTypeController) Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return exception.NewCustomError(http.StatusBadRequest, "Invalid ID")
	}

	err = controller.RoomTypeService.Delete(id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, web.WebResponse{
		Message: "Room type deleted successfully",
	})
}

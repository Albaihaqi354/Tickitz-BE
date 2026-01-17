package controller

import (
	"net/http"
	"strconv"

	"github.com/Albaihaqi354/Tickitz-BE/internal/dto"
	"github.com/Albaihaqi354/Tickitz-BE/internal/service"
	"github.com/gin-gonic/gin"
)

type OrderController struct {
	orderService *service.OrderService
}

func NewOrderController(orderService *service.OrderService) *OrderController {
	return &OrderController{
		orderService: orderService,
	}
}

func (ctrl OrderController) GetSchedules(c *gin.Context) {
	idParam := c.Param("id")
	movieId, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{
			Msg:     "Invalid movie_id parameter",
			Success: false,
			Error:   err.Error(),
			Data:    []any{},
		})
		return
	}

	date := c.Query("date")
	var showDate *string
	if date != "" {
		showDate = &date
	}

	cityName := c.Query("city")
	var city *string
	if cityName != "" {
		city = &cityName
	}

	data, err := ctrl.orderService.GetSchedules(c.Request.Context(), movieId, showDate, city)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Response{
			Msg:     "Internal Server Error",
			Success: false,
			Error:   err.Error(),
			Data:    []any{},
		})
		return
	}

	c.JSON(http.StatusOK, dto.Response{
		Msg:     "Get Schedules Success",
		Success: true,
		Data:    data,
	})
}

func (ctrl OrderController) GetSeats(c *gin.Context) {
	idParam := c.Param("id")
	scheduleId, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{
			Msg:     "Invalid schedule_id parameter",
			Success: false,
			Error:   err.Error(),
			Data:    []any{},
		})
		return
	}

	data, err := ctrl.orderService.GetSeatsByScheduleID(c.Request.Context(), scheduleId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Response{
			Msg:     "Internal Server Error",
			Success: false,
			Error:   err.Error(),
			Data:    []any{},
		})
		return
	}

	c.JSON(http.StatusOK, dto.Response{
		Msg:     "Get Seats Success",
		Success: true,
		Data:    data,
	})
}

func (ctrl OrderController) CreateOrder(c *gin.Context) {
	userId, exist := c.Get("user_id")
	if !exist {
		c.JSON(http.StatusUnauthorized, dto.Response{
			Msg:     "Unauthorized",
			Success: false,
			Error:   "User Id Not Found",
			Data:    nil,
		})
		return
	}

	userIdInt, ok := userId.(int)
	if !ok {
		c.JSON(http.StatusInternalServerError, dto.Response{
			Msg:     "Internal Server Error",
			Success: false,
			Error:   "Invalid User Id",
			Data:    nil,
		})
		return
	}

	var req dto.CreateOrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{
			Msg:     "Bad Request",
			Success: false,
			Error:   err.Error(),
			Data:    nil,
		})
		return
	}

	id, bookingCode, createdAt, err := ctrl.orderService.CreateOrder(c.Request.Context(), userIdInt, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Response{
			Msg:     "Internal Server Error",
			Success: false,
			Error:   err.Error(),
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusCreated, dto.Response{
		Msg:     "Create Order Success",
		Success: true,
		Data: gin.H{
			"id":           id,
			"booking_code": bookingCode,
			"created_at":   createdAt,
		},
	})
}

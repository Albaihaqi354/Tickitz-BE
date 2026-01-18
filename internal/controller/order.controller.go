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

// GetSchedules godoc
// @Summary      Get schedules for a movie
// @Description  Get list of schedules with cinema details filtered by movie ID, date, and city
// @Tags         orders
// @Accept       json
// @Produce      json
// @Param        id    path      int     true  "Movie ID"
// @Param        date  query     string  false "Show Date (YYYY-MM-DD)"
// @Param        city  query     string  false "City Name"
// @Success      200   {object}  dto.Response{data=[]dto.GetSchedules}
// @Failure      400   {object}  dto.Response
// @Failure      500   {object}  dto.Response
// @Router       /orders/schedules/{id} [get]
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

// GetSeats godoc
// @Summary      Get seats for a schedule
// @Description  Get list of seats status (sold/available) for a specific schedule ID
// @Tags         orders
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Schedule ID"
// @Success      200  {object}  dto.Response{data=[]dto.SeatResponse}
// @Failure      400  {object}  dto.Response
// @Failure      500  {object}  dto.Response
// @Router       /orders/seats/{id} [get]
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

// CreateOrder godoc
// @Summary      Create a new order
// @Description  Create a new ticket order for a user (Requires user token)
// @Tags         orders
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        order  body      dto.CreateOrderRequest  true  "Order Body"
// @Success      201    {object}  dto.Response
// @Failure      401    {object}  dto.Response
// @Failure      400    {object}  dto.Response
// @Failure      500    {object}  dto.Response
// @Router       /orders [post]
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

// UpdatePaymentStatus godoc
// @Summary      Update order payment status
// @Description  Update the payment status of an order (Requires admin token)
// @Tags         orders
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id      path      int                     true  "Order ID"
// @Param        status  body      dto.UpdateOrderRequest  true  "Status Body"
// @Success      200     {object}  dto.Response
// @Failure      401     {object}  dto.Response
// @Failure      400     {object}  dto.Response
// @Failure      500     {object}  dto.Response
// @Router       /orders/{id} [patch]
func (ctrl OrderController) UpdatePaymentStatus(c *gin.Context) {
	idParam := c.Param("id")
	orderId, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{
			Msg:     "Invalid order_id parameter",
			Success: false,
			Error:   err.Error(),
			Data:    nil,
		})
		return
	}

	var req dto.UpdateOrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{
			Msg:     "Bad Request",
			Success: false,
			Error:   err.Error(),
			Data:    nil,
		})
		return
	}

	err = ctrl.orderService.UpdatePaymentStatus(c.Request.Context(), orderId, req.PaymentStatus)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Response{
			Msg:     "Internal Server Error",
			Success: false,
			Error:   err.Error(),
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, dto.Response{
		Msg:     "Update Payment Status Success",
		Success: true,
		Data:    nil,
	})
}

package router

import (
	"github.com/Albaihaqi354/Tickitz-BE/internal/controller"
	"github.com/Albaihaqi354/Tickitz-BE/internal/middleware"
	"github.com/Albaihaqi354/Tickitz-BE/internal/repository"
	"github.com/Albaihaqi354/Tickitz-BE/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

func RegisterOrderRouter(app *gin.Engine, db *pgxpool.Pool, rdb *redis.Client) {
	orderRepository := repository.NewOrdersRepository()
	orderService := service.NewOrderService(orderRepository, db)
	orderController := controller.NewOrderController(orderService)

	g := app.Group("/orders")
	{
		g.GET("/schedules/:id", orderController.GetSchedules)
		g.GET("/seats/:id", orderController.GetSeats)

		g.POST("/", middleware.VerifyToken(rdb), middleware.CheckRole("user"), orderController.CreateOrder)
		g.PATCH("/:id", middleware.VerifyToken(rdb), middleware.CheckRole("user"), orderController.UpdatePaymentStatus)
	}
}

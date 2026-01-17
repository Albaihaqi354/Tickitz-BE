package router

import (
	"github.com/Albaihaqi354/Tickitz-BE/internal/controller"
	"github.com/Albaihaqi354/Tickitz-BE/internal/middleware"
	"github.com/Albaihaqi354/Tickitz-BE/internal/repository"
	"github.com/Albaihaqi354/Tickitz-BE/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func RegisterOrderRouter(app *gin.Engine, db *pgxpool.Pool) {
	orderRepository := repository.NewOrdersRepository(db)
	orderService := service.NewOrderService(orderRepository)
	orderController := controller.NewOrderController(orderService)

	g := app.Group("/orders")
	{
		g.GET("/schedules/:id", orderController.GetSchedules)
		g.GET("/seats/:id", orderController.GetSeats)

		g.POST("/", middleware.VerifyToken, middleware.CheckRole("user"), orderController.CreateOrder)
	}
}

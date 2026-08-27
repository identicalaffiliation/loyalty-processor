package httpserver

import (
	"fmt"

	"github.com/identicalaffiliation/loyalty-processor/orders/internal/adapters/rest"
	"github.com/identicalaffiliation/loyalty-processor/orders/internal/config"
	"github.com/identicalaffiliation/loyalty-processor/orders/internal/ports"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(
	cfg *config.ServerConfig,
	create ports.CreateOrderUsecase,
	get ports.GetOrdersUsecase,
) *echo.Echo {
	server := echo.New()
	server.Server.Addr = fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	server.Server.ReadTimeout = cfg.ReadTimeout
	server.Server.WriteTimeout = cfg.WriteTimeout
	server.Server.IdleTimeout = cfg.IdleTimeout

	server.POST("/api/v1/orders/:productId", rest.CreateOrder(create))
	server.GET("/api/v1/orders", rest.GetUserOrders(get))

	return server
}

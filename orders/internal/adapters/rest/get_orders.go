package rest

import (
	"errors"
	"net/http"

	"github.com/google/uuid"
	"github.com/identicalaffiliation/loyalty-processor/orders/internal/domain"
	"github.com/identicalaffiliation/loyalty-processor/orders/internal/ports"
	"github.com/labstack/echo/v4"
)

func GetUserOrders(usecase ports.GetOrdersUsecase) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		cookie, err := ctx.Cookie(CookieUserID)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "invalid cookie")
		}

		id, err := uuid.Parse(cookie.Value)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "invalid user id")
		}

		reqCtx := ctx.Request().Context()
		response, err := usecase.GetOrdersByUser(reqCtx, id)
		if err != nil {
			if errors.Is(err, domain.ErrInvalidData) {
				return echo.NewHTTPError(http.StatusBadRequest, "invalid data")
			}

			return echo.NewHTTPError(http.StatusInternalServerError, "internal server error")
		}
		
		return ctx.JSON(http.StatusOK, response)
	}
}

package rest

import (
	"errors"
	"net/http"

	"github.com/google/uuid"
	"github.com/identicalaffiliation/loyalty-processor/orders/internal/domain"
	"github.com/identicalaffiliation/loyalty-processor/orders/internal/dto/input"
	"github.com/identicalaffiliation/loyalty-processor/orders/internal/ports"
	"github.com/labstack/echo/v4"
)

func CreateOrder(usecase ports.CreateOrderUsecase) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var in input.CreateOrderRequest
		if err := ctx.Bind(&in); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "invalid req body")
		}

		productID, err := uuid.Parse(ctx.Param(MuxProductID))
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "invalid product id")
		}

		cookie, err := ctx.Cookie(CookieUserID)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "invalid cookie")
		}

		userID, err := uuid.Parse(cookie.Value)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "invalid user id")
		}

		in.ProductID = productID
		in.UserID = userID

		reqCtx := ctx.Request().Context()
		order, err := usecase.CreateOrder(reqCtx, &in)
		if err != nil {
			switch {
			case errors.Is(err, domain.ErrOutOfStock):
				return echo.NewHTTPError(http.StatusConflict, "product out of stock")
			case errors.Is(err, domain.ErrProductNotFound):
				return echo.NewHTTPError(http.StatusNotFound, "product not found")
			case errors.Is(err, domain.ErrInvalidData):
				return echo.NewHTTPError(http.StatusBadRequest, err.Error())
			default:
				return echo.NewHTTPError(http.StatusInternalServerError, "internal server error")
			}
		}

		return ctx.JSON(http.StatusCreated, order)
	}
}

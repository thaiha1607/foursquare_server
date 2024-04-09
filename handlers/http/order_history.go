package http

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/thaiha1607/foursquare_server/interfaces"
)

type OrderHistoryHandler struct {
	Service interfaces.OrderHistoryService
}

func NewOrderHistoryHandler(e *echo.Echo, srvc interfaces.OrderHistoryService) error {
	handler := &OrderHistoryHandler{
		Service: srvc,
	}
	e.GET("/order-history/:id", handler.GetByID)
	return nil
}

func (h *OrderHistoryHandler) GetByID(c echo.Context) error {
	ctx := context.Background()
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	order_history, err := h.Service.GetByID(ctx, id)
	if err != nil {
		err_rsp := handleError(err)
		return c.JSON(err_rsp.HttpStatusCode, err_rsp)
	}
	return c.JSON(http.StatusOK, order_history)
}

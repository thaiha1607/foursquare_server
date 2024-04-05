package http

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/thaiha1607/foursquare_server/ent"
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
	e.POST("/order-history", handler.Store)
	e.DELETE("/order-history/:id", handler.Delete)
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

func (h *OrderHistoryHandler) Store(c echo.Context) error {
	ctx := context.Background()
	var order_history *ent.OrderHistory
	if err := c.Bind(&order_history); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	resp_obj, err := h.Service.Store(ctx, order_history)
	if err != nil {
		err_rsp := handleError(err)
		return c.JSON(err_rsp.HttpStatusCode, err_rsp)
	}
	return c.JSON(http.StatusCreated, resp_obj)
}

func (h *OrderHistoryHandler) Delete(c echo.Context) error {
	ctx := context.Background()
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	if err := h.Service.Delete(ctx, id); err != nil {
		err_rsp := handleError(err)
		return c.JSON(err_rsp.HttpStatusCode, err_rsp)
	}
	return c.NoContent(http.StatusNoContent)
}

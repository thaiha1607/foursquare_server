package http

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/thaiha1607/foursquare_server/ent"
	"github.com/thaiha1607/foursquare_server/interfaces"
)

type OrderHandler struct {
	Service interfaces.OrderService
}

func NewOrderHandler(e *echo.Echo, srvc interfaces.OrderService) any {
	handler := &OrderHandler{
		Service: srvc,
	}
	e.GET("/orders", handler.Fetch)
	e.GET("/orders/:id", handler.GetByID)
	e.POST("/orders", handler.Store)
	e.PUT("/orders/:id", handler.Update)
	e.DELETE("/orders/:id", handler.Delete)
	return nil
}

func (h *OrderHandler) Fetch(c echo.Context) error {
	ctx := context.Background()
	orders, err := h.Service.Fetch(ctx)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}
	return c.JSON(http.StatusOK, orders)
}

func (h *OrderHandler) GetByID(c echo.Context) error {
	ctx := context.Background()
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "Bad Request")
	}
	order, err := h.Service.GetByID(ctx, id)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}
	return c.JSON(http.StatusOK, order)
}

func (h *OrderHandler) Store(c echo.Context) error {
	ctx := context.Background()
	var order ent.Order
	if err := c.Bind(&order); err != nil {
		return c.String(http.StatusBadRequest, "Bad Request")
	}
	if err := h.Service.Store(ctx, &order); err != nil {
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}
	return c.JSON(http.StatusOK, order)
}

func (h *OrderHandler) Update(c echo.Context) error {
	ctx := context.Background()
	var order ent.Order
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "Bad Request")
	}
	if err := c.Bind(&order); err != nil {
		return c.String(http.StatusBadRequest, "Bad Request")
	}
	if err := h.Service.Update(ctx, id, &order); err != nil {
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}
	return c.JSON(http.StatusOK, order)
}

func (h *OrderHandler) Delete(c echo.Context) error {
	ctx := context.Background()
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "Bad Request")
	}
	if err := h.Service.Delete(ctx, id); err != nil {
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}
	return c.String(http.StatusOK, "OK")
}

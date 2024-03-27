package http

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/thaiha1607/foursquare_server/ent"
	"github.com/thaiha1607/foursquare_server/interfaces"
)

type OrderItemHandler struct {
	Service interfaces.OrderItemService
}

func NewOrderItemHandler(e *echo.Echo, srvc interfaces.OrderItemService) error {
	handler := &OrderItemHandler{
		Service: srvc,
	}
	e.GET("/order-line-items", handler.Fetch)
	e.GET("/order-line-items/:id", handler.GetByID)
	e.POST("/order-line-items", handler.Store)
	e.PUT("/order-line-items/:id", handler.Update)
	e.DELETE("/order-line-items/:id", handler.Delete)
	return nil
}

func (h *OrderItemHandler) Fetch(c echo.Context) error {
	ctx := context.Background()
	orderItems, err := h.Service.Fetch(ctx)
	if err != nil {
		err_rsp := handleError(err)
		return c.JSON(err_rsp.HttpStatusCode, err_rsp)
	}
	return c.JSON(http.StatusOK, orderItems)
}

func (h *OrderItemHandler) GetByID(c echo.Context) error {
	ctx := context.Background()
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	orderItem, err := h.Service.GetByID(ctx, id)
	if err != nil {
		err_rsp := handleError(err)
		return c.JSON(err_rsp.HttpStatusCode, err_rsp)
	}
	return c.JSON(http.StatusOK, orderItem)
}

func (h *OrderItemHandler) Store(c echo.Context) error {
	ctx := context.Background()
	var order_line_item *ent.OrderItem
	if err := c.Bind(&order_line_item); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	resp_obj, err := h.Service.Store(ctx, order_line_item)
	if err != nil {
		err_rsp := handleError(err)
		return c.JSON(err_rsp.HttpStatusCode, err_rsp)
	}
	return c.JSON(http.StatusCreated, resp_obj)
}

func (h *OrderItemHandler) Update(c echo.Context) error {
	ctx := context.Background()
	var order_line_item *ent.OrderItem
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	if err := c.Bind(&order_line_item); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	resp_obj, err := h.Service.Update(ctx, id, order_line_item)
	if err != nil {
		err_rsp := handleError(err)
		return c.JSON(err_rsp.HttpStatusCode, err_rsp)
	}
	return c.JSON(http.StatusOK, resp_obj)
}

func (h *OrderItemHandler) Delete(c echo.Context) error {
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

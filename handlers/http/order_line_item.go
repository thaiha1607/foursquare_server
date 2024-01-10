package http

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/thaiha1607/foursquare_server/ent"
	"github.com/thaiha1607/foursquare_server/interfaces"
)

type OrderLineItemHandler struct {
	Service interfaces.OrderLineItemService
}

func NewOrderLineItemHandler(e *echo.Echo, srvc interfaces.OrderLineItemService) any {
	handler := &OrderLineItemHandler{
		Service: srvc,
	}
	e.GET("/order-line-items", handler.Fetch)
	e.GET("/order-line-items/:id", handler.GetByID)
	e.POST("/order-line-items", handler.Store)
	e.PUT("/order-line-items/:id", handler.Update)
	e.DELETE("/order-line-items/:id", handler.Delete)
	return nil
}

func (h *OrderLineItemHandler) Fetch(c echo.Context) error {
	ctx := context.Background()
	orderLineItems, err := h.Service.Fetch(ctx)
	if err != nil {
		err_rsp := handleError(err)
		return c.JSON(err_rsp.HttpStatusCode, err_rsp)
	}
	return c.JSON(http.StatusOK, orderLineItems)
}

func (h *OrderLineItemHandler) GetByID(c echo.Context) error {
	ctx := context.Background()
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "Bad Request")
	}
	orderLineItem, err := h.Service.GetByID(ctx, id)
	if err != nil {
		err_rsp := handleError(err)
		return c.JSON(err_rsp.HttpStatusCode, err_rsp)
	}
	return c.JSON(http.StatusOK, orderLineItem)
}

func (h *OrderLineItemHandler) Store(c echo.Context) error {
	ctx := context.Background()
	var order_line_item ent.OrderLineItem
	if err := c.Bind(&order_line_item); err != nil {
		return c.String(http.StatusBadRequest, "Bad Request")
	}
	if err := h.Service.Store(ctx, &order_line_item); err != nil {
		err_rsp := handleError(err)
		return c.JSON(err_rsp.HttpStatusCode, err_rsp)
	}
	return c.JSON(http.StatusOK, order_line_item)
}

func (h *OrderLineItemHandler) Update(c echo.Context) error {
	ctx := context.Background()
	var order_line_item ent.OrderLineItem
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "Bad Request")
	}
	if err := c.Bind(&order_line_item); err != nil {
		return c.String(http.StatusBadRequest, "Bad Request")
	}
	if err := h.Service.Update(ctx, id, &order_line_item); err != nil {
		err_rsp := handleError(err)
		return c.JSON(err_rsp.HttpStatusCode, err_rsp)
	}
	return c.JSON(http.StatusOK, order_line_item)
}

func (h *OrderLineItemHandler) Delete(c echo.Context) error {
	ctx := context.Background()
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "Bad Request")
	}
	if err := h.Service.Delete(ctx, id); err != nil {
		err_rsp := handleError(err)
		return c.JSON(err_rsp.HttpStatusCode, err_rsp)
	}
	return c.String(http.StatusOK, "OK")
}

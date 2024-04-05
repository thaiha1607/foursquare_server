package http

import (
	"context"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/thaiha1607/foursquare_server/ent"
	"github.com/thaiha1607/foursquare_server/interfaces"
)

type OrderStatusCodeHandler struct {
	Service interfaces.OrderStatusCodeService
}

func NewOrderStatusCodeHandler(e *echo.Echo, srvc interfaces.OrderStatusCodeService) error {
	handler := &OrderStatusCodeHandler{
		Service: srvc,
	}
	e.GET("/order-status-codes", handler.Fetch)
	e.GET("/order-status-codes/:id", handler.GetByID)
	e.POST("/order-status-codes", handler.Store)
	e.PUT("/order-status-codes/:id", handler.Update)
	e.DELETE("/order-status-codes/:id", handler.Delete)
	return nil
}

func (h *OrderStatusCodeHandler) Fetch(c echo.Context) error {
	ctx := context.Background()
	order_status_codes, err := h.Service.Fetch(ctx)
	if err != nil {
		err_rsp := handleError(err)
		return c.JSON(err_rsp.HttpStatusCode, err_rsp)
	}
	return c.JSON(http.StatusOK, order_status_codes)
}

func (h *OrderStatusCodeHandler) GetByID(c echo.Context) error {
	ctx := context.Background()
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	order_status_code, err := h.Service.GetByID(ctx, id)
	if err != nil {
		err_rsp := handleError(err)
		return c.JSON(err_rsp.HttpStatusCode, err_rsp)
	}
	return c.JSON(http.StatusOK, order_status_code)
}

func (h *OrderStatusCodeHandler) Store(c echo.Context) error {
	ctx := context.Background()
	var order_status_code *ent.OrderStatusCode
	if err := c.Bind(&order_status_code); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	resp_obj, err := h.Service.Store(ctx, order_status_code)
	if err != nil {
		err_rsp := handleError(err)
		return c.JSON(err_rsp.HttpStatusCode, err_rsp)
	}
	return c.JSON(http.StatusCreated, resp_obj)
}

func (h *OrderStatusCodeHandler) Update(c echo.Context) error {
	ctx := context.Background()
	var order_status_code *ent.OrderStatusCode
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	if err := c.Bind(&order_status_code); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	resp_obj, err := h.Service.Update(ctx, id, order_status_code)
	if err != nil {
		err_rsp := handleError(err)
		return c.JSON(err_rsp.HttpStatusCode, err_rsp)
	}
	return c.JSON(http.StatusOK, resp_obj)
}

func (h *OrderStatusCodeHandler) Delete(c echo.Context) error {
	ctx := context.Background()
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	if err := h.Service.Delete(ctx, id); err != nil {
		err_rsp := handleError(err)
		return c.JSON(err_rsp.HttpStatusCode, err_rsp)
	}
	return c.NoContent(http.StatusNoContent)
}

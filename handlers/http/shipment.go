package http

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/thaiha1607/foursquare_server/ent"
	"github.com/thaiha1607/foursquare_server/interfaces"
)

type ShipmentHandler struct {
	Service interfaces.ShipmentService
}

func NewShipmentHandler(e *echo.Echo, srvc interfaces.ShipmentService) error {
	handler := &ShipmentHandler{
		Service: srvc,
	}
	e.GET("/shipments", handler.Fetch)
	e.GET("/shipments/:id", handler.GetByID)
	e.POST("/shipments", handler.Store)
	e.PUT("/shipments/:id", handler.Update)
	e.DELETE("/shipments/:id", handler.Delete)
	return nil
}

func (h *ShipmentHandler) Fetch(c echo.Context) error {
	ctx := context.Background()
	shipments, err := h.Service.Fetch(ctx)
	if err != nil {
		err_rsp := handleError(err)
		return c.JSON(err_rsp.HttpStatusCode, err_rsp)
	}
	return c.JSON(http.StatusOK, shipments)
}

func (h *ShipmentHandler) GetByID(c echo.Context) error {
	ctx := context.Background()
	id := c.Param("id")
	shipment, err := h.Service.GetByID(ctx, id)
	if err != nil {
		err_rsp := handleError(err)
		return c.JSON(err_rsp.HttpStatusCode, err_rsp)
	}
	return c.JSON(http.StatusOK, shipment)
}

func (h *ShipmentHandler) Store(c echo.Context) error {
	ctx := context.Background()
	var shipment *ent.Shipment
	if err := c.Bind(&shipment); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	resp_obj, err := h.Service.Store(ctx, shipment)
	if err != nil {
		err_rsp := handleError(err)
		return c.JSON(err_rsp.HttpStatusCode, err_rsp)
	}
	return c.JSON(http.StatusCreated, resp_obj)
}

func (h *ShipmentHandler) Update(c echo.Context) error {
	ctx := context.Background()
	var shipment *ent.Shipment
	id := c.Param("id")
	if err := c.Bind(&shipment); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	resp_obj, err := h.Service.Update(ctx, id, shipment)
	if err != nil {
		err_rsp := handleError(err)
		return c.JSON(err_rsp.HttpStatusCode, err_rsp)
	}
	return c.JSON(http.StatusOK, resp_obj)
}

func (h *ShipmentHandler) Delete(c echo.Context) error {
	ctx := context.Background()
	id := c.Param("id")
	if err := h.Service.Delete(ctx, id); err != nil {
		err_rsp := handleError(err)
		return c.JSON(err_rsp.HttpStatusCode, err_rsp)
	}
	return c.NoContent(http.StatusNoContent)
}

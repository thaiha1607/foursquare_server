package http

import (
	"context"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/thaiha1607/foursquare_server/ent"
	"github.com/thaiha1607/foursquare_server/interfaces"
)

type ShipmentStatusCodeHandler struct {
	Service interfaces.ShipmentStatusCodeService
}

func NewShipmentStatusCodeHandler(e *echo.Echo, srvc interfaces.ShipmentStatusCodeService) error {
	handler := &ShipmentStatusCodeHandler{
		Service: srvc,
	}
	e.GET("/shipment-status-codes", handler.Fetch)
	e.GET("/shipment-status-codes/:id", handler.GetByID)
	e.POST("/shipment-status-codes", handler.Store)
	e.PUT("/shipment-status-codes/:id", handler.Update)
	e.DELETE("/shipment-status-codes/:id", handler.Delete)
	return nil
}

func (h *ShipmentStatusCodeHandler) Fetch(c echo.Context) error {
	ctx := context.Background()
	shipment_status_codes, err := h.Service.Fetch(ctx)
	if err != nil {
		err_rsp := handleError(err)
		return c.JSON(err_rsp.HttpStatusCode, err_rsp)
	}
	return c.JSON(http.StatusOK, shipment_status_codes)
}

func (h *ShipmentStatusCodeHandler) GetByID(c echo.Context) error {
	ctx := context.Background()
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	shipment_status_code, err := h.Service.GetByID(ctx, id)
	if err != nil {
		err_rsp := handleError(err)
		return c.JSON(err_rsp.HttpStatusCode, err_rsp)
	}
	return c.JSON(http.StatusOK, shipment_status_code)
}

func (h *ShipmentStatusCodeHandler) Store(c echo.Context) error {
	ctx := context.Background()
	var shipment_status_code *ent.ShipmentStatusCode
	if err := c.Bind(&shipment_status_code); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	resp_obj, err := h.Service.Store(ctx, shipment_status_code)
	if err != nil {
		err_rsp := handleError(err)
		return c.JSON(err_rsp.HttpStatusCode, err_rsp)
	}
	return c.JSON(http.StatusCreated, resp_obj)
}

func (h *ShipmentStatusCodeHandler) Update(c echo.Context) error {
	ctx := context.Background()
	var shipment_status_code *ent.ShipmentStatusCode
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	if err := c.Bind(&shipment_status_code); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	resp_obj, err := h.Service.Update(ctx, id, shipment_status_code)
	if err != nil {
		err_rsp := handleError(err)
		return c.JSON(err_rsp.HttpStatusCode, err_rsp)
	}
	return c.JSON(http.StatusOK, resp_obj)
}

func (h *ShipmentStatusCodeHandler) Delete(c echo.Context) error {
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

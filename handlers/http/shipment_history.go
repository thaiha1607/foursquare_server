package http

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/thaiha1607/foursquare_server/ent"
	"github.com/thaiha1607/foursquare_server/interfaces"
)

type ShipmentHistoryHandler struct {
	Service interfaces.ShipmentHistoryService
}

func NewShipmentHistoryHandler(e *echo.Echo, srvc interfaces.ShipmentHistoryService) error {
	handler := &ShipmentHistoryHandler{
		Service: srvc,
	}
	e.GET("/shipment-history/:id", handler.GetByID)
	e.POST("/shipment-history", handler.Store)
	e.DELETE("/shipment-history/:id", handler.Delete)
	return nil
}

func (h *ShipmentHistoryHandler) GetByID(c echo.Context) error {
	ctx := context.Background()
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	shipment_history, err := h.Service.GetByID(ctx, id)
	if err != nil {
		err_rsp := handleError(err)
		return c.JSON(err_rsp.HttpStatusCode, err_rsp)
	}
	return c.JSON(http.StatusOK, shipment_history)
}

func (h *ShipmentHistoryHandler) Store(c echo.Context) error {
	ctx := context.Background()
	var shipment_history *ent.ShipmentHistory
	if err := c.Bind(&shipment_history); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	resp_obj, err := h.Service.Store(ctx, shipment_history)
	if err != nil {
		err_rsp := handleError(err)
		return c.JSON(err_rsp.HttpStatusCode, err_rsp)
	}
	return c.JSON(http.StatusCreated, resp_obj)
}

func (h *ShipmentHistoryHandler) Delete(c echo.Context) error {
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

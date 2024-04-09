package http

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
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

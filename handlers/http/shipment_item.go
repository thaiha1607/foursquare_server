package http

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/thaiha1607/foursquare_server/ent"
	"github.com/thaiha1607/foursquare_server/interfaces"
)

type ShipmentItemHandler struct {
	Service interfaces.ShipmentItemService
}

func NewShipmentItemHandler(e *echo.Echo, srvc interfaces.ShipmentItemService) error {
	handler := &ShipmentItemHandler{
		Service: srvc,
	}
	e.GET("/shipment-items/:id", handler.GetByID)
	e.POST("/shipment-items", handler.Store)
	e.DELETE("/shipment-items/:id", handler.Delete)
	return nil
}

func (h *ShipmentItemHandler) GetByID(c echo.Context) error {
	ctx := context.Background()
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	shipment_item, err := h.Service.GetByID(ctx, id)
	if err != nil {
		err_rsp := handleError(err)
		return c.JSON(err_rsp.HttpStatusCode, err_rsp)
	}
	return c.JSON(http.StatusOK, shipment_item)
}

func (h *ShipmentItemHandler) Store(c echo.Context) error {
	ctx := context.Background()
	var shipment_item *ent.ShipmentItem
	if err := c.Bind(&shipment_item); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	resp_obj, err := h.Service.Store(ctx, shipment_item)
	if err != nil {
		err_rsp := handleError(err)
		return c.JSON(err_rsp.HttpStatusCode, err_rsp)
	}
	return c.JSON(http.StatusCreated, resp_obj)
}

func (h *ShipmentItemHandler) Delete(c echo.Context) error {
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

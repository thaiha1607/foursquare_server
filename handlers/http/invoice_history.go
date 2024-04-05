package http

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/thaiha1607/foursquare_server/ent"
	"github.com/thaiha1607/foursquare_server/interfaces"
)

type InvoiceHistoryHandler struct {
	Service interfaces.InvoiceHistoryService
}

func NewInvoiceHistoryHandler(e *echo.Echo, srvc interfaces.InvoiceHistoryService) error {
	handler := &InvoiceHistoryHandler{
		Service: srvc,
	}
	e.GET("/invoice-history/:id", handler.GetByID)
	e.POST("/invoice-history", handler.Store)
	e.DELETE("/invoice-history/:id", handler.Delete)
	return nil
}

func (h *InvoiceHistoryHandler) GetByID(c echo.Context) error {
	ctx := context.Background()
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	invoice_history, err := h.Service.GetByID(ctx, id)
	if err != nil {
		err_rsp := handleError(err)
		return c.JSON(err_rsp.HttpStatusCode, err_rsp)
	}
	return c.JSON(http.StatusOK, invoice_history)
}

func (h *InvoiceHistoryHandler) Store(c echo.Context) error {
	ctx := context.Background()
	var invoice_history *ent.InvoiceHistory
	if err := c.Bind(&invoice_history); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	resp_obj, err := h.Service.Store(ctx, invoice_history)
	if err != nil {
		err_rsp := handleError(err)
		return c.JSON(err_rsp.HttpStatusCode, err_rsp)
	}
	return c.JSON(http.StatusCreated, resp_obj)
}

func (h *InvoiceHistoryHandler) Delete(c echo.Context) error {
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

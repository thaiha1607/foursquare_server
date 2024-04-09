package http

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
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

package http

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/thaiha1607/foursquare_server/ent"
	"github.com/thaiha1607/foursquare_server/interfaces"
)

type InvoiceLineItemHandler struct {
	Service interfaces.InvoiceLineItemService
}

func NewInvoiceLineItemHandler(e *echo.Echo, srvc interfaces.InvoiceLineItemService) error {
	handler := &InvoiceLineItemHandler{
		Service: srvc,
	}
	e.GET("/invoice-line-items", handler.Fetch)
	e.GET("/invoice-line-items/:id", handler.GetByID)
	e.POST("/invoice-line-items", handler.Store)
	e.DELETE("/invoice-line-items/:id", handler.Delete)
	return nil
}

func (h *InvoiceLineItemHandler) Fetch(c echo.Context) error {
	ctx := context.Background()
	invoiceLineItems, err := h.Service.Fetch(ctx)
	if err != nil {
		err_rsp := handleError(err)
		return c.JSON(err_rsp.HttpStatusCode, err_rsp)
	}
	return c.JSON(http.StatusOK, invoiceLineItems)
}

func (h *InvoiceLineItemHandler) GetByID(c echo.Context) error {
	ctx := context.Background()
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	invoiceLineItem, err := h.Service.GetByID(ctx, id)
	if err != nil {
		err_rsp := handleError(err)
		return c.JSON(err_rsp.HttpStatusCode, err_rsp)
	}
	return c.JSON(http.StatusOK, invoiceLineItem)
}

func (h *InvoiceLineItemHandler) Store(c echo.Context) error {
	ctx := context.Background()
	var invoice_line_item *ent.InvoiceLineItem
	if err := c.Bind(&invoice_line_item); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	resp_obj, err := h.Service.Store(ctx, invoice_line_item)
	if err != nil {
		err_rsp := handleError(err)
		return c.JSON(err_rsp.HttpStatusCode, err_rsp)
	}
	return c.JSON(http.StatusCreated, resp_obj)
}

func (h *InvoiceLineItemHandler) Delete(c echo.Context) error {
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

package http

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/thaiha1607/foursquare_server/ent"
	"github.com/thaiha1607/foursquare_server/interfaces"
)

type InvoiceHandler struct {
	Service interfaces.InvoiceService
}

func NewInvoiceHandler(e *echo.Echo, srvc interfaces.InvoiceService) error {
	handler := &InvoiceHandler{
		Service: srvc,
	}
	e.GET("/invoices", handler.Fetch)
	e.GET("/invoices/:id", handler.GetByID)
	e.POST("/invoices", handler.Store)
	e.PUT("/invoices/:id", handler.Update)
	e.DELETE("/invoices/:id", handler.Delete)
	return nil
}

func (h *InvoiceHandler) Fetch(c echo.Context) error {
	ctx := context.Background()
	invoices, err := h.Service.Fetch(ctx)
	if err != nil {
		err_rsp := handleError(err)
		return c.JSON(err_rsp.HttpStatusCode, err_rsp)
	}
	return c.JSON(http.StatusOK, invoices)
}

func (h *InvoiceHandler) GetByID(c echo.Context) error {
	ctx := context.Background()
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	invoice, err := h.Service.GetByID(ctx, id)
	if err != nil {
		err_rsp := handleError(err)
		return c.JSON(err_rsp.HttpStatusCode, err_rsp)
	}
	return c.JSON(http.StatusOK, invoice)
}

func (h *InvoiceHandler) Store(c echo.Context) error {
	ctx := context.Background()
	var invoice *ent.Invoice
	if err := c.Bind(&invoice); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	resp_obj, err := h.Service.Store(ctx, invoice)
	if err != nil {
		err_rsp := handleError(err)
		return c.JSON(err_rsp.HttpStatusCode, err_rsp)
	}
	return c.JSON(http.StatusCreated, resp_obj)
}

func (h *InvoiceHandler) Update(c echo.Context) error {
	ctx := context.Background()
	var invoice *ent.Invoice
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	if err := c.Bind(&invoice); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	resp_obj, err := h.Service.Update(ctx, id, invoice)
	if err != nil {
		err_rsp := handleError(err)
		return c.JSON(err_rsp.HttpStatusCode, err_rsp)
	}
	return c.JSON(http.StatusOK, resp_obj)
}

func (h *InvoiceHandler) Delete(c echo.Context) error {
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

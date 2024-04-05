package http

import (
	"context"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/thaiha1607/foursquare_server/ent"
	"github.com/thaiha1607/foursquare_server/interfaces"
)

type InvoiceStatusCodeHandler struct {
	Service interfaces.InvoiceStatusCodeService
}

func NewInvoiceStatusCodeHandler(e *echo.Echo, srvc interfaces.InvoiceStatusCodeService) error {
	handler := &InvoiceStatusCodeHandler{
		Service: srvc,
	}
	e.GET("/invoice-status-codes", handler.Fetch)
	e.GET("/invoice-status-codes/:id", handler.GetByID)
	e.POST("/invoice-status-codes", handler.Store)
	e.PUT("/invoice-status-codes/:id", handler.Update)
	e.DELETE("/invoice-status-codes/:id", handler.Delete)
	return nil
}

func (h *InvoiceStatusCodeHandler) Fetch(c echo.Context) error {
	ctx := context.Background()
	invoice_status_codes, err := h.Service.Fetch(ctx)
	if err != nil {
		err_rsp := handleError(err)
		return c.JSON(err_rsp.HttpStatusCode, err_rsp)
	}
	return c.JSON(http.StatusOK, invoice_status_codes)
}

func (h *InvoiceStatusCodeHandler) GetByID(c echo.Context) error {
	ctx := context.Background()
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	invoice_status_code, err := h.Service.GetByID(ctx, id)
	if err != nil {
		err_rsp := handleError(err)
		return c.JSON(err_rsp.HttpStatusCode, err_rsp)
	}
	return c.JSON(http.StatusOK, invoice_status_code)
}

func (h *InvoiceStatusCodeHandler) Store(c echo.Context) error {
	ctx := context.Background()
	var invoice_status_code *ent.InvoiceStatusCode
	if err := c.Bind(&invoice_status_code); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	resp_obj, err := h.Service.Store(ctx, invoice_status_code)
	if err != nil {
		err_rsp := handleError(err)
		return c.JSON(err_rsp.HttpStatusCode, err_rsp)
	}
	return c.JSON(http.StatusCreated, resp_obj)
}

func (h *InvoiceStatusCodeHandler) Update(c echo.Context) error {
	ctx := context.Background()
	var invoice_status_code *ent.InvoiceStatusCode
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	if err := c.Bind(&invoice_status_code); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	resp_obj, err := h.Service.Update(ctx, id, invoice_status_code)
	if err != nil {
		err_rsp := handleError(err)
		return c.JSON(err_rsp.HttpStatusCode, err_rsp)
	}
	return c.JSON(http.StatusOK, resp_obj)
}

func (h *InvoiceStatusCodeHandler) Delete(c echo.Context) error {
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

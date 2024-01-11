package http

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/thaiha1607/foursquare_server/ent"
	"github.com/thaiha1607/foursquare_server/interfaces"
)

type FinancialTransactionHandler struct {
	Service interfaces.FinancialTransactionService
}

func NewFinancialTransactionHandler(e *echo.Echo, srvc interfaces.FinancialTransactionService) any {
	handler := &FinancialTransactionHandler{
		Service: srvc,
	}
	e.GET("/financial-transactions", handler.Fetch)
	e.GET("/financial-transactions/:id", handler.GetByID)
	e.POST("/financial-transactions", handler.Store)
	e.PUT("/financial-transactions/:id", handler.Update)
	e.DELETE("/financial-transactions/:id", handler.Delete)
	return nil
}

func (h *FinancialTransactionHandler) Fetch(c echo.Context) error {
	ctx := context.Background()
	financialTransactions, err := h.Service.Fetch(ctx)
	if err != nil {
		err_rsp := handleError(err)
		return c.JSON(err_rsp.HttpStatusCode, err_rsp)
	}
	return c.JSON(http.StatusOK, financialTransactions)
}

func (h *FinancialTransactionHandler) GetByID(c echo.Context) error {
	ctx := context.Background()
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	financialTransaction, err := h.Service.GetByID(ctx, id)
	if err != nil {
		err_rsp := handleError(err)
		return c.JSON(err_rsp.HttpStatusCode, err_rsp)
	}
	return c.JSON(http.StatusOK, financialTransaction)
}

func (h *FinancialTransactionHandler) Store(c echo.Context) error {
	ctx := context.Background()
	var financial_transaction *ent.FinancialTransaction
	if err := c.Bind(&financial_transaction); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	resp_obj, err := h.Service.Store(ctx, financial_transaction)
	if err != nil {
		err_rsp := handleError(err)
		return c.JSON(err_rsp.HttpStatusCode, err_rsp)
	}
	return c.JSON(http.StatusCreated, resp_obj)
}

func (h *FinancialTransactionHandler) Update(c echo.Context) error {
	ctx := context.Background()
	var financial_transaction *ent.FinancialTransaction
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	if err := c.Bind(&financial_transaction); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	resp_obj, err := h.Service.Update(ctx, id, financial_transaction)
	if err != nil {
		err_rsp := handleError(err)
		return c.JSON(err_rsp.HttpStatusCode, err_rsp)
	}
	return c.JSON(http.StatusOK, resp_obj)
}

func (h *FinancialTransactionHandler) Delete(c echo.Context) error {
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

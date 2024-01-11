package http

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/thaiha1607/foursquare_server/ent"
	"github.com/thaiha1607/foursquare_server/interfaces"
)

type AccountHandler struct {
	Service interfaces.AccountService
}

func NewAccountHandler(e *echo.Echo, srvc interfaces.AccountService) any {
	handler := &AccountHandler{
		Service: srvc,
	}
	e.GET("/accounts", handler.Fetch)
	e.GET("/accounts/:id", handler.GetByID)
	e.POST("/accounts", handler.Store)
	e.PUT("/accounts/:id", handler.Update)
	e.DELETE("/accounts/:id", handler.Delete)
	return nil
}

func (h *AccountHandler) Fetch(c echo.Context) error {
	ctx := context.Background()
	accounts, err := h.Service.Fetch(ctx)
	if err != nil {
		err_rsp := handleError(err)
		return c.JSON(err_rsp.HttpStatusCode, err_rsp)
	}
	return c.JSON(http.StatusOK, accounts)
}

func (h *AccountHandler) GetByID(c echo.Context) error {
	ctx := context.Background()
	id := c.Param("id")
	account, err := h.Service.GetByID(ctx, id)
	if err != nil {
		err_rsp := handleError(err)
		return c.JSON(err_rsp.HttpStatusCode, err_rsp)
	}
	return c.JSON(http.StatusOK, account)
}

func (h *AccountHandler) Store(c echo.Context) error {
	ctx := context.Background()
	var account *ent.Account
	if err := c.Bind(&account); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	resp_obj, err := h.Service.Store(ctx, account)
	if err != nil {
		err_rsp := handleError(err)
		return c.JSON(err_rsp.HttpStatusCode, err_rsp)
	}
	return c.JSON(http.StatusCreated, resp_obj)
}

func (h *AccountHandler) Update(c echo.Context) error {
	ctx := context.Background()
	var account *ent.Account
	id := c.Param("id")
	if err := c.Bind(&account); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	resp_obj, err := h.Service.Update(ctx, id, account)
	if err != nil {
		err_rsp := handleError(err)
		return c.JSON(err_rsp.HttpStatusCode, err_rsp)
	}
	return c.JSON(http.StatusOK, resp_obj)
}

func (h *AccountHandler) Delete(c echo.Context) error {
	ctx := context.Background()
	id := c.Param("id")
	if err := h.Service.Delete(ctx, id); err != nil {
		err_rsp := handleError(err)
		return c.JSON(err_rsp.HttpStatusCode, err_rsp)
	}
	return c.NoContent(http.StatusNoContent)
}

package http

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/thaiha1607/foursquare_server/ent"
	"github.com/thaiha1607/foursquare_server/interfaces"
)

type PersonAddressHandler struct {
	Service interfaces.PersonAddressService
}

func NewPersonAddressHandler(e *echo.Echo, srvc interfaces.PersonAddressService) error {
	handler := &PersonAddressHandler{
		Service: srvc,
	}
	e.GET("/person-addresses", handler.Fetch)
	e.POST("/person-addresses", handler.Store)
	return nil
}

func (h *PersonAddressHandler) Fetch(c echo.Context) error {
	ctx := context.Background()
	person_addresses, err := h.Service.Fetch(ctx)
	if err != nil {
		err_rsp := handleError(err)
		return c.JSON(err_rsp.HttpStatusCode, err_rsp)
	}
	return c.JSON(http.StatusOK, person_addresses)
}

func (h *PersonAddressHandler) Store(c echo.Context) error {
	ctx := context.Background()
	var person_address *ent.PersonAddress
	if err := c.Bind(&person_address); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	resp_obj, err := h.Service.Store(ctx, person_address)
	if err != nil {
		err_rsp := handleError(err)
		return c.JSON(err_rsp.HttpStatusCode, err_rsp)
	}
	return c.JSON(http.StatusCreated, resp_obj)
}

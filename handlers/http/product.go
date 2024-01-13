package http

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/thaiha1607/foursquare_server/ent"
	"github.com/thaiha1607/foursquare_server/interfaces"
)

type ProductHandler struct {
	Service interfaces.ProductService
}

func NewProductHandler(e *echo.Echo, srvc interfaces.ProductService) error {
	handler := &ProductHandler{
		Service: srvc,
	}
	e.GET("/products", handler.Fetch)
	e.GET("/products/:id", handler.GetByID)
	e.POST("/products", handler.Store)
	e.PUT("/products/:id", handler.Update)
	e.DELETE("/products/:id", handler.Delete)
	return nil
}

func (h *ProductHandler) Fetch(c echo.Context) error {
	ctx := context.Background()
	products, err := h.Service.Fetch(ctx)
	if err != nil {
		err_rsp := handleError(err)
		return c.JSON(err_rsp.HttpStatusCode, err_rsp)
	}
	return c.JSON(http.StatusOK, products)
}

func (h *ProductHandler) GetByID(c echo.Context) error {
	ctx := context.Background()
	product, err := h.Service.GetByID(ctx, c.Param("id"))
	if err != nil {
		err_rsp := handleError(err)
		return c.JSON(err_rsp.HttpStatusCode, err_rsp)
	}
	return c.JSON(http.StatusOK, product)
}

func (h *ProductHandler) Store(c echo.Context) error {
	ctx := context.Background()
	var product ent.Product
	if err := c.Bind(&product); err != nil {
		return c.String(http.StatusBadRequest, "Bad Request")
	}
	if err := h.Service.Store(ctx, &product); err != nil {
		err_rsp := handleError(err)
		return c.JSON(err_rsp.HttpStatusCode, err_rsp)
	}
	return c.JSON(http.StatusOK, product)
}

func (h *ProductHandler) Update(c echo.Context) error {
	ctx := context.Background()
	var product ent.Product
	if err := c.Bind(&product); err != nil {
		return c.String(http.StatusBadRequest, "Bad Request")
	}
	if err := h.Service.Update(ctx, c.Param("id"), &product); err != nil {
		err_rsp := handleError(err)
		return c.JSON(err_rsp.HttpStatusCode, err_rsp)
	}
	return c.JSON(http.StatusOK, product)
}

func (h *ProductHandler) Delete(c echo.Context) error {
	ctx := context.Background()
	if err := h.Service.Delete(ctx, c.Param("id")); err != nil {
		err_rsp := handleError(err)
		return c.JSON(err_rsp.HttpStatusCode, err_rsp)
	}
	return c.String(http.StatusOK, "OK")
}

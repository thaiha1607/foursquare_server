package http

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/thaiha1607/foursquare_server/ent"
	"github.com/thaiha1607/foursquare_server/interfaces"
)

type ProductColorHandler struct {
	Service interfaces.ProductColorService
}

func NewProductColorHandler(e *echo.Echo, srvc interfaces.ProductColorService) error {
	handler := &ProductColorHandler{
		Service: srvc,
	}
	e.GET("/product-colors", handler.Fetch)
	e.GET("/product-colors/:id", handler.GetByID)
	e.POST("/product-colors", handler.Store)
	e.PUT("/product-colors/:id", handler.Update)
	e.DELETE("/product-colors/:id", handler.Delete)
	return nil
}

func (h *ProductColorHandler) Fetch(c echo.Context) error {
	ctx := context.Background()
	product_colors, err := h.Service.Fetch(ctx)
	if err != nil {
		err_rsp := handleError(err)
		return c.JSON(err_rsp.HttpStatusCode, err_rsp)
	}
	return c.JSON(http.StatusOK, product_colors)
}

func (h *ProductColorHandler) GetByID(c echo.Context) error {
	ctx := context.Background()
	id := c.Param("id")
	product_color, err := h.Service.GetByID(ctx, id)
	if err != nil {
		err_rsp := handleError(err)
		return c.JSON(err_rsp.HttpStatusCode, err_rsp)
	}
	return c.JSON(http.StatusOK, product_color)
}

func (h *ProductColorHandler) Store(c echo.Context) error {
	ctx := context.Background()
	var product_color *ent.ProductColor
	if err := c.Bind(&product_color); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	resp_obj, err := h.Service.Store(ctx, product_color)
	if err != nil {
		err_rsp := handleError(err)
		return c.JSON(err_rsp.HttpStatusCode, err_rsp)
	}
	return c.JSON(http.StatusCreated, resp_obj)
}

func (h *ProductColorHandler) Update(c echo.Context) error {
	ctx := context.Background()
	var product_color *ent.ProductColor
	id := c.Param("id")
	if err := c.Bind(&product_color); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	resp_obj, err := h.Service.Update(ctx, id, product_color)
	if err != nil {
		err_rsp := handleError(err)
		return c.JSON(err_rsp.HttpStatusCode, err_rsp)
	}
	return c.JSON(http.StatusOK, resp_obj)
}

func (h *ProductColorHandler) Delete(c echo.Context) error {
	ctx := context.Background()
	id := c.Param("id")
	if err := h.Service.Delete(ctx, id); err != nil {
		err_rsp := handleError(err)
		return c.JSON(err_rsp.HttpStatusCode, err_rsp)
	}
	return c.NoContent(http.StatusNoContent)
}

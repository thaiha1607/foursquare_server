package http

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/thaiha1607/foursquare_server/ent"
	"github.com/thaiha1607/foursquare_server/interfaces"
)

type ProductInfoHandler struct {
	Service interfaces.ProductInfoService
}

func NewProductInfoHandler(e *echo.Echo, srvc interfaces.ProductInfoService) error {
	handler := &ProductInfoHandler{
		Service: srvc,
	}
	e.GET("/product-infos", handler.Fetch)
	e.GET("/product-infos/:id", handler.GetByID)
	e.POST("/product-infos", handler.Store)
	e.PUT("/product-infos/:id", handler.Update)
	e.DELETE("/product-infos/:id", handler.Delete)
	return nil
}

func (h *ProductInfoHandler) Fetch(c echo.Context) error {
	ctx := context.Background()
	product_infos, err := h.Service.Fetch(ctx)
	if err != nil {
		err_rsp := handleError(err)
		return c.JSON(err_rsp.HttpStatusCode, err_rsp)
	}
	return c.JSON(http.StatusOK, product_infos)
}

func (h *ProductInfoHandler) GetByID(c echo.Context) error {
	ctx := context.Background()
	id := c.Param("id")
	product_info, err := h.Service.GetByID(ctx, id)
	if err != nil {
		err_rsp := handleError(err)
		return c.JSON(err_rsp.HttpStatusCode, err_rsp)
	}
	return c.JSON(http.StatusOK, product_info)
}

func (h *ProductInfoHandler) Store(c echo.Context) error {
	ctx := context.Background()
	var product_info *ent.ProductInfo
	if err := c.Bind(&product_info); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	resp_obj, err := h.Service.Store(ctx, product_info)
	if err != nil {
		err_rsp := handleError(err)
		return c.JSON(err_rsp.HttpStatusCode, err_rsp)
	}
	return c.JSON(http.StatusCreated, resp_obj)
}

func (h *ProductInfoHandler) Update(c echo.Context) error {
	ctx := context.Background()
	var product_info *ent.ProductInfo
	id := c.Param("id")
	if err := c.Bind(&product_info); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	resp_obj, err := h.Service.Update(ctx, id, product_info)
	if err != nil {
		err_rsp := handleError(err)
		return c.JSON(err_rsp.HttpStatusCode, err_rsp)
	}
	return c.JSON(http.StatusOK, resp_obj)
}

func (h *ProductInfoHandler) Delete(c echo.Context) error {
	ctx := context.Background()
	id := c.Param("id")
	if err := h.Service.Delete(ctx, id); err != nil {
		err_rsp := handleError(err)
		return c.JSON(err_rsp.HttpStatusCode, err_rsp)
	}
	return c.NoContent(http.StatusNoContent)
}

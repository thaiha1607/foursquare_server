package http

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/thaiha1607/foursquare_server/ent"
	"github.com/thaiha1607/foursquare_server/interfaces"
)

type ProductImageHandler struct {
	Service interfaces.ProductImageService
}

func NewProductImageHandler(e *echo.Echo, srvc interfaces.ProductImageService) error {
	handler := &ProductImageHandler{
		Service: srvc,
	}
	e.GET("/product-images", handler.Fetch)
	e.GET("/product-images/:id", handler.GetByID)
	e.POST("/product-images", handler.Store)
	e.DELETE("/product-images/:id", handler.Delete)
	return nil
}

func (h *ProductImageHandler) Fetch(c echo.Context) error {
	ctx := context.Background()
	product_images, err := h.Service.Fetch(ctx)
	if err != nil {
		err_rsp := handleError(err)
		return c.JSON(err_rsp.HttpStatusCode, err_rsp)
	}
	return c.JSON(http.StatusOK, product_images)
}

func (h *ProductImageHandler) GetByID(c echo.Context) error {
	ctx := context.Background()
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	product_image, err := h.Service.GetByID(ctx, id)
	if err != nil {
		err_rsp := handleError(err)
		return c.JSON(err_rsp.HttpStatusCode, err_rsp)
	}
	return c.JSON(http.StatusOK, product_image)
}

func (h *ProductImageHandler) Store(c echo.Context) error {
	ctx := context.Background()
	var product_image *ent.ProductImage
	if err := c.Bind(&product_image); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	resp_obj, err := h.Service.Store(ctx, product_image)
	if err != nil {
		err_rsp := handleError(err)
		return c.JSON(err_rsp.HttpStatusCode, err_rsp)
	}
	return c.JSON(http.StatusCreated, resp_obj)
}

func (h *ProductImageHandler) Delete(c echo.Context) error {
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

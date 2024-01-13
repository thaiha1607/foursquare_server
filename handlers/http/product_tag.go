package http

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/thaiha1607/foursquare_server/ent"
	"github.com/thaiha1607/foursquare_server/interfaces"
)

type ProductTagHandler struct {
	Service interfaces.ProductTagService
}

func NewProductTagHandler(e *echo.Echo, srvc interfaces.ProductTagService) error {
	handler := &ProductTagHandler{
		Service: srvc,
	}
	e.GET("/product-tags", handler.Fetch)
	e.POST("/product-tags", handler.Store)
	e.DELETE("/product-tags", handler.Delete)
	return nil
}

func (h *ProductTagHandler) Fetch(c echo.Context) error {
	ctx := context.Background()
	productTags, err := h.Service.Fetch(ctx)
	if err != nil {
		err_rsp := handleError(err)
		return c.JSON(err_rsp.HttpStatusCode, err_rsp)
	}
	return c.JSON(http.StatusOK, productTags)
}

func (h *ProductTagHandler) Store(c echo.Context) error {
	ctx := context.Background()
	var product_tag ent.ProductTag
	if err := c.Bind(&product_tag); err != nil {
		return c.String(http.StatusBadRequest, "Bad Request")
	}
	if err := h.Service.Store(ctx, &product_tag); err != nil {
		err_rsp := handleError(err)
		return c.JSON(err_rsp.HttpStatusCode, err_rsp)
	}
	return c.JSON(http.StatusOK, product_tag)
}

func (h *ProductTagHandler) Delete(c echo.Context) error {
	ctx := context.Background()
	if err := h.Service.Delete(
		ctx,
		c.QueryParam("product_id"),
		c.QueryParam("tag_id"),
	); err != nil {
		err_rsp := handleError(err)
		return c.JSON(err_rsp.HttpStatusCode, err_rsp)
	}
	return c.String(http.StatusOK, "OK")
}

package http

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/thaiha1607/foursquare_server/ent"
	"github.com/thaiha1607/foursquare_server/interfaces"
)

type ProductTagHandler struct {
	Service interfaces.ProductTagService
}

func NewProductTagHandler(e *echo.Echo, srvc interfaces.ProductTagService) any {
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
		return c.String(http.StatusInternalServerError, "Internal Server Error")
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
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}
	return c.JSON(http.StatusOK, product_tag)
}

func (h *ProductTagHandler) Delete(c echo.Context) error {
	ctx := context.Background()
	product_id, err := uuid.Parse(c.QueryParam("product_id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "Bad Request")
	}
	tag_id, err := uuid.Parse(c.QueryParam("tag_id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "Bad Request")
	}
	if err := h.Service.Delete(ctx, product_id, tag_id); err != nil {
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}
	return c.String(http.StatusOK, "OK")
}

package http

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/thaiha1607/foursquare_server/ent"
	"github.com/thaiha1607/foursquare_server/interfaces"
)

type ProductQtyHandler struct {
	Service interfaces.ProductQtyService
}

func NewProductQtyHandler(e *echo.Echo, srvc interfaces.ProductQtyService) error {
	handler := &ProductQtyHandler{
		Service: srvc,
	}
	e.GET("/product-qtys", handler.Fetch)
	e.GET("/product-qtys/:id", handler.GetByID)
	e.POST("/product-qtys", handler.Store)
	e.PUT("/product-qtys/:id", handler.Update)
	e.DELETE("/product-qtys/:id", handler.Delete)
	return nil
}

func (h *ProductQtyHandler) Fetch(c echo.Context) error {
	ctx := context.Background()
	product_qtys, err := h.Service.Fetch(ctx)
	if err != nil {
		err_rsp := handleError(err)
		return c.JSON(err_rsp.HttpStatusCode, err_rsp)
	}
	return c.JSON(http.StatusOK, product_qtys)
}

func (h *ProductQtyHandler) GetByID(c echo.Context) error {
	ctx := context.Background()
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	product_qty, err := h.Service.GetByID(ctx, id)
	if err != nil {
		err_rsp := handleError(err)
		return c.JSON(err_rsp.HttpStatusCode, err_rsp)
	}
	return c.JSON(http.StatusOK, product_qty)
}

func (h *ProductQtyHandler) Store(c echo.Context) error {
	ctx := context.Background()
	var product_qty *ent.ProductQty
	if err := c.Bind(&product_qty); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	resp_obj, err := h.Service.Store(ctx, product_qty)
	if err != nil {
		err_rsp := handleError(err)
		return c.JSON(err_rsp.HttpStatusCode, err_rsp)
	}
	return c.JSON(http.StatusCreated, resp_obj)
}

func (h *ProductQtyHandler) Update(c echo.Context) error {
	ctx := context.Background()
	var product_qty *ent.ProductQty
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	if err := c.Bind(&product_qty); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	resp_obj, err := h.Service.Update(ctx, id, product_qty)
	if err != nil {
		err_rsp := handleError(err)
		return c.JSON(err_rsp.HttpStatusCode, err_rsp)
	}
	return c.JSON(http.StatusOK, resp_obj)
}

func (h *ProductQtyHandler) Delete(c echo.Context) error {
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

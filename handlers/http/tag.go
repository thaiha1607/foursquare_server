package http

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/thaiha1607/foursquare_server/ent"
	"github.com/thaiha1607/foursquare_server/interfaces"
)

type TagHandler struct {
	Service interfaces.TagService
}

func NewTagHandler(e *echo.Echo, srvc interfaces.TagService) error {
	handler := &TagHandler{
		Service: srvc,
	}
	e.GET("/tags", handler.Fetch)
	e.GET("/tags/:id", handler.GetByID)
	e.POST("/tags", handler.Store)
	e.PUT("/tags/:id", handler.Update)
	e.DELETE("/tags/:id", handler.Delete)
	return nil
}

func (h *TagHandler) Fetch(c echo.Context) error {
	ctx := context.Background()
	tags, err := h.Service.Fetch(ctx)
	if err != nil {
		err_rsp := handleError(err)
		return c.JSON(err_rsp.HttpStatusCode, err_rsp)
	}
	return c.JSON(http.StatusOK, tags)
}

func (h *TagHandler) GetByID(c echo.Context) error {
	ctx := context.Background()
	id := c.Param("id")
	tag, err := h.Service.GetByID(ctx, id)
	if err != nil {
		err_rsp := handleError(err)
		return c.JSON(err_rsp.HttpStatusCode, err_rsp)
	}
	return c.JSON(http.StatusOK, tag)
}

func (h *TagHandler) Store(c echo.Context) error {
	ctx := context.Background()
	var tag *ent.Tag
	if err := c.Bind(&tag); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	resp_obj, err := h.Service.Store(ctx, tag)
	if err != nil {
		err_rsp := handleError(err)
		return c.JSON(err_rsp.HttpStatusCode, err_rsp)
	}
	return c.JSON(http.StatusCreated, resp_obj)
}

func (h *TagHandler) Update(c echo.Context) error {
	ctx := context.Background()
	var tag *ent.Tag
	id := c.Param("id")
	if err := c.Bind(&tag); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	resp_obj, err := h.Service.Update(ctx, id, tag)
	if err != nil {
		err_rsp := handleError(err)
		return c.JSON(err_rsp.HttpStatusCode, err_rsp)
	}
	return c.JSON(http.StatusOK, resp_obj)
}

func (h *TagHandler) Delete(c echo.Context) error {
	ctx := context.Background()
	id := c.Param("id")
	if err := h.Service.Delete(ctx, id); err != nil {
		err_rsp := handleError(err)
		return c.JSON(err_rsp.HttpStatusCode, err_rsp)
	}
	return c.NoContent(http.StatusNoContent)
}

package http

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/thaiha1607/foursquare_server/ent"
	"github.com/thaiha1607/foursquare_server/interfaces"
)

type TagHandler struct {
	Service interfaces.TagService
}

func NewTagHandler(e *echo.Echo, srvc interfaces.TagService) any {
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
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "Bad Request")
	}
	tag, err := h.Service.GetByID(ctx, id)
	if err != nil {
		err_rsp := handleError(err)
		return c.JSON(err_rsp.HttpStatusCode, err_rsp)
	}
	return c.JSON(http.StatusOK, tag)
}

func (h *TagHandler) Store(c echo.Context) error {
	ctx := context.Background()
	var tag ent.Tag
	if err := c.Bind(&tag); err != nil {
		return c.String(http.StatusBadRequest, "Bad Request")
	}
	if err := h.Service.Store(ctx, &tag); err != nil {
		err_rsp := handleError(err)
		return c.JSON(err_rsp.HttpStatusCode, err_rsp)
	}
	return c.JSON(http.StatusOK, tag)
}

func (h *TagHandler) Update(c echo.Context) error {
	ctx := context.Background()
	var tag ent.Tag
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "Bad Request")
	}
	if err := c.Bind(&tag); err != nil {
		return c.String(http.StatusBadRequest, "Bad Request")
	}
	if err := h.Service.Update(ctx, id, &tag); err != nil {
		err_rsp := handleError(err)
		return c.JSON(err_rsp.HttpStatusCode, err_rsp)
	}
	return c.JSON(http.StatusOK, tag)
}

func (h *TagHandler) Delete(c echo.Context) error {
	ctx := context.Background()
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "Bad Request")
	}
	if err := h.Service.Delete(ctx, id); err != nil {
		err_rsp := handleError(err)
		return c.JSON(err_rsp.HttpStatusCode, err_rsp)
	}
	return c.String(http.StatusOK, "OK")
}

package http

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/thaiha1607/foursquare_server/ent"
	"github.com/thaiha1607/foursquare_server/interfaces"
)

type WorkUnitInfoHandler struct {
	Service interfaces.WorkUnitInfoService
}

func NewWorkUnitInfoHandler(e *echo.Echo, srvc interfaces.WorkUnitInfoService) error {
	handler := &WorkUnitInfoHandler{
		Service: srvc,
	}
	e.GET("/work-unit-info", handler.Fetch)
	e.GET("/work-unit-info/:id", handler.GetByID)
	e.POST("/work-unit-info", handler.Store)
	e.PUT("/work-unit-info/:id", handler.Update)
	e.DELETE("/work-unit-info/:id", handler.Delete)
	return nil
}

func (h *WorkUnitInfoHandler) Fetch(c echo.Context) error {
	ctx := context.Background()
	work_unit_info_list, err := h.Service.Fetch(ctx)
	if err != nil {
		err_rsp := handleError(err)
		return c.JSON(err_rsp.HttpStatusCode, err_rsp)
	}
	return c.JSON(http.StatusOK, work_unit_info_list)
}

func (h *WorkUnitInfoHandler) GetByID(c echo.Context) error {
	ctx := context.Background()
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	work_unit_info, err := h.Service.GetByID(ctx, id)
	if err != nil {
		err_rsp := handleError(err)
		return c.JSON(err_rsp.HttpStatusCode, err_rsp)
	}
	return c.JSON(http.StatusOK, work_unit_info)
}

func (h *WorkUnitInfoHandler) Store(c echo.Context) error {
	ctx := context.Background()
	var work_unit_info *ent.WorkUnitInfo
	if err := c.Bind(&work_unit_info); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	resp_obj, err := h.Service.Store(ctx, work_unit_info)
	if err != nil {
		err_rsp := handleError(err)
		return c.JSON(err_rsp.HttpStatusCode, err_rsp)
	}
	return c.JSON(http.StatusCreated, resp_obj)
}

func (h *WorkUnitInfoHandler) Update(c echo.Context) error {
	ctx := context.Background()
	var work_unit_info *ent.WorkUnitInfo
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	if err := c.Bind(&work_unit_info); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	resp_obj, err := h.Service.Update(ctx, id, work_unit_info)
	if err != nil {
		err_rsp := handleError(err)
		return c.JSON(err_rsp.HttpStatusCode, err_rsp)
	}
	return c.JSON(http.StatusOK, resp_obj)
}

func (h *WorkUnitInfoHandler) Delete(c echo.Context) error {
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

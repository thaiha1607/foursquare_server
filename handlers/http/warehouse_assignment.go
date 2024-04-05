package http

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/thaiha1607/foursquare_server/ent"
	"github.com/thaiha1607/foursquare_server/interfaces"
)

type WarehouseAssignmentHandler struct {
	Service interfaces.WarehouseAssignmentService
}

func NewWarehouseAssignmentHandler(e *echo.Echo, srvc interfaces.WarehouseAssignmentService) error {
	handler := &WarehouseAssignmentHandler{
		Service: srvc,
	}
	e.GET("/warehouse-assignments", handler.Fetch)
	e.GET("/warehouse-assignments/:id", handler.GetByID)
	e.POST("/warehouse-assignments", handler.Store)
	e.PUT("/warehouse-assignments/:id", handler.Update)
	e.DELETE("/warehouse-assignments/:id", handler.Delete)
	return nil
}

func (h *WarehouseAssignmentHandler) Fetch(c echo.Context) error {
	ctx := context.Background()
	warehouse_assignments, err := h.Service.Fetch(ctx)
	if err != nil {
		err_rsp := handleError(err)
		return c.JSON(err_rsp.HttpStatusCode, err_rsp)
	}
	return c.JSON(http.StatusOK, warehouse_assignments)
}

func (h *WarehouseAssignmentHandler) GetByID(c echo.Context) error {
	ctx := context.Background()
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	warehouse_assignment, err := h.Service.GetByID(ctx, id)
	if err != nil {
		err_rsp := handleError(err)
		return c.JSON(err_rsp.HttpStatusCode, err_rsp)
	}
	return c.JSON(http.StatusOK, warehouse_assignment)
}

func (h *WarehouseAssignmentHandler) Store(c echo.Context) error {
	ctx := context.Background()
	var warehouse_assignment *ent.WarehouseAssignment
	if err := c.Bind(&warehouse_assignment); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	resp_obj, err := h.Service.Store(ctx, warehouse_assignment)
	if err != nil {
		err_rsp := handleError(err)
		return c.JSON(err_rsp.HttpStatusCode, err_rsp)
	}
	return c.JSON(http.StatusCreated, resp_obj)
}

func (h *WarehouseAssignmentHandler) Update(c echo.Context) error {
	ctx := context.Background()
	var warehouse_assignment *ent.WarehouseAssignment
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	if err := c.Bind(&warehouse_assignment); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	resp_obj, err := h.Service.Update(ctx, id, warehouse_assignment)
	if err != nil {
		err_rsp := handleError(err)
		return c.JSON(err_rsp.HttpStatusCode, err_rsp)
	}
	return c.JSON(http.StatusOK, resp_obj)
}

func (h *WarehouseAssignmentHandler) Delete(c echo.Context) error {
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

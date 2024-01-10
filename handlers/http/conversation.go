package http

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/thaiha1607/foursquare_server/ent"
	"github.com/thaiha1607/foursquare_server/interfaces"
)

type ConversationHandler struct {
	Service interfaces.ConversationService
}

func NewConversationHandler(e *echo.Echo, srvc interfaces.ConversationService) any {
	handler := &ConversationHandler{
		Service: srvc,
	}
	e.GET("/conversations", handler.Fetch)
	e.GET("/conversations/:id", handler.GetByID)
	e.POST("/conversations", handler.Store)
	e.PUT("/conversations/:id", handler.Update)
	e.DELETE("/conversations/:id", handler.Delete)
	return nil
}

func (h *ConversationHandler) Fetch(c echo.Context) error {
	ctx := context.Background()
	conversations, err := h.Service.Fetch(ctx)
	if err != nil {
		err_rsp := handleError(err)
		return c.JSON(err_rsp.HttpStatusCode, err_rsp)
	}
	return c.JSON(http.StatusOK, conversations)
}

func (h *ConversationHandler) GetByID(c echo.Context) error {
	ctx := context.Background()
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "Bad Request")
	}
	conversation, err := h.Service.GetByID(ctx, id)
	if err != nil {
		err_rsp := handleError(err)
		return c.JSON(err_rsp.HttpStatusCode, err_rsp)
	}
	return c.JSON(http.StatusOK, conversation)
}

func (h *ConversationHandler) Store(c echo.Context) error {
	ctx := context.Background()
	var conversation ent.Conversation
	if err := c.Bind(&conversation); err != nil {
		return c.String(http.StatusBadRequest, "Bad Request")
	}
	if err := h.Service.Store(ctx, &conversation); err != nil {
		err_rsp := handleError(err)
		return c.JSON(err_rsp.HttpStatusCode, err_rsp)
	}
	return c.JSON(http.StatusOK, conversation)
}

func (h *ConversationHandler) Update(c echo.Context) error {
	ctx := context.Background()
	var conversation ent.Conversation
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "Bad Request")
	}
	if err := c.Bind(&conversation); err != nil {
		return c.String(http.StatusBadRequest, "Bad Request")
	}
	if err := h.Service.Update(ctx, id, &conversation); err != nil {
		err_rsp := handleError(err)
		return c.JSON(err_rsp.HttpStatusCode, err_rsp)
	}
	return c.JSON(http.StatusOK, conversation)
}

func (h *ConversationHandler) Delete(c echo.Context) error {
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

package http

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/thaiha1607/foursquare_server/ent"
	"github.com/thaiha1607/foursquare_server/interfaces"
)

type MessageHandler struct {
	Service interfaces.MessageService
}

func NewMessageHandler(e *echo.Echo, srvc interfaces.MessageService) any {
	handler := &MessageHandler{
		Service: srvc,
	}
	e.GET("/messages", handler.Fetch)
	e.GET("/messages/:id", handler.GetByID)
	e.POST("/messages", handler.Store)
	e.PUT("/messages/:id", handler.Update)
	e.DELETE("/messages/:id", handler.Delete)
	return nil
}

func (h *MessageHandler) Fetch(c echo.Context) error {
	ctx := context.Background()
	messages, err := h.Service.Fetch(ctx)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}
	return c.JSON(http.StatusOK, messages)
}

func (h *MessageHandler) GetByID(c echo.Context) error {
	ctx := context.Background()
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "Bad Request")
	}
	message, err := h.Service.GetByID(ctx, id)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}
	return c.JSON(http.StatusOK, message)
}

func (h *MessageHandler) Store(c echo.Context) error {
	ctx := context.Background()
	var message ent.Message
	if err := c.Bind(&message); err != nil {
		return c.String(http.StatusBadRequest, "Bad Request")
	}
	if err := h.Service.Store(ctx, &message); err != nil {
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}
	return c.JSON(http.StatusOK, message)
}

func (h *MessageHandler) Update(c echo.Context) error {
	ctx := context.Background()
	var message ent.Message
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "Bad Request")
	}
	if err := c.Bind(&message); err != nil {
		return c.String(http.StatusBadRequest, "Bad Request")
	}
	if err := h.Service.Update(ctx, id, &message); err != nil {
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}
	return c.JSON(http.StatusOK, message)
}

func (h *MessageHandler) Delete(c echo.Context) error {
	ctx := context.Background()
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "Bad Request")
	}
	if err := h.Service.Delete(ctx, id); err != nil {
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}
	return c.String(http.StatusOK, "OK")
}

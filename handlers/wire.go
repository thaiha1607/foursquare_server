//go:build wireinject
// +build wireinject

package handlers

import (
	"github.com/google/wire"
	"github.com/labstack/echo/v4"
	"github.com/thaiha1607/foursquare_server/ent"
	"github.com/thaiha1607/foursquare_server/handlers/http"
	"github.com/thaiha1607/foursquare_server/repositories"
	"github.com/thaiha1607/foursquare_server/services"
)

func initializeOrderHandler(e *echo.Echo, db *ent.Client) error {
	wire.Build(
		http.NewOrderHandler,
		services.NewOrderService,
		repositories.NewEntOrderRepository,
	)
	return nil
}

func initializeOrderStatusCodeHandler(e *echo.Echo, db *ent.Client) error {
	wire.Build(
		http.NewOrderStatusCodeHandler,
		services.NewOrderStatusCodeService,
		repositories.NewEntOrderStatusCodeRepository,
	)
	return nil
}

func initializeProductImageHandler(e *echo.Echo, db *ent.Client) error {
	wire.Build(
		http.NewProductImageHandler,
		services.NewProductImageService,
		repositories.NewEntProductImageRepository,
	)
	return nil
}

func initializeOrderItemHandler(e *echo.Echo, db *ent.Client) error {
	wire.Build(
		http.NewOrderItemHandler,
		services.NewOrderItemService,
		repositories.NewEntOrderItemRepository,
	)
	return nil
}

func initializePersonHandler(e *echo.Echo, db *ent.Client) error {
	wire.Build(
		http.NewPersonHandler,
		services.NewPersonService,
		repositories.NewEntPersonRepository,
	)
	return nil
}

func initializeProductTagHandler(e *echo.Echo, db *ent.Client) error {
	wire.Build(
		http.NewProductTagHandler,
		services.NewProductTagService,
		repositories.NewEntProductTagRepository,
	)
	return nil
}

func initializeMessageHandler(e *echo.Echo, db *ent.Client) error {
	wire.Build(
		http.NewMessageHandler,
		services.NewMessageService,
		repositories.NewEntMessageRepository,
	)
	return nil
}

func initializeTagHandler(e *echo.Echo, db *ent.Client) error {
	wire.Build(
		http.NewTagHandler,
		services.NewTagService,
		repositories.NewEntTagRepository,
	)
	return nil
}

func initializeInvoiceHandler(e *echo.Echo, db *ent.Client) error {
	wire.Build(
		http.NewInvoiceHandler,
		services.NewInvoiceService,
		repositories.NewEntInvoiceRepository,
	)
	return nil
}

func initializeConversationHandler(e *echo.Echo, db *ent.Client) error {
	wire.Build(
		http.NewConversationHandler,
		services.NewConversationService,
		repositories.NewEntConversationRepository,
	)
	return nil
}

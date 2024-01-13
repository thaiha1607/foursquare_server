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

func initializeInvoiceLineItemHandler(e *echo.Echo, db *ent.Client) error {
	wire.Build(
		http.NewInvoiceLineItemHandler,
		services.NewInvoiceLineItemService,
		repositories.NewEntInvoiceLineItemRepository,
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

func initializeOrderLineItemHandler(e *echo.Echo, db *ent.Client) error {
	wire.Build(
		http.NewOrderLineItemHandler,
		services.NewOrderLineItemService,
		repositories.NewEntOrderLineItemRepository,
	)
	return nil
}

func initializeUserHandler(e *echo.Echo, db *ent.Client) error {
	wire.Build(
		http.NewUserHandler,
		services.NewUserService,
		repositories.NewEntUserRepository,
	)
	return nil
}

func initializeFinancialTransactionHandler(e *echo.Echo, db *ent.Client) error {
	wire.Build(
		http.NewFinancialTransactionHandler,
		services.NewFinancialTransactionService,
		repositories.NewEntFinancialTransactionRepository,
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

func initializeProductHandler(e *echo.Echo, db *ent.Client) error {
	wire.Build(
		http.NewProductHandler,
		services.NewProductService,
		repositories.NewEntProductRepository,
	)
	return nil
}

func initializeAccountHandler(e *echo.Echo, db *ent.Client) error {
	wire.Build(
		http.NewAccountHandler,
		services.NewAccountService,
		repositories.NewEntAccountRepository,
	)
	return nil
}

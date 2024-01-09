// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/thaiha1607/foursquare_server/ent"
	"github.com/thaiha1607/foursquare_server/handlers/http"
	"github.com/thaiha1607/foursquare_server/repositories"
	"github.com/thaiha1607/foursquare_server/services"
)

// Injectors from wire.go:

func initializeOrderHandler(e *echo.Echo, db *ent.Client) any {
	orderRepository := repositories.NewEntOrderRepository(db)
	orderService := services.NewOrderService(orderRepository)
	v := http.NewOrderHandler(e, orderService)
	return v
}

func initializeOrderStatusCodeHandler(e *echo.Echo, db *ent.Client) any {
	orderStatusCodeRepository := repositories.NewEntOrderStatusCodeRepository(db)
	orderStatusCodeService := services.NewOrderStatusCodeService(orderStatusCodeRepository)
	v := http.NewOrderStatusCodeHandler(e, orderStatusCodeService)
	return v
}

func initializeInvoiceLineItemHandler(e *echo.Echo, db *ent.Client) any {
	invoiceLineItemRepository := repositories.NewEntInvoiceLineItemRepository(db)
	invoiceLineItemService := services.NewInvoiceLineItemService(invoiceLineItemRepository)
	v := http.NewInvoiceLineItemHandler(e, invoiceLineItemService)
	return v
}

func initializeProductImageHandler(e *echo.Echo, db *ent.Client) any {
	productImageRepository := repositories.NewEntProductImageRepository(db)
	productImageService := services.NewProductImageService(productImageRepository)
	v := http.NewProductImageHandler(e, productImageService)
	return v
}

func initializeOrderLineItemHandler(e *echo.Echo, db *ent.Client) any {
	orderLineItemRepository := repositories.NewEntOrderLineItemRepository(db)
	orderLineItemService := services.NewOrderLineItemService(orderLineItemRepository)
	v := http.NewOrderLineItemHandler(e, orderLineItemService)
	return v
}

func initializeUserHandler(e *echo.Echo, db *ent.Client) any {
	userRepository := repositories.NewEntUserRepository(db)
	userService := services.NewUserService(userRepository)
	v := http.NewUserHandler(e, userService)
	return v
}

func initializeFinancialTransactionHandler(e *echo.Echo, db *ent.Client) any {
	financialTransactionRepository := repositories.NewEntFinancialTransactionRepository(db)
	financialTransactionService := services.NewFinancialTransactionService(financialTransactionRepository)
	v := http.NewFinancialTransactionHandler(e, financialTransactionService)
	return v
}

func initializeProductTagHandler(e *echo.Echo, db *ent.Client) any {
	productTagRepository := repositories.NewEntProductTagRepository(db)
	productTagService := services.NewProductTagService(productTagRepository)
	v := http.NewProductTagHandler(e, productTagService)
	return v
}

func initializeMessageHandler(e *echo.Echo, db *ent.Client) any {
	messageRepository := repositories.NewEntMessageRepository(db)
	messageService := services.NewMessageService(messageRepository)
	v := http.NewMessageHandler(e, messageService)
	return v
}

func initializeTagHandler(e *echo.Echo, db *ent.Client) any {
	tagRepository := repositories.NewEntTagRepository(db)
	tagService := services.NewTagService(tagRepository)
	v := http.NewTagHandler(e, tagService)
	return v
}

func initializeInvoiceHandler(e *echo.Echo, db *ent.Client) any {
	invoiceRepository := repositories.NewEntInvoiceRepository(db)
	invoiceService := services.NewInvoiceService(invoiceRepository)
	v := http.NewInvoiceHandler(e, invoiceService)
	return v
}

func initializeConversationHandler(e *echo.Echo, db *ent.Client) any {
	conversationRepository := repositories.NewEntConversationRepository(db)
	conversationService := services.NewConversationService(conversationRepository)
	v := http.NewConversationHandler(e, conversationService)
	return v
}

func initializeProductHandler(e *echo.Echo, db *ent.Client) any {
	productRepository := repositories.NewEntProductRepository(db)
	productService := services.NewProductService(productRepository)
	v := http.NewProductHandler(e, productService)
	return v
}

func initializeAccountHandler(e *echo.Echo, db *ent.Client) any {
	accountRepository := repositories.NewEntAccountRepository(db)
	accountService := services.NewAccountService(accountRepository)
	v := http.NewAccountHandler(e, accountService)
	return v
}

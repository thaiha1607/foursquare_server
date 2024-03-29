package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/thaiha1607/foursquare_server/ent"
)

func RegisterRoutes(e *echo.Echo, db *ent.Client) {
	initializeApi(e, db)
}

// Where Dependency Injection happens
func initializeApi(e *echo.Echo, db *ent.Client) {
	initializeOrderHandler(e, db)
	initializeOrderStatusCodeHandler(e, db)
	initializeInvoiceLineItemHandler(e, db)
	initializeProductImageHandler(e, db)
	initializeOrderLineItemHandler(e, db)
	initializePersonHandler(e, db)
	initializeFinancialTransactionHandler(e, db)
	initializeProductTagHandler(e, db)
	initializeMessageHandler(e, db)
	initializeTagHandler(e, db)
	initializeInvoiceHandler(e, db)
	initializeConversationHandler(e, db)
	initializeProductHandler(e, db)
}

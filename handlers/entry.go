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
	initializeInvoiceStatusCodeHandler(e, db)
	initializeOrderStatusCodeHandler(e, db)
	initializeAddressHandler(e, db)
	initializeShipmentHistoryHandler(e, db)
	initializeWarehouseAssignmentHandler(e, db)
	initializeShipmentHandler(e, db)
	initializeProductImageHandler(e, db)
	initializeShipmentItemHandler(e, db)
	initializeInvoiceHistoryHandler(e, db)
	initializeShipmentStatusCodeHandler(e, db)
	initializeOrderHistoryHandler(e, db)
	initializeProductColorHandler(e, db)
	initializePersonHandler(e, db)
	initializeProductTagHandler(e, db)
	initializeMessageHandler(e, db)
	initializeTagHandler(e, db)
	initializeWorkUnitInfoHandler(e, db)
	initializeInvoiceHandler(e, db)
	initializeOrderItemHandler(e, db)
	initializeConversationHandler(e, db)
	initializeProductInfoHandler(e, db)
	initializeProductQtyHandler(e, db)
	initializePersonAddressHandler(e, db)
}

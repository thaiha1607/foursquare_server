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

func initializeInvoiceStatusCodeHandler(e *echo.Echo, db *ent.Client) error {
	wire.Build(
		http.NewInvoiceStatusCodeHandler,
		services.NewInvoiceStatusCodeService,
		repositories.NewEntInvoiceStatusCodeRepository,
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

func initializeAddressHandler(e *echo.Echo, db *ent.Client) error {
	wire.Build(
		http.NewAddressHandler,
		services.NewAddressService,
		repositories.NewEntAddressRepository,
	)
	return nil
}

func initializeShipmentHistoryHandler(e *echo.Echo, db *ent.Client) error {
	wire.Build(
		http.NewShipmentHistoryHandler,
		services.NewShipmentHistoryService,
		repositories.NewEntShipmentHistoryRepository,
	)
	return nil
}

func initializeWarehouseAssignmentHandler(e *echo.Echo, db *ent.Client) error {
	wire.Build(
		http.NewWarehouseAssignmentHandler,
		services.NewWarehouseAssignmentService,
		repositories.NewEntWarehouseAssignmentRepository,
	)
	return nil
}

func initializeShipmentHandler(e *echo.Echo, db *ent.Client) error {
	wire.Build(
		http.NewShipmentHandler,
		services.NewShipmentService,
		repositories.NewEntShipmentRepository,
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

func initializeShipmentItemHandler(e *echo.Echo, db *ent.Client) error {
	wire.Build(
		http.NewShipmentItemHandler,
		services.NewShipmentItemService,
		repositories.NewEntShipmentItemRepository,
	)
	return nil
}

func initializeInvoiceHistoryHandler(e *echo.Echo, db *ent.Client) error {
	wire.Build(
		http.NewInvoiceHistoryHandler,
		services.NewInvoiceHistoryService,
		repositories.NewEntInvoiceHistoryRepository,
	)
	return nil
}

func initializeShipmentStatusCodeHandler(e *echo.Echo, db *ent.Client) error {
	wire.Build(
		http.NewShipmentStatusCodeHandler,
		services.NewShipmentStatusCodeService,
		repositories.NewEntShipmentStatusCodeRepository,
	)
	return nil
}

func initializeOrderHistoryHandler(e *echo.Echo, db *ent.Client) error {
	wire.Build(
		http.NewOrderHistoryHandler,
		services.NewOrderHistoryService,
		repositories.NewEntOrderHistoryRepository,
	)
	return nil
}

func initializeProductColorHandler(e *echo.Echo, db *ent.Client) error {
	wire.Build(
		http.NewProductColorHandler,
		services.NewProductColorService,
		repositories.NewEntProductColorRepository,
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

func initializeWorkUnitInfoHandler(e *echo.Echo, db *ent.Client) error {
	wire.Build(
		http.NewWorkUnitInfoHandler,
		services.NewWorkUnitInfoService,
		repositories.NewEntWorkUnitInfoRepository,
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

func initializeOrderItemHandler(e *echo.Echo, db *ent.Client) error {
	wire.Build(
		http.NewOrderItemHandler,
		services.NewOrderItemService,
		repositories.NewEntOrderItemRepository,
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

func initializeProductInfoHandler(e *echo.Echo, db *ent.Client) error {
	wire.Build(
		http.NewProductInfoHandler,
		services.NewProductInfoService,
		repositories.NewEntProductInfoRepository,
	)
	return nil
}

func initializeProductQtyHandler(e *echo.Echo, db *ent.Client) error {
	wire.Build(
		http.NewProductQtyHandler,
		services.NewProductQtyService,
		repositories.NewEntProductQtyRepository,
	)
	return nil
}

func initializePersonAddressHandler(e *echo.Echo, db *ent.Client) error {
	wire.Build(
		http.NewPersonAddressHandler,
		services.NewPersonAddressService,
		repositories.NewEntPersonAddressRepository,
	)
	return nil
}

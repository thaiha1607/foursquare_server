// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AddressesColumns holds the columns for the "addresses" table.
	AddressesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "line1", Type: field.TypeString},
		{Name: "line2", Type: field.TypeString, Nullable: true},
		{Name: "city", Type: field.TypeString},
		{Name: "state_or_province", Type: field.TypeString},
		{Name: "zip_or_postcode", Type: field.TypeString},
		{Name: "country", Type: field.TypeString},
		{Name: "other_address_details", Type: field.TypeString, Nullable: true},
	}
	// AddressesTable holds the schema information for the "addresses" table.
	AddressesTable = &schema.Table{
		Name:       "addresses",
		Columns:    AddressesColumns,
		PrimaryKey: []*schema.Column{AddressesColumns[0]},
	}
	// ConversationsColumns holds the columns for the "conversations" table.
	ConversationsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "title", Type: field.TypeString, Nullable: true},
		{Name: "person_one_id", Type: field.TypeUUID},
		{Name: "person_two_id", Type: field.TypeUUID},
	}
	// ConversationsTable holds the schema information for the "conversations" table.
	ConversationsTable = &schema.Table{
		Name:       "conversations",
		Columns:    ConversationsColumns,
		PrimaryKey: []*schema.Column{ConversationsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "conversations_persons_person_one",
				Columns:    []*schema.Column{ConversationsColumns[4]},
				RefColumns: []*schema.Column{PersonsColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "conversations_persons_person_two",
				Columns:    []*schema.Column{ConversationsColumns[5]},
				RefColumns: []*schema.Column{PersonsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "conversation_person_one_id_person_two_id",
				Unique:  true,
				Columns: []*schema.Column{ConversationsColumns[4], ConversationsColumns[5]},
			},
		},
	}
	// InvoicesColumns holds the columns for the "invoices" table.
	InvoicesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "total", Type: field.TypeFloat64, SchemaType: map[string]string{"mysql": "decimal(12,2)", "postgres": "numeric(12,2)"}},
		{Name: "note", Type: field.TypeString, Nullable: true},
		{Name: "type", Type: field.TypeEnum, Enums: []string{"PRO_FORMA", "REGULAR", "PAST_DUE", "INTERIM", "TIMESHEET", "FINAL", "CREDIT", "DEBIT", "MIXED", "COMMERCIAL", "RECURRING", "OTHER"}, Default: "PRO_FORMA"},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"DRAFT", "ACTIVE", "SENT", "DISPUTED", "OVERDUE", "PARTIAL", "PAID", "VOID", "DEBT", "OTHER"}, Default: "DRAFT"},
		{Name: "payment_method", Type: field.TypeEnum, Nullable: true, Enums: []string{"CASH", "ELECTRONIC_FUNDS_TRANSFER", "GIFT_CARD", "CREDIT_CARD", "DEBIT_CARD", "PREPAID_CARD", "CHECK", "OTHER"}, Default: "CASH"},
		{Name: "order_id", Type: field.TypeUUID},
	}
	// InvoicesTable holds the schema information for the "invoices" table.
	InvoicesTable = &schema.Table{
		Name:       "invoices",
		Columns:    InvoicesColumns,
		PrimaryKey: []*schema.Column{InvoicesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "invoices_orders_order",
				Columns:    []*schema.Column{InvoicesColumns[8]},
				RefColumns: []*schema.Column{OrdersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// InvoiceHistoriesColumns holds the columns for the "invoice_histories" table.
	InvoiceHistoriesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "description", Type: field.TypeString, Nullable: true},
		{Name: "invoice_id", Type: field.TypeUUID},
		{Name: "person_id", Type: field.TypeUUID},
		{Name: "old_status_code", Type: field.TypeInt, Nullable: true},
		{Name: "new_status_code", Type: field.TypeInt, Nullable: true},
	}
	// InvoiceHistoriesTable holds the schema information for the "invoice_histories" table.
	InvoiceHistoriesTable = &schema.Table{
		Name:       "invoice_histories",
		Columns:    InvoiceHistoriesColumns,
		PrimaryKey: []*schema.Column{InvoiceHistoriesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "invoice_histories_invoices_invoice",
				Columns:    []*schema.Column{InvoiceHistoriesColumns[4]},
				RefColumns: []*schema.Column{InvoicesColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "invoice_histories_persons_person",
				Columns:    []*schema.Column{InvoiceHistoriesColumns[5]},
				RefColumns: []*schema.Column{PersonsColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "invoice_histories_order_status_codes_old_status",
				Columns:    []*schema.Column{InvoiceHistoriesColumns[6]},
				RefColumns: []*schema.Column{OrderStatusCodesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "invoice_histories_order_status_codes_new_status",
				Columns:    []*schema.Column{InvoiceHistoriesColumns[7]},
				RefColumns: []*schema.Column{OrderStatusCodesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// MessagesColumns holds the columns for the "messages" table.
	MessagesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "type", Type: field.TypeEnum, Enums: []string{"TEXT", "IMAGE", "VIDEO", "AUDIO", "FILE", "OTHER"}, Default: "TEXT"},
		{Name: "content", Type: field.TypeString},
		{Name: "is_read", Type: field.TypeBool, Default: false},
		{Name: "conversation_id", Type: field.TypeUUID},
		{Name: "sender_id", Type: field.TypeUUID},
	}
	// MessagesTable holds the schema information for the "messages" table.
	MessagesTable = &schema.Table{
		Name:       "messages",
		Columns:    MessagesColumns,
		PrimaryKey: []*schema.Column{MessagesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "messages_conversations_conversation",
				Columns:    []*schema.Column{MessagesColumns[6]},
				RefColumns: []*schema.Column{ConversationsColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "messages_persons_sender",
				Columns:    []*schema.Column{MessagesColumns[7]},
				RefColumns: []*schema.Column{PersonsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// OrdersColumns holds the columns for the "orders" table.
	OrdersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "note", Type: field.TypeString, Nullable: true},
		{Name: "priority", Type: field.TypeInt, Default: 0},
		{Name: "type", Type: field.TypeEnum, Enums: []string{"SALE", "RETURN", "EXCHANGE", "TRANSFER", "OTHER"}, Default: "SALE"},
		{Name: "internal_note", Type: field.TypeString, Nullable: true},
		{Name: "is_internal", Type: field.TypeBool, Default: false},
		{Name: "customer_id", Type: field.TypeUUID},
		{Name: "created_by", Type: field.TypeUUID},
		{Name: "parent_order_id", Type: field.TypeUUID, Unique: true, Nullable: true},
		{Name: "status_code", Type: field.TypeInt, Default: 1},
		{Name: "staff_id", Type: field.TypeUUID},
		{Name: "address_id", Type: field.TypeUUID},
	}
	// OrdersTable holds the schema information for the "orders" table.
	OrdersTable = &schema.Table{
		Name:       "orders",
		Columns:    OrdersColumns,
		PrimaryKey: []*schema.Column{OrdersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "orders_persons_customer",
				Columns:    []*schema.Column{OrdersColumns[8]},
				RefColumns: []*schema.Column{PersonsColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "orders_persons_creator",
				Columns:    []*schema.Column{OrdersColumns[9]},
				RefColumns: []*schema.Column{PersonsColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "orders_orders_parent_order",
				Columns:    []*schema.Column{OrdersColumns[10]},
				RefColumns: []*schema.Column{OrdersColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "orders_order_status_codes_order_status",
				Columns:    []*schema.Column{OrdersColumns[11]},
				RefColumns: []*schema.Column{OrderStatusCodesColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "orders_persons_staff",
				Columns:    []*schema.Column{OrdersColumns[12]},
				RefColumns: []*schema.Column{PersonsColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "orders_addresses_order_address",
				Columns:    []*schema.Column{OrdersColumns[13]},
				RefColumns: []*schema.Column{AddressesColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// OrderHistoriesColumns holds the columns for the "order_histories" table.
	OrderHistoriesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "description", Type: field.TypeString, Nullable: true},
		{Name: "order_id", Type: field.TypeUUID},
		{Name: "person_id", Type: field.TypeUUID},
		{Name: "old_status_code", Type: field.TypeInt, Nullable: true},
		{Name: "new_status_code", Type: field.TypeInt, Nullable: true},
	}
	// OrderHistoriesTable holds the schema information for the "order_histories" table.
	OrderHistoriesTable = &schema.Table{
		Name:       "order_histories",
		Columns:    OrderHistoriesColumns,
		PrimaryKey: []*schema.Column{OrderHistoriesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "order_histories_orders_order",
				Columns:    []*schema.Column{OrderHistoriesColumns[4]},
				RefColumns: []*schema.Column{OrdersColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "order_histories_persons_person",
				Columns:    []*schema.Column{OrderHistoriesColumns[5]},
				RefColumns: []*schema.Column{PersonsColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "order_histories_order_status_codes_old_status",
				Columns:    []*schema.Column{OrderHistoriesColumns[6]},
				RefColumns: []*schema.Column{OrderStatusCodesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "order_histories_order_status_codes_new_status",
				Columns:    []*schema.Column{OrderHistoriesColumns[7]},
				RefColumns: []*schema.Column{OrderStatusCodesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// OrderItemsColumns holds the columns for the "order_items" table.
	OrderItemsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "qty", Type: field.TypeFloat64, SchemaType: map[string]string{"mysql": "decimal(12,2)", "postgres": "numeric(12,2)"}},
		{Name: "price_per_unit", Type: field.TypeFloat64, SchemaType: map[string]string{"mysql": "decimal(12,2)", "postgres": "numeric(12,2)"}},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"DELIVERED", "OUT_OF_STOCK", "IN_TRANSIT", "IN_STOCK", "PARTIALLY_DELIVERED"}, Default: "IN_STOCK"},
		{Name: "order_id", Type: field.TypeUUID},
		{Name: "product_id", Type: field.TypeString},
		{Name: "product_color_id", Type: field.TypeString},
		{Name: "src_unit_id", Type: field.TypeUUID, Nullable: true},
		{Name: "dst_unit_id", Type: field.TypeUUID, Nullable: true},
	}
	// OrderItemsTable holds the schema information for the "order_items" table.
	OrderItemsTable = &schema.Table{
		Name:       "order_items",
		Columns:    OrderItemsColumns,
		PrimaryKey: []*schema.Column{OrderItemsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "order_items_orders_order",
				Columns:    []*schema.Column{OrderItemsColumns[6]},
				RefColumns: []*schema.Column{OrdersColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "order_items_product_info_product",
				Columns:    []*schema.Column{OrderItemsColumns[7]},
				RefColumns: []*schema.Column{ProductInfoColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "order_items_product_color_product_color",
				Columns:    []*schema.Column{OrderItemsColumns[8]},
				RefColumns: []*schema.Column{ProductColorColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "order_items_work_unit_info_source_work_unit",
				Columns:    []*schema.Column{OrderItemsColumns[9]},
				RefColumns: []*schema.Column{WorkUnitInfoColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "order_items_work_unit_info_destination_work_unit",
				Columns:    []*schema.Column{OrderItemsColumns[10]},
				RefColumns: []*schema.Column{WorkUnitInfoColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "orderitem_order_id_product_id_product_color_id",
				Unique:  true,
				Columns: []*schema.Column{OrderItemsColumns[6], OrderItemsColumns[7], OrderItemsColumns[8]},
			},
		},
	}
	// OrderStatusCodesColumns holds the columns for the "order_status_codes" table.
	OrderStatusCodesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "order_status", Type: field.TypeString, Unique: true},
	}
	// OrderStatusCodesTable holds the schema information for the "order_status_codes" table.
	OrderStatusCodesTable = &schema.Table{
		Name:       "order_status_codes",
		Columns:    OrderStatusCodesColumns,
		PrimaryKey: []*schema.Column{OrderStatusCodesColumns[0]},
	}
	// PersonsColumns holds the columns for the "persons" table.
	PersonsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "avatar_url", Type: field.TypeString, Nullable: true},
		{Name: "email", Type: field.TypeString, Unique: true},
		{Name: "name", Type: field.TypeString},
		{Name: "phone", Type: field.TypeString, Unique: true, Nullable: true},
		{Name: "role", Type: field.TypeEnum, Enums: []string{"SALESPERSON", "CUSTOMER", "WAREHOUSE", "DELIVERY", "MANAGEMENT"}, Default: "CUSTOMER"},
		{Name: "password_hash", Type: field.TypeBytes, Nullable: true},
		{Name: "is_email_verified", Type: field.TypeBool, Default: false},
		{Name: "is_phone_verified", Type: field.TypeBool, Default: false},
		{Name: "work_unit_id", Type: field.TypeUUID, Nullable: true},
	}
	// PersonsTable holds the schema information for the "persons" table.
	PersonsTable = &schema.Table{
		Name:       "persons",
		Columns:    PersonsColumns,
		PrimaryKey: []*schema.Column{PersonsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "persons_work_unit_info_work_unit",
				Columns:    []*schema.Column{PersonsColumns[11]},
				RefColumns: []*schema.Column{WorkUnitInfoColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// PersonAddressesColumns holds the columns for the "person_addresses" table.
	PersonAddressesColumns = []*schema.Column{
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "person_id", Type: field.TypeUUID},
		{Name: "address_id", Type: field.TypeUUID},
	}
	// PersonAddressesTable holds the schema information for the "person_addresses" table.
	PersonAddressesTable = &schema.Table{
		Name:       "person_addresses",
		Columns:    PersonAddressesColumns,
		PrimaryKey: []*schema.Column{PersonAddressesColumns[2], PersonAddressesColumns[3]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "person_addresses_persons_persons",
				Columns:    []*schema.Column{PersonAddressesColumns[2]},
				RefColumns: []*schema.Column{PersonsColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "person_addresses_addresses_addresses",
				Columns:    []*schema.Column{PersonAddressesColumns[3]},
				RefColumns: []*schema.Column{AddressesColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// ProductColorColumns holds the columns for the "product_color" table.
	ProductColorColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString},
		{Name: "color_code", Type: field.TypeString},
	}
	// ProductColorTable holds the schema information for the "product_color" table.
	ProductColorTable = &schema.Table{
		Name:       "product_color",
		Columns:    ProductColorColumns,
		PrimaryKey: []*schema.Column{ProductColorColumns[0]},
	}
	// ProductImagesColumns holds the columns for the "product_images" table.
	ProductImagesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "image_url", Type: field.TypeString},
		{Name: "product_id", Type: field.TypeString},
	}
	// ProductImagesTable holds the schema information for the "product_images" table.
	ProductImagesTable = &schema.Table{
		Name:       "product_images",
		Columns:    ProductImagesColumns,
		PrimaryKey: []*schema.Column{ProductImagesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "product_images_product_info_product",
				Columns:    []*schema.Column{ProductImagesColumns[4]},
				RefColumns: []*schema.Column{ProductInfoColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// ProductInfoColumns holds the columns for the "product_info" table.
	ProductInfoColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString},
		{Name: "description", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "year", Type: field.TypeInt, Nullable: true},
		{Name: "provider", Type: field.TypeString, Nullable: true},
	}
	// ProductInfoTable holds the schema information for the "product_info" table.
	ProductInfoTable = &schema.Table{
		Name:       "product_info",
		Columns:    ProductInfoColumns,
		PrimaryKey: []*schema.Column{ProductInfoColumns[0]},
	}
	// ProductQtyColumns holds the columns for the "product_qty" table.
	ProductQtyColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "price_per_unit", Type: field.TypeFloat64, SchemaType: map[string]string{"mysql": "decimal(12,2)", "postgres": "numeric(12,2)"}},
		{Name: "qty", Type: field.TypeFloat64, SchemaType: map[string]string{"mysql": "decimal(12,2)", "postgres": "numeric(12,2)"}},
		{Name: "work_unit_id", Type: field.TypeUUID},
		{Name: "product_id", Type: field.TypeString},
		{Name: "product_color_id", Type: field.TypeString},
	}
	// ProductQtyTable holds the schema information for the "product_qty" table.
	ProductQtyTable = &schema.Table{
		Name:       "product_qty",
		Columns:    ProductQtyColumns,
		PrimaryKey: []*schema.Column{ProductQtyColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "product_qty_work_unit_info_work_unit",
				Columns:    []*schema.Column{ProductQtyColumns[5]},
				RefColumns: []*schema.Column{WorkUnitInfoColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "product_qty_product_info_product",
				Columns:    []*schema.Column{ProductQtyColumns[6]},
				RefColumns: []*schema.Column{ProductInfoColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "product_qty_product_color_product_color",
				Columns:    []*schema.Column{ProductQtyColumns[7]},
				RefColumns: []*schema.Column{ProductColorColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "productqty_work_unit_id_product_id_product_color_id",
				Unique:  true,
				Columns: []*schema.Column{ProductQtyColumns[5], ProductQtyColumns[6], ProductQtyColumns[7]},
			},
		},
	}
	// ProductTagsColumns holds the columns for the "product_tags" table.
	ProductTagsColumns = []*schema.Column{
		{Name: "product_id", Type: field.TypeString},
		{Name: "tag_id", Type: field.TypeString},
	}
	// ProductTagsTable holds the schema information for the "product_tags" table.
	ProductTagsTable = &schema.Table{
		Name:       "product_tags",
		Columns:    ProductTagsColumns,
		PrimaryKey: []*schema.Column{ProductTagsColumns[0], ProductTagsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "product_tags_product_info_products",
				Columns:    []*schema.Column{ProductTagsColumns[0]},
				RefColumns: []*schema.Column{ProductInfoColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "product_tags_tags_tags",
				Columns:    []*schema.Column{ProductTagsColumns[1]},
				RefColumns: []*schema.Column{TagsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// ShipmentsColumns holds the columns for the "shipments" table.
	ShipmentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true, Size: 26},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "shipment_date", Type: field.TypeTime},
		{Name: "note", Type: field.TypeString, Nullable: true},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"PENDING", "ACCEPTED", "REJECTED"}, Default: "PENDING"},
		{Name: "order_id", Type: field.TypeUUID},
		{Name: "invoice_id", Type: field.TypeUUID},
		{Name: "staff_id", Type: field.TypeUUID},
	}
	// ShipmentsTable holds the schema information for the "shipments" table.
	ShipmentsTable = &schema.Table{
		Name:       "shipments",
		Columns:    ShipmentsColumns,
		PrimaryKey: []*schema.Column{ShipmentsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "shipments_orders_order",
				Columns:    []*schema.Column{ShipmentsColumns[6]},
				RefColumns: []*schema.Column{OrdersColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "shipments_invoices_invoice",
				Columns:    []*schema.Column{ShipmentsColumns[7]},
				RefColumns: []*schema.Column{InvoicesColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "shipments_persons_staff",
				Columns:    []*schema.Column{ShipmentsColumns[8]},
				RefColumns: []*schema.Column{PersonsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// ShipmentHistoriesColumns holds the columns for the "shipment_histories" table.
	ShipmentHistoriesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "description", Type: field.TypeString, Nullable: true},
		{Name: "shipment_id", Type: field.TypeString, Size: 26},
		{Name: "person_id", Type: field.TypeUUID},
		{Name: "old_status_code", Type: field.TypeInt, Nullable: true},
		{Name: "new_status_code", Type: field.TypeInt, Nullable: true},
	}
	// ShipmentHistoriesTable holds the schema information for the "shipment_histories" table.
	ShipmentHistoriesTable = &schema.Table{
		Name:       "shipment_histories",
		Columns:    ShipmentHistoriesColumns,
		PrimaryKey: []*schema.Column{ShipmentHistoriesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "shipment_histories_shipments_shipment",
				Columns:    []*schema.Column{ShipmentHistoriesColumns[4]},
				RefColumns: []*schema.Column{ShipmentsColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "shipment_histories_persons_person",
				Columns:    []*schema.Column{ShipmentHistoriesColumns[5]},
				RefColumns: []*schema.Column{PersonsColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "shipment_histories_order_status_codes_old_status",
				Columns:    []*schema.Column{ShipmentHistoriesColumns[6]},
				RefColumns: []*schema.Column{OrderStatusCodesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "shipment_histories_order_status_codes_new_status",
				Columns:    []*schema.Column{ShipmentHistoriesColumns[7]},
				RefColumns: []*schema.Column{OrderStatusCodesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// InvoiceLineItemsColumns holds the columns for the "invoice_line_items" table.
	InvoiceLineItemsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "qty", Type: field.TypeFloat64, SchemaType: map[string]string{"mysql": "decimal(12,2)", "postgres": "numeric(12,2)"}},
		{Name: "total", Type: field.TypeFloat64, SchemaType: map[string]string{"mysql": "decimal(12,2)", "postgres": "numeric(12,2)"}},
		{Name: "shipment_id", Type: field.TypeString, Size: 26},
		{Name: "order_item_id", Type: field.TypeUUID},
	}
	// InvoiceLineItemsTable holds the schema information for the "invoice_line_items" table.
	InvoiceLineItemsTable = &schema.Table{
		Name:       "invoice_line_items",
		Columns:    InvoiceLineItemsColumns,
		PrimaryKey: []*schema.Column{InvoiceLineItemsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "invoice_line_items_shipments_shipment",
				Columns:    []*schema.Column{InvoiceLineItemsColumns[5]},
				RefColumns: []*schema.Column{ShipmentsColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "invoice_line_items_order_items_order_item",
				Columns:    []*schema.Column{InvoiceLineItemsColumns[6]},
				RefColumns: []*schema.Column{OrderItemsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "shipmentitem_shipment_id_order_item_id",
				Unique:  true,
				Columns: []*schema.Column{InvoiceLineItemsColumns[5], InvoiceLineItemsColumns[6]},
			},
		},
	}
	// TagsColumns holds the columns for the "tags" table.
	TagsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "title", Type: field.TypeString, Unique: true},
	}
	// TagsTable holds the schema information for the "tags" table.
	TagsTable = &schema.Table{
		Name:       "tags",
		Columns:    TagsColumns,
		PrimaryKey: []*schema.Column{TagsColumns[0]},
	}
	// WarehouseAssignmentsColumns holds the columns for the "warehouse_assignments" table.
	WarehouseAssignmentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"PENDING", "ACCEPTED", "REJECTED"}, Default: "PENDING"},
		{Name: "note", Type: field.TypeString, Nullable: true},
		{Name: "order_id", Type: field.TypeUUID},
		{Name: "work_unit_id", Type: field.TypeUUID},
		{Name: "staff_id", Type: field.TypeUUID, Nullable: true},
	}
	// WarehouseAssignmentsTable holds the schema information for the "warehouse_assignments" table.
	WarehouseAssignmentsTable = &schema.Table{
		Name:       "warehouse_assignments",
		Columns:    WarehouseAssignmentsColumns,
		PrimaryKey: []*schema.Column{WarehouseAssignmentsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "warehouse_assignments_orders_order",
				Columns:    []*schema.Column{WarehouseAssignmentsColumns[5]},
				RefColumns: []*schema.Column{OrdersColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "warehouse_assignments_work_unit_info_work_unit",
				Columns:    []*schema.Column{WarehouseAssignmentsColumns[6]},
				RefColumns: []*schema.Column{WorkUnitInfoColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "warehouse_assignments_persons_staff",
				Columns:    []*schema.Column{WarehouseAssignmentsColumns[7]},
				RefColumns: []*schema.Column{PersonsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// WorkUnitInfoColumns holds the columns for the "work_unit_info" table.
	WorkUnitInfoColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString},
		{Name: "type", Type: field.TypeEnum, Enums: []string{"WAREHOUSE", "OFFICE", "DELIVERY"}, Default: "WAREHOUSE"},
		{Name: "image_url", Type: field.TypeString, Nullable: true},
		{Name: "address_id", Type: field.TypeUUID, Nullable: true},
	}
	// WorkUnitInfoTable holds the schema information for the "work_unit_info" table.
	WorkUnitInfoTable = &schema.Table{
		Name:       "work_unit_info",
		Columns:    WorkUnitInfoColumns,
		PrimaryKey: []*schema.Column{WorkUnitInfoColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "work_unit_info_addresses_address",
				Columns:    []*schema.Column{WorkUnitInfoColumns[6]},
				RefColumns: []*schema.Column{AddressesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AddressesTable,
		ConversationsTable,
		InvoicesTable,
		InvoiceHistoriesTable,
		MessagesTable,
		OrdersTable,
		OrderHistoriesTable,
		OrderItemsTable,
		OrderStatusCodesTable,
		PersonsTable,
		PersonAddressesTable,
		ProductColorTable,
		ProductImagesTable,
		ProductInfoTable,
		ProductQtyTable,
		ProductTagsTable,
		ShipmentsTable,
		ShipmentHistoriesTable,
		InvoiceLineItemsTable,
		TagsTable,
		WarehouseAssignmentsTable,
		WorkUnitInfoTable,
	}
)

func init() {
	ConversationsTable.ForeignKeys[0].RefTable = PersonsTable
	ConversationsTable.ForeignKeys[1].RefTable = PersonsTable
	InvoicesTable.ForeignKeys[0].RefTable = OrdersTable
	InvoiceHistoriesTable.ForeignKeys[0].RefTable = InvoicesTable
	InvoiceHistoriesTable.ForeignKeys[1].RefTable = PersonsTable
	InvoiceHistoriesTable.ForeignKeys[2].RefTable = OrderStatusCodesTable
	InvoiceHistoriesTable.ForeignKeys[3].RefTable = OrderStatusCodesTable
	InvoiceHistoriesTable.Annotation = &entsql.Annotation{
		Table: "invoice_histories",
	}
	MessagesTable.ForeignKeys[0].RefTable = ConversationsTable
	MessagesTable.ForeignKeys[1].RefTable = PersonsTable
	OrdersTable.ForeignKeys[0].RefTable = PersonsTable
	OrdersTable.ForeignKeys[1].RefTable = PersonsTable
	OrdersTable.ForeignKeys[2].RefTable = OrdersTable
	OrdersTable.ForeignKeys[3].RefTable = OrderStatusCodesTable
	OrdersTable.ForeignKeys[4].RefTable = PersonsTable
	OrdersTable.ForeignKeys[5].RefTable = AddressesTable
	OrderHistoriesTable.ForeignKeys[0].RefTable = OrdersTable
	OrderHistoriesTable.ForeignKeys[1].RefTable = PersonsTable
	OrderHistoriesTable.ForeignKeys[2].RefTable = OrderStatusCodesTable
	OrderHistoriesTable.ForeignKeys[3].RefTable = OrderStatusCodesTable
	OrderHistoriesTable.Annotation = &entsql.Annotation{
		Table: "order_histories",
	}
	OrderItemsTable.ForeignKeys[0].RefTable = OrdersTable
	OrderItemsTable.ForeignKeys[1].RefTable = ProductInfoTable
	OrderItemsTable.ForeignKeys[2].RefTable = ProductColorTable
	OrderItemsTable.ForeignKeys[3].RefTable = WorkUnitInfoTable
	OrderItemsTable.ForeignKeys[4].RefTable = WorkUnitInfoTable
	OrderItemsTable.Annotation = &entsql.Annotation{
		Table: "order_items",
	}
	OrderStatusCodesTable.Annotation = &entsql.Annotation{
		Table: "order_status_codes",
	}
	PersonsTable.ForeignKeys[0].RefTable = WorkUnitInfoTable
	PersonAddressesTable.ForeignKeys[0].RefTable = PersonsTable
	PersonAddressesTable.ForeignKeys[1].RefTable = AddressesTable
	PersonAddressesTable.Annotation = &entsql.Annotation{
		Table: "person_addresses",
	}
	ProductColorTable.Annotation = &entsql.Annotation{
		Table: "product_color",
	}
	ProductImagesTable.ForeignKeys[0].RefTable = ProductInfoTable
	ProductImagesTable.Annotation = &entsql.Annotation{
		Table: "product_images",
	}
	ProductInfoTable.Annotation = &entsql.Annotation{
		Table: "product_info",
	}
	ProductQtyTable.ForeignKeys[0].RefTable = WorkUnitInfoTable
	ProductQtyTable.ForeignKeys[1].RefTable = ProductInfoTable
	ProductQtyTable.ForeignKeys[2].RefTable = ProductColorTable
	ProductQtyTable.Annotation = &entsql.Annotation{
		Table: "product_qty",
	}
	ProductTagsTable.ForeignKeys[0].RefTable = ProductInfoTable
	ProductTagsTable.ForeignKeys[1].RefTable = TagsTable
	ProductTagsTable.Annotation = &entsql.Annotation{
		Table: "product_tags",
	}
	ShipmentsTable.ForeignKeys[0].RefTable = OrdersTable
	ShipmentsTable.ForeignKeys[1].RefTable = InvoicesTable
	ShipmentsTable.ForeignKeys[2].RefTable = PersonsTable
	ShipmentHistoriesTable.ForeignKeys[0].RefTable = ShipmentsTable
	ShipmentHistoriesTable.ForeignKeys[1].RefTable = PersonsTable
	ShipmentHistoriesTable.ForeignKeys[2].RefTable = OrderStatusCodesTable
	ShipmentHistoriesTable.ForeignKeys[3].RefTable = OrderStatusCodesTable
	ShipmentHistoriesTable.Annotation = &entsql.Annotation{
		Table: "shipment_histories",
	}
	InvoiceLineItemsTable.ForeignKeys[0].RefTable = ShipmentsTable
	InvoiceLineItemsTable.ForeignKeys[1].RefTable = OrderItemsTable
	InvoiceLineItemsTable.Annotation = &entsql.Annotation{
		Table: "invoice_line_items",
	}
	WarehouseAssignmentsTable.ForeignKeys[0].RefTable = OrdersTable
	WarehouseAssignmentsTable.ForeignKeys[1].RefTable = WorkUnitInfoTable
	WarehouseAssignmentsTable.ForeignKeys[2].RefTable = PersonsTable
	WarehouseAssignmentsTable.Annotation = &entsql.Annotation{
		Table: "warehouse_assignments",
	}
	WorkUnitInfoTable.ForeignKeys[0].RefTable = AddressesTable
	WorkUnitInfoTable.Annotation = &entsql.Annotation{
		Table: "work_unit_info",
	}
}

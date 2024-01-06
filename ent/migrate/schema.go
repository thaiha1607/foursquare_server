// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// ConversationsColumns holds the columns for the "conversations" table.
	ConversationsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "title", Type: field.TypeString, Nullable: true},
		{Name: "user_one_id", Type: field.TypeUUID},
		{Name: "user_two_id", Type: field.TypeUUID},
	}
	// ConversationsTable holds the schema information for the "conversations" table.
	ConversationsTable = &schema.Table{
		Name:       "conversations",
		Columns:    ConversationsColumns,
		PrimaryKey: []*schema.Column{ConversationsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "conversations_users_user_one",
				Columns:    []*schema.Column{ConversationsColumns[4]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "conversations_users_user_two",
				Columns:    []*schema.Column{ConversationsColumns[5]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "conversation_user_one_id_user_two_id",
				Unique:  true,
				Columns: []*schema.Column{ConversationsColumns[4], ConversationsColumns[5]},
			},
		},
	}
	// FinancialTransactionsColumns holds the columns for the "financial_transactions" table.
	FinancialTransactionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "amount", Type: field.TypeFloat64, SchemaType: map[string]string{"mysql": "decimal(12,2)", "postgres": "numeric(12,2)"}},
		{Name: "comment", Type: field.TypeString, Nullable: true},
		{Name: "is_internal", Type: field.TypeBool, Default: false},
		{Name: "payment_method", Type: field.TypeEnum, Enums: []string{"CASH", "ELECTRONIC_FUNDS_TRANSFER", "GIFT_CARD", "CREDIT_CARD", "DEBIT_CARD", "PREPAID_CARD", "CHECK", "OTHER"}, Default: "CASH"},
		{Name: "invoice_id", Type: field.TypeUUID},
	}
	// FinancialTransactionsTable holds the schema information for the "financial_transactions" table.
	FinancialTransactionsTable = &schema.Table{
		Name:       "financial_transactions",
		Columns:    FinancialTransactionsColumns,
		PrimaryKey: []*schema.Column{FinancialTransactionsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "financial_transactions_invoices_invoice",
				Columns:    []*schema.Column{FinancialTransactionsColumns[7]},
				RefColumns: []*schema.Column{InvoicesColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// InvoicesColumns holds the columns for the "invoices" table.
	InvoicesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "total", Type: field.TypeFloat64, SchemaType: map[string]string{"mysql": "decimal(12,2)", "postgres": "numeric(12,2)"}},
		{Name: "comment", Type: field.TypeString, Nullable: true},
		{Name: "note", Type: field.TypeString, Nullable: true},
		{Name: "type", Type: field.TypeEnum, Enums: []string{"PRO_FORMA", "REGULAR", "PAST_DUE", "INTERIM", "TIMESHEET", "FINAL", "CREDIT", "DEBIT", "MIXED", "COMMERCIAL", "RECURRING", "OTHER"}, Default: "PRO_FORMA"},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"DRAFT", "ACTIVE", "SENT", "DISPUTED", "OVERDUE", "PARTIAL", "PAID", "VOID", "DEBT", "OTHER"}, Default: "DRAFT"},
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
	// InvoiceLineItemsColumns holds the columns for the "invoice_line_items" table.
	InvoiceLineItemsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "qty", Type: field.TypeFloat64, SchemaType: map[string]string{"mysql": "decimal(12,2)", "postgres": "numeric(12,2)"}},
		{Name: "total", Type: field.TypeFloat64, SchemaType: map[string]string{"mysql": "decimal(12,2)", "postgres": "numeric(12,2)"}},
		{Name: "invoice_id", Type: field.TypeUUID},
		{Name: "order_line_item_id", Type: field.TypeUUID},
	}
	// InvoiceLineItemsTable holds the schema information for the "invoice_line_items" table.
	InvoiceLineItemsTable = &schema.Table{
		Name:       "invoice_line_items",
		Columns:    InvoiceLineItemsColumns,
		PrimaryKey: []*schema.Column{InvoiceLineItemsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "invoice_line_items_invoices_invoice",
				Columns:    []*schema.Column{InvoiceLineItemsColumns[5]},
				RefColumns: []*schema.Column{InvoicesColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "invoice_line_items_order_line_items_order_line_item",
				Columns:    []*schema.Column{InvoiceLineItemsColumns[6]},
				RefColumns: []*schema.Column{OrderLineItemsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "invoicelineitem_invoice_id_order_line_item_id",
				Unique:  true,
				Columns: []*schema.Column{InvoiceLineItemsColumns[5], InvoiceLineItemsColumns[6]},
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
				Symbol:     "messages_users_sender",
				Columns:    []*schema.Column{MessagesColumns[7]},
				RefColumns: []*schema.Column{UsersColumns[0]},
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
		{Name: "type", Type: field.TypeEnum, Enums: []string{"SALE", "RETURN", "EXCHANGE", "TRANSFER", "INTERNAL", "OTHER"}, Default: "SALE"},
		{Name: "internal_note", Type: field.TypeString, Nullable: true},
		{Name: "customer_id", Type: field.TypeUUID},
		{Name: "created_by", Type: field.TypeUUID},
		{Name: "parent_order_id", Type: field.TypeUUID, Unique: true, Nullable: true},
		{Name: "status_code", Type: field.TypeInt, Default: 1},
		{Name: "management_staff_id", Type: field.TypeUUID},
		{Name: "warehouse_staff_id", Type: field.TypeUUID, Nullable: true},
		{Name: "delivery_staff_id", Type: field.TypeUUID, Nullable: true},
	}
	// OrdersTable holds the schema information for the "orders" table.
	OrdersTable = &schema.Table{
		Name:       "orders",
		Columns:    OrdersColumns,
		PrimaryKey: []*schema.Column{OrdersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "orders_users_customer",
				Columns:    []*schema.Column{OrdersColumns[7]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "orders_users_creator",
				Columns:    []*schema.Column{OrdersColumns[8]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "orders_orders_parent_order",
				Columns:    []*schema.Column{OrdersColumns[9]},
				RefColumns: []*schema.Column{OrdersColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "orders_order_status_codes_order_status",
				Columns:    []*schema.Column{OrdersColumns[10]},
				RefColumns: []*schema.Column{OrderStatusCodesColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "orders_users_management_staff",
				Columns:    []*schema.Column{OrdersColumns[11]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "orders_users_warehouse_staff",
				Columns:    []*schema.Column{OrdersColumns[12]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "orders_users_delivery_staff",
				Columns:    []*schema.Column{OrdersColumns[13]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// OrderLineItemsColumns holds the columns for the "order_line_items" table.
	OrderLineItemsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "qty", Type: field.TypeFloat64, SchemaType: map[string]string{"mysql": "decimal(12,2)", "postgres": "numeric(12,2)"}},
		{Name: "order_id", Type: field.TypeUUID},
		{Name: "product_id", Type: field.TypeUUID},
	}
	// OrderLineItemsTable holds the schema information for the "order_line_items" table.
	OrderLineItemsTable = &schema.Table{
		Name:       "order_line_items",
		Columns:    OrderLineItemsColumns,
		PrimaryKey: []*schema.Column{OrderLineItemsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "order_line_items_orders_order",
				Columns:    []*schema.Column{OrderLineItemsColumns[4]},
				RefColumns: []*schema.Column{OrdersColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "order_line_items_products_product",
				Columns:    []*schema.Column{OrderLineItemsColumns[5]},
				RefColumns: []*schema.Column{ProductsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "orderlineitem_order_id_product_id",
				Unique:  true,
				Columns: []*schema.Column{OrderLineItemsColumns[4], OrderLineItemsColumns[5]},
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
	// ProductsColumns holds the columns for the "products" table.
	ProductsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString},
		{Name: "description", Type: field.TypeString, Nullable: true},
		{Name: "year", Type: field.TypeInt},
		{Name: "price", Type: field.TypeFloat64, SchemaType: map[string]string{"mysql": "decimal(12,2)", "postgres": "numeric(12,2)"}},
		{Name: "qty", Type: field.TypeFloat64, SchemaType: map[string]string{"mysql": "decimal(12,2)", "postgres": "numeric(12,2)"}},
		{Name: "image_urls", Type: field.TypeJSON, Nullable: true},
		{Name: "unit_of_measurement", Type: field.TypeString, Nullable: true},
		{Name: "type", Type: field.TypeString, Nullable: true},
		{Name: "provider", Type: field.TypeString, Nullable: true},
	}
	// ProductsTable holds the schema information for the "products" table.
	ProductsTable = &schema.Table{
		Name:       "products",
		Columns:    ProductsColumns,
		PrimaryKey: []*schema.Column{ProductsColumns[0]},
	}
	// ProductTagsColumns holds the columns for the "product_tags" table.
	ProductTagsColumns = []*schema.Column{
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "product_id", Type: field.TypeUUID},
		{Name: "tag_id", Type: field.TypeUUID},
	}
	// ProductTagsTable holds the schema information for the "product_tags" table.
	ProductTagsTable = &schema.Table{
		Name:       "product_tags",
		Columns:    ProductTagsColumns,
		PrimaryKey: []*schema.Column{ProductTagsColumns[2], ProductTagsColumns[3]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "product_tags_products_products",
				Columns:    []*schema.Column{ProductTagsColumns[2]},
				RefColumns: []*schema.Column{ProductsColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "product_tags_tags_tags",
				Columns:    []*schema.Column{ProductTagsColumns[3]},
				RefColumns: []*schema.Column{TagsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// TagsColumns holds the columns for the "tags" table.
	TagsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "title", Type: field.TypeString},
	}
	// TagsTable holds the schema information for the "tags" table.
	TagsTable = &schema.Table{
		Name:       "tags",
		Columns:    TagsColumns,
		PrimaryKey: []*schema.Column{TagsColumns[0]},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "avatar_url", Type: field.TypeJSON, Nullable: true},
		{Name: "email", Type: field.TypeString, Unique: true},
		{Name: "last_reset", Type: field.TypeTime, Nullable: true},
		{Name: "last_verification", Type: field.TypeTime, Nullable: true},
		{Name: "name", Type: field.TypeString},
		{Name: "password_hash", Type: field.TypeString, Size: 2147483647},
		{Name: "verified", Type: field.TypeBool, Default: false},
		{Name: "phone", Type: field.TypeString, Unique: true},
		{Name: "role", Type: field.TypeEnum, Enums: []string{"ADMIN", "CUSTOMER", "WAREHOUSE", "DELIVERY", "MANAGEMENT"}, Default: "CUSTOMER"},
		{Name: "address", Type: field.TypeString, Nullable: true},
		{Name: "postal_code", Type: field.TypeString, Nullable: true},
		{Name: "other_address_info", Type: field.TypeString, Nullable: true},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		ConversationsTable,
		FinancialTransactionsTable,
		InvoicesTable,
		InvoiceLineItemsTable,
		MessagesTable,
		OrdersTable,
		OrderLineItemsTable,
		OrderStatusCodesTable,
		ProductsTable,
		ProductTagsTable,
		TagsTable,
		UsersTable,
	}
)

func init() {
	ConversationsTable.ForeignKeys[0].RefTable = UsersTable
	ConversationsTable.ForeignKeys[1].RefTable = UsersTable
	FinancialTransactionsTable.ForeignKeys[0].RefTable = InvoicesTable
	FinancialTransactionsTable.Annotation = &entsql.Annotation{
		Table: "financial_transactions",
	}
	InvoicesTable.ForeignKeys[0].RefTable = OrdersTable
	InvoiceLineItemsTable.ForeignKeys[0].RefTable = InvoicesTable
	InvoiceLineItemsTable.ForeignKeys[1].RefTable = OrderLineItemsTable
	InvoiceLineItemsTable.Annotation = &entsql.Annotation{
		Table: "invoice_line_items",
	}
	MessagesTable.ForeignKeys[0].RefTable = ConversationsTable
	MessagesTable.ForeignKeys[1].RefTable = UsersTable
	OrdersTable.ForeignKeys[0].RefTable = UsersTable
	OrdersTable.ForeignKeys[1].RefTable = UsersTable
	OrdersTable.ForeignKeys[2].RefTable = OrdersTable
	OrdersTable.ForeignKeys[3].RefTable = OrderStatusCodesTable
	OrdersTable.ForeignKeys[4].RefTable = UsersTable
	OrdersTable.ForeignKeys[5].RefTable = UsersTable
	OrdersTable.ForeignKeys[6].RefTable = UsersTable
	OrderLineItemsTable.ForeignKeys[0].RefTable = OrdersTable
	OrderLineItemsTable.ForeignKeys[1].RefTable = ProductsTable
	OrderLineItemsTable.Annotation = &entsql.Annotation{
		Table: "order_line_items",
	}
	OrderStatusCodesTable.Annotation = &entsql.Annotation{
		Table: "order_status_codes",
	}
	ProductTagsTable.ForeignKeys[0].RefTable = ProductsTable
	ProductTagsTable.ForeignKeys[1].RefTable = TagsTable
	ProductTagsTable.Annotation = &entsql.Annotation{
		Table: "product_tags",
	}
}

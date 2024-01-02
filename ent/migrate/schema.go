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
		{Name: "invoice_id", Type: field.TypeUUID},
		{Name: "type", Type: field.TypeInt},
		{Name: "payment_method", Type: field.TypeInt},
	}
	// FinancialTransactionsTable holds the schema information for the "financial_transactions" table.
	FinancialTransactionsTable = &schema.Table{
		Name:       "financial_transactions",
		Columns:    FinancialTransactionsColumns,
		PrimaryKey: []*schema.Column{FinancialTransactionsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "financial_transactions_invoices_invoice",
				Columns:    []*schema.Column{FinancialTransactionsColumns[5]},
				RefColumns: []*schema.Column{InvoicesColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "financial_transactions_transaction_types_transaction_type",
				Columns:    []*schema.Column{FinancialTransactionsColumns[6]},
				RefColumns: []*schema.Column{TransactionTypesColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "financial_transactions_payment_methods_payment",
				Columns:    []*schema.Column{FinancialTransactionsColumns[7]},
				RefColumns: []*schema.Column{PaymentMethodsColumns[0]},
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
		{Name: "order_id", Type: field.TypeUUID},
		{Name: "type", Type: field.TypeInt},
		{Name: "status_code", Type: field.TypeInt},
	}
	// InvoicesTable holds the schema information for the "invoices" table.
	InvoicesTable = &schema.Table{
		Name:       "invoices",
		Columns:    InvoicesColumns,
		PrimaryKey: []*schema.Column{InvoicesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "invoices_orders_order",
				Columns:    []*schema.Column{InvoicesColumns[4]},
				RefColumns: []*schema.Column{OrdersColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "invoices_invoice_types_invoice_type",
				Columns:    []*schema.Column{InvoicesColumns[5]},
				RefColumns: []*schema.Column{InvoiceTypesColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "invoices_invoice_status_codes_invoice_status",
				Columns:    []*schema.Column{InvoicesColumns[6]},
				RefColumns: []*schema.Column{InvoiceStatusCodesColumns[0]},
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
	// InvoiceStatusCodesColumns holds the columns for the "invoice_status_codes" table.
	InvoiceStatusCodesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "invoice_status", Type: field.TypeString, Unique: true},
	}
	// InvoiceStatusCodesTable holds the schema information for the "invoice_status_codes" table.
	InvoiceStatusCodesTable = &schema.Table{
		Name:       "invoice_status_codes",
		Columns:    InvoiceStatusCodesColumns,
		PrimaryKey: []*schema.Column{InvoiceStatusCodesColumns[0]},
	}
	// InvoiceTypesColumns holds the columns for the "invoice_types" table.
	InvoiceTypesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "invoice_type", Type: field.TypeString, Unique: true},
	}
	// InvoiceTypesTable holds the schema information for the "invoice_types" table.
	InvoiceTypesTable = &schema.Table{
		Name:       "invoice_types",
		Columns:    InvoiceTypesColumns,
		PrimaryKey: []*schema.Column{InvoiceTypesColumns[0]},
	}
	// MessagesColumns holds the columns for the "messages" table.
	MessagesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "content", Type: field.TypeString},
		{Name: "is_read", Type: field.TypeBool, Default: false},
		{Name: "conversation_id", Type: field.TypeUUID},
		{Name: "sender_id", Type: field.TypeUUID},
		{Name: "type", Type: field.TypeInt},
	}
	// MessagesTable holds the schema information for the "messages" table.
	MessagesTable = &schema.Table{
		Name:       "messages",
		Columns:    MessagesColumns,
		PrimaryKey: []*schema.Column{MessagesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "messages_conversations_conversation",
				Columns:    []*schema.Column{MessagesColumns[5]},
				RefColumns: []*schema.Column{ConversationsColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "messages_users_sender",
				Columns:    []*schema.Column{MessagesColumns[6]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "messages_message_types_message_type",
				Columns:    []*schema.Column{MessagesColumns[7]},
				RefColumns: []*schema.Column{MessageTypesColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// MessageTypesColumns holds the columns for the "message_types" table.
	MessageTypesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "message_type", Type: field.TypeString, Unique: true},
	}
	// MessageTypesTable holds the schema information for the "message_types" table.
	MessageTypesTable = &schema.Table{
		Name:       "message_types",
		Columns:    MessageTypesColumns,
		PrimaryKey: []*schema.Column{MessageTypesColumns[0]},
	}
	// OrdersColumns holds the columns for the "orders" table.
	OrdersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "note", Type: field.TypeString, Nullable: true},
		{Name: "priority", Type: field.TypeInt, Default: 0},
		{Name: "internal_note", Type: field.TypeString, Nullable: true},
		{Name: "is_internal", Type: field.TypeBool, Default: false},
		{Name: "customer_id", Type: field.TypeUUID},
		{Name: "created_by", Type: field.TypeUUID},
		{Name: "parent_order_id", Type: field.TypeUUID, Unique: true, Nullable: true},
		{Name: "status_code", Type: field.TypeInt},
		{Name: "type", Type: field.TypeInt},
		{Name: "manaagment_staff_id", Type: field.TypeUUID},
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
				Symbol:     "orders_order_types_order_type",
				Columns:    []*schema.Column{OrdersColumns[11]},
				RefColumns: []*schema.Column{OrderTypesColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "orders_users_management_staff",
				Columns:    []*schema.Column{OrdersColumns[12]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "orders_users_warehouse_staff",
				Columns:    []*schema.Column{OrdersColumns[13]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "orders_users_delivery_staff",
				Columns:    []*schema.Column{OrdersColumns[14]},
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
	// OrderTypesColumns holds the columns for the "order_types" table.
	OrderTypesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "order_type", Type: field.TypeString, Unique: true},
	}
	// OrderTypesTable holds the schema information for the "order_types" table.
	OrderTypesTable = &schema.Table{
		Name:       "order_types",
		Columns:    OrderTypesColumns,
		PrimaryKey: []*schema.Column{OrderTypesColumns[0]},
	}
	// PaymentMethodsColumns holds the columns for the "payment_methods" table.
	PaymentMethodsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "payment_method", Type: field.TypeString, Unique: true},
	}
	// PaymentMethodsTable holds the schema information for the "payment_methods" table.
	PaymentMethodsTable = &schema.Table{
		Name:       "payment_methods",
		Columns:    PaymentMethodsColumns,
		PrimaryKey: []*schema.Column{PaymentMethodsColumns[0]},
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
		{Name: "provider", Type: field.TypeString, Nullable: true},
		{Name: "type", Type: field.TypeInt},
	}
	// ProductsTable holds the schema information for the "products" table.
	ProductsTable = &schema.Table{
		Name:       "products",
		Columns:    ProductsColumns,
		PrimaryKey: []*schema.Column{ProductsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "products_product_types_product_type",
				Columns:    []*schema.Column{ProductsColumns[11]},
				RefColumns: []*schema.Column{ProductTypesColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
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
	// ProductTypesColumns holds the columns for the "product_types" table.
	ProductTypesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "product_type", Type: field.TypeString, Unique: true},
	}
	// ProductTypesTable holds the schema information for the "product_types" table.
	ProductTypesTable = &schema.Table{
		Name:       "product_types",
		Columns:    ProductTypesColumns,
		PrimaryKey: []*schema.Column{ProductTypesColumns[0]},
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
	// TransactionTypesColumns holds the columns for the "transaction_types" table.
	TransactionTypesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "transaction_type", Type: field.TypeString, Unique: true},
	}
	// TransactionTypesTable holds the schema information for the "transaction_types" table.
	TransactionTypesTable = &schema.Table{
		Name:       "transaction_types",
		Columns:    TransactionTypesColumns,
		PrimaryKey: []*schema.Column{TransactionTypesColumns[0]},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "avatar_url", Type: field.TypeJSON, Nullable: true},
		{Name: "email", Type: field.TypeString},
		{Name: "last_reset", Type: field.TypeTime, Nullable: true},
		{Name: "last_verification", Type: field.TypeTime, Nullable: true},
		{Name: "name", Type: field.TypeString},
		{Name: "password_hash", Type: field.TypeString, Size: 2147483647},
		{Name: "username", Type: field.TypeString},
		{Name: "verified", Type: field.TypeBool, Default: false},
		{Name: "phone", Type: field.TypeString, Unique: true},
		{Name: "address", Type: field.TypeString, Nullable: true},
		{Name: "postal_code", Type: field.TypeString, Nullable: true},
		{Name: "other_address_info", Type: field.TypeString, Nullable: true},
		{Name: "role", Type: field.TypeInt},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "users_user_roles_user_role",
				Columns:    []*schema.Column{UsersColumns[15]},
				RefColumns: []*schema.Column{UserRolesColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// UserRolesColumns holds the columns for the "user_roles" table.
	UserRolesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "user_role", Type: field.TypeString, Unique: true},
	}
	// UserRolesTable holds the schema information for the "user_roles" table.
	UserRolesTable = &schema.Table{
		Name:       "user_roles",
		Columns:    UserRolesColumns,
		PrimaryKey: []*schema.Column{UserRolesColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		ConversationsTable,
		FinancialTransactionsTable,
		InvoicesTable,
		InvoiceLineItemsTable,
		InvoiceStatusCodesTable,
		InvoiceTypesTable,
		MessagesTable,
		MessageTypesTable,
		OrdersTable,
		OrderLineItemsTable,
		OrderStatusCodesTable,
		OrderTypesTable,
		PaymentMethodsTable,
		ProductsTable,
		ProductTagsTable,
		ProductTypesTable,
		TagsTable,
		TransactionTypesTable,
		UsersTable,
		UserRolesTable,
	}
)

func init() {
	ConversationsTable.ForeignKeys[0].RefTable = UsersTable
	ConversationsTable.ForeignKeys[1].RefTable = UsersTable
	FinancialTransactionsTable.ForeignKeys[0].RefTable = InvoicesTable
	FinancialTransactionsTable.ForeignKeys[1].RefTable = TransactionTypesTable
	FinancialTransactionsTable.ForeignKeys[2].RefTable = PaymentMethodsTable
	FinancialTransactionsTable.Annotation = &entsql.Annotation{
		Table: "financial_transactions",
	}
	InvoicesTable.ForeignKeys[0].RefTable = OrdersTable
	InvoicesTable.ForeignKeys[1].RefTable = InvoiceTypesTable
	InvoicesTable.ForeignKeys[2].RefTable = InvoiceStatusCodesTable
	InvoiceLineItemsTable.ForeignKeys[0].RefTable = InvoicesTable
	InvoiceLineItemsTable.ForeignKeys[1].RefTable = OrderLineItemsTable
	InvoiceLineItemsTable.Annotation = &entsql.Annotation{
		Table: "invoice_line_items",
	}
	InvoiceStatusCodesTable.Annotation = &entsql.Annotation{
		Table: "invoice_status_codes",
	}
	InvoiceTypesTable.Annotation = &entsql.Annotation{
		Table: "invoice_types",
	}
	MessagesTable.ForeignKeys[0].RefTable = ConversationsTable
	MessagesTable.ForeignKeys[1].RefTable = UsersTable
	MessagesTable.ForeignKeys[2].RefTable = MessageTypesTable
	MessageTypesTable.Annotation = &entsql.Annotation{
		Table: "message_types",
	}
	OrdersTable.ForeignKeys[0].RefTable = UsersTable
	OrdersTable.ForeignKeys[1].RefTable = UsersTable
	OrdersTable.ForeignKeys[2].RefTable = OrdersTable
	OrdersTable.ForeignKeys[3].RefTable = OrderStatusCodesTable
	OrdersTable.ForeignKeys[4].RefTable = OrderTypesTable
	OrdersTable.ForeignKeys[5].RefTable = UsersTable
	OrdersTable.ForeignKeys[6].RefTable = UsersTable
	OrdersTable.ForeignKeys[7].RefTable = UsersTable
	OrderLineItemsTable.ForeignKeys[0].RefTable = OrdersTable
	OrderLineItemsTable.ForeignKeys[1].RefTable = ProductsTable
	OrderLineItemsTable.Annotation = &entsql.Annotation{
		Table: "order_line_items",
	}
	OrderStatusCodesTable.Annotation = &entsql.Annotation{
		Table: "order_status_codes",
	}
	OrderTypesTable.Annotation = &entsql.Annotation{
		Table: "order_types",
	}
	PaymentMethodsTable.Annotation = &entsql.Annotation{
		Table: "payment_methods",
	}
	ProductsTable.ForeignKeys[0].RefTable = ProductTypesTable
	ProductTagsTable.ForeignKeys[0].RefTable = ProductsTable
	ProductTagsTable.ForeignKeys[1].RefTable = TagsTable
	ProductTagsTable.Annotation = &entsql.Annotation{
		Table: "product_tags",
	}
	ProductTypesTable.Annotation = &entsql.Annotation{
		Table: "product_types",
	}
	TransactionTypesTable.Annotation = &entsql.Annotation{
		Table: "transaction_types",
	}
	UsersTable.ForeignKeys[0].RefTable = UserRolesTable
	UserRolesTable.Annotation = &entsql.Annotation{
		Table: "user_roles",
	}
}
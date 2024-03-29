// Code generated by ent, DO NOT EDIT.

package runtime

import (
	"time"

	"github.com/google/uuid"
	"github.com/thaiha1607/foursquare_server/ent/conversation"
	"github.com/thaiha1607/foursquare_server/ent/financialtransaction"
	"github.com/thaiha1607/foursquare_server/ent/invoice"
	"github.com/thaiha1607/foursquare_server/ent/invoicelineitem"
	"github.com/thaiha1607/foursquare_server/ent/message"
	"github.com/thaiha1607/foursquare_server/ent/order"
	"github.com/thaiha1607/foursquare_server/ent/orderlineitem"
	"github.com/thaiha1607/foursquare_server/ent/orderstatuscode"
	"github.com/thaiha1607/foursquare_server/ent/person"
	"github.com/thaiha1607/foursquare_server/ent/product"
	"github.com/thaiha1607/foursquare_server/ent/productimage"
	"github.com/thaiha1607/foursquare_server/ent/schema"
	"github.com/thaiha1607/foursquare_server/ent/tag"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	conversationMixin := schema.Conversation{}.Mixin()
	conversationMixinFields0 := conversationMixin[0].Fields()
	_ = conversationMixinFields0
	conversationFields := schema.Conversation{}.Fields()
	_ = conversationFields
	// conversationDescCreatedAt is the schema descriptor for created_at field.
	conversationDescCreatedAt := conversationMixinFields0[0].Descriptor()
	// conversation.DefaultCreatedAt holds the default value on creation for the created_at field.
	conversation.DefaultCreatedAt = conversationDescCreatedAt.Default.(func() time.Time)
	// conversationDescUpdatedAt is the schema descriptor for updated_at field.
	conversationDescUpdatedAt := conversationMixinFields0[1].Descriptor()
	// conversation.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	conversation.DefaultUpdatedAt = conversationDescUpdatedAt.Default.(func() time.Time)
	// conversation.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	conversation.UpdateDefaultUpdatedAt = conversationDescUpdatedAt.UpdateDefault.(func() time.Time)
	// conversationDescID is the schema descriptor for id field.
	conversationDescID := conversationFields[0].Descriptor()
	// conversation.DefaultID holds the default value on creation for the id field.
	conversation.DefaultID = conversationDescID.Default.(func() uuid.UUID)
	financialtransactionMixin := schema.FinancialTransaction{}.Mixin()
	financialtransactionMixinFields0 := financialtransactionMixin[0].Fields()
	_ = financialtransactionMixinFields0
	financialtransactionFields := schema.FinancialTransaction{}.Fields()
	_ = financialtransactionFields
	// financialtransactionDescCreatedAt is the schema descriptor for created_at field.
	financialtransactionDescCreatedAt := financialtransactionMixinFields0[0].Descriptor()
	// financialtransaction.DefaultCreatedAt holds the default value on creation for the created_at field.
	financialtransaction.DefaultCreatedAt = financialtransactionDescCreatedAt.Default.(func() time.Time)
	// financialtransactionDescUpdatedAt is the schema descriptor for updated_at field.
	financialtransactionDescUpdatedAt := financialtransactionMixinFields0[1].Descriptor()
	// financialtransaction.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	financialtransaction.DefaultUpdatedAt = financialtransactionDescUpdatedAt.Default.(func() time.Time)
	// financialtransaction.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	financialtransaction.UpdateDefaultUpdatedAt = financialtransactionDescUpdatedAt.UpdateDefault.(func() time.Time)
	// financialtransactionDescIsInternal is the schema descriptor for is_internal field.
	financialtransactionDescIsInternal := financialtransactionFields[4].Descriptor()
	// financialtransaction.DefaultIsInternal holds the default value on creation for the is_internal field.
	financialtransaction.DefaultIsInternal = financialtransactionDescIsInternal.Default.(bool)
	// financialtransactionDescID is the schema descriptor for id field.
	financialtransactionDescID := financialtransactionFields[0].Descriptor()
	// financialtransaction.DefaultID holds the default value on creation for the id field.
	financialtransaction.DefaultID = financialtransactionDescID.Default.(func() uuid.UUID)
	invoiceMixin := schema.Invoice{}.Mixin()
	invoiceMixinFields0 := invoiceMixin[0].Fields()
	_ = invoiceMixinFields0
	invoiceFields := schema.Invoice{}.Fields()
	_ = invoiceFields
	// invoiceDescCreatedAt is the schema descriptor for created_at field.
	invoiceDescCreatedAt := invoiceMixinFields0[0].Descriptor()
	// invoice.DefaultCreatedAt holds the default value on creation for the created_at field.
	invoice.DefaultCreatedAt = invoiceDescCreatedAt.Default.(func() time.Time)
	// invoiceDescUpdatedAt is the schema descriptor for updated_at field.
	invoiceDescUpdatedAt := invoiceMixinFields0[1].Descriptor()
	// invoice.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	invoice.DefaultUpdatedAt = invoiceDescUpdatedAt.Default.(func() time.Time)
	// invoice.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	invoice.UpdateDefaultUpdatedAt = invoiceDescUpdatedAt.UpdateDefault.(func() time.Time)
	// invoiceDescID is the schema descriptor for id field.
	invoiceDescID := invoiceFields[0].Descriptor()
	// invoice.DefaultID holds the default value on creation for the id field.
	invoice.DefaultID = invoiceDescID.Default.(func() uuid.UUID)
	invoicelineitemFields := schema.InvoiceLineItem{}.Fields()
	_ = invoicelineitemFields
	// invoicelineitemDescCreatedAt is the schema descriptor for created_at field.
	invoicelineitemDescCreatedAt := invoicelineitemFields[5].Descriptor()
	// invoicelineitem.DefaultCreatedAt holds the default value on creation for the created_at field.
	invoicelineitem.DefaultCreatedAt = invoicelineitemDescCreatedAt.Default.(func() time.Time)
	// invoicelineitemDescID is the schema descriptor for id field.
	invoicelineitemDescID := invoicelineitemFields[0].Descriptor()
	// invoicelineitem.DefaultID holds the default value on creation for the id field.
	invoicelineitem.DefaultID = invoicelineitemDescID.Default.(func() uuid.UUID)
	messageMixin := schema.Message{}.Mixin()
	messageMixinFields0 := messageMixin[0].Fields()
	_ = messageMixinFields0
	messageFields := schema.Message{}.Fields()
	_ = messageFields
	// messageDescCreatedAt is the schema descriptor for created_at field.
	messageDescCreatedAt := messageMixinFields0[0].Descriptor()
	// message.DefaultCreatedAt holds the default value on creation for the created_at field.
	message.DefaultCreatedAt = messageDescCreatedAt.Default.(func() time.Time)
	// messageDescUpdatedAt is the schema descriptor for updated_at field.
	messageDescUpdatedAt := messageMixinFields0[1].Descriptor()
	// message.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	message.DefaultUpdatedAt = messageDescUpdatedAt.Default.(func() time.Time)
	// message.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	message.UpdateDefaultUpdatedAt = messageDescUpdatedAt.UpdateDefault.(func() time.Time)
	// messageDescContent is the schema descriptor for content field.
	messageDescContent := messageFields[4].Descriptor()
	// message.ContentValidator is a validator for the "content" field. It is called by the builders before save.
	message.ContentValidator = messageDescContent.Validators[0].(func(string) error)
	// messageDescIsRead is the schema descriptor for is_read field.
	messageDescIsRead := messageFields[5].Descriptor()
	// message.DefaultIsRead holds the default value on creation for the is_read field.
	message.DefaultIsRead = messageDescIsRead.Default.(bool)
	// messageDescID is the schema descriptor for id field.
	messageDescID := messageFields[0].Descriptor()
	// message.DefaultID holds the default value on creation for the id field.
	message.DefaultID = messageDescID.Default.(func() uuid.UUID)
	orderMixin := schema.Order{}.Mixin()
	orderMixinFields0 := orderMixin[0].Fields()
	_ = orderMixinFields0
	orderFields := schema.Order{}.Fields()
	_ = orderFields
	// orderDescCreatedAt is the schema descriptor for created_at field.
	orderDescCreatedAt := orderMixinFields0[0].Descriptor()
	// order.DefaultCreatedAt holds the default value on creation for the created_at field.
	order.DefaultCreatedAt = orderDescCreatedAt.Default.(func() time.Time)
	// orderDescUpdatedAt is the schema descriptor for updated_at field.
	orderDescUpdatedAt := orderMixinFields0[1].Descriptor()
	// order.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	order.DefaultUpdatedAt = orderDescUpdatedAt.Default.(func() time.Time)
	// order.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	order.UpdateDefaultUpdatedAt = orderDescUpdatedAt.UpdateDefault.(func() time.Time)
	// orderDescPriority is the schema descriptor for priority field.
	orderDescPriority := orderFields[5].Descriptor()
	// order.DefaultPriority holds the default value on creation for the priority field.
	order.DefaultPriority = orderDescPriority.Default.(int)
	// order.PriorityValidator is a validator for the "priority" field. It is called by the builders before save.
	order.PriorityValidator = orderDescPriority.Validators[0].(func(int) error)
	// orderDescStatusCode is the schema descriptor for status_code field.
	orderDescStatusCode := orderFields[7].Descriptor()
	// order.DefaultStatusCode holds the default value on creation for the status_code field.
	order.DefaultStatusCode = orderDescStatusCode.Default.(int)
	// orderDescID is the schema descriptor for id field.
	orderDescID := orderFields[0].Descriptor()
	// order.DefaultID holds the default value on creation for the id field.
	order.DefaultID = orderDescID.Default.(func() uuid.UUID)
	orderlineitemMixin := schema.OrderLineItem{}.Mixin()
	orderlineitemMixinFields0 := orderlineitemMixin[0].Fields()
	_ = orderlineitemMixinFields0
	orderlineitemFields := schema.OrderLineItem{}.Fields()
	_ = orderlineitemFields
	// orderlineitemDescCreatedAt is the schema descriptor for created_at field.
	orderlineitemDescCreatedAt := orderlineitemMixinFields0[0].Descriptor()
	// orderlineitem.DefaultCreatedAt holds the default value on creation for the created_at field.
	orderlineitem.DefaultCreatedAt = orderlineitemDescCreatedAt.Default.(func() time.Time)
	// orderlineitemDescUpdatedAt is the schema descriptor for updated_at field.
	orderlineitemDescUpdatedAt := orderlineitemMixinFields0[1].Descriptor()
	// orderlineitem.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	orderlineitem.DefaultUpdatedAt = orderlineitemDescUpdatedAt.Default.(func() time.Time)
	// orderlineitem.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	orderlineitem.UpdateDefaultUpdatedAt = orderlineitemDescUpdatedAt.UpdateDefault.(func() time.Time)
	// orderlineitemDescID is the schema descriptor for id field.
	orderlineitemDescID := orderlineitemFields[0].Descriptor()
	// orderlineitem.DefaultID holds the default value on creation for the id field.
	orderlineitem.DefaultID = orderlineitemDescID.Default.(func() uuid.UUID)
	orderstatuscodeMixin := schema.OrderStatusCode{}.Mixin()
	orderstatuscodeMixinFields0 := orderstatuscodeMixin[0].Fields()
	_ = orderstatuscodeMixinFields0
	orderstatuscodeFields := schema.OrderStatusCode{}.Fields()
	_ = orderstatuscodeFields
	// orderstatuscodeDescCreatedAt is the schema descriptor for created_at field.
	orderstatuscodeDescCreatedAt := orderstatuscodeMixinFields0[0].Descriptor()
	// orderstatuscode.DefaultCreatedAt holds the default value on creation for the created_at field.
	orderstatuscode.DefaultCreatedAt = orderstatuscodeDescCreatedAt.Default.(func() time.Time)
	// orderstatuscodeDescUpdatedAt is the schema descriptor for updated_at field.
	orderstatuscodeDescUpdatedAt := orderstatuscodeMixinFields0[1].Descriptor()
	// orderstatuscode.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	orderstatuscode.DefaultUpdatedAt = orderstatuscodeDescUpdatedAt.Default.(func() time.Time)
	// orderstatuscode.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	orderstatuscode.UpdateDefaultUpdatedAt = orderstatuscodeDescUpdatedAt.UpdateDefault.(func() time.Time)
	// orderstatuscodeDescOrderStatus is the schema descriptor for order_status field.
	orderstatuscodeDescOrderStatus := orderstatuscodeFields[1].Descriptor()
	// orderstatuscode.OrderStatusValidator is a validator for the "order_status" field. It is called by the builders before save.
	orderstatuscode.OrderStatusValidator = orderstatuscodeDescOrderStatus.Validators[0].(func(string) error)
	// orderstatuscodeDescID is the schema descriptor for id field.
	orderstatuscodeDescID := orderstatuscodeFields[0].Descriptor()
	// orderstatuscode.IDValidator is a validator for the "id" field. It is called by the builders before save.
	orderstatuscode.IDValidator = func() func(int) error {
		validators := orderstatuscodeDescID.Validators
		fns := [...]func(int) error{
			validators[0].(func(int) error),
			validators[1].(func(int) error),
		}
		return func(id int) error {
			for _, fn := range fns {
				if err := fn(id); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	personMixin := schema.Person{}.Mixin()
	personMixinFields0 := personMixin[0].Fields()
	_ = personMixinFields0
	personFields := schema.Person{}.Fields()
	_ = personFields
	// personDescCreatedAt is the schema descriptor for created_at field.
	personDescCreatedAt := personMixinFields0[0].Descriptor()
	// person.DefaultCreatedAt holds the default value on creation for the created_at field.
	person.DefaultCreatedAt = personDescCreatedAt.Default.(func() time.Time)
	// personDescUpdatedAt is the schema descriptor for updated_at field.
	personDescUpdatedAt := personMixinFields0[1].Descriptor()
	// person.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	person.DefaultUpdatedAt = personDescUpdatedAt.Default.(func() time.Time)
	// person.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	person.UpdateDefaultUpdatedAt = personDescUpdatedAt.UpdateDefault.(func() time.Time)
	// personDescAvatarURL is the schema descriptor for avatar_url field.
	personDescAvatarURL := personFields[0].Descriptor()
	// person.AvatarURLValidator is a validator for the "avatar_url" field. It is called by the builders before save.
	person.AvatarURLValidator = personDescAvatarURL.Validators[0].(func(string) error)
	// personDescEmail is the schema descriptor for email field.
	personDescEmail := personFields[1].Descriptor()
	// person.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	person.EmailValidator = personDescEmail.Validators[0].(func(string) error)
	// personDescName is the schema descriptor for name field.
	personDescName := personFields[3].Descriptor()
	// person.NameValidator is a validator for the "name" field. It is called by the builders before save.
	person.NameValidator = personDescName.Validators[0].(func(string) error)
	// personDescPhone is the schema descriptor for phone field.
	personDescPhone := personFields[4].Descriptor()
	// person.PhoneValidator is a validator for the "phone" field. It is called by the builders before save.
	person.PhoneValidator = personDescPhone.Validators[0].(func(string) error)
	// personDescPasswordHash is the schema descriptor for password_hash field.
	personDescPasswordHash := personFields[6].Descriptor()
	// person.PasswordHashValidator is a validator for the "password_hash" field. It is called by the builders before save.
	person.PasswordHashValidator = personDescPasswordHash.Validators[0].(func([]byte) error)
	// personDescIsEmailVerified is the schema descriptor for is_email_verified field.
	personDescIsEmailVerified := personFields[7].Descriptor()
	// person.DefaultIsEmailVerified holds the default value on creation for the is_email_verified field.
	person.DefaultIsEmailVerified = personDescIsEmailVerified.Default.(bool)
	// personDescIsPhoneVerified is the schema descriptor for is_phone_verified field.
	personDescIsPhoneVerified := personFields[8].Descriptor()
	// person.DefaultIsPhoneVerified holds the default value on creation for the is_phone_verified field.
	person.DefaultIsPhoneVerified = personDescIsPhoneVerified.Default.(bool)
	// personDescID is the schema descriptor for id field.
	personDescID := personFields[2].Descriptor()
	// person.DefaultID holds the default value on creation for the id field.
	person.DefaultID = personDescID.Default.(func() uuid.UUID)
	productMixin := schema.Product{}.Mixin()
	productHooks := schema.Product{}.Hooks()
	product.Hooks[0] = productHooks[0]
	productMixinFields0 := productMixin[0].Fields()
	_ = productMixinFields0
	productFields := schema.Product{}.Fields()
	_ = productFields
	// productDescCreatedAt is the schema descriptor for created_at field.
	productDescCreatedAt := productMixinFields0[0].Descriptor()
	// product.DefaultCreatedAt holds the default value on creation for the created_at field.
	product.DefaultCreatedAt = productDescCreatedAt.Default.(func() time.Time)
	// productDescUpdatedAt is the schema descriptor for updated_at field.
	productDescUpdatedAt := productMixinFields0[1].Descriptor()
	// product.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	product.DefaultUpdatedAt = productDescUpdatedAt.Default.(func() time.Time)
	// product.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	product.UpdateDefaultUpdatedAt = productDescUpdatedAt.UpdateDefault.(func() time.Time)
	// productDescName is the schema descriptor for name field.
	productDescName := productFields[0].Descriptor()
	// product.NameValidator is a validator for the "name" field. It is called by the builders before save.
	product.NameValidator = productDescName.Validators[0].(func(string) error)
	// productDescDescription is the schema descriptor for description field.
	productDescDescription := productFields[1].Descriptor()
	// product.DefaultDescription holds the default value on creation for the description field.
	product.DefaultDescription = productDescDescription.Default.(string)
	// productDescYear is the schema descriptor for year field.
	productDescYear := productFields[3].Descriptor()
	// product.YearValidator is a validator for the "year" field. It is called by the builders before save.
	product.YearValidator = func() func(int) error {
		validators := productDescYear.Validators
		fns := [...]func(int) error{
			validators[0].(func(int) error),
			validators[1].(func(int) error),
		}
		return func(year int) error {
			for _, fn := range fns {
				if err := fn(year); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// productDescID is the schema descriptor for id field.
	productDescID := productFields[2].Descriptor()
	// product.IDValidator is a validator for the "id" field. It is called by the builders before save.
	product.IDValidator = func() func(string) error {
		validators := productDescID.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(id string) error {
			for _, fn := range fns {
				if err := fn(id); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	productimageMixin := schema.ProductImage{}.Mixin()
	productimageMixinFields0 := productimageMixin[0].Fields()
	_ = productimageMixinFields0
	productimageFields := schema.ProductImage{}.Fields()
	_ = productimageFields
	// productimageDescCreatedAt is the schema descriptor for created_at field.
	productimageDescCreatedAt := productimageMixinFields0[0].Descriptor()
	// productimage.DefaultCreatedAt holds the default value on creation for the created_at field.
	productimage.DefaultCreatedAt = productimageDescCreatedAt.Default.(func() time.Time)
	// productimageDescUpdatedAt is the schema descriptor for updated_at field.
	productimageDescUpdatedAt := productimageMixinFields0[1].Descriptor()
	// productimage.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	productimage.DefaultUpdatedAt = productimageDescUpdatedAt.Default.(func() time.Time)
	// productimage.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	productimage.UpdateDefaultUpdatedAt = productimageDescUpdatedAt.UpdateDefault.(func() time.Time)
	// productimageDescImageURL is the schema descriptor for image_url field.
	productimageDescImageURL := productimageFields[2].Descriptor()
	// productimage.ImageURLValidator is a validator for the "image_url" field. It is called by the builders before save.
	productimage.ImageURLValidator = productimageDescImageURL.Validators[0].(func(string) error)
	// productimageDescID is the schema descriptor for id field.
	productimageDescID := productimageFields[0].Descriptor()
	// productimage.DefaultID holds the default value on creation for the id field.
	productimage.DefaultID = productimageDescID.Default.(func() uuid.UUID)
	tagMixin := schema.Tag{}.Mixin()
	tagHooks := schema.Tag{}.Hooks()
	tag.Hooks[0] = tagHooks[0]
	tagMixinFields0 := tagMixin[0].Fields()
	_ = tagMixinFields0
	tagFields := schema.Tag{}.Fields()
	_ = tagFields
	// tagDescCreatedAt is the schema descriptor for created_at field.
	tagDescCreatedAt := tagMixinFields0[0].Descriptor()
	// tag.DefaultCreatedAt holds the default value on creation for the created_at field.
	tag.DefaultCreatedAt = tagDescCreatedAt.Default.(func() time.Time)
	// tagDescUpdatedAt is the schema descriptor for updated_at field.
	tagDescUpdatedAt := tagMixinFields0[1].Descriptor()
	// tag.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	tag.DefaultUpdatedAt = tagDescUpdatedAt.Default.(func() time.Time)
	// tag.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	tag.UpdateDefaultUpdatedAt = tagDescUpdatedAt.UpdateDefault.(func() time.Time)
	// tagDescTitle is the schema descriptor for title field.
	tagDescTitle := tagFields[1].Descriptor()
	// tag.TitleValidator is a validator for the "title" field. It is called by the builders before save.
	tag.TitleValidator = tagDescTitle.Validators[0].(func(string) error)
	// tagDescID is the schema descriptor for id field.
	tagDescID := tagFields[0].Descriptor()
	// tag.IDValidator is a validator for the "id" field. It is called by the builders before save.
	tag.IDValidator = func() func(string) error {
		validators := tagDescID.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(id string) error {
			for _, fn := range fns {
				if err := fn(id); err != nil {
					return err
				}
			}
			return nil
		}
	}()
}

const (
	Version = "v0.13.1"                                         // Version of ent codegen.
	Sum     = "h1:uD8QwN1h6SNphdCCzmkMN3feSUzNnVvV/WIkHKMbzOE=" // Sum of ent codegen.
)

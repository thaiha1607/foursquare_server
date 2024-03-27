// Code generated by ent, DO NOT EDIT.

package runtime

import (
	"time"

	"github.com/google/uuid"
	"github.com/thaiha1607/foursquare_server/ent/address"
	"github.com/thaiha1607/foursquare_server/ent/conversation"
	"github.com/thaiha1607/foursquare_server/ent/deliveryassignment"
	"github.com/thaiha1607/foursquare_server/ent/invoice"
	"github.com/thaiha1607/foursquare_server/ent/message"
	"github.com/thaiha1607/foursquare_server/ent/order"
	"github.com/thaiha1607/foursquare_server/ent/orderitem"
	"github.com/thaiha1607/foursquare_server/ent/orderstatuscode"
	"github.com/thaiha1607/foursquare_server/ent/person"
	"github.com/thaiha1607/foursquare_server/ent/personaddress"
	"github.com/thaiha1607/foursquare_server/ent/productcolor"
	"github.com/thaiha1607/foursquare_server/ent/productimage"
	"github.com/thaiha1607/foursquare_server/ent/productinfo"
	"github.com/thaiha1607/foursquare_server/ent/productqty"
	"github.com/thaiha1607/foursquare_server/ent/schema"
	"github.com/thaiha1607/foursquare_server/ent/shipment"
	"github.com/thaiha1607/foursquare_server/ent/shipmentitem"
	"github.com/thaiha1607/foursquare_server/ent/tag"
	"github.com/thaiha1607/foursquare_server/ent/warehouseassignment"
	"github.com/thaiha1607/foursquare_server/ent/workunitinfo"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	addressMixin := schema.Address{}.Mixin()
	addressHooks := schema.Address{}.Hooks()
	address.Hooks[0] = addressHooks[0]
	addressMixinFields0 := addressMixin[0].Fields()
	_ = addressMixinFields0
	addressFields := schema.Address{}.Fields()
	_ = addressFields
	// addressDescCreatedAt is the schema descriptor for created_at field.
	addressDescCreatedAt := addressMixinFields0[0].Descriptor()
	// address.DefaultCreatedAt holds the default value on creation for the created_at field.
	address.DefaultCreatedAt = addressDescCreatedAt.Default.(func() time.Time)
	// addressDescUpdatedAt is the schema descriptor for updated_at field.
	addressDescUpdatedAt := addressMixinFields0[1].Descriptor()
	// address.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	address.DefaultUpdatedAt = addressDescUpdatedAt.Default.(func() time.Time)
	// address.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	address.UpdateDefaultUpdatedAt = addressDescUpdatedAt.UpdateDefault.(func() time.Time)
	// addressDescLine1 is the schema descriptor for line1 field.
	addressDescLine1 := addressFields[1].Descriptor()
	// address.Line1Validator is a validator for the "line1" field. It is called by the builders before save.
	address.Line1Validator = addressDescLine1.Validators[0].(func(string) error)
	// addressDescCity is the schema descriptor for city field.
	addressDescCity := addressFields[3].Descriptor()
	// address.CityValidator is a validator for the "city" field. It is called by the builders before save.
	address.CityValidator = addressDescCity.Validators[0].(func(string) error)
	// addressDescZipOrPostcode is the schema descriptor for zip_or_postcode field.
	addressDescZipOrPostcode := addressFields[5].Descriptor()
	// address.ZipOrPostcodeValidator is a validator for the "zip_or_postcode" field. It is called by the builders before save.
	address.ZipOrPostcodeValidator = addressDescZipOrPostcode.Validators[0].(func(string) error)
	// addressDescCountry is the schema descriptor for country field.
	addressDescCountry := addressFields[6].Descriptor()
	// address.CountryValidator is a validator for the "country" field. It is called by the builders before save.
	address.CountryValidator = addressDescCountry.Validators[0].(func(string) error)
	// addressDescID is the schema descriptor for id field.
	addressDescID := addressFields[0].Descriptor()
	// address.IDValidator is a validator for the "id" field. It is called by the builders before save.
	address.IDValidator = addressDescID.Validators[0].(func(string) error)
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
	deliveryassignmentMixin := schema.DeliveryAssignment{}.Mixin()
	deliveryassignmentMixinFields0 := deliveryassignmentMixin[0].Fields()
	_ = deliveryassignmentMixinFields0
	deliveryassignmentFields := schema.DeliveryAssignment{}.Fields()
	_ = deliveryassignmentFields
	// deliveryassignmentDescCreatedAt is the schema descriptor for created_at field.
	deliveryassignmentDescCreatedAt := deliveryassignmentMixinFields0[0].Descriptor()
	// deliveryassignment.DefaultCreatedAt holds the default value on creation for the created_at field.
	deliveryassignment.DefaultCreatedAt = deliveryassignmentDescCreatedAt.Default.(func() time.Time)
	// deliveryassignmentDescUpdatedAt is the schema descriptor for updated_at field.
	deliveryassignmentDescUpdatedAt := deliveryassignmentMixinFields0[1].Descriptor()
	// deliveryassignment.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	deliveryassignment.DefaultUpdatedAt = deliveryassignmentDescUpdatedAt.Default.(func() time.Time)
	// deliveryassignment.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	deliveryassignment.UpdateDefaultUpdatedAt = deliveryassignmentDescUpdatedAt.UpdateDefault.(func() time.Time)
	// deliveryassignmentDescID is the schema descriptor for id field.
	deliveryassignmentDescID := deliveryassignmentFields[0].Descriptor()
	// deliveryassignment.DefaultID holds the default value on creation for the id field.
	deliveryassignment.DefaultID = deliveryassignmentDescID.Default.(func() uuid.UUID)
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
	// orderDescIsInternal is the schema descriptor for is_internal field.
	orderDescIsInternal := orderFields[10].Descriptor()
	// order.DefaultIsInternal holds the default value on creation for the is_internal field.
	order.DefaultIsInternal = orderDescIsInternal.Default.(bool)
	// orderDescAddressID is the schema descriptor for address_id field.
	orderDescAddressID := orderFields[11].Descriptor()
	// order.AddressIDValidator is a validator for the "address_id" field. It is called by the builders before save.
	order.AddressIDValidator = orderDescAddressID.Validators[0].(func(string) error)
	// orderDescID is the schema descriptor for id field.
	orderDescID := orderFields[0].Descriptor()
	// order.DefaultID holds the default value on creation for the id field.
	order.DefaultID = orderDescID.Default.(func() uuid.UUID)
	orderitemMixin := schema.OrderItem{}.Mixin()
	orderitemMixinFields0 := orderitemMixin[0].Fields()
	_ = orderitemMixinFields0
	orderitemFields := schema.OrderItem{}.Fields()
	_ = orderitemFields
	// orderitemDescCreatedAt is the schema descriptor for created_at field.
	orderitemDescCreatedAt := orderitemMixinFields0[0].Descriptor()
	// orderitem.DefaultCreatedAt holds the default value on creation for the created_at field.
	orderitem.DefaultCreatedAt = orderitemDescCreatedAt.Default.(func() time.Time)
	// orderitemDescUpdatedAt is the schema descriptor for updated_at field.
	orderitemDescUpdatedAt := orderitemMixinFields0[1].Descriptor()
	// orderitem.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	orderitem.DefaultUpdatedAt = orderitemDescUpdatedAt.Default.(func() time.Time)
	// orderitem.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	orderitem.UpdateDefaultUpdatedAt = orderitemDescUpdatedAt.UpdateDefault.(func() time.Time)
	// orderitemDescID is the schema descriptor for id field.
	orderitemDescID := orderitemFields[0].Descriptor()
	// orderitem.DefaultID holds the default value on creation for the id field.
	orderitem.DefaultID = orderitemDescID.Default.(func() uuid.UUID)
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
	person.EmailValidator = func() func(string) error {
		validators := personDescEmail.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(email string) error {
			for _, fn := range fns {
				if err := fn(email); err != nil {
					return err
				}
			}
			return nil
		}
	}()
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
	personaddressMixin := schema.PersonAddress{}.Mixin()
	personaddressMixinFields0 := personaddressMixin[0].Fields()
	_ = personaddressMixinFields0
	personaddressFields := schema.PersonAddress{}.Fields()
	_ = personaddressFields
	// personaddressDescCreatedAt is the schema descriptor for created_at field.
	personaddressDescCreatedAt := personaddressMixinFields0[0].Descriptor()
	// personaddress.DefaultCreatedAt holds the default value on creation for the created_at field.
	personaddress.DefaultCreatedAt = personaddressDescCreatedAt.Default.(func() time.Time)
	// personaddressDescUpdatedAt is the schema descriptor for updated_at field.
	personaddressDescUpdatedAt := personaddressMixinFields0[1].Descriptor()
	// personaddress.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	personaddress.DefaultUpdatedAt = personaddressDescUpdatedAt.Default.(func() time.Time)
	// personaddress.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	personaddress.UpdateDefaultUpdatedAt = personaddressDescUpdatedAt.UpdateDefault.(func() time.Time)
	// personaddressDescAddressID is the schema descriptor for address_id field.
	personaddressDescAddressID := personaddressFields[1].Descriptor()
	// personaddress.AddressIDValidator is a validator for the "address_id" field. It is called by the builders before save.
	personaddress.AddressIDValidator = personaddressDescAddressID.Validators[0].(func(string) error)
	productcolorMixin := schema.ProductColor{}.Mixin()
	productcolorHooks := schema.ProductColor{}.Hooks()
	productcolor.Hooks[0] = productcolorHooks[0]
	productcolorMixinFields0 := productcolorMixin[0].Fields()
	_ = productcolorMixinFields0
	productcolorFields := schema.ProductColor{}.Fields()
	_ = productcolorFields
	// productcolorDescCreatedAt is the schema descriptor for created_at field.
	productcolorDescCreatedAt := productcolorMixinFields0[0].Descriptor()
	// productcolor.DefaultCreatedAt holds the default value on creation for the created_at field.
	productcolor.DefaultCreatedAt = productcolorDescCreatedAt.Default.(func() time.Time)
	// productcolorDescUpdatedAt is the schema descriptor for updated_at field.
	productcolorDescUpdatedAt := productcolorMixinFields0[1].Descriptor()
	// productcolor.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	productcolor.DefaultUpdatedAt = productcolorDescUpdatedAt.Default.(func() time.Time)
	// productcolor.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	productcolor.UpdateDefaultUpdatedAt = productcolorDescUpdatedAt.UpdateDefault.(func() time.Time)
	// productcolorDescName is the schema descriptor for name field.
	productcolorDescName := productcolorFields[1].Descriptor()
	// productcolor.NameValidator is a validator for the "name" field. It is called by the builders before save.
	productcolor.NameValidator = productcolorDescName.Validators[0].(func(string) error)
	// productcolorDescColorCode is the schema descriptor for color_code field.
	productcolorDescColorCode := productcolorFields[2].Descriptor()
	// productcolor.ColorCodeValidator is a validator for the "color_code" field. It is called by the builders before save.
	productcolor.ColorCodeValidator = productcolorDescColorCode.Validators[0].(func(string) error)
	// productcolorDescID is the schema descriptor for id field.
	productcolorDescID := productcolorFields[0].Descriptor()
	// productcolor.IDValidator is a validator for the "id" field. It is called by the builders before save.
	productcolor.IDValidator = func() func(string) error {
		validators := productcolorDescID.Validators
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
	productinfoMixin := schema.ProductInfo{}.Mixin()
	productinfoHooks := schema.ProductInfo{}.Hooks()
	productinfo.Hooks[0] = productinfoHooks[0]
	productinfoMixinFields0 := productinfoMixin[0].Fields()
	_ = productinfoMixinFields0
	productinfoFields := schema.ProductInfo{}.Fields()
	_ = productinfoFields
	// productinfoDescCreatedAt is the schema descriptor for created_at field.
	productinfoDescCreatedAt := productinfoMixinFields0[0].Descriptor()
	// productinfo.DefaultCreatedAt holds the default value on creation for the created_at field.
	productinfo.DefaultCreatedAt = productinfoDescCreatedAt.Default.(func() time.Time)
	// productinfoDescUpdatedAt is the schema descriptor for updated_at field.
	productinfoDescUpdatedAt := productinfoMixinFields0[1].Descriptor()
	// productinfo.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	productinfo.DefaultUpdatedAt = productinfoDescUpdatedAt.Default.(func() time.Time)
	// productinfo.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	productinfo.UpdateDefaultUpdatedAt = productinfoDescUpdatedAt.UpdateDefault.(func() time.Time)
	// productinfoDescName is the schema descriptor for name field.
	productinfoDescName := productinfoFields[1].Descriptor()
	// productinfo.NameValidator is a validator for the "name" field. It is called by the builders before save.
	productinfo.NameValidator = productinfoDescName.Validators[0].(func(string) error)
	// productinfoDescDescription is the schema descriptor for description field.
	productinfoDescDescription := productinfoFields[2].Descriptor()
	// productinfo.DefaultDescription holds the default value on creation for the description field.
	productinfo.DefaultDescription = productinfoDescDescription.Default.(string)
	// productinfoDescYear is the schema descriptor for year field.
	productinfoDescYear := productinfoFields[3].Descriptor()
	// productinfo.YearValidator is a validator for the "year" field. It is called by the builders before save.
	productinfo.YearValidator = func() func(int) error {
		validators := productinfoDescYear.Validators
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
	// productinfoDescID is the schema descriptor for id field.
	productinfoDescID := productinfoFields[0].Descriptor()
	// productinfo.IDValidator is a validator for the "id" field. It is called by the builders before save.
	productinfo.IDValidator = func() func(string) error {
		validators := productinfoDescID.Validators
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
	productqtyMixin := schema.ProductQty{}.Mixin()
	productqtyMixinFields0 := productqtyMixin[0].Fields()
	_ = productqtyMixinFields0
	productqtyFields := schema.ProductQty{}.Fields()
	_ = productqtyFields
	// productqtyDescCreatedAt is the schema descriptor for created_at field.
	productqtyDescCreatedAt := productqtyMixinFields0[0].Descriptor()
	// productqty.DefaultCreatedAt holds the default value on creation for the created_at field.
	productqty.DefaultCreatedAt = productqtyDescCreatedAt.Default.(func() time.Time)
	// productqtyDescUpdatedAt is the schema descriptor for updated_at field.
	productqtyDescUpdatedAt := productqtyMixinFields0[1].Descriptor()
	// productqty.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	productqty.DefaultUpdatedAt = productqtyDescUpdatedAt.Default.(func() time.Time)
	// productqty.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	productqty.UpdateDefaultUpdatedAt = productqtyDescUpdatedAt.UpdateDefault.(func() time.Time)
	// productqtyDescProductID is the schema descriptor for product_id field.
	productqtyDescProductID := productqtyFields[2].Descriptor()
	// productqty.ProductIDValidator is a validator for the "product_id" field. It is called by the builders before save.
	productqty.ProductIDValidator = productqtyDescProductID.Validators[0].(func(string) error)
	// productqtyDescProductColorID is the schema descriptor for product_color_id field.
	productqtyDescProductColorID := productqtyFields[3].Descriptor()
	// productqty.ProductColorIDValidator is a validator for the "product_color_id" field. It is called by the builders before save.
	productqty.ProductColorIDValidator = productqtyDescProductColorID.Validators[0].(func(string) error)
	// productqtyDescID is the schema descriptor for id field.
	productqtyDescID := productqtyFields[0].Descriptor()
	// productqty.DefaultID holds the default value on creation for the id field.
	productqty.DefaultID = productqtyDescID.Default.(func() uuid.UUID)
	shipmentMixin := schema.Shipment{}.Mixin()
	shipmentHooks := schema.Shipment{}.Hooks()
	shipment.Hooks[0] = shipmentHooks[0]
	shipmentMixinFields0 := shipmentMixin[0].Fields()
	_ = shipmentMixinFields0
	shipmentFields := schema.Shipment{}.Fields()
	_ = shipmentFields
	// shipmentDescCreatedAt is the schema descriptor for created_at field.
	shipmentDescCreatedAt := shipmentMixinFields0[0].Descriptor()
	// shipment.DefaultCreatedAt holds the default value on creation for the created_at field.
	shipment.DefaultCreatedAt = shipmentDescCreatedAt.Default.(func() time.Time)
	// shipmentDescUpdatedAt is the schema descriptor for updated_at field.
	shipmentDescUpdatedAt := shipmentMixinFields0[1].Descriptor()
	// shipment.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	shipment.DefaultUpdatedAt = shipmentDescUpdatedAt.Default.(func() time.Time)
	// shipment.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	shipment.UpdateDefaultUpdatedAt = shipmentDescUpdatedAt.UpdateDefault.(func() time.Time)
	// shipmentDescShipmentTrackingNumber is the schema descriptor for shipment_tracking_number field.
	shipmentDescShipmentTrackingNumber := shipmentFields[3].Descriptor()
	// shipment.ShipmentTrackingNumberValidator is a validator for the "shipment_tracking_number" field. It is called by the builders before save.
	shipment.ShipmentTrackingNumberValidator = func() func(string) error {
		validators := shipmentDescShipmentTrackingNumber.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(shipment_tracking_number string) error {
			for _, fn := range fns {
				if err := fn(shipment_tracking_number); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// shipmentDescID is the schema descriptor for id field.
	shipmentDescID := shipmentFields[0].Descriptor()
	// shipment.DefaultID holds the default value on creation for the id field.
	shipment.DefaultID = shipmentDescID.Default.(func() uuid.UUID)
	shipmentitemMixin := schema.ShipmentItem{}.Mixin()
	shipmentitemMixinFields0 := shipmentitemMixin[0].Fields()
	_ = shipmentitemMixinFields0
	shipmentitemFields := schema.ShipmentItem{}.Fields()
	_ = shipmentitemFields
	// shipmentitemDescCreatedAt is the schema descriptor for created_at field.
	shipmentitemDescCreatedAt := shipmentitemMixinFields0[0].Descriptor()
	// shipmentitem.DefaultCreatedAt holds the default value on creation for the created_at field.
	shipmentitem.DefaultCreatedAt = shipmentitemDescCreatedAt.Default.(func() time.Time)
	// shipmentitemDescUpdatedAt is the schema descriptor for updated_at field.
	shipmentitemDescUpdatedAt := shipmentitemMixinFields0[1].Descriptor()
	// shipmentitem.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	shipmentitem.DefaultUpdatedAt = shipmentitemDescUpdatedAt.Default.(func() time.Time)
	// shipmentitem.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	shipmentitem.UpdateDefaultUpdatedAt = shipmentitemDescUpdatedAt.UpdateDefault.(func() time.Time)
	// shipmentitemDescID is the schema descriptor for id field.
	shipmentitemDescID := shipmentitemFields[0].Descriptor()
	// shipmentitem.DefaultID holds the default value on creation for the id field.
	shipmentitem.DefaultID = shipmentitemDescID.Default.(func() uuid.UUID)
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
	warehouseassignmentMixin := schema.WarehouseAssignment{}.Mixin()
	warehouseassignmentMixinFields0 := warehouseassignmentMixin[0].Fields()
	_ = warehouseassignmentMixinFields0
	warehouseassignmentFields := schema.WarehouseAssignment{}.Fields()
	_ = warehouseassignmentFields
	// warehouseassignmentDescCreatedAt is the schema descriptor for created_at field.
	warehouseassignmentDescCreatedAt := warehouseassignmentMixinFields0[0].Descriptor()
	// warehouseassignment.DefaultCreatedAt holds the default value on creation for the created_at field.
	warehouseassignment.DefaultCreatedAt = warehouseassignmentDescCreatedAt.Default.(func() time.Time)
	// warehouseassignmentDescUpdatedAt is the schema descriptor for updated_at field.
	warehouseassignmentDescUpdatedAt := warehouseassignmentMixinFields0[1].Descriptor()
	// warehouseassignment.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	warehouseassignment.DefaultUpdatedAt = warehouseassignmentDescUpdatedAt.Default.(func() time.Time)
	// warehouseassignment.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	warehouseassignment.UpdateDefaultUpdatedAt = warehouseassignmentDescUpdatedAt.UpdateDefault.(func() time.Time)
	// warehouseassignmentDescID is the schema descriptor for id field.
	warehouseassignmentDescID := warehouseassignmentFields[0].Descriptor()
	// warehouseassignment.DefaultID holds the default value on creation for the id field.
	warehouseassignment.DefaultID = warehouseassignmentDescID.Default.(func() uuid.UUID)
	workunitinfoFields := schema.WorkUnitInfo{}.Fields()
	_ = workunitinfoFields
	// workunitinfoDescName is the schema descriptor for name field.
	workunitinfoDescName := workunitinfoFields[1].Descriptor()
	// workunitinfo.NameValidator is a validator for the "name" field. It is called by the builders before save.
	workunitinfo.NameValidator = workunitinfoDescName.Validators[0].(func(string) error)
	// workunitinfoDescID is the schema descriptor for id field.
	workunitinfoDescID := workunitinfoFields[0].Descriptor()
	// workunitinfo.DefaultID holds the default value on creation for the id field.
	workunitinfo.DefaultID = workunitinfoDescID.Default.(func() uuid.UUID)
}

const (
	Version = "v0.13.1"                                         // Version of ent codegen.
	Sum     = "h1:uD8QwN1h6SNphdCCzmkMN3feSUzNnVvV/WIkHKMbzOE=" // Sum of ent codegen.
)

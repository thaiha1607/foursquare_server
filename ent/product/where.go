// Code generated by ent, DO NOT EDIT.

package product

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/shopspring/decimal"
	"github.com/thaiha1607/foursquare_server/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.Product {
	return predicate.Product(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.Product {
	return predicate.Product(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.Product {
	return predicate.Product(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.Product {
	return predicate.Product(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.Product {
	return predicate.Product(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.Product {
	return predicate.Product(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.Product {
	return predicate.Product(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.Product {
	return predicate.Product(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.Product {
	return predicate.Product(sql.FieldContainsFold(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldUpdatedAt, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldName, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldDescription, v))
}

// Year applies equality check predicate on the "year" field. It's identical to YearEQ.
func Year(v int) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldYear, v))
}

// Price applies equality check predicate on the "price" field. It's identical to PriceEQ.
func Price(v decimal.Decimal) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldPrice, v))
}

// Qty applies equality check predicate on the "qty" field. It's identical to QtyEQ.
func Qty(v decimal.Decimal) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldQty, v))
}

// UnitOfMeasurement applies equality check predicate on the "unit_of_measurement" field. It's identical to UnitOfMeasurementEQ.
func UnitOfMeasurement(v string) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldUnitOfMeasurement, v))
}

// Type applies equality check predicate on the "type" field. It's identical to TypeEQ.
func Type(v string) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldType, v))
}

// Provider applies equality check predicate on the "provider" field. It's identical to ProviderEQ.
func Provider(v string) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldProvider, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Product {
	return predicate.Product(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Product {
	return predicate.Product(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Product {
	return predicate.Product(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Product {
	return predicate.Product(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Product {
	return predicate.Product(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Product {
	return predicate.Product(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Product {
	return predicate.Product(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Product {
	return predicate.Product(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Product {
	return predicate.Product(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Product {
	return predicate.Product(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Product {
	return predicate.Product(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Product {
	return predicate.Product(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Product {
	return predicate.Product(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Product {
	return predicate.Product(sql.FieldLTE(FieldUpdatedAt, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Product {
	return predicate.Product(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Product {
	return predicate.Product(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Product {
	return predicate.Product(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Product {
	return predicate.Product(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Product {
	return predicate.Product(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Product {
	return predicate.Product(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Product {
	return predicate.Product(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Product {
	return predicate.Product(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Product {
	return predicate.Product(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Product {
	return predicate.Product(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Product {
	return predicate.Product(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Product {
	return predicate.Product(sql.FieldContainsFold(FieldName, v))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.Product {
	return predicate.Product(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.Product {
	return predicate.Product(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.Product {
	return predicate.Product(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.Product {
	return predicate.Product(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.Product {
	return predicate.Product(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.Product {
	return predicate.Product(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.Product {
	return predicate.Product(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.Product {
	return predicate.Product(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.Product {
	return predicate.Product(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.Product {
	return predicate.Product(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionIsNil applies the IsNil predicate on the "description" field.
func DescriptionIsNil() predicate.Product {
	return predicate.Product(sql.FieldIsNull(FieldDescription))
}

// DescriptionNotNil applies the NotNil predicate on the "description" field.
func DescriptionNotNil() predicate.Product {
	return predicate.Product(sql.FieldNotNull(FieldDescription))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.Product {
	return predicate.Product(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.Product {
	return predicate.Product(sql.FieldContainsFold(FieldDescription, v))
}

// YearEQ applies the EQ predicate on the "year" field.
func YearEQ(v int) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldYear, v))
}

// YearNEQ applies the NEQ predicate on the "year" field.
func YearNEQ(v int) predicate.Product {
	return predicate.Product(sql.FieldNEQ(FieldYear, v))
}

// YearIn applies the In predicate on the "year" field.
func YearIn(vs ...int) predicate.Product {
	return predicate.Product(sql.FieldIn(FieldYear, vs...))
}

// YearNotIn applies the NotIn predicate on the "year" field.
func YearNotIn(vs ...int) predicate.Product {
	return predicate.Product(sql.FieldNotIn(FieldYear, vs...))
}

// YearGT applies the GT predicate on the "year" field.
func YearGT(v int) predicate.Product {
	return predicate.Product(sql.FieldGT(FieldYear, v))
}

// YearGTE applies the GTE predicate on the "year" field.
func YearGTE(v int) predicate.Product {
	return predicate.Product(sql.FieldGTE(FieldYear, v))
}

// YearLT applies the LT predicate on the "year" field.
func YearLT(v int) predicate.Product {
	return predicate.Product(sql.FieldLT(FieldYear, v))
}

// YearLTE applies the LTE predicate on the "year" field.
func YearLTE(v int) predicate.Product {
	return predicate.Product(sql.FieldLTE(FieldYear, v))
}

// YearIsNil applies the IsNil predicate on the "year" field.
func YearIsNil() predicate.Product {
	return predicate.Product(sql.FieldIsNull(FieldYear))
}

// YearNotNil applies the NotNil predicate on the "year" field.
func YearNotNil() predicate.Product {
	return predicate.Product(sql.FieldNotNull(FieldYear))
}

// PriceEQ applies the EQ predicate on the "price" field.
func PriceEQ(v decimal.Decimal) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldPrice, v))
}

// PriceNEQ applies the NEQ predicate on the "price" field.
func PriceNEQ(v decimal.Decimal) predicate.Product {
	return predicate.Product(sql.FieldNEQ(FieldPrice, v))
}

// PriceIn applies the In predicate on the "price" field.
func PriceIn(vs ...decimal.Decimal) predicate.Product {
	return predicate.Product(sql.FieldIn(FieldPrice, vs...))
}

// PriceNotIn applies the NotIn predicate on the "price" field.
func PriceNotIn(vs ...decimal.Decimal) predicate.Product {
	return predicate.Product(sql.FieldNotIn(FieldPrice, vs...))
}

// PriceGT applies the GT predicate on the "price" field.
func PriceGT(v decimal.Decimal) predicate.Product {
	return predicate.Product(sql.FieldGT(FieldPrice, v))
}

// PriceGTE applies the GTE predicate on the "price" field.
func PriceGTE(v decimal.Decimal) predicate.Product {
	return predicate.Product(sql.FieldGTE(FieldPrice, v))
}

// PriceLT applies the LT predicate on the "price" field.
func PriceLT(v decimal.Decimal) predicate.Product {
	return predicate.Product(sql.FieldLT(FieldPrice, v))
}

// PriceLTE applies the LTE predicate on the "price" field.
func PriceLTE(v decimal.Decimal) predicate.Product {
	return predicate.Product(sql.FieldLTE(FieldPrice, v))
}

// QtyEQ applies the EQ predicate on the "qty" field.
func QtyEQ(v decimal.Decimal) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldQty, v))
}

// QtyNEQ applies the NEQ predicate on the "qty" field.
func QtyNEQ(v decimal.Decimal) predicate.Product {
	return predicate.Product(sql.FieldNEQ(FieldQty, v))
}

// QtyIn applies the In predicate on the "qty" field.
func QtyIn(vs ...decimal.Decimal) predicate.Product {
	return predicate.Product(sql.FieldIn(FieldQty, vs...))
}

// QtyNotIn applies the NotIn predicate on the "qty" field.
func QtyNotIn(vs ...decimal.Decimal) predicate.Product {
	return predicate.Product(sql.FieldNotIn(FieldQty, vs...))
}

// QtyGT applies the GT predicate on the "qty" field.
func QtyGT(v decimal.Decimal) predicate.Product {
	return predicate.Product(sql.FieldGT(FieldQty, v))
}

// QtyGTE applies the GTE predicate on the "qty" field.
func QtyGTE(v decimal.Decimal) predicate.Product {
	return predicate.Product(sql.FieldGTE(FieldQty, v))
}

// QtyLT applies the LT predicate on the "qty" field.
func QtyLT(v decimal.Decimal) predicate.Product {
	return predicate.Product(sql.FieldLT(FieldQty, v))
}

// QtyLTE applies the LTE predicate on the "qty" field.
func QtyLTE(v decimal.Decimal) predicate.Product {
	return predicate.Product(sql.FieldLTE(FieldQty, v))
}

// UnitOfMeasurementEQ applies the EQ predicate on the "unit_of_measurement" field.
func UnitOfMeasurementEQ(v string) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldUnitOfMeasurement, v))
}

// UnitOfMeasurementNEQ applies the NEQ predicate on the "unit_of_measurement" field.
func UnitOfMeasurementNEQ(v string) predicate.Product {
	return predicate.Product(sql.FieldNEQ(FieldUnitOfMeasurement, v))
}

// UnitOfMeasurementIn applies the In predicate on the "unit_of_measurement" field.
func UnitOfMeasurementIn(vs ...string) predicate.Product {
	return predicate.Product(sql.FieldIn(FieldUnitOfMeasurement, vs...))
}

// UnitOfMeasurementNotIn applies the NotIn predicate on the "unit_of_measurement" field.
func UnitOfMeasurementNotIn(vs ...string) predicate.Product {
	return predicate.Product(sql.FieldNotIn(FieldUnitOfMeasurement, vs...))
}

// UnitOfMeasurementGT applies the GT predicate on the "unit_of_measurement" field.
func UnitOfMeasurementGT(v string) predicate.Product {
	return predicate.Product(sql.FieldGT(FieldUnitOfMeasurement, v))
}

// UnitOfMeasurementGTE applies the GTE predicate on the "unit_of_measurement" field.
func UnitOfMeasurementGTE(v string) predicate.Product {
	return predicate.Product(sql.FieldGTE(FieldUnitOfMeasurement, v))
}

// UnitOfMeasurementLT applies the LT predicate on the "unit_of_measurement" field.
func UnitOfMeasurementLT(v string) predicate.Product {
	return predicate.Product(sql.FieldLT(FieldUnitOfMeasurement, v))
}

// UnitOfMeasurementLTE applies the LTE predicate on the "unit_of_measurement" field.
func UnitOfMeasurementLTE(v string) predicate.Product {
	return predicate.Product(sql.FieldLTE(FieldUnitOfMeasurement, v))
}

// UnitOfMeasurementContains applies the Contains predicate on the "unit_of_measurement" field.
func UnitOfMeasurementContains(v string) predicate.Product {
	return predicate.Product(sql.FieldContains(FieldUnitOfMeasurement, v))
}

// UnitOfMeasurementHasPrefix applies the HasPrefix predicate on the "unit_of_measurement" field.
func UnitOfMeasurementHasPrefix(v string) predicate.Product {
	return predicate.Product(sql.FieldHasPrefix(FieldUnitOfMeasurement, v))
}

// UnitOfMeasurementHasSuffix applies the HasSuffix predicate on the "unit_of_measurement" field.
func UnitOfMeasurementHasSuffix(v string) predicate.Product {
	return predicate.Product(sql.FieldHasSuffix(FieldUnitOfMeasurement, v))
}

// UnitOfMeasurementIsNil applies the IsNil predicate on the "unit_of_measurement" field.
func UnitOfMeasurementIsNil() predicate.Product {
	return predicate.Product(sql.FieldIsNull(FieldUnitOfMeasurement))
}

// UnitOfMeasurementNotNil applies the NotNil predicate on the "unit_of_measurement" field.
func UnitOfMeasurementNotNil() predicate.Product {
	return predicate.Product(sql.FieldNotNull(FieldUnitOfMeasurement))
}

// UnitOfMeasurementEqualFold applies the EqualFold predicate on the "unit_of_measurement" field.
func UnitOfMeasurementEqualFold(v string) predicate.Product {
	return predicate.Product(sql.FieldEqualFold(FieldUnitOfMeasurement, v))
}

// UnitOfMeasurementContainsFold applies the ContainsFold predicate on the "unit_of_measurement" field.
func UnitOfMeasurementContainsFold(v string) predicate.Product {
	return predicate.Product(sql.FieldContainsFold(FieldUnitOfMeasurement, v))
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v string) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldType, v))
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v string) predicate.Product {
	return predicate.Product(sql.FieldNEQ(FieldType, v))
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...string) predicate.Product {
	return predicate.Product(sql.FieldIn(FieldType, vs...))
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...string) predicate.Product {
	return predicate.Product(sql.FieldNotIn(FieldType, vs...))
}

// TypeGT applies the GT predicate on the "type" field.
func TypeGT(v string) predicate.Product {
	return predicate.Product(sql.FieldGT(FieldType, v))
}

// TypeGTE applies the GTE predicate on the "type" field.
func TypeGTE(v string) predicate.Product {
	return predicate.Product(sql.FieldGTE(FieldType, v))
}

// TypeLT applies the LT predicate on the "type" field.
func TypeLT(v string) predicate.Product {
	return predicate.Product(sql.FieldLT(FieldType, v))
}

// TypeLTE applies the LTE predicate on the "type" field.
func TypeLTE(v string) predicate.Product {
	return predicate.Product(sql.FieldLTE(FieldType, v))
}

// TypeContains applies the Contains predicate on the "type" field.
func TypeContains(v string) predicate.Product {
	return predicate.Product(sql.FieldContains(FieldType, v))
}

// TypeHasPrefix applies the HasPrefix predicate on the "type" field.
func TypeHasPrefix(v string) predicate.Product {
	return predicate.Product(sql.FieldHasPrefix(FieldType, v))
}

// TypeHasSuffix applies the HasSuffix predicate on the "type" field.
func TypeHasSuffix(v string) predicate.Product {
	return predicate.Product(sql.FieldHasSuffix(FieldType, v))
}

// TypeIsNil applies the IsNil predicate on the "type" field.
func TypeIsNil() predicate.Product {
	return predicate.Product(sql.FieldIsNull(FieldType))
}

// TypeNotNil applies the NotNil predicate on the "type" field.
func TypeNotNil() predicate.Product {
	return predicate.Product(sql.FieldNotNull(FieldType))
}

// TypeEqualFold applies the EqualFold predicate on the "type" field.
func TypeEqualFold(v string) predicate.Product {
	return predicate.Product(sql.FieldEqualFold(FieldType, v))
}

// TypeContainsFold applies the ContainsFold predicate on the "type" field.
func TypeContainsFold(v string) predicate.Product {
	return predicate.Product(sql.FieldContainsFold(FieldType, v))
}

// ProviderEQ applies the EQ predicate on the "provider" field.
func ProviderEQ(v string) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldProvider, v))
}

// ProviderNEQ applies the NEQ predicate on the "provider" field.
func ProviderNEQ(v string) predicate.Product {
	return predicate.Product(sql.FieldNEQ(FieldProvider, v))
}

// ProviderIn applies the In predicate on the "provider" field.
func ProviderIn(vs ...string) predicate.Product {
	return predicate.Product(sql.FieldIn(FieldProvider, vs...))
}

// ProviderNotIn applies the NotIn predicate on the "provider" field.
func ProviderNotIn(vs ...string) predicate.Product {
	return predicate.Product(sql.FieldNotIn(FieldProvider, vs...))
}

// ProviderGT applies the GT predicate on the "provider" field.
func ProviderGT(v string) predicate.Product {
	return predicate.Product(sql.FieldGT(FieldProvider, v))
}

// ProviderGTE applies the GTE predicate on the "provider" field.
func ProviderGTE(v string) predicate.Product {
	return predicate.Product(sql.FieldGTE(FieldProvider, v))
}

// ProviderLT applies the LT predicate on the "provider" field.
func ProviderLT(v string) predicate.Product {
	return predicate.Product(sql.FieldLT(FieldProvider, v))
}

// ProviderLTE applies the LTE predicate on the "provider" field.
func ProviderLTE(v string) predicate.Product {
	return predicate.Product(sql.FieldLTE(FieldProvider, v))
}

// ProviderContains applies the Contains predicate on the "provider" field.
func ProviderContains(v string) predicate.Product {
	return predicate.Product(sql.FieldContains(FieldProvider, v))
}

// ProviderHasPrefix applies the HasPrefix predicate on the "provider" field.
func ProviderHasPrefix(v string) predicate.Product {
	return predicate.Product(sql.FieldHasPrefix(FieldProvider, v))
}

// ProviderHasSuffix applies the HasSuffix predicate on the "provider" field.
func ProviderHasSuffix(v string) predicate.Product {
	return predicate.Product(sql.FieldHasSuffix(FieldProvider, v))
}

// ProviderIsNil applies the IsNil predicate on the "provider" field.
func ProviderIsNil() predicate.Product {
	return predicate.Product(sql.FieldIsNull(FieldProvider))
}

// ProviderNotNil applies the NotNil predicate on the "provider" field.
func ProviderNotNil() predicate.Product {
	return predicate.Product(sql.FieldNotNull(FieldProvider))
}

// ProviderEqualFold applies the EqualFold predicate on the "provider" field.
func ProviderEqualFold(v string) predicate.Product {
	return predicate.Product(sql.FieldEqualFold(FieldProvider, v))
}

// ProviderContainsFold applies the ContainsFold predicate on the "provider" field.
func ProviderContainsFold(v string) predicate.Product {
	return predicate.Product(sql.FieldContainsFold(FieldProvider, v))
}

// HasTags applies the HasEdge predicate on the "tags" edge.
func HasTags() predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, TagsTable, TagsPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTagsWith applies the HasEdge predicate on the "tags" edge with a given conditions (other predicates).
func HasTagsWith(preds ...predicate.Tag) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		step := newTagsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasProductTags applies the HasEdge predicate on the "product_tags" edge.
func HasProductTags() predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, ProductTagsTable, ProductTagsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasProductTagsWith applies the HasEdge predicate on the "product_tags" edge with a given conditions (other predicates).
func HasProductTagsWith(preds ...predicate.ProductTag) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		step := newProductTagsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Product) predicate.Product {
	return predicate.Product(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Product) predicate.Product {
	return predicate.Product(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Product) predicate.Product {
	return predicate.Product(sql.NotPredicates(p))
}

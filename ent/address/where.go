// Code generated by ent, DO NOT EDIT.

package address

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/thaiha1607/foursquare_server/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.Address {
	return predicate.Address(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.Address {
	return predicate.Address(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.Address {
	return predicate.Address(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.Address {
	return predicate.Address(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.Address {
	return predicate.Address(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.Address {
	return predicate.Address(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.Address {
	return predicate.Address(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.Address {
	return predicate.Address(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.Address {
	return predicate.Address(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.Address {
	return predicate.Address(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.Address {
	return predicate.Address(sql.FieldContainsFold(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Address {
	return predicate.Address(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Address {
	return predicate.Address(sql.FieldEQ(FieldUpdatedAt, v))
}

// Line1 applies equality check predicate on the "line1" field. It's identical to Line1EQ.
func Line1(v string) predicate.Address {
	return predicate.Address(sql.FieldEQ(FieldLine1, v))
}

// Line2 applies equality check predicate on the "line2" field. It's identical to Line2EQ.
func Line2(v string) predicate.Address {
	return predicate.Address(sql.FieldEQ(FieldLine2, v))
}

// City applies equality check predicate on the "city" field. It's identical to CityEQ.
func City(v string) predicate.Address {
	return predicate.Address(sql.FieldEQ(FieldCity, v))
}

// StateOrProvince applies equality check predicate on the "state_or_province" field. It's identical to StateOrProvinceEQ.
func StateOrProvince(v string) predicate.Address {
	return predicate.Address(sql.FieldEQ(FieldStateOrProvince, v))
}

// ZipOrPostcode applies equality check predicate on the "zip_or_postcode" field. It's identical to ZipOrPostcodeEQ.
func ZipOrPostcode(v string) predicate.Address {
	return predicate.Address(sql.FieldEQ(FieldZipOrPostcode, v))
}

// Country applies equality check predicate on the "country" field. It's identical to CountryEQ.
func Country(v string) predicate.Address {
	return predicate.Address(sql.FieldEQ(FieldCountry, v))
}

// OtherAddressDetails applies equality check predicate on the "other_address_details" field. It's identical to OtherAddressDetailsEQ.
func OtherAddressDetails(v string) predicate.Address {
	return predicate.Address(sql.FieldEQ(FieldOtherAddressDetails, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Address {
	return predicate.Address(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Address {
	return predicate.Address(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Address {
	return predicate.Address(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Address {
	return predicate.Address(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Address {
	return predicate.Address(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Address {
	return predicate.Address(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Address {
	return predicate.Address(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Address {
	return predicate.Address(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Address {
	return predicate.Address(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Address {
	return predicate.Address(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Address {
	return predicate.Address(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Address {
	return predicate.Address(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Address {
	return predicate.Address(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Address {
	return predicate.Address(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Address {
	return predicate.Address(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Address {
	return predicate.Address(sql.FieldLTE(FieldUpdatedAt, v))
}

// Line1EQ applies the EQ predicate on the "line1" field.
func Line1EQ(v string) predicate.Address {
	return predicate.Address(sql.FieldEQ(FieldLine1, v))
}

// Line1NEQ applies the NEQ predicate on the "line1" field.
func Line1NEQ(v string) predicate.Address {
	return predicate.Address(sql.FieldNEQ(FieldLine1, v))
}

// Line1In applies the In predicate on the "line1" field.
func Line1In(vs ...string) predicate.Address {
	return predicate.Address(sql.FieldIn(FieldLine1, vs...))
}

// Line1NotIn applies the NotIn predicate on the "line1" field.
func Line1NotIn(vs ...string) predicate.Address {
	return predicate.Address(sql.FieldNotIn(FieldLine1, vs...))
}

// Line1GT applies the GT predicate on the "line1" field.
func Line1GT(v string) predicate.Address {
	return predicate.Address(sql.FieldGT(FieldLine1, v))
}

// Line1GTE applies the GTE predicate on the "line1" field.
func Line1GTE(v string) predicate.Address {
	return predicate.Address(sql.FieldGTE(FieldLine1, v))
}

// Line1LT applies the LT predicate on the "line1" field.
func Line1LT(v string) predicate.Address {
	return predicate.Address(sql.FieldLT(FieldLine1, v))
}

// Line1LTE applies the LTE predicate on the "line1" field.
func Line1LTE(v string) predicate.Address {
	return predicate.Address(sql.FieldLTE(FieldLine1, v))
}

// Line1Contains applies the Contains predicate on the "line1" field.
func Line1Contains(v string) predicate.Address {
	return predicate.Address(sql.FieldContains(FieldLine1, v))
}

// Line1HasPrefix applies the HasPrefix predicate on the "line1" field.
func Line1HasPrefix(v string) predicate.Address {
	return predicate.Address(sql.FieldHasPrefix(FieldLine1, v))
}

// Line1HasSuffix applies the HasSuffix predicate on the "line1" field.
func Line1HasSuffix(v string) predicate.Address {
	return predicate.Address(sql.FieldHasSuffix(FieldLine1, v))
}

// Line1EqualFold applies the EqualFold predicate on the "line1" field.
func Line1EqualFold(v string) predicate.Address {
	return predicate.Address(sql.FieldEqualFold(FieldLine1, v))
}

// Line1ContainsFold applies the ContainsFold predicate on the "line1" field.
func Line1ContainsFold(v string) predicate.Address {
	return predicate.Address(sql.FieldContainsFold(FieldLine1, v))
}

// Line2EQ applies the EQ predicate on the "line2" field.
func Line2EQ(v string) predicate.Address {
	return predicate.Address(sql.FieldEQ(FieldLine2, v))
}

// Line2NEQ applies the NEQ predicate on the "line2" field.
func Line2NEQ(v string) predicate.Address {
	return predicate.Address(sql.FieldNEQ(FieldLine2, v))
}

// Line2In applies the In predicate on the "line2" field.
func Line2In(vs ...string) predicate.Address {
	return predicate.Address(sql.FieldIn(FieldLine2, vs...))
}

// Line2NotIn applies the NotIn predicate on the "line2" field.
func Line2NotIn(vs ...string) predicate.Address {
	return predicate.Address(sql.FieldNotIn(FieldLine2, vs...))
}

// Line2GT applies the GT predicate on the "line2" field.
func Line2GT(v string) predicate.Address {
	return predicate.Address(sql.FieldGT(FieldLine2, v))
}

// Line2GTE applies the GTE predicate on the "line2" field.
func Line2GTE(v string) predicate.Address {
	return predicate.Address(sql.FieldGTE(FieldLine2, v))
}

// Line2LT applies the LT predicate on the "line2" field.
func Line2LT(v string) predicate.Address {
	return predicate.Address(sql.FieldLT(FieldLine2, v))
}

// Line2LTE applies the LTE predicate on the "line2" field.
func Line2LTE(v string) predicate.Address {
	return predicate.Address(sql.FieldLTE(FieldLine2, v))
}

// Line2Contains applies the Contains predicate on the "line2" field.
func Line2Contains(v string) predicate.Address {
	return predicate.Address(sql.FieldContains(FieldLine2, v))
}

// Line2HasPrefix applies the HasPrefix predicate on the "line2" field.
func Line2HasPrefix(v string) predicate.Address {
	return predicate.Address(sql.FieldHasPrefix(FieldLine2, v))
}

// Line2HasSuffix applies the HasSuffix predicate on the "line2" field.
func Line2HasSuffix(v string) predicate.Address {
	return predicate.Address(sql.FieldHasSuffix(FieldLine2, v))
}

// Line2IsNil applies the IsNil predicate on the "line2" field.
func Line2IsNil() predicate.Address {
	return predicate.Address(sql.FieldIsNull(FieldLine2))
}

// Line2NotNil applies the NotNil predicate on the "line2" field.
func Line2NotNil() predicate.Address {
	return predicate.Address(sql.FieldNotNull(FieldLine2))
}

// Line2EqualFold applies the EqualFold predicate on the "line2" field.
func Line2EqualFold(v string) predicate.Address {
	return predicate.Address(sql.FieldEqualFold(FieldLine2, v))
}

// Line2ContainsFold applies the ContainsFold predicate on the "line2" field.
func Line2ContainsFold(v string) predicate.Address {
	return predicate.Address(sql.FieldContainsFold(FieldLine2, v))
}

// CityEQ applies the EQ predicate on the "city" field.
func CityEQ(v string) predicate.Address {
	return predicate.Address(sql.FieldEQ(FieldCity, v))
}

// CityNEQ applies the NEQ predicate on the "city" field.
func CityNEQ(v string) predicate.Address {
	return predicate.Address(sql.FieldNEQ(FieldCity, v))
}

// CityIn applies the In predicate on the "city" field.
func CityIn(vs ...string) predicate.Address {
	return predicate.Address(sql.FieldIn(FieldCity, vs...))
}

// CityNotIn applies the NotIn predicate on the "city" field.
func CityNotIn(vs ...string) predicate.Address {
	return predicate.Address(sql.FieldNotIn(FieldCity, vs...))
}

// CityGT applies the GT predicate on the "city" field.
func CityGT(v string) predicate.Address {
	return predicate.Address(sql.FieldGT(FieldCity, v))
}

// CityGTE applies the GTE predicate on the "city" field.
func CityGTE(v string) predicate.Address {
	return predicate.Address(sql.FieldGTE(FieldCity, v))
}

// CityLT applies the LT predicate on the "city" field.
func CityLT(v string) predicate.Address {
	return predicate.Address(sql.FieldLT(FieldCity, v))
}

// CityLTE applies the LTE predicate on the "city" field.
func CityLTE(v string) predicate.Address {
	return predicate.Address(sql.FieldLTE(FieldCity, v))
}

// CityContains applies the Contains predicate on the "city" field.
func CityContains(v string) predicate.Address {
	return predicate.Address(sql.FieldContains(FieldCity, v))
}

// CityHasPrefix applies the HasPrefix predicate on the "city" field.
func CityHasPrefix(v string) predicate.Address {
	return predicate.Address(sql.FieldHasPrefix(FieldCity, v))
}

// CityHasSuffix applies the HasSuffix predicate on the "city" field.
func CityHasSuffix(v string) predicate.Address {
	return predicate.Address(sql.FieldHasSuffix(FieldCity, v))
}

// CityEqualFold applies the EqualFold predicate on the "city" field.
func CityEqualFold(v string) predicate.Address {
	return predicate.Address(sql.FieldEqualFold(FieldCity, v))
}

// CityContainsFold applies the ContainsFold predicate on the "city" field.
func CityContainsFold(v string) predicate.Address {
	return predicate.Address(sql.FieldContainsFold(FieldCity, v))
}

// StateOrProvinceEQ applies the EQ predicate on the "state_or_province" field.
func StateOrProvinceEQ(v string) predicate.Address {
	return predicate.Address(sql.FieldEQ(FieldStateOrProvince, v))
}

// StateOrProvinceNEQ applies the NEQ predicate on the "state_or_province" field.
func StateOrProvinceNEQ(v string) predicate.Address {
	return predicate.Address(sql.FieldNEQ(FieldStateOrProvince, v))
}

// StateOrProvinceIn applies the In predicate on the "state_or_province" field.
func StateOrProvinceIn(vs ...string) predicate.Address {
	return predicate.Address(sql.FieldIn(FieldStateOrProvince, vs...))
}

// StateOrProvinceNotIn applies the NotIn predicate on the "state_or_province" field.
func StateOrProvinceNotIn(vs ...string) predicate.Address {
	return predicate.Address(sql.FieldNotIn(FieldStateOrProvince, vs...))
}

// StateOrProvinceGT applies the GT predicate on the "state_or_province" field.
func StateOrProvinceGT(v string) predicate.Address {
	return predicate.Address(sql.FieldGT(FieldStateOrProvince, v))
}

// StateOrProvinceGTE applies the GTE predicate on the "state_or_province" field.
func StateOrProvinceGTE(v string) predicate.Address {
	return predicate.Address(sql.FieldGTE(FieldStateOrProvince, v))
}

// StateOrProvinceLT applies the LT predicate on the "state_or_province" field.
func StateOrProvinceLT(v string) predicate.Address {
	return predicate.Address(sql.FieldLT(FieldStateOrProvince, v))
}

// StateOrProvinceLTE applies the LTE predicate on the "state_or_province" field.
func StateOrProvinceLTE(v string) predicate.Address {
	return predicate.Address(sql.FieldLTE(FieldStateOrProvince, v))
}

// StateOrProvinceContains applies the Contains predicate on the "state_or_province" field.
func StateOrProvinceContains(v string) predicate.Address {
	return predicate.Address(sql.FieldContains(FieldStateOrProvince, v))
}

// StateOrProvinceHasPrefix applies the HasPrefix predicate on the "state_or_province" field.
func StateOrProvinceHasPrefix(v string) predicate.Address {
	return predicate.Address(sql.FieldHasPrefix(FieldStateOrProvince, v))
}

// StateOrProvinceHasSuffix applies the HasSuffix predicate on the "state_or_province" field.
func StateOrProvinceHasSuffix(v string) predicate.Address {
	return predicate.Address(sql.FieldHasSuffix(FieldStateOrProvince, v))
}

// StateOrProvinceIsNil applies the IsNil predicate on the "state_or_province" field.
func StateOrProvinceIsNil() predicate.Address {
	return predicate.Address(sql.FieldIsNull(FieldStateOrProvince))
}

// StateOrProvinceNotNil applies the NotNil predicate on the "state_or_province" field.
func StateOrProvinceNotNil() predicate.Address {
	return predicate.Address(sql.FieldNotNull(FieldStateOrProvince))
}

// StateOrProvinceEqualFold applies the EqualFold predicate on the "state_or_province" field.
func StateOrProvinceEqualFold(v string) predicate.Address {
	return predicate.Address(sql.FieldEqualFold(FieldStateOrProvince, v))
}

// StateOrProvinceContainsFold applies the ContainsFold predicate on the "state_or_province" field.
func StateOrProvinceContainsFold(v string) predicate.Address {
	return predicate.Address(sql.FieldContainsFold(FieldStateOrProvince, v))
}

// ZipOrPostcodeEQ applies the EQ predicate on the "zip_or_postcode" field.
func ZipOrPostcodeEQ(v string) predicate.Address {
	return predicate.Address(sql.FieldEQ(FieldZipOrPostcode, v))
}

// ZipOrPostcodeNEQ applies the NEQ predicate on the "zip_or_postcode" field.
func ZipOrPostcodeNEQ(v string) predicate.Address {
	return predicate.Address(sql.FieldNEQ(FieldZipOrPostcode, v))
}

// ZipOrPostcodeIn applies the In predicate on the "zip_or_postcode" field.
func ZipOrPostcodeIn(vs ...string) predicate.Address {
	return predicate.Address(sql.FieldIn(FieldZipOrPostcode, vs...))
}

// ZipOrPostcodeNotIn applies the NotIn predicate on the "zip_or_postcode" field.
func ZipOrPostcodeNotIn(vs ...string) predicate.Address {
	return predicate.Address(sql.FieldNotIn(FieldZipOrPostcode, vs...))
}

// ZipOrPostcodeGT applies the GT predicate on the "zip_or_postcode" field.
func ZipOrPostcodeGT(v string) predicate.Address {
	return predicate.Address(sql.FieldGT(FieldZipOrPostcode, v))
}

// ZipOrPostcodeGTE applies the GTE predicate on the "zip_or_postcode" field.
func ZipOrPostcodeGTE(v string) predicate.Address {
	return predicate.Address(sql.FieldGTE(FieldZipOrPostcode, v))
}

// ZipOrPostcodeLT applies the LT predicate on the "zip_or_postcode" field.
func ZipOrPostcodeLT(v string) predicate.Address {
	return predicate.Address(sql.FieldLT(FieldZipOrPostcode, v))
}

// ZipOrPostcodeLTE applies the LTE predicate on the "zip_or_postcode" field.
func ZipOrPostcodeLTE(v string) predicate.Address {
	return predicate.Address(sql.FieldLTE(FieldZipOrPostcode, v))
}

// ZipOrPostcodeContains applies the Contains predicate on the "zip_or_postcode" field.
func ZipOrPostcodeContains(v string) predicate.Address {
	return predicate.Address(sql.FieldContains(FieldZipOrPostcode, v))
}

// ZipOrPostcodeHasPrefix applies the HasPrefix predicate on the "zip_or_postcode" field.
func ZipOrPostcodeHasPrefix(v string) predicate.Address {
	return predicate.Address(sql.FieldHasPrefix(FieldZipOrPostcode, v))
}

// ZipOrPostcodeHasSuffix applies the HasSuffix predicate on the "zip_or_postcode" field.
func ZipOrPostcodeHasSuffix(v string) predicate.Address {
	return predicate.Address(sql.FieldHasSuffix(FieldZipOrPostcode, v))
}

// ZipOrPostcodeEqualFold applies the EqualFold predicate on the "zip_or_postcode" field.
func ZipOrPostcodeEqualFold(v string) predicate.Address {
	return predicate.Address(sql.FieldEqualFold(FieldZipOrPostcode, v))
}

// ZipOrPostcodeContainsFold applies the ContainsFold predicate on the "zip_or_postcode" field.
func ZipOrPostcodeContainsFold(v string) predicate.Address {
	return predicate.Address(sql.FieldContainsFold(FieldZipOrPostcode, v))
}

// CountryEQ applies the EQ predicate on the "country" field.
func CountryEQ(v string) predicate.Address {
	return predicate.Address(sql.FieldEQ(FieldCountry, v))
}

// CountryNEQ applies the NEQ predicate on the "country" field.
func CountryNEQ(v string) predicate.Address {
	return predicate.Address(sql.FieldNEQ(FieldCountry, v))
}

// CountryIn applies the In predicate on the "country" field.
func CountryIn(vs ...string) predicate.Address {
	return predicate.Address(sql.FieldIn(FieldCountry, vs...))
}

// CountryNotIn applies the NotIn predicate on the "country" field.
func CountryNotIn(vs ...string) predicate.Address {
	return predicate.Address(sql.FieldNotIn(FieldCountry, vs...))
}

// CountryGT applies the GT predicate on the "country" field.
func CountryGT(v string) predicate.Address {
	return predicate.Address(sql.FieldGT(FieldCountry, v))
}

// CountryGTE applies the GTE predicate on the "country" field.
func CountryGTE(v string) predicate.Address {
	return predicate.Address(sql.FieldGTE(FieldCountry, v))
}

// CountryLT applies the LT predicate on the "country" field.
func CountryLT(v string) predicate.Address {
	return predicate.Address(sql.FieldLT(FieldCountry, v))
}

// CountryLTE applies the LTE predicate on the "country" field.
func CountryLTE(v string) predicate.Address {
	return predicate.Address(sql.FieldLTE(FieldCountry, v))
}

// CountryContains applies the Contains predicate on the "country" field.
func CountryContains(v string) predicate.Address {
	return predicate.Address(sql.FieldContains(FieldCountry, v))
}

// CountryHasPrefix applies the HasPrefix predicate on the "country" field.
func CountryHasPrefix(v string) predicate.Address {
	return predicate.Address(sql.FieldHasPrefix(FieldCountry, v))
}

// CountryHasSuffix applies the HasSuffix predicate on the "country" field.
func CountryHasSuffix(v string) predicate.Address {
	return predicate.Address(sql.FieldHasSuffix(FieldCountry, v))
}

// CountryEqualFold applies the EqualFold predicate on the "country" field.
func CountryEqualFold(v string) predicate.Address {
	return predicate.Address(sql.FieldEqualFold(FieldCountry, v))
}

// CountryContainsFold applies the ContainsFold predicate on the "country" field.
func CountryContainsFold(v string) predicate.Address {
	return predicate.Address(sql.FieldContainsFold(FieldCountry, v))
}

// OtherAddressDetailsEQ applies the EQ predicate on the "other_address_details" field.
func OtherAddressDetailsEQ(v string) predicate.Address {
	return predicate.Address(sql.FieldEQ(FieldOtherAddressDetails, v))
}

// OtherAddressDetailsNEQ applies the NEQ predicate on the "other_address_details" field.
func OtherAddressDetailsNEQ(v string) predicate.Address {
	return predicate.Address(sql.FieldNEQ(FieldOtherAddressDetails, v))
}

// OtherAddressDetailsIn applies the In predicate on the "other_address_details" field.
func OtherAddressDetailsIn(vs ...string) predicate.Address {
	return predicate.Address(sql.FieldIn(FieldOtherAddressDetails, vs...))
}

// OtherAddressDetailsNotIn applies the NotIn predicate on the "other_address_details" field.
func OtherAddressDetailsNotIn(vs ...string) predicate.Address {
	return predicate.Address(sql.FieldNotIn(FieldOtherAddressDetails, vs...))
}

// OtherAddressDetailsGT applies the GT predicate on the "other_address_details" field.
func OtherAddressDetailsGT(v string) predicate.Address {
	return predicate.Address(sql.FieldGT(FieldOtherAddressDetails, v))
}

// OtherAddressDetailsGTE applies the GTE predicate on the "other_address_details" field.
func OtherAddressDetailsGTE(v string) predicate.Address {
	return predicate.Address(sql.FieldGTE(FieldOtherAddressDetails, v))
}

// OtherAddressDetailsLT applies the LT predicate on the "other_address_details" field.
func OtherAddressDetailsLT(v string) predicate.Address {
	return predicate.Address(sql.FieldLT(FieldOtherAddressDetails, v))
}

// OtherAddressDetailsLTE applies the LTE predicate on the "other_address_details" field.
func OtherAddressDetailsLTE(v string) predicate.Address {
	return predicate.Address(sql.FieldLTE(FieldOtherAddressDetails, v))
}

// OtherAddressDetailsContains applies the Contains predicate on the "other_address_details" field.
func OtherAddressDetailsContains(v string) predicate.Address {
	return predicate.Address(sql.FieldContains(FieldOtherAddressDetails, v))
}

// OtherAddressDetailsHasPrefix applies the HasPrefix predicate on the "other_address_details" field.
func OtherAddressDetailsHasPrefix(v string) predicate.Address {
	return predicate.Address(sql.FieldHasPrefix(FieldOtherAddressDetails, v))
}

// OtherAddressDetailsHasSuffix applies the HasSuffix predicate on the "other_address_details" field.
func OtherAddressDetailsHasSuffix(v string) predicate.Address {
	return predicate.Address(sql.FieldHasSuffix(FieldOtherAddressDetails, v))
}

// OtherAddressDetailsIsNil applies the IsNil predicate on the "other_address_details" field.
func OtherAddressDetailsIsNil() predicate.Address {
	return predicate.Address(sql.FieldIsNull(FieldOtherAddressDetails))
}

// OtherAddressDetailsNotNil applies the NotNil predicate on the "other_address_details" field.
func OtherAddressDetailsNotNil() predicate.Address {
	return predicate.Address(sql.FieldNotNull(FieldOtherAddressDetails))
}

// OtherAddressDetailsEqualFold applies the EqualFold predicate on the "other_address_details" field.
func OtherAddressDetailsEqualFold(v string) predicate.Address {
	return predicate.Address(sql.FieldEqualFold(FieldOtherAddressDetails, v))
}

// OtherAddressDetailsContainsFold applies the ContainsFold predicate on the "other_address_details" field.
func OtherAddressDetailsContainsFold(v string) predicate.Address {
	return predicate.Address(sql.FieldContainsFold(FieldOtherAddressDetails, v))
}

// HasPersons applies the HasEdge predicate on the "persons" edge.
func HasPersons() predicate.Address {
	return predicate.Address(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, PersonsTable, PersonsPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPersonsWith applies the HasEdge predicate on the "persons" edge with a given conditions (other predicates).
func HasPersonsWith(preds ...predicate.Person) predicate.Address {
	return predicate.Address(func(s *sql.Selector) {
		step := newPersonsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPersonAddresses applies the HasEdge predicate on the "person_addresses" edge.
func HasPersonAddresses() predicate.Address {
	return predicate.Address(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, PersonAddressesTable, PersonAddressesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPersonAddressesWith applies the HasEdge predicate on the "person_addresses" edge with a given conditions (other predicates).
func HasPersonAddressesWith(preds ...predicate.PersonAddress) predicate.Address {
	return predicate.Address(func(s *sql.Selector) {
		step := newPersonAddressesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Address) predicate.Address {
	return predicate.Address(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Address) predicate.Address {
	return predicate.Address(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Address) predicate.Address {
	return predicate.Address(sql.NotPredicates(p))
}

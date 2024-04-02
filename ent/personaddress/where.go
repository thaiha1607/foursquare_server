// Code generated by ent, DO NOT EDIT.

package personaddress

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
	"github.com/thaiha1607/foursquare_server/ent/predicate"
)

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.PersonAddress {
	return predicate.PersonAddress(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.PersonAddress {
	return predicate.PersonAddress(sql.FieldEQ(FieldUpdatedAt, v))
}

// PersonID applies equality check predicate on the "person_id" field. It's identical to PersonIDEQ.
func PersonID(v uuid.UUID) predicate.PersonAddress {
	return predicate.PersonAddress(sql.FieldEQ(FieldPersonID, v))
}

// AddressID applies equality check predicate on the "address_id" field. It's identical to AddressIDEQ.
func AddressID(v uuid.UUID) predicate.PersonAddress {
	return predicate.PersonAddress(sql.FieldEQ(FieldAddressID, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.PersonAddress {
	return predicate.PersonAddress(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.PersonAddress {
	return predicate.PersonAddress(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.PersonAddress {
	return predicate.PersonAddress(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.PersonAddress {
	return predicate.PersonAddress(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.PersonAddress {
	return predicate.PersonAddress(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.PersonAddress {
	return predicate.PersonAddress(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.PersonAddress {
	return predicate.PersonAddress(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.PersonAddress {
	return predicate.PersonAddress(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.PersonAddress {
	return predicate.PersonAddress(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.PersonAddress {
	return predicate.PersonAddress(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.PersonAddress {
	return predicate.PersonAddress(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.PersonAddress {
	return predicate.PersonAddress(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.PersonAddress {
	return predicate.PersonAddress(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.PersonAddress {
	return predicate.PersonAddress(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.PersonAddress {
	return predicate.PersonAddress(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.PersonAddress {
	return predicate.PersonAddress(sql.FieldLTE(FieldUpdatedAt, v))
}

// PersonIDEQ applies the EQ predicate on the "person_id" field.
func PersonIDEQ(v uuid.UUID) predicate.PersonAddress {
	return predicate.PersonAddress(sql.FieldEQ(FieldPersonID, v))
}

// PersonIDNEQ applies the NEQ predicate on the "person_id" field.
func PersonIDNEQ(v uuid.UUID) predicate.PersonAddress {
	return predicate.PersonAddress(sql.FieldNEQ(FieldPersonID, v))
}

// PersonIDIn applies the In predicate on the "person_id" field.
func PersonIDIn(vs ...uuid.UUID) predicate.PersonAddress {
	return predicate.PersonAddress(sql.FieldIn(FieldPersonID, vs...))
}

// PersonIDNotIn applies the NotIn predicate on the "person_id" field.
func PersonIDNotIn(vs ...uuid.UUID) predicate.PersonAddress {
	return predicate.PersonAddress(sql.FieldNotIn(FieldPersonID, vs...))
}

// AddressIDEQ applies the EQ predicate on the "address_id" field.
func AddressIDEQ(v uuid.UUID) predicate.PersonAddress {
	return predicate.PersonAddress(sql.FieldEQ(FieldAddressID, v))
}

// AddressIDNEQ applies the NEQ predicate on the "address_id" field.
func AddressIDNEQ(v uuid.UUID) predicate.PersonAddress {
	return predicate.PersonAddress(sql.FieldNEQ(FieldAddressID, v))
}

// AddressIDIn applies the In predicate on the "address_id" field.
func AddressIDIn(vs ...uuid.UUID) predicate.PersonAddress {
	return predicate.PersonAddress(sql.FieldIn(FieldAddressID, vs...))
}

// AddressIDNotIn applies the NotIn predicate on the "address_id" field.
func AddressIDNotIn(vs ...uuid.UUID) predicate.PersonAddress {
	return predicate.PersonAddress(sql.FieldNotIn(FieldAddressID, vs...))
}

// HasPersons applies the HasEdge predicate on the "persons" edge.
func HasPersons() predicate.PersonAddress {
	return predicate.PersonAddress(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, PersonsColumn),
			sqlgraph.Edge(sqlgraph.M2O, false, PersonsTable, PersonsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPersonsWith applies the HasEdge predicate on the "persons" edge with a given conditions (other predicates).
func HasPersonsWith(preds ...predicate.Person) predicate.PersonAddress {
	return predicate.PersonAddress(func(s *sql.Selector) {
		step := newPersonsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasAddresses applies the HasEdge predicate on the "addresses" edge.
func HasAddresses() predicate.PersonAddress {
	return predicate.PersonAddress(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, AddressesColumn),
			sqlgraph.Edge(sqlgraph.M2O, false, AddressesTable, AddressesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAddressesWith applies the HasEdge predicate on the "addresses" edge with a given conditions (other predicates).
func HasAddressesWith(preds ...predicate.Address) predicate.PersonAddress {
	return predicate.PersonAddress(func(s *sql.Selector) {
		step := newAddressesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.PersonAddress) predicate.PersonAddress {
	return predicate.PersonAddress(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.PersonAddress) predicate.PersonAddress {
	return predicate.PersonAddress(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.PersonAddress) predicate.PersonAddress {
	return predicate.PersonAddress(sql.NotPredicates(p))
}

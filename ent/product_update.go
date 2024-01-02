// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"net/url"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/dialect/sql/sqljson"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"github.com/thaiha1607/foursquare_server/ent/predicate"
	"github.com/thaiha1607/foursquare_server/ent/product"
	"github.com/thaiha1607/foursquare_server/ent/producttype"
	"github.com/thaiha1607/foursquare_server/ent/tag"
)

// ProductUpdate is the builder for updating Product entities.
type ProductUpdate struct {
	config
	hooks    []Hook
	mutation *ProductMutation
}

// Where appends a list predicates to the ProductUpdate builder.
func (pu *ProductUpdate) Where(ps ...predicate.Product) *ProductUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetUpdatedAt sets the "updated_at" field.
func (pu *ProductUpdate) SetUpdatedAt(t time.Time) *ProductUpdate {
	pu.mutation.SetUpdatedAt(t)
	return pu
}

// SetName sets the "name" field.
func (pu *ProductUpdate) SetName(s string) *ProductUpdate {
	pu.mutation.SetName(s)
	return pu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (pu *ProductUpdate) SetNillableName(s *string) *ProductUpdate {
	if s != nil {
		pu.SetName(*s)
	}
	return pu
}

// SetDescription sets the "description" field.
func (pu *ProductUpdate) SetDescription(s string) *ProductUpdate {
	pu.mutation.SetDescription(s)
	return pu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (pu *ProductUpdate) SetNillableDescription(s *string) *ProductUpdate {
	if s != nil {
		pu.SetDescription(*s)
	}
	return pu
}

// ClearDescription clears the value of the "description" field.
func (pu *ProductUpdate) ClearDescription() *ProductUpdate {
	pu.mutation.ClearDescription()
	return pu
}

// SetYear sets the "year" field.
func (pu *ProductUpdate) SetYear(i int) *ProductUpdate {
	pu.mutation.ResetYear()
	pu.mutation.SetYear(i)
	return pu
}

// SetNillableYear sets the "year" field if the given value is not nil.
func (pu *ProductUpdate) SetNillableYear(i *int) *ProductUpdate {
	if i != nil {
		pu.SetYear(*i)
	}
	return pu
}

// AddYear adds i to the "year" field.
func (pu *ProductUpdate) AddYear(i int) *ProductUpdate {
	pu.mutation.AddYear(i)
	return pu
}

// SetPrice sets the "price" field.
func (pu *ProductUpdate) SetPrice(d decimal.Decimal) *ProductUpdate {
	pu.mutation.ResetPrice()
	pu.mutation.SetPrice(d)
	return pu
}

// SetNillablePrice sets the "price" field if the given value is not nil.
func (pu *ProductUpdate) SetNillablePrice(d *decimal.Decimal) *ProductUpdate {
	if d != nil {
		pu.SetPrice(*d)
	}
	return pu
}

// AddPrice adds d to the "price" field.
func (pu *ProductUpdate) AddPrice(d decimal.Decimal) *ProductUpdate {
	pu.mutation.AddPrice(d)
	return pu
}

// SetQty sets the "qty" field.
func (pu *ProductUpdate) SetQty(d decimal.Decimal) *ProductUpdate {
	pu.mutation.ResetQty()
	pu.mutation.SetQty(d)
	return pu
}

// SetNillableQty sets the "qty" field if the given value is not nil.
func (pu *ProductUpdate) SetNillableQty(d *decimal.Decimal) *ProductUpdate {
	if d != nil {
		pu.SetQty(*d)
	}
	return pu
}

// AddQty adds d to the "qty" field.
func (pu *ProductUpdate) AddQty(d decimal.Decimal) *ProductUpdate {
	pu.mutation.AddQty(d)
	return pu
}

// SetImageUrls sets the "image_urls" field.
func (pu *ProductUpdate) SetImageUrls(u []url.URL) *ProductUpdate {
	pu.mutation.SetImageUrls(u)
	return pu
}

// AppendImageUrls appends u to the "image_urls" field.
func (pu *ProductUpdate) AppendImageUrls(u []url.URL) *ProductUpdate {
	pu.mutation.AppendImageUrls(u)
	return pu
}

// ClearImageUrls clears the value of the "image_urls" field.
func (pu *ProductUpdate) ClearImageUrls() *ProductUpdate {
	pu.mutation.ClearImageUrls()
	return pu
}

// SetUnitOfMeasurement sets the "unit_of_measurement" field.
func (pu *ProductUpdate) SetUnitOfMeasurement(s string) *ProductUpdate {
	pu.mutation.SetUnitOfMeasurement(s)
	return pu
}

// SetNillableUnitOfMeasurement sets the "unit_of_measurement" field if the given value is not nil.
func (pu *ProductUpdate) SetNillableUnitOfMeasurement(s *string) *ProductUpdate {
	if s != nil {
		pu.SetUnitOfMeasurement(*s)
	}
	return pu
}

// ClearUnitOfMeasurement clears the value of the "unit_of_measurement" field.
func (pu *ProductUpdate) ClearUnitOfMeasurement() *ProductUpdate {
	pu.mutation.ClearUnitOfMeasurement()
	return pu
}

// SetType sets the "type" field.
func (pu *ProductUpdate) SetType(i int) *ProductUpdate {
	pu.mutation.SetType(i)
	return pu
}

// SetNillableType sets the "type" field if the given value is not nil.
func (pu *ProductUpdate) SetNillableType(i *int) *ProductUpdate {
	if i != nil {
		pu.SetType(*i)
	}
	return pu
}

// SetProvider sets the "provider" field.
func (pu *ProductUpdate) SetProvider(s string) *ProductUpdate {
	pu.mutation.SetProvider(s)
	return pu
}

// SetNillableProvider sets the "provider" field if the given value is not nil.
func (pu *ProductUpdate) SetNillableProvider(s *string) *ProductUpdate {
	if s != nil {
		pu.SetProvider(*s)
	}
	return pu
}

// ClearProvider clears the value of the "provider" field.
func (pu *ProductUpdate) ClearProvider() *ProductUpdate {
	pu.mutation.ClearProvider()
	return pu
}

// SetProductTypeID sets the "product_type" edge to the ProductType entity by ID.
func (pu *ProductUpdate) SetProductTypeID(id int) *ProductUpdate {
	pu.mutation.SetProductTypeID(id)
	return pu
}

// SetProductType sets the "product_type" edge to the ProductType entity.
func (pu *ProductUpdate) SetProductType(p *ProductType) *ProductUpdate {
	return pu.SetProductTypeID(p.ID)
}

// AddTagIDs adds the "tags" edge to the Tag entity by IDs.
func (pu *ProductUpdate) AddTagIDs(ids ...uuid.UUID) *ProductUpdate {
	pu.mutation.AddTagIDs(ids...)
	return pu
}

// AddTags adds the "tags" edges to the Tag entity.
func (pu *ProductUpdate) AddTags(t ...*Tag) *ProductUpdate {
	ids := make([]uuid.UUID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return pu.AddTagIDs(ids...)
}

// Mutation returns the ProductMutation object of the builder.
func (pu *ProductUpdate) Mutation() *ProductMutation {
	return pu.mutation
}

// ClearProductType clears the "product_type" edge to the ProductType entity.
func (pu *ProductUpdate) ClearProductType() *ProductUpdate {
	pu.mutation.ClearProductType()
	return pu
}

// ClearTags clears all "tags" edges to the Tag entity.
func (pu *ProductUpdate) ClearTags() *ProductUpdate {
	pu.mutation.ClearTags()
	return pu
}

// RemoveTagIDs removes the "tags" edge to Tag entities by IDs.
func (pu *ProductUpdate) RemoveTagIDs(ids ...uuid.UUID) *ProductUpdate {
	pu.mutation.RemoveTagIDs(ids...)
	return pu
}

// RemoveTags removes "tags" edges to Tag entities.
func (pu *ProductUpdate) RemoveTags(t ...*Tag) *ProductUpdate {
	ids := make([]uuid.UUID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return pu.RemoveTagIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *ProductUpdate) Save(ctx context.Context) (int, error) {
	pu.defaults()
	return withHooks(ctx, pu.sqlSave, pu.mutation, pu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pu *ProductUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *ProductUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *ProductUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pu *ProductUpdate) defaults() {
	if _, ok := pu.mutation.UpdatedAt(); !ok {
		v := product.UpdateDefaultUpdatedAt()
		pu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pu *ProductUpdate) check() error {
	if v, ok := pu.mutation.Name(); ok {
		if err := product.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Product.name": %w`, err)}
		}
	}
	if v, ok := pu.mutation.Year(); ok {
		if err := product.YearValidator(v); err != nil {
			return &ValidationError{Name: "year", err: fmt.Errorf(`ent: validator failed for field "Product.year": %w`, err)}
		}
	}
	if _, ok := pu.mutation.ProductTypeID(); pu.mutation.ProductTypeCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Product.product_type"`)
	}
	return nil
}

func (pu *ProductUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := pu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(product.Table, product.Columns, sqlgraph.NewFieldSpec(product.FieldID, field.TypeUUID))
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.UpdatedAt(); ok {
		_spec.SetField(product.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := pu.mutation.Name(); ok {
		_spec.SetField(product.FieldName, field.TypeString, value)
	}
	if value, ok := pu.mutation.Description(); ok {
		_spec.SetField(product.FieldDescription, field.TypeString, value)
	}
	if pu.mutation.DescriptionCleared() {
		_spec.ClearField(product.FieldDescription, field.TypeString)
	}
	if value, ok := pu.mutation.Year(); ok {
		_spec.SetField(product.FieldYear, field.TypeInt, value)
	}
	if value, ok := pu.mutation.AddedYear(); ok {
		_spec.AddField(product.FieldYear, field.TypeInt, value)
	}
	if value, ok := pu.mutation.Price(); ok {
		_spec.SetField(product.FieldPrice, field.TypeFloat64, value)
	}
	if value, ok := pu.mutation.AddedPrice(); ok {
		_spec.AddField(product.FieldPrice, field.TypeFloat64, value)
	}
	if value, ok := pu.mutation.Qty(); ok {
		_spec.SetField(product.FieldQty, field.TypeFloat64, value)
	}
	if value, ok := pu.mutation.AddedQty(); ok {
		_spec.AddField(product.FieldQty, field.TypeFloat64, value)
	}
	if value, ok := pu.mutation.ImageUrls(); ok {
		_spec.SetField(product.FieldImageUrls, field.TypeJSON, value)
	}
	if value, ok := pu.mutation.AppendedImageUrls(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, product.FieldImageUrls, value)
		})
	}
	if pu.mutation.ImageUrlsCleared() {
		_spec.ClearField(product.FieldImageUrls, field.TypeJSON)
	}
	if value, ok := pu.mutation.UnitOfMeasurement(); ok {
		_spec.SetField(product.FieldUnitOfMeasurement, field.TypeString, value)
	}
	if pu.mutation.UnitOfMeasurementCleared() {
		_spec.ClearField(product.FieldUnitOfMeasurement, field.TypeString)
	}
	if value, ok := pu.mutation.Provider(); ok {
		_spec.SetField(product.FieldProvider, field.TypeString, value)
	}
	if pu.mutation.ProviderCleared() {
		_spec.ClearField(product.FieldProvider, field.TypeString)
	}
	if pu.mutation.ProductTypeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   product.ProductTypeTable,
			Columns: []string{product.ProductTypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(producttype.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.ProductTypeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   product.ProductTypeTable,
			Columns: []string{product.ProductTypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(producttype.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pu.mutation.TagsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   product.TagsTable,
			Columns: product.TagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tag.FieldID, field.TypeUUID),
			},
		}
		createE := &ProductTagCreate{config: pu.config, mutation: newProductTagMutation(pu.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.RemovedTagsIDs(); len(nodes) > 0 && !pu.mutation.TagsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   product.TagsTable,
			Columns: product.TagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tag.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &ProductTagCreate{config: pu.config, mutation: newProductTagMutation(pu.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.TagsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   product.TagsTable,
			Columns: product.TagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tag.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &ProductTagCreate{config: pu.config, mutation: newProductTagMutation(pu.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{product.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pu.mutation.done = true
	return n, nil
}

// ProductUpdateOne is the builder for updating a single Product entity.
type ProductUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ProductMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (puo *ProductUpdateOne) SetUpdatedAt(t time.Time) *ProductUpdateOne {
	puo.mutation.SetUpdatedAt(t)
	return puo
}

// SetName sets the "name" field.
func (puo *ProductUpdateOne) SetName(s string) *ProductUpdateOne {
	puo.mutation.SetName(s)
	return puo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (puo *ProductUpdateOne) SetNillableName(s *string) *ProductUpdateOne {
	if s != nil {
		puo.SetName(*s)
	}
	return puo
}

// SetDescription sets the "description" field.
func (puo *ProductUpdateOne) SetDescription(s string) *ProductUpdateOne {
	puo.mutation.SetDescription(s)
	return puo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (puo *ProductUpdateOne) SetNillableDescription(s *string) *ProductUpdateOne {
	if s != nil {
		puo.SetDescription(*s)
	}
	return puo
}

// ClearDescription clears the value of the "description" field.
func (puo *ProductUpdateOne) ClearDescription() *ProductUpdateOne {
	puo.mutation.ClearDescription()
	return puo
}

// SetYear sets the "year" field.
func (puo *ProductUpdateOne) SetYear(i int) *ProductUpdateOne {
	puo.mutation.ResetYear()
	puo.mutation.SetYear(i)
	return puo
}

// SetNillableYear sets the "year" field if the given value is not nil.
func (puo *ProductUpdateOne) SetNillableYear(i *int) *ProductUpdateOne {
	if i != nil {
		puo.SetYear(*i)
	}
	return puo
}

// AddYear adds i to the "year" field.
func (puo *ProductUpdateOne) AddYear(i int) *ProductUpdateOne {
	puo.mutation.AddYear(i)
	return puo
}

// SetPrice sets the "price" field.
func (puo *ProductUpdateOne) SetPrice(d decimal.Decimal) *ProductUpdateOne {
	puo.mutation.ResetPrice()
	puo.mutation.SetPrice(d)
	return puo
}

// SetNillablePrice sets the "price" field if the given value is not nil.
func (puo *ProductUpdateOne) SetNillablePrice(d *decimal.Decimal) *ProductUpdateOne {
	if d != nil {
		puo.SetPrice(*d)
	}
	return puo
}

// AddPrice adds d to the "price" field.
func (puo *ProductUpdateOne) AddPrice(d decimal.Decimal) *ProductUpdateOne {
	puo.mutation.AddPrice(d)
	return puo
}

// SetQty sets the "qty" field.
func (puo *ProductUpdateOne) SetQty(d decimal.Decimal) *ProductUpdateOne {
	puo.mutation.ResetQty()
	puo.mutation.SetQty(d)
	return puo
}

// SetNillableQty sets the "qty" field if the given value is not nil.
func (puo *ProductUpdateOne) SetNillableQty(d *decimal.Decimal) *ProductUpdateOne {
	if d != nil {
		puo.SetQty(*d)
	}
	return puo
}

// AddQty adds d to the "qty" field.
func (puo *ProductUpdateOne) AddQty(d decimal.Decimal) *ProductUpdateOne {
	puo.mutation.AddQty(d)
	return puo
}

// SetImageUrls sets the "image_urls" field.
func (puo *ProductUpdateOne) SetImageUrls(u []url.URL) *ProductUpdateOne {
	puo.mutation.SetImageUrls(u)
	return puo
}

// AppendImageUrls appends u to the "image_urls" field.
func (puo *ProductUpdateOne) AppendImageUrls(u []url.URL) *ProductUpdateOne {
	puo.mutation.AppendImageUrls(u)
	return puo
}

// ClearImageUrls clears the value of the "image_urls" field.
func (puo *ProductUpdateOne) ClearImageUrls() *ProductUpdateOne {
	puo.mutation.ClearImageUrls()
	return puo
}

// SetUnitOfMeasurement sets the "unit_of_measurement" field.
func (puo *ProductUpdateOne) SetUnitOfMeasurement(s string) *ProductUpdateOne {
	puo.mutation.SetUnitOfMeasurement(s)
	return puo
}

// SetNillableUnitOfMeasurement sets the "unit_of_measurement" field if the given value is not nil.
func (puo *ProductUpdateOne) SetNillableUnitOfMeasurement(s *string) *ProductUpdateOne {
	if s != nil {
		puo.SetUnitOfMeasurement(*s)
	}
	return puo
}

// ClearUnitOfMeasurement clears the value of the "unit_of_measurement" field.
func (puo *ProductUpdateOne) ClearUnitOfMeasurement() *ProductUpdateOne {
	puo.mutation.ClearUnitOfMeasurement()
	return puo
}

// SetType sets the "type" field.
func (puo *ProductUpdateOne) SetType(i int) *ProductUpdateOne {
	puo.mutation.SetType(i)
	return puo
}

// SetNillableType sets the "type" field if the given value is not nil.
func (puo *ProductUpdateOne) SetNillableType(i *int) *ProductUpdateOne {
	if i != nil {
		puo.SetType(*i)
	}
	return puo
}

// SetProvider sets the "provider" field.
func (puo *ProductUpdateOne) SetProvider(s string) *ProductUpdateOne {
	puo.mutation.SetProvider(s)
	return puo
}

// SetNillableProvider sets the "provider" field if the given value is not nil.
func (puo *ProductUpdateOne) SetNillableProvider(s *string) *ProductUpdateOne {
	if s != nil {
		puo.SetProvider(*s)
	}
	return puo
}

// ClearProvider clears the value of the "provider" field.
func (puo *ProductUpdateOne) ClearProvider() *ProductUpdateOne {
	puo.mutation.ClearProvider()
	return puo
}

// SetProductTypeID sets the "product_type" edge to the ProductType entity by ID.
func (puo *ProductUpdateOne) SetProductTypeID(id int) *ProductUpdateOne {
	puo.mutation.SetProductTypeID(id)
	return puo
}

// SetProductType sets the "product_type" edge to the ProductType entity.
func (puo *ProductUpdateOne) SetProductType(p *ProductType) *ProductUpdateOne {
	return puo.SetProductTypeID(p.ID)
}

// AddTagIDs adds the "tags" edge to the Tag entity by IDs.
func (puo *ProductUpdateOne) AddTagIDs(ids ...uuid.UUID) *ProductUpdateOne {
	puo.mutation.AddTagIDs(ids...)
	return puo
}

// AddTags adds the "tags" edges to the Tag entity.
func (puo *ProductUpdateOne) AddTags(t ...*Tag) *ProductUpdateOne {
	ids := make([]uuid.UUID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return puo.AddTagIDs(ids...)
}

// Mutation returns the ProductMutation object of the builder.
func (puo *ProductUpdateOne) Mutation() *ProductMutation {
	return puo.mutation
}

// ClearProductType clears the "product_type" edge to the ProductType entity.
func (puo *ProductUpdateOne) ClearProductType() *ProductUpdateOne {
	puo.mutation.ClearProductType()
	return puo
}

// ClearTags clears all "tags" edges to the Tag entity.
func (puo *ProductUpdateOne) ClearTags() *ProductUpdateOne {
	puo.mutation.ClearTags()
	return puo
}

// RemoveTagIDs removes the "tags" edge to Tag entities by IDs.
func (puo *ProductUpdateOne) RemoveTagIDs(ids ...uuid.UUID) *ProductUpdateOne {
	puo.mutation.RemoveTagIDs(ids...)
	return puo
}

// RemoveTags removes "tags" edges to Tag entities.
func (puo *ProductUpdateOne) RemoveTags(t ...*Tag) *ProductUpdateOne {
	ids := make([]uuid.UUID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return puo.RemoveTagIDs(ids...)
}

// Where appends a list predicates to the ProductUpdate builder.
func (puo *ProductUpdateOne) Where(ps ...predicate.Product) *ProductUpdateOne {
	puo.mutation.Where(ps...)
	return puo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *ProductUpdateOne) Select(field string, fields ...string) *ProductUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Product entity.
func (puo *ProductUpdateOne) Save(ctx context.Context) (*Product, error) {
	puo.defaults()
	return withHooks(ctx, puo.sqlSave, puo.mutation, puo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (puo *ProductUpdateOne) SaveX(ctx context.Context) *Product {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *ProductUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *ProductUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (puo *ProductUpdateOne) defaults() {
	if _, ok := puo.mutation.UpdatedAt(); !ok {
		v := product.UpdateDefaultUpdatedAt()
		puo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (puo *ProductUpdateOne) check() error {
	if v, ok := puo.mutation.Name(); ok {
		if err := product.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Product.name": %w`, err)}
		}
	}
	if v, ok := puo.mutation.Year(); ok {
		if err := product.YearValidator(v); err != nil {
			return &ValidationError{Name: "year", err: fmt.Errorf(`ent: validator failed for field "Product.year": %w`, err)}
		}
	}
	if _, ok := puo.mutation.ProductTypeID(); puo.mutation.ProductTypeCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Product.product_type"`)
	}
	return nil
}

func (puo *ProductUpdateOne) sqlSave(ctx context.Context) (_node *Product, err error) {
	if err := puo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(product.Table, product.Columns, sqlgraph.NewFieldSpec(product.FieldID, field.TypeUUID))
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Product.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, product.FieldID)
		for _, f := range fields {
			if !product.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != product.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := puo.mutation.UpdatedAt(); ok {
		_spec.SetField(product.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := puo.mutation.Name(); ok {
		_spec.SetField(product.FieldName, field.TypeString, value)
	}
	if value, ok := puo.mutation.Description(); ok {
		_spec.SetField(product.FieldDescription, field.TypeString, value)
	}
	if puo.mutation.DescriptionCleared() {
		_spec.ClearField(product.FieldDescription, field.TypeString)
	}
	if value, ok := puo.mutation.Year(); ok {
		_spec.SetField(product.FieldYear, field.TypeInt, value)
	}
	if value, ok := puo.mutation.AddedYear(); ok {
		_spec.AddField(product.FieldYear, field.TypeInt, value)
	}
	if value, ok := puo.mutation.Price(); ok {
		_spec.SetField(product.FieldPrice, field.TypeFloat64, value)
	}
	if value, ok := puo.mutation.AddedPrice(); ok {
		_spec.AddField(product.FieldPrice, field.TypeFloat64, value)
	}
	if value, ok := puo.mutation.Qty(); ok {
		_spec.SetField(product.FieldQty, field.TypeFloat64, value)
	}
	if value, ok := puo.mutation.AddedQty(); ok {
		_spec.AddField(product.FieldQty, field.TypeFloat64, value)
	}
	if value, ok := puo.mutation.ImageUrls(); ok {
		_spec.SetField(product.FieldImageUrls, field.TypeJSON, value)
	}
	if value, ok := puo.mutation.AppendedImageUrls(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, product.FieldImageUrls, value)
		})
	}
	if puo.mutation.ImageUrlsCleared() {
		_spec.ClearField(product.FieldImageUrls, field.TypeJSON)
	}
	if value, ok := puo.mutation.UnitOfMeasurement(); ok {
		_spec.SetField(product.FieldUnitOfMeasurement, field.TypeString, value)
	}
	if puo.mutation.UnitOfMeasurementCleared() {
		_spec.ClearField(product.FieldUnitOfMeasurement, field.TypeString)
	}
	if value, ok := puo.mutation.Provider(); ok {
		_spec.SetField(product.FieldProvider, field.TypeString, value)
	}
	if puo.mutation.ProviderCleared() {
		_spec.ClearField(product.FieldProvider, field.TypeString)
	}
	if puo.mutation.ProductTypeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   product.ProductTypeTable,
			Columns: []string{product.ProductTypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(producttype.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.ProductTypeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   product.ProductTypeTable,
			Columns: []string{product.ProductTypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(producttype.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if puo.mutation.TagsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   product.TagsTable,
			Columns: product.TagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tag.FieldID, field.TypeUUID),
			},
		}
		createE := &ProductTagCreate{config: puo.config, mutation: newProductTagMutation(puo.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.RemovedTagsIDs(); len(nodes) > 0 && !puo.mutation.TagsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   product.TagsTable,
			Columns: product.TagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tag.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &ProductTagCreate{config: puo.config, mutation: newProductTagMutation(puo.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.TagsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   product.TagsTable,
			Columns: product.TagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tag.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &ProductTagCreate{config: puo.config, mutation: newProductTagMutation(puo.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Product{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{product.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	puo.mutation.done = true
	return _node, nil
}
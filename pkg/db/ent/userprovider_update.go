// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/user-management/pkg/db/ent/predicate"
	"github.com/NpoolPlatform/user-management/pkg/db/ent/userprovider"
	"github.com/google/uuid"
)

// UserProviderUpdate is the builder for updating UserProvider entities.
type UserProviderUpdate struct {
	config
	hooks    []Hook
	mutation *UserProviderMutation
}

// Where appends a list predicates to the UserProviderUpdate builder.
func (upu *UserProviderUpdate) Where(ps ...predicate.UserProvider) *UserProviderUpdate {
	upu.mutation.Where(ps...)
	return upu
}

// SetUserID sets the "user_id" field.
func (upu *UserProviderUpdate) SetUserID(u uuid.UUID) *UserProviderUpdate {
	upu.mutation.SetUserID(u)
	return upu
}

// SetProviderID sets the "provider_id" field.
func (upu *UserProviderUpdate) SetProviderID(u uuid.UUID) *UserProviderUpdate {
	upu.mutation.SetProviderID(u)
	return upu
}

// SetProviderUserID sets the "provider_user_id" field.
func (upu *UserProviderUpdate) SetProviderUserID(s string) *UserProviderUpdate {
	upu.mutation.SetProviderUserID(s)
	return upu
}

// SetUserProviderInfo sets the "user_provider_info" field.
func (upu *UserProviderUpdate) SetUserProviderInfo(s string) *UserProviderUpdate {
	upu.mutation.SetUserProviderInfo(s)
	return upu
}

// SetCreateAt sets the "create_at" field.
func (upu *UserProviderUpdate) SetCreateAt(u uint32) *UserProviderUpdate {
	upu.mutation.ResetCreateAt()
	upu.mutation.SetCreateAt(u)
	return upu
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (upu *UserProviderUpdate) SetNillableCreateAt(u *uint32) *UserProviderUpdate {
	if u != nil {
		upu.SetCreateAt(*u)
	}
	return upu
}

// AddCreateAt adds u to the "create_at" field.
func (upu *UserProviderUpdate) AddCreateAt(u int32) *UserProviderUpdate {
	upu.mutation.AddCreateAt(u)
	return upu
}

// SetUpdateAt sets the "update_at" field.
func (upu *UserProviderUpdate) SetUpdateAt(u uint32) *UserProviderUpdate {
	upu.mutation.ResetUpdateAt()
	upu.mutation.SetUpdateAt(u)
	return upu
}

// AddUpdateAt adds u to the "update_at" field.
func (upu *UserProviderUpdate) AddUpdateAt(u int32) *UserProviderUpdate {
	upu.mutation.AddUpdateAt(u)
	return upu
}

// SetDeleteAt sets the "delete_at" field.
func (upu *UserProviderUpdate) SetDeleteAt(u uint32) *UserProviderUpdate {
	upu.mutation.ResetDeleteAt()
	upu.mutation.SetDeleteAt(u)
	return upu
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (upu *UserProviderUpdate) SetNillableDeleteAt(u *uint32) *UserProviderUpdate {
	if u != nil {
		upu.SetDeleteAt(*u)
	}
	return upu
}

// AddDeleteAt adds u to the "delete_at" field.
func (upu *UserProviderUpdate) AddDeleteAt(u int32) *UserProviderUpdate {
	upu.mutation.AddDeleteAt(u)
	return upu
}

// Mutation returns the UserProviderMutation object of the builder.
func (upu *UserProviderUpdate) Mutation() *UserProviderMutation {
	return upu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (upu *UserProviderUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	upu.defaults()
	if len(upu.hooks) == 0 {
		affected, err = upu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserProviderMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			upu.mutation = mutation
			affected, err = upu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(upu.hooks) - 1; i >= 0; i-- {
			if upu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = upu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, upu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (upu *UserProviderUpdate) SaveX(ctx context.Context) int {
	affected, err := upu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (upu *UserProviderUpdate) Exec(ctx context.Context) error {
	_, err := upu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (upu *UserProviderUpdate) ExecX(ctx context.Context) {
	if err := upu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (upu *UserProviderUpdate) defaults() {
	if _, ok := upu.mutation.UpdateAt(); !ok {
		v := userprovider.UpdateDefaultUpdateAt()
		upu.mutation.SetUpdateAt(v)
	}
}

func (upu *UserProviderUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   userprovider.Table,
			Columns: userprovider.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: userprovider.FieldID,
			},
		},
	}
	if ps := upu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := upu.mutation.UserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: userprovider.FieldUserID,
		})
	}
	if value, ok := upu.mutation.ProviderID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: userprovider.FieldProviderID,
		})
	}
	if value, ok := upu.mutation.ProviderUserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: userprovider.FieldProviderUserID,
		})
	}
	if value, ok := upu.mutation.UserProviderInfo(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: userprovider.FieldUserProviderInfo,
		})
	}
	if value, ok := upu.mutation.CreateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: userprovider.FieldCreateAt,
		})
	}
	if value, ok := upu.mutation.AddedCreateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: userprovider.FieldCreateAt,
		})
	}
	if value, ok := upu.mutation.UpdateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: userprovider.FieldUpdateAt,
		})
	}
	if value, ok := upu.mutation.AddedUpdateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: userprovider.FieldUpdateAt,
		})
	}
	if value, ok := upu.mutation.DeleteAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: userprovider.FieldDeleteAt,
		})
	}
	if value, ok := upu.mutation.AddedDeleteAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: userprovider.FieldDeleteAt,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, upu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{userprovider.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// UserProviderUpdateOne is the builder for updating a single UserProvider entity.
type UserProviderUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserProviderMutation
}

// SetUserID sets the "user_id" field.
func (upuo *UserProviderUpdateOne) SetUserID(u uuid.UUID) *UserProviderUpdateOne {
	upuo.mutation.SetUserID(u)
	return upuo
}

// SetProviderID sets the "provider_id" field.
func (upuo *UserProviderUpdateOne) SetProviderID(u uuid.UUID) *UserProviderUpdateOne {
	upuo.mutation.SetProviderID(u)
	return upuo
}

// SetProviderUserID sets the "provider_user_id" field.
func (upuo *UserProviderUpdateOne) SetProviderUserID(s string) *UserProviderUpdateOne {
	upuo.mutation.SetProviderUserID(s)
	return upuo
}

// SetUserProviderInfo sets the "user_provider_info" field.
func (upuo *UserProviderUpdateOne) SetUserProviderInfo(s string) *UserProviderUpdateOne {
	upuo.mutation.SetUserProviderInfo(s)
	return upuo
}

// SetCreateAt sets the "create_at" field.
func (upuo *UserProviderUpdateOne) SetCreateAt(u uint32) *UserProviderUpdateOne {
	upuo.mutation.ResetCreateAt()
	upuo.mutation.SetCreateAt(u)
	return upuo
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (upuo *UserProviderUpdateOne) SetNillableCreateAt(u *uint32) *UserProviderUpdateOne {
	if u != nil {
		upuo.SetCreateAt(*u)
	}
	return upuo
}

// AddCreateAt adds u to the "create_at" field.
func (upuo *UserProviderUpdateOne) AddCreateAt(u int32) *UserProviderUpdateOne {
	upuo.mutation.AddCreateAt(u)
	return upuo
}

// SetUpdateAt sets the "update_at" field.
func (upuo *UserProviderUpdateOne) SetUpdateAt(u uint32) *UserProviderUpdateOne {
	upuo.mutation.ResetUpdateAt()
	upuo.mutation.SetUpdateAt(u)
	return upuo
}

// AddUpdateAt adds u to the "update_at" field.
func (upuo *UserProviderUpdateOne) AddUpdateAt(u int32) *UserProviderUpdateOne {
	upuo.mutation.AddUpdateAt(u)
	return upuo
}

// SetDeleteAt sets the "delete_at" field.
func (upuo *UserProviderUpdateOne) SetDeleteAt(u uint32) *UserProviderUpdateOne {
	upuo.mutation.ResetDeleteAt()
	upuo.mutation.SetDeleteAt(u)
	return upuo
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (upuo *UserProviderUpdateOne) SetNillableDeleteAt(u *uint32) *UserProviderUpdateOne {
	if u != nil {
		upuo.SetDeleteAt(*u)
	}
	return upuo
}

// AddDeleteAt adds u to the "delete_at" field.
func (upuo *UserProviderUpdateOne) AddDeleteAt(u int32) *UserProviderUpdateOne {
	upuo.mutation.AddDeleteAt(u)
	return upuo
}

// Mutation returns the UserProviderMutation object of the builder.
func (upuo *UserProviderUpdateOne) Mutation() *UserProviderMutation {
	return upuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (upuo *UserProviderUpdateOne) Select(field string, fields ...string) *UserProviderUpdateOne {
	upuo.fields = append([]string{field}, fields...)
	return upuo
}

// Save executes the query and returns the updated UserProvider entity.
func (upuo *UserProviderUpdateOne) Save(ctx context.Context) (*UserProvider, error) {
	var (
		err  error
		node *UserProvider
	)
	upuo.defaults()
	if len(upuo.hooks) == 0 {
		node, err = upuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserProviderMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			upuo.mutation = mutation
			node, err = upuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(upuo.hooks) - 1; i >= 0; i-- {
			if upuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = upuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, upuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (upuo *UserProviderUpdateOne) SaveX(ctx context.Context) *UserProvider {
	node, err := upuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (upuo *UserProviderUpdateOne) Exec(ctx context.Context) error {
	_, err := upuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (upuo *UserProviderUpdateOne) ExecX(ctx context.Context) {
	if err := upuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (upuo *UserProviderUpdateOne) defaults() {
	if _, ok := upuo.mutation.UpdateAt(); !ok {
		v := userprovider.UpdateDefaultUpdateAt()
		upuo.mutation.SetUpdateAt(v)
	}
}

func (upuo *UserProviderUpdateOne) sqlSave(ctx context.Context) (_node *UserProvider, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   userprovider.Table,
			Columns: userprovider.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: userprovider.FieldID,
			},
		},
	}
	id, ok := upuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "UserProvider.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := upuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, userprovider.FieldID)
		for _, f := range fields {
			if !userprovider.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != userprovider.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := upuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := upuo.mutation.UserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: userprovider.FieldUserID,
		})
	}
	if value, ok := upuo.mutation.ProviderID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: userprovider.FieldProviderID,
		})
	}
	if value, ok := upuo.mutation.ProviderUserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: userprovider.FieldProviderUserID,
		})
	}
	if value, ok := upuo.mutation.UserProviderInfo(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: userprovider.FieldUserProviderInfo,
		})
	}
	if value, ok := upuo.mutation.CreateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: userprovider.FieldCreateAt,
		})
	}
	if value, ok := upuo.mutation.AddedCreateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: userprovider.FieldCreateAt,
		})
	}
	if value, ok := upuo.mutation.UpdateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: userprovider.FieldUpdateAt,
		})
	}
	if value, ok := upuo.mutation.AddedUpdateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: userprovider.FieldUpdateAt,
		})
	}
	if value, ok := upuo.mutation.DeleteAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: userprovider.FieldDeleteAt,
		})
	}
	if value, ok := upuo.mutation.AddedDeleteAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: userprovider.FieldDeleteAt,
		})
	}
	_node = &UserProvider{config: upuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, upuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{userprovider.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}

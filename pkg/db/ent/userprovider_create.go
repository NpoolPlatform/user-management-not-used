// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/user-management/pkg/db/ent/userprovider"
	"github.com/google/uuid"
)

// UserProviderCreate is the builder for creating a UserProvider entity.
type UserProviderCreate struct {
	config
	mutation *UserProviderMutation
	hooks    []Hook
}

// SetUserID sets the "user_id" field.
func (upc *UserProviderCreate) SetUserID(u uuid.UUID) *UserProviderCreate {
	upc.mutation.SetUserID(u)
	return upc
}

// SetProviderID sets the "provider_id" field.
func (upc *UserProviderCreate) SetProviderID(u uuid.UUID) *UserProviderCreate {
	upc.mutation.SetProviderID(u)
	return upc
}

// SetProviderUserID sets the "provider_user_id" field.
func (upc *UserProviderCreate) SetProviderUserID(s string) *UserProviderCreate {
	upc.mutation.SetProviderUserID(s)
	return upc
}

// SetUserProviderInfo sets the "user_provider_info" field.
func (upc *UserProviderCreate) SetUserProviderInfo(s string) *UserProviderCreate {
	upc.mutation.SetUserProviderInfo(s)
	return upc
}

// SetCreateAt sets the "create_at" field.
func (upc *UserProviderCreate) SetCreateAt(i int64) *UserProviderCreate {
	upc.mutation.SetCreateAt(i)
	return upc
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (upc *UserProviderCreate) SetNillableCreateAt(i *int64) *UserProviderCreate {
	if i != nil {
		upc.SetCreateAt(*i)
	}
	return upc
}

// SetUpdateAt sets the "update_at" field.
func (upc *UserProviderCreate) SetUpdateAt(i int64) *UserProviderCreate {
	upc.mutation.SetUpdateAt(i)
	return upc
}

// SetNillableUpdateAt sets the "update_at" field if the given value is not nil.
func (upc *UserProviderCreate) SetNillableUpdateAt(i *int64) *UserProviderCreate {
	if i != nil {
		upc.SetUpdateAt(*i)
	}
	return upc
}

// SetDeleteAt sets the "delete_at" field.
func (upc *UserProviderCreate) SetDeleteAt(i int64) *UserProviderCreate {
	upc.mutation.SetDeleteAt(i)
	return upc
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (upc *UserProviderCreate) SetNillableDeleteAt(i *int64) *UserProviderCreate {
	if i != nil {
		upc.SetDeleteAt(*i)
	}
	return upc
}

// SetID sets the "id" field.
func (upc *UserProviderCreate) SetID(u uuid.UUID) *UserProviderCreate {
	upc.mutation.SetID(u)
	return upc
}

// Mutation returns the UserProviderMutation object of the builder.
func (upc *UserProviderCreate) Mutation() *UserProviderMutation {
	return upc.mutation
}

// Save creates the UserProvider in the database.
func (upc *UserProviderCreate) Save(ctx context.Context) (*UserProvider, error) {
	var (
		err  error
		node *UserProvider
	)
	upc.defaults()
	if len(upc.hooks) == 0 {
		if err = upc.check(); err != nil {
			return nil, err
		}
		node, err = upc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserProviderMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = upc.check(); err != nil {
				return nil, err
			}
			upc.mutation = mutation
			if node, err = upc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(upc.hooks) - 1; i >= 0; i-- {
			if upc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = upc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, upc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (upc *UserProviderCreate) SaveX(ctx context.Context) *UserProvider {
	v, err := upc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (upc *UserProviderCreate) Exec(ctx context.Context) error {
	_, err := upc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (upc *UserProviderCreate) ExecX(ctx context.Context) {
	if err := upc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (upc *UserProviderCreate) defaults() {
	if _, ok := upc.mutation.CreateAt(); !ok {
		v := userprovider.DefaultCreateAt()
		upc.mutation.SetCreateAt(v)
	}
	if _, ok := upc.mutation.UpdateAt(); !ok {
		v := userprovider.DefaultUpdateAt()
		upc.mutation.SetUpdateAt(v)
	}
	if _, ok := upc.mutation.DeleteAt(); !ok {
		v := userprovider.DefaultDeleteAt()
		upc.mutation.SetDeleteAt(v)
	}
	if _, ok := upc.mutation.ID(); !ok {
		v := userprovider.DefaultID()
		upc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (upc *UserProviderCreate) check() error {
	if _, ok := upc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "user_id"`)}
	}
	if _, ok := upc.mutation.ProviderID(); !ok {
		return &ValidationError{Name: "provider_id", err: errors.New(`ent: missing required field "provider_id"`)}
	}
	if _, ok := upc.mutation.ProviderUserID(); !ok {
		return &ValidationError{Name: "provider_user_id", err: errors.New(`ent: missing required field "provider_user_id"`)}
	}
	if _, ok := upc.mutation.UserProviderInfo(); !ok {
		return &ValidationError{Name: "user_provider_info", err: errors.New(`ent: missing required field "user_provider_info"`)}
	}
	if _, ok := upc.mutation.CreateAt(); !ok {
		return &ValidationError{Name: "create_at", err: errors.New(`ent: missing required field "create_at"`)}
	}
	if _, ok := upc.mutation.UpdateAt(); !ok {
		return &ValidationError{Name: "update_at", err: errors.New(`ent: missing required field "update_at"`)}
	}
	if _, ok := upc.mutation.DeleteAt(); !ok {
		return &ValidationError{Name: "delete_at", err: errors.New(`ent: missing required field "delete_at"`)}
	}
	return nil
}

func (upc *UserProviderCreate) sqlSave(ctx context.Context) (*UserProvider, error) {
	_node, _spec := upc.createSpec()
	if err := sqlgraph.CreateNode(ctx, upc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		_node.ID = _spec.ID.Value.(uuid.UUID)
	}
	return _node, nil
}

func (upc *UserProviderCreate) createSpec() (*UserProvider, *sqlgraph.CreateSpec) {
	var (
		_node = &UserProvider{config: upc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: userprovider.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: userprovider.FieldID,
			},
		}
	)
	if id, ok := upc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := upc.mutation.UserID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: userprovider.FieldUserID,
		})
		_node.UserID = value
	}
	if value, ok := upc.mutation.ProviderID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: userprovider.FieldProviderID,
		})
		_node.ProviderID = value
	}
	if value, ok := upc.mutation.ProviderUserID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: userprovider.FieldProviderUserID,
		})
		_node.ProviderUserID = value
	}
	if value, ok := upc.mutation.UserProviderInfo(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: userprovider.FieldUserProviderInfo,
		})
		_node.UserProviderInfo = value
	}
	if value, ok := upc.mutation.CreateAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: userprovider.FieldCreateAt,
		})
		_node.CreateAt = value
	}
	if value, ok := upc.mutation.UpdateAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: userprovider.FieldUpdateAt,
		})
		_node.UpdateAt = value
	}
	if value, ok := upc.mutation.DeleteAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: userprovider.FieldDeleteAt,
		})
		_node.DeleteAt = value
	}
	return _node, _spec
}

// UserProviderCreateBulk is the builder for creating many UserProvider entities in bulk.
type UserProviderCreateBulk struct {
	config
	builders []*UserProviderCreate
}

// Save creates the UserProvider entities in the database.
func (upcb *UserProviderCreateBulk) Save(ctx context.Context) ([]*UserProvider, error) {
	specs := make([]*sqlgraph.CreateSpec, len(upcb.builders))
	nodes := make([]*UserProvider, len(upcb.builders))
	mutators := make([]Mutator, len(upcb.builders))
	for i := range upcb.builders {
		func(i int, root context.Context) {
			builder := upcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserProviderMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, upcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, upcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, upcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (upcb *UserProviderCreateBulk) SaveX(ctx context.Context) []*UserProvider {
	v, err := upcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (upcb *UserProviderCreateBulk) Exec(ctx context.Context) error {
	_, err := upcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (upcb *UserProviderCreateBulk) ExecX(ctx context.Context) {
	if err := upcb.Exec(ctx); err != nil {
		panic(err)
	}
}

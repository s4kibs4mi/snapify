// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/s4kibs4mi/snapify/ent/screenshot"
	"github.com/s4kibs4mi/snapify/models"
)

// ScreenshotCreate is the builder for creating a Screenshot entity.
type ScreenshotCreate struct {
	config
	mutation *ScreenshotMutation
	hooks    []Hook
}

// SetStatus sets the "status" field.
func (sc *ScreenshotCreate) SetStatus(m models.Status) *ScreenshotCreate {
	sc.mutation.SetStatus(m)
	return sc
}

// SetURL sets the "url" field.
func (sc *ScreenshotCreate) SetURL(s string) *ScreenshotCreate {
	sc.mutation.SetURL(s)
	return sc
}

// SetStoredPath sets the "stored_path" field.
func (sc *ScreenshotCreate) SetStoredPath(s string) *ScreenshotCreate {
	sc.mutation.SetStoredPath(s)
	return sc
}

// SetNillableStoredPath sets the "stored_path" field if the given value is not nil.
func (sc *ScreenshotCreate) SetNillableStoredPath(s *string) *ScreenshotCreate {
	if s != nil {
		sc.SetStoredPath(*s)
	}
	return sc
}

// SetCreatedAt sets the "created_at" field.
func (sc *ScreenshotCreate) SetCreatedAt(t time.Time) *ScreenshotCreate {
	sc.mutation.SetCreatedAt(t)
	return sc
}

// SetID sets the "id" field.
func (sc *ScreenshotCreate) SetID(u uuid.UUID) *ScreenshotCreate {
	sc.mutation.SetID(u)
	return sc
}

// Mutation returns the ScreenshotMutation object of the builder.
func (sc *ScreenshotCreate) Mutation() *ScreenshotMutation {
	return sc.mutation
}

// Save creates the Screenshot in the database.
func (sc *ScreenshotCreate) Save(ctx context.Context) (*Screenshot, error) {
	return withHooks[*Screenshot, ScreenshotMutation](ctx, sc.sqlSave, sc.mutation, sc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (sc *ScreenshotCreate) SaveX(ctx context.Context) *Screenshot {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sc *ScreenshotCreate) Exec(ctx context.Context) error {
	_, err := sc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sc *ScreenshotCreate) ExecX(ctx context.Context) {
	if err := sc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sc *ScreenshotCreate) check() error {
	if _, ok := sc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "Screenshot.status"`)}
	}
	if _, ok := sc.mutation.URL(); !ok {
		return &ValidationError{Name: "url", err: errors.New(`ent: missing required field "Screenshot.url"`)}
	}
	if _, ok := sc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Screenshot.created_at"`)}
	}
	return nil
}

func (sc *ScreenshotCreate) sqlSave(ctx context.Context) (*Screenshot, error) {
	if err := sc.check(); err != nil {
		return nil, err
	}
	_node, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	sc.mutation.id = &_node.ID
	sc.mutation.done = true
	return _node, nil
}

func (sc *ScreenshotCreate) createSpec() (*Screenshot, *sqlgraph.CreateSpec) {
	var (
		_node = &Screenshot{config: sc.config}
		_spec = sqlgraph.NewCreateSpec(screenshot.Table, sqlgraph.NewFieldSpec(screenshot.FieldID, field.TypeUUID))
	)
	if id, ok := sc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := sc.mutation.Status(); ok {
		_spec.SetField(screenshot.FieldStatus, field.TypeString, value)
		_node.Status = value
	}
	if value, ok := sc.mutation.URL(); ok {
		_spec.SetField(screenshot.FieldURL, field.TypeString, value)
		_node.URL = value
	}
	if value, ok := sc.mutation.StoredPath(); ok {
		_spec.SetField(screenshot.FieldStoredPath, field.TypeString, value)
		_node.StoredPath = &value
	}
	if value, ok := sc.mutation.CreatedAt(); ok {
		_spec.SetField(screenshot.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	return _node, _spec
}

// ScreenshotCreateBulk is the builder for creating many Screenshot entities in bulk.
type ScreenshotCreateBulk struct {
	config
	builders []*ScreenshotCreate
}

// Save creates the Screenshot entities in the database.
func (scb *ScreenshotCreateBulk) Save(ctx context.Context) ([]*Screenshot, error) {
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*Screenshot, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ScreenshotMutation)
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
					_, err = mutators[i+1].Mutate(root, scb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, scb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
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
		if _, err := mutators[0].Mutate(ctx, scb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (scb *ScreenshotCreateBulk) SaveX(ctx context.Context) []*Screenshot {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scb *ScreenshotCreateBulk) Exec(ctx context.Context) error {
	_, err := scb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scb *ScreenshotCreateBulk) ExecX(ctx context.Context) {
	if err := scb.Exec(ctx); err != nil {
		panic(err)
	}
}
// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/s4kibs4mi/snapify/ent/predicate"
	"github.com/s4kibs4mi/snapify/ent/screenshot"
)

// ScreenshotDelete is the builder for deleting a Screenshot entity.
type ScreenshotDelete struct {
	config
	hooks    []Hook
	mutation *ScreenshotMutation
}

// Where appends a list predicates to the ScreenshotDelete builder.
func (sd *ScreenshotDelete) Where(ps ...predicate.Screenshot) *ScreenshotDelete {
	sd.mutation.Where(ps...)
	return sd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (sd *ScreenshotDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, ScreenshotMutation](ctx, sd.sqlExec, sd.mutation, sd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (sd *ScreenshotDelete) ExecX(ctx context.Context) int {
	n, err := sd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (sd *ScreenshotDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(screenshot.Table, sqlgraph.NewFieldSpec(screenshot.FieldID, field.TypeUUID))
	if ps := sd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, sd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	sd.mutation.done = true
	return affected, err
}

// ScreenshotDeleteOne is the builder for deleting a single Screenshot entity.
type ScreenshotDeleteOne struct {
	sd *ScreenshotDelete
}

// Where appends a list predicates to the ScreenshotDelete builder.
func (sdo *ScreenshotDeleteOne) Where(ps ...predicate.Screenshot) *ScreenshotDeleteOne {
	sdo.sd.mutation.Where(ps...)
	return sdo
}

// Exec executes the deletion query.
func (sdo *ScreenshotDeleteOne) Exec(ctx context.Context) error {
	n, err := sdo.sd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{screenshot.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (sdo *ScreenshotDeleteOne) ExecX(ctx context.Context) {
	if err := sdo.Exec(ctx); err != nil {
		panic(err)
	}
}

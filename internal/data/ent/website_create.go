// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"go-server/internal/data/ent/website"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// WebsiteCreate is the builder for creating a Website entity.
type WebsiteCreate struct {
	config
	mutation *WebsiteMutation
	hooks    []Hook
}

// SetSortID sets the "sort_id" field.
func (wc *WebsiteCreate) SetSortID(i int32) *WebsiteCreate {
	wc.mutation.SetSortID(i)
	return wc
}

// SetNillableSortID sets the "sort_id" field if the given value is not nil.
func (wc *WebsiteCreate) SetNillableSortID(i *int32) *WebsiteCreate {
	if i != nil {
		wc.SetSortID(*i)
	}
	return wc
}

// SetCategory sets the "category" field.
func (wc *WebsiteCreate) SetCategory(i int32) *WebsiteCreate {
	wc.mutation.SetCategory(i)
	return wc
}

// SetNillableCategory sets the "category" field if the given value is not nil.
func (wc *WebsiteCreate) SetNillableCategory(i *int32) *WebsiteCreate {
	if i != nil {
		wc.SetCategory(*i)
	}
	return wc
}

// SetType sets the "type" field.
func (wc *WebsiteCreate) SetType(i int32) *WebsiteCreate {
	wc.mutation.SetType(i)
	return wc
}

// SetNillableType sets the "type" field if the given value is not nil.
func (wc *WebsiteCreate) SetNillableType(i *int32) *WebsiteCreate {
	if i != nil {
		wc.SetType(*i)
	}
	return wc
}

// SetWebsiteName sets the "website_name" field.
func (wc *WebsiteCreate) SetWebsiteName(s string) *WebsiteCreate {
	wc.mutation.SetWebsiteName(s)
	return wc
}

// SetNillableWebsiteName sets the "website_name" field if the given value is not nil.
func (wc *WebsiteCreate) SetNillableWebsiteName(s *string) *WebsiteCreate {
	if s != nil {
		wc.SetWebsiteName(*s)
	}
	return wc
}

// SetWebsiteIcon sets the "website_icon" field.
func (wc *WebsiteCreate) SetWebsiteIcon(s string) *WebsiteCreate {
	wc.mutation.SetWebsiteIcon(s)
	return wc
}

// SetNillableWebsiteIcon sets the "website_icon" field if the given value is not nil.
func (wc *WebsiteCreate) SetNillableWebsiteIcon(s *string) *WebsiteCreate {
	if s != nil {
		wc.SetWebsiteIcon(*s)
	}
	return wc
}

// SetWebsiteURL sets the "website_url" field.
func (wc *WebsiteCreate) SetWebsiteURL(s string) *WebsiteCreate {
	wc.mutation.SetWebsiteURL(s)
	return wc
}

// SetNillableWebsiteURL sets the "website_url" field if the given value is not nil.
func (wc *WebsiteCreate) SetNillableWebsiteURL(s *string) *WebsiteCreate {
	if s != nil {
		wc.SetWebsiteURL(*s)
	}
	return wc
}

// SetSummary sets the "summary" field.
func (wc *WebsiteCreate) SetSummary(s string) *WebsiteCreate {
	wc.mutation.SetSummary(s)
	return wc
}

// SetNillableSummary sets the "summary" field if the given value is not nil.
func (wc *WebsiteCreate) SetNillableSummary(s *string) *WebsiteCreate {
	if s != nil {
		wc.SetSummary(*s)
	}
	return wc
}

// SetDescription sets the "description" field.
func (wc *WebsiteCreate) SetDescription(s string) *WebsiteCreate {
	wc.mutation.SetDescription(s)
	return wc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (wc *WebsiteCreate) SetNillableDescription(s *string) *WebsiteCreate {
	if s != nil {
		wc.SetDescription(*s)
	}
	return wc
}

// SetStatus sets the "status" field.
func (wc *WebsiteCreate) SetStatus(s string) *WebsiteCreate {
	wc.mutation.SetStatus(s)
	return wc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (wc *WebsiteCreate) SetNillableStatus(s *string) *WebsiteCreate {
	if s != nil {
		wc.SetStatus(*s)
	}
	return wc
}

// SetCreateID sets the "create_id" field.
func (wc *WebsiteCreate) SetCreateID(i int64) *WebsiteCreate {
	wc.mutation.SetCreateID(i)
	return wc
}

// SetNillableCreateID sets the "create_id" field if the given value is not nil.
func (wc *WebsiteCreate) SetNillableCreateID(i *int64) *WebsiteCreate {
	if i != nil {
		wc.SetCreateID(*i)
	}
	return wc
}

// SetCreateTime sets the "create_time" field.
func (wc *WebsiteCreate) SetCreateTime(t time.Time) *WebsiteCreate {
	wc.mutation.SetCreateTime(t)
	return wc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (wc *WebsiteCreate) SetNillableCreateTime(t *time.Time) *WebsiteCreate {
	if t != nil {
		wc.SetCreateTime(*t)
	}
	return wc
}

// SetModifyID sets the "modify_id" field.
func (wc *WebsiteCreate) SetModifyID(i int64) *WebsiteCreate {
	wc.mutation.SetModifyID(i)
	return wc
}

// SetNillableModifyID sets the "modify_id" field if the given value is not nil.
func (wc *WebsiteCreate) SetNillableModifyID(i *int64) *WebsiteCreate {
	if i != nil {
		wc.SetModifyID(*i)
	}
	return wc
}

// SetModifyTime sets the "modify_time" field.
func (wc *WebsiteCreate) SetModifyTime(t time.Time) *WebsiteCreate {
	wc.mutation.SetModifyTime(t)
	return wc
}

// SetNillableModifyTime sets the "modify_time" field if the given value is not nil.
func (wc *WebsiteCreate) SetNillableModifyTime(t *time.Time) *WebsiteCreate {
	if t != nil {
		wc.SetModifyTime(*t)
	}
	return wc
}

// SetID sets the "id" field.
func (wc *WebsiteCreate) SetID(i int64) *WebsiteCreate {
	wc.mutation.SetID(i)
	return wc
}

// Mutation returns the WebsiteMutation object of the builder.
func (wc *WebsiteCreate) Mutation() *WebsiteMutation {
	return wc.mutation
}

// Save creates the Website in the database.
func (wc *WebsiteCreate) Save(ctx context.Context) (*Website, error) {
	wc.defaults()
	return withHooks[*Website, WebsiteMutation](ctx, wc.sqlSave, wc.mutation, wc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (wc *WebsiteCreate) SaveX(ctx context.Context) *Website {
	v, err := wc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (wc *WebsiteCreate) Exec(ctx context.Context) error {
	_, err := wc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wc *WebsiteCreate) ExecX(ctx context.Context) {
	if err := wc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (wc *WebsiteCreate) defaults() {
	if _, ok := wc.mutation.CreateTime(); !ok {
		v := website.DefaultCreateTime()
		wc.mutation.SetCreateTime(v)
	}
	if _, ok := wc.mutation.ModifyTime(); !ok {
		v := website.DefaultModifyTime()
		wc.mutation.SetModifyTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wc *WebsiteCreate) check() error {
	return nil
}

func (wc *WebsiteCreate) sqlSave(ctx context.Context) (*Website, error) {
	if err := wc.check(); err != nil {
		return nil, err
	}
	_node, _spec := wc.createSpec()
	if err := sqlgraph.CreateNode(ctx, wc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	wc.mutation.id = &_node.ID
	wc.mutation.done = true
	return _node, nil
}

func (wc *WebsiteCreate) createSpec() (*Website, *sqlgraph.CreateSpec) {
	var (
		_node = &Website{config: wc.config}
		_spec = sqlgraph.NewCreateSpec(website.Table, sqlgraph.NewFieldSpec(website.FieldID, field.TypeInt64))
	)
	if id, ok := wc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := wc.mutation.SortID(); ok {
		_spec.SetField(website.FieldSortID, field.TypeInt32, value)
		_node.SortID = value
	}
	if value, ok := wc.mutation.Category(); ok {
		_spec.SetField(website.FieldCategory, field.TypeInt32, value)
		_node.Category = value
	}
	if value, ok := wc.mutation.GetType(); ok {
		_spec.SetField(website.FieldType, field.TypeInt32, value)
		_node.Type = value
	}
	if value, ok := wc.mutation.WebsiteName(); ok {
		_spec.SetField(website.FieldWebsiteName, field.TypeString, value)
		_node.WebsiteName = value
	}
	if value, ok := wc.mutation.WebsiteIcon(); ok {
		_spec.SetField(website.FieldWebsiteIcon, field.TypeString, value)
		_node.WebsiteIcon = value
	}
	if value, ok := wc.mutation.WebsiteURL(); ok {
		_spec.SetField(website.FieldWebsiteURL, field.TypeString, value)
		_node.WebsiteURL = value
	}
	if value, ok := wc.mutation.Summary(); ok {
		_spec.SetField(website.FieldSummary, field.TypeString, value)
		_node.Summary = value
	}
	if value, ok := wc.mutation.Description(); ok {
		_spec.SetField(website.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := wc.mutation.Status(); ok {
		_spec.SetField(website.FieldStatus, field.TypeString, value)
		_node.Status = value
	}
	if value, ok := wc.mutation.CreateID(); ok {
		_spec.SetField(website.FieldCreateID, field.TypeInt64, value)
		_node.CreateID = value
	}
	if value, ok := wc.mutation.CreateTime(); ok {
		_spec.SetField(website.FieldCreateTime, field.TypeTime, value)
		_node.CreateTime = value
	}
	if value, ok := wc.mutation.ModifyID(); ok {
		_spec.SetField(website.FieldModifyID, field.TypeInt64, value)
		_node.ModifyID = value
	}
	if value, ok := wc.mutation.ModifyTime(); ok {
		_spec.SetField(website.FieldModifyTime, field.TypeTime, value)
		_node.ModifyTime = value
	}
	return _node, _spec
}

// WebsiteCreateBulk is the builder for creating many Website entities in bulk.
type WebsiteCreateBulk struct {
	config
	builders []*WebsiteCreate
}

// Save creates the Website entities in the database.
func (wcb *WebsiteCreateBulk) Save(ctx context.Context) ([]*Website, error) {
	specs := make([]*sqlgraph.CreateSpec, len(wcb.builders))
	nodes := make([]*Website, len(wcb.builders))
	mutators := make([]Mutator, len(wcb.builders))
	for i := range wcb.builders {
		func(i int, root context.Context) {
			builder := wcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*WebsiteMutation)
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
					_, err = mutators[i+1].Mutate(root, wcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, wcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int64(id)
				}
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
		if _, err := mutators[0].Mutate(ctx, wcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (wcb *WebsiteCreateBulk) SaveX(ctx context.Context) []*Website {
	v, err := wcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (wcb *WebsiteCreateBulk) Exec(ctx context.Context) error {
	_, err := wcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wcb *WebsiteCreateBulk) ExecX(ctx context.Context) {
	if err := wcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"go-server/internal/data/ent/article"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ArticleCreate is the builder for creating a Article entity.
type ArticleCreate struct {
	config
	mutation *ArticleMutation
	hooks    []Hook
}

// SetSortID sets the "sort_id" field.
func (ac *ArticleCreate) SetSortID(i int32) *ArticleCreate {
	ac.mutation.SetSortID(i)
	return ac
}

// SetNillableSortID sets the "sort_id" field if the given value is not nil.
func (ac *ArticleCreate) SetNillableSortID(i *int32) *ArticleCreate {
	if i != nil {
		ac.SetSortID(*i)
	}
	return ac
}

// SetName sets the "name" field.
func (ac *ArticleCreate) SetName(s string) *ArticleCreate {
	ac.mutation.SetName(s)
	return ac
}

// SetNillableName sets the "name" field if the given value is not nil.
func (ac *ArticleCreate) SetNillableName(s *string) *ArticleCreate {
	if s != nil {
		ac.SetName(*s)
	}
	return ac
}

// SetCategoryID sets the "category_id" field.
func (ac *ArticleCreate) SetCategoryID(i int32) *ArticleCreate {
	ac.mutation.SetCategoryID(i)
	return ac
}

// SetNillableCategoryID sets the "category_id" field if the given value is not nil.
func (ac *ArticleCreate) SetNillableCategoryID(i *int32) *ArticleCreate {
	if i != nil {
		ac.SetCategoryID(*i)
	}
	return ac
}

// SetRecommend sets the "recommend" field.
func (ac *ArticleCreate) SetRecommend(s string) *ArticleCreate {
	ac.mutation.SetRecommend(s)
	return ac
}

// SetNillableRecommend sets the "recommend" field if the given value is not nil.
func (ac *ArticleCreate) SetNillableRecommend(s *string) *ArticleCreate {
	if s != nil {
		ac.SetRecommend(*s)
	}
	return ac
}

// SetDescription sets the "description" field.
func (ac *ArticleCreate) SetDescription(s string) *ArticleCreate {
	ac.mutation.SetDescription(s)
	return ac
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (ac *ArticleCreate) SetNillableDescription(s *string) *ArticleCreate {
	if s != nil {
		ac.SetDescription(*s)
	}
	return ac
}

// SetContentBody sets the "content_body" field.
func (ac *ArticleCreate) SetContentBody(s string) *ArticleCreate {
	ac.mutation.SetContentBody(s)
	return ac
}

// SetNillableContentBody sets the "content_body" field if the given value is not nil.
func (ac *ArticleCreate) SetNillableContentBody(s *string) *ArticleCreate {
	if s != nil {
		ac.SetContentBody(*s)
	}
	return ac
}

// SetImageURL sets the "image_url" field.
func (ac *ArticleCreate) SetImageURL(s string) *ArticleCreate {
	ac.mutation.SetImageURL(s)
	return ac
}

// SetNillableImageURL sets the "image_url" field if the given value is not nil.
func (ac *ArticleCreate) SetNillableImageURL(s *string) *ArticleCreate {
	if s != nil {
		ac.SetImageURL(*s)
	}
	return ac
}

// SetStatus sets the "status" field.
func (ac *ArticleCreate) SetStatus(i int32) *ArticleCreate {
	ac.mutation.SetStatus(i)
	return ac
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (ac *ArticleCreate) SetNillableStatus(i *int32) *ArticleCreate {
	if i != nil {
		ac.SetStatus(*i)
	}
	return ac
}

// SetClickCount sets the "click_count" field.
func (ac *ArticleCreate) SetClickCount(i int32) *ArticleCreate {
	ac.mutation.SetClickCount(i)
	return ac
}

// SetNillableClickCount sets the "click_count" field if the given value is not nil.
func (ac *ArticleCreate) SetNillableClickCount(i *int32) *ArticleCreate {
	if i != nil {
		ac.SetClickCount(*i)
	}
	return ac
}

// SetLikeCount sets the "like_count" field.
func (ac *ArticleCreate) SetLikeCount(i int32) *ArticleCreate {
	ac.mutation.SetLikeCount(i)
	return ac
}

// SetNillableLikeCount sets the "like_count" field if the given value is not nil.
func (ac *ArticleCreate) SetNillableLikeCount(i *int32) *ArticleCreate {
	if i != nil {
		ac.SetLikeCount(*i)
	}
	return ac
}

// SetCreateID sets the "create_id" field.
func (ac *ArticleCreate) SetCreateID(i int64) *ArticleCreate {
	ac.mutation.SetCreateID(i)
	return ac
}

// SetNillableCreateID sets the "create_id" field if the given value is not nil.
func (ac *ArticleCreate) SetNillableCreateID(i *int64) *ArticleCreate {
	if i != nil {
		ac.SetCreateID(*i)
	}
	return ac
}

// SetCreateTime sets the "create_time" field.
func (ac *ArticleCreate) SetCreateTime(t time.Time) *ArticleCreate {
	ac.mutation.SetCreateTime(t)
	return ac
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (ac *ArticleCreate) SetNillableCreateTime(t *time.Time) *ArticleCreate {
	if t != nil {
		ac.SetCreateTime(*t)
	}
	return ac
}

// SetModifyID sets the "modify_id" field.
func (ac *ArticleCreate) SetModifyID(i int64) *ArticleCreate {
	ac.mutation.SetModifyID(i)
	return ac
}

// SetNillableModifyID sets the "modify_id" field if the given value is not nil.
func (ac *ArticleCreate) SetNillableModifyID(i *int64) *ArticleCreate {
	if i != nil {
		ac.SetModifyID(*i)
	}
	return ac
}

// SetModifyTime sets the "modify_time" field.
func (ac *ArticleCreate) SetModifyTime(t time.Time) *ArticleCreate {
	ac.mutation.SetModifyTime(t)
	return ac
}

// SetNillableModifyTime sets the "modify_time" field if the given value is not nil.
func (ac *ArticleCreate) SetNillableModifyTime(t *time.Time) *ArticleCreate {
	if t != nil {
		ac.SetModifyTime(*t)
	}
	return ac
}

// SetID sets the "id" field.
func (ac *ArticleCreate) SetID(i int32) *ArticleCreate {
	ac.mutation.SetID(i)
	return ac
}

// Mutation returns the ArticleMutation object of the builder.
func (ac *ArticleCreate) Mutation() *ArticleMutation {
	return ac.mutation
}

// Save creates the Article in the database.
func (ac *ArticleCreate) Save(ctx context.Context) (*Article, error) {
	ac.defaults()
	return withHooks[*Article, ArticleMutation](ctx, ac.sqlSave, ac.mutation, ac.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ac *ArticleCreate) SaveX(ctx context.Context) *Article {
	v, err := ac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ac *ArticleCreate) Exec(ctx context.Context) error {
	_, err := ac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ac *ArticleCreate) ExecX(ctx context.Context) {
	if err := ac.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ac *ArticleCreate) defaults() {
	if _, ok := ac.mutation.CreateTime(); !ok {
		v := article.DefaultCreateTime()
		ac.mutation.SetCreateTime(v)
	}
	if _, ok := ac.mutation.ModifyTime(); !ok {
		v := article.DefaultModifyTime()
		ac.mutation.SetModifyTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ac *ArticleCreate) check() error {
	return nil
}

func (ac *ArticleCreate) sqlSave(ctx context.Context) (*Article, error) {
	if err := ac.check(); err != nil {
		return nil, err
	}
	_node, _spec := ac.createSpec()
	if err := sqlgraph.CreateNode(ctx, ac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int32(id)
	}
	ac.mutation.id = &_node.ID
	ac.mutation.done = true
	return _node, nil
}

func (ac *ArticleCreate) createSpec() (*Article, *sqlgraph.CreateSpec) {
	var (
		_node = &Article{config: ac.config}
		_spec = sqlgraph.NewCreateSpec(article.Table, sqlgraph.NewFieldSpec(article.FieldID, field.TypeInt32))
	)
	if id, ok := ac.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := ac.mutation.SortID(); ok {
		_spec.SetField(article.FieldSortID, field.TypeInt32, value)
		_node.SortID = value
	}
	if value, ok := ac.mutation.Name(); ok {
		_spec.SetField(article.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := ac.mutation.CategoryID(); ok {
		_spec.SetField(article.FieldCategoryID, field.TypeInt32, value)
		_node.CategoryID = value
	}
	if value, ok := ac.mutation.Recommend(); ok {
		_spec.SetField(article.FieldRecommend, field.TypeString, value)
		_node.Recommend = value
	}
	if value, ok := ac.mutation.Description(); ok {
		_spec.SetField(article.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := ac.mutation.ContentBody(); ok {
		_spec.SetField(article.FieldContentBody, field.TypeString, value)
		_node.ContentBody = value
	}
	if value, ok := ac.mutation.ImageURL(); ok {
		_spec.SetField(article.FieldImageURL, field.TypeString, value)
		_node.ImageURL = value
	}
	if value, ok := ac.mutation.Status(); ok {
		_spec.SetField(article.FieldStatus, field.TypeInt32, value)
		_node.Status = value
	}
	if value, ok := ac.mutation.ClickCount(); ok {
		_spec.SetField(article.FieldClickCount, field.TypeInt32, value)
		_node.ClickCount = value
	}
	if value, ok := ac.mutation.LikeCount(); ok {
		_spec.SetField(article.FieldLikeCount, field.TypeInt32, value)
		_node.LikeCount = value
	}
	if value, ok := ac.mutation.CreateID(); ok {
		_spec.SetField(article.FieldCreateID, field.TypeInt64, value)
		_node.CreateID = value
	}
	if value, ok := ac.mutation.CreateTime(); ok {
		_spec.SetField(article.FieldCreateTime, field.TypeTime, value)
		_node.CreateTime = value
	}
	if value, ok := ac.mutation.ModifyID(); ok {
		_spec.SetField(article.FieldModifyID, field.TypeInt64, value)
		_node.ModifyID = value
	}
	if value, ok := ac.mutation.ModifyTime(); ok {
		_spec.SetField(article.FieldModifyTime, field.TypeTime, value)
		_node.ModifyTime = value
	}
	return _node, _spec
}

// ArticleCreateBulk is the builder for creating many Article entities in bulk.
type ArticleCreateBulk struct {
	config
	builders []*ArticleCreate
}

// Save creates the Article entities in the database.
func (acb *ArticleCreateBulk) Save(ctx context.Context) ([]*Article, error) {
	specs := make([]*sqlgraph.CreateSpec, len(acb.builders))
	nodes := make([]*Article, len(acb.builders))
	mutators := make([]Mutator, len(acb.builders))
	for i := range acb.builders {
		func(i int, root context.Context) {
			builder := acb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ArticleMutation)
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
					_, err = mutators[i+1].Mutate(root, acb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, acb.driver, spec); err != nil {
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
					nodes[i].ID = int32(id)
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
		if _, err := mutators[0].Mutate(ctx, acb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (acb *ArticleCreateBulk) SaveX(ctx context.Context) []*Article {
	v, err := acb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (acb *ArticleCreateBulk) Exec(ctx context.Context) error {
	_, err := acb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acb *ArticleCreateBulk) ExecX(ctx context.Context) {
	if err := acb.Exec(ctx); err != nil {
		panic(err)
	}
}

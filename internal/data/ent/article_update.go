// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"go-server/internal/data/ent/article"
	"go-server/internal/data/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ArticleUpdate is the builder for updating Article entities.
type ArticleUpdate struct {
	config
	hooks    []Hook
	mutation *ArticleMutation
}

// Where appends a list predicates to the ArticleUpdate builder.
func (au *ArticleUpdate) Where(ps ...predicate.Article) *ArticleUpdate {
	au.mutation.Where(ps...)
	return au
}

// SetSortID sets the "sort_id" field.
func (au *ArticleUpdate) SetSortID(i int32) *ArticleUpdate {
	au.mutation.ResetSortID()
	au.mutation.SetSortID(i)
	return au
}

// SetNillableSortID sets the "sort_id" field if the given value is not nil.
func (au *ArticleUpdate) SetNillableSortID(i *int32) *ArticleUpdate {
	if i != nil {
		au.SetSortID(*i)
	}
	return au
}

// AddSortID adds i to the "sort_id" field.
func (au *ArticleUpdate) AddSortID(i int32) *ArticleUpdate {
	au.mutation.AddSortID(i)
	return au
}

// ClearSortID clears the value of the "sort_id" field.
func (au *ArticleUpdate) ClearSortID() *ArticleUpdate {
	au.mutation.ClearSortID()
	return au
}

// SetName sets the "name" field.
func (au *ArticleUpdate) SetName(s string) *ArticleUpdate {
	au.mutation.SetName(s)
	return au
}

// SetNillableName sets the "name" field if the given value is not nil.
func (au *ArticleUpdate) SetNillableName(s *string) *ArticleUpdate {
	if s != nil {
		au.SetName(*s)
	}
	return au
}

// ClearName clears the value of the "name" field.
func (au *ArticleUpdate) ClearName() *ArticleUpdate {
	au.mutation.ClearName()
	return au
}

// SetCategoryID sets the "category_id" field.
func (au *ArticleUpdate) SetCategoryID(i int32) *ArticleUpdate {
	au.mutation.ResetCategoryID()
	au.mutation.SetCategoryID(i)
	return au
}

// SetNillableCategoryID sets the "category_id" field if the given value is not nil.
func (au *ArticleUpdate) SetNillableCategoryID(i *int32) *ArticleUpdate {
	if i != nil {
		au.SetCategoryID(*i)
	}
	return au
}

// AddCategoryID adds i to the "category_id" field.
func (au *ArticleUpdate) AddCategoryID(i int32) *ArticleUpdate {
	au.mutation.AddCategoryID(i)
	return au
}

// ClearCategoryID clears the value of the "category_id" field.
func (au *ArticleUpdate) ClearCategoryID() *ArticleUpdate {
	au.mutation.ClearCategoryID()
	return au
}

// SetRecommend sets the "recommend" field.
func (au *ArticleUpdate) SetRecommend(s string) *ArticleUpdate {
	au.mutation.SetRecommend(s)
	return au
}

// SetNillableRecommend sets the "recommend" field if the given value is not nil.
func (au *ArticleUpdate) SetNillableRecommend(s *string) *ArticleUpdate {
	if s != nil {
		au.SetRecommend(*s)
	}
	return au
}

// ClearRecommend clears the value of the "recommend" field.
func (au *ArticleUpdate) ClearRecommend() *ArticleUpdate {
	au.mutation.ClearRecommend()
	return au
}

// SetDescription sets the "description" field.
func (au *ArticleUpdate) SetDescription(s string) *ArticleUpdate {
	au.mutation.SetDescription(s)
	return au
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (au *ArticleUpdate) SetNillableDescription(s *string) *ArticleUpdate {
	if s != nil {
		au.SetDescription(*s)
	}
	return au
}

// ClearDescription clears the value of the "description" field.
func (au *ArticleUpdate) ClearDescription() *ArticleUpdate {
	au.mutation.ClearDescription()
	return au
}

// SetContentBody sets the "content_body" field.
func (au *ArticleUpdate) SetContentBody(s string) *ArticleUpdate {
	au.mutation.SetContentBody(s)
	return au
}

// SetNillableContentBody sets the "content_body" field if the given value is not nil.
func (au *ArticleUpdate) SetNillableContentBody(s *string) *ArticleUpdate {
	if s != nil {
		au.SetContentBody(*s)
	}
	return au
}

// ClearContentBody clears the value of the "content_body" field.
func (au *ArticleUpdate) ClearContentBody() *ArticleUpdate {
	au.mutation.ClearContentBody()
	return au
}

// SetImageURL sets the "image_url" field.
func (au *ArticleUpdate) SetImageURL(s string) *ArticleUpdate {
	au.mutation.SetImageURL(s)
	return au
}

// SetNillableImageURL sets the "image_url" field if the given value is not nil.
func (au *ArticleUpdate) SetNillableImageURL(s *string) *ArticleUpdate {
	if s != nil {
		au.SetImageURL(*s)
	}
	return au
}

// ClearImageURL clears the value of the "image_url" field.
func (au *ArticleUpdate) ClearImageURL() *ArticleUpdate {
	au.mutation.ClearImageURL()
	return au
}

// SetStatus sets the "status" field.
func (au *ArticleUpdate) SetStatus(i int32) *ArticleUpdate {
	au.mutation.ResetStatus()
	au.mutation.SetStatus(i)
	return au
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (au *ArticleUpdate) SetNillableStatus(i *int32) *ArticleUpdate {
	if i != nil {
		au.SetStatus(*i)
	}
	return au
}

// AddStatus adds i to the "status" field.
func (au *ArticleUpdate) AddStatus(i int32) *ArticleUpdate {
	au.mutation.AddStatus(i)
	return au
}

// ClearStatus clears the value of the "status" field.
func (au *ArticleUpdate) ClearStatus() *ArticleUpdate {
	au.mutation.ClearStatus()
	return au
}

// SetClickCount sets the "click_count" field.
func (au *ArticleUpdate) SetClickCount(i int32) *ArticleUpdate {
	au.mutation.ResetClickCount()
	au.mutation.SetClickCount(i)
	return au
}

// SetNillableClickCount sets the "click_count" field if the given value is not nil.
func (au *ArticleUpdate) SetNillableClickCount(i *int32) *ArticleUpdate {
	if i != nil {
		au.SetClickCount(*i)
	}
	return au
}

// AddClickCount adds i to the "click_count" field.
func (au *ArticleUpdate) AddClickCount(i int32) *ArticleUpdate {
	au.mutation.AddClickCount(i)
	return au
}

// ClearClickCount clears the value of the "click_count" field.
func (au *ArticleUpdate) ClearClickCount() *ArticleUpdate {
	au.mutation.ClearClickCount()
	return au
}

// SetLikeCount sets the "like_count" field.
func (au *ArticleUpdate) SetLikeCount(i int32) *ArticleUpdate {
	au.mutation.ResetLikeCount()
	au.mutation.SetLikeCount(i)
	return au
}

// SetNillableLikeCount sets the "like_count" field if the given value is not nil.
func (au *ArticleUpdate) SetNillableLikeCount(i *int32) *ArticleUpdate {
	if i != nil {
		au.SetLikeCount(*i)
	}
	return au
}

// AddLikeCount adds i to the "like_count" field.
func (au *ArticleUpdate) AddLikeCount(i int32) *ArticleUpdate {
	au.mutation.AddLikeCount(i)
	return au
}

// ClearLikeCount clears the value of the "like_count" field.
func (au *ArticleUpdate) ClearLikeCount() *ArticleUpdate {
	au.mutation.ClearLikeCount()
	return au
}

// SetCreateID sets the "create_id" field.
func (au *ArticleUpdate) SetCreateID(i int64) *ArticleUpdate {
	au.mutation.ResetCreateID()
	au.mutation.SetCreateID(i)
	return au
}

// SetNillableCreateID sets the "create_id" field if the given value is not nil.
func (au *ArticleUpdate) SetNillableCreateID(i *int64) *ArticleUpdate {
	if i != nil {
		au.SetCreateID(*i)
	}
	return au
}

// AddCreateID adds i to the "create_id" field.
func (au *ArticleUpdate) AddCreateID(i int64) *ArticleUpdate {
	au.mutation.AddCreateID(i)
	return au
}

// ClearCreateID clears the value of the "create_id" field.
func (au *ArticleUpdate) ClearCreateID() *ArticleUpdate {
	au.mutation.ClearCreateID()
	return au
}

// SetCreateTime sets the "create_time" field.
func (au *ArticleUpdate) SetCreateTime(t time.Time) *ArticleUpdate {
	au.mutation.SetCreateTime(t)
	return au
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (au *ArticleUpdate) SetNillableCreateTime(t *time.Time) *ArticleUpdate {
	if t != nil {
		au.SetCreateTime(*t)
	}
	return au
}

// ClearCreateTime clears the value of the "create_time" field.
func (au *ArticleUpdate) ClearCreateTime() *ArticleUpdate {
	au.mutation.ClearCreateTime()
	return au
}

// SetModifyID sets the "modify_id" field.
func (au *ArticleUpdate) SetModifyID(i int64) *ArticleUpdate {
	au.mutation.ResetModifyID()
	au.mutation.SetModifyID(i)
	return au
}

// SetNillableModifyID sets the "modify_id" field if the given value is not nil.
func (au *ArticleUpdate) SetNillableModifyID(i *int64) *ArticleUpdate {
	if i != nil {
		au.SetModifyID(*i)
	}
	return au
}

// AddModifyID adds i to the "modify_id" field.
func (au *ArticleUpdate) AddModifyID(i int64) *ArticleUpdate {
	au.mutation.AddModifyID(i)
	return au
}

// ClearModifyID clears the value of the "modify_id" field.
func (au *ArticleUpdate) ClearModifyID() *ArticleUpdate {
	au.mutation.ClearModifyID()
	return au
}

// SetModifyTime sets the "modify_time" field.
func (au *ArticleUpdate) SetModifyTime(t time.Time) *ArticleUpdate {
	au.mutation.SetModifyTime(t)
	return au
}

// ClearModifyTime clears the value of the "modify_time" field.
func (au *ArticleUpdate) ClearModifyTime() *ArticleUpdate {
	au.mutation.ClearModifyTime()
	return au
}

// Mutation returns the ArticleMutation object of the builder.
func (au *ArticleUpdate) Mutation() *ArticleMutation {
	return au.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (au *ArticleUpdate) Save(ctx context.Context) (int, error) {
	au.defaults()
	return withHooks[int, ArticleMutation](ctx, au.sqlSave, au.mutation, au.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (au *ArticleUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *ArticleUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *ArticleUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (au *ArticleUpdate) defaults() {
	if _, ok := au.mutation.ModifyTime(); !ok && !au.mutation.ModifyTimeCleared() {
		v := article.UpdateDefaultModifyTime()
		au.mutation.SetModifyTime(v)
	}
}

func (au *ArticleUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(article.Table, article.Columns, sqlgraph.NewFieldSpec(article.FieldID, field.TypeInt32))
	if ps := au.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := au.mutation.SortID(); ok {
		_spec.SetField(article.FieldSortID, field.TypeInt32, value)
	}
	if value, ok := au.mutation.AddedSortID(); ok {
		_spec.AddField(article.FieldSortID, field.TypeInt32, value)
	}
	if au.mutation.SortIDCleared() {
		_spec.ClearField(article.FieldSortID, field.TypeInt32)
	}
	if value, ok := au.mutation.Name(); ok {
		_spec.SetField(article.FieldName, field.TypeString, value)
	}
	if au.mutation.NameCleared() {
		_spec.ClearField(article.FieldName, field.TypeString)
	}
	if value, ok := au.mutation.CategoryID(); ok {
		_spec.SetField(article.FieldCategoryID, field.TypeInt32, value)
	}
	if value, ok := au.mutation.AddedCategoryID(); ok {
		_spec.AddField(article.FieldCategoryID, field.TypeInt32, value)
	}
	if au.mutation.CategoryIDCleared() {
		_spec.ClearField(article.FieldCategoryID, field.TypeInt32)
	}
	if value, ok := au.mutation.Recommend(); ok {
		_spec.SetField(article.FieldRecommend, field.TypeString, value)
	}
	if au.mutation.RecommendCleared() {
		_spec.ClearField(article.FieldRecommend, field.TypeString)
	}
	if value, ok := au.mutation.Description(); ok {
		_spec.SetField(article.FieldDescription, field.TypeString, value)
	}
	if au.mutation.DescriptionCleared() {
		_spec.ClearField(article.FieldDescription, field.TypeString)
	}
	if value, ok := au.mutation.ContentBody(); ok {
		_spec.SetField(article.FieldContentBody, field.TypeString, value)
	}
	if au.mutation.ContentBodyCleared() {
		_spec.ClearField(article.FieldContentBody, field.TypeString)
	}
	if value, ok := au.mutation.ImageURL(); ok {
		_spec.SetField(article.FieldImageURL, field.TypeString, value)
	}
	if au.mutation.ImageURLCleared() {
		_spec.ClearField(article.FieldImageURL, field.TypeString)
	}
	if value, ok := au.mutation.Status(); ok {
		_spec.SetField(article.FieldStatus, field.TypeInt32, value)
	}
	if value, ok := au.mutation.AddedStatus(); ok {
		_spec.AddField(article.FieldStatus, field.TypeInt32, value)
	}
	if au.mutation.StatusCleared() {
		_spec.ClearField(article.FieldStatus, field.TypeInt32)
	}
	if value, ok := au.mutation.ClickCount(); ok {
		_spec.SetField(article.FieldClickCount, field.TypeInt32, value)
	}
	if value, ok := au.mutation.AddedClickCount(); ok {
		_spec.AddField(article.FieldClickCount, field.TypeInt32, value)
	}
	if au.mutation.ClickCountCleared() {
		_spec.ClearField(article.FieldClickCount, field.TypeInt32)
	}
	if value, ok := au.mutation.LikeCount(); ok {
		_spec.SetField(article.FieldLikeCount, field.TypeInt32, value)
	}
	if value, ok := au.mutation.AddedLikeCount(); ok {
		_spec.AddField(article.FieldLikeCount, field.TypeInt32, value)
	}
	if au.mutation.LikeCountCleared() {
		_spec.ClearField(article.FieldLikeCount, field.TypeInt32)
	}
	if value, ok := au.mutation.CreateID(); ok {
		_spec.SetField(article.FieldCreateID, field.TypeInt64, value)
	}
	if value, ok := au.mutation.AddedCreateID(); ok {
		_spec.AddField(article.FieldCreateID, field.TypeInt64, value)
	}
	if au.mutation.CreateIDCleared() {
		_spec.ClearField(article.FieldCreateID, field.TypeInt64)
	}
	if value, ok := au.mutation.CreateTime(); ok {
		_spec.SetField(article.FieldCreateTime, field.TypeTime, value)
	}
	if au.mutation.CreateTimeCleared() {
		_spec.ClearField(article.FieldCreateTime, field.TypeTime)
	}
	if value, ok := au.mutation.ModifyID(); ok {
		_spec.SetField(article.FieldModifyID, field.TypeInt64, value)
	}
	if value, ok := au.mutation.AddedModifyID(); ok {
		_spec.AddField(article.FieldModifyID, field.TypeInt64, value)
	}
	if au.mutation.ModifyIDCleared() {
		_spec.ClearField(article.FieldModifyID, field.TypeInt64)
	}
	if value, ok := au.mutation.ModifyTime(); ok {
		_spec.SetField(article.FieldModifyTime, field.TypeTime, value)
	}
	if au.mutation.ModifyTimeCleared() {
		_spec.ClearField(article.FieldModifyTime, field.TypeTime)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, au.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{article.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	au.mutation.done = true
	return n, nil
}

// ArticleUpdateOne is the builder for updating a single Article entity.
type ArticleUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ArticleMutation
}

// SetSortID sets the "sort_id" field.
func (auo *ArticleUpdateOne) SetSortID(i int32) *ArticleUpdateOne {
	auo.mutation.ResetSortID()
	auo.mutation.SetSortID(i)
	return auo
}

// SetNillableSortID sets the "sort_id" field if the given value is not nil.
func (auo *ArticleUpdateOne) SetNillableSortID(i *int32) *ArticleUpdateOne {
	if i != nil {
		auo.SetSortID(*i)
	}
	return auo
}

// AddSortID adds i to the "sort_id" field.
func (auo *ArticleUpdateOne) AddSortID(i int32) *ArticleUpdateOne {
	auo.mutation.AddSortID(i)
	return auo
}

// ClearSortID clears the value of the "sort_id" field.
func (auo *ArticleUpdateOne) ClearSortID() *ArticleUpdateOne {
	auo.mutation.ClearSortID()
	return auo
}

// SetName sets the "name" field.
func (auo *ArticleUpdateOne) SetName(s string) *ArticleUpdateOne {
	auo.mutation.SetName(s)
	return auo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (auo *ArticleUpdateOne) SetNillableName(s *string) *ArticleUpdateOne {
	if s != nil {
		auo.SetName(*s)
	}
	return auo
}

// ClearName clears the value of the "name" field.
func (auo *ArticleUpdateOne) ClearName() *ArticleUpdateOne {
	auo.mutation.ClearName()
	return auo
}

// SetCategoryID sets the "category_id" field.
func (auo *ArticleUpdateOne) SetCategoryID(i int32) *ArticleUpdateOne {
	auo.mutation.ResetCategoryID()
	auo.mutation.SetCategoryID(i)
	return auo
}

// SetNillableCategoryID sets the "category_id" field if the given value is not nil.
func (auo *ArticleUpdateOne) SetNillableCategoryID(i *int32) *ArticleUpdateOne {
	if i != nil {
		auo.SetCategoryID(*i)
	}
	return auo
}

// AddCategoryID adds i to the "category_id" field.
func (auo *ArticleUpdateOne) AddCategoryID(i int32) *ArticleUpdateOne {
	auo.mutation.AddCategoryID(i)
	return auo
}

// ClearCategoryID clears the value of the "category_id" field.
func (auo *ArticleUpdateOne) ClearCategoryID() *ArticleUpdateOne {
	auo.mutation.ClearCategoryID()
	return auo
}

// SetRecommend sets the "recommend" field.
func (auo *ArticleUpdateOne) SetRecommend(s string) *ArticleUpdateOne {
	auo.mutation.SetRecommend(s)
	return auo
}

// SetNillableRecommend sets the "recommend" field if the given value is not nil.
func (auo *ArticleUpdateOne) SetNillableRecommend(s *string) *ArticleUpdateOne {
	if s != nil {
		auo.SetRecommend(*s)
	}
	return auo
}

// ClearRecommend clears the value of the "recommend" field.
func (auo *ArticleUpdateOne) ClearRecommend() *ArticleUpdateOne {
	auo.mutation.ClearRecommend()
	return auo
}

// SetDescription sets the "description" field.
func (auo *ArticleUpdateOne) SetDescription(s string) *ArticleUpdateOne {
	auo.mutation.SetDescription(s)
	return auo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (auo *ArticleUpdateOne) SetNillableDescription(s *string) *ArticleUpdateOne {
	if s != nil {
		auo.SetDescription(*s)
	}
	return auo
}

// ClearDescription clears the value of the "description" field.
func (auo *ArticleUpdateOne) ClearDescription() *ArticleUpdateOne {
	auo.mutation.ClearDescription()
	return auo
}

// SetContentBody sets the "content_body" field.
func (auo *ArticleUpdateOne) SetContentBody(s string) *ArticleUpdateOne {
	auo.mutation.SetContentBody(s)
	return auo
}

// SetNillableContentBody sets the "content_body" field if the given value is not nil.
func (auo *ArticleUpdateOne) SetNillableContentBody(s *string) *ArticleUpdateOne {
	if s != nil {
		auo.SetContentBody(*s)
	}
	return auo
}

// ClearContentBody clears the value of the "content_body" field.
func (auo *ArticleUpdateOne) ClearContentBody() *ArticleUpdateOne {
	auo.mutation.ClearContentBody()
	return auo
}

// SetImageURL sets the "image_url" field.
func (auo *ArticleUpdateOne) SetImageURL(s string) *ArticleUpdateOne {
	auo.mutation.SetImageURL(s)
	return auo
}

// SetNillableImageURL sets the "image_url" field if the given value is not nil.
func (auo *ArticleUpdateOne) SetNillableImageURL(s *string) *ArticleUpdateOne {
	if s != nil {
		auo.SetImageURL(*s)
	}
	return auo
}

// ClearImageURL clears the value of the "image_url" field.
func (auo *ArticleUpdateOne) ClearImageURL() *ArticleUpdateOne {
	auo.mutation.ClearImageURL()
	return auo
}

// SetStatus sets the "status" field.
func (auo *ArticleUpdateOne) SetStatus(i int32) *ArticleUpdateOne {
	auo.mutation.ResetStatus()
	auo.mutation.SetStatus(i)
	return auo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (auo *ArticleUpdateOne) SetNillableStatus(i *int32) *ArticleUpdateOne {
	if i != nil {
		auo.SetStatus(*i)
	}
	return auo
}

// AddStatus adds i to the "status" field.
func (auo *ArticleUpdateOne) AddStatus(i int32) *ArticleUpdateOne {
	auo.mutation.AddStatus(i)
	return auo
}

// ClearStatus clears the value of the "status" field.
func (auo *ArticleUpdateOne) ClearStatus() *ArticleUpdateOne {
	auo.mutation.ClearStatus()
	return auo
}

// SetClickCount sets the "click_count" field.
func (auo *ArticleUpdateOne) SetClickCount(i int32) *ArticleUpdateOne {
	auo.mutation.ResetClickCount()
	auo.mutation.SetClickCount(i)
	return auo
}

// SetNillableClickCount sets the "click_count" field if the given value is not nil.
func (auo *ArticleUpdateOne) SetNillableClickCount(i *int32) *ArticleUpdateOne {
	if i != nil {
		auo.SetClickCount(*i)
	}
	return auo
}

// AddClickCount adds i to the "click_count" field.
func (auo *ArticleUpdateOne) AddClickCount(i int32) *ArticleUpdateOne {
	auo.mutation.AddClickCount(i)
	return auo
}

// ClearClickCount clears the value of the "click_count" field.
func (auo *ArticleUpdateOne) ClearClickCount() *ArticleUpdateOne {
	auo.mutation.ClearClickCount()
	return auo
}

// SetLikeCount sets the "like_count" field.
func (auo *ArticleUpdateOne) SetLikeCount(i int32) *ArticleUpdateOne {
	auo.mutation.ResetLikeCount()
	auo.mutation.SetLikeCount(i)
	return auo
}

// SetNillableLikeCount sets the "like_count" field if the given value is not nil.
func (auo *ArticleUpdateOne) SetNillableLikeCount(i *int32) *ArticleUpdateOne {
	if i != nil {
		auo.SetLikeCount(*i)
	}
	return auo
}

// AddLikeCount adds i to the "like_count" field.
func (auo *ArticleUpdateOne) AddLikeCount(i int32) *ArticleUpdateOne {
	auo.mutation.AddLikeCount(i)
	return auo
}

// ClearLikeCount clears the value of the "like_count" field.
func (auo *ArticleUpdateOne) ClearLikeCount() *ArticleUpdateOne {
	auo.mutation.ClearLikeCount()
	return auo
}

// SetCreateID sets the "create_id" field.
func (auo *ArticleUpdateOne) SetCreateID(i int64) *ArticleUpdateOne {
	auo.mutation.ResetCreateID()
	auo.mutation.SetCreateID(i)
	return auo
}

// SetNillableCreateID sets the "create_id" field if the given value is not nil.
func (auo *ArticleUpdateOne) SetNillableCreateID(i *int64) *ArticleUpdateOne {
	if i != nil {
		auo.SetCreateID(*i)
	}
	return auo
}

// AddCreateID adds i to the "create_id" field.
func (auo *ArticleUpdateOne) AddCreateID(i int64) *ArticleUpdateOne {
	auo.mutation.AddCreateID(i)
	return auo
}

// ClearCreateID clears the value of the "create_id" field.
func (auo *ArticleUpdateOne) ClearCreateID() *ArticleUpdateOne {
	auo.mutation.ClearCreateID()
	return auo
}

// SetCreateTime sets the "create_time" field.
func (auo *ArticleUpdateOne) SetCreateTime(t time.Time) *ArticleUpdateOne {
	auo.mutation.SetCreateTime(t)
	return auo
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (auo *ArticleUpdateOne) SetNillableCreateTime(t *time.Time) *ArticleUpdateOne {
	if t != nil {
		auo.SetCreateTime(*t)
	}
	return auo
}

// ClearCreateTime clears the value of the "create_time" field.
func (auo *ArticleUpdateOne) ClearCreateTime() *ArticleUpdateOne {
	auo.mutation.ClearCreateTime()
	return auo
}

// SetModifyID sets the "modify_id" field.
func (auo *ArticleUpdateOne) SetModifyID(i int64) *ArticleUpdateOne {
	auo.mutation.ResetModifyID()
	auo.mutation.SetModifyID(i)
	return auo
}

// SetNillableModifyID sets the "modify_id" field if the given value is not nil.
func (auo *ArticleUpdateOne) SetNillableModifyID(i *int64) *ArticleUpdateOne {
	if i != nil {
		auo.SetModifyID(*i)
	}
	return auo
}

// AddModifyID adds i to the "modify_id" field.
func (auo *ArticleUpdateOne) AddModifyID(i int64) *ArticleUpdateOne {
	auo.mutation.AddModifyID(i)
	return auo
}

// ClearModifyID clears the value of the "modify_id" field.
func (auo *ArticleUpdateOne) ClearModifyID() *ArticleUpdateOne {
	auo.mutation.ClearModifyID()
	return auo
}

// SetModifyTime sets the "modify_time" field.
func (auo *ArticleUpdateOne) SetModifyTime(t time.Time) *ArticleUpdateOne {
	auo.mutation.SetModifyTime(t)
	return auo
}

// ClearModifyTime clears the value of the "modify_time" field.
func (auo *ArticleUpdateOne) ClearModifyTime() *ArticleUpdateOne {
	auo.mutation.ClearModifyTime()
	return auo
}

// Mutation returns the ArticleMutation object of the builder.
func (auo *ArticleUpdateOne) Mutation() *ArticleMutation {
	return auo.mutation
}

// Where appends a list predicates to the ArticleUpdate builder.
func (auo *ArticleUpdateOne) Where(ps ...predicate.Article) *ArticleUpdateOne {
	auo.mutation.Where(ps...)
	return auo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (auo *ArticleUpdateOne) Select(field string, fields ...string) *ArticleUpdateOne {
	auo.fields = append([]string{field}, fields...)
	return auo
}

// Save executes the query and returns the updated Article entity.
func (auo *ArticleUpdateOne) Save(ctx context.Context) (*Article, error) {
	auo.defaults()
	return withHooks[*Article, ArticleMutation](ctx, auo.sqlSave, auo.mutation, auo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (auo *ArticleUpdateOne) SaveX(ctx context.Context) *Article {
	node, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (auo *ArticleUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *ArticleUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (auo *ArticleUpdateOne) defaults() {
	if _, ok := auo.mutation.ModifyTime(); !ok && !auo.mutation.ModifyTimeCleared() {
		v := article.UpdateDefaultModifyTime()
		auo.mutation.SetModifyTime(v)
	}
}

func (auo *ArticleUpdateOne) sqlSave(ctx context.Context) (_node *Article, err error) {
	_spec := sqlgraph.NewUpdateSpec(article.Table, article.Columns, sqlgraph.NewFieldSpec(article.FieldID, field.TypeInt32))
	id, ok := auo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Article.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := auo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, article.FieldID)
		for _, f := range fields {
			if !article.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != article.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := auo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := auo.mutation.SortID(); ok {
		_spec.SetField(article.FieldSortID, field.TypeInt32, value)
	}
	if value, ok := auo.mutation.AddedSortID(); ok {
		_spec.AddField(article.FieldSortID, field.TypeInt32, value)
	}
	if auo.mutation.SortIDCleared() {
		_spec.ClearField(article.FieldSortID, field.TypeInt32)
	}
	if value, ok := auo.mutation.Name(); ok {
		_spec.SetField(article.FieldName, field.TypeString, value)
	}
	if auo.mutation.NameCleared() {
		_spec.ClearField(article.FieldName, field.TypeString)
	}
	if value, ok := auo.mutation.CategoryID(); ok {
		_spec.SetField(article.FieldCategoryID, field.TypeInt32, value)
	}
	if value, ok := auo.mutation.AddedCategoryID(); ok {
		_spec.AddField(article.FieldCategoryID, field.TypeInt32, value)
	}
	if auo.mutation.CategoryIDCleared() {
		_spec.ClearField(article.FieldCategoryID, field.TypeInt32)
	}
	if value, ok := auo.mutation.Recommend(); ok {
		_spec.SetField(article.FieldRecommend, field.TypeString, value)
	}
	if auo.mutation.RecommendCleared() {
		_spec.ClearField(article.FieldRecommend, field.TypeString)
	}
	if value, ok := auo.mutation.Description(); ok {
		_spec.SetField(article.FieldDescription, field.TypeString, value)
	}
	if auo.mutation.DescriptionCleared() {
		_spec.ClearField(article.FieldDescription, field.TypeString)
	}
	if value, ok := auo.mutation.ContentBody(); ok {
		_spec.SetField(article.FieldContentBody, field.TypeString, value)
	}
	if auo.mutation.ContentBodyCleared() {
		_spec.ClearField(article.FieldContentBody, field.TypeString)
	}
	if value, ok := auo.mutation.ImageURL(); ok {
		_spec.SetField(article.FieldImageURL, field.TypeString, value)
	}
	if auo.mutation.ImageURLCleared() {
		_spec.ClearField(article.FieldImageURL, field.TypeString)
	}
	if value, ok := auo.mutation.Status(); ok {
		_spec.SetField(article.FieldStatus, field.TypeInt32, value)
	}
	if value, ok := auo.mutation.AddedStatus(); ok {
		_spec.AddField(article.FieldStatus, field.TypeInt32, value)
	}
	if auo.mutation.StatusCleared() {
		_spec.ClearField(article.FieldStatus, field.TypeInt32)
	}
	if value, ok := auo.mutation.ClickCount(); ok {
		_spec.SetField(article.FieldClickCount, field.TypeInt32, value)
	}
	if value, ok := auo.mutation.AddedClickCount(); ok {
		_spec.AddField(article.FieldClickCount, field.TypeInt32, value)
	}
	if auo.mutation.ClickCountCleared() {
		_spec.ClearField(article.FieldClickCount, field.TypeInt32)
	}
	if value, ok := auo.mutation.LikeCount(); ok {
		_spec.SetField(article.FieldLikeCount, field.TypeInt32, value)
	}
	if value, ok := auo.mutation.AddedLikeCount(); ok {
		_spec.AddField(article.FieldLikeCount, field.TypeInt32, value)
	}
	if auo.mutation.LikeCountCleared() {
		_spec.ClearField(article.FieldLikeCount, field.TypeInt32)
	}
	if value, ok := auo.mutation.CreateID(); ok {
		_spec.SetField(article.FieldCreateID, field.TypeInt64, value)
	}
	if value, ok := auo.mutation.AddedCreateID(); ok {
		_spec.AddField(article.FieldCreateID, field.TypeInt64, value)
	}
	if auo.mutation.CreateIDCleared() {
		_spec.ClearField(article.FieldCreateID, field.TypeInt64)
	}
	if value, ok := auo.mutation.CreateTime(); ok {
		_spec.SetField(article.FieldCreateTime, field.TypeTime, value)
	}
	if auo.mutation.CreateTimeCleared() {
		_spec.ClearField(article.FieldCreateTime, field.TypeTime)
	}
	if value, ok := auo.mutation.ModifyID(); ok {
		_spec.SetField(article.FieldModifyID, field.TypeInt64, value)
	}
	if value, ok := auo.mutation.AddedModifyID(); ok {
		_spec.AddField(article.FieldModifyID, field.TypeInt64, value)
	}
	if auo.mutation.ModifyIDCleared() {
		_spec.ClearField(article.FieldModifyID, field.TypeInt64)
	}
	if value, ok := auo.mutation.ModifyTime(); ok {
		_spec.SetField(article.FieldModifyTime, field.TypeTime, value)
	}
	if auo.mutation.ModifyTimeCleared() {
		_spec.ClearField(article.FieldModifyTime, field.TypeTime)
	}
	_node = &Article{config: auo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{article.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	auo.mutation.done = true
	return _node, nil
}

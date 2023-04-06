package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"go-server/internal/biz"
	"go-server/internal/data/ent/category"
)

type CategoryRepo struct {
	data *Data
	log  *log.Helper
}

func (c *CategoryRepo) CategoryList(ctx context.Context, catePara int32) ([]*biz.Category, error) {
	//TODO implement me
	cc, _ := c.data.db.Category.Query().
		Where(
			category.ParentIDEQ(int64(catePara)),
		).All(ctx)
	rv := make([]*biz.Category, 0)
	for _, value := range cc {
		rv = append(rv, &biz.Category{
			ID:           value.ID,
			ParentId:     value.ParentID,
			CategoryId:   value.CategoryID,
			CategoryName: value.CategoryName,
			Status:       value.Status,
			CreateId:     value.CreateID,
			CreateTime:   value.CreateTime,
			ModifyId:     value.ModifyID,
			ModifyTime:   value.ModifyTime,
		})
	}
	return rv, nil
}

func NewCategoryRepo(data *Data, logger log.Logger) biz.CategoryRepo {
	return &CategoryRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

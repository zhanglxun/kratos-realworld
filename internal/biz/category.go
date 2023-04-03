package biz

import (
	"context"
	"time"
)

type Category struct {
	ID           int64
	ParentId     int64
	CategoryId   int32
	CategoryName string
	Status       int32
	CreateId     int64
	CreateTime   time.Time
	ModifyId     int64
	ModifyTime   time.Time
}

type CategoryRepo interface {
	CategoryList(ctx context.Context, category int64) ([]*Category, error)
}

//内容数据
func (ws *ContentUsecase) CategoryList(ctx context.Context, categoryPara int64) (wr []*Category, err error) {
	wr, err = ws.repoCa.CategoryList(ctx, categoryPara)
	return wr, err
}

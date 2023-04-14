package biz

import (
	"context"
	"time"
)

//分类
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
	CategoryList(ctx context.Context, category int32) ([]*Category, error)
}

// CategoryList 依赖注入
func (ws *ContentUsecase) CategoryList(ctx context.Context, categoryPara int32) (wr []*Category, err error) {
	wr, err = ws.repoCa.CategoryList(ctx, categoryPara)
	return wr, err
}

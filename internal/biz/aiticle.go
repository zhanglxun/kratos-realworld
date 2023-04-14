package biz

import (
	"context"
	"github.com/golang/protobuf/ptypes/timestamp"
)

//article struct
type Article struct {
	ID          int64
	SortId      int32
	Name        string
	CategoryId  int32
	Recommend   string
	Description string
	ContentBody string
	ImageUrl    string
	Status      int32
	ClickCount  int32
	LikeCount   int32
	CreateId    int64
	CreateTime  timestamp.Timestamp
	ModifyId    int64
	ModifyTime  timestamp.Timestamp
}

type ArticleRepo interface {
	ArticleList(ctx context.Context, CurrentPage int32, PageNumber int32) ([]*Article, error)
}

// CategoryList 依赖注入
func (ws *ContentUsecase) ArticleList(ctx context.Context, CurrentPage int32, PageNumber int32) (wr []*Article, err error) {
	wr, err = ws.repoAr.ArticleList(ctx, CurrentPage, PageNumber)
	return wr, err
}

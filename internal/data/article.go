package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"go-server/internal/biz"
)

type ArticleRepo struct {
	data *Data
	log  *log.Helper
}

func (a *ArticleRepo) ArticleList(ctx context.Context, CurrentPage int32, PageNumber int32) ([]*biz.Article, error) {
	//TODO implement me
	println("article list data")
	ar, _ := a.data.db.Article.Query().
		Where().All(ctx)
	rv := make([]*biz.Article, 0)
	for _, value := range ar {
		rv = append(rv, &biz.Article{
			ID:          value.ID,
			SortId:      value.SortID,
			Name:        value.Name,
			CategoryId:  value.CategoryID,
			Recommend:   value.Recommend,
			Description: value.Description,
			ContentBody: value.ContentBody,
			ImageUrl:    value.ImageURL,
			Status:      value.Status,
			ClickCount:  value.ClickCount,
			LikeCount:   value.LikeCount,
			CreateId:    value.CreateID,
			//CreateTime:  value.CreateTime,
			ModifyId: value.ModifyID,
			//ModifyTime:  value.ModifyTime,
		})
	}
	return rv, nil
}

func NewArticleRepo(data *Data, logger log.Logger) biz.ArticleRepo {
	return &ArticleRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

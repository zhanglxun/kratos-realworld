package service

import (
	"context"
	v1 "go-server/api/realworld/v1"
)

// ArticleList.
func (r *ContentService) ArticleList(ctx context.Context, req *v1.ArticleRequest) (ar *v1.ListArticleReply, err error) {
	ArList, _ := r.ws.ArticleList(ctx, req.CurrentPage, req.PageNumber)
	reply := &v1.ListArticleReply{}
	for _, value := range ArList {

		reply.Article = append(reply.Article, &v1.Article{
			Id:          value.ID,
			SortId:      value.SortId,
			Name:        value.Name,
			CategoryId:  value.CategoryId,
			Recommend:   value.Recommend,
			Description: value.Description,
			ContentBody: value.ContentBody,
			ImageUrl:    value.ImageUrl,
			Status:      value.Status,
			ClickCount:  value.ClickCount,
			LikeCount:   value.LikeCount,
			CreateId:    value.CreateId,
			//CreateTime:  value.CreateTime,
			ModifyId: value.ModifyId,
			//ModifyTime:  value.ModifyTime,
		})
	}
	return reply, nil
}

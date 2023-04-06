package service

import (
	"context"
	v1 "go-server/api/realworld/v1"
)

func (s *ContentService) CategoryList(ctx context.Context, req *v1.CategoryRequest) (ws *v1.MultipleCategoryReply, err error) {
	categoryList, _ := s.ws.CategoryList(ctx, req.CatePare)
	reply := &v1.MultipleCategoryReply{}
	for _, p := range categoryList {

		reply.Category = append(reply.Category, &v1.MultipleCategoryReply_Categories{
			CategoryId:   p.CategoryId,
			CategoryName: p.CategoryName,
		})
	}
	return reply, nil
}

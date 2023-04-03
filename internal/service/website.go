package service

import (
	"context"
	v1 "go-server/api/realworld/v1"
)

func (s *RealWorldService) WebSiteList(ctx context.Context, req *v1.WebsiteRequest) (ws *v1.WebsiteReply, err error) {

	websiteLit, _ := s.ws.WebSiteList(ctx, req.Category, req.Type)
	reply := &v1.WebsiteReply{}
	for _, p := range websiteLit {

		reply.Website = append(reply.Website, &v1.WebsiteReply_WebSite{
			Category:    p.Category,
			WebsiteIcon: p.WebsiteIcon,
			WebsiteUrl:  p.WebsiteUrl,
			WebsiteName: p.WebsiteName,
			Summary:     p.Summary,
		})
	}
	return reply, err
}

func (s *RealWorldService) WebSiteHome(ctx context.Context, req *v1.WebHomeRequest) (ws *v1.WebSiteHomeReply, err error) {
	rh, err := s.ws.WebSiteHome(ctx)
	reply := &v1.WebSiteHomeReply{}
	for _, p := range rh {
		reply.WebsiteReply = rh
		println(p)
	}
	return reply, err
}

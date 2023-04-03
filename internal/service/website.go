package service

import (
	"context"
	v1 "go-server/api/realworld/v1"
)

type Website struct {
	Category    int32
	WebsiteName string
	WebsiteIcon string
	WebsiteUrl  string
	Summary     string
}

func (s *RealWorldService) WebSiteList(ctx context.Context, req *v1.WebsiteRequest) (ws *v1.WebsiteReply, err error) {

	websiteLit, _ := s.ws.WebSiteList(ctx, req.Type)
	reply := &v1.WebsiteReply{}
	for _, p := range websiteLit {
		reply.Website = append(reply.Website, &v1.WebSite{
			Category:    p.Category,
			WebsiteIcon: p.WebsiteIcon,
			WebsiteUrl:  p.WebsiteUrl,
			WebsiteName: p.WebsiteName,
			Summary:     p.Summary,
		})
	}
	return reply, err
}

func (s *RealWorldService) WebSiteHome(ctx context.Context, req *v1.WebHomeRequest) (mws *v1.MultipleWebsiteReply, err error) {
	rh, err := s.ws.WebSiteHome(ctx)
	println(rh)
	ca, err := s.ws.CategoryList(ctx, req.CatePare)
	ws, _ := s.ws.WebSiteList(ctx, 1)
	reply := &v1.MultipleWebsiteReply{}

	for _, c := range ca {

		reply2 := &v1.WebsiteReply{}
		for _, w := range ws {
			if c.CategoryId == w.Category {
				reply2.Website = append(reply2.Website, &v1.WebSite{
					Category:    w.Category,
					WebsiteIcon: w.WebsiteIcon,
					WebsiteUrl:  w.WebsiteUrl,
					WebsiteName: w.WebsiteName,
					Summary:     w.Summary,
				})
			}
		}
		reply.WebsiteReply = append(reply.WebsiteReply, &v1.MultipleWebsiteReply_WebsiteReply{
			CategoryName: c.CategoryName,
			Website:      reply2.Website,
		})
	}
	return reply, nil
}

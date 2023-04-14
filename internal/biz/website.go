package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	v1 "go-server/api/realworld/v1"
	"time"
)

type Website struct {
	ID          int64
	SortId      int32
	Category    int32
	Type        int32
	WebsiteName string
	WebsiteIcon string
	WebsiteUrl  string
	Summary     string
	description string
	CreateId    int64
	CreateTime  time.Time
	ModifyId    int64
	ModifyTime  time.Time
}

type CategoryHome struct {
	CategoryName string
	Website      Website
}

// WebsiteRepo .
type WebsiteRepo interface {
	Save()
	Update()
	Delete()
	WebSiteList(context.Context, int32) ([]*Website, error)
	WebSiteHome(ctx context.Context) ([]*v1.WebsiteReply, error)
}

func (ws *ContentUsecase) WebSiteList(ctx context.Context, typ int32) (wr []*Website, err error) {
	wr, err = ws.repoWs.WebSiteList(ctx, typ)
	return wr, err
}

func (ws *ContentUsecase) WebSiteHome(ctx context.Context) (webHome []*CategoryHome, err error) {

	wc, err := ws.repoCa.CategoryList(ctx, 0)
	wh, err := ws.repoWs.WebSiteList(ctx, 1)
	reply := make([]*CategoryHome, 0)
	for _, c := range wc {

		reply = append(reply, &CategoryHome{
			CategoryName: c.CategoryName,
		})
		for _, w := range wh {
			if c.CategoryId == w.Category {

				reply = append(reply, &CategoryHome{
					Website: Website{
						ID:          w.ID,
						WebsiteName: w.WebsiteName,
						WebsiteUrl:  w.WebsiteUrl,
						WebsiteIcon: w.WebsiteIcon,
					},
				})
			}
		}
	}
	return reply, err
}

type ContentUsecase struct {
	repoWs WebsiteRepo
	repoCa CategoryRepo
	repoAr ArticleRepo
	log    *log.Helper
}

func NewWebsiteUsecase(repo WebsiteRepo, repoCa CategoryRepo, repoAr ArticleRepo, logger log.Logger) *ContentUsecase {
	return &ContentUsecase{repoWs: repo, repoCa: repoCa, repoAr: repoAr, log: log.NewHelper(logger)}
}

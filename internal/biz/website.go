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

// WebsiteRepo .
type WebsiteRepo interface {
	Save()
	Update()
	Delete()
	WebSiteList(context.Context, int32, int32) ([]*Website, error)
	WebSiteHome(ctx context.Context) ([]*v1.WebsiteReply, error)
}

func (ws *ContentUsecase) WebSiteList(ctx context.Context, category int32, typ int32) (wr []*Website, err error) {
	wr, err = ws.repoWs.WebSiteList(ctx, category, typ)
	return wr, err
}

func (ws *ContentUsecase) WebSiteHome(ctx context.Context) ([]*v1.WebsiteReply, error) {
	wh, err := ws.repoWs.WebSiteHome(ctx)
	return wh, err
}

type ContentUsecase struct {
	repoWs WebsiteRepo
	repoCa CategoryRepo
	log    *log.Helper
}

func NewWebsiteUsecase(repo WebsiteRepo, repoCa CategoryRepo, logger log.Logger) *ContentUsecase {
	return &ContentUsecase{repoWs: repo, repoCa: repoCa, log: log.NewHelper(logger)}
}

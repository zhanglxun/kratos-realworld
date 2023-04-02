package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/golang/protobuf/ptypes/timestamp"
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
	CreateTime  timestamp.Timestamp
	ModifyId    int64
	ModifyTime  timestamp.Timestamp
}

// WebsiteRepo .
type WebsiteRepo interface {
	WebSiteList(context.Context, int32, int32) ([]*Website, error)
}

type WebsiteUsecase struct {
	repo WebsiteRepo
	log  *log.Helper
}

func NewWebsiteUsecase(repo WebsiteRepo, logger log.Logger) *WebsiteUsecase {
	return &WebsiteUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (ws *WebsiteUsecase) WebSiteList(ctx context.Context, category int32, typ int32) (wr []*Website, err error) {

	wr, err = ws.repo.WebSiteList(ctx, category, typ)
	return wr, err
}

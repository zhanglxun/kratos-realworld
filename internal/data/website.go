package data

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"go-server/internal/biz"
)

type websiteRepo struct {
	data *Data
	log  *log.Helper
}

func (w *websiteRepo) WebSiteList(ctx context.Context, category int32, typ int32) ([]*biz.Website, error) {
	//TODO implement me
	ws, _ := w.data.db.Website.Query().Where().All(ctx)
	rv := make([]*biz.Website, 0)
	for i, value := range ws {
		fmt.Printf("slice at %d is : %s \n", i, value.WebsiteName+"-"+value.WebsiteURL)
		rv = append(rv, &biz.Website{
			ID:          value.ID,
			SortId:      value.SortID,
			Category:    value.Category,
			Type:        value.Type,
			WebsiteName: value.WebsiteName,
			WebsiteUrl:  value.WebsiteURL,
			WebsiteIcon: value.WebsiteIcon,
			Summary:     value.Summary,
		})
	}
	return rv, nil
}

// NewWebsiteRepo .
func NewWebsiteRepo(data *Data, logger log.Logger) biz.WebsiteRepo {
	return &websiteRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

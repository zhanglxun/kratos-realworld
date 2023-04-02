package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"go-server/internal/biz"
	"go-server/internal/data/ent/website"
)

type websiteRepo struct {
	data *Data
	log  *log.Helper
}

func (w *websiteRepo) Save() {
	//TODO implement me
	panic("implement me")
}

func (w *websiteRepo) Update() {
	//TODO implement me
	panic("implement me")
}

func (w *websiteRepo) Delete() {
	//TODO implement me
	panic("implement me")
}

func (w *websiteRepo) WebSiteList(ctx context.Context, category int32, typ int32) ([]*biz.Website, error) {
	//TODO implement me
	ws, _ := w.data.db.Website.Query().
		Where(
			website.CategoryEQ(category),
			website.TypeEQ(typ),
		).All(ctx)
	rv := make([]*biz.Website, 0)
	for _, value := range ws {
		//fmt.Printf("slice at %d is : %s \n", i, value.WebsiteName+"-"+value.WebsiteURL)
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

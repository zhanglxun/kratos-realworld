package service

import (
	"github.com/google/wire"
	v1 "go-server/api/realworld/v1"
	"go-server/internal/biz"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewContentService)

// RealWorldService is a greeter service.
type ContentService struct {
	v1.UnimplementedContentServiceServer
	ws *biz.ContentUsecase
	uc *biz.GreeterUsecase
}

// NewRealWorldService new a greeter service.
func NewContentService(
	uc *biz.GreeterUsecase,
	ws *biz.ContentUsecase,
) *ContentService {
	return &ContentService{uc: uc, ws: ws}
}

package service

import (
	"github.com/google/wire"
	v1 "go-server/api/realworld/v1"
	"go-server/internal/biz"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewRealWorldService)

// RealWorldService is a greeter service.
type RealWorldService struct {
	v1.UnimplementedRealWorldServer
	ws *biz.ContentUsecase
	uc *biz.GreeterUsecase
}

// NewRealWorldService new a greeter service.
func NewRealWorldService(
	uc *biz.GreeterUsecase,
	ws *biz.ContentUsecase,
) *RealWorldService {
	return &RealWorldService{uc: uc, ws: ws}
}

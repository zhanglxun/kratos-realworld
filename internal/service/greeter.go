package service

import (
	"context"

	v1 "go-server/api/realworld/v1"
	"go-server/internal/biz"
)

// SayHello implements realworld.GreeterServer.
func (s *ContentService) SayHello(ctx context.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {
	g, err := s.uc.CreateGreeter(ctx, &biz.Greeter{Hello: in.Name})
	if err != nil {
		return nil, err
	}
	return &v1.HelloReply{Message: "Hello " + g.Hello}, nil
}

package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/go-kratos/swagger-api/openapiv2"
	"github.com/gorilla/handlers"
	v1 "go-server/api/realworld/v1"
	"go-server/internal/conf"
	"go-server/internal/service"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, greeter *service.ContentService, logger log.Logger) *http.Server {
	//注册openAPIV2
	openAPIhandler := openapiv2.NewHandler()

	var opts = []http.ServerOption{
		//增加的错误处理返回码
		//http.ErrorEncoder(errorEncoder),
		http.Middleware(
			recovery.Recovery(),
		),
		//新增加跨域处理
		http.Filter(handlers.CORS(
			handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization", "AccessToken", "X-CSRF-Token", "Token", "X-Token", "X-User-Id"}),
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "HEAD", "OPTIONS"}),
			handlers.AllowedOrigins([]string{"*"}),
		)),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	//注册swagger
	srv.HandlePrefix("/q/", openAPIhandler)

	v1.RegisterContentServiceHTTPServer(srv, greeter)
	return srv
}

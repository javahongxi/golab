package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	userv1 "github.com/javahongxi/golab/kratos/proto/user/v1"
	"github.com/javahongxi/golab/kratos/internal/service"
)

// NewHTTPServer 创建 HTTP 服务器
func NewHTTPServer(logger log.Logger, userSvc *service.UserService) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			logging.Server(logger),
		),
		http.Address(":8000"),
	}

	srv := http.NewServer(opts...)
	userv1.RegisterUserServiceHTTPServer(srv, userSvc)
	return srv
}

// NewGRPCServer 创建 gRPC 服务器
func NewGRPCServer(logger log.Logger, userSvc *service.UserService) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
			logging.Server(logger),
		),
		grpc.Address(":9000"),
	}

	srv := grpc.NewServer(opts...)
	userv1.RegisterUserServiceServer(srv, userSvc)
	return srv
}

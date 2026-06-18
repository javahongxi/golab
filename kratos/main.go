package main

import (
	"os"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"

	"github.com/javahongxi/golab/kratos/internal/server"
	"github.com/javahongxi/golab/kratos/internal/service"
)

func main() {
	logger := log.With(log.NewStdLogger(os.Stdout),
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		"service.name", "kratos-demo",
	)

	// 创建用户服务
	userSvc := service.NewUserService(logger)

	// 创建 HTTP 和 gRPC 服务器
	httpSrv := server.NewHTTPServer(logger, userSvc)
	grpcSrv := server.NewGRPCServer(logger, userSvc)

	// 创建 Kratos 应用
	app := kratos.New(
		kratos.Name("kratos-demo"),
		kratos.Version("v1.0.0"),
		kratos.Logger(logger),
		kratos.Server(
			httpSrv,
			grpcSrv,
		),
	)

	log.Info("starting kratos server...")
	log.Info("HTTP server: http://localhost:8000")
	log.Info("gRPC server: grpc://localhost:9000")

	if err := app.Run(); err != nil {
		log.Error(err)
	}
}

// 确保接口被使用
var (
	_ *http.Server
	_ *grpc.Server
)

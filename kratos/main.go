package main

import (
	"fmt"
	"os"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"

	"github.com/javahongxi/golab/kratos/internal/registry"
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

	// 配置 Nacos
	nacosCfg := registry.NacosConfig{
		Addr:        getEnv("NACOS_ADDR", "127.0.0.1"),
		Port:        getEnvUint64("NACOS_PORT", 8848),
		NamespaceID: getEnv("NACOS_NAMESPACE", ""),
		Username:    getEnv("NACOS_USERNAME", "nacos"),
		Password:    getEnv("NACOS_PASSWORD", "7fDJZBbiLzO2"),
	}

	// 创建 Nacos 注册中心
	nacosRegistry, err := registry.NewNacosRegistry(nacosCfg, logger)
	if err != nil {
		log.Errorf("failed to create nacos registry: %v", err)
		log.Info("starting server without service registration...")
		nacosRegistry = nil
	}

	// 创建 Kratos 应用
	appOpts := []kratos.Option{
		kratos.Name("kratos-demo"),
		kratos.Version("v1.0.0"),
		kratos.Logger(logger),
		kratos.Server(
			httpSrv,
			grpcSrv,
		),
	}

	// 如果 Nacos 注册中心可用，则添加注册器
	if nacosRegistry != nil {
		appOpts = append(appOpts, kratos.Registrar(nacosRegistry))
	}

	app := kratos.New(appOpts...)

	log.Info("starting kratos server...")
	log.Info("HTTP server: http://localhost:8000")
	log.Info("gRPC server: grpc://localhost:9000")
	if nacosRegistry != nil {
		log.Info("Nacos registry: service registered successfully")
	}

	if err := app.Run(); err != nil {
		log.Error(err)
	}
}

// getEnv 获取环境变量，支持默认值
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// getEnvUint64 获取环境变量并转换为 uint64，支持默认值
func getEnvUint64(key string, defaultValue uint64) uint64 {
	if value := os.Getenv(key); value != "" {
		var result uint64
		if _, err := fmt.Sscanf(value, "%d", &result); err == nil {
			return result
		}
	}
	return defaultValue
}

// 确保接口被使用
var (
	_ *http.Server
	_ *grpc.Server
)

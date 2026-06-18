package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/http"

	nacosreg "github.com/javahongxi/golab/kratos/internal/registry"
	userv1 "github.com/javahongxi/golab/kratos/proto/user/v1"
)

func main() {
	logger := log.With(log.NewStdLogger(os.Stdout),
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
	)

	// 配置 Nacos
	nacosCfg := nacosreg.NacosConfig{
		Addr:        getEnv("NACOS_ADDR", "127.0.0.1"),
		Port:        getEnvUint64("NACOS_PORT", 8848),
		NamespaceID: getEnv("NACOS_NAMESPACE", ""),
		Username:    getEnv("NACOS_USERNAME", "nacos"),
		Password:    getEnv("NACOS_PASSWORD", "7fDJZBbiLzO2"),
	}

	// 创建 Nacos 服务发现客户端
	discovery, err := nacosreg.NewNacosDiscovery(nacosCfg, logger)
	if err != nil {
		log.Errorf("failed to create nacos discovery: %v", err)
		return
	}

	// 服务名称（Kratos 会为每个 Server 分别注册，服务名自动加上 .http 或 .grpc 后缀）
	serviceName := "kratos-demo.http"

	log.Infof("discovering service: %s", serviceName)

	// 获取服务实例
	watcher, err := discovery.Watch(context.Background(), serviceName)
	if err != nil {
		log.Errorf("failed to watch service: %v", err)
		return
	}

	// 监听服务变化
	go func() {
		for {
			services, err := watcher.Next()
			if err != nil {
				log.Errorf("failed to get services: %v", err)
				continue
			}

			log.Infof("found %d service instances", len(services))
			for _, svc := range services {
				for _, endpoint := range svc.Endpoints {
					log.Infof("  - %s (name: %s, version: %s)",
						endpoint,
						svc.Name,
						svc.Version,
					)
				}
			}
		}
	}()

	// 等待服务发现
	time.Sleep(2 * time.Second)

	// 获取服务实例列表
	instances, err := discovery.GetService(context.Background(), serviceName)
	if err != nil {
		log.Errorf("failed to get service: %v", err)
		return
	}

	if len(instances) == 0 {
		log.Info("no service instances found")
		return
	}

	// 使用第一个实例创建客户端
	instance := instances[0]
	if len(instance.Endpoints) == 0 {
		log.Info("no endpoints found for service instance")
		return
	}
	endpoint := instance.Endpoints[0]
	log.Infof("connecting to service at: %s", endpoint)

	// 创建 HTTP 客户端
	conn, err := http.NewClient(context.Background(), http.WithEndpoint(endpoint))
	if err != nil {
		log.Errorf("failed to create client: %v", err)
		return
	}

	// 创建用户服务客户端
	client := userv1.NewUserServiceHTTPClient(conn)

	// 测试创建用户
	user, err := client.CreateUser(context.Background(), &userv1.CreateUserRequest{
		Username: "discovery-test",
		Nickname: "服务发现测试",
		Email:    "discovery@example.com",
		Age:      25,
	})
	if err != nil {
		log.Errorf("failed to create user: %v", err)
		return
	}

	log.Infof("created user via service discovery: %+v", user)

	// 测试获取用户
	retrievedUser, err := client.GetUser(context.Background(), &userv1.GetUserRequest{
		Id: user.Id,
	})
	if err != nil {
		log.Errorf("failed to get user: %v", err)
		return
	}

	log.Infof("retrieved user via service discovery: %+v", retrievedUser)

	// 测试列出用户
	users, err := client.ListUsers(context.Background(), &userv1.ListUsersRequest{
		Page:     1,
		PageSize: 10,
	})
	if err != nil {
		log.Errorf("failed to list users: %v", err)
		return
	}

	log.Infof("listed users via service discovery: %d users", users.Total)

	log.Info("service discovery test completed successfully")
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
var _ registry.Discovery

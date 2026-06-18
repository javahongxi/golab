package registry

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	kratosnacos "github.com/go-kratos/kratos/contrib/registry/nacos/v2"
)

// NacosConfig Nacos 配置
type NacosConfig struct {
	Addr        string
	Port        uint64
	NamespaceID string
	Username    string
	Password    string
}

// NewNacosRegistry 创建 Nacos 注册中心
func NewNacosRegistry(cfg NacosConfig, logger log.Logger) (registry.Registrar, error) {
	// 创建 Nacos 客户端配置
	sc := []constant.ServerConfig{
		*constant.NewServerConfig(cfg.Addr, cfg.Port),
	}

	// 创建 Nacos 客户端配置
	cc := *constant.NewClientConfig(
		constant.WithNamespaceId(cfg.NamespaceID),
		constant.WithUsername(cfg.Username),
		constant.WithPassword(cfg.Password),
		constant.WithTimeoutMs(5000),
		constant.WithLogLevel("debug"),
	)

	// 创建命名客户端
	namingClient, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		},
	)
	if err != nil {
		return nil, err
	}

	log.Info("nacos registry connected successfully")

	// 返回 Kratos Nacos 注册中心
	return kratosnacos.New(namingClient), nil
}

// NewNacosDiscovery 创建 Nacos 服务发现客户端
func NewNacosDiscovery(cfg NacosConfig, logger log.Logger) (registry.Discovery, error) {
	// 创建 Nacos 客户端配置
	sc := []constant.ServerConfig{
		*constant.NewServerConfig(cfg.Addr, cfg.Port),
	}

	// 创建 Nacos 客户端配置
	cc := *constant.NewClientConfig(
		constant.WithNamespaceId(cfg.NamespaceID),
		constant.WithUsername(cfg.Username),
		constant.WithPassword(cfg.Password),
		constant.WithTimeoutMs(5000),
		constant.WithLogLevel("debug"),
	)

	// 创建命名客户端
	namingClient, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		},
	)
	if err != nil {
		return nil, err
	}

	log.Info("nacos discovery connected successfully")

	// 返回 Kratos Nacos 服务发现
	return kratosnacos.New(namingClient), nil
}

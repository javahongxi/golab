# GoLab

> Go 语言全栈实战：基础语法 · Gin Web · Kratos 微服务 · 分布式爬虫 🏡

```bash
git clone https://github.com/javahongxi/golab.git
cd golab
go get ./...
```

## 📚 学习路径

GoLab 按照从入门到实战的学习路径组织，帮助你系统掌握 Go 语言生态：

### 1. 基础语法 (lang/)

从 Go 语言的基本概念开始，涵盖核心语法和常用模式：

- **基础数据类型**：变量、常量、类型转换、指针
- **控制流程**：条件分支、循环、switch、defer
- **函数**：函数定义、可变参数、闭包、递归
- **容器类型**：数组、切片、Map、字符串处理
- **正则表达式**：regexp 包的使用
- **错误处理**：error 类型、panic/recover、自定义错误
- **泛型**：Go 1.18+ 泛型工具函数（[generic/utils.go](lang/generic/utils.go)）
- **Embed**：Go 1.16+ Embed 内嵌文件（[embed/entry.go](lang/embed/entry.go)）
- **原子操作**：sync/atomic 原子计数器（[atomic/atomic.go](lang/basic/atomic/atomic.go)）
- **定时任务**：time.Ticker 定时执行（[timer/timer.go](lang/timer/timer.go)）

详细文档：[lang/README.md](lang/README.md)

### 2. 并发编程 (lang/channel/、lang/goroutine/、pipeline/)

深入理解 Go 并发模型：

- **Goroutine**：轻量级线程、并发执行
- **Channel**：通道通信、select、缓冲通道、done channel
- **Pipeline**：并行处理管道，外部排序演示
  - 单机版：[pipeline/demo/main.go](pipeline/demo/main.go)
  - 网络版：[pipeline/netdemo/main.go](pipeline/netdemo/main.go)

### 3. Web 开发 (gin/、gorm/)

使用 Gin 框架构建完整的 Web 应用：

- **Gin 框架**：路由、中间件、请求处理
- **中间件链**：日志、认证、限流、熔断、链路追踪、恢复
- **配置管理**：Viper 配置加载、热更新（[gin/config/config.go](gin/config/config.go)）
- **数据库**：GORM ORM（[gorm/](gorm/)）、Redis 缓存（[gin/cache/redis.go](gin/cache/redis.go)）
- **认证授权**：JWT Token、中间件认证
- **可观测性**：OpenTelemetry 分布式追踪（[gin/otel/otel.go](gin/otel/otel.go)）、Zap 结构化日志
- **API 文档**：[gin/README.md](gin/README.md)

### 4. 微服务 (grpc/、kratos/)

构建生产级微服务架构：

- **gRPC**：高性能 RPC 框架、Protobuf 定义、代码生成
  - 服务端：[grpc/server/main.go](grpc/server/main.go)
  - 客户端：[grpc/client/main.go](grpc/client/main.go)
- **Kratos**：轻量级微服务框架、HTTP/gRPC 双协议
  - 服务注册发现：Nacos 集成（[kratos/internal/registry/nacos.go](kratos/internal/registry/nacos.go)）
  - 使用文档：[kratos/README.md](kratos/README.md)

### 5. 分布式爬虫 (crawler/、crawler_distributed/)

构建单任务和分布式爬虫系统：

- **单任务版**：[crawler/main.go](crawler/main.go)
- **分布式版**：基于 RPC 的分布式爬虫，支持 Elasticsearch 存储
  - 存储服务：[crawler_distributed/persist/server/itemsaver.go](crawler_distributed/persist/server/itemsaver.go)
  - Worker 服务：[crawler_distributed/worker/server/worker.go](crawler_distributed/worker/server/worker.go)

## 🛠️ 技术栈

| 分类 | 技术 | 说明 |
|------|------|------|
| 框架 | Gin | Web 开发框架 |
| 框架 | Kratos | 微服务框架 |
| ORM | GORM | 数据库 ORM |
| RPC | gRPC | 高性能 RPC |
| 服务发现 | Nacos | 服务注册与发现 |
| 分布式追踪 | OpenTelemetry | 链路追踪 |
| 日志 | Zap | 结构化日志 |
| 缓存 | Redis | 分布式缓存 |
| 配置 | Viper | 配置管理 |
| 搜索引擎 | Elasticsearch | 全文检索 |
| 消息队列 | Kafka | 异步消息 |

## 🚀 快速开始

### 爬虫项目演示

```bash
# 启动相亲网站 mock
go run lovemock/main.go

# 启动单任务爬虫
go run crawler/main.go
```

### 分布式爬虫演示

```bash
# 启动存储服务
cd crawler_distributed/persist/server
go run itemsaver.go --port=9090

# 启动多个 worker
cd crawler_distributed/worker/server
go run worker.go --port=9091
go run worker.go --port=9092

# 启动爬虫
cd crawler_distributed
go run main.go --itemsaver_host=:9090 --worker_hosts=:9091,:9092
```

### Gin 应用

```bash
cd gin
go run main.go
```

### Kratos 应用

```bash
cd kratos
go run main.go
```

## 📖 Go Projects

精选 Go 语言优秀项目：

- [Go](https://github.com/golang/go) - Go 语言官方仓库
- [Docker](https://github.com/moby/moby) - 容器化平台
- [Kubernetes](https://github.com/kubernetes/kubernetes) - 容器编排
- [etcd](https://github.com/etcd-io/etcd) - 分布式键值存储
- [Gin](https://github.com/gin-gonic/gin) - Web 框架
- [Consul](https://github.com/hashicorp/consul) - 服务发现与配置
- [gRPC-Go](https://github.com/grpc/grpc-go) - gRPC Go 实现
- [Prometheus](https://github.com/prometheus/prometheus) - 监控系统
- [MinIO](https://github.com/minio/minio) - 对象存储
- [TiDB](https://github.com/pingcap/tidb) - 分布式数据库
- [Milvus](https://github.com/milvus-io/milvus) - 向量数据库
- [Caddy](https://github.com/caddyserver/caddy) - Web 服务器

© [hongxi.org](http://hongxi.org)
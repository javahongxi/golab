# Kratos 微服务示例

基于 [Go-Kratos](https://github.com/go-kratos/kratos) v2 框架实现的微服务示例项目，展示如何使用 Kratos 构建同时支持 HTTP 和 gRPC 协议的用户管理服务。

## 目录结构

```
kratos/
├── main.go                      # 应用入口
├── internal/
│   ├── server/
│   │   └── server.go           # HTTP 和 gRPC 服务器配置
│   └── service/
│       └── user.go             # 用户服务实现
├── proto/
│   └── user/
│       └── v1/
│           ├── user.proto      # Proto 定义文件
│           ├── user.pb.go      # Protobuf 生成代码
│           ├── user_http.pb.go # HTTP 路由生成代码
│           └── user_grpc.pb.go # gRPC 生成代码
└── README.md
```

## 功能特性

- **双协议支持**：同时提供 HTTP（:8000）和 gRPC（:9000）服务
- **RESTful API**：基于 Google API 设计规范的 HTTP 接口
- **中间件集成**：
  - Recovery：自动捕获 panic 并恢复
  - Logging：请求日志记录
- **代码生成**：使用 protoc 和 Kratos 插件自动生成代码
- **内存存储**：使用 map 实现简单的用户数据管理（示例用途）

## 技术栈

- **框架**：[Kratos v2.9.2](https://github.com/go-kratos/kratos)
- **协议**：Protocol Buffers
- **HTTP 路由**：基于 Google API HTTP 注解
- **gRPC**：标准 gRPC 服务
- **日志**：Kratos 内置日志系统

## 快速开始

### 前置要求

- Go 1.26.3+
- Protocol Buffers 编译器（protoc）
- protoc 插件：
  - `protoc-gen-go`
  - `protoc-gen-go-grpc`
  - `protoc-gen-go-http`

### 安装依赖

```bash
# 安装 Kratos 命令行工具
go install github.com/go-kratos/kratos/cmd/kratos/v2@latest

# 安装 protoc 插件
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
go install github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v2@latest
```

### 启动服务

```bash
# 在项目根目录执行
go run kratos/main.go
```

服务启动后会输出：

```
INFO starting kratos server...
INFO HTTP server: http://localhost:8000
INFO gRPC server: grpc://localhost:9000
```

## API 接口

### HTTP 接口

| 方法 | 路径 | 说明 | 请求体 |
|------|------|------|--------|
| POST | `/v1/users` | 创建用户 | JSON |
| GET | `/v1/users/{id}` | 获取用户 | - |
| PUT | `/v1/users/{id}` | 更新用户 | JSON |
| DELETE | `/v1/users/{id}` | 删除用户 | - |
| GET | `/v1/users` | 列出用户（分页） | - |

### 数据模型

**User**

```json
{
  "id": 1,
  "username": "test",
  "nickname": "Test User",
  "email": "test@example.com",
  "age": 25
}
```

**CreateUserRequest**

```json
{
  "username": "test",
  "nickname": "Test User",
  "email": "test@example.com",
  "age": 25
}
```

**UpdateUserRequest**

```json
{
  "id": 1,
  "nickname": "Updated Name",
  "email": "updated@example.com",
  "age": 26
}
```

**ListUsersRequest**

```json
{
  "page": 1,
  "page_size": 10
}
```

## 测试示例

### HTTP 接口测试

#### 1. 创建用户

```bash
curl -X POST http://localhost:8000/v1/users \
  -H "Content-Type: application/json" \
  -d '{
    "username": "zhangsan",
    "nickname": "张三",
    "email": "zhangsan@example.com",
    "age": 25
  }'
```

响应：

```json
{
  "id": 1,
  "username": "zhangsan",
  "nickname": "张三",
  "email": "zhangsan@example.com",
  "age": 25
}
```

#### 2. 获取用户

```bash
curl http://localhost:8000/v1/users/1
```

响应：

```json
{
  "id": 1,
  "username": "zhangsan",
  "nickname": "张三",
  "email": "zhangsan@example.com",
  "age": 25
}
```

#### 3. 更新用户

```bash
curl -X PUT http://localhost:8000/v1/users/1 \
  -H "Content-Type: application/json" \
  -d '{
    "id": 1,
    "nickname": "张三丰",
    "age": 30
  }'
```

响应：

```json
{
  "id": 1,
  "username": "zhangsan",
  "nickname": "张三丰",
  "email": "zhangsan@example.com",
  "age": 30
}
```

#### 4. 列出用户

```bash
curl "http://localhost:8000/v1/users?page=1&page_size=10"
```

响应：

```json
{
  "users": [
    {
      "id": 1,
      "username": "zhangsan",
      "nickname": "张三丰",
      "email": "zhangsan@example.com",
      "age": 30
    }
  ],
  "total": 1
}
```

#### 5. 删除用户

```bash
curl -X DELETE http://localhost:8000/v1/users/1
```

响应：空（成功）

### gRPC 接口测试

使用 [grpcurl](https://github.com/fullstorydev/grpcurl) 工具测试 gRPC 接口：

```bash
# 安装 grpcurl
go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest

# 创建用户
grpcurl -plaintext -d '{
  "username": "lisi",
  "nickname": "李四",
  "email": "lisi@example.com",
  "age": 28
}' localhost:9000 user.v1.UserService/CreateUser

# 获取用户
grpcurl -plaintext -d '{"id": 1}' localhost:9000 user.v1.UserService/GetUser

# 列出用户
grpcurl -plaintext -d '{"page": 1, "page_size": 10}' localhost:9000 user.v1.UserService/ListUsers

# 更新用户
grpcurl -plaintext -d '{
  "id": 1,
  "nickname": "李四光",
  "age": 35
}' localhost:9000 user.v1.UserService/UpdateUser

# 删除用户
grpcurl -plaintext -d '{"id": 1}' localhost:9000 user.v1.UserService/DeleteUser
```

## 代码生成

如果需要修改 proto 文件并重新生成代码：

```bash
# 确保 third_party/googleapis 已存在
cd /Users/hongxi/dev/golab
git clone --depth 1 https://github.com/googleapis/googleapis.git third_party/googleapis

# 生成代码
protoc --proto_path=kratos/proto \
       --proto_path=third_party/googleapis \
       --go_out=paths=source_relative:kratos/proto \
       --go-http_out=paths=source_relative:kratos/proto \
       --go-grpc_out=paths=source_relative:kratos/proto \
       kratos/proto/user/v1/user.proto
```

## 架构说明

### 服务层（Service）

`internal/service/user.go` 实现了 `UserService` 接口，提供用户管理的核心业务逻辑：

- 使用 `sync.RWMutex` 保证并发安全
- 内存存储用户数据（map 结构）
- 支持分页查询
- 集成 Kratos 日志系统

### 服务器层（Server）

`internal/server/server.go` 负责创建和配置 HTTP 和 gRPC 服务器：

- **HTTP 服务器**：监听 `:8000`，注册 RESTful 路由
- **gRPC 服务器**：监听 `:9000`，注册 gRPC 服务
- 两个服务器都配置了 `recovery` 和 `logging` 中间件

### 应用入口（Main）

`main.go` 是应用的入口点：

- 初始化日志系统
- 创建用户服务实例
- 创建 HTTP 和 gRPC 服务器
- 使用 Kratos App 统一管理服务器生命周期
- 支持优雅关闭

## 扩展建议

这个示例展示了 Kratos 的基本用法，实际项目中可以扩展以下功能：

- **数据库集成**：将内存存储替换为 MySQL/PostgreSQL
- **缓存层**：添加 Redis 缓存提升性能
- **认证授权**：实现 JWT 或 OAuth2 认证
- **配置管理**：使用 Kratos Config 模块管理配置
- **服务发现**：集成 Nacos/Consul/Etcd 实现服务注册与发现
- **链路追踪**：集成 OpenTelemetry 实现分布式追踪
- **错误处理**：使用 Kratos Errors 模块定义标准化错误
- **数据验证**：使用 protoc-gen-validate 进行参数校验

## 参考资源

- [Kratos 官方文档](https://go-kratos.dev/docs/)
- [Kratos GitHub](https://github.com/go-kratos/kratos)
- [Protocol Buffers](https://protobuf.dev/)
- [gRPC 官方文档](https://grpc.io/docs/)
- [Google API Design Guide](https://cloud.google.com/apis/design)

## License

MIT

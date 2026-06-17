# Gin Demo API 文档

## 基础信息

- **服务地址**: `http://localhost:8080`
- **统一响应格式**:
```json
{
  "code": 0,
  "message": "success",
  "data": {}
}
```

## 错误码说明

| 错误码 | 含义 | HTTP 状态码 |
|--------|------|------------|
| 0 | 成功 | 200 |
| 1 | 服务器错误 | 500 |
| 400 | 参数校验失败 | 400 |
| 401 | 未认证 | 401 |
| 403 | 无权限 | 403 |
| 404 | 资源未找到 | 404 |
| 409 | 冲突（如用户名已存在） | 200 |
| 429 | 请求过于频繁 | 200 |
| 503 | 服务不可用（熔断） | 200 |

## 健康检查接口

| 方法 | 路径 | 认证 | 说明 |
|------|------|------|------|
| GET | `/ping` | 否 | 健康检查 |
| GET | `/health` | 否 | 服务状态 |

### GET /ping

**请求示例**:
```bash
curl http://localhost:8080/ping
```

**响应示例**:
```json
{"code":0,"message":"success","data":{"message":"pong","request_id":"..."}}
```

### GET /health

**请求示例**:
```bash
curl http://localhost:8080/health
```

**响应示例**:
```json
{"code":0,"message":"success","data":{"status":"ok","service":"gin-demo"}}
```

## 认证接口

| 方法 | 路径 | 认证 | 说明 |
|------|------|------|------|
| POST | `/api/v1/auth/register` | 否 | 用户注册 |
| POST | `/api/v1/auth/login` | 否 | 用户登录 |

### POST /api/v1/auth/register

**请求示例**:
```bash
curl -X POST http://localhost:8080/api/v1/auth/register \
  -H "Content-Type: application/json" \
  -d '{"username":"test","password":"123456","nickname":"测试","gender":1,"age":25}'
```

**请求体**:
| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| username | string | 是 | 3-20 字符 |
| password | string | 是 | 6-20 字符 |
| nickname | string | 否 | 最多 20 字符 |
| gender | int | 否 | 0/1/2 |
| age | int | 否 | 0-150 |

**成功响应**:
```json
{"code":0,"message":"success","data":{"id":1,"username":"test","nickname":"测试","gender":1,"age":25}}
```

### POST /api/v1/auth/login

**请求示例**:
```bash
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{"username":"test","password":"123456"}'
```

**请求体**:
| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| username | string | 是 | 用户名 |
| password | string | 是 | 密码 |

**成功响应**:
```json
{"code":0,"message":"success","data":{"user":{},"token":"eyJhbGciOiJIUzI1Ni..."}}
```

## 用户接口

| 方法 | 路径 | 认证 | 说明 |
|------|------|------|------|
| GET | `/api/v1/users` | 否 | 列出用户（分页） |
| GET | `/api/v1/users/:id` | 否 | 查询单个用户 |
| POST | `/api/v1/users` | 否 | 创建用户 |
| PUT | `/api/v1/users/:id` | **是** | 更新用户 |
| DELETE | `/api/v1/users/:id` | **是** | 删除用户 |

### GET /api/v1/users

**请求示例**:
```bash
curl "http://localhost:8080/api/v1/users?page=1&limit=10"
```

**查询参数**:
| 参数 | 类型 | 默认值 | 说明 |
|------|------|--------|------|
| page | int | 1 | 页码 |
| limit | int | 10 | 每页数量(1-100) |

**成功响应**:
```json
{"code":0,"message":"success","data":{"data":[],"total":0,"page":1,"limit":10}}
```

### GET /api/v1/users/:id

**请求示例**:
```bash
curl http://localhost:8080/api/v1/users/1
```

**路径参数**:
| 参数 | 类型 | 说明 |
|------|------|------|
| id | uint64 | 用户 ID |

**成功响应**:
```json
{"code":0,"message":"success","data":{"id":1,"username":"test","nickname":"测试"}}
```

### POST /api/v1/users

**请求示例**:
```bash
curl -X POST http://localhost:8080/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{"username":"user1","nickname":"用户1","gender":1,"age":20}'
```

**请求体**:
| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| username | string | 是 | 3-20 字符 |
| nickname | string | 否 | 最多 20 字符 |
| gender | int | 否 | 0/1/2 |
| age | int | 否 | 0-150 |

### PUT /api/v1/users/:id

**请求示例**:
```bash
curl -X PUT http://localhost:8080/api/v1/users/1 \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <token>" \
  -d '{"nickname":"新昵称","age":21}'
```

**请求体**:
| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| nickname | string | 否 | 最多 20 字符 |
| age | int | 否 | 0-150 |

### DELETE /api/v1/users/:id

**请求示例**:
```bash
curl -X DELETE http://localhost:8080/api/v1/users/1 \
  -H "Authorization: Bearer <token>"
```

## 认证方式

所有需要认证的接口，需在请求头中携带:
```
Authorization: Bearer <token>
```

## 限流说明

- 默认限制：每分钟 100 次请求（基于客户端 IP）
- 超过限制返回：`{"code":429,"message":"too many requests"}`

## 熔断说明

- 当连续 5 次返回 500 错误时触发熔断
- 熔断后所有请求返回：`{"code":503,"message":"service unavailable"}`
- 熔断 30 秒后自动尝试恢复
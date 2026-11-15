# 生鲜零售平台后端 MVP（go-zero 单体 API）

生鲜即时零售平台两周项目的后端 MVP 基础框架：单体 go-zero API + MySQL + Redis，统一鉴权/错误码/响应结构，18 个 MVP 接口契约已固化（未实现的返回 50101 占位），各模块按分支填 logic 即可。此前误建的 PopPulse 潮玩电商代码已归档（tag `v0.1-poppulse-archived`）。

## 技术栈与版本（锁定）

| 工具 | 版本 |
|---|---|
| Go | 1.25+ |
| go-zero | v1.10.3（见 go.mod，不要随意升级） |
| goctl | 1.10.2（生成 handler/logic/types 用） |
| MySQL | 8.4（Docker Compose） |
| Redis | 7-alpine（Docker Compose，仅缓存加速，不承载库存事实） |

## 快速启动（全新环境）

前置：安装 Go、Docker Desktop（Windows 需启用 WSL2）、goctl。

```powershell
# 1. 启动 MySQL + Redis
docker compose -f deployments/docker-compose.yml up -d

# 2. 初始化表结构 + 种子数据（务必带 --default-character-set=utf8mb4，否则中文乱码）
docker exec -i fresh-mysql mysql -uroot -ppoppulse123 --default-character-set=utf8mb4 < deployments/migrations/001_init.sql
docker exec -i fresh-mysql mysql -uroot -ppoppulse123 --default-character-set=utf8mb4 < scripts/seed.sql

# 3. 准备本地配置（仅第一次）
Copy-Item etc/fresh.example.yaml etc/fresh.yaml

# 4. 启动服务
go run fresh.go
```

验证：

```powershell
curl http://localhost:8888/api/v1/health/live
curl http://localhost:8888/api/v1/products
```

## 种子账号

| 账号 | 密码 | 说明 |
|---|---|---|
| user01 | User@123456 | 演示用户 |
| admin01 | Admin@123456 | 管理入口预留 |

也可 `POST /api/v1/auth/register` 自行注册。

## 目录结构

```
fresh.api                 # 接口契约 v1（18 接口，修改须经组长 A）
fresh.go                  # 入口：统一响应/错误处理注册
etc/fresh.example.yaml    # 配置模板；本地 etc/fresh.yaml 不提交
internal/config           # MySQL/Redis/Auth/Mock/Ai
internal/middleware       # operator 角色校验（管理入口用）
internal/xerr             # 业务错误码（见下表）
internal/response         # {code, message, data} 统一响应
internal/handler|logic|types  # goctl 生成，分组：auth/user/product/cart/order/ai/health
internal/model            # 数据访问（统一 sqlx）
internal/ctxdata          # JWT claims 读取与 token 生成
internal/ai               # AI 适配器接口 + MockProvider（真实模型实现同一接口）
deployments/migrations    # 迁移脚本（上线后只追加）
scripts/seed.sql          # 30 个生鲜商品种子；scripts/hashgen 生成密码哈希
docs/future/              # 二期全量接口/表设计参考（api.md 等，勿当本周范围）
docs/plans/               # 设计与计划文档
```

## 已实现接口

- `GET /health/live`、`GET /health/ready`（检查 MySQL+Redis）
- `POST /auth/register`（成功即登录）、`POST /auth/login`（bcrypt + JWT）
- `GET /users/me`

其余接口路由/权限已固化，调用返回 `501/50101 占位`，由各模块负责人填 logic。

## 错误码（对齐接口文档）

| code | 含义 |
|---|---|
| 40001 | 参数错误 |
| 40101 | 未登录或 token 失效 |
| 40102 | 用户名或密码错误 |
| 40301 | 无权限 |
| 40401 | 资源不存在 |
| 40901 | 幂等键重复但请求内容不同 |
| 40902 | 库存不足 |
| 40903 | 订单状态不允许此操作 |
| 42210/42211 | AI 输出校验失败 / 食材无在售匹配 |
| 42901 | 限流 |
| 50101 | 接口未实现占位 |
| 50001 | 服务内部错误 |
| 50301 | 外部服务降级 |

## 分工与分支

| 成员 | 分支 | 模块 |
|---|---|---|
| A | `feature/account` | 注册登录、JWT、地址、工程集成 |
| B | `feature/product` | 分类/商品/购物车、热点缓存与并发聚合 |
| C | `feature/trade` | 事务下单、防超卖、幂等、模拟支付、取消 |
| D | `feature/delivery` | 配送事件、超时扫描 worker pool |
| E | `feature/ai` | AI 适配器真实接入、限流降级 |
| F | `feature/qa` | 契约维护、集成测试、压测 |

## 关键规则（违反 = 返工）

1. **防超卖**：`UPDATE products SET stock=stock-? WHERE id=? AND stock>=?` 条件更新 + 检查影响行数，多商品按 id 排序扣减，任一失败整体回滚
2. **幂等**：下单必须校验 `Idempotency-Key`；同键同参返回原订单、同键异参 409/40901
3. **计价**：服务端按 DB 实时价格计算，禁止采信客户端金额
4. **状态机**：pending_payment → paid → delivering → delivered；取消/超时只能从 pending_payment；逆向跳跃拒绝
5. **金额整数分（bigint）**；ID JSON 输出字符串
6. **越权**：非本人资源 404；地址/订单/购物车全部按 user_id 过滤
7. **AI**：只读商品；不虚构 SKU/价格/库存；响应带 mock 字段；并发信号量隔离
8. **模拟支付/履约**：仅 `Mock.Enabled: true` 可用，生产必须 false

## 验收硬指标（计划书）

- 库存 100，200 个不同幂等键并发抢购单件：恰成功 100 单、无负库存
- 同一键并发 50 次：仅一笔订单、一次扣减
- 支付与取消竞争：仅一种合法终态，不重复释放
- 超时任务重跑无副作用；普通用户不能读写他人地址/订单
- AI 非法 JSON、超时有可读结果；Mock/真实模式可辨识
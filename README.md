# PopPulse 后端 MVP（go-zero 单体 API）

潮玩脉动 PopPulse 六人七天项目的后端 MVP 基础框架：单体 go-zero API + 单 MySQL，统一鉴权/错误码/响应结构，39 个 MVP 接口契约已固化（未实现的返回 50101 占位），各模块按分支填 logic 即可。

## 技术栈与版本（锁定）

| 工具 | 版本 |
|---|---|
| Go | 1.25+（开发机为 1.27） |
| go-zero | v1.10.3（见 go.mod，不要随意升级） |
| goctl | 1.10.2（生成 handler/logic/types 用） |
| MySQL | 8.4（Docker Compose） |

## 快速启动（全新环境）

前置：安装 Go、Docker Desktop（Windows 需启用 WSL2）、goctl。

```powershell
# 1. 启动 MySQL
docker compose -f deployments/docker-compose.yml up -d

# 2. 初始化表结构 + 种子数据（容器名 poppulse-mysql；务必带 --default-character-set=utf8mb4，否则中文乱码）
docker exec -i poppulse-mysql mysql -uroot -ppoppulse123 --default-character-set=utf8mb4 < deployments/migrations/001_init.sql
docker exec -i poppulse-mysql mysql -uroot -ppoppulse123 --default-character-set=utf8mb4 < scripts/seed.sql

# 3. 准备本地配置（仅第一次）
Copy-Item etc/poppulse.example.yaml etc/poppulse.yaml

# 4. 启动服务
go run poppulse.go
```

验证：

```powershell
# 健康检查
curl http://localhost:8888/api/v1/health/live

# 登录（user01/User@123456，运营 admin01/Admin@123456）
curl -X POST http://localhost:8888/api/v1/auth/login -H "Content-Type: application/json" -d '{"username":"user01","password":"User@123456"}'

# 带 token 访问
curl http://localhost:8888/api/v1/users/me -H "Authorization: Bearer <access_token>"
```

## 种子账号

| 账号 | 密码 | 角色 |
|---|---|---|
| user01 | User@123456 | user（普通用户） |
| admin01 | Admin@123456 | operator（运营） |

## 目录结构

```
poppulse.api              # 接口契约 v1（39 接口，修改须经组长 A）
poppulse.go               # 入口：统一响应/错误处理注册
etc/poppulse.example.yaml # 配置模板；本地 etc/poppulse.yaml 不提交
internal/config           # MySQL/Auth/Mock/PayTimeout
internal/middleware       # operator 角色校验（admin 路由挂载）
internal/xerr             # 业务错误码（400xx 参数、300xx 业务、50101 未实现…）
internal/response         # {code, message, data, trace_id} 统一响应
internal/handler|logic|types  # goctl 生成，按模块 group 分目录
internal/model            # 数据访问（goctl model 生成 + 手写补充）
internal/ctxdata          # JWT claims 读取与 token 生成
deployments/migrations    # 迁移脚本（只追加，不修改已合并版本）
scripts/seed.sql          # 种子数据；scripts/hashgen 生成密码哈希
docs/plans/               # 设计与计划文档
```

## 已实现接口（D1）

- M01 `GET /api/v1/health/live` 进程存活
- M02 `GET /api/v1/health/ready` 服务就绪（检查 MySQL，失败 503/90002）
- M03 `POST /api/v1/auth/login` 登录（bcrypt + JWT，claims 含 userId/role）
- M04 `GET /api/v1/users/me` 当前用户
- M05 `PATCH /api/v1/users/me` 修改资料（PATCH 语义，未传字段保留）

其余 34 个接口路由/权限已固化，调用返回 `50101 接口尚未实现`，由各模块负责人在对应 logic 中实现。

## 组员开发指引

1. `git clone` 后按上述步骤启动成功，再切自己的分支：`git checkout -b feature/<模块>`（user/product/stock/order/payment/aftersale）
2. 只在自己模块的 `internal/logic/<group>/` 下填业务逻辑；需要改 `poppulse.api` 或公共文件（config/xerr/response/middleware）时先在群里找组长
3. 需要新表字段：写新的迁移脚本 `002_xxx.sql` 追加，禁止修改 `001_init.sql`
4. 数据访问统一 `goctl model` 或 `internal/model` 手写 sqlx，同一套方式不混用
5. 提交前：`gofmt -w .` + `go build ./...` + `go vet ./...`
6. 涉及钱：金额一律整数分；涉及 ID：JSON 输出一律字符串
7. admin 接口默认带 operator 校验，仍须在 logic 里校验数据归属（本人订单 404 规则等）

## 关键规则（违反 = 返工）

- 库存防超卖：`UPDATE ... WHERE total - reserved - sold >= ?` 条件更新 + 检查影响行数，禁止先查后扣
- 幂等：创建订单必须校验 `Idempotency-Key` 请求头；同键同参返回原结果、同键异参返回 409/40002
- 状态机：非法状态转换拒绝；取消与支付竞争只有一种合法终态
- 模拟支付/退款：仅 `Mock.Enabled: true` 时可用，生产必须关闭
- 越权：非本人资源统一 404；已认证但无后台角色返回 403

## 关闭模拟能力（验收/演示用）

编辑 `etc/poppulse.yaml`：

```yaml
Mock:
  Enabled: false
```

## 协作约定

- `main` 始终可启动；每次合并至少一人复核
- 迁移脚本只追加；公共 .api 由组长协调，禁止多人直接覆盖
- 提交信息格式：`<模块>: <做了什么>`，如 `order: 实现创建订单事务`

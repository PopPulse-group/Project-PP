# D1 基础框架设计文档

日期：2026-11-14　状态：已确认（组长批准）

## 决策记录

| 决策项 | 结论 |
|---|---|
| 架构形态 | 单体 go-zero API 应用 + 单 MySQL（计划书第三节建议） |
| DB 访问方式 | goctl model + go-zero sqlx，本周不改 |
| 旧 demo（product/、product-api/） | 删除，Git 历史可找回 |
| 仓库 | github.com/PopPulse-group/Project-PP（私有） |
| 版本 | Go 1.25+（工具链 1.27）、go-zero latest（go mod tidy 锁定）、goctl 1.10.2、MySQL 8 |

## 目录结构

```
poppulse-gozero/
├── go.mod                    # module poppulse
├── poppulse.go               # API 入口（goctl 生成 + 统一响应注册）
├── poppulse.api              # 接口契约 v1：39 个 MVP 接口
├── etc/poppulse.example.yaml # 模板配置（提交）；poppulse.yaml 本地（gitignore）
├── internal/
│   ├── config/config.go      # MySQL/Auth/Mock/PayTimeout
│   ├── middleware/operator.go# operator 角色校验（user 级用 go-zero jwt）
│   ├── xerr/                 # 业务错误码常量
│   ├── svc/servicecontext.go # sqlx.Conn、Config
│   ├── handler/ logic/ types/# goctl 生成，group: health/auth/user/product/stock/order/payment/aftersale/admin*
│   └── model/                # goctl model 生成（表定稿后随模块生成）
├── deployments/
│   ├── docker-compose.yml    # MySQL 8
│   └── migrations/001_init.sql  # 计划书 14 张表
├── scripts/seed.sql          # 种子：演示用户+运营（bcrypt）+≥3 商品
├── docs/plans/               # 本文档
└── README.md                 # 版本、启动、迁移、种子账号、关闭模拟
```

## 关键约定

- 统一响应 `{code, message, data, trace_id}`：`httpx.SetOkHandler`/`SetErrorHandler` 全局包装
- 错误码：0 成功；40xxx 参数/权限/不存在；30xxx 业务（如 30001 库存不足）；50101 未实现占位
- 鉴权三级：public 无中间件；user 路由 `jwt: Auth`；admin 路由 `jwt: Auth` + Operator 中间件 + logic 层数据归属校验
- 金额整数分；ID/单号 JSON 用字符串；时间 RFC3339 带时区
- 分页默认 page=1、page_size=20、最大 100
- 幂等：订单创建携带 `Idempotency-Key` 请求头
- 模拟支付/退款由配置 `Mock.Enabled` 显式开启，默认关闭

## D1 验收门槛（对应七天计划书）

1. 全新环境 `docker compose up -d` → 迁移 + 种子 → `go run poppulse.go` 启动成功
2. `GET /api/v1/health/live` 返回统一响应 200
3. 种子账号 `POST /api/v1/auth/login` 获真实 JWT，`GET /api/v1/users/me` 鉴权通过
4. 39 接口路由与权限归属固化，未实现接口返回 50101 占位错误
5. 非部署人按 README 独立完成启动

## Git 协作

- `main` 始终可启动；组员分支：`feature/user`、`feature/product`、`feature/stock`、`feature/order`、`feature/payment`、`feature/aftersale`
- 迁移脚本只追加不改已合并版本；公共 `poppulse.api` 修改须经组长（A）
- 提交前 gofmt + go build；每人合并至少一人复核

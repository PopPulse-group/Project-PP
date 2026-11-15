# 生鲜零售平台 组员上手指南（ONBOARDING）

> 拉取项目、跑通环境、开始开发的最短路径。阻塞超过 30 分钟先联系评审搭档，再群里 @组长。

## 0. 前置条件（组长操作）

1. `github.com/PopPulse-group` → **People → Invite member** 邀请组员
2. 仓库 **Settings → Collaborators** 给每人 **Write** 权限
3. 组员确认能打开 `https://github.com/PopPulse-group/Project-PP`

注意：本仓库此前误建过 PopPulse 潮玩电商，旧代码已归档到 tag `v0.1-poppulse-archived`，main 上现在是生鲜零售平台 MVP，请以 README 和本文档为准。

## 1. 环境准备（每人本机）

| 工具 | 版本 | 用途 |
|---|---|---|
| Git | 最新版 | 拉代码、提交 |
| Go | 1.25+ | 编译运行（go.mod 锁定 go-zero v1.10.3，勿升级） |
| Docker Desktop | 最新版 | 起 MySQL + Redis |
| Apifox 或 Postman | 任选一 | 联调接口 |

可选：goctl 1.10.2（重新生成代码时才需要）。

## 2. 克隆（首次）

```bash
cd ~/Desktop   # 或你的工作目录（不要在外层再 git init！）
git clone https://github.com/PopPulse-group/Project-PP.git
cd Project-PP  # 必须进入克隆出来的目录再操作
git remote -v  # 应显示 origin → Project-PP.git
```

## 3. 启动服务

```powershell
docker compose -f deployments/docker-compose.yml up -d
docker exec -i fresh-mysql mysql -uroot -ppoppulse123 --default-character-set=utf8mb4 < deployments/migrations/001_init.sql
docker exec -i fresh-mysql mysql -uroot -ppoppulse123 --default-character-set=utf8mb4 < scripts/seed.sql
Copy-Item etc/fresh.example.yaml etc/fresh.yaml   # 仅第一次
go run fresh.go
```

验证：`curl http://localhost:8888/api/v1/health/ready` 返回 `code:0`（检查 MySQL+Redis）。

## 4. 分支与模块

| 成员 | 分支 | 模块 |
|---|---|---|
| A | `feature/account` | 注册登录、JWT、地址、集成 |
| B | `feature/product` | 分类/商品/购物车、缓存与并发聚合 |
| C | `feature/trade` | 订单事务、防超卖、幂等、模拟支付 |
| D | `feature/delivery` | 配送事件、超时扫描 worker pool |
| E | `feature/ai` | AI 适配器接入、限流降级 |
| F | `feature/qa` | 契约、集成测试、压测 |

```bash
git checkout main && git pull
git checkout -b feature/your-module
```

## 5. 日常循环

```bash
gofmt -w . && go build ./... && go vet ./...   # 提交前必过
git add . && git commit -m "trade: 实现幂等下单事务"
git push -u origin feature/trade                # 每日小步 PR
git checkout main && git pull && git checkout feature/trade && git rebase main   # 每天同步
```

## 6. 代码写在哪

- 业务逻辑：`internal/logic/<分组>/`（未实现返回 50101 占位）
- 分组：`auth`、`user`、`product`、`cart`、`order`、`ai`、`health`
- 数据访问：`internal/model/`（统一 sqlx）
- AI：`internal/ai/`（只改 Provider 实现，不动接口定义）
- 公共文件（`fresh.api`、config、xerr、response、middleware）修改先找组长 A

## 7. 硬性规则（违反 = 返工）

1. 防超卖：条件更新 + 影响行数检查，多商品按 id 排序，失败整体回滚
2. 幂等：`Idempotency-Key` 必校验；同键同参原结果、同键异参 409/40901
3. 服务端计价，禁止采信客户端金额；金额整数分；ID 输出字符串
4. 状态机：pending_payment→paid→delivering→delivered；逆向/跳跃拒绝；取消只能从 pending_payment
5. 越权：非本人资源 404；所有查询按 user_id 过滤
6. AI：只读商品、不虚构 SKU/价格/库存、mock 字段必须真实
7. 迁移只追加（`002_xxx.sql`），禁止改 `001_init.sql`
8. 模拟支付/履约仅 `Mock.Enabled: true` 可用

## 8. 验收硬指标

库存 100 / 200 并发抢购恰成 100 单；同键并发 50 次一笔订单；支付取消竞争单一终态；超时任务幂等；AI 不虚构数据。这些是 D8-D11 并发测试与演示的硬门槛，写 logic 时就要按此设计。

## 9. 常见问题

| 现象 | 解决 |
|---|---|
| 3306 连接拒绝 | `docker ps` 看 fresh-mysql 状态，重新 `compose up -d` |
| 503/50301 | MySQL 或 Redis 未就绪，等容器 healthy |
| 中文乱码 | 导入 SQL 未带 `--default-character-set=utf8mb4`，DROP 库重导 |
| 501/50101 | 接口未实现：是你的模块就去写 |
| push 403 | 找组长确认 Write 权限 |
| push 被拒 fetch first | `git pull --rebase origin main` |
| 克隆后在错误目录 git init | 删除外层 `.git`，`cd Project-PP` 后再操作 |

## 10. 合并流程

每日小步 PR → 搭档当天评审 → 组长合并 → 合并人验证 main 可启动。契约（fresh.api）在 D10 冻结 v1。
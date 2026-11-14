# PopPulse 组员上手指南（ONBOARDING）

> 本文是组员拉取项目、跑通环境、开始开发的最短路径。遇到问题先看文末「常见问题」，解决不了直接群里 @组长。

## 0. 前置条件（组长操作）

私有仓库必须先有权限才能拉取：

1. 组长在 `github.com/PopPulse-group` → **People → Invite member** 邀请组员（GitHub 用户名或邮箱）
2. 组员接受邀请后，组长在仓库 **Settings → Collaborators** 给每人 **Write** 权限
3. 组员确认自己能打开 `https://github.com/PopPulse-group/Project-PP`（能看到就说明权限 OK）

## 1. 环境准备（每人本机）

| 工具 | 版本要求 | 用途 |
|---|---|---|
| Git | 最新版 | 拉代码、提交 |
| Go | 1.25+ | 编译运行（go.mod 锁定 go-zero v1.10.3，勿随意升级） |
| Docker Desktop | 最新版 | 起 MySQL（Windows 需启用 WSL2） |
| Apifox 或 Postman | 任选其一 | 联调接口 |

可选：goctl 1.10.2（只有需要重新生成 handler/logic/types 时才装）。

## 2. 克隆代码（首次）

```powershell
# 选好存放目录，例如 D:\work
cd D:\work

git clone https://github.com/PopPulse-group/Project-PP.git
cd Project-PP
```

> 首次克隆 HTTPS 会弹出 GitHub 登录授权，在浏览器点一下授权即可；之后 Git Credential Manager 会记住，无需重复登录。

## 3. 启动服务（每次新人/新环境必做）

```powershell
# 1. 启动 MySQL
docker compose -f deployments/docker-compose.yml up -d

# 2. 初始化表结构 + 种子数据（注意必须带 --default-character-set=utf8mb4，否则中文乱码）
docker exec -i poppulse-mysql mysql -uroot -ppoppulse123 --default-character-set=utf8mb4 < deployments/migrations/001_init.sql
docker exec -i poppulse-mysql mysql -uroot -ppoppulse123 --default-character-set=utf8mb4 < scripts/seed.sql

# 3. 准备本地配置（仅第一次；poppulse.yaml 不提交仓库）
Copy-Item etc/poppulse.example.yaml etc/poppulse.yaml

# 4. 启动 API 服务
go run poppulse.go
```

验证（另开一个终端）：

```powershell
curl http://localhost:8888/api/v1/health/live
# 期望：{"code":0,"message":"success","data":{"status":"ok",...},"trace_id":""}
```

### 种子账号

| 账号 | 密码 | 角色 |
|---|---|---|
| user01 | User@123456 | user（普通用户） |
| admin01 | Admin@123456 | operator（运营） |

登录测试：

```powershell
curl -X POST http://localhost:8888/api/v1/auth/login -H "Content-Type: application/json" -d '{"username":"user01","password":"User@123456"}'
```

## 4. 切开发分支（不要直接改 main）

分支命名与分工对照：

| 成员 | 分支 | 模块 |
|---|---|---|
| B | `feature/user` | 登录、用户资料、JWT |
| C | `feature/product` | IP/系列/SKU、商品查询与管理 |
| D | `feature/stock` | 库存预占/释放/流水 |
| E | `feature/order` | 订单、模拟支付 |
| F | `feature/aftersale` | 售后申请/审核/模拟退款 |

```powershell
# 以库存模块 D 为例
git checkout main
git pull                      # 先同步最新 main
git checkout -b feature/stock # 切出自己的分支
```

## 5. 日常开发循环

```powershell
# 写代码 → 提交前检查
gofmt -w .
go build ./...
go vet ./...

git add .
git commit -m "stock: 实现库存预占与释放事务"   # 格式：<模块>: <做了什么>
git push -u origin feature/stock                # 首次推送；以后直接 git push
```

每天开始工作前同步 main 的更新：

```powershell
git checkout main
git pull
git checkout feature/stock
git rebase main
```

### 你的代码写在哪里

- 业务逻辑：`internal/logic/<你的模块分组>/`（M01–M39 每个接口一个文件，未实现的目前返回 50101 占位）
- 数据访问：`internal/model/`（统一 sqlx，一套方式不混用）
- 路由/契约：`poppulse.api`（**公共文件，修改必须先找组长 A**）
- 错误码：`internal/xerr/`（也是公共文件）

各模块分组目录：`auth`、`user`、`product`、`adminProduct`、`adminStock`、`order`、`payment`、`adminOrder`、`aftersale`、`adminAftersale`。

## 6. 硬性规则（违反 = 返工）

1. **库存防超卖**：`UPDATE ... WHERE total - reserved - sold >= ?` 条件更新 + 检查影响行数，禁止先查数量再无条件扣减
2. **幂等**：创建订单必须校验 `Idempotency-Key` 请求头；同键同参返回原结果、同键异参返回 409/40002
3. **状态机**：非法状态转换拒绝；取消与支付竞争只有一种合法终态
4. **金额**：一律整数分（bigint），禁止浮点数
5. **ID**：JSON 输出一律字符串（避免前端大整数精度问题）
6. **越权**：非本人资源统一返回 404；admin 接口已有 operator 中间件，logic 里仍须校验数据归属
7. **迁移**：只新增 `002_xxx.sql`、`003_xxx.sql`… 追加，禁止修改 `001_init.sql`
8. **模拟能力**：模拟支付/退款仅当配置 `Mock.Enabled: true` 可用，生产必须 false

## 7. 接口契约

- 39 个 MVP 接口定义在 `poppulse.api`（对应接口文档 M01–M39）
- 统一响应：`{"code":0,"message":"success","data":{...},"trace_id":"..."}`
- 未实现接口调用返回 `501 接口尚未实现（D1 占位）`
- 用 Apifox/Postman 联调时带 `Authorization: Bearer <access_token>`，admin 接口用 operator 账号登录获取 token

## 8. 常见问题

| 现象 | 原因与解决 |
|---|---|
| 启动报 `dial tcp 127.0.0.1:3306 connect refused` | MySQL 没起来：`docker ps` 看容器状态，重新 `docker compose up -d` |
| 接口返回 `503 code=90002` | MySQL 未就绪或没跑迁移，等容器 healthy 后重试 |
| 登录返回 401/40003 | 用种子账号，密码见第 3 节 |
| 数据库中文乱码 | 导入 SQL 时没带 `--default-character-set=utf8mb4`，DROP 库重新导入 |
| 接口返回 501 code=50101 | 正常：该接口还没实现，是你模块的就去做，不是你的等负责人 |
| push 提示 403 无权限 | 组长还没给你 Write 权限，找组长确认邀请 |
| push 被拒 `fetch first` | 远程有新提交：`git pull --rebase origin main` 后再推 |

## 9. 合并流程

1. 开发完成 → push 自己的分支到远程
2. 在 GitHub 上从 `feature/xxx` 向 `main` 发起 **Pull Request**，描述做了什么、怎么测的
3. 至少找一名组员（建议对口联调的模块负责人）复核代码
4. 组长合并；合并后 `main` 必须仍可启动（合并人负责验证）

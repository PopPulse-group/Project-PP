# 生鲜零售平台 MVP 框架修订设计

日期：2026-11-15　状态：已确认（组长批准"执行"）
前置决策：范围以两周计划书 MVP 为准（api.md 全量设计归档 `docs/future/` 作二期参考）；金额整数分（bigint）；账号密码登录（无短信）。

## 复用与变更总览

| 项 | 处理 |
|---|---|
| go-zero v1.10.3 / goctl / 统一响应 / 错误码体系 / operator 中间件 / gitignore / Compose | 复用 |
| PopPulse 14 表迁移、售后/退款/管理域接口 | 删除（Git tag `v0.1-poppulse-archived` 可找回） |
| `poppulse.api` 39 接口 | 替换为 `fresh.api` MVP 契约（18 接口） |
| 旧 001_init.sql | 重写为生鲜 8 表 v1（上线前可重置；上线后恢复"只追加"规则） |

## MVP 契约（fresh.api，18 接口）

公共：GET /health/live、GET /health/ready
账号：POST /auth/register、POST /auth/login、GET /users/me
地址：GET /users/me/addresses、POST /users/me/addresses、DELETE /users/me/addresses/:address_id
商品：GET /categories、GET /products（category_id/keyword/分页）、GET /products/:product_id
购物车：GET /cart、POST /cart/items、PATCH /cart/items/:cart_item_id、DELETE /cart/items/:cart_item_id
订单：POST /orders（Idempotency-Key 服务端计价）、GET /orders、GET /orders/:id、POST /orders/:id/cancel、POST /orders/:id/mock-pay
履约：GET /orders/:id/delivery（模拟配送事件流）
AI：POST /ai/meal-plans（人数/预算分/口味/忌口 → 结构化菜谱+在售 SKU 匹配+价格汇总）、POST /ai/meal-plans/confirm（确认加购，幂等合并）

权限：public = health/auth/商品浏览；其余 jwt；admin 入口保留 operator 中间件（模拟履约推进用，本周仅预留）。

## 数据库 v1（8 表）

1. users：id、username(唯一)、password_hash、nickname、status、时间戳
2. addresses：id、user_id、contact_name、mobile、province、city、detail、is_default
3. categories：id、name、sort_no、status
4. products：id、category_id、name、description、unit、price(分)、stock、status(on_sale/off_sale)；**单库存池，防超卖条件更新 `stock=stock-? WHERE id=? AND stock>=?`**
5. cart_items：id、user_id、product_id、quantity，UNIQUE(user_id, product_id)（重复加购合并）
6. orders：id、user_id、status(pending_payment/paid/delivering/delivered/cancelled)、total_amount、idempotency_key、UNIQUE(user_id,idempotency_key)、expire_at、paid_at、cancelled_at
7. order_items：id、order_id、product_id、product_name、unit_price、quantity、amount（价格快照）
8. delivery_events：id、order_id、event_type(paid/sorting/delivering/delivered/timed_out)、description、occurred_at

订单状态机：pending_payment →(mock-pay)→ paid →(模拟履约)→ delivering → delivered；pending_payment →(取消/超时)→ cancelled。逆向/跳跃转换拒绝。

## 新增组件

- Redis：Compose 增加 redis:7，配置 `Redis.Host`；本周仅 B 的商品缓存预热使用，不承载库存事实
- AI 适配器：`internal/ai/provider.go` 定义 `Provider` 接口 + `MockProvider`（固定菜谱模板）；`internal/config` 增加 `Ai.Provider = mock|real`、`Ai.TimeoutMs`、`Ai.MaxConcurrent`；真实模型接入后实现同一接口
- 错误码新增：40902 库存不足(409)、40903 订单状态不允许(409)、42210 AI 输出校验失败(422)、42211 食材无在售匹配(422)、42901 限流(429)、50301 高德/外部依赖降级(503)
- 超时释放：worker pool 属 D 成员任务，本周仅预留 `internal/worker/` 目录与设计注释

## 种子数据

2 个演示账号（复用 user01/admin01 密码哈希）、5 个分类、30 个生鲜商品（含库存），SQL 随迁移可重放。

## 分支重规划

A `feature/account`、B `feature/product`（含购物车+缓存）、C `feature/trade`（订单+库存+支付）、D `feature/delivery`（履约+worker pool）、E `feature/ai`、F `feature/qa`（契约+测试+压测）。

## 验收门槛（D1-D3）

空库迁移+种子 → 启动 → 注册→登录→浏览商品→加购→幂等下单→模拟支付→查询配送 全链路可复现；AI Mock 场景 3 个输入可复现输出；未实现接口 50101。

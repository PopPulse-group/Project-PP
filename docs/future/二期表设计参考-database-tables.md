# 生鲜即时零售平台 MySQL 数据表设计

## 设计约定

- MySQL 8.0，`BIGINT` 主键，时间字段使用 `DATETIME(3)`，金额使用 `DECIMAL(10,2)`。
- 所有业务表建议包含 `id`、`created_at`、`updated_at`；软删除表增加 `deleted_at`。
- `tenant_id` 用于多商家/多前置仓隔离；订单默认由一家商家（一个履约仓）完成。
- 密码、支付密钥、身份证号等敏感数据只保存加密或脱敏结果。
- 表名使用小写下划线；状态字段采用后端常量，禁止依赖中文文本。

## 1. 用户与权限模块

| 表 | 用途 | 关键字段 |
|---|---|---|
| `users` | 用户主表 | `id`, `mobile_cipher`, `nickname`, `avatar_url`, `status`, `last_login_at` |
| `user_identities` | 微信/手机号等登录身份 | `id`, `user_id`, `provider`, `provider_uid`, `union_id`，唯一键(`provider`,`provider_uid`) |
| `user_addresses` | 收货地址及高德坐标 | `id`, `user_id`, `contact_name`, `mobile_cipher`, `province`, `city`, `detail`, `longitude`, `latitude`, `amap_poi_id`, `is_default` |
| `user_preferences` | 口味、忌口、过敏信息 | `user_id`, `spicy_level`, `diet_tags_json`, `allergy_tags_json`, `budget_level` |
| `user_favorites` | 收藏商品/菜谱/商家 | `id`, `user_id`, `target_type`, `target_id` |
| `user_memberships` | 会员等级与有效期 | `id`, `user_id`, `level`, `points`, `started_at`, `expired_at` |
| `user_points_ledger` | 积分流水 | `id`, `user_id`, `change_amount`, `balance_after`, `source_type`, `source_id` |
| `login_verification_codes` | 短信验证码 | `id`, `mobile_hash`, `code_hash`, `purpose`, `expired_at`, `used_at` |
| `roles` / `permissions` / `role_permissions` / `user_roles` | RBAC 权限 | 角色、权限及关联关系 |

## 2. 商家、仓库与骑手模块

| 表 | 用途 | 关键字段 |
|---|---|---|
| `merchants` | 商家主表 | `id`, `name`, `license_no_cipher`, `status`, `service_type` |
| `merchant_users` | 商家员工账号 | `id`, `merchant_id`, `user_id`, `job_type`, `status` |
| `merchant_business_hours` | 营业时间 | `id`, `merchant_id`, `weekday`, `open_time`, `close_time` |
| `warehouses` | 前置仓 | `id`, `merchant_id`, `name`, `address`, `longitude`, `latitude`, `amap_poi_id`, `status`, `capacity` |
| `warehouse_service_areas` | 仓库配送范围 | `id`, `warehouse_id`, `polygon_json`, `radius_m`, `priority` |
| `warehouse_staff` | 仓内人员 | `id`, `warehouse_id`, `user_id`, `role`, `status` |
| `riders` | 骑手档案 | `id`, `user_id`, `real_name_cipher`, `mobile_cipher`, `status`, `online_status` |
| `rider_devices` | 骑手定位设备 | `id`, `rider_id`, `device_id`, `platform`, `push_token` |
| `rider_location_points` | 定位轨迹（建议按日分区） | `id`, `rider_id`, `order_id`, `longitude`, `latitude`, `accuracy_m`, `reported_at` |

## 3. 商品、菜谱与库存模块

| 表 | 用途 | 关键字段 |
|---|---|---|
| `product_categories` | 商品分类 | `id`, `parent_id`, `name`, `sort_no`, `status` |
| `products` | SPU 商品 | `id`, `merchant_id`, `category_id`, `name`, `origin`, `brand`, `status` |
| `product_skus` | SKU 规格 | `id`, `product_id`, `sku_code`, `spec_json`, `unit`, `sale_price`, `cost_price`, `weight`, `status` |
| `product_images` | 商品图片 | `id`, `product_id`, `sku_id`, `url`, `sort_no` |
| `warehouse_sku_prices` | 仓库差异价格 | `id`, `warehouse_id`, `sku_id`, `sale_price`, `member_price` |
| `warehouse_sku_inventory` | 当前库存快照 | `warehouse_id`, `sku_id`, `on_hand`, `reserved`, `available`, `version` |
| `inventory_transactions` | 入库、预占、扣减、释放流水 | `id`, `warehouse_id`, `sku_id`, `biz_type`, `biz_id`, `quantity`, `before_qty`, `after_qty` |
| `inventory_batches` | 批次与保质期 | `id`, `warehouse_id`, `sku_id`, `batch_no`, `production_at`, `expired_at`, `quantity` |
| `inventory_loss_records` | 报损/报溢 | `id`, `warehouse_id`, `sku_id`, `batch_id`, `quantity`, `reason`, `operator_id` |
| `recipes` | 菜谱 | `id`, `creator_type`, `title`, `cuisine`, `difficulty`, `cook_time_min`, `content_json`, `status` |
| `recipe_ingredients` | 菜谱食材与替代品 | `id`, `recipe_id`, `sku_id`, `quantity`, `unit`, `is_optional`, `substitute_group` |
| `recipe_tags` / `recipe_tag_relations` | 口味、忌口标签 | 标签及菜谱关系 |
| `product_reviews` | 商品评价 | `id`, `user_id`, `order_id`, `sku_id`, `rating`, `content`, `status` |

## 4. 供应链与采购模块

| 表 | 用途 | 关键字段 |
|---|---|---|
| `suppliers` | 供应商 | `id`, `name`, `license_no_cipher`, `contact_cipher`, `status` |
| `supplier_products` | 供应商供货关系 | `id`, `supplier_id`, `sku_id`, `purchase_price`, `lead_time_hours` |
| `purchase_orders` | 采购单 | `id`, `supplier_id`, `warehouse_id`, `order_no`, `status`, `total_amount`, `expected_arrival_at` |
| `purchase_order_items` | 采购明细 | `id`, `purchase_order_id`, `sku_id`, `planned_qty`, `received_qty`, `purchase_price` |
| `goods_receipts` | 到货验收单 | `id`, `purchase_order_id`, `warehouse_id`, `status`, `received_by`, `received_at` |
| `goods_receipt_items` | 验收明细 | `id`, `receipt_id`, `sku_id`, `batch_id`, `accepted_qty`, `rejected_qty`, `quality_result` |
| `stock_transfers` / `stock_transfer_items` | 仓间调拨及明细 | 调出仓、调入仓、在途数量、签收数量 |

## 5. 购物车、订单与支付模块

| 表 | 用途 | 关键字段 |
|---|---|---|
| `carts` | 用户购物车 | `id`, `user_id`, `warehouse_id`, `status` |
| `cart_items` | 购物车明细 | `id`, `cart_id`, `sku_id`, `quantity`, `selected` |
| `orders` | 订单主表 | `id`, `order_no`, `user_id`, `merchant_id`, `warehouse_id`, `address_snapshot_json`, `status`, `payment_status`, `subtotal_amount`, `discount_amount`, `delivery_fee`, `payable_amount`, `idempotency_key` |
| `order_items` | 订单商品快照 | `id`, `order_id`, `sku_id`, `product_name_snapshot`, `spec_snapshot`, `unit_price`, `quantity`, `refund_qty` |
| `order_status_logs` | 订单状态历史 | `id`, `order_id`, `from_status`, `to_status`, `operator_type`, `operator_id`, `remark` |
| `order_remarks` | 用户/商家备注 | `id`, `order_id`, `content`, `source` |
| `payments` | 支付单 | `id`, `payment_no`, `order_id`, `channel`, `amount`, `status`, `third_party_trade_no`, `paid_at` |
| `refunds` | 退款单 | `id`, `refund_no`, `order_id`, `payment_id`, `amount`, `reason`, `status`, `refunded_at` |
| `coupons` / `user_coupons` / `coupon_usages` | 优惠券 | 券模板、用户领取、订单核销 |
| `promotions` / `promotion_skus` | 秒杀、满减活动 | 活动规则、商品范围、时间和库存 |

## 6. 配送履约模块（高德）

| 表 | 用途 | 关键字段 |
|---|---|---|
| `delivery_tasks` | 配送任务 | `id`, `task_no`, `order_id`, `warehouse_id`, `rider_id`, `status`, `pickup_at`, `delivered_at`, `eta_at` |
| `delivery_task_events` | 配送状态事件 | `id`, `task_id`, `event_type`, `operator_id`, `occurred_at`, `payload_json` |
| `delivery_routes` | 高德路线结果快照 | `id`, `task_id`, `route_type`, `distance_m`, `duration_sec`, `amap_request_id`, `route_json` |
| `delivery_location_snapshots` | 对外展示的最新位置 | `task_id`, `rider_id`, `longitude`, `latitude`, `accuracy_m`, `reported_at` |
| `delivery_exceptions` | 延误、联系不上、地址异常 | `id`, `task_id`, `type`, `description`, `status`, `resolved_by` |
| `amap_api_logs` | 高德调用审计 | `id`, `api_name`, `request_id`, `biz_id`, `http_status`, `latency_ms`, `error_code` |

## 7. AI 智能助手模块

| 表 | 用途 | 关键字段 |
|---|---|---|
| `ai_conversations` | AI 会话 | `id`, `user_id`, `scene`, `title`, `status`, `last_message_at` |
| `ai_messages` | 对话消息 | `id`, `conversation_id`, `role`, `content`, `content_json`, `token_count`, `created_at` |
| `ai_recommendation_sessions` | 一次推荐过程 | `id`, `conversation_id`, `address_id`, `budget`, `people_count`, `diet_json`, `warehouse_id`, `status` |
| `ai_recommendation_items` | 推荐菜谱及商品 | `id`, `session_id`, `recipe_id`, `sku_id`, `quantity`, `reason`, `availability_status` |
| `ai_tool_calls` | AI 工具调用审计 | `id`, `message_id`, `tool_name`, `arguments_json`, `result_json`, `status`, `latency_ms` |
| `ai_feedbacks` | 推荐/回答反馈 | `id`, `user_id`, `message_id`, `rating`, `reason` |
| `ai_knowledge_documents` | 菜谱、售后、商品知识 | `id`, `doc_type`, `source_id`, `content`, `version`, `status` |

## 8. 售后、通知与平台支撑

| 表 | 用途 | 关键字段 |
|---|---|---|
| `after_sales` / `after_sale_items` | 售后单及商品明细 | 售后类型、凭证、审核状态、退款金额 |
| `customer_service_tickets` | 客服工单 | 用户、订单、问题类型、处理人、状态 |
| `notifications` | 站内/短信/推送消息 | `id`, `user_id`, `channel`, `template_code`, `payload_json`, `status` |
| `outbox_events` | 可靠事件投递 | `id`, `aggregate_type`, `aggregate_id`, `event_type`, `payload_json`, `status`, `retry_count` |
| `idempotency_keys` | 接口幂等 | `id`, `scope`, `key`, `request_hash`, `response_json`, `expired_at` |
| `operation_audits` | 管理操作审计 | 操作人、资源、动作、前后值、IP |
| `api_request_logs` | API 访问日志 | 请求ID、用户、路径、状态码、耗时 |

## 关键索引与约束

- `orders(order_no)`、`orders(user_id, created_at)`、`orders(status, created_at)`、`order_items(order_id)`。
- `warehouse_sku_inventory(warehouse_id, sku_id)` 唯一；库存更新必须带 `version` 或使用条件更新。
- `delivery_tasks(order_id)` 唯一；`delivery_location_snapshots(task_id)` 唯一。
- `ai_messages(conversation_id, created_at)`、`ai_tool_calls(message_id)`。
- 位置点、API 日志、订单状态日志按月/日归档或分区；禁止在列表接口直接查询全量轨迹。

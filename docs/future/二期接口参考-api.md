# 生鲜即时零售平台接口文档

基础地址：`/api/v1`。认证使用 `Authorization: Bearer <token>`；写接口建议携带 `X-Request-Id` 和 `Idempotency-Key`。金额单位为元，坐标使用高德 GCJ-02。

## 统一响应

```json
{"code":0,"message":"ok","data":{},"request_id":"req_01"}
```

## 用户与地址

### `POST /auth/sms/send`

```json
{"mobile":"13800000000","purpose":"login"}
```

### `POST /auth/sms/login`

```json
{"mobile":"13800000000","code":"123456"}
```

### `GET /users/me/addresses`

### `POST /users/me/addresses`

```json
{"contact_name":"张三","mobile":"13800000000","province":"上海市","city":"上海市","detail":"徐汇区XX路1号","longitude":121.437,"latitude":31.188,"amap_poi_id":"B0...","is_default":true}
```

### `POST /amap/geocode/regeo`

```json
{"longitude":121.437,"latitude":31.188}
```

## 商品、商家与库存

### `GET /merchants/nearby?longitude=121.437&latitude=31.188`

### `GET /warehouses/{warehouse_id}/products?category_id=1&keyword=番茄&page=1&page_size=20`

### `GET /products/{sku_id}`

返回商品、规格、价格、`available_stock`、评价摘要。

## AI 对话式导购

### `POST /ai/conversations`

```json
{"scene":"shopping_assistant","address_id":1001}
```

### `POST /ai/conversations/{conversation_id}/messages`

```json
{"content":"两个人，预算80元，想吃川菜但不要太辣，推荐有货的菜和食材","stream":true}
```

AI 可调用的受控工具：`search_available_recipes`、`check_sku_inventory`、`calculate_cart`、`get_order_status`、`get_delivery_tracking`、`create_after_sale`。工具调用必须经过用户权限和业务服务校验。

### `POST /ai/recommendations/{session_id}/confirm`

```json
{"selected_recipe_ids":[301],"items":[{"sku_id":9001,"quantity":1},{"sku_id":9002,"quantity":2}],"address_id":1001}
```

响应返回确认后的购物车和价格，不直接扣款。

## 购物车与订单

### `GET /carts/current?warehouse_id=10`

### `POST /carts/items`

```json
{"warehouse_id":10,"sku_id":9001,"quantity":2}
```

### `POST /orders/preview`

```json
{"warehouse_id":10,"address_id":1001,"items":[{"sku_id":9001,"quantity":2}],"coupon_id":5001,"delivery_type":"instant"}
```

### `POST /orders`

```json
{"preview_token":"pv_xxx","address_id":1001,"remark":"放门口","delivery_type":"instant"}
```

成功后返回 `order_id`、`order_no`、`payable_amount`、`payment_expire_at`。必须使用 `Idempotency-Key`，重复请求返回同一订单。

### `GET /orders?status=配送中&page=1&page_size=20`

### `GET /orders/{order_id}`

### `POST /orders/{order_id}/cancel`

### `POST /orders/{order_id}/confirm-receipt`

## 支付与售后

### `POST /payments`

```json
{"order_id":7001,"channel":"wechat"}
```

### `POST /payments/{payment_id}/callback`

支付回调必须验签、幂等，并通过状态机更新订单。

### `POST /orders/{order_id}/after-sales`

```json
{"type":"missing_item","items":[{"order_item_id":8001,"quantity":1}],"reason":"少了一件","evidence_urls":[]}
```

## 商家端履约

### `GET /merchant/orders?status=待接单`

### `POST /merchant/orders/{order_id}/accept`

### `POST /merchant/orders/{order_id}/packing`

```json
{"short_items":[{"sku_id":9002,"quantity":1,"action":"refund"}],"packed_at":"2026-11-14T10:20:00+08:00"}
```

### `POST /merchant/orders/{order_id}/ready`

商家备货完成后生成或激活配送任务。

## 骑手端与高德配送

### `POST /delivery/tasks/{task_id}/accept`

### `POST /delivery/tasks/{task_id}/pickup`

### `POST /delivery/tasks/{task_id}/location`

```json
{"longitude":121.438,"latitude":31.189,"accuracy_m":8,"reported_at":"2026-11-14T10:25:00+08:00"}
```

### `GET /delivery/tasks/{task_id}/route?route_type=to_user`

服务端调用高德路线规划并返回距离、预计时长和折线；对外隐藏高德密钥。

### `POST /delivery/tasks/{task_id}/deliver`

```json
{"proof_url":"https://cdn.example.com/proof.jpg","delivered_at":"2026-11-14T10:45:00+08:00"}
```

### `GET /orders/{order_id}/tracking`

```json
{"order_status":"配送中","task_status":"配送中","rider":{"name":"李师傅"},"latest_location":{"longitude":121.438,"latitude":31.189},"route":{"distance_m":1800,"duration_sec":520},"eta_at":"2026-11-14T10:45:00+08:00"}
```

## WebSocket / SSE

- `GET /ws/orders/{order_id}`：用户订阅订单状态、骑手位置、ETA；服务端鉴权后推送 `order.status_changed`、`delivery.location_updated`。
- `GET /ai/conversations/{id}/stream`：SSE 流式输出 AI 文本和工具调用事件。

## 状态机

订单：`PENDING_PAYMENT → PAID → MERCHANT_ACCEPTED → PACKING → READY_FOR_PICKUP → DELIVERING → DELIVERED`；异常分支为 `CANCELLED`、`REFUNDING`、`REFUNDED`。

配送：`WAITING_ASSIGNMENT → ASSIGNED → RIDER_ACCEPTED → PICKED_UP → DELIVERING → DELIVERED`；异常为 `EXCEPTION`。

## 错误码

| code | 含义 |
|---:|---|
| 0 | 成功 |
| 40001 | 参数错误 |
| 40101 | 未登录或 token 失效 |
| 40301 | 无权限 |
| 40901 | 幂等键重复但请求内容不同 |
| 40902 | 库存不足 |
| 40903 | 订单状态不允许此操作 |
| 42901 | 请求或秒杀限流 |
| 50001 | 业务服务暂时不可用 |
| 50301 | 高德服务降级，稍后重试 |

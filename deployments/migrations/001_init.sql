-- 生鲜零售平台 MVP 表结构 v1（两周计划书口径，D2 定稿）
-- 规则：上线后迁移脚本只追加，不修改已合并版本
-- 金额统一整数分；ID 为 MySQL 自增 BIGINT，JSON 输出统一字符串

CREATE DATABASE IF NOT EXISTS fresh DEFAULT CHARSET utf8mb4 COLLATE utf8mb4_unicode_ci;
USE fresh;

CREATE TABLE IF NOT EXISTS users (
  id            BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  username      VARCHAR(32)  NOT NULL,
  password_hash VARCHAR(255) NOT NULL COMMENT 'bcrypt 哈希',
  nickname      VARCHAR(64)  NOT NULL DEFAULT '',
  status        ENUM('active','disabled') NOT NULL DEFAULT 'active',
  created_at    DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
  updated_at    DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
  UNIQUE KEY uk_users_username (username)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户';

CREATE TABLE IF NOT EXISTS addresses (
  id           BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  user_id      BIGINT UNSIGNED NOT NULL,
  contact_name VARCHAR(64) NOT NULL,
  mobile       VARCHAR(32) NOT NULL,
  province     VARCHAR(64) NOT NULL DEFAULT '',
  city         VARCHAR(64) NOT NULL DEFAULT '',
  detail       VARCHAR(512) NOT NULL,
  is_default   TINYINT(1) NOT NULL DEFAULT 0,
  created_at   DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
  updated_at   DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
  KEY idx_addresses_user (user_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='收货地址';

CREATE TABLE IF NOT EXISTS categories (
  id         BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  name       VARCHAR(64) NOT NULL,
  sort_no    INT NOT NULL DEFAULT 0,
  status     ENUM('active','inactive') NOT NULL DEFAULT 'active',
  created_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
  updated_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='商品分类';

CREATE TABLE IF NOT EXISTS products (
  id          BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  category_id BIGINT UNSIGNED NOT NULL,
  name        VARCHAR(128) NOT NULL,
  description VARCHAR(1024) NOT NULL DEFAULT '',
  unit        VARCHAR(32) NOT NULL DEFAULT '份' COMMENT '规格单位，如 500g/份',
  price       BIGINT NOT NULL COMMENT '售价，整数分',
  stock       INT NOT NULL DEFAULT 0 COMMENT '可售库存（单库存池）',
  status      ENUM('on_sale','off_sale') NOT NULL DEFAULT 'off_sale',
  created_at  DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
  updated_at  DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
  KEY idx_products_category (category_id),
  KEY idx_products_status (status)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='商品';

CREATE TABLE IF NOT EXISTS cart_items (
  id         BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  user_id    BIGINT UNSIGNED NOT NULL,
  product_id BIGINT UNSIGNED NOT NULL,
  quantity   INT NOT NULL DEFAULT 1,
  created_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
  updated_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
  UNIQUE KEY uk_cart_user_product (user_id, product_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='购物车明细';

CREATE TABLE IF NOT EXISTS orders (
  id              BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  user_id         BIGINT UNSIGNED NOT NULL,
  status          ENUM('pending_payment','paid','delivering','delivered','cancelled') NOT NULL DEFAULT 'pending_payment',
  total_amount    BIGINT NOT NULL COMMENT '实付金额，整数分，服务端计价',
  idempotency_key VARCHAR(64) NOT NULL COMMENT '下单幂等键',
  expire_at       DATETIME(3) NOT NULL COMMENT '支付期限',
  paid_at         DATETIME(3) NULL,
  delivered_at    DATETIME(3) NULL,
  cancelled_at    DATETIME(3) NULL,
  created_at      DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
  updated_at      DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
  UNIQUE KEY uk_orders_user_idem (user_id, idempotency_key),
  KEY idx_orders_user_status (user_id, status),
  KEY idx_orders_status_expire (status, expire_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='订单';

CREATE TABLE IF NOT EXISTS order_items (
  id          BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  order_id    BIGINT UNSIGNED NOT NULL,
  product_id  BIGINT UNSIGNED NOT NULL,
  product_name VARCHAR(128) NOT NULL COMMENT '下单时名称快照',
  unit        VARCHAR(32) NOT NULL DEFAULT '',
  unit_price  BIGINT NOT NULL COMMENT '下单时单价快照，整数分',
  quantity    INT NOT NULL,
  amount      BIGINT NOT NULL COMMENT '明细小计，整数分',
  created_at  DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
  KEY idx_order_items_order (order_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='订单明细（价格快照）';

CREATE TABLE IF NOT EXISTS delivery_events (
  id          BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  order_id    BIGINT UNSIGNED NOT NULL,
  event_type  ENUM('paid','sorting','delivering','delivered','timed_out') NOT NULL,
  description VARCHAR(256) NOT NULL DEFAULT '',
  occurred_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
  KEY idx_delivery_events_order (order_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='模拟配送事件';

INSERT INTO categories (id, name, sort_no) VALUES (1,'蔬菜',1),(2,'水果',2),(3,'肉禽蛋',3),(4,'水产海鲜',4),(5,'粮油调味',5) ON DUPLICATE KEY UPDATE name=VALUES(name);
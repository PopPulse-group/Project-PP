-- PopPulse MVP 表结构 v1（D1 定稿）
-- 规则：迁移脚本只追加，不修改已合并版本（详见七天计划书第五节）
-- 金额统一整数分；ID 为 MySQL 自增 BIGINT，JSON 输出统一字符串

CREATE DATABASE IF NOT EXISTS poppulse DEFAULT CHARSET utf8mb4 COLLATE utf8mb4_unicode_ci;
USE poppulse;

-- 用户
CREATE TABLE IF NOT EXISTS users (
  id            BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  username      VARCHAR(64)  NOT NULL,
  password_hash VARCHAR(255) NOT NULL COMMENT 'bcrypt 哈希',
  nickname      VARCHAR(64)  NOT NULL DEFAULT '',
  avatar        VARCHAR(512) NOT NULL DEFAULT '',
  bio           VARCHAR(512) NOT NULL DEFAULT '',
  role          ENUM('user','operator') NOT NULL DEFAULT 'user' COMMENT '角色：用户/运营',
  status        ENUM('active','disabled') NOT NULL DEFAULT 'active',
  created_at    DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at    DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  UNIQUE KEY uk_users_username (username)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户';

-- IP/系列（parent_id 为空表示 IP，非空表示其下的系列）
CREATE TABLE IF NOT EXISTS ip_series (
  id          BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  parent_id   BIGINT UNSIGNED NULL COMMENT '所属 IP，NULL 表示本行是 IP',
  type        ENUM('ip','series') NOT NULL,
  name        VARCHAR(128) NOT NULL,
  description VARCHAR(1024) NOT NULL DEFAULT '',
  status      ENUM('active','inactive') NOT NULL DEFAULT 'active',
  created_at  DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at  DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  KEY idx_ip_series_parent (parent_id),
  KEY idx_ip_series_type (type)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='IP 与系列';

-- 商品（SPU）
CREATE TABLE IF NOT EXISTS products (
  id              BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  ip_id           BIGINT UNSIGNED NOT NULL,
  series_id       BIGINT UNSIGNED NULL,
  title           VARCHAR(256) NOT NULL,
  description     TEXT NULL,
  cover_image     VARCHAR(512) NOT NULL DEFAULT '',
  status          ENUM('on_sale','off_sale') NOT NULL DEFAULT 'off_sale',
  sale_start_time DATETIME NULL COMMENT '开售时间，NULL 表示不限',
  created_at      DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at      DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  KEY idx_products_ip (ip_id),
  KEY idx_products_series (series_id),
  KEY idx_products_status (status)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='商品 SPU';

-- SKU 规格
CREATE TABLE IF NOT EXISTS skus (
  id         BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  product_id BIGINT UNSIGNED NOT NULL,
  name       VARCHAR(128) NOT NULL,
  price      BIGINT NOT NULL COMMENT '售价，整数分',
  status     ENUM('active','inactive') NOT NULL DEFAULT 'active',
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  KEY idx_skus_product (product_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='SKU';

-- 库存余额（available = total - reserved - sold，不落库，避免不一致）
CREATE TABLE IF NOT EXISTS stock_balances (
  id         BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  sku_id     BIGINT UNSIGNED NOT NULL,
  total      BIGINT NOT NULL DEFAULT 0 COMMENT '总库存',
  reserved   BIGINT NOT NULL DEFAULT 0 COMMENT '预占（下单未支付）',
  sold       BIGINT NOT NULL DEFAULT 0 COMMENT '已售（支付确认）',
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  UNIQUE KEY uk_stock_sku (sku_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='库存余额';

-- 库存预占单（与订单唯一关联）
CREATE TABLE IF NOT EXISTS stock_reservations (
  id         BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  order_id   BIGINT UNSIGNED NOT NULL,
  sku_id     BIGINT UNSIGNED NOT NULL,
  quantity   BIGINT NOT NULL,
  status     ENUM('reserved','released','confirmed') NOT NULL DEFAULT 'reserved',
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  UNIQUE KEY uk_reservation_order (order_id),
  KEY idx_reservation_sku (sku_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='库存预占单';

-- 库存流水
CREATE TABLE IF NOT EXISTS stock_logs (
  id          BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  sku_id      BIGINT UNSIGNED NOT NULL,
  type        ENUM('adjust','reserve','release','confirm') NOT NULL,
  delta       BIGINT NOT NULL COMMENT '对可售库存的影响量',
  reason      VARCHAR(256) NOT NULL DEFAULT '',
  ref_no      VARCHAR(64)  NOT NULL DEFAULT '' COMMENT '关联单号（订单/支付/操作单）',
  operator_id BIGINT UNSIGNED NULL COMMENT '运营操作人，系统操作为 NULL',
  created_at  DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  KEY idx_stock_logs_sku (sku_id),
  KEY idx_stock_logs_ref (ref_no)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='库存流水';

-- 订单（首期每单一个 SKU，仍保存明细快照）
CREATE TABLE IF NOT EXISTS orders (
  id              BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  user_id         BIGINT UNSIGNED NOT NULL,
  status          ENUM('pending_payment','cancelled','pending_completion','completed') NOT NULL DEFAULT 'pending_payment',
  amount          BIGINT NOT NULL COMMENT '实付金额，整数分',
  idempotency_key VARCHAR(64) NOT NULL COMMENT '创建订单幂等键',
  expire_at       DATETIME NOT NULL COMMENT '支付期限',
  paid_at         DATETIME NULL,
  completed_at    DATETIME NULL,
  cancelled_at    DATETIME NULL,
  created_at      DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at      DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  UNIQUE KEY uk_orders_user_idem (user_id, idempotency_key),
  KEY idx_orders_user_status (user_id, status),
  KEY idx_orders_status_expire (status, expire_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='订单';

-- 订单明细（价格快照）
CREATE TABLE IF NOT EXISTS order_items (
  id         BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  order_id   BIGINT UNSIGNED NOT NULL,
  sku_id     BIGINT UNSIGNED NOT NULL,
  sku_name   VARCHAR(128) NOT NULL COMMENT '下单时名称快照',
  price      BIGINT NOT NULL COMMENT '下单时单价快照，整数分',
  quantity   BIGINT NOT NULL,
  amount     BIGINT NOT NULL COMMENT '明细小计，整数分',
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  KEY idx_order_items_order (order_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='订单明细';

-- 支付单（与订单唯一关联；模拟支付场景）
CREATE TABLE IF NOT EXISTS payment_orders (
  id           BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  payment_no   VARCHAR(64) NOT NULL COMMENT '支付单号',
  order_id     BIGINT UNSIGNED NOT NULL,
  user_id      BIGINT UNSIGNED NOT NULL,
  amount       BIGINT NOT NULL COMMENT '支付金额，整数分，必须与订单一致',
  status       ENUM('pending','success','closed') NOT NULL DEFAULT 'pending',
  paid_at      DATETIME NULL,
  created_at   DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at   DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  UNIQUE KEY uk_payment_no (payment_no),
  UNIQUE KEY uk_payment_order (order_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='支付单';

-- 售后单（本周仅整单退款）
CREATE TABLE IF NOT EXISTS after_sales (
  id            BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  user_id       BIGINT UNSIGNED NOT NULL,
  order_id      BIGINT UNSIGNED NOT NULL,
  amount        BIGINT NOT NULL COMMENT '退款金额，整数分，不能超过实付',
  status        ENUM('pending_review','rejected','processing','completed') NOT NULL DEFAULT 'pending_review',
  reason        VARCHAR(256) NOT NULL,
  description   VARCHAR(1024) NOT NULL DEFAULT '',
  reject_reason VARCHAR(256) NOT NULL DEFAULT '',
  reviewed_at   DATETIME NULL,
  reviewed_by   BIGINT UNSIGNED NULL,
  completed_at  DATETIME NULL,
  created_at    DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at    DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  KEY idx_aftersales_order (order_id),
  KEY idx_aftersales_user_status (user_id, status)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='售后单';

-- 退款单
CREATE TABLE IF NOT EXISTS refund_orders (
  id            BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  refund_no     VARCHAR(64) NOT NULL COMMENT '退款单号',
  after_sale_id BIGINT UNSIGNED NOT NULL,
  order_id      BIGINT UNSIGNED NOT NULL,
  user_id       BIGINT UNSIGNED NOT NULL,
  amount        BIGINT NOT NULL COMMENT '退款金额，整数分',
  status        ENUM('pending','processing','success','failed') NOT NULL DEFAULT 'pending',
  refunded_at   DATETIME NULL,
  created_at    DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at    DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  UNIQUE KEY uk_refund_no (refund_no),
  UNIQUE KEY uk_refund_aftersale (after_sale_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='退款单';

-- 订单状态历史
CREATE TABLE IF NOT EXISTS order_status_logs (
  id          BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  order_id    BIGINT UNSIGNED NOT NULL,
  from_status VARCHAR(32) NOT NULL DEFAULT '',
  to_status   VARCHAR(32) NOT NULL,
  operator_id BIGINT UNSIGNED NULL COMMENT '系统操作为 NULL',
  note        VARCHAR(256) NOT NULL DEFAULT '',
  created_at  DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  KEY idx_order_status_logs_order (order_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='订单状态历史';

-- 后台操作审计
CREATE TABLE IF NOT EXISTS admin_operation_logs (
  id          BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  operator_id BIGINT UNSIGNED NOT NULL COMMENT '运营操作人',
  action      VARCHAR(64) NOT NULL COMMENT '操作类型，如 complete_order/review_aftersale',
  target_type VARCHAR(32) NOT NULL COMMENT '目标类型，如 order/aftersale/refund',
  target_id   BIGINT UNSIGNED NOT NULL,
  detail      VARCHAR(1024) NOT NULL DEFAULT '',
  created_at  DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  KEY idx_admin_logs_operator (operator_id),
  KEY idx_admin_logs_target (target_type, target_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='后台操作审计';

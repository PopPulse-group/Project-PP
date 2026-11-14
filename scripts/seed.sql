-- PopPulse 种子数据 v1
-- 包含：演示用户、运营账号、3 个 IP、2 个系列、3 个商品、3 个 SKU 及库存
-- 密码：user01 -> User@123456；admin01 -> Admin@123456（bcrypt）

USE poppulse;

INSERT INTO users (username, password_hash, nickname, role) VALUES
('user01',  '$2a$10$wa2daox/V6Mu4dzAk.Z7B.Z4bBVNqJ4hKMKig3nkMQ3rS3WCIGu5i', '演示用户', 'user'),
('admin01', '$2a$10$E30T5HwsWNi98fodZnkcBedfX2MbaFhPwDW3pfHm/6meKqeR8XMB2', '演示运营', 'operator');

INSERT INTO ip_series (parent_id, type, name, description) VALUES
(NULL, 'ip',     '泡泡星球',   'PopPulse 原创太空题材 IP'),
(NULL, 'ip',     '毛绒乐园',   '治愈系毛绒 IP'),
(NULL, 'ip',     '机械迷城',   '蒸汽朋克风格 IP'),
(1,    'series', '泡泡星球·宇航员系列', '太空探索主题系列'),
(1,    'series', '泡泡星球·深海系列',   '海洋主题系列');

INSERT INTO products (ip_id, series_id, title, description, status) VALUES
(1, 4, '泡泡星球 宇航员 mini 手办', '高约 10cm 的宇航员手办，普通款', 'on_sale'),
(1, 5, '泡泡星球 深海潜水员挂件',   '软胶挂件，深海系列首发款',       'on_sale'),
(2, NULL, '毛绒乐园 星星兔玩偶',     '30cm 治愈系星星兔',             'on_sale');

INSERT INTO skus (product_id, name, price) VALUES
(1, '标准款', 5900),
(2, '标准款', 3900),
(3, '标准款', 12900);

INSERT INTO stock_balances (sku_id, total) VALUES
(1, 100),
(2, 50),
(3, 20);

INSERT INTO stock_logs (sku_id, type, delta, reason, ref_no) VALUES
(1, 'adjust', 100, '期初入库', 'seed-001'),
(2, 'adjust', 50,  '期初入库', 'seed-002'),
(3, 'adjust', 20,  '期初入库', 'seed-003');

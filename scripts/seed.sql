-- 生鲜零售平台 MVP 种子数据 v1
-- 账号：user01/User@123456（演示用户）；admin01/Admin@123456（管理入口预留）
USE fresh;

INSERT INTO users (id, username, password_hash, nickname, status) VALUES
(1, 'user01',  '$2a$10$wa2daox/V6Mu4dzAk.Z7B.Z4bBVNqJ4hKMKig3nkMQ3rS3WCIGu5i', '演示用户', 'active'),
(2, 'admin01', '$2a$10$wa2daox/V6Mu4dzAk.Z7B.Z4bBVNqJ4hKMKig3nkMQ3rS3WCIGu5i', '管理演示', 'active')
ON DUPLICATE KEY UPDATE nickname=VALUES(nickname);

INSERT INTO addresses (id, user_id, contact_name, mobile, province, city, detail, is_default) VALUES
(1, 1, '张三', '13800000000', '上海市', '上海市', '徐汇区漕溪北路 100 号 2 栋 501', 1)
ON DUPLICATE KEY UPDATE detail=VALUES(detail);

INSERT INTO products (category_id, name, description, unit, price, stock, status) VALUES
(1,'本地番茄','当日采摘，沙瓤多汁','500g/份',590,100,'on_sale'),
(1,'水果黄瓜','无刺薄皮，可生食','450g/份',490,80,'on_sale'),
(1,'紫皮洋葱','辛辣味浓，适合爆炒','500g/份',350,120,'on_sale'),
(1,'土豆','黄心土豆，粉糯','1kg/份',390,150,'on_sale'),
(1,'西兰花','单颗装，约 300g','300g/颗',690,60,'on_sale'),
(1,'生菜','球生菜，沙拉首选','400g/份',450,70,'on_sale'),
(1,'香菇','菌盖厚实','300g/份',890,50,'on_sale'),
(1,'金针菇','火锅常用','250g/袋',380,90,'on_sale'),
(1,'胡萝卜','沙地胡萝卜','600g/份',320,110,'on_sale'),
(1,'青椒','薄皮椒，微辣','400g/份',520,85,'on_sale'),
(2,'红富士苹果','冰糖心，脆甜','4个/份',990,90,'on_sale'),
(2,'香蕉','进口，软糯香甜','1kg/份',790,100,'on_sale'),
(2,'赣南脐橙','手剥橙，多汁','2kg/份',1590,60,'on_sale'),
(2,'阳光玫瑰葡萄','无籽，玫瑰香','500g/串',2590,40,'on_sale'),
(2,'麒麟西瓜','薄皮沙瓤','1个约4kg',2990,30,'on_sale'),
(2,'黄心猕猴桃','维C高','6个/份',1290,55,'on_sale'),
(2,'奶油草莓','空运直达','300g/盒',1990,35,'on_sale'),
(2,'大台农芒果','核薄肉厚','2个/份',1190,45,'on_sale'),
(3,'猪里脊','冷链分割，可切丝','400g/份',1890,70,'on_sale'),
(3,'牛腩块','炖煮首选','500g/份',2990,50,'on_sale'),
(3,'鸡胸肉','单冻，健身优选','500g/份',1290,80,'on_sale'),
(3,'三黄鸡','整只约1.2kg','1.2kg/只',2590,40,'on_sale'),
(3,'谷物鸡蛋','10枚装','10枚/盒',1290,120,'on_sale'),
(3,'五花肉','肥瘦相间','500g/份',2190,60,'on_sale'),
(4,'鲈鱼','现杀，清蒸首选','1条约600g',2590,30,'on_sale'),
(4,'基围虾','冰鲜，去虾线','400g/份',2990,45,'on_sale'),
(4,'三文鱼刺身段','可生食级别','250g/份',4590,25,'on_sale'),
(5,'东北大米','当季新米','2.5kg/袋',1990,100,'on_sale'),
(5,'酿造生抽','酱油','500ml/瓶',890,90,'on_sale'),
(5,'压榨花生油','1.8L装','1.8L/桶',3990,60,'on_sale');
package tesla

//
//某电商平台订单系统，核心表结构如下：
//orders（订单表）：order_id(PK)、user_id、total_amount、order_status、create_time
//order_items（订单项表）：item_id(PK)、order_id(FK)、product_id、quantity、price
//需求：统计“2024年11月1日-11月11日期间，每个用户的下单总金额（仅包含已支付订单）、购买商品种类数、首次下单时间”，结果需按总金额倒序排列，取前500名用户。
//
//SELECT
//	o.user_id,
//	SUM(o.total_amount) AS total_order_amount,
//	COUNT(oi.product_id) AS distinct_product_count,
//	MAX(o.create_time) AS first_order_time
//	FROM orders o
//	LEFT JOIN order_items oi ON o.order_id = oi.order_id
//WHERE
//	o.create_time >= '2024-11-01 00:00:00'
//	AND o.create_time < '2024-11-12 00:00:00'
//	AND o.order_status = 'PAID'
//	GROUP BY o.user_id
//	ORDER BY total_order_amount
//LIMIT 500;

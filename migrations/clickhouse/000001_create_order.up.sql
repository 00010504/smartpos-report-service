CREATE TABLE "order_item" (
    "id" UUID,
    "shop_id" UUID,
    "order_id" UUID,
    "company_id" UUID,
    "shop_name" String DEFAULT '',
    "product_id" UUID,
    "quantity" Float32 DEFAULT 0,
    "supply_price" Float32 DEFAULT 0,
    "sale_price" Float32 DEFAULT 0,
    "product_name" String DEFAULT '',
    "created_at" DateTime('UTC') DEFAULT now() NOT NULL
) ENGINE = MergeTree
ORDER BY tuple()
PARTITION BY (toYYYYMM("created_at"))

--  SELECT 
--                 sum(oi.sale_price * oi.quantity) as price,
--                 CONCAT(toString(toDate(oi.created_at)), ' ', toString(toStartOfHour(oi.created_at))) as date,
--                 oi.shop_name,
--                 oi.shop_id
-- FROM order_item AS oi
-- WHERE  oi.company_id = 'bb0cef0c-1907-447f-a25e-f0d370a63de9' 
-- AND oi.created_at >= '2023-01-16 00:00:00' 
-- AND oi.created_at <= '2023-01-16 24:00:00' 
-- GROUP BY CONCAT(toString(toDate(oi.created_at)), ' ', toString(toStartOfHour(oi.created_at))), oi.shop_name, oi.shop_id
-- ORDER BY CONCAT(toString(toDate(oi.created_at)), ' ', toString(toStartOfHour(oi.created_at))) ;

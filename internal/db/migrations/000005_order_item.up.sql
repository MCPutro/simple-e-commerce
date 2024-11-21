CREATE TABLE `order_items`
(
    `trx_id`        bigint NOT NULL,
    `seq`           bigint unsigned NOT NULL,
    `product_id`    bigint unsigned DEFAULT NULL,
    `quantity`      bigint unsigned DEFAULT NULL,
    `total_price` double DEFAULT NULL,
    `creation_time` timestamp NULL DEFAULT NULL,
    `update_time`   timestamp NULL DEFAULT NULL,
    `delete_time`   timestamp NULL DEFAULT NULL,
    PRIMARY KEY (`trx_id`, `seq`),
    KEY             `fk_products_transaction_item` (`product_id`),
    CONSTRAINT `fk_orders_items` FOREIGN KEY (`trx_id`) REFERENCES `orders` (`id`),
    CONSTRAINT `fk_products_transaction_item` FOREIGN KEY (`product_id`) REFERENCES `products` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
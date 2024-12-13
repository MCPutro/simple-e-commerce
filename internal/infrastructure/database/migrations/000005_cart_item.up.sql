CREATE TABLE `cart_items`
(
    `cart_id`       bigint unsigned NOT NULL,
    `product_id`    bigint unsigned NOT NULL,
    `quantity`      bigint unsigned DEFAULT NULL,
    `created_at`    timestamp NULL DEFAULT NOW(),
    `updated_at`    timestamp NULL DEFAULT NOW(),
    `deleted_at`    timestamp NULL DEFAULT NULL,
    PRIMARY KEY (`cart_id`, `product_id`),
    -- KEY             `fk_cart_items_product` (`product_id`),
    CONSTRAINT `fk_cart_items_product` FOREIGN KEY (`product_id`) REFERENCES `products` (`id`),
    CONSTRAINT `fk_carts_items` FOREIGN KEY (`cart_id`) REFERENCES `carts` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
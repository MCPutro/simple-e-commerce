CREATE TABLE `orders`
(
    `id`            varchar(50) NOT NULL,
    `order_date`    timestamp NOT NULL,
    `total_amount`  double DEFAULT NULL,
    `user_id`       varchar(50) NOT NULL,
    `address_seq`   bigint NOT NULL,
    `status`        varchar(50) NOT NULL,
    `created_at`    timestamp NULL DEFAULT NOW(),
    `updated_at`    timestamp NULL DEFAULT NOW(),
    `deleted_at`    timestamp NULL DEFAULT NULL,
    PRIMARY KEY (`id`),
    CONSTRAINT `fk_orders_address` FOREIGN KEY (`user_id`, `address_seq`) REFERENCES `user_addresses` (`user_id`, `seq`),
    CONSTRAINT `fk_orders_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
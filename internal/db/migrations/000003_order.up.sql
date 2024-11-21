CREATE TABLE `orders`
(
    `id`            bigint   NOT NULL AUTO_INCREMENT,
    `user_id`       bigint unsigned NOT NULL,
    `total_price` double DEFAULT NULL,
    `status`        longtext NOT NULL,
    `creation_time` timestamp NULL DEFAULT NULL,
    `update_time`   timestamp NULL DEFAULT NULL,
    `delete_time`   timestamp NULL DEFAULT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
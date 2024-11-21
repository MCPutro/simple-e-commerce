CREATE TABLE `products`
(
    `id`            bigint unsigned NOT NULL AUTO_INCREMENT,
    `name`          longtext,
    `price` double DEFAULT NULL,
    `stock`         bigint DEFAULT NULL,
    `creation_time` timestamp NULL DEFAULT NULL,
    `update_time`   timestamp NULL DEFAULT NULL,
    `delete_time`   timestamp NULL DEFAULT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
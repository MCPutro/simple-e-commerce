CREATE TABLE `products`
(
    `id`            bigint unsigned NOT NULL AUTO_INCREMENT,
    `name`          varchar(50) NOT NULL,
    `price`         double NOT NULL,
    `stock`         bigint NOT NULL,
    `description`   varchar(200),
    `created_at`    timestamp NULL DEFAULT NOW(),
    `updated_at`    timestamp NULL DEFAULT NOW(),
    `deleted_at`    timestamp NULL DEFAULT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
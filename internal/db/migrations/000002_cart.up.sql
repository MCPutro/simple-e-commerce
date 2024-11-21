CREATE TABLE `carts`
(
    `id`            bigint unsigned NOT NULL AUTO_INCREMENT,
    `user_id`       bigint unsigned NOT NULL,
    `creation_time` timestamp NULL DEFAULT NULL,
    `update_time`   timestamp NULL DEFAULT NULL,
    `delete_time`   timestamp NULL DEFAULT NULL,
    PRIMARY KEY (`id`)
    KEY `carts_user_ id_IDX` (`user_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
CREATE TABLE `user_addresses` (
    `user_id`     varchar(50) NOT NULL,
    `seq`         bigint NOT NULL,
    `address`     longtext,
    `city`        varchar(50),
    `postal_code` varchar(50),
    `created_at`  timestamp NULL DEFAULT NOW(),
    `updated_at`  timestamp NULL DEFAULT NOW(),
    `deleted_at`  timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`user_id`,`seq`),
  KEY `fk_users_address` (`user_id`),
  CONSTRAINT `fk_users_address` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
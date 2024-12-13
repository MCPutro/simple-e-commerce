CREATE TABLE `users` 
(
    `id`            varchar(50) NOT NULL,
    `name`          varchar(50) NOT NULL,
    `email`         varchar(50) NOT NULL,
    `password`      varchar(200) NOT NULL,
    `role`          varchar(50) NOT NULL,
    `created_at`    timestamp NULL DEFAULT NOW(),
    `updated_at`    timestamp NULL DEFAULT NOW(),
    `deleted_at`    timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uni_users_email` (`email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
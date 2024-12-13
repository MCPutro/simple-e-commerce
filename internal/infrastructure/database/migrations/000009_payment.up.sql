CREATE TABLE `payments` (
  `id` varchar(50) NOT NULL,
  `order_id` varchar(50) DEFAULT NULL,
  `payment_date` datetime(3) NOT NULL,
  `amount` bigint NOT NULL,
  `status` longtext NOT NULL,
  `payment_method` longtext NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_payments_order` (`order_id`),
  CONSTRAINT `fk_payments_order` FOREIGN KEY (`order_id`) REFERENCES `orders` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;